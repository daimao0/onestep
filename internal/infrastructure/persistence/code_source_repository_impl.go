package persistence

import (
	"gorm.io/gorm"
	"onestep/internal/domain/codesource/model"
	"onestep/internal/infrastructure/database"
	"onestep/internal/infrastructure/mapping"
	"onestep/internal/infrastructure/po"
)

// @Author CY Yan
// @Date 2025/5/30 16:20

type CodeSourceRepositoryImpl struct {
	db *gorm.DB
}

func NewCodeSourceRepositoryImpl() *CodeSourceRepositoryImpl {
	return &CodeSourceRepositoryImpl{
		db: database.GetDB(),
	}
}

func (p *CodeSourceRepositoryImpl) Insert(source *model.CodeSource) error {
	codeSourcePO := mapping.CodeSourceToCodeSourcePO(source)
	return p.db.Create(codeSourcePO).Error
}

func (p *CodeSourceRepositoryImpl) GetById(id int) *model.CodeSource {
	codeSourcePO := &po.CodeSourcePO{}
	p.db.First(codeSourcePO, id)
	return mapping.CodeSourcePOToCodeSource(codeSourcePO)
}
