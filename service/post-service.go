package service

import "alcohol/entity"

type PostService interface {
	Save(entity.Post) entity.Post
	FindAll() []entity.Post
}

type postService struct {
	posts []entity.Post
}

func New() PostService {
	return &postService{}
}

func (service *postService) Save(post entity.Post) entity.Post {
	service.posts = append(service.posts, post)
	return post
}

func (service *postService) FindAll() []entity.Post {
	return service.posts
}
