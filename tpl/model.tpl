package model

import "gorm.io/gorm"

type T{{ .StructName }} struct {
	gorm.Model
}

func (m *T{{ .StructName }}) TableName() string {
    return "t_{{ .StructNameSnakeCase }}"
}
