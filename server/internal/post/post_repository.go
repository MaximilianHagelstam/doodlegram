package post

import "database/sql"

type Repository interface {
	GetPosts() (*[]Post, error)
}

type repository struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetPosts() (*[]Post, error) {
	posts := &[]Post{
		{ID: "acb123", Caption: "Test"},
	}
	return posts, nil
}
