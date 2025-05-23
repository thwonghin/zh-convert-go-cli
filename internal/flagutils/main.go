package flagutils

import (
	"flag"
	"fmt"
	"reflect"
	"strings"
)

func BindFlagsFromStruct(s any, stringFlags map[string]*string, boolFlags map[string]*bool, intFlags map[string]*int, stringSliceFlags map[string]*string) {
	t := reflect.TypeOf(s).Elem()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		name := field.Tag.Get("json")
		if name == "" || name == "-" {
			continue
		}
		name = strings.Split(name, ",")[0]

		switch field.Type.Kind() {
		case reflect.Ptr:
			elem := field.Type.Elem()
			switch elem.Kind() {
			case reflect.String:
				stringFlags[name] = flag.String(name, "", fmt.Sprintf("set %s", name))
			case reflect.Int:
				intFlags[name] = flag.Int(name, 0, fmt.Sprintf("set %s", name))
			case reflect.Bool:
				boolFlags[name] = flag.Bool(name, false, fmt.Sprintf("set %s", name))
			case reflect.Slice:
				if elem.Elem().Kind() == reflect.String {
					stringSliceFlags[name] = flag.String(name, "", fmt.Sprintf("comma-separated list for %s", name))
				}
			}
		case reflect.Slice:
			if field.Type.Elem().Kind() == reflect.String {
				stringSliceFlags[name] = flag.String(name, "", fmt.Sprintf("comma-separated list for %s", name))
			}
		case reflect.String:
			stringFlags[name] = flag.String(name, "", fmt.Sprintf("set %s", name))
		case reflect.Bool:
			boolFlags[name] = flag.Bool(name, false, fmt.Sprintf("set %s", name))
		case reflect.Int:
			intFlags[name] = flag.Int(name, 0, fmt.Sprintf("set %s", name))
		}
	}
}

func PopulateStructFromFlags(ptr any, stringFlags map[string]*string, boolFlags map[string]*bool, intFlags map[string]*int, stringSliceFlags map[string]*string) {
	v := reflect.ValueOf(ptr).Elem()
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		name := field.Tag.Get("json")
		if name == "" || name == "-" {
			continue
		}
		name = strings.Split(name, ",")[0]
		fv := v.Field(i)

		switch fv.Kind() {
		case reflect.Ptr:
			elemKind := fv.Type().Elem().Kind()
			switch elemKind {
			case reflect.String:
				if val, ok := stringFlags[name]; ok && *val != "" {
					fv.Set(reflect.ValueOf(val))
				}
			case reflect.Int:
				if val, ok := intFlags[name]; ok {
					fv.Set(reflect.ValueOf(val))
				}
			case reflect.Bool:
				if val, ok := boolFlags[name]; ok {
					fv.Set(reflect.ValueOf(val))
				}
			case reflect.Slice:
				if val, ok := stringSliceFlags[name]; ok && *val != "" {
					list := strings.Split(*val, ",")
					fv.Set(reflect.ValueOf(&list))
				}
			}
		case reflect.Slice:
			if val, ok := stringSliceFlags[name]; ok && *val != "" {
				list := strings.Split(*val, ",")
				fv.Set(reflect.ValueOf(list))
			}
		case reflect.String:
			if val, ok := stringFlags[name]; ok {
				fv.SetString(*val)
			}
		case reflect.Int:
			if val, ok := intFlags[name]; ok {
				fv.SetInt(int64(*val))
			}
		case reflect.Bool:
			if val, ok := boolFlags[name]; ok {
				fv.SetBool(*val)
			}
		}
	}
}
