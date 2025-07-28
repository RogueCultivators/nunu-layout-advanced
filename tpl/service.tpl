package service

import (
    "context"
	"{{ .ProjectName }}/internal/model"
	"{{ .ProjectName }}/internal/repository"
)

type S{{ .StructName }}Service interface {
	Get{{ .StructName }}(ctx context.Context, id int64) (*model.T{{ .StructName }}, error)
}
func NewS{{ .StructName }}Service(
    service *Service,
    r{{ .StructName }}Repository repository.R{{ .StructName }}Repository,
) S{{ .StructName }}Service {
	return &s{{ .StructName }}Service{
		Service:        service,
		r{{ .StructName }}Repository: r{{ .StructName }}Repository,
	}
}

type s{{ .StructName }}Service struct {
	*Service
	r{{ .StructName }}Repository repository.R{{ .StructName }}Repository
}

func (s *s{{ .StructName }}Service) Get{{ .StructName }}(ctx context.Context, id int64) (*model.T{{ .StructName }}, error) {
	return s.r{{ .StructName }}Repository.Get{{ .StructName }}(ctx, id)
}
