package user

import "gorm.io/gorm"

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) Store {
	return Store{db: db}
}

func (s Store) Create() error {
	return nil
}
