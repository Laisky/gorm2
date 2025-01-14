package utils

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"unicode"
)

var gormSourceDir string

var logFileIgnored = []*regexp.Regexp{
	regexp.MustCompile(`(?:Laisky|jinzhu|gorm\.io)/gorm(?:@.*)?/.*.go`),
}

var logFileNotIgnore = regexp.MustCompile(`(?:Laisky|jinzhu|gorm\.io)/gorm(?:@.*)?/.*test.go`)

// AddLogFileIgnoreStackPattern file name pattern that will not ignored
func AddLogFileIgnoreStackPattern(logFilePattern ...*regexp.Regexp) {
	logFileIgnored = append(logFileIgnored, logFilePattern...)
}

func isLogFileIgnore(file string) bool {
	for _, r := range logFileIgnored {
		if r.MatchString(file) {
			return true
		}
	}

	return false
}

func init() {
	_, file, _, _ := runtime.Caller(0)
	// compatible solution to get gorm source directory with various operating systems
	gormSourceDir = regexp.MustCompile(`utils.utils\.go`).ReplaceAllString(file, "")
}

// FileWithLineNum return the file name and line number of the current file
func FileWithLineNum() string {
	// the second caller usually from gorm internal, so set i start from 2
	for i := 2; i < 15; i++ {
		_, file, line, ok := runtime.Caller(i)
		if ok && (logFileNotIgnore.MatchString(file) || !isLogFileIgnore(file)) {
			return file + ":" + strconv.FormatInt(int64(line), 10)
		}
	}

	return ""
}

func IsValidDBNameChar(c rune) bool {
	return !unicode.IsLetter(c) && !unicode.IsNumber(c) && c != '.' && c != '*' && c != '_' && c != '$' && c != '@'
}

func CheckTruth(val interface{}) bool {
	if v, ok := val.(bool); ok {
		return v
	}

	if v, ok := val.(string); ok {
		v = strings.ToLower(v)
		return v != "false"
	}

	return !reflect.ValueOf(val).IsZero()
}

func ToStringKey(values ...interface{}) string {
	results := make([]string, len(values))

	for idx, value := range values {
		if valuer, ok := value.(driver.Valuer); ok {
			value, _ = valuer.Value()
		}

		switch v := value.(type) {
		case string:
			results[idx] = v
		case []byte:
			results[idx] = string(v)
		case uint:
			results[idx] = strconv.FormatUint(uint64(v), 10)
		default:
			results[idx] = fmt.Sprint(reflect.Indirect(reflect.ValueOf(v)).Interface())
		}
	}

	return strings.Join(results, "_")
}

func AssertEqual(src, dst interface{}) bool {
	if !reflect.DeepEqual(src, dst) {
		if valuer, ok := src.(driver.Valuer); ok {
			src, _ = valuer.Value()
		}

		if valuer, ok := dst.(driver.Valuer); ok {
			dst, _ = valuer.Value()
		}

		return reflect.DeepEqual(src, dst)
	}
	return true
}

func ToString(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	case int:
		return strconv.FormatInt(int64(v), 10)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	}
	return ""
}

func ExistsIn(a string, list *[]string) bool {
	if list == nil {
		return false
	}
	for _, b := range *list {
		if b == a {
			return true
		}
	}
	return false
}
