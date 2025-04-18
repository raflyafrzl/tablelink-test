package usecase

import "test_tablelink/repository"

type UseCase struct {
	repo *repository.Repository
}

func NewUsecase(repo *repository.Repository) *UseCase {
	return &UseCase{repo: repo}
}
