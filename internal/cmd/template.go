package cmd

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"

	"github.com/spf13/cobra"
)

var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "Renders a file from a template file and data input file",
	Long: `Renders a file based on the template file and data input file provided.
For example the command below will render a file based on the template file and the data file provided
and put the rendered file in the a directory named rendered in the current directory.

  duckie template --data data.json --file template.tpl -o $PWD/rendered
`,
	Run: func(cmd *cobra.Command, args []string) {
		files, err := cmd.Flags().GetStringArray("file")
		if err != nil {
			panic(err)
		}

		dataFile, err := cmd.Flags().GetString("data")
		if err != nil {
			panic(err)
		}

		out, err := cmd.Flags().GetString("out")
		if err != nil {
			panic(err)
		}

		var data interface{}

		dataFileBytes, err := os.ReadFile(dataFile)
		if err != nil {
			panic(err)
		}

		if err := json.Unmarshal(dataFileBytes, &data); err != nil {
			panic(err)
		}

		for _, file := range files {
			t, err := template.ParseFiles(file)
			if err != nil {
				panic(err)
			}

			outF, err := os.Create(fmt.Sprintf("%s/%s", out, file))
			if err != nil {
				panic(err)
			}

			err = t.Execute(outF, data)
			if err != nil {
				panic(err)
			}

			outF.Close()
		}
	},
}

func init() {
	rootCmd.AddCommand(templateCmd)

	templateCmd.Flags().StringArrayP("file", "f", []string{}, "Template files to be used for rendering.")
	templateCmd.Flags().String("data", "", "Input file containing data to be used for rendering templates.")
	templateCmd.Flags().StringP("out", "o", "", "Output directory to create rendered files.")
}
