package utils

import (
	"errors"
	"strconv"
	"strings"

	"github.com/fatih/structs"
)

func ConvertStrToInt(s string) (int, error) {
	switch s {
	case "0":
		return 0, errors.New("args not is zero")
	default:
		return strconv.Atoi(s)
	}
}

func SplitUriStrToMapTmp(s, sep1, sep2, sep3 string) map[string]string {
	var m = make(map[string]string, 5)
	arr := strings.Split(s, sep1)
	for _, value := range strings.Split(arr[1], sep2) {
		argsArr := strings.Split(value, sep3)
		if argsArr[1] != "" {
			m[argsArr[0]] = argsArr[1]
		}
	}
	return m

}

// struct to map， safe(exclude zero)
func ConvertStructToMapSafe(s *structs.Struct) (m map[string]interface{}) {
	m = make(map[string]interface{}, 5)
	for _, f := range s.Fields() {
		if f.IsExported() {
			if !f.IsZero() {
				m[f.Name()] = f.Value()
			}
		}
	}
	return
}

// struct to map， safe(exclude zero)
func ConvertStructToMap(s interface{}) (m map[string]interface{}) {
	return structs.Map(s)
}
