package model

type ProficiencyLevel uint

const (
	ProficiencyLevel1 ProficiencyLevel = 1
	ProficiencyLevel2 ProficiencyLevel = 2
	ProficiencyLevel3 ProficiencyLevel = 3
	ProficiencyLevel4 ProficiencyLevel = 4
	ProficiencyLevel5 ProficiencyLevel = 5
)

type WordInput struct {
	Uid                uint
	Word               string
	Definition         string
	ExampleSentence    string
	ChineseTranslation string
	Pronunciation      string
	ProficiencyLevel   ProficiencyLevel
}
