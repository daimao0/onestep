package mapping

import (
	"onestep/internal/domain/codesource/model"
	"onestep/internal/infrastructure/po"
)

// @Author CY Yan
// @Date 2025/5/30 16:24

func CodeSourcePOToCodeSource(po *po.CodeSourcePO) *model.CodeSource {
	return &model.CodeSource{
		Id:                  po.Id,
		CreatedAt:           po.CreatedAt,
		UpdatedAt:           po.UpdatedAt,
		Desc:                po.Desc,
		RemoteURL:           po.RemoteURL,
		PersonalAccessToken: po.PersonalAccessToken,
		Username:            po.Username,
		Password:            po.PersonalAccessToken,
	}
}

func CodeSourceToCodeSourcePO(source *model.CodeSource) *po.CodeSourcePO {
	return &po.CodeSourcePO{
		Model: po.Model{
			Id:        source.Id,
			CreatedAt: source.CreatedAt,
			UpdatedAt: source.UpdatedAt,
		},
		Desc:                source.Desc,
		RemoteURL:           source.RemoteURL,
		PersonalAccessToken: source.PersonalAccessToken,
		Username:            source.Username,
	}
}
