// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// WordsDao is the data access object for table words.
type WordsDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns WordsColumns // columns contains all the column names of Table for convenient usage.
}

// WordsColumns defines and stores column names for table words.
type WordsColumns struct {
	Id                 string //
	Uid                string //
	Word               string //
	Definition         string //
	ExampleSentence    string //
	ChineseTranslation string //
	Pronunciation      string //
	ProficiencyLevel   string //
	CreatedAt          string //
	UpdatedAt          string //
}

// wordsColumns holds the columns for table words.
var wordsColumns = WordsColumns{
	Id:                 "id",
	Uid:                "uid",
	Word:               "word",
	Definition:         "definition",
	ExampleSentence:    "example_sentence",
	ChineseTranslation: "chinese_translation",
	Pronunciation:      "pronunciation",
	ProficiencyLevel:   "proficiency_level",
	CreatedAt:          "created_at",
	UpdatedAt:          "updated_at",
}

// NewWordsDao creates and returns a new DAO object for table data access.
func NewWordsDao() *WordsDao {
	return &WordsDao{
		group:   "default",
		table:   "words",
		columns: wordsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *WordsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *WordsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *WordsDao) Columns() WordsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *WordsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *WordsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *WordsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
