package words

import (
    "context"

    v1 "star/api/words/v1"
    "star/internal/dao"
    "star/internal/model/do"
    "star/internal/model/entity"
)

type Words struct {
}

type CreateInput struct {
    Uid                uint
    Word               string
    Definition         string
    ExampleSentence    string
    ChineseTranslation string
    Pronunciation      string
    ProficiencyLevel   v1.ProficiencyLevel
}

func New() *Words {
    return &Words{}
}

func (w *Words) Create(ctx context.Context, in CreateInput) error {
    _, err := dao.Words.Ctx(ctx).Data(do.Words{
        Uid:                in.Uid,
        Word:               in.Word,
        Definition:         in.Definition,
        ExampleSentence:    in.ExampleSentence,
        ChineseTranslation: in.ChineseTranslation,
        Pronunciation:      in.Pronunciation,
        ProficiencyLevel:   in.ProficiencyLevel,
    }).Insert()
    if err != nil {
        return err
    }
    return nil
}

type UpdateInput struct {
    Word               string
    Definition         string
    ExampleSentence    string
    ChineseTranslation string
    Pronunciation      string
    ProficiencyLevel   v1.ProficiencyLevel
}

func (w *Words) Update(ctx context.Context, id uint, in UpdateInput) error {
    var (
        cls = dao.Words.Columns()
        orm = dao.Words.Ctx(ctx)
    )
    orm = orm.Data(do.Words{
        Word:               in.Word,
        Definition:         in.Definition,
        ExampleSentence:    in.ExampleSentence,
        ChineseTranslation: in.ChineseTranslation,
        Pronunciation:      in.Pronunciation,
        ProficiencyLevel:   in.ProficiencyLevel,
    }).Where(cls.Id, id)
    _, err := orm.Update()
    if err != nil {
        return err
    }
    return nil
}

type ListInput struct {
    Uid  uint
    Word string
    Page int
    Size int
}

func (w *Words) List(ctx context.Context, in ListInput) (list []entity.Words, total int, err error) {
    // 对于查询初始值的处理
    if in.Page == 0 {
        in.Page = 1
    }
    if in.Size == 0 {
        in.Size = 15
    }

    var (
        cls = dao.Words.Columns()
        orm = dao.Words.Ctx(ctx)
    )
    // 组成查询链
    if in.Uid > 0 {
        orm = orm.Where(cls.Uid, in.Uid)
    }

    // 模糊查询
    if len(in.Word) != 0 {
        orm = orm.WhereLike(cls.Word, "%"+in.Word+"%")
    }
    orm = orm.OrderDesc(cls.CreatedAt).OrderDesc(cls.Id).Page(in.Page, in.Size)
    if err = orm.ScanAndCount(&list, &total, true); err != nil {
        return
    }
    return
}

func (w *Words) Detail(ctx context.Context, uid, id uint) (word *entity.Words, err error) {
    var (
        cls = dao.Words.Columns()
        orm = dao.Words.Ctx(ctx)
    )
    orm = orm.Where(cls.Id, id)
    if uid > 0 {
        orm = orm.Where(cls.Uid, uid)
    }
    err = orm.Scan(&word)
    return
}

func (w *Words) Delete(ctx context.Context, uid, id uint) (err error) {
    var (
        cls = dao.Words.Columns()
        orm = dao.Words.Ctx(ctx)
    )
    orm = orm.Where(cls.Id, id)
    if uid > 0 {
        orm = orm.Where(cls.Uid, uid)
    }
    _, err = orm.Delete()
    return
}
