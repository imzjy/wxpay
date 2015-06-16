package wxpay

import (
	"fmt"
	"reflect"
)

// ToXmlString convert the map[string]string to xml string
func ToXmlString(param map[string]string) string {
	xml := "<xml>"
	for k, v := range param {
		xml = xml + fmt.Sprintf("<%s>%s</%s>", k, v, k)
	}
	xml = xml + "</xml>"

	return xml
}

// ToMap convert the xml struct to map[string]string
func ToMap(in interface{}) (map[string]string, error) {
	out := make(map[string]string)

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// we only accept structs
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("ToMap only accepts structs; got %T", v)
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// gets us a StructField
		fi := typ.Field(i)
		if tagv := fi.Tag.Get("xml"); tagv != "" && tagv != "xml" {
			// set key of map to value in struct field
			out[tagv] = v.Field(i).String()
		}
	}
	return out, nil
}
