package blog_tag

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) IRepo {
	return Repository{db: db}
}