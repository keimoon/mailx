package fake

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/keimoon/mailx"
)

func TestFake(t *testing.T) {
	b := &bytes.Buffer{}
	provider := &Fake{b}
	mailx.SetProvider(provider)
	message := mailx.NewMessage().SetFrom("Example", "example.com")
	mailx.Send(message)
	fmt.Println(b.String())
}
