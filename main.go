package main

import(
	"os"

	"github.com/Xiongzj5/cloud-go/routes"
	flag "github.com/spf13/pflag"
)

const ( 
	//Default host configuration
	defaultHost string = "localhost"
	//Default port configuration
	defaultPort string = "8080"
)

func main() {
	// Get the router from routes model
	r := routes.Router()

	//Trying to fetch the host configuration in environment variables
	prot = os.Getenv("PORT")
	if len(port) == 0{
		port = defaultPort
	}

	// Trying to fetch the host & port configuration in the flags
	hostFlag := flag.StringP("hostname", "h", defaultHost, "The host to deploy at")
	portFlag := flag.StringP("port", "p", defaultPort, "The port for listening")
	flag.Parse()
	if len(*hostFlag) != 0 {
		host = *hostFlag
	}
	if len(*portFlag) != 0 {
		port = *portFlag
	}
	// Run the server @host:port
	r.Run(host + ":" + port)
}