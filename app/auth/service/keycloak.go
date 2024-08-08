package service

import (
	"chat/app/auth/domain"
	"chat/app/auth/repository"
)

func NewKeycloak(IdentityProviderRepository repository.IdentityProvider) Keycloak {
	return Keycloak{
		identityProviderRepository: IdentityProviderRepository,
	}
}

type Keycloak struct {
	identityProviderRepository repository.IdentityProvider
}

func (k Keycloak) Login(username string, password string) (domain.JWT, error) {
	return k.identityProviderRepository.GetJWT(username, password)
}
