package service

import "chat/app/auth/repository"

func NewKeycloak(IdentityProviderRepository repository.IdentityProvider) Keycloak {
	return Keycloak{
		IdentityProviderRepository: IdentityProviderRepository,
	}
}

type Keycloak struct {
	IdentityProviderRepository repository.IdentityProvider
}

func (k Keycloak) Login(username string, password string) (accessToken string, err error) {
	return k.IdentityProviderRepository.GetAccessToken(username, password)
}
