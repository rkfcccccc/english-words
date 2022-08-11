package user

import (
	"context"
	"errors"
	"fmt"
	"regexp"

	"github.com/rkfcccccc/english_words/shared_pkg/dsync"
	"golang.org/x/crypto/bcrypt"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

var ErrAlreadyExists = errors.New("email already in use")
var ErrInvalidEmail = errors.New("invalid email")
var ErrInvalidPassword = errors.New("invalid password")
var ErrNotFound = errors.New("user was not found")

type Service struct {
	repo Repository
	sync dsync.Client
}

func NewService(repo Repository, sync dsync.Client) *Service {
	return &Service{repo, sync}
}

func (service *Service) checkEmail(ctx context.Context, email string) error {
	if emailRegex.FindString(email) == "" {
		return ErrInvalidEmail
	}

	if len(email) > 64 {
		return ErrInvalidEmail
	}

	user, err := service.repo.GetByEmail(ctx, email)
	if err != nil {
		return fmt.Errorf("service.GetByEmail: %v", err)
	}

	if user != nil {
		return ErrAlreadyExists
	}

	return nil
}

func (service *Service) checkPassword(_ context.Context, password string) error {
	if len(password) > 72 || len(password) < 6 {
		return ErrInvalidPassword
	}

	return nil
}

func (service *Service) hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("bcrypt.GenerateFromPassword: %v", err)
	}

	return string(hash), err
}

func (service *Service) CanCreate(ctx context.Context, email, password string) (bool, error) {
	if err := service.checkPassword(ctx, password); err != nil {
		return false, err
	}

	if err := service.checkEmail(ctx, email); err != nil {
		return false, err
	}

	return true, nil
}

func (service *Service) UpdatePassword(ctx context.Context, userId int, password string) error {
	if err := service.checkPassword(ctx, password); err != nil {
		return err
	}

	hash, err := service.hashPassword(password)
	if err != nil {
		return fmt.Errorf("bcrypt.GenerateFromPassword: %v", err)
	}

	if err := service.repo.UpdatePassword(ctx, userId, hash); err != nil {
		return fmt.Errorf("repo.UpdatePassword: %v", err)
	}

	return nil
}

func (service *Service) Create(ctx context.Context, email, password string) (int, error) {
	mutex := service.sync.NewMutex(fmt.Sprintf("user_%s", email))
	if err := mutex.Lock(); err != nil {
		return -1, fmt.Errorf("mutex.Lock: %v", err)
	}

	defer mutex.Unlock()

	ok, err := service.CanCreate(ctx, email, password)
	if err != nil {
		return -1, err
	} else if !ok {
		return -1, errors.New("cannot create an account")
	}

	hash, err := service.hashPassword(password)
	if err != nil {
		return -1, fmt.Errorf("bcrypt.GenerateFromPassword: %v", err)
	}

	return service.repo.Create(ctx, email, string(hash))
}

func (service *Service) GetById(ctx context.Context, userId int) (*User, error) {
	user, err := service.repo.GetById(ctx, userId)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, ErrNotFound
	}

	return user, nil
}

func (service *Service) GetByEmail(ctx context.Context, email string) (*User, error) {
	user, err := service.repo.GetByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, ErrNotFound
	}

	return user, nil
}

func (service *Service) GetByEmailAndPassword(ctx context.Context, email, password string) (*User, error) {
	user, err := service.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, fmt.Errorf("bcrypt.CompareHashAndPassword: %v", err)
	}

	return user, nil
}

func (service *Service) Delete(ctx context.Context, userId int) error {
	return service.repo.Delete(ctx, userId)
}
