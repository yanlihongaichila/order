package model

import (
	"github.com/yanlihongaichila/framework/mysql"
)

func Migrate() error {
	err := mysql.Db.AutoMigrate(NewOrder())
	if err != nil {
		return err
	}

	return nil
}
