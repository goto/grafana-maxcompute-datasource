import { DataQuery, DataSourceJsonData } from '@grafana/data';

export const defaultQuery: Partial<MCQuery> = {};

export enum QueryType {
  SQL = 'sql',
  BUILDER = 'builder',
}

export interface MCQueryBase extends DataQuery {}

export enum Format {
  TIMESERIES = 0,
  TABLE = 1,
  LOGS = 2,
  TRACE = 3,
  AUTO = 4,
}

export interface MCSQLQuery extends MCQueryBase {
  queryType: QueryType.SQL;
  rawSql: string;
  
  format: Format;
  selectedFormat: Format;
  expand?: boolean;
}

export interface MCBuilderQuery extends MCQueryBase {
  queryType: QueryType.BUILDER;
  rawSql: string;
  
  format: Format;
  selectedFormat: Format;
}

export type MCQuery = MCSQLQuery | MCBuilderQuery;

// TODO: add query builder support later...

/**
 * These are options configured for each DataSource instance
 */

export interface MCConfig extends DataSourceJsonData {
  endpoint: string;
  projectName: string;

  accessKeyId: string;

  tcpConnectionTimeout?: number;
  httpTimeout?: number;
  tunnelEndpoint?: string;
  tunnelQuotaName?: string;

  others?: CustomOption[];
}

export interface CustomOption {
  key: string;
  value: string;
}

/**
 * Value that is used in the backend, but never sent over HTTP to the frontend
 */
export interface MCSecureConfig {
  accessKeySecret: string;
  stsToken?: string;
}

export const defaultQueryType: QueryType = QueryType.SQL;
export const defaultMCSQLQuery: Omit<MCSQLQuery, 'refId'> = {
  queryType: QueryType.SQL,
  rawSql: '',
  format: Format.TABLE,
  selectedFormat: Format.TABLE,
}

export const defaultMCBuilderQuery: Omit<MCBuilderQuery, 'refId'> = {
  queryType: QueryType.BUILDER,
  rawSql: '',
  format: Format.TABLE,
  selectedFormat: Format.TABLE,
}
