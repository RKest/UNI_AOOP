package main

import (
	"aoop_lab1/cmd/lab2/db"
	"context"
	"github.com/jackc/pgx/v5"
	"log"
	"net/http"
)

func main() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "postgres://postgres:qwer1234@localhost:5432/projects")
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	queries := db.New(conn)
	handler := NewHandler(ctx, queries)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/projects/{projectId}", handler.GetProject)
	mux.HandleFunc("GET /api/projects/page/{page}/size/{size}/sort/{sort}", handler.GetProjects)
	mux.HandleFunc("GET /api/projects/name/{name}/page/{page}/size/{size}/sort/{sort}", handler.GetProjectsByName)
	mux.HandleFunc("POST /api/projects", handler.CreateProject)
	mux.HandleFunc("PUT /api/projects/{projectId}", handler.UpdateProject)
	mux.HandleFunc("DELETE /api/projects/{projectId}", handler.DeleteProject)

	log.Println("http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", mux))
}
