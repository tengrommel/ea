package services

import "ea/bookstore_users_api/domain/users"

func CreateUser(user users.User) (*users.User, error) {
	return &user, nil
}
