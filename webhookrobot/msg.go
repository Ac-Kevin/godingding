package webhookrobot

import "encoding/json"

type Msg struct {
	Msgtype    string        `json:"msgtype"`
	At         MsgAt         `json:"at"`
	Text       MsgText       `json:"text"`
	Markdown   MsgMarkdown   `json:"markdown"`
	ActionCard MsgActionCard `json:"actionCard"`
	FeedCard   MsgFeedCard   `json:"feedCard"`
	Link       MsgLink       `json:"link"`
}

// ToJSONString def .
func (m *Msg) ToJSONString() string {
	bytes, _ := json.Marshal(m)
	return string(bytes)
}

func (m *Msg) SetText(text MsgText, at MsgAt) {
	m.Msgtype = "text"
	m.At = at
	m.Text = text
}

type MsgAt struct {
	AtMobiles []string `json:"atMobiles"`
	AtUserIds []string `json:"atUserIds"`
	IsAtAll   bool     `json:"isAtAll"`
}
type MsgText struct {
	Content string `json:"content"`
}

type MsgMarkdown struct {
	Text  string `json:"text"`
	Title string `json:"title"`
}

type MsgLink struct {
	MessageURL string `json:"messageUrl"`
	PicURL     string `json:"picUrl"`
	Text       string `json:"text"`
	Title      string `json:"title"`
}

type MsgActionCard struct {
	BtnOrientation string `json:"btnOrientation"`
	SingleTitle    string `json:"singleTitle"`
	SingleURL      string `json:"singleURL"`
	Text           string `json:"text"`
	Title          string `json:"title"`
}

type MsgFeedCard struct {
	Links []struct {
		MessageURL string `json:"messageURL"`
		PicURL     string `json:"picURL"`
		Title      string `json:"title"`
	} `json:"links"`
}
