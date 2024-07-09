package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/mu10x/gqlgen-sqlc-example/gqlgen"
	"github.com/mu10x/gqlgen-sqlc-example/pg"
)

func main() {
	dbURI := fmt.Sprintf("host=%s port=%d user=%s database=%s sslmode=disable", "localhost", 5432, "app-default", "gqlgen_sqlc_example_db")
	pool, err := pg.Open(dbURI)
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	repo := pg.NewRepository(pool)

	mux := http.NewServeMux()
	mux.Handle("/", gqlgen.NewPlaygroundHandler("/query"))
	mux.Handle("/query", gqlgen.NewHandler(repo))

	port := ":8080"
	fmt.Fprintf(os.Stdout, "🚀 Server ready at http://localhost%s\n", port)
	fmt.Fprintln(os.Stderr, http.ListenAndServe(port, mux))
}
