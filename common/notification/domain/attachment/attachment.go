package attachment

import "strings"

const (
	attachmentIDPrefix string = "notice"
)

// Attachment は対話式メッセージのオプションの構造体です
type Attachment struct {
	color     string
	selectBox SelectBox
	button    Button
}

// NewAttachment は対話式メッセージのオプションを初期化します
func NewAttachment(sb SelectBox, b Button) Attachment {
	return Attachment{
		color:     "#ff6633",
		selectBox: sb,
		button:    b,
	}
}

// Color はメッセージの色を返します
func (d Attachment) Color() string {
	return d.color
}

// SelectBox はセレクトボックスの構造体を返します
func (d Attachment) SelectBox() SelectBox {
	return d.selectBox
}

// Button はボタンの構造体を返します
func (d Attachment) Button() Button {
	return d.button
}

// ID はオプションのIDを返します
func (d Attachment) ID() string {
	s := attachmentIDPrefix
	sb := d.SelectBox()
	s += strings.Title(sb.ActionName().Name())
	s += strings.Title(sb.ActionType().Name())
	return s
}
