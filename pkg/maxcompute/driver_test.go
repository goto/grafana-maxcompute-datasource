package maxcompute_test

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/ManassehZhou/maxcompute-datasource/pkg/converters"
	"github.com/ManassehZhou/maxcompute-datasource/pkg/maxcompute"
	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/data"
	"github.com/grafana/grafana-plugin-sdk-go/data/sqlutil"
	"github.com/stretchr/testify/require"
	"gotest.tools/assert"
)

func toJson(obj interface{}) (json.RawMessage, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return nil, errors.New("unable to marshal")
	}
	var rawJSON json.RawMessage
	err = json.Unmarshal(bytes, &rawJSON)
	if err != nil {
		return nil, errors.New("unable to unmarshal")
	}
	return rawJSON, nil
}

func TestConnect(t *testing.T) {
	config, _ := toJson(map[string]interface{}{
		"accessKeyId": os.Getenv("ALIBABACLOUD_ACCESS_KEY_ID"),
		"endpoint":    os.Getenv("MAXCOMPUTE_ENDPOINT"),
		"projectName": os.Getenv("MAXCOMPUTE_PROJECT"),
	})

	mc := maxcompute.MaxComputeDriver{}
	t.Run("should not error when valid settings passed", func(t *testing.T) {
		secure := map[string]string{}
		secure["accessKeySecret"] = os.Getenv("ALIBABACLOUD_ACCESS_KEY_SECRET")
		settings := backend.DataSourceInstanceSettings{JSONData: config, DecryptedSecureJSONData: secure}
		_, err := mc.Connect(context.TODO(), settings, json.RawMessage{})
		assert.Equal(t, nil, err)
	})

	config, _ = toJson(map[string]interface{}{
		"accessKeyId":          os.Getenv("ALIBABACLOUD_ACCESS_KEY_ID"),
		"endpoint":             os.Getenv("MAXCOMPUTE_ENDPOINT"),
		"projectName":          os.Getenv("MAXCOMPUTE_PROJECT"),
		"tcpConnectionTimeout": 20,
	})

	t.Run("should not error when valid settings passed - with tcp timeout as number", func(t *testing.T) {
		secure := map[string]string{}
		secure["accessKeySecret"] = os.Getenv("ALIBABACLOUD_ACCESS_KEY_SECRET")
		settings := backend.DataSourceInstanceSettings{JSONData: config, DecryptedSecureJSONData: secure}
		_, err := mc.Connect(context.TODO(), settings, json.RawMessage{})
		assert.Equal(t, nil, err)
	})
}

func setupConnection(t *testing.T) (*sql.DB, error) {
	mc := maxcompute.MaxComputeDriver{}

	config, _ := toJson(map[string]interface{}{
		"accessKeyId": os.Getenv("ALIBABACLOUD_ACCESS_KEY_ID"),
		"endpoint":    os.Getenv("MAXCOMPUTE_ENDPOINT"),
		"projectName": os.Getenv("MAXCOMPUTE_PROJECT"),
		"others": []maxcompute.CustomOption{
			{Key: "odps.sql.submit.mode", Value: "script"},
		},
	})

	secure := map[string]string{}
	secure["accessKeySecret"] = os.Getenv("ALIBABACLOUD_ACCESS_KEY_SECRET")
	settings := backend.DataSourceInstanceSettings{JSONData: config, DecryptedSecureJSONData: secure}
	return mc.Connect(context.TODO(), settings, json.RawMessage{})
}

func TestSql(t *testing.T) {
	t.Run("type & converter test", func(t *testing.T) {
		db, err := setupConnection(t)
		if err != nil {
			panic(err)
		}
		rows, err := db.Query(`
		set odps.sql.type.system.odps2=true;
		select 1Y as tinyint, 
			2S as smallint, 
			1000 as int, 
			3L as long, 
			unhex('FA34E10293CB42848573A4E39937F479') as binary,
			3.14F as float,
			3.14D as double,
			3.5BD as decimal,
			cast("abcd" as varchar(4)) as varchar4,
			cast("abcd" as char(4)) as char4,
			"abcd" as string,
			DATE'2017-11-11' as date,
			DATETIME'2017-11-11 00:00:01' as datetime,
			TIMESTAMP'2017-11-11 00:00:02.123456789' as timestamp,
			True as bool,
			array(struct(1, 2), struct(3, 4)) as array,
			map("k1", "v1","k2","v2") as map,
			named_struct('x', 1,'y',2) as struct
		;`)
		require.NoError(t, err)

		frame, err := sqlutil.FrameFromRows(rows, 1, converters.MaxComputeConverters...)

		require.NotNil(t, frame)
		require.NoError(t, err)
		assert.Equal(t, 18, len(frame.Fields))

		type test struct {
			id       int
			typeName data.FieldType
			value    interface{}
		}
		tests := []test{
			{id: 0, typeName: data.FieldTypeNullableInt8, value: ptrOf(int8(1))},
			{id: 1, typeName: data.FieldTypeNullableInt16, value: ptrOf(int16(2))},
			{id: 2, typeName: data.FieldTypeNullableInt32, value: ptrOf(int32(1000))},
			{id: 3, typeName: data.FieldTypeNullableInt64, value: ptrOf(int64(3))},
			{id: 4, typeName: data.FieldTypeNullableJSON, value: ptrOf(json.RawMessage(`"+jThApPLQoSFc6TjmTf0eQ=="`))},
			{id: 5, typeName: data.FieldTypeNullableFloat32, value: ptrOf(float32(3.14))},
			{id: 6, typeName: data.FieldTypeNullableFloat64, value: ptrOf(float64(3.14))},
			{id: 7, typeName: data.FieldTypeNullableString, value: ptrOf(string("3.5"))},
			{id: 8, typeName: data.FieldTypeNullableString, value: ptrOf(string("abcd"))},
			{id: 9, typeName: data.FieldTypeNullableString, value: ptrOf(string("abcd"))},
			{id: 10, typeName: data.FieldTypeNullableString, value: ptrOf(string("abcd"))},
			{id: 11, typeName: data.FieldTypeNullableTime, value: ptrOf(time.Date(2017, time.November, 11, 00, 00, 00, 00000000, time.UTC).UTC())},
			// {id: 12, typeName: data.FieldTypeNullableTime, value: ptrOf(time.Date(2017, time.November, 11, 00, 00, 01, 00000000, time.UTC).UTC())},
			// {id: 13, typeName: data.FieldTypeNullableTime, value: ptrOf(time.Date(2017, time.November, 11, 00, 00, 02, 123456789, time.UTC).UTC())},
			{id: 14, typeName: data.FieldTypeNullableBool, value: ptrOf(bool(true))},
			{id: 15, typeName: data.FieldTypeNullableString, value: ptrOf(string("array(struct<col1:1,col2:2>, struct<col1:3,col2:4>)"))},
			{id: 16, typeName: data.FieldTypeNullableString, value: ptrOf(string("map('k1', 'v1', 'k2', 'v2')"))},
			{id: 17, typeName: data.FieldTypeNullableString, value: ptrOf(string("struct<x:1,y:2>"))},
		}

		for i, tc := range tests {
			t.Run(fmt.Sprintf("[%v/%v]", i+1, len(tests)), func(t *testing.T) {
				assert.Equal(t, frame.Fields[tc.id].Type(), tc.typeName)
				assert.DeepEqual(t, tc.value, frame.Fields[tc.id].At(0))
			})
		}
	})
}

func ptrOf[K any](val K) *K {
	return &val
}
