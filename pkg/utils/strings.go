package utils

import (
	"strconv"
)

func IntToString(i int) string {
	return strconv.FormatInt(int64(i), 10)
}

func Int32ToString(i int32) string {
	return strconv.FormatInt(int64(i), 10)
}

func Uint32ToString(i uint32) string {
	return strconv.FormatInt(int64(i), 10)
}

func StringToInt(i string) int {
	j, _ := strconv.Atoi(i)
	return j
}

func StringToInt64(i string) int64 {
	j, _ := strconv.ParseInt(i, 10, 64)
	return j
}

func StringToInt32(i string) int32 {
	j, _ := strconv.ParseInt(i, 10, 64)
	return int32(j)
}

// judge a string whether in the string list
func IsContain(target string, List []string) bool {
	for _, element := range List {
		if target == element {
			return true
		}
	}
	return false
}

func IsContainInt32(target int32, List []int32) bool {
	for _, element := range List {
		if target == element {
			return true
		}
	}
	return false
}

func IsContainInt(target int, List []int) bool {
	for _, element := range List {
		if target == element {
			return true
		}
	}
	return false
}

func InterfaceArrayToStringArray(data []interface{}) (i []string) {
	for _, param := range data {
		i = append(i, param.(string))
	}
	return i
}
