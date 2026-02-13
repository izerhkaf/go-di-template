package service

import "template/internal/config"

type TemplateService interface {
}

type templateService struct {
	Config *config.Config
}

func NewTemplateService(
	config *config.Config,
) TemplateService {
	return &templateService{
		Config: config,
	}
}
