package common

const CurrentUser = "current_user"

// jwt
type TokenPayload struct {
	UId   int    `json:"user_id"`
	URole string `json:"role"`
}

func (t TokenPayload) UserId() int {
	return t.UId
}

func (t TokenPayload) Role() string {
	return t.URole
}

// tạo requester
type Requester interface {
	UserId() int
	GetEmail() string
	GetRole() string
}

func IsAdmin(requester Requester) bool {
	return requester.GetRole() == "admin"
}

type DBType int

const (
	DBTypeItem DBType = 1
	DBTypeUser DBType = 2
)
