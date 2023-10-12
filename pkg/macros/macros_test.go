package macros_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/ManassehZhou/maxcompute-datasource/pkg/maxcompute"
	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/sqlds/v3"
	"github.com/stretchr/testify/require"
	"gotest.tools/assert"
)

type MaxComputeDriver struct {
	sqlds.Driver
}

type MockDB struct {
	maxcompute.MaxComputeDriver
}

func (h *MaxComputeDriver) Macros() sqlds.Macros {
	var C = maxcompute.MaxComputeDriver{}

	return C.Macros()
}

func TestInterpolate(t *testing.T) {
	from, _ := time.Parse("2006-01-02T15:04:05.000Z", "2014-11-12T11:45:26.371Z")
	to, _ := time.Parse("2006-01-02T15:04:05.000Z", "2015-11-12T11:45:26.371Z")
	type test struct {
		name   string
		input  string
		output string
	}
	tests := []test{
		{input: "select * from foo where $__timeFilter(time)", output: "select * from foo where time >= '2014-11-12 11:45:26' AND time <= '2015-11-12 11:45:26'", name: "test timeFilter"},
		{input: "select * from foo where $__timeFilter(cast(sth as timestamp))", output: "select * from foo where cast(sth as timestamp) >= '2014-11-12 11:45:26' AND cast(sth as timestamp) <= '2015-11-12 11:45:26'", name: "test timeFilter"},
		{input: "select * from foo where $__timeFilter(cast(sth as timestamp) )", output: "select * from foo where cast(sth as timestamp) >= '2014-11-12 11:45:26' AND cast(sth as timestamp) <= '2015-11-12 11:45:26'", name: "test timeFilter with empty spaces"},
		{input: "select * from foo where $__timeTo(time)", output: "select * from foo where time <= '2015-11-12 11:45:26'", name: "test timeTo macro"},
		{input: "select * from foo where $__timeFrom(time)", output: "select * from foo where time >= '2014-11-12 11:45:26'", name: "test timeFrom macro"},
		{input: "select * from foo where $__timeFrom(cast(sth as timestamp))", output: "select * from foo where cast(sth as timestamp) >= '2014-11-12 11:45:26'", name: "test timeFrom macro"},
		{input: "select $__timeGroup(time,minute), * from foo", output: "select datepart(time, 'mi') as time_minute,datepart(time, 'hh') as time_hour,datepart(time, 'dd') as time_day,datepart(time, 'mm') as time_month,datepart(time, 'yyyy') as time_year, * from foo", name: "test timeGroup macro"},
	}
	for i, tc := range tests {
		driver := MockDB{}
		t.Run(fmt.Sprintf("[%d/%d] %s", i+1, len(tests), tc.name), func(t *testing.T) {
			query := &sqlds.Query{
				RawSQL: tc.input,
				TimeRange: backend.TimeRange{
					From: from,
					To:   to,
				},
			}
			interpolatedQuery, err := sqlds.Interpolate(&driver, query)
			require.Nil(t, err)
			assert.Equal(t, tc.output, interpolatedQuery)
		})
	}
}
