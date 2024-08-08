package e2esdk

import (
	"chat/app/auth/config"

	"github.com/Nerzal/gocloak/v13"
)

func DefaultSeed() *KeycloakData {
	return &KeycloakData{
		Realm:  DefaultRealm(),
		Client: DefaultClient(),
		Users:  Users(),
	}
}

func DefaultClient() *Client {
	return &Client{
		Client:    config.KcClient(),
		RealmName: DefaultRealm().Realm,
	}
}

func DefaultRealm() *Realm {
	return &Realm{
		Roles:               &[]gocloak.Role{*ChatRealmRole()},
		RealmRepresentation: config.KcRealm(),
	}
}

func ChatRealmRole() *gocloak.Role {
	name := "chat-role"
	description := "${role_chat-role}"
	return &gocloak.Role{
		Name:        &name,
		Description: &description,
	}

}

func ChatUser() *User {
	username := "tester"
	password := "password"
	email := "test@email.com"
	firstName := "automated"
	lastName := "tester"
	enabled := true
	emailVerified := true
	realmRoles := []string{
		*ChatRealmRole().Name,
	}
	return &User{
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
		RealmName: DefaultRealm().Realm,
	}
}

func Users() *[]*User {
	return &[]*User{
		ChatUser(),
	}
}
