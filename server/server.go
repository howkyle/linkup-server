package server

//specifies server behaviors
type Server interface {
	//Configures routes and services
	Init()
	//Starts listening on configured port
	Start()
}

type Configurer interface {
	// //returns server configurations
	// Config() interface{}
}
