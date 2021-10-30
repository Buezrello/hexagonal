package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"example.com/hexagonal/domain"
	"example.com/hexagonal/service"
	"github.com/gorilla/mux"
)

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" ||
		os.Getenv("SERVER_PORT") == "" {

		log.Fatal("Environment variables not defined...")
	}
}

func Start() {
	router := mux.NewRouter()

	sanityCheck()

	// writing
	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	// starting server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))

	// run with (if env vars not exported)

	// SERVER_ADDRESS=localhost SERVER_PORT=8000 go run main.go

	// or export one time env vars

	// export SERVER_ADDRESS=localhost
	// export SERVER_PORT=8000
	// export DB_USER=root
	// export DB_PASSWD=admin
	// export DB_ADDR=localhost
	// export DB_PORT=3306
	// export DB_NAME=banking

	// go run main.go
}
