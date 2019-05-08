package repository

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/elgiavilla/kredivo/contact"
	"github.com/elgiavilla/kredivo/models"
	"github.com/jinzhu/gorm"
)

type dbContactRepo struct {
	db *gorm.DB
}

func NewContactRepo(Conn *gorm.DB) contact.Repository {
	return &dbContactRepo{Conn}
}

func (c *dbContactRepo) Insert(contactModel models.Contact) (*models.Contact, error) {
	row := new(models.Contact)
	d := c.db.Debug().Create(&contactModel).Scan(&row)
	if d.Error != nil {
		return nil, d.Error
	}
	return row, nil
}

func (c *dbContactRepo) GetAll(page int, limit int) (*pagination.Paginator, error) {
	var contactList []*models.Contact
	list := c.db.Debug().Find(&contactList)
	if list.Error != nil {
		return nil, list.Error
	}
	paginator := pagination.Paging(&pagination.Param{
		DB:      list,
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"id desc"},
	}, &contactList)
	return paginator, nil
}

func (c *dbContactRepo) GetById(id uint) (*models.Contact, error) {
	row := new(models.Contact)
	d := c.db.Debug().Where("id = ?", id).Find(&row)
	if d.Error != nil {
		return nil, d.Error
	}
	return row, nil
}

func (c *dbContactRepo) Delete(id uint) error {
	row := new(models.Contact)
	d := c.db.Debug().Where("id = ?", id).Delete(&row)
	if d.Error != nil {
		return d.Error
	}
	return nil
}

func (c *dbContactRepo) Update(contactModel models.Contact) (*models.Contact, error) {
	row := new(models.Contact)
	update := c.db.Debug().Model(&contactModel).Update(&contactModel).Scan(&row)
	if update.Error != nil {
		return nil, update.Error
	}
	return row, nil
}
