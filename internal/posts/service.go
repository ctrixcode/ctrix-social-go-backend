package posts

import (
	"fmt"
)

type Service interface {
	CreatePost(post *Post) error
	GetPostByID(id string) (*Post, error)
	GetPostByIDWithFields(id string, fields []string) (*Post, error)
	UpdatePost(post *Post) error
	DeletePost(id string) error
}

type service struct {
	repo PostRepository
}

func NewService(repo PostRepository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreatePost(post *Post) error {
	err := s.repo.CreatePost(post)
	if err != nil {
		return fmt.Errorf("failed to create post: %w", err)
	}
	return nil
}

func (s *service) GetPostByID(id string) (*Post, error) {
	post, err := s.repo.GetPostByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get post by ID: %w", err)
	}
	return post, nil
}

func (s *service) GetPostByIDWithFields(id string, fields []string) (*Post, error) {
	post, err := s.repo.GetPostByIDWithFields(id, fields)
	if err != nil {
		return nil, fmt.Errorf("failed to get post by ID with fields: %w", err)
	}
	return post, nil
}

func (s *service) UpdatePost(post *Post) error {
	err := s.repo.UpdatePost(post)
	if err != nil {
		return fmt.Errorf("failed to update post: %w", err)
	}
	return nil
}

func (s *service) DeletePost(id string) error {
	err := s.repo.DeletePost(id)
	if err != nil {
		return fmt.Errorf("failed to delete post: %w", err)
	}
	return nil
}
