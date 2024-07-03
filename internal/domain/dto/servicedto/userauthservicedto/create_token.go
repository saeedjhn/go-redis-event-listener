package userauthservicedto

import (
	"github.com/saeedjhn/go-redis-event-listener/internal/domain/entity"
)

type CreateTokenRequest struct {
	User entity.User
}

type CreateTokenResponse struct {
	Token string
}
