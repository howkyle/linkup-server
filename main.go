package main //linkup main package
import "github.com/howkyle/linkup-server/server"

func main() {
	config := server.NewConfig("./config.yml")
	s := server.New(config)
	s.Init()
	s.Start()
}
