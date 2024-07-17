package e2esdk

import (
	"chat/app/auth/domain"
	"context"
	"log"

	"github.com/Nerzal/gocloak/v13"
)

const (
	adminUser     = "admin"
	adminPass     = "password"
	defaultReamID = "master"
)

type KeycloakData struct {
	Realm   gocloak.RealmRepresentation
	Clients []gocloak.Client
	Users   []domain.User
}

func NewKeycloak(baseUrl string, seedData KeycloakData) (*Keycloak, error) {
	client := gocloak.NewClient(baseUrl)
	ctx := context.Background()
	token, err := client.LoginAdmin(ctx, adminUser, adminPass, defaultReamID)
	if err != nil {
		return nil, err
	}

	return &Keycloak{
		Client:   client,
		Token:    token,
		Ctx:      context.Background(),
		SeedData: seedData,
	}, nil
}

type Keycloak struct {
	Token    *gocloak.JWT
	Client   *gocloak.GoCloak
	Ctx      context.Context
	SeedData KeycloakData
}

func (k *Keycloak) Setup() error {
	err := k.createRealms()
	if err != nil {
		return err
	}

	err = k.createClients()
	if err != nil {
		return err
	}

	err = k.createUsers()
	if err != nil {
		return err
	}

	return nil
}

func (k *Keycloak) Teardown() {
	err := k.deleteRealms()
	if err != nil {
		log.Println(err)
	}
}

func (k *Keycloak) createRealms() error {
	_, err := k.Client.CreateRealm(k.Ctx, k.Token.AccessToken, k.SeedData.Realm)
	if err != nil {
		return err
	}
	return nil
}

func (k *Keycloak) createClients() error {
	for _, client := range k.SeedData.Clients {
		_, err := k.Client.CreateClient(k.Ctx, k.Token.AccessToken, *k.SeedData.Realm.Realm, client)
		if err != nil {
			return err
		}
	}
	return nil
}

func (k *Keycloak) createUsers() error {
	for _, user := range k.SeedData.Users {
		userID, err := k.Client.CreateUser(k.Ctx, k.Token.AccessToken, *k.SeedData.Realm.Realm, user.KCUser)
		if err != nil {
			return err
		}
		err = k.Client.SetPassword(k.Ctx, k.Token.AccessToken, userID, *k.SeedData.Realm.Realm, *user.Password, false)
		if err != nil {
			return err
		}
	}
	return nil
}

func (k *Keycloak) deleteRealms() error {
	err := k.Client.DeleteRealm(k.Ctx, k.Token.AccessToken, *k.SeedData.Realm.Realm)
	if err != nil {
		return err
	}

	return nil
}
