package contact

type Repository interface {
	Create(contact Contact) error
	List() ([]Contact, error)
	Update(contact Contact) error
	Delete(contact Contact) error
	GetById(id int) (Contact, error)
	GetByName(name string) (Contact, error)
	GetByEmail(email string) (Contact, error)
	GetByPhone(telefone string) (Contact, error)
}