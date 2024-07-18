package graph

import (
	"chat/app/auth/config"
	"chat/app/auth/repository"
	"chat/app/auth/service"
)

var ()

type diContainer struct {
	keycloakRepository repository.Keycloak
	keycloakService    service.Keycloak
}

func newDIContainer() diContainer {
	keycloakRepository := repository.NewKeycloak(config.KcUrl(), *config.KcRealm().Realm)
	keycloakService := service.NewKeycloak(keycloakRepository)

	return diContainer{
		keycloakRepository: keycloakRepository,
		keycloakService:    keycloakService,
	}
}
