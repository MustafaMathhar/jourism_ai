package profile

import (
	"log"
	"net/http"
	"strconv"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/uptrace/bun"
	"goyave.dev/goyave/v4"

	"github.com/MustafaMathhar/jourism_ai/datastore/models"
)

type Controller struct{ DB *bun.DB }

func (pc *Controller) Show(res *goyave.Response, req *goyave.Request) {
	id, err := strconv.Atoi(req.Params["id"])
	if err != nil {
		res.Status(http.StatusNotFound)
		res.Error(err)
	}
	profile := new(models.Profile)
	err = pc.DB.NewSelect().
		Model(profile).
		Relation("Country").
		Where("profile.id = ?", id).
		Limit(1).
		Scan(req.Request().Context())
	if err != nil {
		res.Status(http.StatusNotFound)
		res.Error(err)
	} else {
		res.JSON(http.StatusOK, &profile)
	}
}

func (pc *Controller) Update(res *goyave.Response, req *goyave.Request) {
	id, err := strconv.Atoi(req.Params["id"])
	if err != nil {
		res.Status(http.StatusNotFound)
		res.Error(err)
	}
	profile := new(models.Profile)
	err = pc.DB.NewSelect().
		Model(profile).
		Relation("Country").
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

func (pc *Controller) UpdateLikedAttractions(res *goyave.Response, req *goyave.Request) {
	id, err := strconv.Atoi(req.Params["id"])
	if err != nil {
		res.Status(http.StatusNotFound)
		res.Error(err)
	}
	likedID := req.String("attractionId")
	respDB, err := pc.DB.NewUpdate().
		Model((*models.Attraction)(nil)).
		Set("profile_id = ?", id).
		Where("id = ?", likedID).
		Exec(req.Request().Context())
	if err != nil {
		res.Status(http.StatusNotFound)
		res.Error(err)
	} else {
		res.JSON(http.StatusOK, &respDB)
	}
}

func (pc *Controller) IndexLikedAttractions(res *goyave.Response, req *goyave.Request) {
	id, err := strconv.Atoi(req.Params["id"])
	if err != nil {
		res.Status(http.StatusNotFound)
		res.Error(err)
	}
	var liked []models.Attraction
	respDB, err := pc.DB.
		NewSelect().
		Model(&liked).Where("profile_id = ?", id).
		ScanAndCount(req.Request().Context())
	if err != nil {
		res.Status(http.StatusNotFound)
		res.Error(err)
	} else {
		res.JSON(http.StatusOK, struct {
			Liked *[]models.Attraction `json:"liked"`
			Count int                  `json:"count"`
		}{
			Count: respDB,
			Liked: &liked,
		})
	}
}

func (pc *Controller) DestroyLikedAttractions(res *goyave.Response, req *goyave.Request) {
	id, err := strconv.Atoi(req.Params["id"])
	if err != nil {
		res.Status(http.StatusNotFound)
		res.Error(err)
	}
	attractionIds := req.Data["attractionIds"]
	respDB, err := pc.DB.
		NewUpdate().
		Model((*models.Attraction)(nil)).
		Set("profile_id = NULL").
		Where("id IN (?)", bun.In(attractionIds)).
		Where("profile_id = ?", id).
		Exec(req.Request().Context())
	if err != nil {
		res.Status(http.StatusNotFound)
		res.Error(err)
	} else {
		res.JSON(http.StatusOK, &respDB)
	}
}

func (cc *Controller) IndexPlans(res *goyave.Response, req *goyave.Request) {
	id, err := strconv.Atoi(req.Params["id"])
	if err != nil {
		res.Status(http.StatusNotFound)
		res.Error(err)
	}
	var plans []models.Plan
	err = cc.DB.NewSelect().
		Model(&plans).
		Relation("Days", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.NewSelect().Relation("Attractions")
		}).
		Where("profile_id = ?", id).
		Scan(req.Request().Context())
	if err != nil {
		res.Status(http.StatusNotFound)
		log.Println(err)
		res.Error(err)
	} else {
		res.JSON(http.StatusOK, &plans)
	}
}

func (cc *Controller) StorePlan(res *goyave.Response, req *goyave.Request) {
	id, err := strconv.Atoi(req.Params["id"])
	if err != nil {
		res.Status(http.StatusNotFound)
		res.Error(err)
	}

	id64 := int64(id)
	plan := &models.Plan{
		ProfileID: &id64,
		Name:      req.String("name"),
	}
	rest, errs := cc.DB.NewInsert().
		Model(plan).
		Exec(req.Request().Context())

	if err != nil {
		res.Status(http.StatusNotFound)
		res.Error(errs)
	} else {
		res.JSON(http.StatusOK, &rest)
	}
}

func (cc *Controller) UpdatePlans(res *goyave.Response, req *goyave.Request) {
	id, err := strconv.Atoi(req.Params["planId"])
	if err != nil {
		res.Status(http.StatusNotFound)
		res.Error(err)
	}
	id64 := int64(id)
	daysPlanned := req.Data["days"].(float64)
	var days []models.Day
	const (
		alphabet = "0123456789abcdefghijklmnopqrstuvwxyz"
		length   = 15
	)
	for i := int32(0); i < int32(daysPlanned); i++ {
		id := gonanoid.MustGenerate(alphabet, length)
		days = append(days, models.Day{ID: id, PlanID: id64})
	}
	_, err = cc.DB.NewInsert().Model(&days).Exec(req.Request().Context())

	if err != nil {
		res.Status(http.StatusNotFound)
		res.Error(err)
	} else {
		res.JSON(http.StatusOK, &days)
	}
}
