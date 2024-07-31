package main

import (
	"context"
	"github.com/go-playground/validator/v10"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"messages/api"
	"messages/domain/services"
	"messages/infrastructure/repositories/message"
	"messages/pkg/configuration"
	"messages/pkg/database"
	"messages/pkg/http"
	"messages/pkg/migration"
)

type config struct {
	Http struct {
		Server http.Config
	}
	Db        database.Config
	Migration migration.Config
}

func newConfig(configuration *configuration.Configuration) (*config, error) {
	c := new(config)
	err := configuration.Unmarshal(c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func main() {
	opts := []fx.Option{
		fx.Provide(zap.NewProduction),
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
		fx.Provide(configuration.New),
		fx.Provide(http.NewServer),
		fx.Provide(newConfig),
		fx.Provide(func(c *config) http.Config {
			return c.Http.Server
		}),
		fx.Provide(func(c *config) database.Config {
			return c.Db
		}),
		fx.Provide(func(c *config) migration.Config { return c.Migration }),
		fx.Provide(fx.Annotate(func(c *config) string { return c.Db.PostgresDSN() }, fx.ResultTags(`name:"dsn"`))),
		fx.Provide(validator.New),
		fx.Provide(database.New),
		fx.Provide(fx.Annotate(message.NewRepository,
			fx.As(new(services.CreateMessageRepository)),
			fx.As(new(services.DeleteMessageByIDRepository)),
			fx.As(new(services.FetchMessageByIDRepository)),
			fx.As(new(services.FetchMessagesByUserIDRepository)),
			fx.As(new(services.UpdateMessageByIDRepository)),
		)),
		fx.Provide(
			services.NewCreateMessage,
			services.NewDeleteMessageByID,
			services.NewFetchMessageByID,
			services.NewFetchMessagesByUserID,
			services.NewUpdateMessageByID,
		),

		fx.Invoke(func(lc fx.Lifecycle, server *http.Server) {
			lc.Append(fx.Hook{
				OnStart: server.OnStart,
				OnStop:  server.OnStop,
			})
		}),
		fx.Invoke(fx.Annotate(migration.Do, fx.ParamTags("", "", `name:"dsn"`))),
		fx.Invoke(api.Registry),
	}

	app := fx.New(opts...)
	err := app.Start(context.Background())
	if err != nil {
		panic(err)
	}

	<-app.Done()

	err = app.Stop(context.Background())
	if err != nil {
		panic(err)
	}
}
