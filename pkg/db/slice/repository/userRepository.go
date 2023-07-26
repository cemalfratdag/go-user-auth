package repository

//
//import (
//	"cfd/myapp/internal/common/errors"
//	"cfd/myapp/internal/core/domain/entity"
//)
//
//type UserRepository struct {
//	users     []entity.User
//	lastIndex int
//}
//
//func NewUserRepository() *UserRepository {
//	users := []entity.User{
//		{ID: 1, Name: "user1", Email: "user1@mail.com", Password: "Password1", Birthdate: "1978-01-01", Age: 45, Fee: 150},
//		{ID: 2, Name: "user2", Email: "user2@mail.com", Password: "Password1", Birthdate: "1993-01-01", Age: 30, Fee: 100},
//	}
//	return &UserRepository{
//		users:     users,
//		lastIndex: 2,
//	}
//}
//
//func (r *UserRepository) GetUserByID(ID int) (*entity.User, error) {
//	for _, user := range r.users {
//		if user.ID == ID {
//			foundUser := user
//			return &foundUser, nil
//		}
//	}
//	return nil, errors.ErrGetByID
//}
//
//func (r *UserRepository) GetUserByEmail(email string) (*entity.User, error) {
//	for _, user := range r.users {
//		if user.Email == email {
//			foundUser := user
//			return &foundUser, nil
//		}
//	}
//	return nil, errors.ErrGetByEmail
//}
//
//func (r *UserRepository) CreateUser(user entity.User) (int, error) {
//	r.lastIndex++
//	user.ID = r.lastIndex
//	r.users = append(r.users, user)
//	return user.ID, nil
//}
//
//func (r *UserRepository) UpdateUser(userID int, newUser entity.User) error {
//	for i, user := range r.users {
//		if user.ID == userID {
//			existingUser := &r.users[i]
//			existingUser.Name = newUser.Name
//			existingUser.Email = newUser.Email
//			existingUser.Password = newUser.Password
//			existingUser.Birthdate = newUser.Birthdate
//			existingUser.Age = newUser.Age
//			existingUser.Fee = newUser.Fee
//			return nil
//		}
//	}
//	return errors.ErrUpdate
//}
//
//func (r *UserRepository) DeleteUser(ID int) error {
//	for i, u := range r.users {
//		if u.ID == ID {
//			r.users = append(r.users[:i], r.users[i+1:]...)
//		}
//	}
//	return nil
//}
