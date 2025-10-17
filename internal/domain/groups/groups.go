package groups

import (
	"contactlist/internal/domain/contact"
)

type Group struct {
	ID int
	Name string
	Contacts []contact.Contact
}