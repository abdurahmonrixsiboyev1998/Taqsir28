package utlis

import (
	"encoding/json"
	"os"
	"taqsir/model"
)

func ReadFromFile(filename string) ([]model.User, error) {
	res, err := os.ReadFile(filename)
	if err != nil {
		return nil, os.ErrClosed
	}

	var user []model.User
	err = json.Unmarshal(res, &user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
