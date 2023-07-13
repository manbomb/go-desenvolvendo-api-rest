package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manbomb/go-curso-api-rest/controllers"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.CriaUmaNovaPersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
