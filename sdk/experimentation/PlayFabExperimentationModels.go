package experimentation

import "time"
                    
// AnalysisTaskState 
type AnalysisTaskState string
  
const (
     AnalysisTaskStateWaiting AnalysisTaskState = "Waiting"
     AnalysisTaskStateReadyForSubmission AnalysisTaskState = "ReadyForSubmission"
     AnalysisTaskStateSubmittingToPipeline AnalysisTaskState = "SubmittingToPipeline"
     AnalysisTaskStateRunning AnalysisTaskState = "Running"
     AnalysisTaskStateCompleted AnalysisTaskState = "Completed"
     AnalysisTaskStateFailed AnalysisTaskState = "Failed"
     AnalysisTaskStateCanceled AnalysisTaskState = "Canceled"
      )
// CreateExclusionGroupRequest given a title entity token and exclusion group details, will create a new exclusion group for the title.
type CreateExclusionGroupRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Description description of the exclusion group.
    Description string `json:"Description,omitempty"`
    // Name friendly name of the exclusion group.
    Name string `json:"Name,omitempty"`
}

// CreateExclusionGroupResult 
type CreateExclusionGroupResultModel struct {
    // ExclusionGroupId identifier of the exclusion group.
    ExclusionGroupId string `json:"ExclusionGroupId,omitempty"`
}

// CreateExperimentRequest given a title entity token and experiment details, will create a new experiment for the title.
type CreateExperimentRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Description description of the experiment.
    Description string `json:"Description,omitempty"`
    // Duration the duration of the experiment, in days.
    Duration uint32 `json:"Duration,omitempty"`
    // EndDate when experiment should end.
    EndDate time.Time `json:"EndDate,omitempty"`
    // ExclusionGroupId id of the exclusion group.
    ExclusionGroupId string `json:"ExclusionGroupId,omitempty"`
    // ExclusionGroupTrafficAllocation percentage of exclusion group traffic that will see this experiment.
    ExclusionGroupTrafficAllocation uint32 `json:"ExclusionGroupTrafficAllocation,omitempty"`
    // ExperimentType type of experiment.
    ExperimentType ExperimentType `json:"ExperimentType,omitempty"`
    // Name friendly name of the experiment.
    Name string `json:"Name,omitempty"`
    // SegmentId id of the segment to which this experiment applies. Defaults to the 'All Players' segment.
    SegmentId string `json:"SegmentId,omitempty"`
    // StartDate when experiment should start.
    StartDate time.Time `json:"StartDate,omitempty"`
    // TitlePlayerAccountTestIds list of title player account IDs that automatically receive treatments in the experiment, but are not included when
// calculating experiment metrics.
    TitlePlayerAccountTestIds []string `json:"TitlePlayerAccountTestIds,omitempty"`
    // Variants list of variants for the experiment.
    Variants []VariantModel `json:"Variants,omitempty"`
}

// CreateExperimentResult 
type CreateExperimentResultModel struct {
    // ExperimentId the ID of the new experiment.
    ExperimentId string `json:"ExperimentId,omitempty"`
}

// DeleteExclusionGroupRequest given an entity token and an exclusion group ID this API deletes the exclusion group.
type DeleteExclusionGroupRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // ExclusionGroupId the ID of the exclusion group to delete.
    ExclusionGroupId string `json:"ExclusionGroupId,omitempty"`
}

// DeleteExperimentRequest given an entity token and an experiment ID this API deletes the experiment. A running experiment must be stopped before
// it can be deleted.
type DeleteExperimentRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // ExperimentId the ID of the experiment to delete.
    ExperimentId string `json:"ExperimentId,omitempty"`
}

// EmptyResponse 
type EmptyResponseModel struct {
}

// EntityKey combined entity type and ID structure which uniquely identifies a single entity.
type EntityKeyModel struct {
    // Id unique ID of the entity.
    Id string `json:"Id,omitempty"`
    // Type entity type. See https://docs.microsoft.com/gaming/playfab/features/data/entities/available-built-in-entity-types
    Type string `json:"Type,omitempty"`
}

// ExclusionGroupTrafficAllocation 
type ExclusionGroupTrafficAllocationModel struct {
    // ExperimentId id of the experiment.
    ExperimentId string `json:"ExperimentId,omitempty"`
    // TrafficAllocation percentage of exclusion group traffic that will see this experiment.
    TrafficAllocation uint32 `json:"TrafficAllocation,omitempty"`
}

// Experiment 
type ExperimentModel struct {
    // Description description of the experiment.
    Description string `json:"Description,omitempty"`
    // Duration the duration of the experiment, in days.
    Duration uint32 `json:"Duration,omitempty"`
    // EndDate when experiment should end/was ended.
    EndDate time.Time `json:"EndDate,omitempty"`
    // ExclusionGroupId id of the exclusion group for this experiment.
    ExclusionGroupId string `json:"ExclusionGroupId,omitempty"`
    // ExclusionGroupTrafficAllocation percentage of exclusion group traffic that will see this experiment.
    ExclusionGroupTrafficAllocation uint32 `json:"ExclusionGroupTrafficAllocation,omitempty"`
    // ExperimentType type of experiment.
    ExperimentType ExperimentType `json:"ExperimentType,omitempty"`
    // Id id of the experiment.
    Id string `json:"Id,omitempty"`
    // Name friendly name of the experiment.
    Name string `json:"Name,omitempty"`
    // SegmentId id of the segment to which this experiment applies. Defaults to the 'All Players' segment.
    SegmentId string `json:"SegmentId,omitempty"`
    // StartDate when experiment should start/was started.
    StartDate time.Time `json:"StartDate,omitempty"`
    // State state experiment is currently in.
    State ExperimentState `json:"State,omitempty"`
    // TitlePlayerAccountTestIds list of title player account IDs that automatically receive treatments in the experiment, but are not included when
// calculating experiment metrics.
    TitlePlayerAccountTestIds []string `json:"TitlePlayerAccountTestIds,omitempty"`
    // Variants list of variants for the experiment.
    Variants []VariantModel `json:"Variants,omitempty"`
}

// ExperimentExclusionGroup 
type ExperimentExclusionGroupModel struct {
    // Description description of the exclusion group.
    Description string `json:"Description,omitempty"`
    // ExclusionGroupId id of the exclusion group.
    ExclusionGroupId string `json:"ExclusionGroupId,omitempty"`
    // Name friendly name of the exclusion group.
    Name string `json:"Name,omitempty"`
}

// ExperimentState 
type ExperimentState string
  
const (
     ExperimentStateNew ExperimentState = "New"
     ExperimentStateStarted ExperimentState = "Started"
     ExperimentStateStopped ExperimentState = "Stopped"
     ExperimentStateDeleted ExperimentState = "Deleted"
      )
// ExperimentType 
type ExperimentType string
  
const (
     ExperimentTypeActive ExperimentType = "Active"
     ExperimentTypeSnapshot ExperimentType = "Snapshot"
      )
// GetExclusionGroupsRequest given a title entity token will return the list of all exclusion groups for a title.
type GetExclusionGroupsRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
}

// GetExclusionGroupsResult 
type GetExclusionGroupsResultModel struct {
    // ExclusionGroups list of exclusion groups for the title.
    ExclusionGroups []ExperimentExclusionGroupModel `json:"ExclusionGroups,omitempty"`
}

// GetExclusionGroupTrafficRequest given a title entity token and an exclusion group ID, will return the list of traffic allocations for the exclusion
// group.
type GetExclusionGroupTrafficRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // ExclusionGroupId the ID of the exclusion group.
    ExclusionGroupId string `json:"ExclusionGroupId,omitempty"`
}

// GetExclusionGroupTrafficResult 
type GetExclusionGroupTrafficResultModel struct {
    // TrafficAllocations list of traffic allocations for the exclusion group.
    TrafficAllocations []ExclusionGroupTrafficAllocationModel `json:"TrafficAllocations,omitempty"`
}

// GetExperimentsRequest given a title entity token will return the list of all experiments for a title, including scheduled, started, stopped or
// completed experiments.
type GetExperimentsRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
}

// GetExperimentsResult 
type GetExperimentsResultModel struct {
    // Experiments list of experiments for the title.
    Experiments []ExperimentModel `json:"Experiments,omitempty"`
}

// GetLatestScorecardRequest given a title entity token and experiment details, will return the latest available scorecard.
type GetLatestScorecardRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // ExperimentId the ID of the experiment.
    ExperimentId string `json:"ExperimentId,omitempty"`
}

// GetLatestScorecardResult 
type GetLatestScorecardResultModel struct {
    // Scorecard scorecard for the experiment of the title.
    Scorecard *ScorecardModel `json:"Scorecard,omitempty"`
}

// GetTreatmentAssignmentRequest given a title player or a title entity token, returns the treatment variants and variables assigned to the entity across
// all running experiments
type GetTreatmentAssignmentRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
}

// GetTreatmentAssignmentResult 
type GetTreatmentAssignmentResultModel struct {
    // TreatmentAssignment treatment assignment for the entity.
    TreatmentAssignment *TreatmentAssignmentModel `json:"TreatmentAssignment,omitempty"`
}

// MetricData 
type MetricDataModel struct {
    // ConfidenceIntervalEnd the upper bound of the confidence interval for the relative delta (Delta.RelativeValue).
    ConfidenceIntervalEnd float64 `json:"ConfidenceIntervalEnd,omitempty"`
    // ConfidenceIntervalStart the lower bound of the confidence interval for the relative delta (Delta.RelativeValue).
    ConfidenceIntervalStart float64 `json:"ConfidenceIntervalStart,omitempty"`
    // DeltaAbsoluteChange the absolute delta between TreatmentStats.Average and ControlStats.Average.
    DeltaAbsoluteChange float32 `json:"DeltaAbsoluteChange,omitempty"`
    // DeltaRelativeChange the relative delta ratio between TreatmentStats.Average and ControlStats.Average.
    DeltaRelativeChange float32 `json:"DeltaRelativeChange,omitempty"`
    // InternalName the machine name of the metric.
    InternalName string `json:"InternalName,omitempty"`
    // Movement indicates if a movement was detected on that metric.
    Movement string `json:"Movement,omitempty"`
    // Name the readable name of the metric.
    Name string `json:"Name,omitempty"`
    // PMove the expectation that a movement is real
    PMove float32 `json:"PMove,omitempty"`
    // PValue the p-value resulting from the statistical test run for this metric
    PValue float32 `json:"PValue,omitempty"`
    // PValueThreshold the threshold for observing sample ratio mismatch.
    PValueThreshold float32 `json:"PValueThreshold,omitempty"`
    // StatSigLevel indicates if the movement is statistically significant.
    StatSigLevel string `json:"StatSigLevel,omitempty"`
    // StdDev observed standard deviation value of the metric.
    StdDev float32 `json:"StdDev,omitempty"`
    // Value observed average value of the metric.
    Value float32 `json:"Value,omitempty"`
}

// Scorecard 
type ScorecardModel struct {
    // DateGenerated represents the date the scorecard was generated.
    DateGenerated string `json:"DateGenerated,omitempty"`
    // Duration represents the duration of scorecard analysis.
    Duration string `json:"Duration,omitempty"`
    // EventsProcessed represents the number of events processed for the generation of this scorecard
    EventsProcessed float64 `json:"EventsProcessed,omitempty"`
    // ExperimentId id of the experiment.
    ExperimentId string `json:"ExperimentId,omitempty"`
    // ExperimentName friendly name of the experiment.
    ExperimentName string `json:"ExperimentName,omitempty"`
    // LatestJobStatus represents the latest compute job status.
    LatestJobStatus AnalysisTaskState `json:"LatestJobStatus,omitempty"`
    // SampleRatioMismatch represents the presence of a sample ratio mismatch in the scorecard data.
    SampleRatioMismatch bool `json:"SampleRatioMismatch"`
    // ScorecardDataRows scorecard containing list of analysis.
    ScorecardDataRows []ScorecardDataRowModel `json:"ScorecardDataRows,omitempty"`
}

// ScorecardDataRow 
type ScorecardDataRowModel struct {
    // IsControl represents whether the variant is control or not.
    IsControl bool `json:"IsControl"`
    // MetricDataRows data of the analysis with the internal name of the metric as the key and an object of metric data as value.
    MetricDataRows map[string]MetricDataModel `json:"MetricDataRows,omitempty"`
    // PlayerCount represents the player count in the variant.
    PlayerCount uint32 `json:"PlayerCount,omitempty"`
    // VariantName name of the variant of analysis.
    VariantName string `json:"VariantName,omitempty"`
}

// StartExperimentRequest given a title entity token and an experiment ID, this API starts the experiment.
type StartExperimentRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // ExperimentId the ID of the experiment to start.
    ExperimentId string `json:"ExperimentId,omitempty"`
}

// StopExperimentRequest given a title entity token and an experiment ID, this API stops the experiment if it is running.
type StopExperimentRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // ExperimentId the ID of the experiment to stop.
    ExperimentId string `json:"ExperimentId,omitempty"`
}

// TreatmentAssignment 
type TreatmentAssignmentModel struct {
    // Variables list of the experiment variables.
    Variables []VariableModel `json:"Variables,omitempty"`
    // Variants list of the experiment variants.
    Variants []string `json:"Variants,omitempty"`
}

// UpdateExclusionGroupRequest given an entity token and exclusion group details this API updates the exclusion group.
type UpdateExclusionGroupRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Description description of the exclusion group.
    Description string `json:"Description,omitempty"`
    // ExclusionGroupId the ID of the exclusion group to update.
    ExclusionGroupId string `json:"ExclusionGroupId,omitempty"`
    // Name friendly name of the exclusion group.
    Name string `json:"Name,omitempty"`
}

// UpdateExperimentRequest given a title entity token and experiment details, this API updates the experiment. If an experiment is already running,
// only the description and duration properties can be updated.
type UpdateExperimentRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Description description of the experiment.
    Description string `json:"Description,omitempty"`
    // Duration the duration of the experiment, in days.
    Duration uint32 `json:"Duration,omitempty"`
    // EndDate when experiment should end.
    EndDate time.Time `json:"EndDate,omitempty"`
    // ExclusionGroupId id of the exclusion group.
    ExclusionGroupId string `json:"ExclusionGroupId,omitempty"`
    // ExclusionGroupTrafficAllocation percentage of exclusion group traffic that will see this experiment.
    ExclusionGroupTrafficAllocation uint32 `json:"ExclusionGroupTrafficAllocation,omitempty"`
    // ExperimentType type of experiment.
    ExperimentType ExperimentType `json:"ExperimentType,omitempty"`
    // Id id of the experiment.
    Id string `json:"Id,omitempty"`
    // Name friendly name of the experiment.
    Name string `json:"Name,omitempty"`
    // SegmentId id of the segment to which this experiment applies. Defaults to the 'All Players' segment.
    SegmentId string `json:"SegmentId,omitempty"`
    // StartDate when experiment should start.
    StartDate time.Time `json:"StartDate,omitempty"`
    // TitlePlayerAccountTestIds list of title player account IDs that automatically receive treatments in the experiment, but are not included when
// calculating experiment metrics.
    TitlePlayerAccountTestIds []string `json:"TitlePlayerAccountTestIds,omitempty"`
    // Variants list of variants for the experiment.
    Variants []VariantModel `json:"Variants,omitempty"`
}

// Variable 
type VariableModel struct {
    // Name name of the variable.
    Name string `json:"Name,omitempty"`
    // Value value of the variable.
    Value string `json:"Value,omitempty"`
}

// Variant 
type VariantModel struct {
    // Description description of the variant.
    Description string `json:"Description,omitempty"`
    // Id id of the variant.
    Id string `json:"Id,omitempty"`
    // IsControl specifies if variant is control for experiment.
    IsControl bool `json:"IsControl"`
    // Name name of the variant.
    Name string `json:"Name,omitempty"`
    // TitleDataOverrideLabel id of the TitleDataOverride to use with this variant.
    TitleDataOverrideLabel string `json:"TitleDataOverrideLabel,omitempty"`
    // TrafficPercentage percentage of target audience traffic that will see this variant.
    TrafficPercentage uint32 `json:"TrafficPercentage,omitempty"`
    // Variables variables returned by this variant.
    Variables []VariableModel `json:"Variables,omitempty"`
}
