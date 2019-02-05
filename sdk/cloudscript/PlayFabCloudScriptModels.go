package cloudscript
 
// CloudScriptRevisionOption 
type CloudScriptRevisionOption string
  
const (
     CloudScriptRevisionOptionLive CloudScriptRevisionOption = "Live"
     CloudScriptRevisionOptionLatest CloudScriptRevisionOption = "Latest"
     CloudScriptRevisionOptionSpecific CloudScriptRevisionOption = "Specific"
      )
// EntityKey combined entity type and ID structure which uniquely identifies a single entity.
type EntityKeyModel struct {
    // Id unique ID of the entity.
    Id string `json:"Id,omitempty"`
    // Type entity type. See https://api.playfab.com/docs/tutorials/entities/entitytypes
    Type string `json:"Type,omitempty"`
}

// ExecuteCloudScriptResult 
type ExecuteCloudScriptResultModel struct {
    // APIRequestsIssued number of PlayFab API requests issued by the CloudScript function
    APIRequestsIssued int32 `json:"APIRequestsIssued,omitempty"`
    // Error information about the error, if any, that occurred during execution
    Error ScriptExecutionErrorModel `json:"Error,omitempty"`
    // ExecutionTimeSeconds 
    ExecutionTimeSeconds float64 `json:"ExecutionTimeSeconds,omitempty"`
    // FunctionName the name of the function that executed
    FunctionName string `json:"FunctionName,omitempty"`
    // FunctionResult the object returned from the CloudScript function, if any
    FunctionResult interface{} `json:"FunctionResult,omitempty"`
    // FunctionResultTooLarge flag indicating if the FunctionResult was too large and was subsequently dropped from this event. This only occurs if
// the total event size is larger than 350KB.
    FunctionResultTooLarge bool `json:"FunctionResultTooLarge,omitempty"`
    // HttpRequestsIssued number of external HTTP requests issued by the CloudScript function
    HttpRequestsIssued int32 `json:"HttpRequestsIssued,omitempty"`
    // Logs entries logged during the function execution. These include both entries logged in the function code using log.info()
// and log.error() and error entries for API and HTTP request failures.
    Logs []LogStatementModel `json:"Logs,omitempty"`
    // LogsTooLarge flag indicating if the logs were too large and were subsequently dropped from this event. This only occurs if the total
// event size is larger than 350KB after the FunctionResult was removed.
    LogsTooLarge bool `json:"LogsTooLarge,omitempty"`
    // MemoryConsumedBytes 
    MemoryConsumedBytes uint32 `json:"MemoryConsumedBytes,omitempty"`
    // ProcessorTimeSeconds processor time consumed while executing the function. This does not include time spent waiting on API calls or HTTP
// requests.
    ProcessorTimeSeconds float64 `json:"ProcessorTimeSeconds,omitempty"`
    // Revision the revision of the CloudScript that executed
    Revision int32 `json:"Revision,omitempty"`
}

// ExecuteEntityCloudScriptRequest executes CloudScript with the entity profile that is defined in the request.
type ExecuteEntityCloudScriptRequestModel struct {
    // Entity the entity to perform this action on.
    Entity EntityKeyModel `json:"Entity,omitempty"`
    // FunctionName the name of the CloudScript function to execute
    FunctionName string `json:"FunctionName,omitempty"`
    // FunctionParameter object that is passed in to the function as the first argument
    FunctionParameter interface{} `json:"FunctionParameter,omitempty"`
    // GeneratePlayStreamEvent generate a 'entity_executed_cloudscript' PlayStream event containing the results of the function execution and other
// contextual information. This event will show up in the PlayStream debugger console for the player in Game Manager.
    GeneratePlayStreamEvent bool `json:"GeneratePlayStreamEvent,omitempty"`
    // RevisionSelection option for which revision of the CloudScript to execute. 'Latest' executes the most recently created revision, 'Live'
// executes the current live, published revision, and 'Specific' executes the specified revision. The default value is
// 'Specific', if the SpecificRevision parameter is specified, otherwise it is 'Live'.
    RevisionSelection CloudScriptRevisionOption `json:"RevisionSelection,omitempty"`
    // SpecificRevision the specific revision to execute, when RevisionSelection is set to 'Specific'
    SpecificRevision int32 `json:"SpecificRevision,omitempty"`
}

// LogStatement 
type LogStatementModel struct {
    // Data optional object accompanying the message as contextual information
    Data interface{} `json:"Data,omitempty"`
    // Level 'Debug', 'Info', or 'Error'
    Level string `json:"Level,omitempty"`
    // Message 
    Message string `json:"Message,omitempty"`
}

// ScriptExecutionError 
type ScriptExecutionErrorModel struct {
    // Error error code, such as CloudScriptNotFound, JavascriptException, CloudScriptFunctionArgumentSizeExceeded,
// CloudScriptAPIRequestCountExceeded, CloudScriptAPIRequestError, or CloudScriptHTTPRequestError
    Error string `json:"Error,omitempty"`
    // Message details about the error
    Message string `json:"Message,omitempty"`
    // StackTrace point during the execution of the script at which the error occurred, if any
    StackTrace string `json:"StackTrace,omitempty"`
}
