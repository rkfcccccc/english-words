package verification

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/rkfcccccc/english_words/verification/pkg/dsync"
	"github.com/rkfcccccc/english_words/verification/pkg/mail"
)

const (
	MAX_REQUESTS_PER_TYPE = 3
	MAX_ATTEMPTS          = 3
	CODE_LENGTH           = 4 // keep in mind to change database field size when changing this
	CODE_TTL              = time.Minute * 5
)

var ErrTooManyRequests = errors.New("user has too many verification requests")
var ErrNotFound = errors.New("verification request was not found")
var ErrNoAttemptsLeft = errors.New("no attempts left")

type Service struct {
	repo Repository
	sync dsync.Client
	mail *mail.Client
}

func NewService(repo Repository, sync dsync.Client, mail *mail.Client) *Service {
	return &Service{repo, sync, mail}
}

func generateRandomCode() int {
	max := 1
	for i := 0; i < CODE_LENGTH; i++ {
		max *= 10
	}

	return rand.Intn(max)
}

/* Deletes expired verification codes with 10% chance */
func (service *Service) deleteAllExpired(ctx context.Context) error {
	if rand.Intn(100) >= 10 {
		return nil
	}

	log.Println("deleted expired")
	return service.repo.DeleteAllExpired(ctx)
}

func (service *Service) sendMail(ctx context.Context, email string, typeId, code int) error {
	paddedCode := fmt.Sprintf("%0*d", CODE_LENGTH, code)
	subject := fmt.Sprintf("%s - your verification code", paddedCode)
	content := fmt.Sprintf("Use this code to verify your email: %s. It will expire in %d minutes.", paddedCode, int(CODE_TTL/time.Minute))
	return service.mail.Send(email, subject, content)
}

func (service *Service) SendCode(ctx context.Context, email string, typeId int) (string, error) {
	if err := service.deleteAllExpired(ctx); err != nil {
		return "", fmt.Errorf("service.deleteAllExpired: %v", err)
	}

	mutex := service.sync.NewMutex(fmt.Sprintf("verification_send:%s", email))
	if err := mutex.Lock(); err != nil {
		return "", fmt.Errorf("mutex.Lock: %v", err)
	}

	defer mutex.Unlock()

	entries, err := service.repo.GetByEmail(ctx, email, typeId)
	if err != nil {
		return "", fmt.Errorf("repo.GetByEmail: %v", err)
	}

	if len(entries) >= MAX_REQUESTS_PER_TYPE {
		return "", ErrTooManyRequests
	}

	code := generateRandomCode()
	requestId, err := service.repo.Create(ctx, email, typeId, code, MAX_ATTEMPTS, CODE_TTL)

	if err != nil {
		return "", fmt.Errorf("repo.Create: %v", err)
	}

	if err := service.sendMail(ctx, email, typeId, code); err != nil {
		return "", fmt.Errorf("service.sendMail: %v", err)
	}

	return requestId, nil
}

func (service *Service) Verify(ctx context.Context, requestId string, code int) (bool, error) {
	if err := service.deleteAllExpired(ctx); err != nil {
		return false, fmt.Errorf("service.deleteAllExpired: %v", err)
	}

	mutex := service.sync.NewMutex(fmt.Sprintf("verification_verify:%s", requestId))
	if err := mutex.Lock(); err != nil {
		return false, fmt.Errorf("mutex.Lock: %v", err)
	}

	defer mutex.Unlock()

	entry, err := service.repo.GetById(ctx, requestId)

	if err != nil {
		return false, fmt.Errorf("repo.GetById: %v", err)
	}

	if entry == nil {
		return false, ErrNotFound
	}

	if entry.Attempts == 0 {
		return false, ErrNoAttemptsLeft
	}

	if entry.Code != code {
		if err := service.repo.SetAttempts(ctx, requestId, entry.Attempts-1); err != nil {
			return false, fmt.Errorf("repo.SetAttempts: %v", err)
		}

		return false, nil
	}

	if err := service.repo.Delete(ctx, requestId); err != nil {
		return false, fmt.Errorf("repo.Delete: %v", err)
	}

	return true, nil
}
