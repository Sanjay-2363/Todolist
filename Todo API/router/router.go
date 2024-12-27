package router

import (
	"net/http"
	"todolist/handler"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Routers Function is used to define the routes and start the server
func RoutersFunction() {

	router := mux.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://127.0.0.1:5500"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
	})

	handle := c.Handler(router)
	router.HandleFunc("/addtask", handler.AddTask).Methods(http.MethodPost)
	router.HandleFunc("/showalltask", handler.ShowAllTask).Methods(http.MethodGet)
	router.HandleFunc("/deletetask/{task_id}", handler.DeleteTask).Methods(http.MethodDelete)
	router.HandleFunc("/pendingtask", handler.Pendingtask).Methods(http.MethodGet)
	router.HandleFunc("/completedtask", handler.CompletedTask).Methods(http.MethodGet)
	router.HandleFunc("/checkbox/{task_id}", handler.CheckBox).Methods(http.MethodPut)
	http.ListenAndServe(":8080", handle)
}
