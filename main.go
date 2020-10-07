package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/matthausen/go_elastic/router"
)

func main() {}
	r := router.Router()
	fmt.Println("App running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
