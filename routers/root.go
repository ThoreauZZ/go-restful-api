package routers

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/ThoreauZZ/go-restful-api/model"
	"github.com/ThoreauZZ/go-restful-api/handler"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello compose api!\n"))
}
func proxyApi() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		thoreau := model.User{
			Name:  "thoreau",
			Email: "zz.thoreau@gmail.com",
			Ext:   1,
		}
		handler.RstHandler(w, thoreau)
	}
	return http.HandlerFunc(fn)
}
func apiHandler() http.Handler {
	return http.Handler(handler.CheckRequest(proxyApi()))
}

func UserResource() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", HomeHandler)
	sub := router.PathPrefix("/api/v1").Subrouter()
	//sub.HandleFunc("/{path:[\\w/]+}", proxyApi)
	sub.Handle("/{path:[\\w/]+}", apiHandler())
	return router
}
