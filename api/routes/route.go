package routes

import (
	"github.com/uptrace/bun"
	"goyave.dev/goyave/v4"

	"github.com/MustafaMathhar/jourism_ai/api/controllers/attraction"
	"github.com/MustafaMathhar/jourism_ai/api/controllers/category"
	"github.com/MustafaMathhar/jourism_ai/api/controllers/profile"
	"github.com/MustafaMathhar/jourism_ai/datastore/models"
)

type DataStore struct{ DB *bun.DB }

func (ds *DataStore) Register(router *goyave.Router) {
	CategoryRoutes(router, ds.DB)
	AttractionRoutes(router, ds.DB)
	ProfileRoutes(router, ds.DB)
}

// Handler function for the "/hello" route
func CategoryRoutes(router *goyave.Router, db *bun.DB) {
	db.RegisterModel((*models.AttractionsToCategories)(nil))
	cc := &category.Controller{DB: db}
	categoryRouter := router.Subrouter("/categories")
	categoryRouter.Get("", cc.Index)
	categoryRouter.Get("/{id:[0-9]+}", cc.Show)
}

func AttractionRoutes(router *goyave.Router, db *bun.DB) {
	ac := &attraction.Controller{DB: db}
	attractionRouter := router.Subrouter("/attractions")
	attractionRouter.Get("", ac.Index)
	attractionRouter.Get("/{id:[0-9]+}", ac.Show)
}

func ProfileRoutes(router *goyave.Router, db *bun.DB) {
	db.RegisterModel(
		(*models.AttractionsToDays)(nil),
		(*models.ProfilesToAttractions)(nil),
	)
	pc := &profile.Controller{DB: db}
	profileRouter := router.Subrouter("/profile/{id:[0-9]+}")
	profileRouter.Get("", pc.Show)

	likedAttractions := profileRouter.Group().Subrouter("/liked")
	likedAttractions.Put("/{attractionId:[0-9]+}", pc.UpdateLikedAttractions)
	likedAttractions.Get("", pc.IndexLikedAttractions)
	likedAttractions.Get("/{attractionId:[0-9]+}", pc.IsLikedAttraction)
	likedAttractions.Delete("/{attractionId:[0-9]+}", pc.DestroyLikedAttractions)
	plans := profileRouter.Group().Subrouter("/plans")
	plans.Get("", pc.IndexPlans)
	plans.Post("", pc.StorePlan)
	plans.Put("/{planId:[0-9]+}", pc.UpdatePlans)

	plans.Get("/{planId:[0-9]+}", pc.IndexDays)
	plans.Get("/{planId:[0-9]+}/{day}", pc.ShowDay)
	plans.Post("/{planId:[0-9]+}/{day}", pc.UpdateDay)
}
