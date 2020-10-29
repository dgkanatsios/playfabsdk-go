package insights

import "time"
                    
// InsightsEmptyRequest 
type InsightsEmptyRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
}

// InsightsGetDetailsResponse 
type InsightsGetDetailsResponseModel struct {
    // DataUsageMb amount of data (in MB) currently used by Insights.
    DataUsageMb uint32 `json:"DataUsageMb"`
    // ErrorMessage details of any error that occurred while retrieving Insights details.
    ErrorMessage string `json:"ErrorMessage"`
    // Limits allowed range of values for performance level and data storage retention.
    Limits *InsightsGetLimitsResponseModel `json:"Limits,omitempty"`
    // PendingOperations list of pending Insights operations for the title.
    PendingOperations []InsightsGetOperationStatusResponseModel `json:"PendingOperations,omitempty"`
    // PerformanceLevel current Insights performance level setting.
    PerformanceLevel int32 `json:"PerformanceLevel"`
    // RetentionDays current Insights data storage retention value in days.
    RetentionDays int32 `json:"RetentionDays"`
}

// InsightsGetLimitsResponse 
type InsightsGetLimitsResponseModel struct {
    // DefaultPerformanceLevel default Insights performance level.
    DefaultPerformanceLevel int32 `json:"DefaultPerformanceLevel"`
    // DefaultStorageRetentionDays default Insights data storage retention days.
    DefaultStorageRetentionDays int32 `json:"DefaultStorageRetentionDays"`
    // StorageMaxRetentionDays maximum allowed data storage retention days.
    StorageMaxRetentionDays int32 `json:"StorageMaxRetentionDays"`
    // StorageMinRetentionDays minimum allowed data storage retention days.
    StorageMinRetentionDays int32 `json:"StorageMinRetentionDays"`
    // SubMeters list of Insights submeter limits for the allowed performance levels.
    SubMeters []InsightsPerformanceLevelModel `json:"SubMeters,omitempty"`
}

// InsightsGetOperationStatusRequest returns the current status for the requested operation id.
type InsightsGetOperationStatusRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // OperationId id of the Insights operation.
    OperationId string `json:"OperationId"`
}

// InsightsGetOperationStatusResponse 
type InsightsGetOperationStatusResponseModel struct {
    // Message optional message related to the operation details.
    Message string `json:"Message"`
    // OperationCompletedTime time the operation was completed.
    OperationCompletedTime time.Time `json:"OperationCompletedTime"`
    // OperationId id of the Insights operation.
    OperationId string `json:"OperationId"`
    // OperationLastUpdated time the operation status was last updated.
    OperationLastUpdated time.Time `json:"OperationLastUpdated"`
    // OperationStartedTime time the operation started.
    OperationStartedTime time.Time `json:"OperationStartedTime"`
    // OperationType the type of operation, SetPerformance or SetStorageRetention.
    OperationType string `json:"OperationType"`
    // OperationValue the value requested for the operation.
    OperationValue int32 `json:"OperationValue"`
    // Status current status of the operation.
    Status string `json:"Status"`
}

// InsightsGetPendingOperationsRequest returns a list of operations that are in the pending state for the requested operation type.
type InsightsGetPendingOperationsRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // OperationType the type of pending operations requested, or blank for all operation types.
    OperationType string `json:"OperationType"`
}

// InsightsGetPendingOperationsResponse 
type InsightsGetPendingOperationsResponseModel struct {
    // PendingOperations list of pending Insights operations.
    PendingOperations []InsightsGetOperationStatusResponseModel `json:"PendingOperations,omitempty"`
}

// InsightsOperationResponse 
type InsightsOperationResponseModel struct {
    // Message optional message related to the operation details.
    Message string `json:"Message"`
    // OperationId id of the Insights operation.
    OperationId string `json:"OperationId"`
    // OperationType the type of operation, SetPerformance or SetStorageRetention.
    OperationType string `json:"OperationType"`
}

// InsightsPerformanceLevel 
type InsightsPerformanceLevelModel struct {
    // ActiveEventExports number of allowed active event exports.
    ActiveEventExports int32 `json:"ActiveEventExports"`
    // CacheSizeMB maximum cache size.
    CacheSizeMB int32 `json:"CacheSizeMB"`
    // Concurrency maximum number of concurrent queries.
    Concurrency int32 `json:"Concurrency"`
    // CreditsPerMinute number of Insights credits consumed per minute.
    CreditsPerMinute float64 `json:"CreditsPerMinute"`
    // EventsPerSecond maximum events per second.
    EventsPerSecond int32 `json:"EventsPerSecond"`
    // Level performance level.
    Level int32 `json:"Level"`
    // MaxMemoryPerQueryMB maximum amount of memory allowed per query.
    MaxMemoryPerQueryMB int32 `json:"MaxMemoryPerQueryMB"`
    // VirtualCpuCores amount of compute power allocated for queries and operations.
    VirtualCpuCores int32 `json:"VirtualCpuCores"`
}

// InsightsSetPerformanceRequest sets the performance level to the requested value. Use the GetLimits method to get the allowed values.
type InsightsSetPerformanceRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // PerformanceLevel the Insights performance level to apply to the title.
    PerformanceLevel int32 `json:"PerformanceLevel"`
}

// InsightsSetStorageRetentionRequest sets the data storage retention to the requested value. Use the GetLimits method to get the range of allowed values.
type InsightsSetStorageRetentionRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // RetentionDays the Insights data storage retention value (in days) to apply to the title.
    RetentionDays int32 `json:"RetentionDays"`
}
