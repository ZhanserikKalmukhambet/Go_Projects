package seacrh_filters

import (
	"assignment_1/data"
	"net/http"
	"strconv"
	"strings"
)

func SearchByName(r *http.Request) []data.Product {
	pattern := r.FormValue("name")

	var filteredProducts []data.Product

	for _, prod := range data.GetListOfProducts() {
		if strings.Contains(prod.Name, pattern) {
			filteredProducts = append(filteredProducts, prod)
		}
	}

	return filteredProducts
}

func SearchByPrice(r *http.Request) []data.Product {
	min, _ := strconv.ParseFloat(r.FormValue("min"), 10)
	max, _ := strconv.ParseFloat(r.FormValue("max"), 10)

	var filteredProducts []data.Product

	for _, prod := range data.GetListOfProducts() {
		if min <= prod.Price && prod.Price <= max {
			filteredProducts = append(filteredProducts, prod)
		}
	}
	return filteredProducts
}

func SearchByRating(r *http.Request) []data.Product {
	min, _ := strconv.ParseFloat(r.FormValue("min"), 10)
	max, _ := strconv.ParseFloat(r.FormValue("max"), 10)

	var filteredProducts []data.Product

	for _, prod := range data.GetListOfProducts() {
		if min <= prod.Rating && prod.Rating <= max {
			filteredProducts = append(filteredProducts, prod)
		}
	}
	return filteredProducts
}

func RateProduct(r *http.Request) []data.Product {
	id, _ := strconv.ParseInt(r.FormValue("id"), 10, 64)
	rating, _ := strconv.ParseFloat(r.FormValue("rating"), 64)

	var filteredProducts []data.Product

	for _, prod := range data.GetListOfProducts() {
		if prod.ID == id {
			prod.SizeOfRatings++
			prod.Rating = (prod.Rating + rating) / float64(prod.SizeOfRatings)
		}
		filteredProducts = append(filteredProducts, prod)
	}

	return filteredProducts
}
