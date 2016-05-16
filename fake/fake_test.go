package fake

import (
	"bytes"
	"testing"

	"github.com/keimoon/mailx"
)

func TestFake(t *testing.T) {
	b := &bytes.Buffer{}
	provider := &Fake{b}
	mailx.SetProvider(provider)
	message := mailx.NewMessage().
		SetFrom("Example", "me@example.com").
		AddTo("to1@example.com").AddTo("to2@example.com").
		AddCC("cc1@example.com").AddCC("cc2@example.com").
		AddBCC("bcc1@example.com").AddBCC("bcc2@example.com").
		AddHeader("RECP_TO", "me@example.com").AddHeader("TrackID", "1").AddHeader("TrackID", "2").
		SetSubject("Test subject").
		SetText("Text body").
		SetHTML("<b>HTML body</b>").
		AddAtachment("attachment1", "text/json", bytes.NewReader([]byte("Attachment 1 body"))).
		AddAtachment("attachment2", "text/html", bytes.NewReader([]byte("Attachment 2 body"))).
		AddInline("inline1", "text/json", bytes.NewReader([]byte("Inline 1 body"))).
		AddInline("inline2", "text/html", bytes.NewReader([]byte("Inline 2 body")))
	mailx.Send(message)
	expectedOutput := `FROM: Example <me@example.com>
TO: to1@example.com, to2@example.com
CC: cc1@example.com, cc2@example.com
BCC: bcc1@example.com, bcc2@example.com

HEADERS:
- RECP_TO: me@example.com
- TrackID: 1
- TrackID: 2

SUBJECT: Test subject

TEXT:
Text body

HTML:
<b>HTML body</b>

ATTACHMENT:
- Name: attachment1
  - Mime: text/json
  - Size: 17
  - Sample: [65 116 116 97 99 104 109 101 110 116 32 49 32 98 111 100 121]
- Name: attachment2
  - Mime: text/html
  - Size: 17
  - Sample: [65 116 116 97 99 104 109 101 110 116 32 50 32 98 111 100 121]

INLINES:
- Name: inline1
  - Mime: text/json
  - Size: 13
  - Sample: [73 110 108 105 110 101 32 49 32 98 111 100 121]
- Name: inline2
  - Mime: text/html
  - Size: 13
  - Sample: [73 110 108 105 110 101 32 50 32 98 111 100 121]`
	if b.String() != expectedOutput {
		t.Fatal("not as expected")
	}
}
