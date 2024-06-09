package internal

import (
	"aoop_lab1/cmd/lab2/db"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

type Handler interface {
	GetProject(res http.ResponseWriter, req *http.Request)
	GetProjects(res http.ResponseWriter, req *http.Request)
	GetProjectsByName(res http.ResponseWriter, req *http.Request)
	CreateProject(res http.ResponseWriter, req *http.Request)
	UpdateProject(res http.ResponseWriter, req *http.Request)
	DeleteProject(res http.ResponseWriter, req *http.Request)
}

type handler struct {
	ctx     context.Context
	queries *db.Queries
}

func NewHandler(ctx context.Context, queries *db.Queries) Handler {
	return &handler{ctx: ctx, queries: queries}
}

func (h *handler) GetProject(res http.ResponseWriter, req *http.Request) {
	projectId, err := strconv.Atoi(req.PathValue("projectId"))
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	project, err := h.queries.GetProject(h.ctx, int32(projectId))
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	if err = json.NewEncoder(res).Encode(project); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *handler) GetProjects(res http.ResponseWriter, req *http.Request) {
	projectPageParams, err := h.parseProjectPageParams(req)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	projects, err := h.queries.GetProjects(h.ctx, db.GetProjectsParams{
		Ord:  projectPageParams.Ord,
		Page: int32(projectPageParams.Page),
		Size: int32(projectPageParams.Size),
	})
	if err = json.NewEncoder(res).Encode(projects); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func (h *handler) GetProjectsByName(res http.ResponseWriter, req *http.Request) {
	name := req.PathValue("name")
	if name == "" {
		http.Error(res, "name is required", http.StatusBadRequest)
		return
	}
	projectPageParams, err := h.parseProjectPageParams(req)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	projects, err := h.queries.GetProjectsByName(h.ctx, db.GetProjectsByNameParams{
		Name: name,
		Ord:  projectPageParams.Ord,
		Page: int32(projectPageParams.Page),
		Size: int32(projectPageParams.Size),
	})
	if err = json.NewEncoder(res).Encode(projects); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

type ProjectPageParams struct {
	Ord  string
	Page int
	Size int
}

func (h *handler) parseProjectPageParams(req *http.Request) (*ProjectPageParams, error) {
	var params ProjectPageParams
	var pageStr, sizeStr string
	var err error
	if pageStr = req.PathValue("page"); pageStr == "" {
		return nil, errors.New("page parameter is required")
	}
	if sizeStr = req.PathValue("size"); sizeStr == "" {
		return nil, errors.New("size parameter is required")
	}
	if params.Ord = req.PathValue("sort"); params.Ord == "" {
		return nil, errors.New("sort parameter is required")
	}
	if params.Page, err = strconv.Atoi(pageStr); err != nil {
		return nil, err
	}
	if params.Size, err = strconv.Atoi(sizeStr); err != nil {
		return nil, err
	}
	return &params, nil
}

func (h *handler) CreateProject(res http.ResponseWriter, req *http.Request) {
	var createProjectParams db.CreateProjectParams
	if err := json.NewDecoder(req.Body).Decode(&createProjectParams); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	if _, err := h.queries.CreateProject(h.ctx, createProjectParams); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func (h *handler) UpdateProject(res http.ResponseWriter, req *http.Request) {
	var updateProjectParam db.UpdateProjectParams
	if err := json.NewDecoder(req.Body).Decode(&updateProjectParam); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.queries.UpdateProject(h.ctx, updateProjectParam); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func (h *handler) DeleteProject(res http.ResponseWriter, req *http.Request) {
	projectId, err := strconv.Atoi(req.PathValue("projectId"))
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	if err = h.queries.DeleteProject(h.ctx, int32(projectId)); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}
