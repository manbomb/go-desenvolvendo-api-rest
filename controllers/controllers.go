package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manbomb/go-curso-api-rest/database"
	"github.com/manbomb/go-curso-api-rest/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	var p []models.Personalidade
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var p models.Personalidade
	database.DB.First(&p, id)
	json.NewEncoder(w).Encode(p)
}

func CriaUmaNovaPersonalidade(w http.ResponseWriter, r *http.Request) {
	var p models.Personalidade
	json.NewDecoder(r.Body).Decode(&p)
	fmt.Println(p)
	database.DB.Create(&p)
	json.NewEncoder(w).Encode(p)
}

func DeletaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade models.Personalidade
	database.DB.Delete(&personalidade, id)
	json.NewEncoder(w).Encode(personalidade)
}

func EditaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)
	json.NewDecoder(r.Body).Decode(&personalidade)
	database.DB.Save(&personalidade)
	json.NewEncoder(w).Encode(personalidade)
}
