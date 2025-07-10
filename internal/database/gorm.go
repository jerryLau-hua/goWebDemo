package database

import (
	"awesomeProject/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewGormConnection 负责根据配置创建GORM数据库连接池
func NewGormConnection(dbConfig config.DatabaseConfig) (*gorm.DB, error) {
	// dsn 来自于我们的配置文件
	dsn := dbConfig.DSN

	// 使用GORM的MySQL驱动来打开数据库连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 获取底层的 *sql.DB 对象来设置连接池参数
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// 设置从配置中读取的连接池参数
	sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConns)
	sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConns)

	// 可以选择在这里 Ping 数据库以验证连接
	if err = sqlDB.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
