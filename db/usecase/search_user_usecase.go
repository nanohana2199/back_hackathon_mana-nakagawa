package usecase

import (
	"db/dao"
	"db/model"
)

type SearchUserUseCase struct {
	DAO *dao.UserDAO
}

func (uc *SearchUserUseCase) Execute(name string) ([]model.User, error) {
	return uc.DAO.FindUserByName(name)
}
