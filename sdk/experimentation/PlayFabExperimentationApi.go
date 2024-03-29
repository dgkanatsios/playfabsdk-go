package experimentation

// This code was generated by a tool. Any changes may be overwritten

import (
    "encoding/json"

    playfab "github.com/dgkanatsios/playfabsdk-go/sdk"

    "github.com/mitchellh/mapstructure"
)

// CreateExclusionGroup creates a new experiment exclusion group for a title.
// https://api.playfab.com/Documentation/Experimentation/method/CreateExclusionGroup
func CreateExclusionGroup(settings *playfab.Settings, postData *CreateExclusionGroupRequestModel, entityToken string) (*CreateExclusionGroupResultModel, error) {
    if entityToken == "" {
        return nil, playfab.NewCustomError("entityToken should not be an empty string", playfab.ErrorGeneric)
    }
    b, errMarshal := json.Marshal(postData)
    if errMarshal != nil {
        return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
    }

    sourceMap, err := playfab.Request(settings, b, "/Experimentation/CreateExclusionGroup", "X-EntityToken", entityToken)
    if err != nil {
        return nil, err
    }
    
    result := &CreateExclusionGroupResultModel{}

    config := mapstructure.DecoderConfig{
        DecodeHook: playfab.StringToDateTimeHook,
        Result:     result,
    }
    
    decoder, errDecoding := mapstructure.NewDecoder(&config)
    if errDecoding != nil {
        return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
    }
   
    errDecoding = decoder.Decode(sourceMap)
    if errDecoding != nil {
        return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
    }

    return result, nil
}

// CreateExperiment creates a new experiment for a title.
// https://api.playfab.com/Documentation/Experimentation/method/CreateExperiment
func CreateExperiment(settings *playfab.Settings, postData *CreateExperimentRequestModel, entityToken string) (*CreateExperimentResultModel, error) {
    if entityToken == "" {
        return nil, playfab.NewCustomError("entityToken should not be an empty string", playfab.ErrorGeneric)
    }
    b, errMarshal := json.Marshal(postData)
    if errMarshal != nil {
        return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
    }

    sourceMap, err := playfab.Request(settings, b, "/Experimentation/CreateExperiment", "X-EntityToken", entityToken)
    if err != nil {
        return nil, err
    }
    
    result := &CreateExperimentResultModel{}

    config := mapstructure.DecoderConfig{
        DecodeHook: playfab.StringToDateTimeHook,
        Result:     result,
    }
    
    decoder, errDecoding := mapstructure.NewDecoder(&config)
    if errDecoding != nil {
        return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
    }
   
    errDecoding = decoder.Decode(sourceMap)
    if errDecoding != nil {
        return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
    }

    return result, nil
}

// DeleteExclusionGroup deletes an existing exclusion group for a title.
// https://api.playfab.com/Documentation/Experimentation/method/DeleteExclusionGroup
func DeleteExclusionGroup(settings *playfab.Settings, postData *DeleteExclusionGroupRequestModel, entityToken string) (*EmptyResponseModel, error) {
    if entityToken == "" {
        return nil, playfab.NewCustomError("entityToken should not be an empty string", playfab.ErrorGeneric)
    }
    b, errMarshal := json.Marshal(postData)
    if errMarshal != nil {
        return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
    }

    sourceMap, err := playfab.Request(settings, b, "/Experimentation/DeleteExclusionGroup", "X-EntityToken", entityToken)
    if err != nil {
        return nil, err
    }
    
    result := &EmptyResponseModel{}

    config := mapstructure.DecoderConfig{
        DecodeHook: playfab.StringToDateTimeHook,
        Result:     result,
    }
    
    decoder, errDecoding := mapstructure.NewDecoder(&config)
    if errDecoding != nil {
        return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
    }
   
    errDecoding = decoder.Decode(sourceMap)
    if errDecoding != nil {
        return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
    }

    return result, nil
}

// DeleteExperiment deletes an existing experiment for a title.
// https://api.playfab.com/Documentation/Experimentation/method/DeleteExperiment
func DeleteExperiment(settings *playfab.Settings, postData *DeleteExperimentRequestModel, entityToken string) (*EmptyResponseModel, error) {
    if entityToken == "" {
        return nil, playfab.NewCustomError("entityToken should not be an empty string", playfab.ErrorGeneric)
    }
    b, errMarshal := json.Marshal(postData)
    if errMarshal != nil {
        return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
    }

    sourceMap, err := playfab.Request(settings, b, "/Experimentation/DeleteExperiment", "X-EntityToken", entityToken)
    if err != nil {
        return nil, err
    }
    
    result := &EmptyResponseModel{}

    config := mapstructure.DecoderConfig{
        DecodeHook: playfab.StringToDateTimeHook,
        Result:     result,
    }
    
    decoder, errDecoding := mapstructure.NewDecoder(&config)
    if errDecoding != nil {
        return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
    }
   
    errDecoding = decoder.Decode(sourceMap)
    if errDecoding != nil {
        return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
    }

    return result, nil
}

// GetExclusionGroups gets the details of all exclusion groups for a title.
// https://api.playfab.com/Documentation/Experimentation/method/GetExclusionGroups
func GetExclusionGroups(settings *playfab.Settings, postData *GetExclusionGroupsRequestModel, entityToken string) (*GetExclusionGroupsResultModel, error) {
    if entityToken == "" {
        return nil, playfab.NewCustomError("entityToken should not be an empty string", playfab.ErrorGeneric)
    }
    b, errMarshal := json.Marshal(postData)
    if errMarshal != nil {
        return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
    }

    sourceMap, err := playfab.Request(settings, b, "/Experimentation/GetExclusionGroups", "X-EntityToken", entityToken)
    if err != nil {
        return nil, err
    }
    
    result := &GetExclusionGroupsResultModel{}

    config := mapstructure.DecoderConfig{
        DecodeHook: playfab.StringToDateTimeHook,
        Result:     result,
    }
    
    decoder, errDecoding := mapstructure.NewDecoder(&config)
    if errDecoding != nil {
        return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
    }
   
    errDecoding = decoder.Decode(sourceMap)
    if errDecoding != nil {
        return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
    }

    return result, nil
}

// GetExclusionGroupTraffic gets the details of all exclusion groups for a title.
// https://api.playfab.com/Documentation/Experimentation/method/GetExclusionGroupTraffic
func GetExclusionGroupTraffic(settings *playfab.Settings, postData *GetExclusionGroupTrafficRequestModel, entityToken string) (*GetExclusionGroupTrafficResultModel, error) {
    if entityToken == "" {
        return nil, playfab.NewCustomError("entityToken should not be an empty string", playfab.ErrorGeneric)
    }
    b, errMarshal := json.Marshal(postData)
    if errMarshal != nil {
        return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
    }

    sourceMap, err := playfab.Request(settings, b, "/Experimentation/GetExclusionGroupTraffic", "X-EntityToken", entityToken)
    if err != nil {
        return nil, err
    }
    
    result := &GetExclusionGroupTrafficResultModel{}

    config := mapstructure.DecoderConfig{
        DecodeHook: playfab.StringToDateTimeHook,
        Result:     result,
    }
    
    decoder, errDecoding := mapstructure.NewDecoder(&config)
    if errDecoding != nil {
        return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
    }
   
    errDecoding = decoder.Decode(sourceMap)
    if errDecoding != nil {
        return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
    }

    return result, nil
}

// GetExperiments gets the details of all experiments for a title.
// https://api.playfab.com/Documentation/Experimentation/method/GetExperiments
func GetExperiments(settings *playfab.Settings, postData *GetExperimentsRequestModel, entityToken string) (*GetExperimentsResultModel, error) {
    if entityToken == "" {
        return nil, playfab.NewCustomError("entityToken should not be an empty string", playfab.ErrorGeneric)
    }
    b, errMarshal := json.Marshal(postData)
    if errMarshal != nil {
        return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
    }

    sourceMap, err := playfab.Request(settings, b, "/Experimentation/GetExperiments", "X-EntityToken", entityToken)
    if err != nil {
        return nil, err
    }
    
    result := &GetExperimentsResultModel{}

    config := mapstructure.DecoderConfig{
        DecodeHook: playfab.StringToDateTimeHook,
        Result:     result,
    }
    
    decoder, errDecoding := mapstructure.NewDecoder(&config)
    if errDecoding != nil {
        return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
    }
   
    errDecoding = decoder.Decode(sourceMap)
    if errDecoding != nil {
        return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
    }

    return result, nil
}

// GetLatestScorecard gets the latest scorecard of the experiment for the title.
// https://api.playfab.com/Documentation/Experimentation/method/GetLatestScorecard
func GetLatestScorecard(settings *playfab.Settings, postData *GetLatestScorecardRequestModel, entityToken string) (*GetLatestScorecardResultModel, error) {
    if entityToken == "" {
        return nil, playfab.NewCustomError("entityToken should not be an empty string", playfab.ErrorGeneric)
    }
    b, errMarshal := json.Marshal(postData)
    if errMarshal != nil {
        return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
    }

    sourceMap, err := playfab.Request(settings, b, "/Experimentation/GetLatestScorecard", "X-EntityToken", entityToken)
    if err != nil {
        return nil, err
    }
    
    result := &GetLatestScorecardResultModel{}

    config := mapstructure.DecoderConfig{
        DecodeHook: playfab.StringToDateTimeHook,
        Result:     result,
    }
    
    decoder, errDecoding := mapstructure.NewDecoder(&config)
    if errDecoding != nil {
        return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
    }
   
    errDecoding = decoder.Decode(sourceMap)
    if errDecoding != nil {
        return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
    }

    return result, nil
}

// GetTreatmentAssignment gets the treatment assignments for a player for every running experiment in the title.
// https://api.playfab.com/Documentation/Experimentation/method/GetTreatmentAssignment
func GetTreatmentAssignment(settings *playfab.Settings, postData *GetTreatmentAssignmentRequestModel, entityToken string) (*GetTreatmentAssignmentResultModel, error) {
    if entityToken == "" {
        return nil, playfab.NewCustomError("entityToken should not be an empty string", playfab.ErrorGeneric)
    }
    b, errMarshal := json.Marshal(postData)
    if errMarshal != nil {
        return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
    }

    sourceMap, err := playfab.Request(settings, b, "/Experimentation/GetTreatmentAssignment", "X-EntityToken", entityToken)
    if err != nil {
        return nil, err
    }
    
    result := &GetTreatmentAssignmentResultModel{}

    config := mapstructure.DecoderConfig{
        DecodeHook: playfab.StringToDateTimeHook,
        Result:     result,
    }
    
    decoder, errDecoding := mapstructure.NewDecoder(&config)
    if errDecoding != nil {
        return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
    }
   
    errDecoding = decoder.Decode(sourceMap)
    if errDecoding != nil {
        return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
    }

    return result, nil
}

// StartExperiment starts an existing experiment for a title.
// https://api.playfab.com/Documentation/Experimentation/method/StartExperiment
func StartExperiment(settings *playfab.Settings, postData *StartExperimentRequestModel, entityToken string) (*EmptyResponseModel, error) {
    if entityToken == "" {
        return nil, playfab.NewCustomError("entityToken should not be an empty string", playfab.ErrorGeneric)
    }
    b, errMarshal := json.Marshal(postData)
    if errMarshal != nil {
        return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
    }

    sourceMap, err := playfab.Request(settings, b, "/Experimentation/StartExperiment", "X-EntityToken", entityToken)
    if err != nil {
        return nil, err
    }
    
    result := &EmptyResponseModel{}

    config := mapstructure.DecoderConfig{
        DecodeHook: playfab.StringToDateTimeHook,
        Result:     result,
    }
    
    decoder, errDecoding := mapstructure.NewDecoder(&config)
    if errDecoding != nil {
        return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
    }
   
    errDecoding = decoder.Decode(sourceMap)
    if errDecoding != nil {
        return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
    }

    return result, nil
}

// StopExperiment stops an existing experiment for a title.
// https://api.playfab.com/Documentation/Experimentation/method/StopExperiment
func StopExperiment(settings *playfab.Settings, postData *StopExperimentRequestModel, entityToken string) (*EmptyResponseModel, error) {
    if entityToken == "" {
        return nil, playfab.NewCustomError("entityToken should not be an empty string", playfab.ErrorGeneric)
    }
    b, errMarshal := json.Marshal(postData)
    if errMarshal != nil {
        return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
    }

    sourceMap, err := playfab.Request(settings, b, "/Experimentation/StopExperiment", "X-EntityToken", entityToken)
    if err != nil {
        return nil, err
    }
    
    result := &EmptyResponseModel{}

    config := mapstructure.DecoderConfig{
        DecodeHook: playfab.StringToDateTimeHook,
        Result:     result,
    }
    
    decoder, errDecoding := mapstructure.NewDecoder(&config)
    if errDecoding != nil {
        return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
    }
   
    errDecoding = decoder.Decode(sourceMap)
    if errDecoding != nil {
        return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
    }

    return result, nil
}

// UpdateExclusionGroup updates an existing exclusion group for a title.
// https://api.playfab.com/Documentation/Experimentation/method/UpdateExclusionGroup
func UpdateExclusionGroup(settings *playfab.Settings, postData *UpdateExclusionGroupRequestModel, entityToken string) (*EmptyResponseModel, error) {
    if entityToken == "" {
        return nil, playfab.NewCustomError("entityToken should not be an empty string", playfab.ErrorGeneric)
    }
    b, errMarshal := json.Marshal(postData)
    if errMarshal != nil {
        return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
    }

    sourceMap, err := playfab.Request(settings, b, "/Experimentation/UpdateExclusionGroup", "X-EntityToken", entityToken)
    if err != nil {
        return nil, err
    }
    
    result := &EmptyResponseModel{}

    config := mapstructure.DecoderConfig{
        DecodeHook: playfab.StringToDateTimeHook,
        Result:     result,
    }
    
    decoder, errDecoding := mapstructure.NewDecoder(&config)
    if errDecoding != nil {
        return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
    }
   
    errDecoding = decoder.Decode(sourceMap)
    if errDecoding != nil {
        return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
    }

    return result, nil
}

// UpdateExperiment updates an existing experiment for a title.
// https://api.playfab.com/Documentation/Experimentation/method/UpdateExperiment
func UpdateExperiment(settings *playfab.Settings, postData *UpdateExperimentRequestModel, entityToken string) (*EmptyResponseModel, error) {
    if entityToken == "" {
        return nil, playfab.NewCustomError("entityToken should not be an empty string", playfab.ErrorGeneric)
    }
    b, errMarshal := json.Marshal(postData)
    if errMarshal != nil {
        return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
    }

    sourceMap, err := playfab.Request(settings, b, "/Experimentation/UpdateExperiment", "X-EntityToken", entityToken)
    if err != nil {
        return nil, err
    }
    
    result := &EmptyResponseModel{}

    config := mapstructure.DecoderConfig{
        DecodeHook: playfab.StringToDateTimeHook,
        Result:     result,
    }
    
    decoder, errDecoding := mapstructure.NewDecoder(&config)
    if errDecoding != nil {
        return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
    }
   
    errDecoding = decoder.Decode(sourceMap)
    if errDecoding != nil {
        return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
    }

    return result, nil
}



