package repository

type IdentityProvider interface {
	GetAccessToken(username string, password string) (string, error)
}
