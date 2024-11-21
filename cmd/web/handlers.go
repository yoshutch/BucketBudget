package main

import (
	"net/http"

	"github.com/yoshutch/BucketBudget/internal/models"
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

func (app *application) createBucket(w http.ResponseWriter, r *http.Request) {
	type formData struct {
		Form struct {
			Name    string
			Balance string
		}
		Base BaseTemplateData
	}
	data := formData{
		Form: struct {
			Name    string
			Balance string
		}{
			Name:    "",
			Balance: "0.00",
		},
		Base: BaseTemplateData{},
	}
	app.render(w, r, http.StatusOK, "createBucket.tmpl.html", data)
}
