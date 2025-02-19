package seeders

import (
	"github.com/shabashab/chattin/apps/chat-server/src/database/models"
	"gorm.io/gorm"
)

type AppsSeeder struct {}

func NewAppsSeeder() (Seeder) {
	return &AppsSeeder{}
}

func (*AppsSeeder) Name() string {
	return "01_apps_seeder"
}

func (*AppsSeeder) Execute(db *gorm.DB) (_ error) {
	apps := []*models.App{
		{Name: "app-01"},
		{Name: "app-02"},
		{Name: "app-03"},
	}

	result := db.Create(apps)

	if(result.Error != nil) {
		return result.Error
	}

	return nil
}