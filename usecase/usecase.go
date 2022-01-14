package usecase

import (
	"context"
	"go-context/repository"
	"log"
)

type Usecase struct {
	repo repository.Repository
}

func NewUsecase(repo repository.Repository) Usecase {
	return Usecase{
		repo: repo,
	}
}

func (u *Usecase) PrintUser(ctx context.Context) {
	users, err := u.repo.GetUser(ctx)
	if err != nil {
		log.Println(err)
	}
	for _, u := range users {
		log.Println("name: ", u.Name, " email:", u.Email)
	}
}
