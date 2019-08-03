package main

import (
	"github.com/justinas/alice"
)

type Middleware = alice.Constructor

type Middlewares struct {
	Billing       ServiceMiddleware
	Auth          ServiceMiddleware
	Position      ServiceMiddleware
	Push          ServiceMiddleware
	ModelEngine   ServiceMiddleware
	ModelsManager ServiceMiddleware
}

type ServiceMiddleware struct {
	Global []Middleware
	Route  map[string][]Middleware
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
	ModelEngine: ServiceMiddleware{
		Global: []Middleware{
			AuthMiddleware,
		},
		Route: nil,
	},
	ModelsManager: ServiceMiddleware{
		Global: []Middleware{
			AuthMiddleware,
		},
		Route: map[string][]Middleware{
			"/character/model": {
				CharacterModelMiddleware,
			},
		},
	},
}
