# MaxCompute Datasource

This is a plugin for Grafana that enables queries to Aliyun MaxCompute.

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
![stability-stable](https://img.shields.io/badge/stability-stable-green.svg)
![GitHub Workflow Status (with event)](https://img.shields.io/github/actions/workflow/status/ManassehZhou/grafana-maxcompute-datasource/ci.yml)
![GitHub Workflow Status (with event)](https://img.shields.io/github/actions/workflow/status/ManassehZhou/grafana-maxcompute-datasource/release.yml)

## Requirements

The plugin requires the user to run Grafana >=10.0.0.

## Getting Started

### Plugin Installation

#### [Recommanded] Installing the Official and Released Plugin on an Existing Grafana With the CLI

The recommended way to install the plugin for most users is to use the grafana CLI:

1. Run this command: `grafana-cli plugins install manassehzhou-maxcompute-datasource`
2. Restart the Grafana server.
3. To make sure the plugin was installed, check the list of installed data sources. Click the
   Plugins item in the main menu. Both core data sources and installed data sources will appear.

#### Latest Version: Installing the newest Plugin Version on an Existing Grafana With the CLI

The grafana-cli can also install plugins from a non-standard URL. This way even plugin versions,
that are not (yet) released to the official Grafana repository can be installed.

1. Run this command:

   ```sh
   # replace the $VERSION part in the URL below with the desired version (e.g. 1.0.0)
   grafana-cli --pluginUrl https://github.com/ManassehZhou/grafana-maxcompute-datasource/releases/download/v$VERSION/manassehzhou-maxcompute-datasource-$VERSION.zip plugins install manassehzhou-maxcompute-datasource
   ```

2. See the recommended installation above (from the restart step)

#### Manual: Installing the Plugin Manually on an Existing Grafana

In case the grafana-cli does not work for whatever reason plugins can also be installed manually.

1. Get the zip file from [Latest release on Github](https://github.com/ManassehZhou/grafana-maxcompute-datasource/releases/latest)
2. Extract the zip file into the data/plugins subdirectory for Grafana:
   `unzip <the_download_zip_file> -d <plugin_dir>/`

   Finding the plugin directory can sometimes be a challenge as this is platform and settings
   dependent. A common location for this on Linux devices is `/var/lib/grafana/plugins/`
3. See the recommended installation above (from the restart step)

## Documentation

### Usage

#### Adding a MaxCompute Datasource

1. Open the side menu by clicking the Grafana icon in the top header.
1. In the side menu under the Dashboards link you should find a link named Data Sources.
1. Click the + Add data source button in the top header.
1. Select MaxCompute from the Type dropdown.

![Configure datasource](https://raw.githubusercontent.com/ManassehZhou/grafana-maxcompute-datasource/main/src/img/config-datasource.png)

### Future Document and Links

- A changelog of the plugin can be found in the [CHANGELOG.md](https://github.com/ManassehZhou/grafana-maxcompute-datasource/blob/main/CHANGELOG.md).
- The plugin in the Grafana registry can be found [here](https://grafana.com/grafana/plugins/manassehzhou-maxcompute-datasource/).
- The official SDK provided by aliyun can be found in [Github](https://github.com/aliyun/aliyun-odps-go-sdk).
- The official document about MaxCompute can be found in [Aliyun](https://www.alibabacloud.com/help/en/maxcompute/).

## Contributing

If you want to contribute to this plugin, you can find more at [Github](https://github.com/ManassehZhou/grafana-maxcompute-datasource/blob/main/DEVELOPMENT.md).
