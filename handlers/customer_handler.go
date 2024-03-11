package handlers

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/ortizdavid/go-nopain/conversion"
	"github.com/ortizdavid/go-null-dbvalues/entities"
	"github.com/ortizdavid/go-null-dbvalues/models"
)

type CustomerHandler struct {
}


func (cus *CustomerHandler) RegisterForm(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/register.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}


func (cus *CustomerHandler) Register(w http.ResponseWriter, r *http.Request) {
	var customerModel models.CustomerModel

	r.ParseForm()
	name := r.PostFormValue("name")
	age := r.PostFormValue("age")
	gender := r.PostFormValue("gender")
	height := r.PostFormValue("height")

	customer := entities.Customer{
		Id:     0,
		Name:   name,
		Age:    conversion.StringToIntOrNil(age),
		Gender: conversion.StringOrNil(gender),
		Height: conversion.StringToFloat32OrNil(height),
	}
	err := customerModel.Create(customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}
	fmt.Fprintf(w, "Customer registered sucessfully: %v", customer)
}