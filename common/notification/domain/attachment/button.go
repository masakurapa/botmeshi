package attachment

// Button はボタンの構造体です
type Button struct {
	// ボタンテキスト
	text string
	// ボタンのスタイル
	style string
	// ボタンの処理内容
	actionName ActionName
	// 対話通知上の表示形式
	actionType ActionType
}

// NewCancelButton はキャンセル用のボタンを返却します
func NewCancelButton() Button {
	return Button{
		text:       "キャンセル",
		style:      "danger",
		actionName: NewActionNameCancel(),
		actionType: NewActionTypeButton(),
	}
}

// Text はボタンテキストを返します
func (d Button) Text() string {
	return d.text
}

// Style はボタンのスタイルを返します
func (d Button) Style() string {
	return d.style
}

// ActionName はボタンの処理内容の構造体を返します
func (d Button) ActionName() ActionName {
	return d.actionName
}

// ActionType はボタンの対話通知上の表示形式の構造体を返します
func (d Button) ActionType() ActionType {
	return d.actionType
}
