package api

import (
	"github.com/sqdron/squad/util"
	"net/url"
	"fmt"
)

type github struct {
	clientId     string
	clientSecret string
}


type RequestAuthToken struct {
	ClientId string
	ClientSecret string
	Code string
	RedirectUrl string
}

type IOAuthApi interface {
	GetAuthUrl() string
	GetToken(RequestAuthToken) string
}

func (p *github) GetAuthUrl() string {
	fmt.Println("GetAuthUrl...")
	data := url.Values{}
	data.Set("client_id", p.clientId)
	data.Set("scope", "email repo")
	data.Set("state", util.GenerateString(10))
	data.Set("redirect_uri", "bar")
	u, _ := url.ParseRequestURI("https://github.com/login/oauth/authorize")
	u.RawQuery = data.Encode()
	fmt.Println("rrr "+ u.String())
	return u.String()
}

func (p *github) GetToken(r RequestAuthToken) string {
	return fmt.Sprintf("%s %s %s %s", r.ClientId, r.ClientSecret, r.Code, r.RedirectUrl)
}

type Route struct {
	path   string
	action interface{}
}

func GithubAPI(clientId string, clientSecret string) IOAuthApi {
	return &github{clientId, clientSecret}
}
