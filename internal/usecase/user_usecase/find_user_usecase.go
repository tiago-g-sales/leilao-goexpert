package user_usecase

import (
	"context"

	entity "github.com/tiago-g-sales/leilao-goexpert/internal/entity/user_entity"
	"github.com/tiago-g-sales/leilao-goexpert/internal/internal_error"
	"github.com/tiago-g-sales/leilao-goexpert/internal/model"
)

type UserUseCase struct {
	UserRepository entity.UserRepositoryInterface
}

type UserUseCaseInterface interface {
	
	FindUserById(ctx context.Context, userId string) (*model.UserOutputDTO, *internal_error.InternalError)

}

func (u *UserUseCase) FindUserById(ctx context.Context, userId string) (*model.UserOutputDTO, *internal_error.InternalError){

	userEntity, err :=  u.UserRepository.FindUserById(ctx, userId)
	if err != nil {
		return nil, err
	}
	return &model.UserOutputDTO{
		Id: userEntity.Id,
		Name: userEntity.Name,
	}, nil
}

