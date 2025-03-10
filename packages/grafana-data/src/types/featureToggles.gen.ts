// NOTE: This file was auto generated.  DO NOT EDIT DIRECTLY!
// To change feature flags, edit:
//  pkg/services/featuremgmt/registry.go
// Then run tests in:
//  pkg/services/featuremgmt/toggles_gen_test.go

/**
 * Describes available feature toggles in Grafana. These can be configured via
 * conf/custom.ini to enable features under development or not yet available in
 * stable version.
 *
 * Only enabled values will be returned in this interface
 *
 * @public
 */
export interface FeatureToggles {
  [name: string]: boolean | undefined; // support any string value

  trimDefaults?: boolean;
  disableEnvelopeEncryption?: boolean;
  serviceAccounts?: boolean;
  database_metrics?: boolean;
  dashboardPreviews?: boolean;
  dashboardPreviewsAdmin?: boolean;
  ['live-config']?: boolean;
  ['live-pipeline']?: boolean;
  ['live-service-web-worker']?: boolean;
  queryOverLive?: boolean;
  panelTitleSearch?: boolean;
  tempoServiceGraph?: boolean;
  tempoApmTable?: boolean;
  prometheus_azure_auth?: boolean;
  prometheusAzureOverrideAudience?: boolean;
  influxdbBackendMigration?: boolean;
  newNavigation?: boolean;
  showFeatureFlagsInUI?: boolean;
  publicDashboards?: boolean;
  lokiLive?: boolean;
  swaggerUi?: boolean;
  featureHighlights?: boolean;
  dashboardComments?: boolean;
  annotationComments?: boolean;
  migrationLocking?: boolean;
  storage?: boolean;
  export?: boolean;
  storageLocalUpload?: boolean;
  azureMonitorResourcePickerForMetrics?: boolean;
  explore2Dashboard?: boolean;
  tracing?: boolean;
  commandPalette?: boolean;
  savedItems?: boolean;
  cloudWatchDynamicLabels?: boolean;
  datasourceQueryMultiStatus?: boolean;
  azureMonitorExperimentalUI?: boolean;
  traceToMetrics?: boolean;
  prometheusStreamingJSONParser?: boolean;
  validateDashboardsOnSave?: boolean;
  autoMigrateGraphPanels?: boolean;
  prometheusWideSeries?: boolean;
  canvasPanelNesting?: boolean;
  cloudMonitoringExperimentalUI?: boolean;
  logRequestsInstrumentedAsUnknown?: boolean;
  dataConnectionsConsole?: boolean;
  internationalization?: boolean;
}
