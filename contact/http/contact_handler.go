package http

import (
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/elgiavilla/kredivo/contact"
	"github.com/elgiavilla/kredivo/models"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

type ReponseError struct {
	Message string `json:"message"`
}

type HttpContactHandler struct {
	ContactService contact.Service
}

func NewContactHandler(e *echo.Echo, contactService contact.Service) {
	handler := &HttpContactHandler{
		ContactService: contactService,
	}
	e.POST("/api/contact", handler.Insert)
	e.GET("/api/contacts", handler.GetAll)
	e.GET("/api/contact/:id", handler.GetById)
	e.DELETE("/api/contact/:id", handler.Delete)
	e.PUT("/api/contact", handler.Update)
}

func (co *HttpContactHandler) Insert(c echo.Context) error {
	var contact models.Contact
	err := c.Bind(&contact)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(getStatusCode(err), ReponseError{Message: err.Error()})
	}

	src, err := file.Open()
	if err != nil {
		c.JSON(getStatusCode(err), ReponseError{Message: err.Error()})
	}
	defer src.Close()

	dst, err := os.Create(file.Filename)
	if err != nil {
		c.JSON(getStatusCode(err), ReponseError{Message: err.Error()})
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		c.JSON(getStatusCode(err), ReponseError{Message: err.Error()})
	}

	contact.Profile_picture = file.Filename

	d, err := co.ContactService.Insert(contact)
	if err != nil {
		c.JSON(getStatusCode(err), ReponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, d)
}

func (co *HttpContactHandler) GetAll(c echo.Context) error {
	pageS := c.QueryParam("page")
	limitS := c.QueryParam("limit")

	page, _ := strconv.Atoi(pageS)
	limit, _ := strconv.Atoi(limitS)

	d, err := co.ContactService.GetAll(page, limit)
	if err != nil {
		return c.JSON(getStatusCode(err), ReponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, d)
}

func (co *HttpContactHandler) GetById(c echo.Context) error {
	idP, _ := strconv.Atoi(c.Param("id"))
	id := uint(idP)

	d, err := co.ContactService.GetById(id)
	if err != nil {
		return c.JSON(getStatusCode(err), ReponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, d)
}

func (co *HttpContactHandler) Delete(c echo.Context) error {
	idP, _ := strconv.Atoi(c.Param("id"))
	id := uint(idP)

	err := co.ContactService.Delete(id)
	if err != nil {
		return c.JSON(getStatusCode(err), ReponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, "Deleted")
}

func (co *HttpContactHandler) Update(c echo.Context) error {
	var contact models.Contact
	err := c.Bind(&contact)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(getStatusCode(err), ReponseError{Message: err.Error()})
	}

	src, err := file.Open()
	if err != nil {
		c.JSON(getStatusCode(err), ReponseError{Message: err.Error()})
	}
	defer src.Close()

	dst, err := os.Create(file.Filename)
	if err != nil {
		c.JSON(getStatusCode(err), ReponseError{Message: err.Error()})
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		c.JSON(getStatusCode(err), ReponseError{Message: err.Error()})
	}

	contact.Profile_picture = file.Filename

	d, err := co.ContactService.Update(contact)
	if err != nil {
		c.JSON(getStatusCode(err), ReponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, d)
}

func getStatusCode(err error) int {

	if err == nil {
		return http.StatusOK
	}
	logrus.Error(err)
	switch err {
	case models.ErrInternalServerError:
		return http.StatusInternalServerError
	case models.ErrNotFound:
		return http.StatusNotFound
	case models.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
