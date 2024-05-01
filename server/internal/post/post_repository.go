package post

type Repository interface {
	GetPosts() (*[]Post, error)
}

type repository struct {
}

func NewRepo() Repository {
	return &repository{}
}

func (r *repository) GetPosts() (*[]Post, error) {
	posts := &[]Post{
		{ID: "acb123", Caption: "Test"},
	}
	return posts, nil
}
