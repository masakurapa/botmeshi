package notification

const (
	ActionNameSelect  = "actionSelect"
	ActionNameCancel  = "actionCancel"
	ActionNameGo      = "actionGo"
	ActionNameDoNotGo = "actionDoNotGo"

	// ActionTypeButton ボタン
	ActionTypeButton = "button"
	// ActionTypeSelect セレクトボックス
	ActionTypeSelect = "select"

	// ButtonStyleDanger ボタンのスタイル
	ButtonStyleDanger = "danger"
)

// Option struct
type Option struct {
	// 送信先
	Target string
	// 送信メッセージ
	Message string
	// 送信メッセージID
	MessageID string
	// 色
	Color string
	// リッチメッセージの送信オプションリスト
	RichMessageOptions []RichMessageOption
}

// RichMessageOption struct
type RichMessageOption struct {
	ActionName    string
	ActionType    string
	Text          string
	Style         string
	SelectOptions []SelectOption
}

type SelectOption struct {
	Text  string
	Value string
}
