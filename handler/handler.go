package handler

import (
	"go-context/usecase"
	"net/http"
)

type handler struct {
	ucase usecase.Usecase
}

func NewHandler(ucase usecase.Usecase) handler {
	return handler{
		ucase: ucase,
	}
}

func (h *handler) HandlerPrintUser(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	go h.ucase.PrintUser(ctx)
}
