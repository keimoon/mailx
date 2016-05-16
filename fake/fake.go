package fake

import (
	"io"
	"io/ioutil"
	"strings"
	"text/template"

	"github.com/keimoon/mailx"
)

type Fake struct {
	w io.Writer
}

func NewProvider(w io.Writer) mailx.Provider {
	return &Fake{w}
}

var tmpl = template.Must(template.New("mailx/fake").Parse(
	`FROM: {{.From}}
TO: {{.To}}
CC: {{.CC}}
BCC: {{.BCC}}

HEADERS:
{{- range $key, $values := .Headers}}
{{- range $values}}
- {{$key}}: {{.}}
{{- end}}
{{- end}}

SUBJECT: {{.Subject}}

TEXT:
{{.Text}}

HTML:
{{.HTML}}

ATTACHMENT:
{{- range $attachment := .Attachments}}
- Name: {{$attachment.Name}}
  - Mime: {{$attachment.Mime}}
  - Size: {{$attachment.Size}}
  - Sample: {{$attachment.Sample}}
{{- end}}

INLINES:
{{- range $attachment := .Inlines}}
- Name: {{$attachment.Name}}
  - Mime: {{$attachment.Mime}}
  - Size: {{$attachment.Size}}
  - Sample: {{$attachment.Sample}}
{{- end}}`))

type attachmentSample struct {
	Name   string
	Mime   string
	Size   int
	Sample []byte
}

func readSampleAttachment(attachment *mailx.Attachment) (*attachmentSample, error) {
	sample := &attachmentSample{
		Name: attachment.Name(),
		Mime: attachment.Mime(),
	}
	b, err := ioutil.ReadAll(attachment.Body())
	if err != nil {
		return nil, err
	}
	sample.Size = len(b)
	if sample.Size < 25 {
		sample.Sample = b
	} else {
		sample.Sample = b[:25]
	}
	return sample, nil
}

func (p *Fake) Send(message *mailx.Message) error {
	attachmentSamples := []*attachmentSample{}
	for _, attachment := range message.Attachments() {
		sample, err := readSampleAttachment(attachment)
		if err != nil {
			return err
		}
		attachmentSamples = append(attachmentSamples, sample)
	}
	inlineSamples := []*attachmentSample{}
	for _, attachment := range message.Inlines() {
		sample, err := readSampleAttachment(attachment)
		if err != nil {
			return err
		}
		inlineSamples = append(inlineSamples, sample)
	}
	data := &struct {
		From        string
		To          string
		CC          string
		BCC         string
		Headers     map[string][]string
		Subject     string
		Text        string
		HTML        string
		Attachments []*attachmentSample
		Inlines     []*attachmentSample
	}{
		From:        message.From().String(),
		To:          strings.Join(message.To(), ", "),
		CC:          strings.Join(message.CC(), ", "),
		BCC:         strings.Join(message.BCC(), ", "),
		Headers:     message.Headers(),
		Subject:     message.Subject(),
		Text:        message.Text(),
		HTML:        message.HTML(),
		Attachments: attachmentSamples,
		Inlines:     inlineSamples,
	}
	tmpl.Execute(p.w, data)
	return nil
}
