package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"

	"golang_web_api_application_sample/pkg/api"
	model "golang_web_api_application_sample/pkg/model"
	"golang_web_api_application_sample/pkg/repository/user"
	service "golang_web_api_application_sample/pkg/service/user"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func main() {

	// Application declaring
	app := App{}

	// Initialize routes
	app.routes()

	// Run
	app.run(":9001")

	/*
		http.HandleFunc("/", handler)
		http.ListenAndServe(":9000", nil)
	*/
}

// App Initializer
func (app *App) initialize(host, port, username, password, dbname string) {

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbname)

	var err error
	app.DB, err = gorm.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	app.Router = mux.NewRouter()
}

// Run
func (app *App) run(address string) {
	fmt.Printf("Server started at %s", address)
	log.Fatal(http.ListenAndServe(address, app.Router))
}

// Routes
func (app *App) routes() {

	// User API's
	userAPI := InitializeUserAPI(app.DB)
	app.Router.HandleFunc("/users", userAPI.GetAllUser()).Methods("GET")
	app.Router.HandleFunc("/users/create", userAPI.CreateUser()).Methods("POST")
	app.Router.HandleFunc("/users/{id:[0-9]+}", userAPI.FindUserByID()).Methods("GET")
	app.Router.HandleFunc("/users/{id:[0-9]+}", userAPI.UpdateUser()).Methods("PUT")
	app.Router.HandleFunc("/users/delete/{id:[0-9]+}", userAPI.DeleteUser()).Methods("DELETE")

}

// Initialize User API
func InitializeUserAPI(db *gorm.DB) api.UserAPI {
	userRepository := user.NewRepository(db)
	userService := service.NewUserService(userRepository)
	userAPI := api.NewUserAPI(userService)
	return userAPI
}

func handler(rw http.ResponseWriter, r *http.Request) {
	page := model.Page{ID: 3, Name: "Kullanıcılar", Description: "Kullanıcı Listesi", URI: "/users"}
	users := loadUsers()
	interests := loadInterest()
	interestMappings := loadInterestMapping()

	var newUsers []model.User
	for _, user := range users {
		for _, interestMappings := range interestMappings {
			if user.ID == interestMappings.UserID {
				for _, interest := range interests {
					if interestMappings.InterestID == interest.ID {
						user.Interests = append(user.Interests, interest)
					}
				}
			}
		}
		newUsers = append(newUsers, user)
	}

	viewModel := model.UserViewModel{Page: page, Users: newUsers}

	t, _ := template.ParseFiles("template/page.html")
	t.Execute(rw, viewModel)
}

func loadUsers() []model.User {
	bytes, _ := ioutil.ReadFile("json/users.json")
	var users []model.User
	json.Unmarshal(bytes, &users)
	return users
}

func loadInterest() []model.Interest {
	bytes, _ := ioutil.ReadFile("json/interests.json")
	var interest []model.Interest
	json.Unmarshal(bytes, &interest)
	return interest
}

func loadInterestMapping() []model.InterestMapping {
	bytes, _ := ioutil.ReadFile("json/userInterestMappings.json")
	var interestMapping []model.InterestMapping
	json.Unmarshal(bytes, &interestMapping)
	return interestMapping
}
