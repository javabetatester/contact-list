package main

import (
	"contactlist/internal/domain/contact"
	"contactlist/internal/domain/groups"
	"contactlist/internal/infrastructure"
	"contactlist/internal/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	PORT := "8080"
	db := infrastructure.NewDb()

	ContactService := contact.Service{
		Repository: &infrastructure.ContactsRepository{DB: db},
	}

	GroupService := groups.Service{
		Repository: &infrastructure.GroupsRepository{DB: db},
	}

	handler := routes.Handler{
		ContactService: &ContactService,
		GroupService:   &GroupService,
	}

	r := gin.Default()
	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	//CONTATOS
	r.POST("/contact", handler.CreateContact)
	r.PUT("/contact/:id", handler.UpdateContact)
	r.GET("/contact", handler.ListContacts)
	r.GET("/contact/:id", handler.GetContactById)
	r.GET("/contact/name", handler.GetContactByName)
	r.GET("/contact/phone", handler.GetContactByPhone)
	r.GET("/contact/email", handler.GetContactByEmail)
	r.DELETE("/contact/:id", handler.DeleteContact)

	//GRUPOS
	r.POST("/group", handler.CreateGroup)
	r.PUT("/group/:id", handler.UpdateGroup)
	r.GET("/group", handler.ListGroups)
	r.GET("/group/:id", handler.GetGroupById)
	r.GET("/group/name", handler.GetGroupByName)
	r.DELETE("/group/:id", handler.DeleteGroup)
	r.PUT("/group/:id/contact/:contactId", handler.AddContactToGroup)

	if err := r.Run(":" + PORT); err != nil {
		log.Fatalf("cannot start server on port %s: %v", PORT, err)
	}
}
