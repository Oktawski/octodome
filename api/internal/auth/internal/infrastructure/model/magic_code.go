package model

type MagicCode struct {
	ID    uint   `gorm:"primaryKey"`
	Code  string `gorm:"not null;size:6"`
	Email string `gorm:"not null"`
	// TODO: add expires_at or created_at
	// TODO: add base entity with ID, CreatedAt, UpdatedAt, DeletedAt, DeletedById
}

func (MagicCode) TableName() string {
	return "magic_codes"
}

func NewMagicCode(code string, email string) *MagicCode {
	return &MagicCode{
		Code:  code,
		Email: email,
	}
}
