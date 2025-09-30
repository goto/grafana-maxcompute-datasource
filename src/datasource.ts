import { DataSourceInstanceSettings, CoreApp, ScopedVars, VariableSupportType, DataQueryRequest } from '@grafana/data';
import { DataSourceWithBackend, getTemplateSrv } from '@grafana/runtime';

import { MCQuery, MCConfig, defaultMCSQLQuery } from './types';
import { SQLEditor } from './components/SQLEditor'
import { uniqueId } from 'lodash';

export class DataSource extends DataSourceWithBackend<MCQuery, MCConfig> {
  constructor(instanceSettings: DataSourceInstanceSettings<MCConfig>) {
    super(instanceSettings);
    this.variables = {
      getType: () => VariableSupportType.Custom,
      editor: SQLEditor as any,
      query: (request: DataQueryRequest<MCQuery>) => {
        const queries = request.targets.map((query) => {
          return { ...query, refId: query.refId || uniqueId('tempVar') };
        });
        return this.query({ ...request, targets: queries });
      }
    };
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
