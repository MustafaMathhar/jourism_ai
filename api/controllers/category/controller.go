package category

import (
	"net/http"
	"strconv"

	"github.com/uptrace/bun"
	"goyave.dev/goyave/v4"

	"github.com/MustafaMathhar/jourism_ai/datastore/models"
)

type Controller struct{ DB *bun.DB }

func (cc *Controller) Index(res *goyave.Response, req *goyave.Request) {
	var categories []models.Category
	err := cc.DB.NewSelect().
		Model(&categories).
		Relation("Attractions").
		Scan(req.Request().Context())
	if err != nil {
		res.Status(http.StatusNotFound)
		res.Error(err)
	} else {
		res.JSON(http.StatusOK, categories)
	}
}

func (cc *Controller) Show(res *goyave.Response, req *goyave.Request) {
	id, err := strconv.Atoi(req.Params["id"])
	if err != nil {
		res.Status(http.StatusNotFound)
		res.Error(err)
	}
	category := new(models.Category)
	err = cc.DB.NewSelect().
		Model(&category).
		Relation("Attractions").
		Where("id = ?", id).
		Limit(1).
		Scan(req.Request().Context())
	if err != nil {
		res.Status(http.StatusNotFound)
		res.Error(err)
	} else {
		res.JSON(http.StatusOK, category)
	}
}
