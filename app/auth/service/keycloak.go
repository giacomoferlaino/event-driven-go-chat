package service

import "chat/app/auth/repository"

func NewKeycloak(IdentityProviderRepository repository.IdentityProvider) KeycloakAuth {
	return KeycloakAuth{
		IdentityProviderRepository: IdentityProviderRepository,
	}
}

type KeycloakAuth struct {
	IdentityProviderRepository repository.IdentityProvider
}

func (k KeycloakAuth) Login(username string, password string) (accessToken string, err error) {
	return k.IdentityProviderRepository.GetAccessToken(username, password)
}
