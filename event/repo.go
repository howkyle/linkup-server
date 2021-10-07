package event

import (
	"errors"

	"gorm.io/gorm"
)

var CreationError = errors.New("failed to add event to db")

type repo struct {
	db *gorm.DB
}

func (r repo) Create(e Event) (uint, error) {

	res := r.db.Create(&e)
	if res.Error != nil {
		return 0, CreationError
	}
	return e.ID, nil
}

func (r repo) Retrieve(e interface{}) (Event, error) {
	return Event{}, nil
}

func (r repo) Update(e Event) error {
	return nil
}

func (r repo) Delete(e interface{}) error {
	return nil
}

func NewRepository(db *gorm.DB) Repository {
	return repo{db: db}
}
