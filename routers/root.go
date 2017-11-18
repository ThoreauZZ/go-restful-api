package routers

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"github.com/ThoreauZZ/go-restful-api/model"
	"github.com/ThoreauZZ/go-restful-api/handler"
)
func service(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Host, r.URL)
	thoreau := model.User{
		Name:  "thoreau",
		Email: "zz.thoreau@gmail.com",
		Ext:   1,
	}
	handler.RstHandler(w, thoreau)
}

func UserResource() *mux.Router{
	router := mux.NewRouter().StrictSlash(false)
	sub := router.PathPrefix("/api/v1").Subrouter()
	sub.HandleFunc("/{path}", service)
	return router
}
