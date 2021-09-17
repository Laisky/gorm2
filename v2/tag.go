package gorm

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// ConditionTagFunc func(db, columnName, value)
type ConditionTagFunc func(*DB, string, interface{}) (*DB, error)

var (
	customConditionTags = map[string]ConditionTagFunc{}
)

// AddCustomSQLTagCondition add custom sql tag condition func
func AddCustomSQLTagCondition(op string, f ConditionTagFunc) {
	customConditionTags[op] = f
}

// ApplySQLCondition add tag for ORM struct
//
//   HostName  string  `sql:"column:h.name;op:eq"`
//
// will generate DB with `Where("h.name = ?", <value of this field>)`
func ApplySQLCondition(db *DB, req interface{}) (*DB, error) {
	sv := reflect.ValueOf(req)
	if sv.Kind() == reflect.Ptr {
		sv = sv.Elem()
	}

	var err error
	st := sv.Type()
	for i := 0; i < st.NumField(); i++ {
		if sv.Field(i).IsZero() {
			continue
		}

		tag := st.Field(i).Tag.Get("sql")
		if tag == "" {
			continue
		}

		tagMap := map[string]string{}
		for _, tv := range strings.Split(tag, ";") {
			vs := strings.Split(tv, ":")
			if len(vs) != 2 {
				return nil, fmt.Errorf("unknown tag `%s`", tag)
			}

			tagMap[vs[0]] = vs[1]
		}

		column := tagMap["column"]
		if column == "" {
			continue
		}

		v := sv.Field(i).Interface()
		if db, err = applySQLOp(db, v, tagMap, column); err != nil {
			return nil, err
		}
	}

	return db, nil
}

func applySQLOp(db *DB, v interface{}, tagMap map[string]string, column string) (*DB, error) {
	var err error
	switch op := tagMap["op"]; op {
	case "eq":
		db = db.Where(fmt.Sprintf("%s = ?", column), v)
	case "lt":
		db = db.Where(fmt.Sprintf("%s < ?", column), v)
	case "lte":
		db = db.Where(fmt.Sprintf("%s <= ?", column), v)
	case "gt":
		db = db.Where(fmt.Sprintf("%s > ?", column), v)
	case "gte":
		db = db.Where(fmt.Sprintf("%s >= ?", column), v)
	case "in":
		db = db.Where(fmt.Sprintf("%s IN (?)", column), v)
	case "ints":
		db, err = SQLTagSplitInInt(db, column, v)
	case "strs":
		db, err = SQLTagSplitInStr(db, column, v)
	case "like":
		db = db.Where(fmt.Sprintf("%s like ?", column), MYSQLLike(fmt.Sprint(v)))
	case "like-bin":
		db = db.Where(fmt.Sprintf("BINARY %s like ?", column), MYSQLLike(fmt.Sprint(v)))
	default:
		f, ok := customConditionTags[op]
		if !ok {
			return nil, fmt.Errorf("unknown tag `%s` for `sql:`", op)
		}

		db, err = f(db, column, v)
	}

	return db, err
}

// MYSQLLike 生成 mysql like 语法
func MYSQLLike(v string) string {
	v = strings.ReplaceAll(v, `\`, `\\`)
	v = strings.ReplaceAll(v, `%`, `\%`)
	return "%" + strings.ReplaceAll(v, "_", `\_`) + "%"
}

func SQLTagSplitInStr(db *DB, column string, val interface{}) (*DB, error) {
	vs := TrimEleSpaceAndRemoveEmpty(strings.Split(fmt.Sprint(val), ","))
	if len(vs) != 0 {
		db = db.Where(fmt.Sprintf("%s IN (?)", column), vs)
	}

	return db, nil
}

func SQLTagSplitInInt(db *DB, column string, val interface{}) (*DB, error) {
	var vals []int
	for _, vs := range TrimEleSpaceAndRemoveEmpty(strings.Split(fmt.Sprint(val), ",")) {
		vi, err := strconv.Atoi(vs)
		if err != nil {
			return nil, err
		}

		vals = append(vals, vi)
	}

	if len(vals) != 0 {
		db = db.Where(fmt.Sprintf("%s IN (?)", column), vals)
	}

	return db, nil
}

// TrimEleSpaceAndRemoveEmpty remove duplicate string in slice
func TrimEleSpaceAndRemoveEmpty(vs []string) (r []string) {
	for _, v := range vs {
		v = strings.TrimSpace(v)
		if v != "" {
			r = append(r, v)
		}
	}

	return
}
