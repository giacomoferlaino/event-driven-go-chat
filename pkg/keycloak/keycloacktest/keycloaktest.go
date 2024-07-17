package keycloacktest

import (
	"context"
	"log"

	"github.com/Nerzal/gocloak/v13"
)

const (
	adminUser     = "admin"
	adminPass     = "password"
	defaultReamID = "master"
)

func New() (*KeycloakTest, error) {
	client := gocloak.NewClient("http://localhost:8081")
	ctx := context.Background()
	token, err := client.LoginAdmin(ctx, adminUser, adminPass, defaultReamID)
	if err != nil {
		return nil, err
	}

	return &KeycloakTest{
		Client: client,
		Token:  token,
		Ctx:    context.Background(),
	}, nil
}

type KeycloakTest struct {
	Token  *gocloak.JWT
	Client *gocloak.GoCloak
	Ctx    context.Context
}

func (k *KeycloakTest) Setup() error {
	err := k.createRealms()
	if err != nil {
		return err
	}

	err = k.createUsers()
	if err != nil {
		return err
	}

	return nil
}

func (k *KeycloakTest) Teardown() {
	err := k.deleteRealms()
	if err != nil {
		log.Println(err)
	}
}

func (k *KeycloakTest) createRealms() error {
	for _, realm := range realms() {
		_, err := k.Client.CreateRealm(k.Ctx, k.Token.AccessToken, realm)
		if err != nil {
			return err
		}
	}
	return nil
}

func (k *KeycloakTest) deleteRealms() error {
	for _, realm := range realms() {
		err := k.Client.DeleteRealm(k.Ctx, k.Token.AccessToken, *realm.Realm)
		if err != nil {
			return err
		}
	}
	return nil
}

func (k *KeycloakTest) createUsers() error {
	for _, user := range users() {
		_, err := k.Client.CreateUser(k.Ctx, k.Token.AccessToken, *e2eRealm().Realm, user)
		if err != nil {
			return err
		}
	}
	return nil
}
