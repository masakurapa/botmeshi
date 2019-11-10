package vo

import "github.com/masakurapa/botmeshi/common/faas/errs"

// SearchQuery は検索クエリの構造体です
type SearchQuery struct {
	value string
}

// NewSearchQuery は検索クエリを初期化します
func NewSearchQuery(v string) (SearchQuery, error) {
	m := SearchQuery{}
	if v == "" {
		return m, errs.NewRequiredArgsError("message")
	}

	m.value = v
	return m, nil
}

// String はテキストメッセージを返します
func (v SearchQuery) String() string {
	return v.value
}
