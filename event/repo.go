package event

import "gorm.io/gorm"

type repo struct {
	db gorm.DB
}

func (r repo) Create(e Event) {

}
