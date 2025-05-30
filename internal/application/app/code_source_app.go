package app

import (
	"onestep/internal/application/cmd"
)

// @Author CY Yan
// @Date 2025/5/29 16:11

// CodeSourceApp is an interface to operate code source
type CodeSourceApp interface {

	// Init code source to local, clone remote git repository and create fat, uat, pre, pro branches
	Init(cmd *cmd.CodeSourceCreateCmd) error
}
