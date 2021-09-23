package main //linkup main package
import "github.com/howkyle/linkup-server/server"

func main() {
	s := server.New()
	s.Init()
	// s.ConfigureServices()
	// s.ConfigureRoutes()
	s.Start()
}
