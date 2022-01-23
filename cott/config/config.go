package config

import "github.com/sirupsen/logrus"

type Config struct {
	Port uint `default:"1004" env:"PORT"`

	ReportFilePath string `default:"report.json" env:"TESTER_REPORT_FILE_PATH"`

	Log                  LogConfig
	DatabaseTesterConfig DatabaseTesterConfig
}

type LogConfig struct {
	Level            logrus.Level `default:"info" env:"LOG_LEVEL"`
	FilePath         string       `default:"/var/log/cott/cott.log" env:"LOG_FILE_PATH"`
	MaxFileSizeInMb  int          `default:"10" env:"LOG_MAX_FILE_SIZE_IN_MB"`
	MaxFilesCount    int          `default:"7" env:"LOG_MAX_FILES_COUNT"`
	MaxFileAgeInDays int          `default:"7" env:"LOG_MAX_FILE_AGE_IN_DAYS"`
	CompressOldFiles bool         `default:"true" env:"LOG_COMPRESS_OLD_FILES"`
}

type DatabaseTesterConfig struct {
	DatabaseName string `default:"cott_db" env:"DATABASE_TESTER_DATABASE_NAME"`
}
