package main

import (
	"flag"
	"fmt"
	"log"
	"log/slog"
	"net/http"
)

func main() {
	log.SetFlags(log.Lshortfile)

	portPtr := flag.Int("port", 8080, "The port to use for the HTTP server")
	exposePtr := flag.Bool("expose", false, "Whatever to expose the service to the network or not")

	flag.Parse()

	var bind string

	if *exposePtr {
		bind = fmt.Sprintf("0.0.0.0:%d", *portPtr)
	} else {
		bind = fmt.Sprintf("127.0.0.1:%d", *portPtr)
	}

	slog.Info(fmt.Sprintf("Serving HTTP on http://%s", bind))

	http.ListenAndServe(bind, nil)
}
