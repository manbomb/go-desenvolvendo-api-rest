package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manbomb/go-curso-api-rest/controllers"
	"github.com/manbomb/go-curso-api-rest/middlewares"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middlewares.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.CriaUmaNovaPersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletaUmaPersonalidade).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", controllers.EditaUmaPersonalidade).Methods("Put")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
