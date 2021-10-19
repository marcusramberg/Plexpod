package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/routers/gorillamux"
)

func main() {
	ctx := context.Background()
	loader := openapi3.NewLoader()
	doc, _ := loader.LoadFromFile("sirikit-cloud-media.json")
	err := doc.Validate(loader.Context)
	if err != nil {
		fmt.Printf("Failed to validate spec: %s\n", err)
		os.Exit(1)
	}
	router, _ := gorillamux.NewRouter(doc)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Path: %s!", r.URL.Path[1:])
		_, _, _ = router.FindRoute(r)
	})

	http.ListenAndServe(":8080", nil)
	// Do something with route.Operation
}
