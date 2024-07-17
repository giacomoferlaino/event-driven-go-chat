package keycloacktest

import "github.com/Nerzal/gocloak/v13"

func chatUser() gocloak.User {
	username := "tester"
	enabled := true
	emailVerified := true
	return gocloak.User{
		Username:      &username,
		Enabled:       &enabled,
		EmailVerified: &emailVerified,
	}
}

func users() []gocloak.User {
	return []gocloak.User{
		chatUser(),
	}
}
