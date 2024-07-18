package e2e

import (
	"chat/app/auth/config"
	"chat/app/auth/domain"

	"github.com/Nerzal/gocloak/v13"
)

func realm() gocloak.RealmRepresentation {
	return config.KcRealm()
}

func chatUser() domain.User {
	username := "tester"
	password := "password"
	email := "test@email.com"
	firstName := "automated"
	lastName := "tester"
	enabled := true
	emailVerified := true
	return domain.User{
		User: gocloak.User{
			Username:      &username,
			Enabled:       &enabled,
			Email:         &email,
			EmailVerified: &emailVerified,
			FirstName:     &firstName,
			LastName:      &lastName,
		},
		Password: &password,
	}
}

func users() []domain.User {
	return []domain.User{
		chatUser(),
	}
}

func clients() []gocloak.Client {
	return []gocloak.Client{
		config.KcClient(),
	}
}
