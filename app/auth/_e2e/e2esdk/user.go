package e2esdk

import "github.com/Nerzal/gocloak/v13"

type User struct {
	gocloak.User
	Password *string `json:"-"`
}
