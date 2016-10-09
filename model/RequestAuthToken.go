package model

type RequestAuthToken struct {
	RequestAuth
	Code  string
	State string
}