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

func (r *UserRepository) Search(search string) ([]*model.User, error) {
	users := make([]*model.User, 0)
	query := r.DB.DB.Model(&users)
	if search != "" {
		searchString := "%" + search + "%"
		query = query.Where("name ILIKE ? OR email ILIKE ?", searchString, searchString)
	}
	err := query.Select()
	return users, err
}

func (r *UserRepository) Find(req model.FindUserRequest) ([]*model.User, int64, error) {
	users := make([]*model.User, 0)
	query := r.DB.DB.Model(&users)
	//
	if req.Search != "" {
		searchString := "%" + req.Search + "%"
		query = query.Where("name ILIKE ? OR email ILIKE ?", searchString, searchString)
	}
	if req.MinAge != 0 {
		query = query.Where("age >= ?", req.MinAge)
	}
	if req.MaxAge != 0 {
		query = query.Where("age <= ?", req.MaxAge)
	}
	if req.Job != "" {
		query = query.Where("job = ?", req.Job)
	}

	total, err := query.Offset(req.PerPage * (req.Page - 1)).Limit(req.PerPage).SelectAndCount()
	return users, int64(total), err
}

func (r *UserRepository) Delete(user *model.User) error {
	_, err := r.DB.DB.Model(user).WherePK().Delete()
	return err
}
