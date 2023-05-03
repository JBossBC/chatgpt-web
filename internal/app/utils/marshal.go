package utils

import (
	"fmt"
	"reflect"
)

// ban the normal type exists pointer type
func Marshal(data map[string]interface{}, structTemplate reflect.Type) (result reflect.Value, err error) {
	defer func() {
		if errRecover := recover(); errRecover != nil {
			err = fmt.Errorf("utils.Marshal analsis failed:%v", errRecover)
		}
	}()
	if structTemplate.Kind() == reflect.Map {
		newResult := data
		return reflect.ValueOf(newResult), nil
	}
	result = reflect.New(structTemplate).Elem()
	fieldNumber := structTemplate.NumField()
	for i := 0; i < fieldNumber; i++ {
		if !result.Field(i).CanSet() {
			continue
		}
		fieldStruct := structTemplate.Field(i)

		if value, ok := data[fieldStruct.Name]; ok {
			elemType := fieldStruct.Type
			var isPtr = false
			switch elemType.Kind() {
			case reflect.Ptr:
				// var subMap map[string]interface{}
				// if subMap, ok = value.(map[string]interface{}); !ok {
				// 	return result, fmt.Errorf("%s should be map[string]interface{} type", elemType)
				// }
				// subValue, err := Marshal(subMap, fieldStruct.Type.Elem())
				// if err != nil {
				// 	return result, err
				// }
				// result.Field(i).Set(subValue)
				elemType = fieldStruct.Type.Elem()
				isPtr = true
				fallthrough
			case reflect.Struct:
				var actualStruct map[string]interface{}
				if mapRes, ok := value.(map[string]interface{}); ok {
					actualStruct = mapRes
				} else {
					if isPtr {
						value = &value
					}
					result.Field(i).Set(reflect.ValueOf(value))
					continue
				}
				subStruct, err := Marshal(actualStruct, fieldStruct.Type)
				if err != nil {
					return result, err
				}
				result.Field(i).Set(subStruct)
			case reflect.Map:
				subValue, err := Marshal(value.(map[string]interface{}), fieldStruct.Type)
				if err != nil {
					return result, err
				}
				result.Field(i).Set(subValue)
			case reflect.Invalid:
				return result, fmt.Errorf("%s is invalid type", elemType)
			case reflect.Chan:
				panic("cant support chan type,defend the memoey safe")
			default:
				result.Field(i).Set(reflect.ValueOf(value))
			}

		}
	}
	return result, nil
}

func UnMarshal(value reflect.Value) (map[string]interface{}, error) {
	var result = make(map[string]interface{})
	if res, ok := value.Interface().(map[string]interface{}); ok {
		return res, nil
	}
	elemSize := value.NumField()
	for i := 0; i < elemSize; i++ {
		curElem := value.Field(i)
		switch curElem.Kind() {
		case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int8:
			result[curElem.Type().Name()] = curElem.Int()
		case reflect.Complex64, reflect.Complex128:
			result[curElem.Type().Name()] = curElem.Complex()
		case reflect.Float32, reflect.Float64:
			result[curElem.Type().Name()] = curElem.Float()
		case reflect.Ptr:
			value, err := UnMarshal(curElem.Elem())
			if err != nil {
				return result, err
			}
			result[curElem.Type().Name()] = value
		case reflect.Struct:
			value, err := UnMarshal(curElem)
			if err != nil {
				return result, err
			}
			result[curElem.Type().Name()] = value
		case reflect.Bool:
			result[curElem.Type().Name()] = curElem.Bool()
		case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint8:
			result[curElem.Type().Name()] = curElem.Uint()
		case reflect.Array, reflect.Slice:
			result[curElem.Type().Name()] = curElem.Slice(0, curElem.Len())
		case reflect.Map:
			result[curElem.Type().Name()] = curElem.Interface().(map[string]interface{})
		case reflect.String:
			result[curElem.Type().Name()] = curElem.String()
		default:
			return result, fmt.Errorf("%s cant support this type", curElem.Type().Name())
		}
	}
	return result, nil
}
