module github.com/Laisky/gorm/v2/tests

go 1.14

require (
	github.com/google/uuid v1.3.0
	github.com/jinzhu/now v1.1.2
	github.com/lib/pq v1.10.2
	gorm.io/driver/mysql v1.1.2
	gorm.io/driver/postgres v1.1.0
	gorm.io/driver/sqlite v1.1.4
	gorm.io/driver/sqlserver v1.0.8
	github.com/Laisky/gorm/v2 v1.21.13
)

replace github.com/Laisky/gorm/v2 => ../