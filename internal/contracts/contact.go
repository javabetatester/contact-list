package contracts

type ContactRequest struct {
	Name     string `gorm:"not null"`
	Email    string `gorm:"unique"`
	Telefone string `gorm:"not null"`
}

type ContactResponse struct {
	ID       int    `gorm:"primaryKey;autoIncrement"`
	Name     string `gorm:"not null"`
	Email    string `gorm:"unique"`
	Telefone string `gorm:"not null"`
}
