package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id        uint   `Json:"id"`
	FirstName string `Json:"first_name"`
	LastName  string `Json:"last_name"`
	Password  []byte `Json:"-"`
	Phone     string `Json:"phone"`
	Email     string `Json:"email"`
}

func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = hashedPassword
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
