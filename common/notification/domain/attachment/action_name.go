package attachment

const (
	// 対話メッセージでの処理内容 - 選択
	actionNameSelect = "actionSelect"
	// 対話メッセージでの処理内容 - キャンセル
	actionNameCancel = "actionCancel"
)

// ActionName は対話式メッセージでの処理内容を表す型です
type ActionName struct {
	name string
}

// NewActionNameSelect は選択処理の型を返します
func NewActionNameSelect() ActionName {
	return ActionName{name: actionNameSelect}
}

// NewActionNameCancel はキャンセル処理の型を返します
func NewActionNameCancel() ActionName {
	return ActionName{name: actionNameCancel}
}

// Name は処理内容の名称を文字列で返します
func (d ActionName) Name() string {
	return d.name
}
