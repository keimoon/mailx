package fake

import (
	"io"
	"text/template"

	"github.com/keimoon/mailx"
)

type Fake struct {
	w io.Writer
}

var tmpl = template.Must(template.New("mailx/fake").Parse(
	`FROM: {{.From}}{{"" -}}
`))

func (p *Fake) Send(message *mailx.Message) error {
	data := &struct {
		From string
	}{
		From: message.From().String(),
	}
	tmpl.Execute(p.w, data)
	return nil
}
