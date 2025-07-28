package handler

import (
	"github.com/gin-gonic/gin"
	"{{ .ProjectName }}/internal/service"
)

type H{{ .StructName }}Handler struct {
	*Handler
	s{{ .StructName }}Service service.S{{ .StructName }}Service
}

func NewH{{ .StructName }}Handler(
    handler *Handler,
    s{{ .StructName }}Service service.S{{ .StructName }}Service,
) *H{{ .StructName }}Handler {
	return &H{{ .StructName }}Handler{
		Handler:      handler,
		s{{ .StructName }}Service: s{{ .StructName }}Service,
	}
}

func (h *H{{ .StructName }}Handler) Get{{ .StructName }}(ctx *gin.Context) {

}
