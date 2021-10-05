package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/howkyle/authman"
	"github.com/howkyle/linkup-server/user"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//wraps custom router implementations
type Router interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

//wraps db
type DB *gorm.DB

type Service interface{}
type ServiceContainer map[string]Service

type server struct {
	router      Router
	db          DB
	config      config
	userService user.Service
	// authManager uman.AuthManager
	// userManager uman.UserManager
}

//configures the servers database connection and application routes
func (s *server) Init() {
	s.db = initDB(s.config.DB)
	configServices(s)
	configRouter(s)
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
	err = db.AutoMigrate(user.User{})
	if err != nil {
		log.Println(err)
		panic("unable to run db migration: " + err.Error())
	}
	return db

}

//configures services
func configServices(s *server) {
	ur := user.NewRepository(s.db)
	authMan := authman.NewJWTAuthManager(s.config.ServerSecret, "pyt", "localhost", time.Minute*15)
	s.userService = user.NewService(ur, authMan)
}

//configures routes and sets server router
func configRouter(s *server) {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "welcome") })
	r.HandleFunc("/signup", user.SignupHandler(s.userService)).Methods("POST")
	r.HandleFunc("/login", user.LoginHandler(s.userService)).Methods("POST")

	s.router = r
}

//returns a new instance of a server with configurations
func New(c Configurer) Server {

	conf, ok := c.(config)
	if !ok {
		panic("invalid configuration")
	}
	return &server{config: conf}
}
