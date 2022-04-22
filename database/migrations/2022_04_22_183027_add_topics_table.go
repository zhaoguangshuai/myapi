package migrations

import (
	"database/sql"
	"gohub/app/models"
	"gohub/app/models/category"
	"gohub/app/models/user"
	"gohub/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type Topic struct {
		models.BaseModel

		Title      string `gorm:"type:varchar(255);not null;index"`
		Body       string `gorm:"type:longtext;not null"`
		UserID     string `gorm:"type:bigint;not null;index"`
		CategoryID string `gorm:"type:bigint;not null;index"`

		// 会创建 user_id 和 category_id 外键的约束
		User     user.User
		Category category.Category

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Topic{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Topic{})
	}

	migrate.Add("2022_04_22_183027_add_topics_table", up, down)
}
