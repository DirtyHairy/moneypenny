package serve

import (
	"os"

	"github.com/dirtyhairy/moneypenny/server"
	"github.com/dirtyhairy/moneypenny/service/persistence"
	"github.com/spf13/cobra"
)

type Options struct {
	Listen  string
	Debug   bool
	Logfile string
}

func Run(cmd *cobra.Command, args []string, options Options) (err error) {
	databaseFile := args[0]

	var p persistence.Provider
	if p, err = persistence.FromSqlite(databaseFile); err != nil {
		return
	}

	defer p.Close()

	logWriter := os.Stdout
	if options.Logfile != "" {
		if logWriter, err = os.OpenFile(options.Logfile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666); err != nil {
			return
		}

		defer logWriter.Close()
	}

	config := server.Config{
		Listen:      options.Listen,
		Debug:       options.Debug,
		Persistence: p,
		LogWriter:   logWriter,
	}

	var s server.Server
	if s, err = server.Create(config); err != nil {
		return
	}

	if err = s.Start(); err != nil {
		return
	}

	return
}
