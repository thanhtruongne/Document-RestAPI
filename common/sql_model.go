package common

type SQLmodel struct {
	Id     int  `json:"-" gorm:"column:id;"`
	FakeID *UID `json:"id" gorm:"-"`
	// Created_at *time.Time `json:"created_at" gorm:"column:created_at;"`
	// Updated_at *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

func (s *SQLmodel) Mask(dbType DBType) {
	uid := NewUID(uint32(s.Id), int(dbType), 1)
	s.FakeID = &uid
}
