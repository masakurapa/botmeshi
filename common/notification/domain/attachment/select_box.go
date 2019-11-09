package attachment

import "github.com/masakurapa/botmeshi/common/notification/domain/vo"

// SelectBox はセレクトボックスの構造体です
type SelectBox struct {
	options SelectBoxOptions
	// セレクトボックス処理内容
	actionName ActionName
	// 対話通知上の表示形式
	actionType ActionType
}

// SelectBoxOption はセレクトボックスの１要素の構造体です
type SelectBoxOption struct {
	text  string
	value string
}

// SelectBoxOptions はセレクトボックスの要素のリストです
type SelectBoxOptions []SelectBoxOption

// NewSelectBox はセレクトボックスの構造体を初期化して返します
func NewSelectBox(values vo.TextValues) SelectBox {
	opts := make(SelectBoxOptions, len(values.All()))
	for i, s := range values.All() {
		opts[i].text = s.Text()
		opts[i].value = s.Value()
	}

	return SelectBox{
		options:    opts,
		actionName: NewActionNameSelect(),
		actionType: NewActionTypeSelectBox(),
	}
}

// Options はセレクトボックスの全オプションを返します
func (d SelectBox) Options() SelectBoxOptions {
	return d.options
}

// ActionName はセレクトボックスの処理内容の構造体を返します
func (d SelectBox) ActionName() ActionName {
	return d.actionName
}

// ActionType はセレクトボックスの対話通知上の表示形式の構造体を返します
func (d SelectBox) ActionType() ActionType {
	return d.actionType
}

// Text はセレクトボックスの１要素の表示名称を返します
func (d SelectBoxOption) Text() string {
	return d.text
}

// Value はセレクトボックスの１要素の値を返します
func (d SelectBoxOption) Value() string {
	return d.value
}
