package config

import (
	"os"
	"time"

	"github.com/Nerzal/gocloak/v13"
)

func KcUrl() string {
	return os.Getenv("KC_BASE_URL")
}

func KcRealm() gocloak.RealmRepresentation {
	id := os.Getenv("KC_CHAT_REALM_ID")
	enabled := true
	accessTokenLifespan := int(time.Minute / time.Second * 20)
	return gocloak.RealmRepresentation{
		Realm:               &id,
		Enabled:             &enabled,
		AccessTokenLifespan: &accessTokenLifespan,
	}
}

func KcClient() gocloak.Client {
	clientId := os.Getenv("KC_CHAT_CLIENT_ID")
	clientSecret := os.Getenv("KC_CHAT_CLIENT_SECRET")
	directAccessGrantsEnabled := true
	return gocloak.Client{
		ClientID:                  &clientId,
		Secret:                    &clientSecret,
		DirectAccessGrantsEnabled: &directAccessGrantsEnabled,
	}
}
