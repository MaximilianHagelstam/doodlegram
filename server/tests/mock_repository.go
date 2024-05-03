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

func (r *repository) GetPostByID(id string) (*data.Post, error) {
	post := &data.Post{
		ID: "acb123", Caption: "Test",
	}
	return post, nil
}

func (r *repository) DeletePost(id string) error {
	return nil
}
