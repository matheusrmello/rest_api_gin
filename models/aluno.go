package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type AlunosDados struct {
	gorm.Model
	Nome string `json:"nome" validate:"nonzero"`
	CPF  string `json:"CPF" validate:"len=11, regexp=^[0-9]*$"`
	RG   string `json:"RG" validate:"len=9, regexp=^[0-9]*$"`
}

func ValidaDados(aluno *AlunosDados) error {
	if err := validator.Validate(aluno); err != nil{
		return err
	}
	return nil
}