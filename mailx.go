package mailx

import "io"

type Provider interface {
	Send(message *Message) error
}

func Send(message *Message) error {
	return defaultProvider.Send(message)
}

func SendSimpleMessage(fromName, fromEmail, to, subject, text string) error {
	return Send(NewSimpleMessage(fromName, fromEmail, to, subject, text))
}

type nullProvider struct {
}

func (p *nullProvider) Send(message *Message) error {
	return nil
}

var defaultProvider Provider = &nullProvider{}

func SetProvider(provider Provider) {
	defaultProvider = provider
}

type Message struct {
	from        *NameAndEmail
	to          []string
	cc          []string
	bcc         []string
	subject     string
	text        string
	html        string
	attachments []*Attachment
	inlines     []*Attachment
	headers     map[string][]string
}

func NewMessage() *Message {
	return &Message{
		headers: make(map[string][]string),
	}
}

func NewSimpleMessage(fromName, fromEmail, to, subject, text string) *Message {
	return NewMessage().SetFrom(fromName, fromEmail).AddTo(to).SetSubject(subject).SetText(text)
}

func (m *Message) SetFrom(name, email string) *Message {
	m.from = NewNameAndEmail(name, email)
	return m
}

func (m *Message) From() *NameAndEmail {
	return m.from
}

func (m *Message) AddTo(to string) *Message {
	m.to = append(m.to, to)
	return m
}

func (m *Message) To() []string {
	return m.to
}

func (m *Message) AddCC(cc string) *Message {
	m.cc = append(m.cc, cc)
	return m
}

func (m *Message) CC() []string {
	return m.cc
}

func (m *Message) AddBCC(bcc string) *Message {
	m.bcc = append(m.bcc, bcc)
	return m
}

func (m *Message) BCC() []string {
	return m.bcc
}

func (m *Message) SetSubject(subject string) *Message {
	m.subject = subject
	return m
}

func (m *Message) Subject() string {
	return m.subject
}

func (m *Message) SetText(text string) *Message {
	m.text = text
	return m
}

func (m *Message) Text() string {
	return m.text
}

func (m *Message) SetHTML(html string) *Message {
	m.html = html
	return m
}

func (m *Message) HTML() string {
	return m.html
}

func (m *Message) AddAtachment(name, mime string, body io.Reader) *Message {
	m.attachments = append(m.attachments, NewAttachment(name, mime, body))
	return m
}

func (m *Message) Attachments() []*Attachment {
	return m.attachments
}

func (m *Message) AddInline(name, mime string, body io.Reader) *Message {
	m.inlines = append(m.inlines, NewAttachment(name, mime, body))
	return m
}

func (m *Message) Inlines() []*Attachment {
	return m.inlines
}

func (m *Message) AddHeader(key, value string) *Message {
	m.headers[key] = append(m.headers[key], value)
	return m
}

func (m *Message) Headers() map[string][]string {
	return m.headers
}

type NameAndEmail struct {
	name  string
	email string
}

func NewNameAndEmail(name, email string) *NameAndEmail {
	return &NameAndEmail{name, email}
}

func (n *NameAndEmail) Name() string {
	return n.name
}

func (n *NameAndEmail) Email() string {
	return n.email
}

func (n *NameAndEmail) String() string {
	if n.name == "" {
		return n.email
	}
	return n.name + " <" + n.email + ">"
}

type Attachment struct {
	name string
	mime string
	body io.Reader
}

func NewAttachment(name, mime string, body io.Reader) *Attachment {
	return &Attachment{name, mime, body}
}

func (a *Attachment) Name() string {
	return a.name
}

func (a *Attachment) Mime() string {
	return a.mime
}

func (a *Attachment) Body() io.Reader {
	return a.body
}
