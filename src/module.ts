import { DataSourcePlugin } from '@grafana/data';
import { DataSource } from './datasource';
import { MCConfigEditor } from './views/ConfigEditor';
import { MCQueryEditor } from './views/QueryEditor';
import { MCQuery, MCConfig } from './types';

export const plugin = new DataSourcePlugin<DataSource, MCQuery, MCConfig>(DataSource)
  .setConfigEditor(MCConfigEditor)
  .setQueryEditor(MCQueryEditor);
