package controllers

import (
	"net/http"
	"strconv"

	"github.com/uptrace/bun"
	"goyave.dev/goyave/v4"

	"github.com/MustafaMathhar/jourism_ai/datastore/models"
)

type ProfileController struct{ DB *bun.DB }

func (pc *ProfileController) Show(res *goyave.Response, req *goyave.Request) {
	id, err := strconv.Atoi(req.Params["id"])
	if err != nil {
		res.Status(http.StatusNotFound)
		res.Error(err)
	}
	profile := new(models.Profile)
	err = pc.DB.NewSelect().
		Model(profile).
		Where("id = ?", id).
		Limit(1).
		Scan(req.Request().Context())
	if err != nil {
		res.Status(http.StatusNotFound)
		res.Error(err)
	} else {
		res.JSON(http.StatusOK, &profile)
	}
}

func (pc *ProfileController) UpdateLikedAttractions(res *goyave.Response, req *goyave.Request) {
	id, err := strconv.Atoi(req.Params["id"])
	if err != nil {
		res.Status(http.StatusNotFound)
		res.Error(err)
	}
	likedID := req.String("attractionId")
	respDB, err := pc.DB.NewUpdate().
		Model((*models.Attraction)(nil)).
		Set("profile_id = ?", id).
		Where("id = ?",likedID).
		Exec(req.Request().Context())
	if err != nil {
		res.Status(http.StatusNotFound)
		res.Error(err)
	} else {
		res.JSON(http.StatusOK, &respDB)
	}
}
