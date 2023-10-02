package profile

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/uptrace/bun"
	"goyave.dev/goyave/v4"

	"github.com/MustafaMathhar/jourism_ai/datastore/models"
)

func (pc *Controller) IndexDays(res *goyave.Response, req *goyave.Request) {
	id, err := strconv.Atoi(req.Params["planId"])
	if err != nil {
		res.Status(http.StatusNotFound)
		res.Error(err)
	}
	plan := new(models.Plan)
	err = pc.DB.NewSelect().
		Model(plan).
		Where("id = ?", id).
		Relation("Days").
		Scan(req.Request().Context())
	if err != nil {
		res.Status(http.StatusNotFound)
		log.Println(err)
		res.Error(err)
	} else {
		res.JSON(http.StatusOK, &plan)
	}
}

func (c *Controller) ShowDay(res *goyave.Response, req *goyave.Request) {
	id, err := strconv.Atoi(req.Params["planId"])
	if err != nil {
		res.Status(http.StatusNotFound)
		res.Error(err)
	}

	dayId := req.Params["day"]
	day := new(models.Day)
	err = c.DB.NewSelect().
		Model(day).
		Relation("Attractions").
		Where("plan_id = ?", id).
		Where("id = ?", dayId).
		Scan(req.Request().Context())
	if err != nil {
		res.Status(http.StatusNotFound)
		log.Println(err)
		res.Error(err)
	} else {
		res.JSON(http.StatusOK, &day)
	}
}

func DayExists(ctx context.Context, db bun.IDB, id int) error {
	res, err := db.NewSelect().Model((*models.Day)(nil)).Where("plan_id = ?", id).Exists(ctx)
	if !res {
		return fmt.Errorf("[ERROR] day does not exist")
	}
	return err
}

func (c *Controller) UpdateDay(res *goyave.Response, req *goyave.Request) {
	dayId := req.Params["day"]
	attraction := req.Data["attraction"].(float64)
	a2d := models.AttractionsToDays{
		DayID:        dayId,
		AttractionID: int64(attraction),
	}
	

	result, err := c.DB.NewInsert().
		Model(&a2d).Exec(req.Request().Context())
	if err != nil {
		res.Status(http.StatusNotFound)
		log.Println(err)
		res.Error(err)
	} else {
		res.JSON(http.StatusOK, &result)
	}
}
