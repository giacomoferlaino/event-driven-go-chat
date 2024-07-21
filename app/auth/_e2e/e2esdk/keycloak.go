package e2esdk

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

type KeycloakData struct {
	Realm  gocloak.RealmRepresentation
	Client gocloak.Client
	Users  []User
}

func NewKeycloak(baseUrl string, seedData KeycloakData) (*Keycloak, error) {
	client := gocloak.NewClient(baseUrl)
	ctx := context.Background()
	adminJWT, err := client.LoginAdmin(ctx, adminUser, adminPass, defaultReamID)
	if err != nil {
		return nil, err
	}

	return &Keycloak{
		Client:    client,
		AdminJWT:  adminJWT,
		Ctx:       context.Background(),
		SeedData:  seedData,
		UsersJWTs: NewJWTs(),
	}, nil
}

type Keycloak struct {
	AdminJWT  *gocloak.JWT
	Client    *gocloak.GoCloak
	Ctx       context.Context
	SeedData  KeycloakData
	UsersJWTs *JWTs
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

	err = k.generateJWTs()
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
	_, err := k.Client.CreateRealm(k.Ctx, k.AdminJWT.AccessToken, k.SeedData.Realm)
	if err != nil {
		return err
	}
	return nil
}

func (k *Keycloak) createClients() error {
	_, err := k.Client.CreateClient(k.Ctx, k.AdminJWT.AccessToken, *k.SeedData.Realm.Realm, k.SeedData.Client)
	if err != nil {
		return err
	}

	return nil
}

func (k *Keycloak) createUsers() error {
	for _, user := range k.SeedData.Users {
		userID, err := k.Client.CreateUser(k.Ctx, k.AdminJWT.AccessToken, *k.SeedData.Realm.Realm, user.User)
		if err != nil {
			return err
		}
		err = k.Client.SetPassword(k.Ctx, k.AdminJWT.AccessToken, userID, *k.SeedData.Realm.Realm, *user.Password, false)
		if err != nil {
			return err
		}
	}
	return nil
}

func (k *Keycloak) generateJWTs() error {
	for _, user := range k.SeedData.Users {
		jwt, err := k.Client.Login(
			k.Ctx,
			*k.SeedData.Client.ClientID,
			*k.SeedData.Client.Secret,
			*k.SeedData.Realm.Realm,
			*user.Username,
			*user.Password,
		)
		if err != nil {
			return err
		}
		k.UsersJWTs.Add(*user.Username, *jwt)
	}
	return nil
}

func (k *Keycloak) deleteRealms() error {
	err := k.Client.DeleteRealm(k.Ctx, k.AdminJWT.AccessToken, *k.SeedData.Realm.Realm)
	if err != nil {
		return err
	}

	return nil
}
