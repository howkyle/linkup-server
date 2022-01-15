package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/howkyle/authman"
	"github.com/howkyle/linkup-server/event"
	"github.com/howkyle/linkup-server/invitation"
	"github.com/howkyle/linkup-server/user"
	"github.com/howkyle/linkup-server/validation"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//wraps custom router implementations
type Router interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type mongodb *mongo.Database

type Service interface{}
type ServiceContainer map[string]Service

type server struct {
	router        Router
	db            mongodb
	config        config
	userService   user.Service
	eventService  event.Service
	inviteService invitation.Service
	authManager   authman.AuthManager
}

//configures the servers database connection and application routes
func (s *server) Init() {
	s.db = initMongo(s.config.DB, s.config.DB_Name)
	configServices(s)
	configRouter(s)
}

//starts the server on the specified port
func (s *server) Start() {
	log.Printf("starting server on port %v", s.config.ServerPort)
	log.Fatal(http.ListenAndServe(s.config.ServerPort, s.router))
}

//initializes mongo db
func initMongo(connection string, database string) mongodb {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	opts := options.Client().ApplyURI(connection)
	log.Println(" creating connection to mongo")
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Println("failed to connect to mongo")
		panic(fmt.Sprintf("failed to create connection to mongo: %v", err))
	}
	log.Println("testing connection to mongo")
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Println("connection test failed")
		panic("failed to verify connection to mongo")
	}
	log.Println("connection verified")

	return client.Database(database)
}

//configures services
func configServices(s *server) {
	ur := user.NewMongoRepository(s.db)
	er := event.NewMongoRepository(s.db)
	ir := invitation.NewMongoRepository(s.db)
	s.authManager = authman.NewJWTAuthManager(s.config.ServerSecret, "pyt", "localhost", time.Minute*15)
	s.userService = user.NewService(ur, s.authManager)
	s.eventService = event.NewService(er)
	s.inviteService = invitation.NewService(ir)
}

//configures routes and sets server router
func configRouter(s *server) {
	v := validation.NewValidator()
	r := mux.NewRouter()
	r.HandleFunc("/signup", user.SignupHandler(s.userService, v)).Methods("POST")
	r.HandleFunc("/login", user.LoginHandler(s.userService, v)).Methods("POST")
	r.HandleFunc("/event", s.authManager.Filter(event.NewEventHandler(s.eventService, v))).Methods("POST")
	r.HandleFunc("/invitation", s.authManager.Filter(invitation.NewInvitationHandler(s.inviteService, s.eventService))).Methods("POST")
	r.HandleFunc("/invitations", s.authManager.Filter(invitation.GetInvitationsHandler(s.inviteService))).Methods("GET")
	r.HandleFunc("/invitation/{id}/accept", s.authManager.Filter(invitation.AcceptHandler(s.inviteService))).Methods("GET")
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
