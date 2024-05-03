package tests

import "doodlegram/internal/data"

type repository struct {
}

func NewMockRepo() data.Repository {
	return &repository{}
}

func (r *repository) GetPosts() (*[]data.Post, error) {
	posts := &[]data.Post{
		{ID: "acb123", Caption: "Test"},
	}
	return posts, nil
}
