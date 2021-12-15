package api

import (
	"encoding/json"
	"golang_web_api_application_sample/pkg/model"
	service "golang_web_api_application_sample/pkg/service/interest"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type InterestAPI struct {
	InterestService service.InterestService
}

// InterestAPI constructor
func NewInterestAPI(i service.InterestService) InterestAPI {
	return InterestAPI{InterestService: i}
}

// Get all interest
func (i InterestAPI) GetAllInterest() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		interests, err := i.InterestService.All()
		if err != nil {
			RespondWithError(rw, http.StatusNotFound, err.Error())
			return
		}
		RespondWithJSON(rw, http.StatusOK, interests)
	}
}

// Find interest by id
func (i InterestAPI) FindInterestByID() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			RespondWithError(rw, http.StatusNotFound, err.Error())
			return
		}
		interest, err := i.InterestService.FindByID(uint(id))
		if err != nil {
			RespondWithError(rw, http.StatusNotFound, err.Error())
			return
		}
		RespondWithJSON(rw, http.StatusOK, interest)
	}
}

// Create a interest
func (i InterestAPI) CreateInterest() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var interestDTO model.InterestDTO

		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()
		if err := decoder.Decode(&interestDTO); err != nil {
			RespondWithError(rw, http.StatusBadRequest, err.Error())
			return
		}

		createInterest, err := i.InterestService.Save(model.ToInterest(&interestDTO))
		if err != nil {
			RespondWithError(rw, http.StatusInternalServerError, err.Error())
			return
		}

		RespondWithJSON(rw, http.StatusCreated, createInterest)
	}
}

// Update a interest
func (i InterestAPI) UpdateInterest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		var interestDTO model.InterestDTO
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()
		if err := decoder.Decode(&interestDTO); err != nil {
			RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		interest, err := i.InterestService.FindByID(uint(id))
		if err != nil {
			RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		interest.Name = interestDTO.Name
		interest.Status = interestDTO.Status
		updatedInterest, err := i.InterestService.Save(interest)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		RespondWithJSON(w, http.StatusOK, model.ToInterestDTO(updatedInterest))
	}
}

// Hard delete a interest
func (i InterestAPI) DeleteInterest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		interest, err := i.InterestService.FindByID(uint(id))
		if err != nil {
			RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		err = i.InterestService.Delete(interest.ID)
		if err != nil {
			RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		type Response struct {
			Message string
		}

		response := Response{
			Message: "Successful!",
		}
		RespondWithJSON(w, http.StatusOK, response)
	}
}

// Migration
func (i InterestAPI) Migrate() {
	err := i.InterestService.Migrate()
	if err != nil {
		log.Println(err)
	}
}
