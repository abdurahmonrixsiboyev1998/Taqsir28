package handler

import (
	"taqsir/db"
	"taqsir/model"
)

type UserHandler struct{}

func (h *UserHandler) CreateUser(user model.User, reply *string) error {
	if err := db.CreateUser(user); err != nil {
		return err
	}
	*reply = "User created successfully"
	return nil
}

func (h *UserHandler) GetUser(id string, reply *model.User) error {
	user, err := db.GetUser(id)
	if err != nil {
		return err
	}
	*reply = user
	return nil
}

func (h *UserHandler) UpdateUser(user model.User, reply *string) error {
	if err := db.UpdateUser(user); err != nil {
		return err
	}
	*reply = "User updated successfully"
	return nil
}

func (h *UserHandler) DeleteUser(id string, reply *string) error {
	if err := db.DeleteUser(id); err != nil {
		return err
	}
	*reply = "User deleted successfully"
	return nil
}

func (h *UserHandler) ListUsers(dummy string, reply *[]model.User) error {
	users, err := db.ListUsers()
	if err != nil {
		return err
	}
	*reply = users
	return nil
}
