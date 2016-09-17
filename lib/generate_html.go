package lib

import (
	"bytes"
	"html/template"
)

type templateData struct {
	BuildErrors string
	Summary     *TestSummary
}

func GenerateHTML(templateStr string, summary *TestSummary, buildErrors string) (string, error) {
	t, err := template.New("webpage").Parse(templateStr)
	if err != nil {
		return "", err
	}

	buf := bytes.Buffer{}
	err = t.Execute(&buf, &templateData{Summary: summary, BuildErrors: buildErrors})
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
