package pkg

import "github.com/lishimeng/owl-messager/internal/db/model"

func TemplateInfoFromModel(src model.MessageTemplate) TemplateInfo {
	return TemplateInfo{
		Id:            src.Id,
		Code:          src.Code,
		Name:          src.Name,
		Category:      src.Category.String(),
		Body:          src.Body,
		Params:        src.Params,
		Provider:      src.Provider.String(),
		CloudTemplate: src.CloudTemplate,
		Description:   src.Description,
	}
}
