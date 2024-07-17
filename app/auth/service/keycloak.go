package service

import "chat/app/auth/repository"

func NewKeycloak(keycloakRepository repository.Keycloak) KeycloakAuth {
	return KeycloakAuth{
		keycloakRepository: keycloakRepository,
	}
}

type KeycloakAuth struct {
	keycloakRepository repository.Keycloak
}

func (k KeycloakAuth) Login(username string, password string) (string, error) {
	return k.keycloakRepository.GetAccessToken(username, password)
}
