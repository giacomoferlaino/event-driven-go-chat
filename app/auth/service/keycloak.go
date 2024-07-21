package service

import (
	"chat/app/auth/domain"
	"chat/app/auth/repository"
)

func NewKeycloak(IdentityProviderRepository repository.IdentityProvider) Keycloak {
	return Keycloak{
		IdentityProviderRepository: IdentityProviderRepository,
	}
}

type Keycloak struct {
	IdentityProviderRepository repository.IdentityProvider
}

func (k Keycloak) Login(username string, password string) (domain.JWT, error) {
	return k.IdentityProviderRepository.GetJWT(username, password)
}
