package v1

type List struct {
    Id               uint             `json:"id"`
    Word             string           `json:"word"`
    Definition       string           `json:"definition"`
    ProficiencyLevel ProficiencyLevel `json:"proficiencyLevel"`
}
