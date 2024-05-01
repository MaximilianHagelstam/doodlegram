package post

type Service interface {
	GetPosts() (*[]Post, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

// Get all posts
func (s *service) GetPosts() (*[]Post, error) {
	return s.repository.GetPosts()
}
