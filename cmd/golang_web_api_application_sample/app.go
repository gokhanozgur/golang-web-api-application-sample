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
	"golang_web_api_application_sample/pkg/repository/interest"
	"golang_web_api_application_sample/pkg/repository/user"
	userInterestRepository "golang_web_api_application_sample/pkg/repository/user_interest"
	interestService "golang_web_api_application_sample/pkg/service/interest"
	userService "golang_web_api_application_sample/pkg/service/user"
	userInterestService "golang_web_api_application_sample/pkg/service/user_interest"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func main() {

	// Application declaring
	app := App{}

	// Initialize Application

	// Get environment variable with viper package
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	app.initialize(
		viper.GetString("APP_DB_HOST"),
		viper.GetString("APP_DB_PORT"),
		viper.GetString("APP_DB_USERNAME"),
		viper.GetString("APP_DB_PASSWORD"),
		viper.GetString("APP_DB_NAME"))

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

	// Posgres connection string
	//connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbname)

	// MySQL Connection string
	connectionString := fmt.Sprintf("%s:%s@/%s", username, password, dbname)

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

	// User API`s
	userAPI := InitializeUserAPI(app.DB)
	app.Router.HandleFunc("/users", userAPI.GetAllUser()).Methods("GET")
	app.Router.HandleFunc("/user/create", userAPI.CreateUser()).Methods("POST")
	app.Router.HandleFunc("/user/{id:[0-9]+}", userAPI.FindUserByID()).Methods("GET")
	app.Router.HandleFunc("/user/update/{id:[0-9]+}", userAPI.UpdateUser()).Methods("PUT")
	app.Router.HandleFunc("/user/delete/{id:[0-9]+}", userAPI.DeleteUser()).Methods("DELETE")

	// Interest API`s
	interestAPI := InitializeInterestAPI(app.DB)
	app.Router.HandleFunc("/interests", interestAPI.GetAllInterest()).Methods("GET")
	app.Router.HandleFunc("/interest/create", interestAPI.CreateInterest()).Methods("POST")
	app.Router.HandleFunc("/interest/{id:[0-9]+}", interestAPI.FindInterestByID()).Methods("GET")
	app.Router.HandleFunc("/interest/update/{id:[0-9]+}", interestAPI.UpdateInterest()).Methods("PUT")
	app.Router.HandleFunc("/interest/delete/{id:[0-9]+}", interestAPI.DeleteInterest()).Methods("DELETE")

	// User Interest API`s
	userInterestAPI := InitializeUserInterestAPI(app.DB)
	app.Router.HandleFunc("/user/interests", userInterestAPI.GetAllUserInterest()).Methods("GET")
	app.Router.HandleFunc("/user/interest/create", userInterestAPI.CreateUserInterest()).Methods("POST")
	app.Router.HandleFunc("/user/{id:[0-9]+}/interests", userInterestAPI.FindByUserID()).Methods("GET")
	app.Router.HandleFunc("/user/interest/{id:[0-9]+}", userInterestAPI.FindByInterestID()).Methods("GET")
	app.Router.HandleFunc("/user/interest/update/{id:[0-9]+}", userInterestAPI.UpdateUserInterest()).Methods("PUT")
	app.Router.HandleFunc("/user/interest/delete/{id:[0-9]+}", userInterestAPI.DeleteUserInterest()).Methods("DELETE")

}

// Initialize User API
func InitializeUserAPI(db *gorm.DB) api.UserAPI {
	userRepository := user.NewRepository(db)
	userService := userService.NewUserService(userRepository)
	userAPI := api.NewUserAPI(userService)
	return userAPI
}

// Initialize Interest API
func InitializeInterestAPI(db *gorm.DB) api.InterestAPI {
	interestRepository := interest.NewRepository(db)
	interestService := interestService.NewInterestService(interestRepository)
	interestAPI := api.NewInterestAPI(interestService)
	return interestAPI
}

// Initialize User Interest API
func InitializeUserInterestAPI(db *gorm.DB) api.UserInterestAPI {
	userInterestRepository := userInterestRepository.NewRepository(db)
	userInterestService := userInterestService.NewUserInterestService(userInterestRepository)
	userInterestAPI := api.NewUserInterestAPI(userInterestService)
	return userInterestAPI
}

func handler(rw http.ResponseWriter, r *http.Request) {
	page := model.Page{ID: 3, Name: "Kullan??c??lar", Description: "Kullan??c?? Listesi", URI: "/users"}
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
