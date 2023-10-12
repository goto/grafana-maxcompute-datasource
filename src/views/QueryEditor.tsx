import React from 'react';
import { QueryEditorProps } from '@grafana/data';
import { DataSource } from '../datasource';
import { Format, MCConfig, MCQuery, QueryType } from '../types';
import { SQLEditor } from 'components/SQLEditor';
import { QueryHeader } from 'components/QueryHeader';

type MCQueryEditorProps = QueryEditorProps<DataSource, MCQuery, MCConfig>;

const MCEditorByType = (props: MCQueryEditorProps) => {
  const { query } = props;

  switch (query.queryType) {
    case QueryType.SQL:
    default:
      return (
        <div data-testid="query-editor-section-sql">
          <SQLEditor {...props} />
        </div>
      );
  };
}

export function MCQueryEditor(props: MCQueryEditorProps) {
  const { query, onChange, onRunQuery } = props;

  React.useEffect(() => {
    if (typeof query.selectedFormat === 'undefined' && query.queryType === QueryType.SQL) {
      const selectedFormat = Format.TABLE;
      const format = Format.AUTO;
      onChange({ ...query, selectedFormat, format })
    }
  }, [query, onChange])

  return (
    <>
      <QueryHeader query={query} onChange={onChange} onRunQuery={onRunQuery} />
      <MCEditorByType {...props} />
    </>
  )
}
