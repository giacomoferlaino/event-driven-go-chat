package e2esdk

import "github.com/Nerzal/gocloak/v13"

func NewJWTs() *JWTs {
	return &JWTs{
		byUsername: make(map[string]gocloak.JWT),
	}
}

type JWTs struct {
	byUsername map[string]gocloak.JWT
}

func (j *JWTs) Add(username string, jwt gocloak.JWT) {
	j.byUsername[username] = jwt
}

func (j *JWTs) ByUsername(username string) (gocloak.JWT, bool) {
	jwt, ok := j.byUsername[username]
	return jwt, ok
}
