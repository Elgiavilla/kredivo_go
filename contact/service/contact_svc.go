package service

import (
	"time"

	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/elgiavilla/kredivo/contact"
	"github.com/elgiavilla/kredivo/models"
)

type contactService struct {
	contactRepo     contact.Repository
	ctxTimeDuration time.Duration
}

func NewContactSvc(c contact.Repository, timout time.Duration) contact.Service {
	return &contactService{
		contactRepo:     c,
		ctxTimeDuration: timout,
	}
}

func (c *contactService) Insert(contactModel models.Contact) (*models.Contact, error) {
	d, err := c.contactRepo.Insert(contactModel)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (c *contactService) GetAll(page int, limit int) (*pagination.Paginator, error) {
	d, err := c.contactRepo.GetAll(page, limit)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (c *contactService) GetById(id uint) (*models.Contact, error) {
	d, err := c.contactRepo.GetById(id)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (c *contactService) Delete(id uint) error {
	d := c.contactRepo.Delete(id)
	return d
}

func (c *contactService) Update(contactModel models.Contact) (*models.Contact, error) {
	d, err := c.contactRepo.Update(contactModel)
	if err != nil {
		return nil, err
	}
	return d, nil
}
