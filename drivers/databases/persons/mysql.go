package persons

import "gorm.io/gorm"

type PersonRepository struct {
	Conn *gorm.DB
}
