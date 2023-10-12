import { E2ESelectors } from '@grafana/e2e-selectors';

export const Components = {
    ConfigEditor: {
        Endpoint: {
            label: 'Endpoint',
            placeholder: 'http://maxcompute/api',
            tooltip: 'MaxCompute service endpoint URL',
        },
        ProjectName: {
            label: 'Project',
            placeholder: 'project_name',
            tooltip: 'MaxCompute project name',
        },
        AccessKeyId: {
            label: 'Access Key ID',
            placeholder: 'LTAI*******',
            tooltip: 'Alibaba Cloud Access Key Id',
        },
        AccessKeySecret: {
            label: 'Access Key Secret',
            placeholder: '******',
            tooltip: 'Alibaba Cloud Access Key Secret',
        },
        STSToken: {
            label: 'STS Token',
            placeholder: '',
            tooltip: 'Alibaba Cloud STS Token(required when auth with RAM role)',
        },
        TcpConnectionTimeout: {
            label: 'Tcp Connection Timeout',
            placeholder: '30',
            tooltip: 'Timeout in second, 0 for inf',
        },
        HttpTimeout: {
            label: 'HttpTimeout',
            placeholder: '0',
            tooltip: 'Timeout in second, 0 for inf',
        },
        TunnelEndpoint: {
            label: 'Tunnel Endpoint',
            placeholder: '',
            tooltip: 'MaxCompute Tunnel Endpoint',
        },
        TunnelQuotaName: {
            label: 'Tunnel Quota Name',
            placeholder: '',
            tooltip: 'MaxCompute Tunnel Quota Name',
        },
        Others: {},
    },
    QueryEditor: {
        CodeEditor: {
            input: () => '.monaco-editor textarea',
            container: 'data-testid-code-editor-container',
            Expand: 'data-testid-code-editor-expand-button',
        },
        Format: {
            label: 'Format',
            tooltip: 'Query Type',
            options: {
                AUTO: 'Auto',
                TABLE: 'Table',
                TIME_SERIES: 'Time Series',
                LOGS: 'Logs',
                TRACE: 'Trace',
            },
        },
        Types: {
            label: 'Query Type',
            tooltip: 'Query Type',
            options: {
                SQLEditor: 'SQL Editor',
                QueryBuilder: 'Query Builder',
            },
            switcher: {
                title: 'Are you sure?',
                body: 'Queries that are too complex for the Query Builder will be altered.',
                confirmText: 'Continue',
                dismissText: 'Cancel',
            },
            cannotConvert: {
                title: 'Cannot convert',
                confirmText: 'Yes',
            },
        },
    },
};

export const selectors: { components: E2ESelectors<typeof Components> } = {
    components: Components,
};
