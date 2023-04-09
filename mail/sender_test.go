package mail

import (
	"testing"

	"github.com/IkehAkinyemi/mono-finance/utils"
	"github.com/stretchr/testify/require"
)

func TestSendEmailWithGmail(t *testing.T) {
	// if testing.Short() {
	// 	t.Skip()
	// }

	config, err := utils.LoadConfig("..")
	require.NoError(t, err)

	sender := NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)

	subject := "A test email"
	content := `
	<h1>Hello gengs💩!</h1>
	<p>This is a test message from <a href="https://github.com/IkehAkinyemi/mono-finance">Mono Finance</a></p>
	`
	to := []string{"alerudivine@gmail.com", "mrikehchukwuka@gmail.com", "umavictor11@gmail.com", "chrisebuberoland@gmail.com "}
	attachFiles := []string{"../README.md"}

	err = sender.SendEmail(subject, content, to, nil, nil, attachFiles)
	require.NoError(t, err)
}