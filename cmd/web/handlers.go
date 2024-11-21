package main

import (
	"net/http"

	"github.com/yoshutch/BucketBudget/internal/models"
	"github.com/yoshutch/BucketBudget/internal/validator"
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

type bucketForm struct {
	Name                string `form:"name"`
	Balance             string `form:"balance"`
	validator.Validator `form:"-"`
}

func (app *application) createBucket(w http.ResponseWriter, r *http.Request) {
	type formData struct {
		Form bucketForm
		Base BaseTemplateData
	}
	data := formData{
		Form: bucketForm{
			Name:    "",
			Balance: "0.00",
		},
		Base: BaseTemplateData{},
	}
	app.render(w, r, http.StatusOK, "createBucket.tmpl.html", data)
}

func (app *application) createBucketPost(w http.ResponseWriter, r *http.Request) {
	var form bucketForm
	err := r.ParseForm()
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	form.Name = r.FormValue("name")
	form.Balance = r.FormValue("balance")

	form.CheckField(validator.IsNotBlank(form.Name), "name", "must not be blank")
	form.CheckField(validator.MaxChars(form.Name, 64), "name", "must be less than 64 characters")
	form.CheckField(validator.IsNotBlank(form.Balance), "balance", "must not be blank")
	// validate regex to prove period in string and only digits?

	type formData struct {
		Form bucketForm
		Base BaseTemplateData
	}
	if !form.Valid() {
		app.logger.Debug("Form not valid!")
		data := formData{
			Form: form,
			Base: BaseTemplateData{},
		}
		app.render(w, r, http.StatusOK, "createBucket.tmpl.html", data)
		return
	}

	balance, err := models.ParseAmountFromString(form.Balance)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	app.bucketService.NewBucket(form.Name, balance)

	app.logger.Debug("Saved successfully!")
	data := formData{
		Form: form,
		Base: BaseTemplateData{Flash: "Saved successfully!"},
	}
	app.render(w, r, http.StatusOK, "createBucket.tmpl.html", data)
	// http.Redirect(w, r, "/snippet/create", http.StatusSeeOther)
}
