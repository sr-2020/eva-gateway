package service

import (
	"github.com/sr-2020/eva-gateway/app/adapter/client"
)

type Service struct {
	Host       string
	Path       string
	Client     client.Client
}

var Services map[string]Service
