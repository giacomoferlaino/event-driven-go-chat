package e2e

import (
	"chat/app/auth/_e2e/e2esdk"
	"chat/app/auth/config"

	"github.com/Nerzal/gocloak/v13"
)

func realm() gocloak.RealmRepresentation {
	return config.KcRealm()
}

func chatUser() e2esdk.User {
	username := "tester"
	password := "password"
	email := "test@email.com"
	firstName := "automated"
	lastName := "tester"
	enabled := true
	emailVerified := true
	return e2esdk.User{
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

func users() []e2esdk.User {
	return []e2esdk.User{
		chatUser(),
	}
}

func client() gocloak.Client {
	return config.KcClient()
}
