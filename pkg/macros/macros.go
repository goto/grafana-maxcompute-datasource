package macros

import (
	"errors"
	"fmt"
	"time"

	"github.com/grafana/sqlds/v3"
)

var (
	ErrorNoArgumentsToMacro           = errors.New("expected minimum of 1 argument. But no argument found")
	ErrorInsufficientArgumentsToMacro = errors.New("expected number of arguments not matching")
)

func invalidArgs(args []string) error {
	return sqlds.DownstreamError(fmt.Errorf("%w: expected 1 argument, received %d", sqlds.ErrorBadArgumentCount, len(args)))
}

// Default time filter for SQL based on the query time range.
// It requires one argument, the time column to filter.
// Example:
//
//	$__timeFilter(time) => "time BETWEEN '2006-01-02 15:04:05' AND '2006-01-02 15:04:05'"
func MacroTimeFilter(query *sqlds.Query, args []string) (string, error) {
	if len(args) != 1 {
		return "", invalidArgs(args)
	}

	var (
		column = args[0]
		from   = query.TimeRange.From.UTC().Format(time.DateTime)
		to     = query.TimeRange.To.UTC().Format(time.DateTime)
	)

	return fmt.Sprintf("%s >= '%s' AND %s <= '%s'", column, from, column, to), nil
}

// Default time filter for SQL based on the starting query time range.
// It requires one argument, the time column to filter.
// Example:
//
//	$__timeFrom(time) => "time > '2006-01-02 15:04:05'"
func MacroTimeFrom(query *sqlds.Query, args []string) (string, error) {
	if len(args) != 1 {
		return "", invalidArgs(args)
	}

	return fmt.Sprintf("%s >= '%s'", args[0], query.TimeRange.From.UTC().Format(time.DateTime)), nil
}

// Default time group for SQL based the given period.
// This basic example is meant to be customized with more complex periods.
// It requires two arguments, the column to filter and the period.
// Example:
//
//	$__timeTo(time, month) => "datepart(year, time), datepart(month, time)'"
func MacroTimeGroup(_ *sqlds.Query, args []string) (string, error) {
	if len(args) != 2 {
		return "", invalidArgs(args)
	}

	column := args[0]

	res := ""
	switch args[1] {
	case "minute":
		res += fmt.Sprintf("datepart(%s, 'mi') as %s_minute,", column, column)
		fallthrough
	case "hour":
		res += fmt.Sprintf("datepart(%s, 'hh') as %s_hour,", column, column)
		fallthrough
	case "day":
		res += fmt.Sprintf("datepart(%s, 'dd') as %s_day,", column, column)
		fallthrough
	case "month":
		res += fmt.Sprintf("datepart(%s, 'mm') as %s_month,", column, column)
		fallthrough
	case "year":
		res += fmt.Sprintf("datepart(%s, 'yyyy') as %s_year", column, column)
	}

	return res, nil
}

// Default time filter for SQL based on the ending query time range.
// It requires one argument, the time column to filter.
// Example:
//
//	$__timeTo(time) => "time < '2006-01-02 15:04:05'"
func MacroTimeTo(query *sqlds.Query, args []string) (string, error) {
	if len(args) != 1 {
		return "", invalidArgs(args)
	}

	return fmt.Sprintf("%s <= '%s'", args[0], query.TimeRange.To.UTC().Format(time.DateTime)), nil
}
