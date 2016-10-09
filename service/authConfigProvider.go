package service

type authProviders struct {
	providers map[string]authConfig
}

type authConfig struct {
	Client      string
	Secret      string
	RedirectUrl string
}

func NewProviders() IAuthProvider{
	return &authProviders{providers:make(map[string]authConfig)}
}

type IAuthProvider interface {
	AddProvider(provider string, client string, secret string, redirectUrl string)
	Get(provider string) authConfig
}

func (c *authProviders) AddProvider(provider string, client string, secret string, redirectUrl string) {
	c.providers[provider] = &authConfig{client, secret, redirectUrl}
}

func (c *authProviders) Get(provider string) authConfig {
	return c[provider]
}