package helloworld

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pedropina/base_api/pkg/server"
)

func Run(addr string) {
	app := server.New()

	app.AddHandler("GET /{name}", func(w http.ResponseWriter, r *http.Request) {
		name := r.PathValue("name")
		fmt.Fprintf(w, "<a>Hello <b>%v</b> </a> \n", name)
	})
	log.Fatal(app.InitServer(addr))
}
