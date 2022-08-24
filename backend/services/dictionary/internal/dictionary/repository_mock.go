package dictionary

import (
	"container/list"
	"context"

	models "github.com/rkfcccccc/english_words/shared_pkg/services/dictionary/models"
)

type repositoryMock struct {
	stringsQueue *list.List
	entriesQueue *list.List
	errQueue     *list.List
}

func NewRepositoryMock() *repositoryMock {
	return &repositoryMock{list.New(), list.New(), list.New()}
}

func (repo *repositoryMock) PushStringResponse(x string) {
	repo.stringsQueue.PushBack(x)
}

func (repo *repositoryMock) pullString() string {
	if repo.stringsQueue.Len() == 0 {
		panic("string response was not added")
	}

	return repo.stringsQueue.Remove(repo.stringsQueue.Front()).(string)
}

func (repo *repositoryMock) PushEntryResponse(x *models.WordEntry) {
	repo.entriesQueue.PushBack(x)
}

func (repo *repositoryMock) pullEntry() *models.WordEntry {
	if repo.entriesQueue.Len() == 0 {
		panic("string response was not added")
	}

	if x := repo.entriesQueue.Remove(repo.entriesQueue.Front()); x != nil {
		return x.(*models.WordEntry)
	}

	return nil
}

func (repo *repositoryMock) pullAllEntries() []*models.WordEntry {
	result := make([]*models.WordEntry, repo.entriesQueue.Len())

	for i := 0; i < len(result); i++ {
		result[i] = repo.entriesQueue.Front().Value.(*models.WordEntry)
		repo.entriesQueue.Remove(repo.entriesQueue.Front())
	}

	return nil
}

func (repo *repositoryMock) PushErrResponse(x error) {
	repo.errQueue.PushBack(x)
}

func (repo *repositoryMock) pullErr() error {
	if repo.errQueue.Len() == 0 {
		panic("string response was not added")
	}

	if x := repo.errQueue.Remove(repo.errQueue.Front()); x != nil {
		return x.(error)
	}

	return nil
}

func (repo *repositoryMock) CreateWordIndex(ctx context.Context) error {
	return repo.pullErr()
}

func (repo *repositoryMock) Create(ctx context.Context, entry *models.WordEntry) (string, error) {
	return repo.pullString(), repo.pullErr()
}

func (repo *repositoryMock) GetById(ctx context.Context, wordId string) (*models.WordEntry, error) {
	return repo.pullEntry(), repo.pullErr()
}

func (repo *repositoryMock) GetByWord(ctx context.Context, word string) (*models.WordEntry, error) {
	return repo.pullEntry(), repo.pullErr()
}

func (repo *repositoryMock) Delete(ctx context.Context, wordId string) error {
	return repo.pullErr()
}

func (repo *repositoryMock) SetPictures(ctx context.Context, wordId string, pictures []models.SourcedPicture) error {
	return repo.pullErr()
}

func (repo *repositoryMock) Search(ctx context.Context, query string) ([]*models.WordEntry, error) {
	return repo.pullAllEntries(), repo.pullErr()
}
