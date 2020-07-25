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
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/k1LoW/lrep/format"
	"github.com/k1LoW/lrep/format/json"
	"github.com/k1LoW/lrep/format/ltsv"
	"github.com/k1LoW/lrep/format/sqlite"
	"github.com/k1LoW/lrep/parser"
	"github.com/k1LoW/lrep/version"
	"github.com/mattn/go-isatty"
	"github.com/spf13/cobra"
)

var (
	regexp  string
	file    string
	fFormat string
	noM0    bool
	noRaw   bool
)

var rootCmd = &cobra.Command{
	Use:   "lrep [REGEXP]",
	Short: "line regular expression parser",
	Long:  `line regular expression parser.`,
	Args: func(cmd *cobra.Command, args []string) error {
		for n, r := range parser.Builtins {
			bool, err := cmd.Flags().GetBool(n)
			if err != nil {
				return err
			}
			if bool {
				if regexp != "" {
					return errors.New("only one built-in regexp can be selected")
				}
				regexp = r.Regexp
			}
		}
		if len(args) == 1 && regexp != "" {
			return errors.New("select either an argument or a built-in regexp")
		}
		if len(args) != 1 && regexp == "" {
			return fmt.Errorf("accepts %d arg(s), received %d", 1, len(args))
		}

		if (isatty.IsTerminal(os.Stdin.Fd()) && file == "") || (!isatty.IsTerminal(os.Stdin.Fd()) && file != "") {
			return fmt.Errorf("%s need either target file(--file) or STDIN", version.Name)
		}
		return nil
	},
	Version: version.Version,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			in *bufio.Reader
			f  format.Formatter
		)
		if len(args) > 0 {
			regexp = args[0]
		}
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		if file != "" {
			fi, err := os.Open(filepath.Clean(file))
			if err != nil {
				printFatalln(cmd, err)
			}
			defer func() {
				_ = fi.Close()
			}()
			in = bufio.NewReader(fi)
		} else {
			in = bufio.NewReader(os.Stdin)
		}

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

		opts := []parser.Option{}
		if noM0 {
			opts = append(opts, parser.NoM0())
		}
		if noRaw {
			opts = append(opts, parser.NoRaw())
		}

		p := parser.New(regexp, opts...)

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
	rootCmd.Flags().StringVarP(&file, "file", "f", "", "input file")

	rootCmd.Flags().StringVarP(&fFormat, "format", "t", "json", "output format")
	if err := rootCmd.RegisterFlagCompletionFunc("format", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"json", "ltsv", "sqlite"}, cobra.ShellCompDirectiveDefault
	}); err != nil {
		printFatalln(rootCmd, err)
	}

	for n, r := range parser.Builtins {
		rootCmd.Flags().BoolP(n, "", false, fmt.Sprintf("[build-in regexp] %s", r.Desc))
	}

	rootCmd.Flags().BoolVarP(&noM0, "no-m0", "", false, "ignore regexp submatches[0]")
	rootCmd.Flags().BoolVarP(&noRaw, "no-raw", "", false, "ignore line raw data")
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
