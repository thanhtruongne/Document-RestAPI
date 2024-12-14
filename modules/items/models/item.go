package models

import (
	"errors"

	"github.com/user/Practice_api/common"
	"github.com/user/Practice_api/modules/users/models"
)

var (
	ErrTitle          = errors.New("Tên không bỏ trống")
	ErrHasBeenDeleted = errors.New("Item đã bị xóa hoặc không đúng")
	ErrInvalidRequest = errors.New("Invalid request")
)

type Test struct {
	common.SQLmodel
	Name   string `json:"name" gorm:"column:name;"`
	Desc   string `json:"desc" gorm:"column:desc;"`
	Status int    `json:"status" gorm:"column:status;"`
	UserID int    `json:"user_id" column:"column:user_id;"`
	Owner  *Users `json:"owner" gorm:"foreignKey:UserID"`
}

type Users struct {
	// Id        int              `json:"id" gorm:"column:id;"`
	common.SQLmodel
	Email     string           `json:"email" gorm:"column:email;"`
	FirstName string           `json:"first_name" gorm:"column:first_name;"`
	LastName  string           `json:"last_name" gorm:"column:last_name;"`
	Status    int              `json:"status" gorm:"column:status;"`
	Role      *models.UserRole `json:"role" gorm:"column:role;"`
}

// genereate UID Test
func (test *Test) Mask() {
	test.SQLmodel.Mask(common.DBTypeItem)
	if v := test.Owner; v != nil {
		v.Mask()
	}
}

// genereate UID User
func (user *Users) Mask() {
	user.SQLmodel.Mask(common.DBTypeUser)
}

// tạo data table name cho struct
func (Test) TableName() string { return "test" }

// tạo struct để tạo data
type ItemCreate struct {
	Id      int    `json:"-" gorm:"column:id;"` // để rỗng k lấy vào
	User_id int    `json:"-" column:"column:user_id;"`
	Name    string `json:"name" gorm:"column:name;"`
	Desc    string `json:"desc"  gorm:"column:desc;" `
	Status  int    `json:"status"  gorm:"column:status;"`
}

// tạo func column name tạo mehtod reciever
func (ItemCreate) TableName() string { return Test{}.TableName() }

type ItemUpdate struct { // nên để pointer vào kiểu du liệu để xác định data và update
	Name   *string `json:"name" gorm:"column:name;" `
	Desc   *string `json:"desc"  gorm:"column:desc;"`
	Status *int    `json:"status"  gorm:"column:status;"`
}

func (ItemUpdate) TableName() string { return Test{}.TableName() }
