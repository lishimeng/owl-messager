package repo

import (
	"time"

	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/owl-messager/internal/db/model"
)

func CreateMailAttachment(row *model.MailAttachment) error {
	return create(row)
}

func GetMailAttachment(org int, attachmentId string) (row model.MailAttachment, err error) {
	err = orm().Model(&model.MailAttachment{}).
		Equal("org", org).
		Equal("attachment_id", attachmentId).
		First(&row)
	return
}

func UpdateMailAttachmentBound(org int, attachmentId string, bound bool, expiresAt time.Time) error {
	return app.Transaction(func(tx persistence.TxContext) error {
		var row model.MailAttachment
		err := tx.Model(&model.MailAttachment{}).
			Equal("org", org).
			Equal("attachment_id", attachmentId).
			First(&row)
		if err != nil {
			return err
		}
		row.Bound = bound
		row.ExpiresAt = expiresAt
		return updateSelect(tx, &row, "Bound", "ExpiresAt")
	})
}

func DeleteMailAttachment(org int, attachmentId string) error {
	return app.Transaction(func(tx persistence.TxContext) error {
		var row model.MailAttachment
		err := tx.Model(&model.MailAttachment{}).
			Equal("org", org).
			Equal("attachment_id", attachmentId).
			First(&row)
		if err != nil {
			return err
		}
		return tx.Delete(&row)
	})
}

func ListExpiredMailAttachments(before time.Time) (rows []model.MailAttachment, err error) {
	err = orm().Model(&model.MailAttachment{}).
		Where("expires_at < ?", before).
		Find(&rows)
	return
}
