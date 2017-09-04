package initdb

import (
	"errors"
	"fmt"
	"os"

	"github.com/dirtyhairy/moneypenny/service/persistence"
	"github.com/spf13/cobra"
)

type Options struct {
	Force bool
}

func Run(cmd *cobra.Command, args []string, options Options) (err error) {
	dbFile := args[0]

	var finfo os.FileInfo
	finfo, err = os.Stat(dbFile)

	switch {
	case err != nil && !os.IsNotExist(err):
		return err

	case err == nil && finfo.IsDir():
		return errors.New(fmt.Sprintf(
			"%s exists and is a directory",
			dbFile,
		))

	case err == nil:
		if options.Force {
			os.Remove(dbFile)
		} else {
			return errors.New(fmt.Sprintf(
				"database %s already exists; please specify --force in order to overwrite it",
				dbFile,
			))
		}
	}

	var p persistence.Provider
	if p, err = persistence.FromSqlite(dbFile); err != nil {
		return
	}

	if err = p.Close(); err != nil {
		return
	}

	fmt.Printf("database successfully initialized\n")

	return
}
