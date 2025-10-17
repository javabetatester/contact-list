package routes

import (
	"contactlist/internal/domain/contact"
	"contactlist/internal/contracts"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ContactRequest(contract *contracts.ContactRequest) *contact.Contact {
	return &contact.Contact{
		Name:     contract.Name,
		Email:    contract.Email,
		Telefone: contract.Telefone,
	}
}

func ContactResponse(contact *contact.Contact) *contracts.ContactResponse {
	return &contracts.ContactResponse{
		ID:       contact.ID,
		Name:     contact.Name,
		Email:    contact.Email,
		Telefone: contact.Telefone,
	}
}

func (h *Handler) CreateContact(ctx *gin.Context) {
	var contact contracts.ContactRequest
	if err := ctx.ShouldBindJSON(&contact); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var contactEntity = *ContactRequest(&contact)
	err := h.ContactService.Create(contactEntity)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	response := ContactResponse(&contactEntity)
	ctx.JSON(http.StatusCreated, response)
}

func (h *Handler) ListContacts(ctx *gin.Context) {
	contacts, err := h.ContactService.List()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	response := make([]*contracts.ContactResponse, 0)
	for _, contact := range contacts {
		response = append(response, ContactResponse(&contact))
	}
	ctx.JSON(http.StatusOK, response)
}

func (h *Handler) GetContactById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	contact, err := h.ContactService.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	response := ContactResponse(&contact)
	ctx.JSON(http.StatusOK, response)
}

func (h *Handler) GetContactByName(ctx *gin.Context) {
	var name string
	if err := ctx.ShouldBindJSON(&name); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	contact, err := h.ContactService.GetByName(name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	response := ContactResponse(&contact)
	ctx.JSON(http.StatusOK, response)
}

func (h *Handler) GetContactByEmail(ctx *gin.Context) {
	var email string
	if err := ctx.ShouldBindJSON(&email); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	contact, err := h.ContactService.GetByEmail(email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	response := ContactResponse(&contact)
	ctx.JSON(http.StatusOK, response)
}

func (h *Handler) GetContactByPhone(ctx *gin.Context) {
	var telefone string
	if err := ctx.ShouldBindJSON(&telefone); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	contact, err := h.ContactService.GetByPhone(telefone)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	response := ContactResponse(&contact)
	ctx.JSON(http.StatusOK, response)
}

func (h *Handler) UpdateContact(ctx *gin.Context) {
	var contact contact.Contact
	if err := ctx.ShouldBindJSON(&contact); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.ContactService.Update(contact)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := ContactResponse(&contact)
	ctx.JSON(http.StatusOK, response)
}

func (h *Handler) DeleteContact(ctx *gin.Context) {
	var contact contact.Contact
	if err := ctx.ShouldBindJSON(&contact); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.ContactService.Delete(contact)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "contact deleted"})
}