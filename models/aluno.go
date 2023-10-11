package models

import "gorm.io/gorm"

type AlunosDados struct {
	gorm.Model
	Nome string `json:"nome"`
	CPF  string `json:"CPF"`
	RG   string `json:"RG"`
}
