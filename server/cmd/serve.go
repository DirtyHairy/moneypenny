// Copyright © 2017 Christian Speckner <cnspeckn@googlemail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"github.com/dirtyhairy/moneypenny/server/cmd/serve"
	"github.com/spf13/cobra"
)

func init() {
	options := serve.Options{}

	// serveCmd represents the serve command
	serveCmd := &cobra.Command{
		Use:   "serve <database>",
		Short: "Start the server.",
		Long:  `Start the moneypenny server.`,
		Run: func(cmd *cobra.Command, args []string) {
			failIf(serve.Run(cmd, args, options))
		},
	}

	RootCmd.AddCommand(serveCmd)

	serveCmd.Args = cobra.ExactArgs(1)

	flags := serveCmd.PersistentFlags()
	flags.StringVarP(&options.Listen, "listen", "l", "localhost:8888", "listen address")
	flags.BoolVarP(&options.Debug, "debug", "d", false, "debug mode")
	flags.StringVar(&options.Logfile, "logfile", "", "log into file")
	flags.StringVar(&options.StaticPath, "static", "", "service static files from this directory")
}