package util

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"reflect"
	"strings"
)

// todo sql injection check
func Update(db sqlx.Ext, table string, arg interface{}, uniquePropertyName string) string {
	fields := dBFields(arg) // e.g. []string{"id", "name", "description"}

	var newValues strings.Builder
	for i, field := range fields {
		if i != 0 {
			newValues.WriteString(",")
		}
		newValues.WriteString(field)
		newValues.WriteString("=:")
		newValues.WriteString(field)
	}

	sql := "UPDATE " + table + " set " + newValues.String() + " where " + uniquePropertyName + "=:" + uniquePropertyName
	return sql
}

func Insert(table string, arg interface{}) string {
	fields := dBFields(arg)         // e.g. []string{"id", "name", "description"}
	csv := fieldsCSV(fields)        // e.g. "id, name, description"
	csvc := fieldsCSVColons(fields) // e.g. ":id, :name, :description"
	sql := "INSERT INTO " + table + " (" + csv + ") VALUES (" + csvc + ")"
	return sql
}

func BulkUpsertWithQuery(db sqlx.Ext, query string, args interface{}) *sql.Result {
	const splicer = " on duplicate key update "

	// query got parsed weirdly
	queryParts := strings.Split(query, splicer)

	str, arg, err := sqlx.BindNamed(sqlx.QUESTION, queryParts[0], args)
	if err != nil {
		panic(err)
	}
	str += splicer + queryParts[1]

	res, err := db.Exec(str, arg...)
	if err != nil {
		panic(err)
	}
	return &res
}

func Upsert(table string, arg interface{}) string {
	fields := dBFields(arg)         // e.g. []string{"id", "name", "description"}
	csv := fieldsCSV(fields)        // e.g. "id, name, description"
	csvc := fieldsCSVColons(fields) // e.g. ":id, :name, :description"

	var newValues strings.Builder
	for i, field := range fields {
		if i != 0 {
			newValues.WriteString(",")
		}
		newValues.WriteString(field)
		newValues.WriteString("=values(")
		newValues.WriteString(field)
		newValues.WriteString(")")
	}

	sql := "INSERT INTO " + table + " (" + csv + ") VALUES (" + csvc + ") on duplicate key update " + newValues.String()
	return sql
}

func fieldsCSV(str []string) string {
	return fieldsCSVPrefix(str, "")
}

func fieldsCSVColons(str []string) string {
	return fieldsCSVPrefix(str, ":")
}

func fieldsCSVPrefix(str []string, prefix string) string {
	var build strings.Builder
	for i, s := range str {
		if i != 0 {
			build.WriteString(",")
		}
		build.WriteString(prefix)
		build.WriteString(s)
	}
	return build.String()
}

// dBFields reflects on a struct and returns the values of fields with `db` tags,
// or a map[string]interface{} and returns the keys.
func dBFields(values interface{}) []string {
	v := reflect.ValueOf(values)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	var fields []string
	if v.Kind() == reflect.Struct {
		for i := 0; i < v.NumField(); i++ {
			field := v.Type().Field(i).Tag.Get("db")
			if field != "" {
				fields = append(fields, field)
			}
		}
		return fields
	}
	if v.Kind() == reflect.Map {
		for _, keyv := range v.MapKeys() {
			fields = append(fields, keyv.String())
		}
		return fields
	}
	panic(fmt.Errorf("dBFields requires a struct or a map, found: %s", v.Kind().String()))
}
