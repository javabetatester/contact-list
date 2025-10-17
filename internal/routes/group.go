package routes

import (
	"contactlist/internal/domain/groups"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateGroup(ctx *gin.Context) {
	var group groups.Group
	if err := ctx.ShouldBindJSON(&group); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	groupEntity := groups.Group{
		Name: group.Name,
	}
	err := h.GroupService.Create(groupEntity)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, &groupEntity)
}

func (h *Handler) ListGroups(ctx *gin.Context) {
	group, err := h.GroupService.List()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	response := make([]*groups.Group, 0)
	for _, group := range group {
		response = append(response, &group)
	}
	ctx.JSON(http.StatusOK, response)
}

func (h *Handler) GetGroupById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	group, err := h.GroupService.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, &group)
}

func (h *Handler) GetGroupByName(ctx *gin.Context) {
	var name string
	if err := ctx.ShouldBindJSON(&name); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	group, err := h.GroupService.GetByName(name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, &group)
}

func (h *Handler) DeleteGroup(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.GroupService.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "group deleted"})
}

func (h *Handler) UpdateGroup(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var group groups.Group
	if err := ctx.ShouldBindJSON(&group); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	groupEntity := groups.Group{
		ID:   id,
		Name: group.Name,
	}
	err = h.GroupService.Update(groupEntity)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, &groupEntity)
}

func (h *Handler) AddContactToGroup(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	contactId, err := strconv.Atoi(ctx.Param("contactId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.GroupService.AddContactToGroup(id, contactId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "contact added to group"})
}
