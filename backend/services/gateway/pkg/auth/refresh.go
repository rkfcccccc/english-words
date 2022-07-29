package auth

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"

	"github.com/rkfcccccc/english_words/shared_pkg/cache"
)

// TODO: somehow rename all these functions because they sound shitty

func generateToken32() string {
	b := make([]byte, 32)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func (helper *Helper) GetRefreshToken(ctx context.Context, token string) (int, error) {
	var user_id int
	err := helper.cache.Get(ctx, fmt.Sprintf("refresh_%s", token), &user_id)

	if errors.Is(err, cache.ErrCacheMiss) {
		return -1, ErrRefreshMiss
	}

	if err != nil {
		return -1, fmt.Errorf("helper.cache.Get: %v", err)
	}

	return user_id, nil
}

func (helper *Helper) DeleteRefreshToken(ctx context.Context, token string) error {
	err := helper.cache.Del(ctx, fmt.Sprintf("refresh_%s", token))

	if err != nil {
		return fmt.Errorf("helper.cache.Del: %v", err)
	}

	return nil
}

func (helper *Helper) generateFreeToken(ctx context.Context) (string, error) {
	token := generateToken32()

	for {
		_, err := helper.GetRefreshToken(ctx, token)

		if errors.Is(err, ErrRefreshMiss) {
			break
		}

		if err != nil {
			return "", fmt.Errorf("helper.GetRefreshToken: %v", err)
		}

		token = generateToken32()
	}

	return token, nil
}

func (helper *Helper) IssueRefreshToken(ctx context.Context, user_id int) (string, error) {
	token, err := helper.generateFreeToken(ctx)

	if err != nil {
		return "", fmt.Errorf("helper.generateFreeToken: %v", err)
	}

	if err := helper.cache.Set(ctx, fmt.Sprintf("refresh_%s", token), user_id, refreshTTL); err != nil {
		return "", fmt.Errorf("helper.cache.Set: %v", err)
	}

	return token, nil
}
