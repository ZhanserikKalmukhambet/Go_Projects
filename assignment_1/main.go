package main

import (
	"assignment_1/data"
	"assignment_1/seacrh_filters"
	"fmt"
	"net/http"
	"text/template"
)

func templating(w http.ResponseWriter, filename string, data interface{}) {
	t, _ := template.ParseFiles(filename)
	t.ExecuteTemplate(w, filename, data)
}

func getSignInPage(w http.ResponseWriter) {
	templating(w, "sign-in.html", nil)
}

func getSignUpPage(w http.ResponseWriter) {
	templating(w, "sign-up.html", nil)
}

func getUser(r *http.Request) data.User {
	email := r.FormValue("email")
	password := r.FormValue("password")
	return data.User{Email: email, Password: password}
}

func signInUser(w http.ResponseWriter, r *http.Request) {
	newUser := getUser(r)
	ok := data.Admin.VerifyUser(newUser)

	if ok {
		getUserPage(w)
		return
	}
	return
}

func signUpUser(w http.ResponseWriter, r *http.Request) {
	newUser := getUser(r)
	err := data.Admin.CreateUser(newUser)

	fileName := "sign-up.html"
	t, _ := template.ParseFiles(fileName)
	if err != nil {
		t.ExecuteTemplate(w, fileName, "New User Sign-up Failure, Try Again!")
		return
	}

	t.ExecuteTemplate(w, fileName, "New User Sign-up Success!")
	return
}

func getUserPage(w http.ResponseWriter) {
	t, _ := template.ParseFiles("list-products.html")
	t.Execute(w, data.GetListOfProducts())
}

func getUserPageFilteredByName(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("list-products.html")
	t.Execute(w, seacrh_filters.SearchByName(r))
}

func getUserPageFilteredByPrice(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("list-products.html")
	t.Execute(w, seacrh_filters.SearchByPrice(r))
}

func getUserPageFilteredByRating(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("list-products.html")
	t.Execute(w, seacrh_filters.SearchByRating(r))
}

func getUserPageAfterRating(w http.ResponseWriter, r *http.Request) {
	data.Products = seacrh_filters.RateProduct(r)
	t, _ := template.ParseFiles("list-products.html")
	t.Execute(w, data.GetListOfProducts())
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/sign-in":
		signInUser(w, r)
	case "/sign-up":
		signUpUser(w, r)
	case "/sign-in-form":
		getSignInPage(w)
	case "/sign-up-form":
		getSignUpPage(w)
	case "/sign-in-show-all":
		getUserPage(w)
	case "/sign-in-search-by-name":
		getUserPageFilteredByName(w, r)
	case "/sign-in-search-by-price":
		getUserPageFilteredByPrice(w, r)
	case "/sign-in-search-by-rating":
		getUserPageFilteredByRating(w, r)
	case "/sign-in-rating":
		getUserPageAfterRating(w, r)
	}
}

func main() {
	fmt.Println("http://localhost:8080")
	http.HandleFunc("/", userHandler)
	http.ListenAndServe(":8080", nil)
}
