package attachment

const (
	// 対話メッセージでの表示形式 - ボタン
	actionTypeButton = "button"
	// 対話メッセージでの表示形式 - セレクトボックス
	actionTypeSelect = "select"
)

// ActionType は対話式メッセージに表示する形式に関する情報です
type ActionType struct {
	name string
}

// NewActionTypeButton は表示形式がボタンのための名称を返します
func NewActionTypeButton() ActionType {
	return ActionType{name: actionTypeButton}
}

// NewActionTypeSelectBox は表示形式がセレクトボックスのための名称を返します
func NewActionTypeSelectBox() ActionType {
	return ActionType{name: actionTypeSelect}
}

// Name は表示形式の名称を文字列で返します
func (d ActionType) Name() string {
	return d.name
}
