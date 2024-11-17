package main

import (
	"net/http"

	"yosbomb.com/bucketbudget/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	type homeData struct {
		Buckets []models.Bucket
		Base    BaseTemplateData
	}
	buckets, err := app.bucketService.GetMyBuckets()
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	data := homeData{
		Buckets: buckets,
		Base:    BaseTemplateData{},
	}

	app.render(w, r, http.StatusOK, "home.tmpl.html", data)
}
