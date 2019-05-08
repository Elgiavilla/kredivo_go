package contact

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/elgiavilla/kredivo/models"
)

type Service interface {
	Insert(contactModel models.Contact) (*models.Contact, error)
	GetAll(page int, limit int) (*pagination.Paginator, error)
	GetById(id uint) (*models.Contact, error)
	Delete(id uint) error
	Update(contactModel models.Contact) (*models.Contact, error)
}
