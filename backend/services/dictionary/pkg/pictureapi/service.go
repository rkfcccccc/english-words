package pictureapi

import (
	"context"
	"log"
	"sync"
)

type Service struct {
	repositories []Repository
}

func NewService(repositories ...Repository) *Service {
	return &Service{repositories}
}

func (service *Service) Search(ctx context.Context, query string) []Picture {
	pictures := []Picture{}

	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}

	wg.Add(len(service.repositories))
	for _, repo := range service.repositories {
		go func(repo Repository) {
			defer wg.Done()

			result, err := repo.Search(ctx, query)
			if err != nil {
				log.Printf("%s repository got error when searching for %q: %v", repo.GetName(), query, err)
				return
			}

			mutex.Lock()
			pictures = append(pictures, result...)
			mutex.Unlock()
		}(repo)
	}

	wg.Wait()
	return pictures
}
