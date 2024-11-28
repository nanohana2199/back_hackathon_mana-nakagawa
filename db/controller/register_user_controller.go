package controller

import (
	"db/usecase"
	"encoding/json"
	"log"
	"net/http"
)

type RegisterUserController struct {
	UseCase *usecase.RegisterUserUseCase
}

func (c *RegisterUserController) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var input struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := c.UseCase.Execute(input.Name, input.Age)
	if err != nil {
		log.Printf("Error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(id))
}
