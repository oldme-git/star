// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Words is the golang structure for table words.
type Words struct {
	Id                 uint        `json:"id"                 orm:"id"                  ` //
	Uid                uint        `json:"uid"                orm:"uid"                 ` //
	Word               string      `json:"word"               orm:"word"                ` //
	Definition         string      `json:"definition"         orm:"definition"          ` //
	ExampleSentence    string      `json:"exampleSentence"    orm:"example_sentence"    ` //
	ChineseTranslation string      `json:"chineseTranslation" orm:"chinese_translation" ` //
	Pronunciation      string      `json:"pronunciation"      orm:"pronunciation"       ` //
	ProficiencyLevel   uint        `json:"proficiencyLevel"   orm:"proficiency_level"   ` //
	CreatedAt          *gtime.Time `json:"createdAt"          orm:"created_at"          ` //
	UpdatedAt          *gtime.Time `json:"updatedAt"          orm:"updated_at"          ` //
}
