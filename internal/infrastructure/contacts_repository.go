package infrastructure

import (
	"contactlist/internal/domain/contact"
	"gorm.io/gorm"
)

type ContactsRepository struct {
	DB *gorm.DB
}

func (r *ContactsRepository) Create(contact contact.Contact) error {
	return r.DB.Create(&contact).Error
}

func (r *ContactsRepository) List() ([]contact.Contact, error) {
	var contacts []contact.Contact
	return contacts, r.DB.Find(&contacts).Error
}

func (r *ContactsRepository) Update(contact contact.Contact) error {
	return r.DB.Save(&contact).Error
}

func (r *ContactsRepository) Delete(contact contact.Contact) error {
	return r.DB.Delete(&contact).Error
}

func (r *ContactsRepository) GetById(id int) (contact.Contact, error) {
	var contact contact.Contact
	return contact, r.DB.First(&contact, id).Error
}

func (r *ContactsRepository) GetByName(name string) (contact.Contact, error) {
	var contact contact.Contact
	return contact, r.DB.Where("name = ?", name).First(&contact).Error
}

func (r *ContactsRepository) GetByPhone(phone string) (contact.Contact, error) {
	var contact contact.Contact
	return contact, r.DB.Where("phone = ?", phone).First(&contact).Error
}

func (r *ContactsRepository) GetByEmail(email string) (contact.Contact, error) {
	var contact contact.Contact
	return contact, r.DB.Where("email = ?", email).First(&contact).Error
}
