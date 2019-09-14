package repository

import "github.com/masakurapa/botmeshi/app/domain/model/notification"

// Notification interface
type Notification interface {
	PostMessage(option notification.Option) error
	PostRichMessage(option notification.Option) error
}
