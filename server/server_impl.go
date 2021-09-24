package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//wraps custom router implementations
type Router interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

//wraps db
type DB interface{}

type server struct {
	router Router
	db     DB
	config config
}

//configures the servers database connection and application routes
func (s *server) Init() {
	s.db = initDB(s.config.DB)
	s.router = configRouter()
}

//starts the server on the specified port
func (s *server) Start() {
	log.Printf("starting server on port %v", s.config.ServerPort)
	log.Fatal(http.ListenAndServe(s.config.ServerPort, s.router))
}

//connects to the database
func initDB(connection string) DB {
	log.Println("connecting to db")
	db, err := gorm.Open(mysql.Open(connection))
	if err != nil {
		log.Println(err)
		panic("unable to connect to database")
	}
	log.Printf("connected to db: %v\n", db.Name())

	log.Println("running db migrations")
	err = db.AutoMigrate()
	if err != nil {
		log.Println(err)
		panic("unable to run db migration: " + err.Error())
	}
	return db

}

//configures routes and returns pointer to router
func configRouter() Router {

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "welcome") })

	return r
}

//returns a new instance of a server with configurations
func New(c Configurer) server {

	conf, ok := c.(config)
	if !ok {
		panic("invalid configuration")
	}
	return server{config: conf}
}
