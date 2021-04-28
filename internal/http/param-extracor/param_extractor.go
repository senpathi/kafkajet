package param_extracor

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
)

type extractorFunc func(key string) (string, bool)

type ParamExtractor interface {
	ExtractHeaders(v interface{}, req *http.Request) error
	ExtractQuery(v interface{}, req *http.Request) error
	ExtractForm(v interface{}, req *http.Request) error
}

type paramExtractor struct{}

func NewParamExtractor() ParamExtractor {
	return paramExtractor{}
}

func (p paramExtractor) ExtractHeaders(v interface{}, req *http.Request) error {
	return p.extract(v, req, func(key string) (string, bool) {
		str := req.Header.Get(key)
		if str == `` {
			return str, false
		}
		return str, true
	})
}

func (p paramExtractor) ExtractQuery(v interface{}, req *http.Request) error {
	return p.extract(v, req, func(key string) (string, bool) {
		str := req.URL.Query().Get(key)
		if str == `` {
			return str, false
		}
		return str, true
	})
}

func (p paramExtractor) ExtractForm(v interface{}, req *http.Request) error {
	return p.extract(v, req, func(key string) (string, bool) {
		str := req.Form.Get(key)
		if str == `` {
			return str, false
		}
		return str, true
	})
}

func (p paramExtractor) extract(v interface{}, req *http.Request, keyExtractor extractorFunc) error {
	t := reflect.TypeOf(v)
	if t.Kind() != reflect.Ptr || v == nil {
		return ErrorNotAssignable(fmt.Errorf(`type of %v is not assignabale, required object reference`, t))
	}

	elem := reflect.ValueOf(v).Elem()
	if elem.Kind() != reflect.Struct {
		return ErrorUnSupportedType(fmt.Errorf(`type of %v is not extractable, required struct object`, elem.Type().String())) //todo err
	}

	t = elem.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		tag, ok := field.Tag.Lookup(`param`)
		if !ok || tag == `-` {
			continue
		}

		fmt.Println(tag, field.Type.Name())
		valueStr, ok := keyExtractor(tag)
		if !ok {
			continue
		}

		switch field.Type.Kind() {
		case reflect.String:
			fmt.Println(`kind string : `, tag)
			elem.Field(i).Set(reflect.ValueOf(valueStr))

		case reflect.Bool:
			fmt.Println(`kind bool : `, tag)
			value, err := strconv.ParseBool(valueStr)
			if err != nil {
				return ErrorUnmarshalType(fmt.Errorf(`error unmarshalling [%v] into [bool] due to %v`, valueStr, err))
			}
			elem.Field(i).Set(reflect.ValueOf(value))

		case reflect.Int32:
			fmt.Println(`kind int32 : `, tag)
			value, err := strconv.Atoi(valueStr)
			if err != nil {
				return ErrorUnmarshalType(fmt.Errorf(`error unmarshalling [%v] into [int32] due to %v`, valueStr, err))
			}
			elem.Field(i).Set(reflect.ValueOf(int32(value)))

		case reflect.Int:
			fmt.Println(`kind int : `, tag)
			value, err := strconv.Atoi(valueStr)
			if err != nil {
				return ErrorUnmarshalType(fmt.Errorf(`error unmarshalling [%v] into [int] due to %v`, valueStr, err))
			}
			elem.Field(i).Set(reflect.ValueOf(value))

		case reflect.Int64:
			fmt.Println(`kind int64 : `, tag)
			value, err := strconv.ParseInt(valueStr, 10, 64)
			if err != nil {
				return ErrorUnmarshalType(fmt.Errorf(`error unmarshalling [%v] into [int64] due to %v`, valueStr, err))
			}
			elem.Field(i).Set(reflect.ValueOf(value))

		case reflect.Float32:
			fmt.Println(`kind float32 : `, tag)
			value, err := strconv.ParseFloat(valueStr, 32)
			if err != nil {
				return ErrorUnmarshalType(fmt.Errorf(`error unmarshalling [%v] into [float32] due to %v`, valueStr, err))
			}
			elem.Field(i).Set(reflect.ValueOf(float32(value)))

		case reflect.Float64:
			fmt.Println(`kind float64 : `, tag)
			value, err := strconv.ParseFloat(valueStr, 64)
			if err != nil {
				return ErrorUnmarshalType(fmt.Errorf(`error unmarshalling [%v] into [float64] due to %v`, valueStr, err))
			}
			elem.Field(i).Set(reflect.ValueOf(value))

		default:
			return ErrorUnSupportedParamType(errors.New(`unsupported param extractor type`))
		}
		fmt.Println(fmt.Sprintf(`%v`, field))
	}

	return nil
}
