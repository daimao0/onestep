package repository

import "onestep/internal/domain/codesource/model"

// @Author CY Yan
// @Date 2025/5/30 16:12

type CodeSourceRepository interface {

	// Insert code-source into persist storage
	Insert(source *model.CodeSource) error

	// GetById retrieves a code-source domain entity by its id.
	GetById(id int) *model.CodeSource
}
