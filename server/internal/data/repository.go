package data

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
	rows, err := r.db.Query("SELECT * FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := []Post{}
	for rows.Next() {
		post := Post{}
		if err := rows.Scan(&post.ID, &post.Caption); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return &posts, nil
}
