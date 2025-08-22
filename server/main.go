package main

import (
	"fmt"
	"go-ai/provider"
	"go-ai/server/handlers"
	"log"
	"net/http"

	"go-ai/pkg/db"

	"github.com/go-chi/chi/v5"
	"github.com/samber/do"
)

func main() {

	fmt.Println("server")

	client, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	injector := do.New()

	do.ProvideValue(injector, client)

	provider.Provide(injector)

	r := chi.NewRouter()

	handlers.Register(injector, r)

	http.ListenAndServe("127.0.0.1:3332", r)

}
