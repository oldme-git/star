package v1

import "star/internal/model"

type List struct {
	Id               uint                   `json:"id"`
	Word             string                 `json:"word"`
	Definition       string                 `json:"definition"`
	ProficiencyLevel model.ProficiencyLevel `json:"proficiencyLevel"`
}
