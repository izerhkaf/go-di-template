package helper

import (
	"crypto/rand"
	"database/sql"
	"math/big"
	"reflect"
	"unicode"
)

func ToNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: s, Valid: true}
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_+=<>?/"
const charsetWithoutSymbo = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateRandomString(length int64, withSymbol bool) string {
	result := make([]byte, length)
	if withSymbol {
		for i := range result {
			index, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
			result[i] = charset[index.Int64()]
		}
	} else {
		for i := range result {
			index, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charsetWithoutSymbo))))
			result[i] = charsetWithoutSymbo[index.Int64()]
		}
	}
	return string(result)
}

func ReplaceDashWithEmptyString[T any](input T) T {
	val := reflect.ValueOf(&input).Elem()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		if field.Kind() == reflect.String && field.CanSet() {
			if field.String() == "-" {
				field.SetString("")
			}
		}

		if field.Kind() == reflect.Struct {
			nested := ReplaceDashWithEmptyString(field.Interface())
			field.Set(reflect.ValueOf(nested))
		}
	}

	return input
}

func StartUpcase(input string) string {
	if len(input) == 0 {
		return input
	}

	runes := []rune(input)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)

}
