package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Middleware func(http.ResponseWriter, *http.Request, httprouter.Params) (*http.Response, error)

type Middlewares struct {
	Billing ServiceMiddleware
	Auth ServiceMiddleware
	Position ServiceMiddleware
	Push ServiceMiddleware
}

type ServiceMiddleware struct {
	Global []Middleware
	Route map[string][]Middleware
}

var MiddlewareMap = Middlewares{
	Auth: ServiceMiddleware{
		Global: nil,
		Route: map[string][]Middleware{
			"/profile": {
				AuthMiddleware,
			},
			"/login": {
				LoginMiddleware,
			},
		},
	},
	Billing: ServiceMiddleware{
		Global: []Middleware{
			AuthMiddleware,
		},
		Route: map[string][]Middleware{
			"/account_info": {
				AccountInfoMiddleware,
			},
			"/transfer": {
				TransferMiddleware,
			},
		},
	},
	Position: ServiceMiddleware{
		Global: []Middleware{
			AuthMiddleware,
		},
		Route: nil,
	},
	Push: ServiceMiddleware{
		Global: []Middleware{
			AuthMiddleware,
		},
		Route: nil,
	},
}
