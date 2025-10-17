package groups

import (
	"contactlist/internal/domain/contact"
)

type Group struct {
	ID       int               `gorm:"primaryKey;autoIncrement"`
	Name     string            `gorm:"not null"`
	Contacts []contact.Contact `gorm:"many2many:group_contacts;"`
}
