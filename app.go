package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"text/template"

	model "golang_web_api_application_sample/models"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)
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
