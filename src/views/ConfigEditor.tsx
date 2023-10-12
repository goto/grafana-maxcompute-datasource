import React, { ChangeEvent, useMemo, useState } from 'react';
import { Button, Field, HorizontalGroup, Input, SecretInput } from '@grafana/ui';
import { DataSourcePluginOptionsEditorProps, onUpdateDatasourceJsonDataOption, onUpdateDatasourceSecureJsonDataOption } from '@grafana/data';
import { CustomOption, MCConfig, MCSecureConfig } from '../types';
import { ConfigSection, ConfigSubSection, DataSourceDescription } from '@grafana/experimental';
import { Divider } from 'components/Divider';
import { Components } from 'selectors';

interface Props extends DataSourcePluginOptionsEditorProps<MCConfig> { }

export function MCConfigEditor(props: Props) {
  const { options, onOptionsChange } = props;
  const { jsonData, secureJsonFields } = options;
  const secureJsonData = (options.secureJsonData || {}) as MCSecureConfig;

  const hasAdditionalSettings = useMemo(
    () =>
      !!(
        options.jsonData.tcpConnectionTimeout ||
        options.jsonData.httpTimeout ||
        options.jsonData.tunnelEndpoint ||
        options.jsonData.tunnelQuotaName ||
        (options.jsonData.others && options.jsonData.others.length !== 0)
      ),
    [options]
  );

  const [otherOptions, setOtherOptions] = useState(jsonData.others || []);

  const onResetAccessKeySecret = () => {
    onOptionsChange({
      ...options,
      secureJsonFields: {
        ...options.secureJsonFields,
        accessKeySecret: false,
      },
      secureJsonData: {
        ...options.secureJsonData,
        accessKeySecret: '',
      },
    });
  };

  const onResetSTSToken = () => {
    onOptionsChange({
      ...options,
      secureJsonFields: {
        ...options.secureJsonFields,
        stsToken: false,
      },
      secureJsonData: {
        ...options.secureJsonData,
        stsToken: '',
      },
    });
  };

  const onOtherOptionsChange = (otherOptions: CustomOption[]) => {
    onOptionsChange({
      ...options,
      jsonData: {
        ...options.jsonData,
        others: otherOptions.filter((s) => !!s.key && !!s.value),
      }
    })
  }


  const isValidUrl = /^(http|https):\/\/(\w+:{0,1}\w*@)?(\S+)(:[0-9]+)?(\/|\/([\w#!:.?+=&%@!\-\/]))?$/.test(
    jsonData.endpoint
  );

  return (
    <>
      <DataSourceDescription
        dataSourceName="MaxCompute"
        docsLink="https://grafana.com/grafana/plugins/manassehzhou-maxcompute-datasource/"
        hasRequiredFields
      />

      <Divider />

      <ConfigSection title="Server">
        <Field
          required
          label={Components.ConfigEditor.Endpoint.label}
          description={Components.ConfigEditor.Endpoint.tooltip}
          invalid={!jsonData.endpoint || !isValidUrl}
          error={'Endpoint is required'}
        >
          <Input
            name='endpoint'
            width={40}
            value={jsonData.endpoint || ''}
            onChange={onUpdateDatasourceJsonDataOption(props, 'endpoint')}
            label={Components.ConfigEditor.Endpoint.label}
            aria-label={Components.ConfigEditor.Endpoint.label}
            placeholder={Components.ConfigEditor.Endpoint.placeholder}
          />
        </Field>

        <Field
          required
          label={Components.ConfigEditor.ProjectName.label}
          description={Components.ConfigEditor.ProjectName.tooltip}
          invalid={!jsonData.projectName}
          error={'Project name is required'}
        >
          <Input
            name='projectName'
            width={40}
            value={jsonData.projectName || ''}
            onChange={onUpdateDatasourceJsonDataOption(props, 'projectName')}
            label={Components.ConfigEditor.ProjectName.label}
            aria-label={Components.ConfigEditor.ProjectName.label}
            placeholder={Components.ConfigEditor.ProjectName.placeholder}
          />
        </Field>
      </ConfigSection>

      <Divider />
      <ConfigSection title="Credentials">
        <Field
          required
          label={Components.ConfigEditor.AccessKeyId.label}
          description={Components.ConfigEditor.AccessKeyId.tooltip}
          invalid={!jsonData.accessKeyId}
          error={'Access Key ID is required'}
        >
          <Input
            name="accessKeyId"
            width={40}
            value={jsonData.accessKeyId || ''}
            onChange={onUpdateDatasourceJsonDataOption(props, 'accessKeyId')}
            label={Components.ConfigEditor.AccessKeyId.label}
            aria-label={Components.ConfigEditor.AccessKeyId.label}
            placeholder={Components.ConfigEditor.AccessKeyId.placeholder}
          />
        </Field>

        <Field
          required
          label={Components.ConfigEditor.AccessKeySecret.label}
          description={Components.ConfigEditor.AccessKeySecret.tooltip}
          invalid={!secureJsonFields.accessKeySecret}
          error={'Access Key Secret is required'}
        >
          <SecretInput
            name="accessKeySecret"
            width={40}
            value={secureJsonData.accessKeySecret || ''}
            label={Components.ConfigEditor.AccessKeySecret.label}
            aria-label={Components.ConfigEditor.AccessKeySecret.label}
            placeholder={Components.ConfigEditor.AccessKeySecret.placeholder}
            onReset={onResetAccessKeySecret}
            onChange={onUpdateDatasourceSecureJsonDataOption(props, 'accessKeySecret')}
            isConfigured={(secureJsonFields && secureJsonFields.accessKeySecret) as boolean}
          />
        </Field>

        <Field
          label={Components.ConfigEditor.STSToken.label}
          description={Components.ConfigEditor.STSToken.tooltip}
        >
          <SecretInput
            name="stsToken"
            width={40}
            value={secureJsonData.stsToken || ''}
            label={Components.ConfigEditor.STSToken.label}
            aria-label={Components.ConfigEditor.STSToken.label}
            placeholder={Components.ConfigEditor.STSToken.placeholder}
            onReset={onResetSTSToken}
            onChange={onUpdateDatasourceSecureJsonDataOption(props, 'stsToken')}
            isConfigured={(secureJsonFields && secureJsonFields.stsToken) as boolean}
          />
        </Field>
      </ConfigSection>

      <Divider />
      <ConfigSection
        title='Additional settings'
        isCollapsible
        isInitiallyOpen={hasAdditionalSettings}
      >
        <Field
          label={Components.ConfigEditor.TcpConnectionTimeout.label}
          description={Components.ConfigEditor.TcpConnectionTimeout.tooltip}
        >
          <Input
            name="tcpConnectionTimeout"
            width={40}
            value={jsonData.tcpConnectionTimeout || ''}
            onChange={onUpdateDatasourceJsonDataOption(props, 'tcpConnectionTimeout')}
            label={Components.ConfigEditor.TcpConnectionTimeout.label}
            aria-label={Components.ConfigEditor.TcpConnectionTimeout.label}
            placeholder={Components.ConfigEditor.TcpConnectionTimeout.placeholder}
            type='number'
          />
        </Field>

        <Field
          label={Components.ConfigEditor.HttpTimeout.label}
          description={Components.ConfigEditor.HttpTimeout.tooltip}
        >
          <Input
            name="httpTimeout"
            width={40}
            value={jsonData.tcpConnectionTimeout || ''}
            onChange={onUpdateDatasourceJsonDataOption(props, 'httpTimeout')}
            label={Components.ConfigEditor.HttpTimeout.label}
            aria-label={Components.ConfigEditor.HttpTimeout.label}
            placeholder={Components.ConfigEditor.HttpTimeout.placeholder}
            type='number'
          />
        </Field>

        <Field
          label={Components.ConfigEditor.TunnelEndpoint.label}
          description={Components.ConfigEditor.TunnelEndpoint.tooltip}
        >
          <Input
            name="tunnelEndpoint"
            width={40}
            value={jsonData.tunnelEndpoint || ''}
            onChange={onUpdateDatasourceJsonDataOption(props, 'tunnelEndpoint')}
            label={Components.ConfigEditor.TunnelEndpoint.label}
            aria-label={Components.ConfigEditor.TunnelEndpoint.label}
            placeholder={Components.ConfigEditor.TunnelEndpoint.placeholder}
          />
        </Field>

        <Field
          label={Components.ConfigEditor.TunnelQuotaName.label}
          description={Components.ConfigEditor.TunnelQuotaName.tooltip}
        >
          <Input
            name="tunnelQuotaName"
            width={40}
            value={jsonData.tunnelQuotaName || ''}
            onChange={onUpdateDatasourceJsonDataOption(props, 'tunnelQuotaName')}
            label={Components.ConfigEditor.TunnelQuotaName.label}
            aria-label={Components.ConfigEditor.TunnelQuotaName.label}
            placeholder={Components.ConfigEditor.TunnelQuotaName.placeholder}
          />
        </Field>

        <ConfigSubSection title="Hints and Other Options">
          {otherOptions.map(({ key, value }, i) => {
            return (
              <HorizontalGroup key={i}>
                <Field label={`Key`} aria-label={`Key`}>
                  <Input
                    value={key}
                    placeholder={'Key'}
                    onChange={(changeEvent: ChangeEvent<HTMLInputElement>) => {
                      let newSettings = otherOptions.concat();
                      newSettings[i] = { key: changeEvent.target.value, value };
                      setOtherOptions(newSettings);
                    }}
                    onBlur={() => {
                      onOtherOptionsChange(otherOptions);
                    }}
                  ></Input>
                </Field>
                <Field label={'Value'} aria-label={`Value`}>
                  <Input
                    value={value}
                    placeholder={'Value'}
                    onChange={(changeEvent: ChangeEvent<HTMLInputElement>) => {
                      let newSettings = otherOptions.concat();
                      newSettings[i] = { key, value: changeEvent.target.value };
                      setOtherOptions(newSettings);
                    }}
                    onBlur={() => {
                      onOtherOptionsChange(otherOptions);
                    }}
                  ></Input>
                </Field>
              </HorizontalGroup>
            );
          })}

          <Button
            variant="secondary"
            icon="plus"
            type="button"
            onClick={() => {
              setOtherOptions([...otherOptions, { key: '', value: '' }])
            }}
          >
            Add custom setting
          </Button>
        </ConfigSubSection>
      </ConfigSection>
    </>
  );
}
