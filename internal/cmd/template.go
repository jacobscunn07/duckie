package cmd

import (
	"github.com/jacobscunn07/duckie/internal/template"
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
		CheckErr(err)

		dataFile, err := cmd.Flags().GetString("data")
		CheckErr(err)

		out, err := cmd.Flags().GetString("out")
		CheckErr(err)

		_, err = template.GenerateTemplates(template.GenerateTemplatesInput{Files: files, DataInputPath: dataFile, OutputPath: out})
		CheckErr(err)
	},
}

func init() {
	rootCmd.AddCommand(templateCmd)

	templateCmd.Flags().StringArrayP("file", "f", []string{}, "Template files to be used for rendering.")
	templateCmd.Flags().String("data", "", "Input file containing data to be used for rendering templates.")
	templateCmd.Flags().StringP("out", "o", "", "Output directory to create rendered files.")
}
