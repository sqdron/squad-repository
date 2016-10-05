package api

import (
	"github.com/sqdron/squad/util"
	"net/url"
	"fmt"
	"time"
	"net/http"
	"strings"
	"io/ioutil"
)

type github struct {
	clientId       string
	clientSecret   string
	clientRedirect string
}

type RequestAuthToken struct {
	Code  string
	State string
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
	data.Set("redirect_uri", p.clientRedirect)
	data.Set("state", util.GenerateString(10))
	u, _ := url.ParseRequestURI("https://github.com/login/oauth/authorize")
	u.RawQuery = data.Encode()
	return u.String()
}

func (p *github) GetToken(r RequestAuthToken) string {
	var client = &http.Client{
		Timeout: time.Second * 10,
	}

	form := url.Values{}
	form.Add("client_id", p.clientId)
	form.Add("client_secret", p.clientSecret)
	form.Add("redirect_uri", p.clientRedirect)
	form.Add("code", r.Code)
	form.Add("state", r.State)
	fmt.Println(r.Code)
	fmt.Println(r.State)

	//TODO: implement JSON Accept
	req, err := http.NewRequest("POST", "https://github.com/login/oauth/access_token", strings.NewReader(form.Encode()))
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	if (resp.StatusCode == http.StatusOK){
		fmt.Println("status ok!")
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
		u,e := url.Parse(p.clientRedirect + "?" + string(body))
		fmt.Println(u)
		if (e != nil){
			panic(e)
		}
		return u.Query().Get("access_token")
	}
	return ""
}

type Route struct {
	path   string
	action interface{}
}

func GithubAPI(clientId string, clientSecret string, clientRedirect string) IOAuthApi {
	return &github{clientId, clientSecret, clientRedirect}
}
