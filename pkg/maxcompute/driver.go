package maxcompute

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/ManassehZhou/maxcompute-datasource/pkg/converters"
	"github.com/ManassehZhou/maxcompute-datasource/pkg/macros"
	_ "github.com/aliyun/aliyun-odps-go-sdk/sqldriver"
	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"github.com/grafana/grafana-plugin-sdk-go/data"
	"github.com/grafana/grafana-plugin-sdk-go/data/sqlutil"
	"github.com/grafana/sqlds/v3"
)

var (
	_ sqlds.Driver = (*MaxComputeDriver)(nil)
)

type MaxComputeDriver struct {
}

// Connect connects to the database. It does not need to call `db.Ping()`
func (*MaxComputeDriver) Connect(_ context.Context, settings backend.DataSourceInstanceSettings, raw json.RawMessage) (*sql.DB, error) {
	log.DefaultLogger.Debug("Creating MaxCompute instance")
	config, err := LoadMaxComputeConfig(settings)
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("odps", config.FormatDsn())
	return db, err
}

// Settings are read whenever the plugin is initialized, or after the data source settings are updated
func (*MaxComputeDriver) Settings(_ context.Context, settings backend.DataSourceInstanceSettings) (res sqlds.DriverSettings) {
	res.FillMode = &data.FillMissing{
		Mode: data.FillModeNull,
	}

	config, err := LoadMaxComputeConfig(settings)
	if err != nil {
		res.Timeout = time.Second * 30
	}

	res.Timeout = config.TcpConnectionTimeout
	return
}

func (*MaxComputeDriver) Macros() sqlds.Macros {
	return map[string]sqlds.MacroFunc{
		"timeFrom":   macros.MacroTimeFrom,
		"timeTo":     macros.MacroTimeTo,
		"timeFilter": macros.MacroTimeFilter,
		"timeGroup":  macros.MacroTimeGroup,
	}
}

func (*MaxComputeDriver) Converters() []sqlutil.Converter {
	return converters.MaxComputeConverters
}
