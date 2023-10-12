package main

import (
	"context"
	"os"

	"github.com/ManassehZhou/maxcompute-datasource/pkg/maxcompute"
	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/datasource"
	"github.com/grafana/grafana-plugin-sdk-go/backend/instancemgmt"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"github.com/grafana/sqlds/v3"
)

func newDatasource(ctx context.Context, settings backend.DataSourceInstanceSettings) (instancemgmt.Instance, error) {
	ds := sqlds.NewDatasource(&maxcompute.MaxComputeDriver{})
	return ds.NewDatasource(ctx, settings)
}

func main() {
	if err := datasource.Manage("manassehzhou-maxcompute-datasource", newDatasource, datasource.ManageOpts{}); err != nil {
		log.DefaultLogger.Error(err.Error())
		os.Exit(1)
	}
}
