package msgraph

import (
	"encoding/json"
	"fmt"
	"net/url"

	b64 "encoding/base64"
)

type Mail struct {
	Message Message `json:"message"`
}

type Message struct {
	ToRecipients  []Recipient `json:"toRecipients"`
	CCRecipients  []Recipient `json:"ccRecipients"`
	BCCRecipients []Recipient `json:"bccRecipients"`
	Subject       string      `json:"subject"`
	Body          Body        `json:"body"`
}

type Recipient struct {
	EmailAddress EmailAddress `json:"emailAddress"`
}

type Body struct {
	ContentType string `json:"contentType"`
	Content     string `json:"content"`
}

// SendMailJSON performs a JSON-based sendMail request
func (g *GraphClient) SendMailJSON(from string, SendMail Mail) error {
	resource := fmt.Sprintf("/users/%v/sendMail", url.PathEscape(from))

	if SendMail.Message.BCCRecipients == nil {
		SendMail.Message.BCCRecipients = []Recipient{}
	}
	if SendMail.Message.CCRecipients == nil {
		SendMail.Message.CCRecipients = []Recipient{}
	}

	data, err := json.Marshal(SendMail)
	if err != nil {
		fmt.Println(err)
	}

	dataStr := string(data)

	err = g.makePOSTAPICall(resource, "application/json", dataStr, nil)
	return err
}

// SendMailMIME performs a text/plain (in MIME format) sendMail request
func (g *GraphClient) SendMailMIME(from string, data []byte) error {
	resource := fmt.Sprintf("/users/%v/sendMail", url.PathEscape(from))
	datab64 := b64.StdEncoding.EncodeToString(data)
	err := g.makePOSTAPICall(resource, "text/plain", datab64, nil)
	return err
}
