/*
Copyright © 2020 Ken'ichiro Oyama <k1lowxb@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/k1LoW/regexq/format"
	"github.com/k1LoW/regexq/format/json"
	"github.com/k1LoW/regexq/format/ltsv"
	"github.com/k1LoW/regexq/format/sqlite"
	"github.com/k1LoW/regexq/parser"
	"github.com/k1LoW/regexq/version"
	"github.com/mattn/go-isatty"
	"github.com/spf13/cobra"
)

var fFormat string

var rootCmd = &cobra.Command{
	Use:   "regexq [REGEXP]",
	Short: "regexq",
	Long:  `regexq.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("accepts %d arg(s), received %d", 1, len(args))
		}
		if isatty.IsTerminal(os.Stdin.Fd()) {
			return fmt.Errorf("%s need STDIN. Please use pipe", version.Name)
		}
		return nil
	},
	Version: version.Version,
	Run: func(cmd *cobra.Command, args []string) {
		var f format.Formatter
		regexp := args[0]
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		in := bufio.NewReader(os.Stdin)
		out := os.Stdout

		switch fFormat {
		case "json":
			f = json.New(out)
		case "ltsv":
			f = ltsv.New(out)
		case "sqlite":
			f = sqlite.New(out)
		default:
			printFatalln(cmd, fmt.Errorf("unsupported format '%s'", fFormat))
		}

		p := parser.New(regexp)

		schema := p.Schema()
		if err := f.WriteSchema(schema); err != nil {
			printFatalln(cmd, err)
		}

	L:
		for {
			s, err := in.ReadString('\n')
			if err == io.EOF {
				break L
			} else if err != nil {
				printFatalln(cmd, err)
			}
			select {
			case <-ctx.Done():
				break L
			default:
				parsed := p.Parse(strings.TrimSuffix(s, "\n"))
				if err := f.Write(schema, parsed); err != nil {
					printFatalln(cmd, err)
				}
			}
		}
	},
}

func Execute() {
	rootCmd.SetOut(os.Stdout)
	rootCmd.SetErr(os.Stderr)

	log.SetOutput(ioutil.Discard)
	if env := os.Getenv("DEBUG"); env != "" {
		debug, err := os.Create(fmt.Sprintf("%s.debug", version.Name))
		if err != nil {
			printFatalln(rootCmd, err)
		}
		log.SetOutput(debug)
	}

	if err := rootCmd.Execute(); err != nil {
		printFatalln(rootCmd, err)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&fFormat, "format", "t", "json", "output format")
}

// https://github.com/spf13/cobra/pull/894
func printErrln(c *cobra.Command, i ...interface{}) {
	c.PrintErr(fmt.Sprintln(i...))
}

func printErrf(c *cobra.Command, format string, i ...interface{}) {
	c.PrintErr(fmt.Sprintf(format, i...))
}

func printFatalln(c *cobra.Command, i ...interface{}) {
	printErrln(c, i...)
	os.Exit(1)
}

func printFatalf(c *cobra.Command, format string, i ...interface{}) {
	printErrf(c, format, i...)
	os.Exit(1)
}
