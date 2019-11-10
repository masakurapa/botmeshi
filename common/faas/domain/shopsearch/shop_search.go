package shopsearch

import (
	"encoding/json"

	"github.com/masakurapa/botmeshi/common/faas/domain/vo"
)

// Params は店検索の外部関数実行パラメーターです
type Params struct {
	// CallBack は外部関数が結果を返す先の情報
	callback vo.NotificationCallbackTarget
	// Query は店舗検索文字列です
	query vo.SearchQuery
}

// New は店検索の外部関数実行パラメーターを初期化します
func New(c vo.NotificationCallbackTarget, q vo.SearchQuery) Params {
	return Params{
		callback: c,
		query:    q,
	}
}

// Unmarshal はJSON文字列をShopSearcの構造体にして返します
func Unmarshal(s string) (*Params, error) {
	var p body
	err := json.Unmarshal([]byte(s), &p)
	if err != nil {
		return nil, err
	}

	callback, err := vo.NewNotificationCallbackTarget(p.Callback)
	if err != nil {
		return nil, err
	}
	query, err := vo.NewSearchQuery(p.Query)
	if err != nil {
		return nil, err
	}

	return &Params{
		callback: callback,
		query:    query,
	}, nil
}

// Marshal はShopSearchをJSONにエンコードした文字列を返します
func (p *Params) Marshal() (string, error) {
	b := body{
		Callback: p.callback.String(),
		Query:    p.query.String(),
	}
	s, err := json.Marshal(&b)
	if err != nil {
		return "", err
	}
	return string(s), nil
}

// Callback は結果返却先の情報を返します
func (p *Params) Callback() vo.NotificationCallbackTarget {
	return p.callback
}

// Query は検索クエリを返します
func (p *Params) Query() vo.SearchQuery {
	return p.query
}
