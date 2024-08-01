package e2esdk

import "github.com/Nerzal/gocloak/v13"

type User struct {
	gocloak.User
	Password  *string
	JWT       *gocloak.JWT
	RealmName *string
}

type Realm struct {
	gocloak.RealmRepresentation
	Roles *[]gocloak.Role
}

type Client struct {
	gocloak.Client
	RealmName *string
}
