package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"text/template"

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

		_, err = GenerateTemplates(GenerateTemplatesInput{Files: files, DataInputPath: dataFile, OutputPath: out})
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(templateCmd)

	templateCmd.Flags().StringArrayP("file", "f", []string{}, "Template files to be used for rendering.")
	templateCmd.Flags().String("data", "", "Input file containing data to be used for rendering templates.")
	templateCmd.Flags().StringP("out", "o", "", "Output directory to create rendered files.")
}

type GenerateTemplatesInput struct {
	Files         []string
	DataInputPath string
	OutputPath    string
}

type GenerateTemplatesOutput struct {
	OutputPath string
}

func GenerateTemplates(input GenerateTemplatesInput) (*GenerateTemplatesOutput, error) {
	var data interface{}

	dataFileBytes, err := os.ReadFile(input.DataInputPath)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(dataFileBytes, &data); err != nil {
		return nil, err
	}

	for _, file := range input.Files {
		t, err := template.ParseFiles(file)
		if err != nil {
			return nil, err
		}

		outF, err := os.Create(fmt.Sprintf("%s/%s", input.OutputPath, file))
		if err != nil {
			return nil, err
		}

		err = t.Execute(outF, data)
		if err != nil {
			return nil, err
		}

		outF.Close()
	}

	return &GenerateTemplatesOutput{OutputPath: input.OutputPath}, nil
}
