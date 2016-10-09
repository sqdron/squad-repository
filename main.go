package main

import (
	"github.com/sqdron/squad"
	"github.com/sqdron/squad-repository/api"
	"github.com/sqdron/squad/configurator"
	"github.com/sqdron/squad-repository/service"
)

type Options struct {
	EndpointUrl   string `json:"app_endpoint" option:"Communication Endpoint URL"`
	ApplicationId string `json:"app_id" option:"Application Identity"`
	GithubClient  string `json:"github_client" option:"Github Client ID"`
	GithubSecret  string `json:"github_secret" option:"Github Client secret"`
	GithubRedirect string `json:"github_redirect" option:"Github Redirect Url"`
}

func main() {
	opts := &Options{}
	cfg := configurator.New()
	cfg.ReadFlags(opts)

	providers := service.NewProviders()
	providers.AddProvider("github", opts.GithubClient, opts.GithubSecret, opts.GithubRedirect)

	var squad = squad.Client(opts.EndpointUrl, opts.ApplicationId)

	authApi := api.AuthAPI(providers)
	squad.Api("AuthUrl").Action(authApi.GetAuthUrl)
	squad.Api("token").Action(authApi.GetToken)

	squad.Activate()
}
