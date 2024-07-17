package domain

import "github.com/Nerzal/gocloak/v13"

type User struct {
	KCUser   gocloak.User
	Password *string `json:"-"`
}
