package template

import (
	"encoding/json"
	"fmt"
	"os"
	"text/template"
)

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
