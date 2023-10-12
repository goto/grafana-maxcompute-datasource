package converters

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"time"

	data2 "github.com/aliyun/aliyun-odps-go-sdk/odps/data"
	"github.com/aliyun/aliyun-odps-go-sdk/sqldriver"
	"github.com/grafana/grafana-plugin-sdk-go/data"
	"github.com/grafana/grafana-plugin-sdk-go/data/sqlutil"
)

type Converter struct {
	convert    func(in interface{}) (interface{}, error)
	fieldType  data.FieldType
	matchRegex *regexp.Regexp
	scanType   reflect.Type
}

func invalidType(name string) error {
	return fmt.Errorf("invalid type - %s", name)
}

var matchRegexes = map[string]*regexp.Regexp{
	"ARRAY":   regexp.MustCompile(`^ARRAY<.*>`),
	"DECIMAL": regexp.MustCompile(`^DECIMAL`),
	"VARCHAR": regexp.MustCompile(`^VARCHAR`),
	"CHAR":    regexp.MustCompile(`^CHAR`),
	"MAP":     regexp.MustCompile(`^MAP<.*>`),
	"STRUCT":  regexp.MustCompile(`^STRUCT<.*>`),
}

var Converters = map[string]Converter{
	"BIGINT": {
		fieldType: data.FieldTypeNullableInt64,
		convert: func(in interface{}) (interface{}, error) {
			if in == nil {
				return int64(0), nil
			}

			if v, ok := in.(int64); ok {
				return v, nil
			}

			if v, ok := in.(*sqldriver.NullInt64); ok {
				if v.IsNull() {
					return (*int64)(nil), nil
				}
				return &v.Int64, nil
			}

			return nil, invalidType("BIGINT")
		},
	},
	"INT": {
		fieldType: data.FieldTypeNullableInt32,
		convert: func(in interface{}) (interface{}, error) {
			if in == nil {
				return int32(0), nil
			}

			if v, ok := in.(int32); ok {
				return v, nil
			}

			if v, ok := in.(*sqldriver.NullInt32); ok {
				if v.IsNull() {
					return (*int32)(nil), nil
				}
				return &v.Int32, nil
			}

			return nil, invalidType("INT")
		},
	},
	"SMALLINT": {
		fieldType: data.FieldTypeNullableInt16,
		convert: func(in interface{}) (interface{}, error) {
			if in == nil {
				return int16(0), nil
			}

			if v, ok := in.(int16); ok {
				return v, nil
			}

			if v, ok := in.(*sqldriver.NullInt16); ok {
				if v.IsNull() {
					return (*int16)(nil), nil
				}
				return &v.Int16, nil
			}

			return nil, invalidType("SMALLINT")
		},
	},
	"TINYINT": {
		fieldType: data.FieldTypeNullableInt8,
		convert: func(in interface{}) (interface{}, error) {
			if in == nil {
				return int8(0), nil
			}

			if v, ok := in.(int8); ok {
				return v, nil
			}

			if v, ok := in.(*sqldriver.NullInt8); ok {
				if v.IsNull() {
					return (*int8)(nil), nil
				}
				return &v.Int8, nil
			}

			return nil, invalidType("TINYINT")
		},
	},
	"DOUBLE": {
		fieldType: data.FieldTypeNullableFloat64,
		convert: func(in interface{}) (interface{}, error) {
			if in == nil {
				return float64(0), nil
			}

			if v, ok := in.(float64); ok {
				return v, nil
			}

			if v, ok := in.(*sqldriver.NullFloat64); ok {
				if v.IsNull() {
					return (*float64)(nil), nil
				}
				return &v.Float64, nil
			}

			return nil, invalidType("DOUBLE")
		},
	},
	"FLOAT": {
		fieldType: data.FieldTypeNullableFloat32,
		convert: func(in interface{}) (interface{}, error) {
			if in == nil {
				return float32(0), nil
			}

			if v, ok := in.(float32); ok {
				return v, nil
			}

			if v, ok := in.(*sqldriver.NullFloat32); ok {
				if v.IsNull() {
					return (*float32)(nil), nil
				}
				return &v.Float32, nil
			}

			return nil, invalidType("FLOAT")
		},
	},
	"STRING": {
		fieldType: data.FieldTypeNullableString,
		convert:   stringConverter,
	},
	"CHAR": {
		fieldType:  data.FieldTypeNullableString,
		matchRegex: matchRegexes["CHAR"],
		convert:    stringConverter,
	},
	"VARCHAR": {
		fieldType:  data.FieldTypeNullableString,
		matchRegex: matchRegexes["VARCHAR"],
		convert:    stringConverter,
	},
	"BINARY": {
		fieldType: data.FieldTypeNullableJSON,
		convert:   jsonConverter,
	},
	"BOOLEAN": {
		fieldType: data.FieldTypeNullableBool,
		convert: func(in interface{}) (interface{}, error) {
			if in == nil {
				return false, nil
			}

			if v, ok := in.(bool); ok {
				return v, nil
			}

			if v, ok := in.(*sqldriver.NullBool); ok {
				if v.IsNull() {
					return (*bool)(nil), nil
				}
				return &v.Bool, nil
			}

			return nil, invalidType("BOOLEAN")
		},
	},
	"DATE": {
		fieldType: data.FieldTypeNullableTime,
		convert: func(in interface{}) (interface{}, error) {
			if in == nil {
				return time.Time{}, nil
			}

			if v, ok := in.(*sqldriver.NullDate); ok {
				if v.IsNull() {
					return (*time.Time)(nil), nil
				}
				return &v.Time, nil
			}

			return nil, invalidType("DATE")
		},
	},
	"DATETIME": {
		fieldType: data.FieldTypeNullableTime,
		convert: func(in interface{}) (interface{}, error) {
			if in == nil {
				return time.Time{}, nil
			}

			if v, ok := in.(*sqldriver.NullDateTime); ok {
				if v.IsNull() {
					return (*time.Time)(nil), nil
				}
				return &v.Time, nil
			}

			return nil, invalidType("DATETIME")
		},
	},
	"TIMESTAMP": {
		fieldType: data.FieldTypeNullableTime,
		convert: func(in interface{}) (interface{}, error) {
			if in == nil {
				return time.Time{}, nil
			}

			if v, ok := in.(*sqldriver.NullTimeStamp); ok {
				if v.IsNull() {
					return (*time.Time)(nil), nil
				}
				return &v.Time, nil
			}

			return nil, invalidType("TIMESTAMP")
		},
	},
	"DECIMAL": {
		fieldType:  data.FieldTypeNullableString,
		scanType:   reflect.TypeOf(sqldriver.Decimal{}),
		matchRegex: matchRegexes["DECIMAL"],
		convert: func(in interface{}) (interface{}, error) {
			if in == nil {
				return "", nil
			}

			if v, ok := in.(*sqldriver.Decimal); ok {
				if v.IsNull() {
					return (*string)(nil), nil
				}
				return makePtrToString(v.String()), nil
			}

			return nil, invalidType("DECIMAL")
		},
	},
	"MAP": {
		fieldType:  data.FieldTypeNullableString,
		scanType:   reflect.TypeOf(sqldriver.Map{}),
		matchRegex: matchRegexes["MAP"],
		convert: func(in interface{}) (interface{}, error) {
			if in == nil {
				return "", nil
			}

			if v, ok := in.(*sqldriver.Map); ok {
				return makePtrToString(v.String()), nil
			}

			return nil, invalidType("MAP")
		},
	},
	"ARRAY": {
		fieldType:  data.FieldTypeNullableString,
		matchRegex: matchRegexes["ARRAY"],
		scanType:   reflect.TypeOf(sqldriver.Array{}),
		convert: func(in interface{}) (interface{}, error) {
			if in == nil {
				return "", nil
			}

			if v, ok := in.(*sqldriver.Array); ok {
				return makePtrToString(v.String()), nil
			}

			return nil, invalidType("ARRAY")
		},
	},
	"STRUCT": {
		fieldType:  data.FieldTypeNullableString,
		scanType:   reflect.TypeOf(sqldriver.Struct{}),
		matchRegex: matchRegexes["STRUCT"],
		convert: func(in interface{}) (interface{}, error) {
			if in == nil {
				return "", nil
			}

			if v, ok := in.(*sqldriver.Struct); ok {
				return makePtrToString(v.String()), nil
			}

			return nil, invalidType("STRUCT")
		},
	},
	"VOID": {
		fieldType: data.FieldTypeNullableBool,
		scanType:  reflect.TypeOf(data2.Null),
		convert: func(in interface{}) (interface{}, error) {
			return nil, nil
		},
	},
	"INTERVAL_DAY_TIME": {
		fieldType: data.FieldTypeNullableString,
		scanType:  reflect.TypeOf(data2.IntervalDayTime{}),
		convert: func(in interface{}) (interface{}, error) {
			if in == nil {
				return "", nil
			}

			if v, ok := in.(*data2.IntervalDayTime); ok {
				return makePtrToString(v.String()), nil
			}

			return nil, invalidType("INTERVAL_DAY_TIME")
		},
	},
	"INTERVAL_YEAR_MONTH": {
		fieldType: data.FieldTypeNullableString,
		scanType:  reflect.TypeOf(data2.IntervalYearMonth(0)),
		convert: func(in interface{}) (interface{}, error) {
			if in == nil {
				return "", nil
			}

			if v, ok := in.(*data2.IntervalDayTime); ok {
				return (makePtrToString)(v.String()), nil
			}

			return nil, invalidType("INTERVAL_YEAR_MONTH")
		},
	},
}

func GetConverter(cn string) sqlutil.Converter {
	converter, ok := Converters[cn]
	if ok {
		return createConverter(cn, converter)
	}
	for name, converter := range Converters {
		if name == cn {
			return createConverter(name, converter)
		}
		if converter.matchRegex != nil && converter.matchRegex.MatchString(cn) {
			return createConverter(name, converter)
		}
	}
	return sqlutil.Converter{}
}

var MaxComputeConverters = MaxcomputeConverters()

func MaxcomputeConverters() []sqlutil.Converter {
	list := make([]sqlutil.Converter, 0, len(Converters))
	for name, converter := range Converters {
		list = append(list, createConverter(name, converter))
	}
	return list
}

func createConverter(name string, converter Converter) sqlutil.Converter {
	return sqlutil.Converter{
		Name:           name,
		InputScanType:  converter.scanType,
		InputTypeRegex: converter.matchRegex,
		InputTypeName:  name,
		FrameConverter: sqlutil.FrameConverter{
			FieldType:     converter.fieldType,
			ConverterFunc: converter.convert,
		},
	}
}

func stringConverter(in interface{}) (interface{}, error) {
	if in == nil {
		return "", nil
	}

	if v, ok := in.(string); ok {
		return v, nil
	}

	if v, ok := in.(*sqldriver.NullString); ok {
		if v.IsNull() {
			return (*string)(nil), nil
		}
		return &v.String, nil
	}

	return nil, invalidType("STRING")
}

func jsonConverter(in interface{}) (interface{}, error) {
	if in == nil {
		return (*string)(nil), nil
	}
	jBytes, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}
	var rawJSON json.RawMessage
	err = json.Unmarshal(jBytes, &rawJSON)
	if err != nil {
		return nil, err
	}
	return &rawJSON, nil
}

func makePtrToString(str string) *string {
	return &str
}
