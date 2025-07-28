package repository

import (
    "context"
	"{{ .ProjectName }}/internal/model"
	"errors"
	"gorm.io/gorm"
	v1 "{{ .ProjectName }}/api/v1"
)

type R{{ .StructName }}Repository interface {
	Get{{ .StructName }}(ctx context.Context, id int64) (*model.T{{ .StructName }}, error)
	Create{{ .StructName }}(ctx context.Context, {{ .StructNameLowerFirst }} *model.T{{ .StructName }}) error
	Update{{ .StructName }}(ctx context.Context, {{ .StructNameLowerFirst }} *model.T{{ .StructName }}) error
	Delete{{ .StructName }}(ctx context.Context, {{ .StructNameLowerFirst }} *model.T{{ .StructName }}) error
}

func NewR{{ .StructName }}Repository(
	repository *Repository,
) R{{ .StructName }}Repository {
	return &r{{ .StructName }}Repository{
		Repository: repository,
	}
}

type r{{ .StructName }}Repository struct {
	*Repository
}

func (r *r{{ .StructName }}Repository) Get{{ .StructName }}(ctx context.Context, id int64) (*model.T{{ .StructName }}, error) {
	var {{ .StructNameLowerFirst }} model.T{{ .StructName }}
	if err := r.DB(ctx).Where("id = ?", id).First(&{{ .StructNameLowerFirst }}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, v1.ErrNotFound
		}
		return nil, err
	}
	return &{{ .StructNameLowerFirst }}, nil
}

func (r *r{{ .StructName }}Repository) Create{{ .StructName }}(ctx context.Context, {{ .StructNameLowerFirst }} *model.T{{ .StructName }}) error {
	if err := r.DB(ctx).Create({{ .StructNameLowerFirst }}).Error; err != nil {
		return err
	}
	return nil
}

func (r *r{{ .StructName }}Repository) Update{{ .StructName }}(ctx context.Context, {{ .StructNameLowerFirst }} *model.T{{ .StructName }}) error {
	if err := r.DB(ctx).Save({{ .StructNameLowerFirst }}).Error; err != nil {
		return err
	}
	return nil
}
func (r *r{{ .StructName }}Repository) Delete{{ .StructName }}(ctx context.Context, {{ .StructNameLowerFirst }} *model.T{{ .StructName }} ) error {
	if err := r.DB(ctx).Where({{ .StructNameLowerFirst }}).Delete(&model.T{{ .StructName }}{}).Error; err != nil {
		return err
	}
	return nil
}
