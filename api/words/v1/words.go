package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type ProficiencyLevel uint

const (
	ProficiencyLevel1 ProficiencyLevel = iota + 1
	ProficiencyLevel2
	ProficiencyLevel3
	ProficiencyLevel4
	ProficiencyLevel5
)

type List struct {
	Id               uint             `json:"id"`
	Word             string           `json:"word"`
	Definition       string           `json:"definition"`
	ProficiencyLevel ProficiencyLevel `json:"proficiencyLevel"`
}

type CreateReq struct {
	g.Meta             `path:"words" method:"post" sm:"创建" tags:"单词"`
	Word               string           `json:"word" v:"required|length:1,100" dc:"单词"`
	Definition         string           `json:"definition" v:"required|length:1,300" dc:"单词定义"`
	ExampleSentence    string           `json:"example_sentence" v:"required|length:1,300" dc:"例句"`
	ChineseTranslation string           `json:"chinese_translation" v:"required|length:1,300" dc:"中文翻译"`
	Pronunciation      string           `json:"pronunciation" v:"required|length:1,100" dc:"发音"`
	ProficiencyLevel   ProficiencyLevel `json:"proficiency_level" v:"required|between:1,5" dc:"熟练度，1最低，5最高"`
}

type CreateRes struct {
}

type UpdateReq struct {
	g.Meta             `path:"words/{id}" method:"put" sm:"更新" tags:"单词"`
	Id                 uint             `json:"id" v:"required"`
	Word               string           `json:"word" v:"required|length:1,100" dc:"单词"`
	Definition         string           `json:"definition" v:"required|length:1,300" dc:"单词定义"`
	ExampleSentence    string           `json:"example_sentence" v:"required|length:1,300" dc:"例句"`
	ChineseTranslation string           `json:"chinese_translation" v:"required|length:1,300" dc:"中文翻译"`
	Pronunciation      string           `json:"pronunciation" v:"required|length:1,100" dc:"发音"`
	ProficiencyLevel   ProficiencyLevel `json:"proficiency_level" v:"required|between:1,5" dc:"熟练度，1最低，5最高"`
}

type UpdateRes struct {
}

type ListReq struct {
	g.Meta `path:"words" method:"get" sm:"列表" tags:"单词"`
	Word   string `json:"word" v:"length:1,100" dc:"模糊查询单词"`
	Page   int    `json:"page" v:"min:1" dc:"页码，默认1"`
	Size   int    `json:"size" v:"between:1,100" dc:"每页数量，默认10"`
}

type ListRes struct {
	List  []List `json:"list"`
	Total uint   `json:"total"`
}

type DetailReq struct {
	g.Meta `path:"words/{id}" method:"get" sm:"详情" tags:"单词"`
	Id     uint `json:"id" v:"required"`
}

type DetailRes struct {
	Id                 uint             `json:"id"`
	Word               string           `json:"word"`
	Definition         string           `json:"definition"`
	ExampleSentence    string           `json:"exampleSentence"`
	ChineseTranslation string           `json:"chineseTranslation"`
	Pronunciation      string           `json:"pronunciation"`
	ProficiencyLevel   ProficiencyLevel `json:"proficiencyLevel"`
	CreatedAt          *gtime.Time      `json:"createdAt"`
	UpdatedAt          *gtime.Time      `json:"updatedAt"`
}

type DeleteReq struct {
	g.Meta `path:"words/{id}" method:"delete" sm:"删除" tags:"单词"`
	Id     uint `json:"id" v:"required"`
}

type DeleteRes struct {
}
