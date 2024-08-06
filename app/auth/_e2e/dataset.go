package e2e

import (
	"chat/app/auth/config"
	"chat/pkg/test/e2esdk"

	"github.com/Nerzal/gocloak/v13"
)

func realm() *e2esdk.Realm {
	return &e2esdk.Realm{
		Roles:               &[]gocloak.Role{*chatRealmRole()},
		RealmRepresentation: config.KcRealm(),
	}
}

func chatRealmRole() *gocloak.Role {
	name := "chat-role"
	description := "${role_chat-role}"
	return &gocloak.Role{
		Name:        &name,
		Description: &description,
	}

}

func chatUser() *e2esdk.User {
	username := "tester"
	password := "password"
	email := "test@email.com"
	firstName := "automated"
	lastName := "tester"
	enabled := true
	emailVerified := true
	realmRoles := []string{
		*chatRealmRole().Name,
	}
	return &e2esdk.User{
		User: gocloak.User{
			Username:      &username,
			Enabled:       &enabled,
			Email:         &email,
			EmailVerified: &emailVerified,
			FirstName:     &firstName,
			LastName:      &lastName,
			RealmRoles:    &realmRoles,
		},
		Password:  &password,
		RealmName: realm().Realm,
	}
}

func users() *[]*e2esdk.User {
	return &[]*e2esdk.User{
		chatUser(),
	}
}

func client() *e2esdk.Client {
	return &e2esdk.Client{
		Client:    config.KcClient(),
		RealmName: realm().Realm,
	}
}
