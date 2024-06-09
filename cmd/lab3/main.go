package main

import (
	"aoop_lab1/cmd/lab2/db"
	"aoop_lab1/cmd/lab3/internal"
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
	handler := internal.NewHandler(ctx, queries)

	mux := http.NewServeMux()

	const staticPath = "cmd/lab3/static"
	if _, err := os.Stat(staticPath); err != nil {
		panic(err)
	}
	auth := internal.Auth{}
	mux.Handle("/", http.FileServer(http.Dir(staticPath)))
	mux.HandleFunc("POST /api/login", auth.Login)

	auth.HandleFuncAuth(mux, "GET /api/projects/{projectId}", handler.GetProject)
	auth.HandleFuncAuth(mux, "GET /api/projects/page/{page}/size/{size}/sort/{sort}", handler.GetProjects)
	auth.HandleFuncAuth(mux, "GET /api/projects/name/{name}/page/{page}/size/{size}/sort/{sort}", handler.GetProjectsByName)
	auth.HandleFuncAuth(mux, "POST /api/projects", handler.CreateProject)
	auth.HandleFuncAuth(mux, "PUT /api/projects/{projectId}", handler.UpdateProject)
	auth.HandleFuncAuth(mux, "DELETE /api/projects/{projectId}", handler.DeleteProject)

	log.Println("http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", mux))
}
