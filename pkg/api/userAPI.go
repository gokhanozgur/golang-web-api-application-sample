package api

import (
	"encoding/json"
	"golang_web_api_application_sample/pkg/model"
	service "golang_web_api_application_sample/pkg/service/user"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UserAPI struct {
	UserService service.UserService
}

// UserAPI constructor
func NewUserAPI(u service.UserService) UserAPI {
	return UserAPI{UserService: u}
}

// Get all user
func (u UserAPI) GetAllUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		users, err := u.UserService.All()
		if err != nil {
			RespondWithError(rw, http.StatusNotFound, err.Error())
			return
		}
		RespondWithJSON(rw, http.StatusOK, users)
	}
}

// Find user by id
func (u UserAPI) FindUserByID() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			RespondWithError(rw, http.StatusNotFound, err.Error())
			return
		}
		user, err := u.UserService.FindByID(uint64(id))
		if err != nil {
			RespondWithError(rw, http.StatusNotFound, err.Error())
			return
		}
		RespondWithJSON(rw, http.StatusOK, user)
	}
}

// Create a user
func (u UserAPI) CreateUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var userDTO model.UserDTO

		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()
		if err := decoder.Decode(&userDTO); err != nil {
			RespondWithError(rw, http.StatusBadRequest, err.Error())
			return
		}

		createUser, err := u.UserService.Save(model.ToUser(&userDTO))
		if err != nil {
			RespondWithError(rw, http.StatusInternalServerError, err.Error())
			return
		}

		RespondWithJSON(rw, http.StatusCreated, createUser)
	}
}

// Update a user
func (u UserAPI) UpdateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		var userDTO model.UserDTO
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()
		if err := decoder.Decode(&userDTO); err != nil {
			RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		user, err := u.UserService.FindByID(uint64(id))
		if err != nil {
			RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		user.FirstName = userDTO.FirstName
		user.LastName = userDTO.LastName
		user.Username = userDTO.Username
		user.Profile = userDTO.Profile
		user.Interests = user.Interests
		updateduser, err := u.UserService.Save(user)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		RespondWithJSON(w, http.StatusOK, model.ToUserDTO(updateduser))
	}
}

// Hard delete a user
func (u UserAPI) DeleteUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		user, err := u.UserService.FindByID(uint64(id))
		if err != nil {
			RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		err = u.UserService.Delete(user.ID)
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
func (u UserAPI) Migrate() {
	err := u.UserService.Migrate()
	if err != nil {
		log.Println(err)
	}
}
