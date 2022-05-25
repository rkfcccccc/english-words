package pictureapi

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/rkfcccccc/english_words/shared_pkg/cache"
)

const cacheTTL = time.Hour

type Service struct {
	repositories []Repository
	cache        cache.Repository
}

func NewService(repositories []Repository, cache cache.Repository) *Service {
	return &Service{repositories, cache}
}

func (service *Service) Search(ctx context.Context, query string) []Picture {
	pictures := []Picture{}

	cacheKey := fmt.Sprintf("pictures-%q", query)
	if service.cache.Get(ctx, cacheKey, pictures) == nil {
		return pictures
	}

	errors := false
	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}

	wg.Add(len(service.repositories))
	for _, repo := range service.repositories {
		go func(repo Repository) {
			defer wg.Done()

			result, err := repo.Search(ctx, query)
			if err != nil {
				errors = true
				log.Printf("%s repository got error when searching for %q: %v", repo.GetName(), query, err)
				return
			}

			mutex.Lock()
			pictures = append(pictures, result...)
			mutex.Unlock()
		}(repo)
	}

	wg.Wait()

	if !errors {
		if err := service.cache.Set(ctx, cacheKey, pictures, cacheTTL); err != nil {
			log.Printf("got error when saving pictures to cache: %v", err)
		}
	}

	return pictures
}
