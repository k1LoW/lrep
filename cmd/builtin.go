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
	"fmt"
	"os"
	"text/template"

	"github.com/k1LoW/lrep/parser"
	"github.com/labstack/gommon/color"
	"github.com/spf13/cobra"
)

// builtinCmd represents the builtin command
var builtinCmd = &cobra.Command{
	Use:   "builtin",
	Short: "show buildin regexp patterns",
	Long:  `show buildin regexp patterns.`,
	Args:  cobra.MaximumNArgs(1),
	ValidArgs: func() []string {
		args := []string{}
		for n := range parser.Builtins {
			args = append(args, n)
		}
		return args
	}(),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			maxlen := 0
			for n := range parser.Builtins {
				if len(n) > maxlen {
					maxlen = len(n)
				}
			}
			for n, r := range parser.Builtins {
				cmd.Printf(fmt.Sprintf("%%-%ds   %%s\n", maxlen), n, r.Desc)
			}
			return
		}
		n := args[0]
		r := parser.Builtins[n]

		tmpl := template.Must(template.New("builtin").Funcs(funcs()).Parse(`{{ "NAME" | bold }}
       {{ .n }} -- {{ .r.Desc }}

{{ "REGEXP" | bold }}
       {{ .r.Regexp | bold }}

{{ "SAMPLE" | bold }}
       {{index .r.Samples 0}}
`))
		params := map[string]interface{}{
			"n": n,
			"r": r,
		}
		if err := tmpl.Execute(os.Stdout, params); err != nil {
			printFatalln(cmd, err)
		}
	},
}

func funcs() map[string]interface{} {
	return map[string]interface{}{
		"bold": func(text string) string {
			return color.Bold(text)
		},
	}
}

func init() {
	rootCmd.AddCommand(builtinCmd)
}
