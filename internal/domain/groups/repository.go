package groups

type Repository interface {
	Create(group Group) error
	List() ([]Group, error)
	Update(group Group) error
	Delete(group Group) error
	GetById(id int) (Group, error)
	GetByName(name string) (Group, error)
}
