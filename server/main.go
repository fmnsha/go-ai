package main

import (
	"fmt"
	"go-ai/pkg/services/member"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {

	fmt.Println("server")

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/member", func(w http.ResponseWriter, r *http.Request) {

		svcs := member.NewMember()

		result, _ := svcs.AddMember(r.Context(), "test")

		w.Write([]byte(result))
	})
	http.ListenAndServe(":3000", r)

}
