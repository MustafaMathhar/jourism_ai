package attraction

import (
	"net/http"
	"strconv"

	"github.com/uptrace/bun"
	"goyave.dev/goyave/v4"

	"github.com/MustafaMathhar/jourism_ai/datastore/models"
)

type Controller struct{ DB *bun.DB }

func (ac *Controller) Index(res *goyave.Response, req *goyave.Request) {
	var attractions []models.Attraction
	err := ac.DB.NewSelect().
		Model(&attractions).
		Scan(req.Request().Context())
	if err != nil {
		res.Status(http.StatusNotFound)
		res.Error(err)
	} else {
		res.JSON(http.StatusOK, &attractions)
	}
}

func (ac *Controller) Show(res *goyave.Response, req *goyave.Request) {
	id, err := strconv.Atoi(req.Params["id"])
	if err != nil {
		res.Status(http.StatusNotFound)
		res.Error(err)
	}
	attraction := new(models.Attraction)
	err = ac.DB.NewSelect().
		Model(&attraction).
		Where("id = ?", id).
		Limit(1).
		Scan(req.Request().Context())
	if err != nil {
		res.Status(http.StatusNotFound)
		res.Error(err)
	} else {
		res.JSON(http.StatusOK, &attraction)
	}
}
