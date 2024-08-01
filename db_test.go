package db

import "testing"

func TestGorm(t *testing.T) {
	count := 0
	Gorm.Select("count(*)").Table("app_user").Scan(&count)
	t.Log(count)
}
