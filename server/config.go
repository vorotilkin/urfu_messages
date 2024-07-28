package main

import (
	"messages/pkg/configuration"
	"messages/pkg/database"
	"messages/pkg/http"
)

type config struct {
	Http struct {
		Server http.Config
	}
	Db database.Config
}

func newConfig(configuration *configuration.Configuration) (*config, error) {
	c := new(config)
	err := configuration.Unmarshal(c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
