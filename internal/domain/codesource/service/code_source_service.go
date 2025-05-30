package service

import "onestep/internal/domain/codesource/model"

// @Author CY Yan
// @Date 2025/5/28 18:34

// CodeSourceService operation git code source
type CodeSourceService interface {

	// Init the code repository and creates the branches: fat, uat, pre, pro.
	Init(source *model.CodeSource) error
}
