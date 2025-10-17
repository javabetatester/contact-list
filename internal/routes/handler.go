package routes

import (
	"contactlist/internal/domain/contact"
	"contactlist/internal/domain/groups"
)

type Handler struct {
	ContactService *contact.Service
	GroupService *groups.Service
}