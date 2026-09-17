package service

import "shelter-platform/internal/shelter/domain"

type Service struct {
	repo domain.Repository
}

func New(repo domain.Repository) *Service {
	return &Service{
		repo: repo,
	}
}
