package api

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"gocrudsample/models"
	"gocrudsample/structs"
	"gocrudsample/utils"
	"io"
	"net/http"
)

func ActorGetListApi(w http.ResponseWriter, r *http.Request) {
	// ActorGetListApi handles the HTTP GET request to retrieve a list of actors from the database.
	// It sets the "Content-Type" header to "application/json" and fetches the actors using the models.DB instance.
	// The list of actors is then encoded as JSON and written to the HTTP response writer.

	// Set the Content-Type
	w.Header().Set("Content-Type", "application/json")

	// Declare actors models
	var actors []models.Actors
	models.DB.Find(&actors)

	// Json Encoder
	err := json.NewEncoder(w).Encode(actors)
	if err != nil {
		fmt.Println("ERROR", err)
	}

}

func ActorCreateAPI(w http.ResponseWriter, r *http.Request) {
	// ActorCreateAPI handles the HTTP POST request to create a new actor in the database.
	// It expects JSON input containing the actor's first name and last name, which are validated
	// using a validator. If the input is valid, a new actor is created and added to the database.
	// The newly created actor is then encoded as JSON and written to the HTTP response writer.

	var input structs.ActorInput

	body, _ := io.ReadAll(r.Body)
	_ = json.Unmarshal(body, &input)

	validate := validator.New()
	err := validate.Struct(input)

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	actors := &models.Actors{
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}

	models.DB.Create(actors)

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(actors)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

}

func ActorGetAPI(w http.ResponseWriter, r *http.Request) {
	// ActorGetAPI handles the HTTP GET request to retrieve a specific actor from the database by its ID.
	// It sets the "Content-Type" header to "application/json" and fetches the actor using the provided ID
	// from the query parameters. If the actor is found, it is encoded as JSON and written to the HTTP response writer.
	// If no actor is found, a "Not Found" response with an appropriate error message is sent.

	w.Header().Set("Content-Type", "application/json")

	// Get the id from query params
	id := mux.Vars(r)["id"]
	var actors models.Actors

	if err := models.DB.Where("id = ?", id).First(&actors).Error; err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Actors not found")
		return
	}

	json.NewEncoder(w).Encode(actors)
}

func ActorUpdateAPI(w http.ResponseWriter, r *http.Request) {
	// ActorUpdateAPI handles the HTTP PUT request to update an existing actor in the database.
	// It sets the "Content-Type" header to "application/json" and fetches the actor using the provided ID
	// from the query parameters. If the actor is found, it updates its attributes based on the JSON input,
	// validates the input data, and saves the changes to the database. The updated actor is then encoded
	// as JSON and written to the HTTP response writer. If the actor is not found or validation fails,
	// appropriate error responses are sent.

	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var actors models.Actors

	if err := models.DB.Where("id = ?", id).First(&actors).Error; err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Actor not found")
		return
	}

	var input structs.ActorInput

	body, _ := io.ReadAll(r.Body)
	_ = json.Unmarshal(body, &input)

	validate := validator.New()
	err := validate.Struct(input)

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	actors.FirstName = input.FirstName
	actors.LastName = input.LastName

	models.DB.Save(&actors)

	json.NewEncoder(w).Encode(actors)
}

func ActorDeleteAPI(w http.ResponseWriter, r *http.Request) {
	// ActorDeleteAPI handles the HTTP DELETE request to delete an existing actor from the database.
	// It sets the "Content-Type" header to "application/json" and fetches the actor using the provided ID
	// from the query parameters. If the actor is found, it is deleted from the database. A successful
	// deletion results in an HTTP status code 204 (No Content) response. If the actor is not found,
	// an appropriate error response is sent.

	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var actors models.Actors

	if err := models.DB.Where("id = ?", id).First(&actors).Error; err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Actor not found")
		return
	}

	models.DB.Delete(&actors)

	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(actors)
}
