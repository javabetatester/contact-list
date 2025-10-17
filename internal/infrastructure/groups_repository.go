package infrastructure

import (
	"contactlist/internal/domain/groups"

	"gorm.io/gorm"
)

type GroupsRepository struct {
	DB *gorm.DB
}

func (r *GroupsRepository) Create(group groups.Group) error {
	return r.DB.Create(&group).Error
}

func (r *GroupsRepository) List() ([]groups.Group, error) {
	var groups []groups.Group
	return groups, r.DB.Preload("Contacts").Find(&groups).Error
}

func (r *GroupsRepository) GetById(id int) (groups.Group, error) {
	var group groups.Group
	return group, r.DB.Preload("Contacts").First(&group, id).Error
}

func (r *GroupsRepository) GetByName(name string) (groups.Group, error) {
	var group groups.Group
	return group, r.DB.Preload("Contacts").Where("name = ?", name).First(&group).Error
}

func (r *GroupsRepository) Update(group groups.Group) error {
	return r.DB.Save(&group).Error
}

func (r *GroupsRepository) Delete(id int) error {
	return r.DB.Delete(&groups.Group{}, id).Error
}

func (r *GroupsRepository) AddContactToGroup(groupID int, contactID int) error {
	return r.DB.Exec("INSERT INTO group_contacts (group_id, contact_id) VALUES (?, ?)", groupID, contactID).Error
}
