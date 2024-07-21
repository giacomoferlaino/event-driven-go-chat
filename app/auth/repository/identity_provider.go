package repository

import "chat/app/auth/domain"

type IdentityProvider interface {
	GetJWT(username string, password string) (domain.JWT, error)
}
