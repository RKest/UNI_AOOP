package main

import (
	"aoop_lab1/cmd/lab2/db"
	"context"
	"github.com/jackc/pgx/v5"
	"log"
	"net/http"
	"os"
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

	const staticPath = "cmd/lab3/static"
	if _, err := os.Stat(staticPath); err != nil {
		panic(err)
	}
	mux.Handle("/", http.FileServer(http.Dir(staticPath)))

	mux.HandleFunc("POST /api/login", Auth)

	HandleFuncAuth(mux, "GET /api/projects/{projectId}", handler.GetProject)
	HandleFuncAuth(mux, "GET /api/projects/page/{page}/size/{size}/sort/{sort}", handler.GetProjects)
	HandleFuncAuth(mux, "GET /api/projects/name/{name}/page/{page}/size/{size}/sort/{sort}", handler.GetProjectsByName)
	HandleFuncAuth(mux, "POST /api/projects", handler.CreateProject)
	HandleFuncAuth(mux, "PUT /api/projects/{projectId}", handler.UpdateProject)
	HandleFuncAuth(mux, "DELETE /api/projects/{projectId}", handler.DeleteProject)

	log.Println("http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", mux))
}
