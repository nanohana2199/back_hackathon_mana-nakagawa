package usecase

import (
	"crypto/rand"
	"db/dao"
	"db/model"
	"errors"
	"log"
	"time"

	"github.com/oklog/ulid"
)

type RegisterUserUseCase struct {
	DAO *dao.UserDAO
}

func generateULID() string {
	t := time.Now().UTC()
	entropy := ulid.Monotonic(rand.Reader, 0)
	id, err := ulid.New(ulid.Timestamp(t), entropy)
	if err != nil {
		log.Fatalf("failed to generate ULID: %v", err)
	}
	return id.String()
}

func (uc *RegisterUserUseCase) Execute(name string, age int) (string, error) {
	if name == "" || len(name) > 50 {
		return "", errors.New("name is either empty or exceeds 50 characters")
	}
	if age < 20 || age > 80 {
		return "", errors.New("age is out of the allowed range (20-80)")
	}

	id := generateULID()
	user := model.User{ID: id, Name: name, Age: age}
	if err := uc.DAO.InsertUser(user); err != nil {
		return "", err
	}
	return id, nil
}
