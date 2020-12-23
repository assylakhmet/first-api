package services

import (
	"first-api/configuration"
	logger "git.homebank.kz/HomeBank/elastic-logger.v2"
)

//Handler Вебсервис
type Handler struct {
	Config *configuration.Configuration
	Logger *logger.Logger
}

//NewHandler Создание нового объекта для веб сервисов
func NewHandler(config *configuration.Configuration, log *logger.Logger) *Handler {
	handler := Handler{
		Config: config,
		Logger: log,
	}
	return &handler
}
