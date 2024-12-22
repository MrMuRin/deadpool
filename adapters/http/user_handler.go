package http

import (
	"deadpool/core/services"
)

type UserHandler struct {
	UserService *services.UserService
}
