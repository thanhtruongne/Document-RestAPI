package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/user/Practice_api/common"
)

const EnityModel = "User"

type UserRole int // set dạng Role theo int

const (
	RoleUser  UserRole = 1 << iota // scale lên double số ---> 1
	RoleAdmin                      // ----> 2
	RoleMinur                      // ----> 4
)

func (role UserRole) String() string {
	switch role {
	case RoleUser:
		return "user"
	case RoleAdmin:
		return "admin"
	case RoleMinur:
		return "minur"
	default:
		return "guest"
	}

}

func (role *UserRole) Scan(value interface{}) error { // dũ liệu từ db
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprintf("Failed to Unmarshal JSON value:", value))
	}

	var r UserRole

	roleValue := string(bytes)
	if roleValue == "user" {
		r = RoleUser
	} else if roleValue == "admin" {
		r = RoleAdmin
	} else if roleValue == "adminur" {
		r = RoleMinur
	}

	*role = r
	return nil
}

func (role *UserRole) Value() (driver.Value, error) {
	if role == nil {
		return nil, nil
	}
	return role.String(), nil
}

func (role *UserRole) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", role.String())), nil
}

type User struct {
	Id        int      `json:"id" gorm:"column:id;"`
	Email     string   `json:"email" gorm:"column:email;"`
	Salt      string   `json:"-" gorm:"column:salt;"`
	Password  string   `json:"-" gorm:"column:password;"`
	FirstName string   `json:"first_name" gorm:"column:first_name;"`
	LastName  string   `json:"last_name" gorm:"column:last_name;"`
	Status    int      `json:"status" gorm:"column:status;"`
	Role      UserRole `json:"role" gorm:"column:role;"`
}

func (user *User) UserId() int {
	return user.Id
}

func (user *User) GetEmail() string {
	return user.Email
}
func (r *User) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Role)
}

func (r *User) GetRole() string {
	return r.Role.String()
}

func (user *User) GetFullName() string {
	return fmt.Sprintf("%s %s", user.FirstName, user.LastName)
}

func (User) TableName() string {
	return "users"
}

type UserCreate struct {
	Id        int      `json:"-" gorm:"column:id;"` // để rỗng k lấy vào
	Email     string   `json:"email" gorm:"column:email;"`
	Password  string   `json:"password" gorm:"column:password;"`
	FirstName string   `json:"first_name" gorm:"column:first_name;"`
	LastName  string   `json:"last_name" gorm:"column:last_name;"`
	Status    int      `json:"status" gorm:"column:status;"`
	Role      UserRole `json:"-" gorm:"column:role;"`
	Salt      string   `json:"-" gorm:"column:salt;"`
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

type UserLogin struct {
	Email    string `json:"email" form:"email" gorm:"column:email;"`
	Password string `json:"password" form:"password" gorm:"column:password"`
}

func (UserLogin) TableName() string {
	return User{}.TableName()
}

var (
	ErrEmailOrPasswordInvalid = common.NewErrorCustomResponse(
		errors.New("Email or Password invalid"),
		"email or password is invalid",
		"ErrUserNameOrPasswordInvalid",
	)

	ErrEmailExisted = common.NewErrorCustomResponse(
		errors.New("Email has already exists"),
		"Email has already exists",
		"ErrEmailIsExists",
	)
)
