package server

type Server interface {
	// ConfigureServices()
	// ConfigureRoutes()
	Init()
	Start()
}

type Configurer interface {
}

func New() Server {
	return nil
}

//impl

type server struct {
}
