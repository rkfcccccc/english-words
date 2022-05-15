package user

import (
	"context"
	"errors"
	"fmt"
	"regexp"

	"github.com/rkfcccccc/english_words/user/pkg/dsync"
	"golang.org/x/crypto/bcrypt"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

var ErrEmailAlreadyUsed = errors.New("email already used")
var ErrInvalidEmail = errors.New("invalid email")
var ErrTooLongPassword = errors.New("too long password")
var ErrTooLongEmail = errors.New("too long email")

type Service struct {
	repo Repository
	sync dsync.Client
}

func NewService(repo Repository, sync dsync.Client) *Service {
	return &Service{repo, sync}
}

func (service *Service) Create(ctx context.Context, email, password string) (int, error) {
	if emailRegex.FindString(email) == "" {
		return -1, ErrInvalidEmail
	}

	if len(password) > 72 {
		return -1, ErrTooLongPassword
	}

	if len(email) > 64 {
		return -1, ErrTooLongEmail
	}

	mutex := service.sync.NewMutex(fmt.Sprintf("user_%s", email))
	if err := mutex.Lock(); err != nil {
		return -1, fmt.Errorf("mutex.Lock: %v", err)
	}

	defer mutex.Unlock()

	user, err := service.GetByEmail(ctx, email)
	if err != nil {
		return -1, fmt.Errorf("service.GetByEmail: %v", err)
	}

	if user != nil {
		return -1, ErrEmailAlreadyUsed
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return -1, fmt.Errorf("bcrypt.GenerateFromPassword: %v", err)
	}

	return service.repo.Create(ctx, email, string(hash))
}

func (service *Service) GetById(ctx context.Context, userId int) (*User, error) {
	return service.repo.GetById(ctx, userId)
}

func (service *Service) GetByEmail(ctx context.Context, email string) (*User, error) {
	return service.repo.GetByEmail(ctx, email)
}

func (service *Service) Delete(ctx context.Context, userId int) error {
	return service.repo.Delete(ctx, userId)
}
