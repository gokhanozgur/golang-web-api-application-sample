package api

import (
	"encoding/json"
	"golang_web_api_application_sample/pkg/model"
	service "golang_web_api_application_sample/pkg/service/user_interest"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UserInterestAPI struct {
	UserInterestService service.UserInterestService
}

// UserInterestAPI constructor
func NewUserInterestAPI(u service.UserInterestService) UserInterestAPI {
	return UserInterestAPI{UserInterestService: u}
}

// Get all user
func (u UserInterestAPI) GetAllUserInterest() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		users, err := u.UserInterestService.All()
		if err != nil {
			RespondWithError(rw, http.StatusNotFound, err.Error())
			return
		}
		RespondWithJSON(rw, http.StatusOK, users)
	}
}

// Find user interest by user id
func (u UserInterestAPI) FindByUserID() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			RespondWithError(rw, http.StatusNotFound, err.Error())
			return
		}
		user, err := u.UserInterestService.FindByUserID(uint64(id))
		if err != nil {
			RespondWithError(rw, http.StatusNotFound, err.Error())
			return
		}
		RespondWithJSON(rw, http.StatusOK, user)
	}
}

// Find user interest by id
func (u UserInterestAPI) FindByInterestID() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			RespondWithError(rw, http.StatusNotFound, err.Error())
			return
		}
		user, err := u.UserInterestService.FindByInterestID(uint(id))
		if err != nil {
			RespondWithError(rw, http.StatusNotFound, err.Error())
			return
		}
		RespondWithJSON(rw, http.StatusOK, user)
	}
}

// Create a user interest
func (u UserInterestAPI) CreateUserInterest() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var userDTO model.UserInterestDTO

		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()
		if err := decoder.Decode(&userDTO); err != nil {
			RespondWithError(rw, http.StatusBadRequest, err.Error())
			return
		}

		createUserInterest, err := u.UserInterestService.Save(model.ToUserInterest(&userDTO))
		if err != nil {
			RespondWithError(rw, http.StatusInternalServerError, err.Error())
			return
		}

		RespondWithJSON(rw, http.StatusCreated, createUserInterest)
	}
}

// Update a user interest
func (u UserInterestAPI) UpdateUserInterest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		var userInterestDTO model.UserInterestDTO
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()
		if err := decoder.Decode(&userInterestDTO); err != nil {
			RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		userInterest, err := u.UserInterestService.FindByID(uint64(id))
		if err != nil {
			RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		userInterest.UserID = userInterestDTO.UserID
		userInterest.InterestID = userInterestDTO.InterestID
		updateduserInterest, err := u.UserInterestService.Save(userInterest)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		RespondWithJSON(w, http.StatusOK, model.UserToUserWithInterestDTO(updateduserInterest))
	}
}

// Hard delete a user interest
func (u UserInterestAPI) DeleteUserInterest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		userInterest, err := u.UserInterestService.FindByID(uint64(id))
		if err != nil {
			RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		err = u.UserInterestService.Delete(userInterest.ID)
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
func (u UserInterestAPI) Migrate() {
	err := u.UserInterestService.Migrate()
	if err != nil {
		log.Println(err)
	}
}
