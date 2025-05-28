package str_util

import (
	"errors"
	"time"
)

// IsBlank check if string is blank
func IsBlank(str string) bool {
	if str == "" {
		return true
	}
	return false
}

// IsNotBlank check if string is not blank
func IsNotBlank(str string) bool {
	return !IsBlank(str)
}

// ParseDatetime parse datetime
func ParseDatetime(val string) (*time.Time, error) {
	// 定义可能的日期格式
	layouts := []string{
		"2006-01-02",
		"2006/01/02",
		"02-Jan-2006",
		"01-02-2006",
		"2006-01-02 15:04",
		"2006-01-02 15:04:05",
		"2006-01-02 15:04:05.000",
		"2006-01-02 15:04:05.000000",
		"2006-01-02 15:04:05.000000000",
		"2006-01-02 15:04:05 +0000 UTC",
		"2006-01-02T15:04:05.000000000+08:00",
		"20060102",
		"2006010201",
		"200601020101",
		"20060102010101",
		"2026年",
		"2026年1月",
		"2026年1月1日",
	}
	for _, layout := range layouts {
		parseTime, err := time.Parse(layout, val)
		if err == nil {
			return &parseTime, nil
		}
	}
	return nil, errors.New("parse datetime error")
}
