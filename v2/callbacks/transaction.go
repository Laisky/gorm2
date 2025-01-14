package callbacks

import (
	"github.com/Laisky/gorm/v2"
)

func BeginTransaction(db *gorm.DB) {
	if !db.Config.SkipDefaultTransaction {
		if tx := db.Begin(); tx.Error == nil {
			db.Statement.ConnPool = tx.Statement.ConnPool
			db.InstanceSet("gorm:started_transaction", true)
		} else if tx.Error == gorm.ErrInvalidTransaction {
			tx.Error = nil
		} else {
			db.Error = tx.Error
		}
	}
}

func CommitOrRollbackTransaction(db *gorm.DB) {
	if !db.Config.SkipDefaultTransaction {
		if _, ok := db.InstanceGet("gorm:started_transaction"); ok {
			if db.Error == nil {
				db.Commit()
			} else {
				db.Rollback()
			}
			db.Statement.ConnPool = db.ConnPool
		}
	}
}
