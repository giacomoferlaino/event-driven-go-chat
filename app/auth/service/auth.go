package service

type Auth interface {
	Login(username string, password string) (string, error)
}
