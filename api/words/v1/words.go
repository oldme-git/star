package v1

import "github.com/gogf/gf/v2/frame/g"

type CreateReq struct {
	g.Meta             `path:"words" method:"post" sm:"创建" tags:"单词"`
	Word               string `json:"word" v:"required|length:1,100" dc:"单词"`
	Definition         string `json:"definition" v:"required|length:1,300" dc:"单词定义"`
	ExampleSentence    string `json:"example_sentence" v:"required|length:1,300" dc:"例句"`
	ChineseTranslation string `json:"chinese_translation" v:"required|length:1,300" dc:"中文翻译"`
	Pronunciation      string `json:"pronunciation" v:"required|length:1,100" dc:"发音"`
	ProficiencyLevel   uint   `json:"proficiency_level" v:"required|between:1,5" dc:"熟练度，1最低，5最高"`
}

type CreateRes struct {
}

type UpdateReq struct {
	g.Meta             `path:"words/{id}" method:"put" sm:"更新" tags:"单词"`
	Id                 uint   `json:"id" v:"required"`
	Word               string `json:"word" v:"required|length:1,100" dc:"单词"`
	Definition         string `json:"definition" v:"required|length:1,300" dc:"单词定义"`
	ExampleSentence    string `json:"example_sentence" v:"required|length:1,300" dc:"例句"`
	ChineseTranslation string `json:"chinese_translation" v:"required|length:1,300" dc:"中文翻译"`
	Pronunciation      string `json:"pronunciation" v:"required|length:1,100" dc:"发音"`
	ProficiencyLevel   uint   `json:"proficiency_level" v:"required|between:1,5" dc:"熟练度，1最低，5最高"`
}

type UpdateRes struct {
}
