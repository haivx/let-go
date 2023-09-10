package mailer

import (
	"bytes"
	"html/template"
)

func renderHTML(data map[string]interface{}, layout string) (string, error) {

	tmpl := template.Must(template.ParseFiles(layout))

	buf := new(bytes.Buffer)

	if err := tmpl.Execute(buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
