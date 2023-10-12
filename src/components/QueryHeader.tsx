import React from 'react';
// import { QueryTypeSwitcher } from 'components/QueryTypeSwitcher';
import { Button } from '@grafana/ui';
import { EditorHeader, FlexItem } from '@grafana/experimental';
import { Format, MCQuery, QueryType } from 'types';
import { FormatSelect } from './FormatSelect';

interface QueryHeaderProps {
  query: MCQuery;
  onChange: (query: MCQuery) => void;
  onRunQuery: () => void;
}

export const QueryHeader = ({ query, onChange, onRunQuery }: QueryHeaderProps) => {
  React.useEffect(() => {
    if (typeof query.selectedFormat === 'undefined' && query.queryType === QueryType.SQL) {
      const selectedFormat = Format.TABLE;
      const format = selectedFormat;
      onChange({ ...query, selectedFormat, format });
    }
  }, [query, onChange]);

  const onFormatChange = (selectedFormat: Format) => {
    switch (query.queryType) {
      case QueryType.SQL:
      default:
        onChange({ ...query, format: selectedFormat, selectedFormat });
    }
  };

  return (
    <EditorHeader>
      <FlexItem grow={1} />
      <Button variant="primary" icon="play" size="sm" onClick={onRunQuery}>
        Run query
      </Button>
      <FormatSelect format={query.selectedFormat ?? Format.AUTO} onChange={onFormatChange} />
    </EditorHeader>
  );
};
