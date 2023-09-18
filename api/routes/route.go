package routes

import (
	"github.com/uptrace/bun"
	"goyave.dev/goyave/v4"

	"github.com/MustafaMathhar/jourism_ai/api/controllers"
)

type DataStore struct{ DB *bun.DB }

func (ds *DataStore) Register(router *goyave.Router) {
	CategoryRoutes(router, ds.DB)
	AttractionRoutes(router, ds.DB)
	ProfileRoutes(router, ds.DB)
}

// Handler function for the "/hello" route
func CategoryRoutes(router *goyave.Router, db *bun.DB) {
	cc := &controllers.CategoryController{DB: db}
	categoryRouter := router.Subrouter("/categories")
	categoryRouter.Get("", cc.Index)
	categoryRouter.Get("/{id:[0-9]+}", cc.Show)
}

func AttractionRoutes(router *goyave.Router, db *bun.DB) {
	ac := &controllers.AttractionController{DB: db}
	attractionRouter := router.Subrouter("/attractions")
	attractionRouter.Get("", ac.Index)
	attractionRouter.Get("/{id:[0-9]+}", ac.Show)
}

func ProfileRoutes(router *goyave.Router, db *bun.DB) {
	pc := &controllers.ProfileController{DB: db}
	profileRouter := router.Subrouter("/profile/{id:[0-9]+}")
	profileRouter.Get("", pc.Show)
	profileRouter.Put("/liked", pc.UpdateLikedAttractions)
}
