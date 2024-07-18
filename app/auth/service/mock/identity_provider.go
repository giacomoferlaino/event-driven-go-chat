package mock

type GetAccessTokenReturn struct {
	Val string
	Err error
}

func NewIdentityProvider() IdentityProvider {
	return IdentityProvider{
		GetAccessTokenReturn: &GetAccessTokenReturn{},
	}
}

type IdentityProvider struct {
	GetAccessTokenReturn *GetAccessTokenReturn
}

func (i IdentityProvider) GetAccessToken(username string, password string) (string, error) {
	return i.GetAccessTokenReturn.Val, i.GetAccessTokenReturn.Err
}
