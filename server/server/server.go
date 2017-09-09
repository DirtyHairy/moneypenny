package server

import (
	"os"

	"github.com/dirtyhairy/moneypenny/server/service/persistence"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

type server struct {
	config      Config
	persistence persistence.Provider
	echo        *echo.Echo
}

func Create(config Config) (_s Server, err error) {
	s := server{
		config:      config,
		persistence: config.Persistence,
		echo:        echo.New(),
	}

	if config.LogWriter == nil {
		config.LogWriter = os.Stderr
	}

	s.echo.Logger.SetOutput(config.LogWriter)
	if config.Debug {
		s.echo.Logger.SetLevel(log.DEBUG)
	}

	s.echo.Use(s.addContextMiddleware())

	s.echo.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: config.LogWriter,
	}))

	s.echo.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		DisablePrintStack: !config.Debug,
	}))

	s.echo.Use(middleware.Gzip())

	s.setupRouting()

	_s = &s
	return
}

func (s *server) Start() error {
	return s.echo.Start(s.config.Listen)
}

func (s *server) addContextMiddleware() echo.MiddlewareFunc {
	return func(handler echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ourContext := context{
				Context:     c,
				persistence: s.persistence,
			}

			return handler(&ourContext)
		}
	}
}
