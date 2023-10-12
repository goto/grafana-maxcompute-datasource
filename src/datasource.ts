import { DataSourceInstanceSettings, CoreApp, ScopedVars } from '@grafana/data';
import { DataSourceWithBackend, getTemplateSrv } from '@grafana/runtime';

import { MCQuery, MCConfig, defaultMCSQLQuery } from './types';

export class DataSource extends DataSourceWithBackend<MCQuery, MCConfig> {
  constructor(instanceSettings: DataSourceInstanceSettings<MCConfig>) {
    super(instanceSettings);
  }

  getDefaultQuery(_: CoreApp): Partial<MCQuery> {
    return defaultMCSQLQuery; 
  }

  filterQuery(query: MCQuery): boolean {
    return query.hide !== true && query.rawSql !== '';
  }

  applyTemplateVariables(query: MCQuery, scopedVars: ScopedVars): Record<string, any> {
    let rawQuery = query.rawSql || '';
    let templateSrv = getTemplateSrv();

    rawQuery = templateSrv.replace(rawQuery, scopedVars)

    return {
      ...query,
      rawSql: rawQuery
    }
    
  }
}
