package utils

import (
	"reflect"
	"strconv"
)

func IsValid(data any) (isValid bool, err error) {
	isValid = true
	dataType := reflect.TypeOf(data)

	for i := 0; i < dataType.NumField(); i++ {
		field := dataType.Field(i)
		requiredTag := field.Tag.Get("required")
		maxTag := field.Tag.Get("max")
		if requiredTag != "" {
			isRequired := false
			isRequired, err = strconv.ParseBool(requiredTag)
			if err != nil {
				isValid = false
				return
			} else if isRequired {
				switch value := reflect.ValueOf(data).Field(i).Interface(); convertedValue := value.(type) {
				case string:
					isValid = convertedValue != ""
					if !isValid {
						return
					}
				default:
					isValid = convertedValue != nil
					if !isValid {
						return
					}
				}
			}
		}
		if maxTag != "" {
			var maxValue int64 = 0
			maxValue, err = strconv.ParseInt(maxTag, 10, 64)
			if err != nil {
				isValid = false
				return
			} else if maxValue != 0 {
				switch value := reflect.ValueOf(data).Field(i).Interface(); convertedValue := value.(type) {
				case string:
					isValid = int64(len(convertedValue)) <= maxValue
					if !isValid {
						return
					}
				case int:
					isValid = int64(convertedValue) <= maxValue
					if !isValid {
						return
					}
				case int64:
					isValid = convertedValue <= maxValue
					if !isValid {
						return
					}
				case int32:
					isValid = int64(convertedValue) <= maxValue
					if !isValid {
						return
					}
				case int16:
					isValid = int64(convertedValue) <= maxValue
					if !isValid {
						return
					}
				case int8:
					isValid = int64(convertedValue) <= maxValue
					if !isValid {
						return
					}
				case uint:
					isValid = int64(convertedValue) <= maxValue
					if !isValid {
						return
					}
				case uint64:
					isValid = int64(convertedValue) <= maxValue
					if !isValid {
						return
					}
				case uint32:
					isValid = int64(convertedValue) <= maxValue
					if !isValid {
						return
					}
				case uint16:
					isValid = int64(convertedValue) <= maxValue
					if !isValid {
						return
					}
				case uint8:
					isValid = int64(convertedValue) <= maxValue
					if !isValid {
						return
					}
				default:
					isValid = convertedValue != nil
					if !isValid {
						return
					}
				}
			}
		}
	}
	return
}
