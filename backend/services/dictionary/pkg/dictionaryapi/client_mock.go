package dictionaryapi

import (
	"container/list"
	"context"
)

type clientMock struct {
	entriesQueue *list.List
	errQueue     *list.List
}

func NewMockClient() *clientMock {
	return &clientMock{list.New(), list.New()}
}

func (c *clientMock) AddResponse(entry *Entry, err error) {
	c.entriesQueue.PushBack(entry)
	c.errQueue.PushBack(err)
}

func (c *clientMock) GetWordEntry(ctx context.Context, language, word string) (*Entry, error) {
	if c.entriesQueue.Len() == 0 || c.errQueue.Len() == 0 {
		panic("no mocked response for this call")
	}

	_entry := c.entriesQueue.Remove(c.entriesQueue.Front())
	_err := c.errQueue.Remove(c.errQueue.Front())

	var entry *Entry
	var err error

	if _entry != nil {
		entry = _entry.(*Entry)
	}

	if _err != nil {
		err = _err.(error)
	}

	return entry, err
}
