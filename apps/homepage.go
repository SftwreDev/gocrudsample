package apps

import (
	"container/list"
	"gocrudsample/utils"
	"net/http"
)

func Homepage(w http.ResponseWriter, r *http.Request) {
	// Homepage handles the HTTP request for the home page.
	// It renders an HTML templates using the provided http.ResponseWriter and http.Request.
	// If there is an error loading the templates, it logs the error and returns.
	// Declare HTML templates here

	utils.RenderTemplate(w, "homepage", list.List{})
}
