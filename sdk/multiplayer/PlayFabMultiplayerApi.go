package multiplayer

// This code was generated by a tool. Any changes may be overwritten

import (
    "encoding/json"

    playfab "github.com/dgkanatsios/playfabsdk-go/sdk"

    "github.com/mitchellh/mapstructure"
)

// CreateBuildWithCustomContainer creates a multiplayer server build with a custom container.
// https://api.playfab.com/Documentation/Multiplayer/method/CreateBuildWithCustomContainer
func CreateBuildWithCustomContainer(settings *playfab.Settings, postData *CreateBuildWithCustomContainerRequestModel, entityToken string) (*CreateBuildWithCustomContainerResponseModel, error) {
    if entityToken == "" {
        return nil, playfab.NewCustomError("entityToken should not be an empty string", playfab.ErrorGeneric)
    }
    b, errMarshal := json.Marshal(postData)
    if errMarshal != nil {
        return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
    }

    sourceMap, err := playfab.Request(settings, b, "/MultiplayerServer/CreateBuildWithCustomContainer", "X-EntityToken", entityToken)
    if err != nil {
        return nil, err
    }
    
    result := &CreateBuildWithCustomContainerResponseModel{}

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

// CreateBuildWithManagedContainer creates a multiplayer server build with a managed container.
// https://api.playfab.com/Documentation/Multiplayer/method/CreateBuildWithManagedContainer
func CreateBuildWithManagedContainer(settings *playfab.Settings, postData *CreateBuildWithManagedContainerRequestModel, entityToken string) (*CreateBuildWithManagedContainerResponseModel, error) {
    if entityToken == "" {
        return nil, playfab.NewCustomError("entityToken should not be an empty string", playfab.ErrorGeneric)
    }
    b, errMarshal := json.Marshal(postData)
    if errMarshal != nil {
        return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
    }

    sourceMap, err := playfab.Request(settings, b, "/MultiplayerServer/CreateBuildWithManagedContainer", "X-EntityToken", entityToken)
    if err != nil {
        return nil, err
    }
    
    result := &CreateBuildWithManagedContainerResponseModel{}

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

// CreateRemoteUser creates a remote user to log on to a VM for a multiplayer server build.
// https://api.playfab.com/Documentation/Multiplayer/method/CreateRemoteUser
func CreateRemoteUser(settings *playfab.Settings, postData *CreateRemoteUserRequestModel, entityToken string) (*CreateRemoteUserResponseModel, error) {
    if entityToken == "" {
        return nil, playfab.NewCustomError("entityToken should not be an empty string", playfab.ErrorGeneric)
    }
    b, errMarshal := json.Marshal(postData)
    if errMarshal != nil {
        return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
    }

    sourceMap, err := playfab.Request(settings, b, "/MultiplayerServer/CreateRemoteUser", "X-EntityToken", entityToken)
    if err != nil {
        return nil, err
    }
    
    result := &CreateRemoteUserResponseModel{}

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

// DeleteAsset deletes a multiplayer server game asset for a title.
// https://api.playfab.com/Documentation/Multiplayer/method/DeleteAsset
func DeleteAsset(settings *playfab.Settings, postData *DeleteAssetRequestModel, entityToken string) (*EmptyResponseModel, error) {
    if entityToken == "" {
        return nil, playfab.NewCustomError("entityToken should not be an empty string", playfab.ErrorGeneric)
    }
    b, errMarshal := json.Marshal(postData)
    if errMarshal != nil {
        return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
    }

    sourceMap, err := playfab.Request(settings, b, "/MultiplayerServer/DeleteAsset", "X-EntityToken", entityToken)
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

// DeleteBuild deletes a multiplayer server build.
// https://api.playfab.com/Documentation/Multiplayer/method/DeleteBuild
func DeleteBuild(settings *playfab.Settings, postData *DeleteBuildRequestModel, entityToken string) (*EmptyResponseModel, error) {
    if entityToken == "" {
        return nil, playfab.NewCustomError("entityToken should not be an empty string", playfab.ErrorGeneric)
    }
    b, errMarshal := json.Marshal(postData)
    if errMarshal != nil {
        return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
    }

    sourceMap, err := playfab.Request(settings, b, "/MultiplayerServer/DeleteBuild", "X-EntityToken", entityToken)
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

// DeleteCertificate deletes a multiplayer server game certificate.
// https://api.playfab.com/Documentation/Multiplayer/method/DeleteCertificate
func DeleteCertificate(settings *playfab.Settings, postData *DeleteCertificateRequestModel, entityToken string) (*EmptyResponseModel, error) {
    if entityToken == "" {
        return nil, playfab.NewCustomError("entityToken should not be an empty string", playfab.ErrorGeneric)
    }
    b, errMarshal := json.Marshal(postData)
    if errMarshal != nil {
        return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
    }

    sourceMap, err := playfab.Request(settings, b, "/MultiplayerServer/DeleteCertificate", "X-EntityToken", entityToken)
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

// DeleteRemoteUser deletes a remote user to log on to a VM for a multiplayer server build.
// https://api.playfab.com/Documentation/Multiplayer/method/DeleteRemoteUser
func DeleteRemoteUser(settings *playfab.Settings, postData *DeleteRemoteUserRequestModel, entityToken string) (*EmptyResponseModel, error) {
    if entityToken == "" {
        return nil, playfab.NewCustomError("entityToken should not be an empty string", playfab.ErrorGeneric)
    }
    b, errMarshal := json.Marshal(postData)
    if errMarshal != nil {
        return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
    }

    sourceMap, err := playfab.Request(settings, b, "/MultiplayerServer/DeleteRemoteUser", "X-EntityToken", entityToken)
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

// EnableMultiplayerServersForTitle enables the multiplayer server feature for a title.
// https://api.playfab.com/Documentation/Multiplayer/method/EnableMultiplayerServersForTitle
func EnableMultiplayerServersForTitle(settings *playfab.Settings, postData *EnableMultiplayerServersForTitleRequestModel, entityToken string) (*EnableMultiplayerServersForTitleResponseModel, error) {
    if entityToken == "" {
        return nil, playfab.NewCustomError("entityToken should not be an empty string", playfab.ErrorGeneric)
    }
    b, errMarshal := json.Marshal(postData)
    if errMarshal != nil {
        return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
    }

    sourceMap, err := playfab.Request(settings, b, "/MultiplayerServer/EnableMultiplayerServersForTitle", "X-EntityToken", entityToken)
    if err != nil {
        return nil, err
    }
    
    result := &EnableMultiplayerServersForTitleResponseModel{}

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

// GetAssetUploadUrl gets the URL to upload assets to.
// https://api.playfab.com/Documentation/Multiplayer/method/GetAssetUploadUrl
func GetAssetUploadUrl(settings *playfab.Settings, postData *GetAssetUploadUrlRequestModel, entityToken string) (*GetAssetUploadUrlResponseModel, error) {
    if entityToken == "" {
        return nil, playfab.NewCustomError("entityToken should not be an empty string", playfab.ErrorGeneric)
    }
    b, errMarshal := json.Marshal(postData)
    if errMarshal != nil {
        return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
    }

    sourceMap, err := playfab.Request(settings, b, "/MultiplayerServer/GetAssetUploadUrl", "X-EntityToken", entityToken)
    if err != nil {
        return nil, err
    }
    
    result := &GetAssetUploadUrlResponseModel{}

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

// GetBuild gets a multiplayer server build.
// https://api.playfab.com/Documentation/Multiplayer/method/GetBuild
func GetBuild(settings *playfab.Settings, postData *GetBuildRequestModel, entityToken string) (*GetBuildResponseModel, error) {
    if entityToken == "" {
        return nil, playfab.NewCustomError("entityToken should not be an empty string", playfab.ErrorGeneric)
    }
    b, errMarshal := json.Marshal(postData)
    if errMarshal != nil {
        return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
    }

    sourceMap, err := playfab.Request(settings, b, "/MultiplayerServer/GetBuild", "X-EntityToken", entityToken)
    if err != nil {
        return nil, err
    }
    
    result := &GetBuildResponseModel{}

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

// GetContainerRegistryCredentials gets the credentials to the container registry.
// https://api.playfab.com/Documentation/Multiplayer/method/GetContainerRegistryCredentials
func GetContainerRegistryCredentials(settings *playfab.Settings, postData *GetContainerRegistryCredentialsRequestModel, entityToken string) (*GetContainerRegistryCredentialsResponseModel, error) {
    if entityToken == "" {
        return nil, playfab.NewCustomError("entityToken should not be an empty string", playfab.ErrorGeneric)
    }
    b, errMarshal := json.Marshal(postData)
    if errMarshal != nil {
        return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
    }

    sourceMap, err := playfab.Request(settings, b, "/MultiplayerServer/GetContainerRegistryCredentials", "X-EntityToken", entityToken)
    if err != nil {
        return nil, err
    }
    
    result := &GetContainerRegistryCredentialsResponseModel{}

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

// GetMultiplayerServerDetails gets multiplayer server session details for a build.
// https://api.playfab.com/Documentation/Multiplayer/method/GetMultiplayerServerDetails
func GetMultiplayerServerDetails(settings *playfab.Settings, postData *GetMultiplayerServerDetailsRequestModel, entityToken string) (*GetMultiplayerServerDetailsResponseModel, error) {
    if entityToken == "" {
        return nil, playfab.NewCustomError("entityToken should not be an empty string", playfab.ErrorGeneric)
    }
    b, errMarshal := json.Marshal(postData)
    if errMarshal != nil {
        return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
    }

    sourceMap, err := playfab.Request(settings, b, "/MultiplayerServer/GetMultiplayerServerDetails", "X-EntityToken", entityToken)
    if err != nil {
        return nil, err
    }
    
    result := &GetMultiplayerServerDetailsResponseModel{}

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

// GetRemoteLoginEndpoint gets a remote login endpoint to a VM that is hosting a multiplayer server build.
// https://api.playfab.com/Documentation/Multiplayer/method/GetRemoteLoginEndpoint
func GetRemoteLoginEndpoint(settings *playfab.Settings, postData *GetRemoteLoginEndpointRequestModel, entityToken string) (*GetRemoteLoginEndpointResponseModel, error) {
    if entityToken == "" {
        return nil, playfab.NewCustomError("entityToken should not be an empty string", playfab.ErrorGeneric)
    }
    b, errMarshal := json.Marshal(postData)
    if errMarshal != nil {
        return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
    }

    sourceMap, err := playfab.Request(settings, b, "/MultiplayerServer/GetRemoteLoginEndpoint", "X-EntityToken", entityToken)
    if err != nil {
        return nil, err
    }
    
    result := &GetRemoteLoginEndpointResponseModel{}

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

// GetTitleEnabledForMultiplayerServersStatus gets the status of whether a title is enabled for the multiplayer server feature.
// https://api.playfab.com/Documentation/Multiplayer/method/GetTitleEnabledForMultiplayerServersStatus
func GetTitleEnabledForMultiplayerServersStatus(settings *playfab.Settings, postData *GetTitleEnabledForMultiplayerServersStatusRequestModel, entityToken string) (*GetTitleEnabledForMultiplayerServersStatusResponseModel, error) {
    if entityToken == "" {
        return nil, playfab.NewCustomError("entityToken should not be an empty string", playfab.ErrorGeneric)
    }
    b, errMarshal := json.Marshal(postData)
    if errMarshal != nil {
        return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
    }

    sourceMap, err := playfab.Request(settings, b, "/MultiplayerServer/GetTitleEnabledForMultiplayerServersStatus", "X-EntityToken", entityToken)
    if err != nil {
        return nil, err
    }
    
    result := &GetTitleEnabledForMultiplayerServersStatusResponseModel{}

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

// ListArchivedMultiplayerServers lists archived multiplayer server sessions for a build.
// https://api.playfab.com/Documentation/Multiplayer/method/ListArchivedMultiplayerServers
func ListArchivedMultiplayerServers(settings *playfab.Settings, postData *ListMultiplayerServersRequestModel, entityToken string) (*ListMultiplayerServersResponseModel, error) {
    if entityToken == "" {
        return nil, playfab.NewCustomError("entityToken should not be an empty string", playfab.ErrorGeneric)
    }
    b, errMarshal := json.Marshal(postData)
    if errMarshal != nil {
        return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
    }

    sourceMap, err := playfab.Request(settings, b, "/MultiplayerServer/ListArchivedMultiplayerServers", "X-EntityToken", entityToken)
    if err != nil {
        return nil, err
    }
    
    result := &ListMultiplayerServersResponseModel{}

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

// ListAssetSummaries lists multiplayer server game assets for a title.
// https://api.playfab.com/Documentation/Multiplayer/method/ListAssetSummaries
func ListAssetSummaries(settings *playfab.Settings, postData *ListAssetSummariesRequestModel, entityToken string) (*ListAssetSummariesResponseModel, error) {
    if entityToken == "" {
        return nil, playfab.NewCustomError("entityToken should not be an empty string", playfab.ErrorGeneric)
    }
    b, errMarshal := json.Marshal(postData)
    if errMarshal != nil {
        return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
    }

    sourceMap, err := playfab.Request(settings, b, "/MultiplayerServer/ListAssetSummaries", "X-EntityToken", entityToken)
    if err != nil {
        return nil, err
    }
    
    result := &ListAssetSummariesResponseModel{}

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

// ListBuildSummaries lists summarized details of all multiplayer server builds for a title.
// https://api.playfab.com/Documentation/Multiplayer/method/ListBuildSummaries
func ListBuildSummaries(settings *playfab.Settings, postData *ListBuildSummariesRequestModel, entityToken string) (*ListBuildSummariesResponseModel, error) {
    if entityToken == "" {
        return nil, playfab.NewCustomError("entityToken should not be an empty string", playfab.ErrorGeneric)
    }
    b, errMarshal := json.Marshal(postData)
    if errMarshal != nil {
        return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
    }

    sourceMap, err := playfab.Request(settings, b, "/MultiplayerServer/ListBuildSummaries", "X-EntityToken", entityToken)
    if err != nil {
        return nil, err
    }
    
    result := &ListBuildSummariesResponseModel{}

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

// ListCertificateSummaries lists multiplayer server game certificates for a title.
// https://api.playfab.com/Documentation/Multiplayer/method/ListCertificateSummaries
func ListCertificateSummaries(settings *playfab.Settings, postData *ListCertificateSummariesRequestModel, entityToken string) (*ListCertificateSummariesResponseModel, error) {
    if entityToken == "" {
        return nil, playfab.NewCustomError("entityToken should not be an empty string", playfab.ErrorGeneric)
    }
    b, errMarshal := json.Marshal(postData)
    if errMarshal != nil {
        return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
    }

    sourceMap, err := playfab.Request(settings, b, "/MultiplayerServer/ListCertificateSummaries", "X-EntityToken", entityToken)
    if err != nil {
        return nil, err
    }
    
    result := &ListCertificateSummariesResponseModel{}

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

// ListContainerImages lists custom container images for a title.
// https://api.playfab.com/Documentation/Multiplayer/method/ListContainerImages
func ListContainerImages(settings *playfab.Settings, postData *ListContainerImagesRequestModel, entityToken string) (*ListContainerImagesResponseModel, error) {
    if entityToken == "" {
        return nil, playfab.NewCustomError("entityToken should not be an empty string", playfab.ErrorGeneric)
    }
    b, errMarshal := json.Marshal(postData)
    if errMarshal != nil {
        return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
    }

    sourceMap, err := playfab.Request(settings, b, "/MultiplayerServer/ListContainerImages", "X-EntityToken", entityToken)
    if err != nil {
        return nil, err
    }
    
    result := &ListContainerImagesResponseModel{}

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

// ListContainerImageTags lists the tags for a custom container image.
// https://api.playfab.com/Documentation/Multiplayer/method/ListContainerImageTags
func ListContainerImageTags(settings *playfab.Settings, postData *ListContainerImageTagsRequestModel, entityToken string) (*ListContainerImageTagsResponseModel, error) {
    if entityToken == "" {
        return nil, playfab.NewCustomError("entityToken should not be an empty string", playfab.ErrorGeneric)
    }
    b, errMarshal := json.Marshal(postData)
    if errMarshal != nil {
        return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
    }

    sourceMap, err := playfab.Request(settings, b, "/MultiplayerServer/ListContainerImageTags", "X-EntityToken", entityToken)
    if err != nil {
        return nil, err
    }
    
    result := &ListContainerImageTagsResponseModel{}

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

// ListMultiplayerServers lists multiplayer server sessions for a build.
// https://api.playfab.com/Documentation/Multiplayer/method/ListMultiplayerServers
func ListMultiplayerServers(settings *playfab.Settings, postData *ListMultiplayerServersRequestModel, entityToken string) (*ListMultiplayerServersResponseModel, error) {
    if entityToken == "" {
        return nil, playfab.NewCustomError("entityToken should not be an empty string", playfab.ErrorGeneric)
    }
    b, errMarshal := json.Marshal(postData)
    if errMarshal != nil {
        return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
    }

    sourceMap, err := playfab.Request(settings, b, "/MultiplayerServer/ListMultiplayerServers", "X-EntityToken", entityToken)
    if err != nil {
        return nil, err
    }
    
    result := &ListMultiplayerServersResponseModel{}

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

// ListQosServers lists quality of service servers.
// https://api.playfab.com/Documentation/Multiplayer/method/ListQosServers
func ListQosServers(settings *playfab.Settings, postData *ListQosServersRequestModel, ) (*ListQosServersResponseModel, error) {

    b, errMarshal := json.Marshal(postData)
    if errMarshal != nil {
        return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
    }

    sourceMap, err := playfab.Request(settings, b, "/MultiplayerServer/ListQosServers", "", "")
    if err != nil {
        return nil, err
    }
    
    result := &ListQosServersResponseModel{}

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

// ListVirtualMachineSummaries lists virtual machines for a title.
// https://api.playfab.com/Documentation/Multiplayer/method/ListVirtualMachineSummaries
func ListVirtualMachineSummaries(settings *playfab.Settings, postData *ListVirtualMachineSummariesRequestModel, entityToken string) (*ListVirtualMachineSummariesResponseModel, error) {
    if entityToken == "" {
        return nil, playfab.NewCustomError("entityToken should not be an empty string", playfab.ErrorGeneric)
    }
    b, errMarshal := json.Marshal(postData)
    if errMarshal != nil {
        return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
    }

    sourceMap, err := playfab.Request(settings, b, "/MultiplayerServer/ListVirtualMachineSummaries", "X-EntityToken", entityToken)
    if err != nil {
        return nil, err
    }
    
    result := &ListVirtualMachineSummariesResponseModel{}

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

// RequestMultiplayerServer request a multiplayer server session. Accepts tokens for title and if game client accesss is enabled, allows game client
// to request a server with player entity token.
// https://api.playfab.com/Documentation/Multiplayer/method/RequestMultiplayerServer
func RequestMultiplayerServer(settings *playfab.Settings, postData *RequestMultiplayerServerRequestModel, entityToken string) (*RequestMultiplayerServerResponseModel, error) {
    if entityToken == "" {
        return nil, playfab.NewCustomError("entityToken should not be an empty string", playfab.ErrorGeneric)
    }
    b, errMarshal := json.Marshal(postData)
    if errMarshal != nil {
        return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
    }

    sourceMap, err := playfab.Request(settings, b, "/MultiplayerServer/RequestMultiplayerServer", "X-EntityToken", entityToken)
    if err != nil {
        return nil, err
    }
    
    result := &RequestMultiplayerServerResponseModel{}

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

// RolloverContainerRegistryCredentials rolls over the credentials to the container registry.
// https://api.playfab.com/Documentation/Multiplayer/method/RolloverContainerRegistryCredentials
func RolloverContainerRegistryCredentials(settings *playfab.Settings, postData *RolloverContainerRegistryCredentialsRequestModel, entityToken string) (*RolloverContainerRegistryCredentialsResponseModel, error) {
    if entityToken == "" {
        return nil, playfab.NewCustomError("entityToken should not be an empty string", playfab.ErrorGeneric)
    }
    b, errMarshal := json.Marshal(postData)
    if errMarshal != nil {
        return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
    }

    sourceMap, err := playfab.Request(settings, b, "/MultiplayerServer/RolloverContainerRegistryCredentials", "X-EntityToken", entityToken)
    if err != nil {
        return nil, err
    }
    
    result := &RolloverContainerRegistryCredentialsResponseModel{}

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

// ShutdownMultiplayerServer shuts down a multiplayer server session.
// https://api.playfab.com/Documentation/Multiplayer/method/ShutdownMultiplayerServer
func ShutdownMultiplayerServer(settings *playfab.Settings, postData *ShutdownMultiplayerServerRequestModel, entityToken string) (*EmptyResponseModel, error) {
    if entityToken == "" {
        return nil, playfab.NewCustomError("entityToken should not be an empty string", playfab.ErrorGeneric)
    }
    b, errMarshal := json.Marshal(postData)
    if errMarshal != nil {
        return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
    }

    sourceMap, err := playfab.Request(settings, b, "/MultiplayerServer/ShutdownMultiplayerServer", "X-EntityToken", entityToken)
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

// UpdateBuildRegions updates a multiplayer server build's regions.
// https://api.playfab.com/Documentation/Multiplayer/method/UpdateBuildRegions
func UpdateBuildRegions(settings *playfab.Settings, postData *UpdateBuildRegionsRequestModel, entityToken string) (*EmptyResponseModel, error) {
    if entityToken == "" {
        return nil, playfab.NewCustomError("entityToken should not be an empty string", playfab.ErrorGeneric)
    }
    b, errMarshal := json.Marshal(postData)
    if errMarshal != nil {
        return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
    }

    sourceMap, err := playfab.Request(settings, b, "/MultiplayerServer/UpdateBuildRegions", "X-EntityToken", entityToken)
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

// UploadCertificate uploads a multiplayer server game certificate.
// https://api.playfab.com/Documentation/Multiplayer/method/UploadCertificate
func UploadCertificate(settings *playfab.Settings, postData *UploadCertificateRequestModel, entityToken string) (*EmptyResponseModel, error) {
    if entityToken == "" {
        return nil, playfab.NewCustomError("entityToken should not be an empty string", playfab.ErrorGeneric)
    }
    b, errMarshal := json.Marshal(postData)
    if errMarshal != nil {
        return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
    }

    sourceMap, err := playfab.Request(settings, b, "/MultiplayerServer/UploadCertificate", "X-EntityToken", entityToken)
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



