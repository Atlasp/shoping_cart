package handler

import "revel_systems_shopping/repository"

type handler struct {
	repo repository.Repository
}

func NewHandler(repo repository.Repository) *handler {
	return &handler{repo: repo}
}
