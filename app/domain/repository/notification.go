package repository

import "github.com/masakurapa/botmeshi/app/domain/model/notification"

// Notification interface
type Notification interface {
	PostMessage(notification.Option) error
	PostRichMessage(notification.Option) error
}
