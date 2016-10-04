package main

import (
	"fmt"
	"github.com/sqdron/squad"
	"github.com/sqdron/squad-repository/api"
	"github.com/sqdron/squad/configurator"
)

type Options struct {
	EndpointUrl   string `json:"app_endpoint" option:"Communication Endpoint URL"`
	ApplicationId string `json:"app_id" option:"Application Identity"`
	GithubClient  string `json:"github_client" option:"Github Client ID"`
	GithubSecret  string `json:"github_secret" option:"Github Client secret"`
}

func main() {
	opts := &Options{}
	cfg := configurator.New()
	cfg.ReadFlags(opts)
	var squad = squad.Client(opts.EndpointUrl, opts.ApplicationId)
	github := api.GithubAPI(opts.GithubClient, opts.GithubSecret)
	fmt.Println(github)
	squad.Api().Action("auth", github.GetAuthUrl)
	//squad.api.request("auth", github.GetAuthUrl)
	squad.Activate()
}
