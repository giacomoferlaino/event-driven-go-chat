package keycloacktest

import "github.com/Nerzal/gocloak/v13"

func e2eRealm() gocloak.RealmRepresentation {
	id := "e2e"
	return gocloak.RealmRepresentation{
		Realm: &id,
	}
}

func realms() []gocloak.RealmRepresentation {
	return []gocloak.RealmRepresentation{
		e2eRealm(),
	}
}
