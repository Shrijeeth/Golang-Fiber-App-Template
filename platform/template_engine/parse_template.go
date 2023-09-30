package template_engine

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

func ParseHTMLTemplate(templateName string, templateData interface{}) ([]byte, error) {
	templateNameStr := fmt.Sprint(templateName)
	templatePath := filepath.Join("./templates/html", templateNameStr)

	templateBytes, err := os.ReadFile(templatePath)
	if err != nil {
		return nil, err
	}

	templateParsed, err := template.New(templateNameStr).Parse(string(templateBytes))
	if err != nil {
		return nil, err
	}

	var renderedTemplate bytes.Buffer
	err = templateParsed.Execute(&renderedTemplate, templateData)
	if err != nil {
		return nil, err
	}

	return renderedTemplate.Bytes(), nil
}
