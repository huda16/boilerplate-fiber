package database

import (
	"boilerplate-api/app/models"
	"boilerplate-api/config"
	"fmt"
	"log"

	do "github.com/dzwvip/gorm-oracle"
	oracle "github.com/godoes/gorm-oracle"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func Connect() {
    options := map[string]string{
		"CONNECTION TIMEOUT": "90",
		"LANGUAGE":           "SIMPLIFIED CHINESE",
		"TERRITORY":          "CHINA",
		"SSL":                "false",
	}
	// oracle://user:password@127.0.0.1:1521/service
	url := oracle.BuildUrl(config.GetEnv("DB_HOST"), 1539, config.GetEnv("DB_SERVICE"), config.GetEnv("DB_USER"), config.GetEnv("DB_PASSWORD"), options)
	dialector := oracle.New(oracle.Config{
		DSN:                     url,
		IgnoreCase:              false, // query conditions are not case-sensitive
		NamingCaseSensitive:     true,  // whether naming is case-sensitive
		VarcharSizeIsCharLength: true,  // whether VARCHAR type size is character length, defaulting to byte length

		// RowNumberAliasForOracle11 is the alias for ROW_NUMBER() in Oracle 11g, defaulting to ROW_NUM
		RowNumberAliasForOracle11: "ROW_NUM",
	})
	db, err := gorm.Open(dialector, &gorm.Config{
		SkipDefaultTransaction:                   true, // 是否禁用默认在事务中执行单次创建、更新、删除操作
		DisableForeignKeyConstraintWhenMigrating: true, // 是否禁止在自动迁移或创建表时自动创建外键约束
		// 自定义命名策略
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase:         true, // 是否不自动转换小写表名
			IdentifierMaxLength: 30,   // Oracle: 30, PostgreSQL:63, MySQL: 64, SQL Server、SQLite、DM: 128
		},
		PrepareStmt:     false, // 创建并缓存预编译语句，启用后可能会报 ORA-01002 错误
		CreateBatchSize: 50,    // 插入数据默认批处理大小
	})
	if err != nil {
		log.Fatalf("Error creating GORM connection: %v", err)
	}

	// set session parameters
	if sqlDB, err := db.DB(); err == nil {
		_, _ = oracle.AddSessionParams(sqlDB, map[string]string{
			"TIME_ZONE":               "+08:00",                       // ALTER SESSION SET TIME_ZONE = '+08:00';
			"NLS_DATE_FORMAT":         "YYYY-MM-DD",                   // ALTER SESSION SET NLS_DATE_FORMAT = 'YYYY-MM-DD';
			"NLS_TIME_FORMAT":         "HH24:MI:SSXFF",                // ALTER SESSION SET NLS_TIME_FORMAT = 'HH24:MI:SS.FF3';
			"NLS_TIMESTAMP_FORMAT":    "YYYY-MM-DD HH24:MI:SSXFF",     // ALTER SESSION SET NLS_TIMESTAMP_FORMAT = 'YYYY-MM-DD HH24:MI:SS.FF3';
			"NLS_TIME_TZ_FORMAT":      "HH24:MI:SS.FF TZR",            // ALTER SESSION SET NLS_TIME_TZ_FORMAT = 'HH24:MI:SS.FF3 TZR';
			"NLS_TIMESTAMP_TZ_FORMAT": "YYYY-MM-DD HH24:MI:SSXFF TZR", // ALTER SESSION SET NLS_TIMESTAMP_TZ_FORMAT = 'YYYY-MM-DD HH24:MI:SS.FF3 TZR';
		})
	}

	// do somethings
    log.Println("Database Connected")
}

func Connect2() {
    dsn := fmt.Sprintf("oracle://%s:%s@%s:%s/%s",
        config.GetEnv("DB_USER"),
        config.GetEnv("DB_PASSWORD"),
        config.GetEnv("DB_HOST"),
        config.GetEnv("DB_PORT"),
        config.GetEnv("DB_SERVICE"))

    db, err := gorm.Open(do.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Error creating GORM connection: %v", err)
    } 
    
    // do somethings
    DB = db
    log.Println("Database 2 Connected")
    db.Debug().AutoMigrate(models.Example{})
}