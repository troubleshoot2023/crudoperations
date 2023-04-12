package cmd

import (
	"log"
	"net/http"

	//"github.com/drone/routes"
	"github.com/gorilla/mux"
	"github.com/troubleshoot2023/DEV2/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
