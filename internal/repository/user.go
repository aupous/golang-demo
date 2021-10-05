package repository

import (
	"awesomeProject/internal/db"
	"awesomeProject/internal/model"
)

type UserRepository struct {
	DB *db.DB
}

func NewUserRepository(db *db.DB) model.UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(user *model.User) error {
	_, err := r.DB.DB.Model(user).Returning("*").Insert()
	return err
}

func (r *UserRepository) Update(user *model.User) error {
	_, err := r.DB.DB.Model(user).WherePK().Update()
	return err
}

func (r *UserRepository) Find() ([]*model.User, error) {
	users := make([]*model.User, 0)
	err := r.DB.DB.Model(users).Select()
	return users, err
}

func (r *UserRepository) Delete(user *model.User) error {
	_, err := r.DB.DB.Model(user).WherePK().Delete()
	return err
}
