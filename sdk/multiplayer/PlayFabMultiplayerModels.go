package multiplayer

import "time"
                    
// AssetReference 
type AssetReferenceModel struct {
    // FileName the asset's file name. This is a filename with the .zip, .tar, or .tar.gz extension.
    FileName string `json:"FileName,omitempty"`
    // MountPath the asset's mount path.
    MountPath string `json:"MountPath,omitempty"`
}

// AssetReferenceParams 
type AssetReferenceParamsModel struct {
    // FileName the asset's file name.
    FileName string `json:"FileName,omitempty"`
    // MountPath the asset's mount path.
    MountPath string `json:"MountPath,omitempty"`
}

// AssetSummary 
type AssetSummaryModel struct {
    // FileName the asset's file name. This is a filename with the .zip, .tar, or .tar.gz extension.
    FileName string `json:"FileName,omitempty"`
    // Metadata the metadata associated with the asset.
    Metadata map[string]string `json:"Metadata,omitempty"`
}

// AzureRegion 
type AzureRegion string
  
const (
     AzureRegionAustraliaEast AzureRegion = "AustraliaEast"
     AzureRegionAustraliaSoutheast AzureRegion = "AustraliaSoutheast"
     AzureRegionBrazilSouth AzureRegion = "BrazilSouth"
     AzureRegionCentralUs AzureRegion = "CentralUs"
     AzureRegionEastAsia AzureRegion = "EastAsia"
     AzureRegionEastUs AzureRegion = "EastUs"
     AzureRegionEastUs2 AzureRegion = "EastUs2"
     AzureRegionJapanEast AzureRegion = "JapanEast"
     AzureRegionJapanWest AzureRegion = "JapanWest"
     AzureRegionNorthCentralUs AzureRegion = "NorthCentralUs"
     AzureRegionNorthEurope AzureRegion = "NorthEurope"
     AzureRegionSouthCentralUs AzureRegion = "SouthCentralUs"
     AzureRegionSoutheastAsia AzureRegion = "SoutheastAsia"
     AzureRegionWestEurope AzureRegion = "WestEurope"
     AzureRegionWestUs AzureRegion = "WestUs"
     AzureRegionChinaEast2 AzureRegion = "ChinaEast2"
     AzureRegionChinaNorth2 AzureRegion = "ChinaNorth2"
      )
// AzureVmSize 
type AzureVmSize string
  
const (
     AzureVmSizeStandard_D1_v2 AzureVmSize = "Standard_D1_v2"
     AzureVmSizeStandard_D2_v2 AzureVmSize = "Standard_D2_v2"
     AzureVmSizeStandard_D3_v2 AzureVmSize = "Standard_D3_v2"
     AzureVmSizeStandard_D4_v2 AzureVmSize = "Standard_D4_v2"
     AzureVmSizeStandard_D5_v2 AzureVmSize = "Standard_D5_v2"
     AzureVmSizeStandard_A1_v2 AzureVmSize = "Standard_A1_v2"
     AzureVmSizeStandard_A2_v2 AzureVmSize = "Standard_A2_v2"
     AzureVmSizeStandard_A4_v2 AzureVmSize = "Standard_A4_v2"
     AzureVmSizeStandard_A8_v2 AzureVmSize = "Standard_A8_v2"
     AzureVmSizeStandard_F1 AzureVmSize = "Standard_F1"
     AzureVmSizeStandard_F2 AzureVmSize = "Standard_F2"
     AzureVmSizeStandard_F4 AzureVmSize = "Standard_F4"
     AzureVmSizeStandard_F8 AzureVmSize = "Standard_F8"
     AzureVmSizeStandard_F16 AzureVmSize = "Standard_F16"
     AzureVmSizeStandard_F2s_v2 AzureVmSize = "Standard_F2s_v2"
     AzureVmSizeStandard_F4s_v2 AzureVmSize = "Standard_F4s_v2"
     AzureVmSizeStandard_F8s_v2 AzureVmSize = "Standard_F8s_v2"
     AzureVmSizeStandard_F16s_v2 AzureVmSize = "Standard_F16s_v2"
     AzureVmSizeStandard_A1 AzureVmSize = "Standard_A1"
     AzureVmSizeStandard_A2 AzureVmSize = "Standard_A2"
     AzureVmSizeStandard_A3 AzureVmSize = "Standard_A3"
     AzureVmSizeStandard_A4 AzureVmSize = "Standard_A4"
      )
// BuildRegion 
type BuildRegionModel struct {
    // CurrentServerStats the current multiplayer server stats for the region.
    CurrentServerStats CurrentServerStatsModel `json:"CurrentServerStats,omitempty"`
    // MaxServers the maximum number of multiplayer servers for the region.
    MaxServers int32 `json:"MaxServers,omitempty"`
    // Region the build region.
    Region AzureRegion `json:"Region,omitempty"`
    // StandbyServers the number of standby multiplayer servers for the region.
    StandbyServers int32 `json:"StandbyServers,omitempty"`
    // Status the status of multiplayer servers in the build region. Valid values are - Unknown, Initialized, Deploying, Deployed,
// Unhealthy.
    Status string `json:"Status,omitempty"`
}

// BuildRegionParams 
type BuildRegionParamsModel struct {
    // MaxServers the maximum number of multiplayer servers for the region.
    MaxServers int32 `json:"MaxServers,omitempty"`
    // Region the build region.
    Region AzureRegion `json:"Region,omitempty"`
    // StandbyServers the number of standby multiplayer servers for the region.
    StandbyServers int32 `json:"StandbyServers,omitempty"`
}

// BuildSummary 
type BuildSummaryModel struct {
    // BuildId the guid string build ID of the build.
    BuildId string `json:"BuildId,omitempty"`
    // BuildName the build name.
    BuildName string `json:"BuildName,omitempty"`
    // CreationTime the time the build was created in UTC.
    CreationTime time.Time `json:"CreationTime,omitempty"`
    // Metadata the metadata of the build.
    Metadata map[string]string `json:"Metadata,omitempty"`
}

// Certificate 
type CertificateModel struct {
    // Base64EncodedValue base64 encoded string contents of the certificate.
    Base64EncodedValue string `json:"Base64EncodedValue,omitempty"`
    // Name a name for the certificate. This is used to reference certificates in build configurations.
    Name string `json:"Name,omitempty"`
    // Password if required for your PFX certificate, use this field to provide a password that will be used to install the certificate
// on the container.
    Password string `json:"Password,omitempty"`
}

// CertificateSummary 
type CertificateSummaryModel struct {
    // Name the name of the certificate.
    Name string `json:"Name,omitempty"`
    // Thumbprint the thumbprint for the certificate.
    Thumbprint string `json:"Thumbprint,omitempty"`
}

// ConnectedPlayer 
type ConnectedPlayerModel struct {
    // PlayerId the player ID of the player connected to the multiplayer server.
    PlayerId string `json:"PlayerId,omitempty"`
}

// ContainerFlavor 
type ContainerFlavor string
  
const (
     ContainerFlavorManagedWindowsServerCore ContainerFlavor = "ManagedWindowsServerCore"
     ContainerFlavorCustomLinux ContainerFlavor = "CustomLinux"
     ContainerFlavorManagedWindowsServerCorePreview ContainerFlavor = "ManagedWindowsServerCorePreview"
      )
// ContainerImageReference 
type ContainerImageReferenceModel struct {
    // ImageName the container image name.
    ImageName string `json:"ImageName,omitempty"`
    // Tag the container tag.
    Tag string `json:"Tag,omitempty"`
}

// CreateBuildWithCustomContainerRequest creates a multiplayer server build with a custom container and returns information about the build creation request.
type CreateBuildWithCustomContainerRequestModel struct {
    // BuildName the build name.
    BuildName string `json:"BuildName,omitempty"`
    // ContainerFlavor the flavor of container to create a build from.
    ContainerFlavor ContainerFlavor `json:"ContainerFlavor,omitempty"`
    // ContainerRepositoryName the name of the container repository.
    ContainerRepositoryName string `json:"ContainerRepositoryName,omitempty"`
    // ContainerRunCommand the container command to run when the multiplayer server has been allocated, including any arguments.
    ContainerRunCommand string `json:"ContainerRunCommand,omitempty"`
    // ContainerTag the tag for the container.
    ContainerTag string `json:"ContainerTag,omitempty"`
    // GameAssetReferences the list of game assets related to the build.
    GameAssetReferences []AssetReferenceParamsModel `json:"GameAssetReferences,omitempty"`
    // GameCertificateReferences the game certificates for the build.
    GameCertificateReferences []GameCertificateReferenceParamsModel `json:"GameCertificateReferences,omitempty"`
    // Metadata metadata to tag the build. The keys are case insensitive. The build metadata is made available to the server through
// Game Server SDK (GSDK).
    Metadata map[string]string `json:"Metadata,omitempty"`
    // MultiplayerServerCountPerVm the number of multiplayer servers to host on a single VM.
    MultiplayerServerCountPerVm int32 `json:"MultiplayerServerCountPerVm,omitempty"`
    // Ports the ports to map the build on.
    Ports []PortModel `json:"Ports,omitempty"`
    // RegionConfigurations the region configurations for the build.
    RegionConfigurations []BuildRegionParamsModel `json:"RegionConfigurations,omitempty"`
    // VmSize the VM size to create the build on.
    VmSize AzureVmSize `json:"VmSize,omitempty"`
}

// CreateBuildWithCustomContainerResponse 
type CreateBuildWithCustomContainerResponseModel struct {
    // BuildId the guid string build ID. Must be unique for every build.
    BuildId string `json:"BuildId,omitempty"`
    // BuildName the build name.
    BuildName string `json:"BuildName,omitempty"`
    // ContainerFlavor the flavor of container of the build.
    ContainerFlavor ContainerFlavor `json:"ContainerFlavor,omitempty"`
    // ContainerRunCommand the container command to run when the multiplayer server has been allocated, including any arguments.
    ContainerRunCommand string `json:"ContainerRunCommand,omitempty"`
    // CreationTime the time the build was created in UTC.
    CreationTime time.Time `json:"CreationTime,omitempty"`
    // CustomGameContainerImage the custom game container image reference information.
    CustomGameContainerImage ContainerImageReferenceModel `json:"CustomGameContainerImage,omitempty"`
    // GameAssetReferences the game assets for the build.
    GameAssetReferences []AssetReferenceModel `json:"GameAssetReferences,omitempty"`
    // GameCertificateReferences the game certificates for the build.
    GameCertificateReferences []GameCertificateReferenceModel `json:"GameCertificateReferences,omitempty"`
    // Metadata the metadata of the build.
    Metadata map[string]string `json:"Metadata,omitempty"`
    // MultiplayerServerCountPerVm the number of multiplayer servers to host on a single VM of the build.
    MultiplayerServerCountPerVm int32 `json:"MultiplayerServerCountPerVm,omitempty"`
    // Ports the ports the build is mapped on.
    Ports []PortModel `json:"Ports,omitempty"`
    // RegionConfigurations the region configuration for the build.
    RegionConfigurations []BuildRegionModel `json:"RegionConfigurations,omitempty"`
    // VmSize the VM size the build was created on.
    VmSize AzureVmSize `json:"VmSize,omitempty"`
}

// CreateBuildWithManagedContainerRequest creates a multiplayer server build with a managed container and returns information about the build creation request.
type CreateBuildWithManagedContainerRequestModel struct {
    // BuildName the build name.
    BuildName string `json:"BuildName,omitempty"`
    // ContainerFlavor the flavor of container to create a build from.
    ContainerFlavor ContainerFlavor `json:"ContainerFlavor,omitempty"`
    // GameAssetReferences the list of game assets related to the build.
    GameAssetReferences []AssetReferenceParamsModel `json:"GameAssetReferences,omitempty"`
    // GameCertificateReferences the game certificates for the build.
    GameCertificateReferences []GameCertificateReferenceParamsModel `json:"GameCertificateReferences,omitempty"`
    // Metadata metadata to tag the build. The keys are case insensitive. The build metadata is made available to the server through
// Game Server SDK (GSDK).
    Metadata map[string]string `json:"Metadata,omitempty"`
    // MultiplayerServerCountPerVm the number of multiplayer servers to host on a single VM.
    MultiplayerServerCountPerVm int32 `json:"MultiplayerServerCountPerVm,omitempty"`
    // Ports the ports to map the build on.
    Ports []PortModel `json:"Ports,omitempty"`
    // RegionConfigurations the region configurations for the build.
    RegionConfigurations []BuildRegionParamsModel `json:"RegionConfigurations,omitempty"`
    // StartMultiplayerServerCommand the command to run when the multiplayer server is started, including any arguments.
    StartMultiplayerServerCommand string `json:"StartMultiplayerServerCommand,omitempty"`
    // VmSize the VM size to create the build on.
    VmSize AzureVmSize `json:"VmSize,omitempty"`
}

// CreateBuildWithManagedContainerResponse 
type CreateBuildWithManagedContainerResponseModel struct {
    // BuildId the guid string build ID. Must be unique for every build.
    BuildId string `json:"BuildId,omitempty"`
    // BuildName the build name.
    BuildName string `json:"BuildName,omitempty"`
    // ContainerFlavor the flavor of container of the build.
    ContainerFlavor ContainerFlavor `json:"ContainerFlavor,omitempty"`
    // CreationTime the time the build was created in UTC.
    CreationTime time.Time `json:"CreationTime,omitempty"`
    // GameAssetReferences the game assets for the build.
    GameAssetReferences []AssetReferenceModel `json:"GameAssetReferences,omitempty"`
    // GameCertificateReferences the game certificates for the build.
    GameCertificateReferences []GameCertificateReferenceModel `json:"GameCertificateReferences,omitempty"`
    // Metadata the metadata of the build.
    Metadata map[string]string `json:"Metadata,omitempty"`
    // MultiplayerServerCountPerVm the number of multiplayer servers to host on a single VM of the build.
    MultiplayerServerCountPerVm int32 `json:"MultiplayerServerCountPerVm,omitempty"`
    // Ports the ports the build is mapped on.
    Ports []PortModel `json:"Ports,omitempty"`
    // RegionConfigurations the region configuration for the build.
    RegionConfigurations []BuildRegionModel `json:"RegionConfigurations,omitempty"`
    // StartMultiplayerServerCommand the command to run when the multiplayer server has been allocated, including any arguments.
    StartMultiplayerServerCommand string `json:"StartMultiplayerServerCommand,omitempty"`
    // VmSize the VM size the build was created on.
    VmSize AzureVmSize `json:"VmSize,omitempty"`
}

// CreateRemoteUserRequest creates a remote user to log on to a VM for a multiplayer server build in a specific region. Returns user credential
// information necessary to log on.
type CreateRemoteUserRequestModel struct {
    // BuildId the guid string build ID of to create the remote user for.
    BuildId string `json:"BuildId,omitempty"`
    // ExpirationTime the expiration time for the remote user created. Defaults to expiring in one day if not specified.
    ExpirationTime time.Time `json:"ExpirationTime,omitempty"`
    // Region the region of virtual machine to create the remote user for.
    Region AzureRegion `json:"Region,omitempty"`
    // Username the username to create the remote user with.
    Username string `json:"Username,omitempty"`
    // VmId the virtual machine ID the multiplayer server is located on.
    VmId string `json:"VmId,omitempty"`
}

// CreateRemoteUserResponse 
type CreateRemoteUserResponseModel struct {
    // ExpirationTime the expiration time for the remote user created.
    ExpirationTime time.Time `json:"ExpirationTime,omitempty"`
    // Password the generated password for the remote user that was created.
    Password string `json:"Password,omitempty"`
    // Username the username for the remote user that was created.
    Username string `json:"Username,omitempty"`
}

// CurrentServerStats 
type CurrentServerStatsModel struct {
    // Active the number of active multiplayer servers.
    Active int32 `json:"Active,omitempty"`
    // Propping the number of multiplayer servers still downloading game resources (such as assets).
    Propping int32 `json:"Propping,omitempty"`
    // StandingBy the number of standingby multiplayer servers.
    StandingBy int32 `json:"StandingBy,omitempty"`
    // Total the total number of multiplayer servers.
    Total int32 `json:"Total,omitempty"`
}

// DeleteAssetRequest deletes a multiplayer server game asset for a title.
type DeleteAssetRequestModel struct {
    // FileName the filename of the asset to delete.
    FileName string `json:"FileName,omitempty"`
}

// DeleteBuildRequest deletes a multiplayer server build.
type DeleteBuildRequestModel struct {
    // BuildId the guid string build ID of the build to delete.
    BuildId string `json:"BuildId,omitempty"`
}

// DeleteCertificateRequest deletes a multiplayer server game certificate.
type DeleteCertificateRequestModel struct {
    // Name the name of the certificate.
    Name string `json:"Name,omitempty"`
}

// DeleteRemoteUserRequest deletes a remote user to log on to a VM for a multiplayer server build in a specific region. Returns user credential
// information necessary to log on.
type DeleteRemoteUserRequestModel struct {
    // BuildId the guid string build ID of the multiplayer server where the remote user is to delete.
    BuildId string `json:"BuildId,omitempty"`
    // Region the region of the multiplayer server where the remote user is to delete.
    Region AzureRegion `json:"Region,omitempty"`
    // Username the username of the remote user to delete.
    Username string `json:"Username,omitempty"`
    // VmId the virtual machine ID the multiplayer server is located on.
    VmId string `json:"VmId,omitempty"`
}

// EmptyResponse 
type EmptyResponseModel struct {
}

// EnableMultiplayerServersForTitleRequest enables the multiplayer server feature for a title and returns the enabled status. The enabled status can be
// Initializing, Enabled, and Disabled. It can up to 20 minutes or more for the title to be enabled for the feature. On
// average, it can take up to 20 minutes for the title to be enabled for the feature.
type EnableMultiplayerServersForTitleRequestModel struct {
}

// EnableMultiplayerServersForTitleResponse 
type EnableMultiplayerServersForTitleResponseModel struct {
    // Status the enabled status for the multiplayer server features for the title.
    Status TitleMultiplayerServerEnabledStatus `json:"Status,omitempty"`
}

// GameCertificateReference 
type GameCertificateReferenceModel struct {
    // GsdkAlias an alias for the game certificate. The game server will reference this alias via GSDK config to retrieve the game
// certificate. This alias is used as an identifier in game server code to allow a new certificate with different Name
// field to be uploaded without the need to change any game server code to reference the new Name.
    GsdkAlias string `json:"GsdkAlias,omitempty"`
    // Name the name of the game certificate. This name should match the name of a certificate that was previously uploaded to this
// title.
    Name string `json:"Name,omitempty"`
}

// GameCertificateReferenceParams 
type GameCertificateReferenceParamsModel struct {
    // GsdkAlias an alias for the game certificate. The game server will reference this alias via GSDK config to retrieve the game
// certificate. This alias is used as an identifier in game server code to allow a new certificate with different Name
// field to be uploaded without the need to change any game server code to reference the new Name.
    GsdkAlias string `json:"GsdkAlias,omitempty"`
    // Name the name of the game certificate. This name should match the name of a certificate that was previously uploaded to this
// title.
    Name string `json:"Name,omitempty"`
}

// GetAssetUploadUrlRequest gets the URL to upload assets to.
type GetAssetUploadUrlRequestModel struct {
    // FileName the asset's file name to get the upload URL for.
    FileName string `json:"FileName,omitempty"`
}

// GetAssetUploadUrlResponse 
type GetAssetUploadUrlResponseModel struct {
    // AssetUploadUrl the asset's upload URL.
    AssetUploadUrl string `json:"AssetUploadUrl,omitempty"`
    // FileName the asset's file name to get the upload URL for.
    FileName string `json:"FileName,omitempty"`
}

// GetBuildRequest returns the details about a multiplayer server build.
type GetBuildRequestModel struct {
    // BuildId the guid string build ID of the build to get.
    BuildId string `json:"BuildId,omitempty"`
}

// GetBuildResponse 
type GetBuildResponseModel struct {
    // BuildId the guid string build ID of the build.
    BuildId string `json:"BuildId,omitempty"`
    // BuildName the build name.
    BuildName string `json:"BuildName,omitempty"`
    // BuildStatus the current build status. Valid values are - Deploying, Deployed, DeletingRegion, Unhealthy.
    BuildStatus string `json:"BuildStatus,omitempty"`
    // ContainerFlavor the flavor of container of he build.
    ContainerFlavor ContainerFlavor `json:"ContainerFlavor,omitempty"`
    // ContainerRunCommand the container command to run when the multiplayer server has been allocated, including any arguments. This only applies
// to custom builds. If the build is a managed build, this field will be null.
    ContainerRunCommand string `json:"ContainerRunCommand,omitempty"`
    // CreationTime the time the build was created in UTC.
    CreationTime time.Time `json:"CreationTime,omitempty"`
    // CustomGameContainerImage the custom game container image for a custom build.
    CustomGameContainerImage ContainerImageReferenceModel `json:"CustomGameContainerImage,omitempty"`
    // GameAssetReferences the game assets for the build.
    GameAssetReferences []AssetReferenceModel `json:"GameAssetReferences,omitempty"`
    // GameCertificateReferences the game certificates for the build.
    GameCertificateReferences []GameCertificateReferenceModel `json:"GameCertificateReferences,omitempty"`
    // Metadata metadata of the build. The keys are case insensitive. The build metadata is made available to the server through Game
// Server SDK (GSDK).
    Metadata map[string]string `json:"Metadata,omitempty"`
    // MultiplayerServerCountPerVm the number of multiplayer servers to hosted on a single VM of the build.
    MultiplayerServerCountPerVm int32 `json:"MultiplayerServerCountPerVm,omitempty"`
    // Ports the ports the build is mapped on.
    Ports []PortModel `json:"Ports,omitempty"`
    // RegionConfigurations the region configuration for the build.
    RegionConfigurations []BuildRegionModel `json:"RegionConfigurations,omitempty"`
    // StartMultiplayerServerCommand the command to run when the multiplayer server has been allocated, including any arguments. This only applies to managed
// builds. If the build is a custom build, this field will be null.
    StartMultiplayerServerCommand string `json:"StartMultiplayerServerCommand,omitempty"`
    // VmSize the VM size the build was created on.
    VmSize AzureVmSize `json:"VmSize,omitempty"`
}

// GetContainerRegistryCredentialsRequest gets credentials to the container registry where game developers can upload custom container images to before creating a
// new build.
type GetContainerRegistryCredentialsRequestModel struct {
}

// GetContainerRegistryCredentialsResponse 
type GetContainerRegistryCredentialsResponseModel struct {
    // DnsName the url of the container registry.
    DnsName string `json:"DnsName,omitempty"`
    // Password the password for accessing the container registry.
    Password string `json:"Password,omitempty"`
    // Username the username for accessing the container registry.
    Username string `json:"Username,omitempty"`
}

// GetMultiplayerServerDetailsRequest gets multiplayer server session details for a build in a specific region.
type GetMultiplayerServerDetailsRequestModel struct {
    // BuildId the guid string build ID of the multiplayer server to get details for.
    BuildId string `json:"BuildId,omitempty"`
    // Region the region the multiplayer server is located in to get details for.
    Region AzureRegion `json:"Region,omitempty"`
    // SessionId the title generated guid string session ID of the multiplayer server to get details for. This is to keep track of
// multiplayer server sessions.
    SessionId string `json:"SessionId,omitempty"`
}

// GetMultiplayerServerDetailsResponse 
type GetMultiplayerServerDetailsResponseModel struct {
    // ConnectedPlayers the connected players in the multiplayer server.
    ConnectedPlayers []ConnectedPlayerModel `json:"ConnectedPlayers,omitempty"`
    // FQDN the fully qualified domain name of the virtual machine that is hosting this multiplayer server.
    FQDN string `json:"FQDN,omitempty"`
    // IPV4Address the IPv4 address of the virtual machine that is hosting this multiplayer server.
    IPV4Address string `json:"IPV4Address,omitempty"`
    // LastStateTransitionTime the time (UTC) at which a change in the multiplayer server state was observed.
    LastStateTransitionTime time.Time `json:"LastStateTransitionTime,omitempty"`
    // Ports the ports the multiplayer server uses.
    Ports []PortModel `json:"Ports,omitempty"`
    // Region the region the multiplayer server is located in.
    Region AzureRegion `json:"Region,omitempty"`
    // ServerId the string server ID of the multiplayer server generated by PlayFab.
    ServerId string `json:"ServerId,omitempty"`
    // SessionId the guid string session ID of the multiplayer server.
    SessionId string `json:"SessionId,omitempty"`
    // State the state of the multiplayer server.
    State string `json:"State,omitempty"`
    // VmId the virtual machine ID that the multiplayer server is located on.
    VmId string `json:"VmId,omitempty"`
}

// GetRemoteLoginEndpointRequest gets a remote login endpoint to a VM that is hosting a multiplayer server build in a specific region.
type GetRemoteLoginEndpointRequestModel struct {
    // BuildId the guid string build ID of the multiplayer server to get remote login information for.
    BuildId string `json:"BuildId,omitempty"`
    // Region the region of the multiplayer server to get remote login information for.
    Region AzureRegion `json:"Region,omitempty"`
    // VmId the virtual machine ID the multiplayer server is located on.
    VmId string `json:"VmId,omitempty"`
}

// GetRemoteLoginEndpointResponse 
type GetRemoteLoginEndpointResponseModel struct {
    // IPV4Address the remote login IPV4 address of multiplayer server.
    IPV4Address string `json:"IPV4Address,omitempty"`
    // Port the remote login port of multiplayer server.
    Port int32 `json:"Port,omitempty"`
}

// GetTitleEnabledForMultiplayerServersStatusRequest gets the status of whether a title is enabled for the multiplayer server feature. The enabled status can be
// Initializing, Enabled, and Disabled.
type GetTitleEnabledForMultiplayerServersStatusRequestModel struct {
}

// GetTitleEnabledForMultiplayerServersStatusResponse 
type GetTitleEnabledForMultiplayerServersStatusResponseModel struct {
    // Status the enabled status for the multiplayer server features for the title.
    Status TitleMultiplayerServerEnabledStatus `json:"Status,omitempty"`
}

// ListAssetSummariesRequest returns a list of multiplayer server game asset summaries for a title.
type ListAssetSummariesRequestModel struct {
    // PageSize the page size for the request.
    PageSize int32 `json:"PageSize,omitempty"`
    // SkipToken the skip token for the paged request.
    SkipToken string `json:"SkipToken,omitempty"`
}

// ListAssetSummariesResponse 
type ListAssetSummariesResponseModel struct {
    // AssetSummaries the list of asset summaries.
    AssetSummaries []AssetSummaryModel `json:"AssetSummaries,omitempty"`
    // PageSize the page size on the response.
    PageSize int32 `json:"PageSize,omitempty"`
    // SkipToken the skip token for the paged response.
    SkipToken string `json:"SkipToken,omitempty"`
}

// ListBuildSummariesRequest returns a list of summarized details of all multiplayer server builds for a title.
type ListBuildSummariesRequestModel struct {
    // PageSize the page size for the request.
    PageSize int32 `json:"PageSize,omitempty"`
    // SkipToken the skip token for the paged request.
    SkipToken string `json:"SkipToken,omitempty"`
}

// ListBuildSummariesResponse 
type ListBuildSummariesResponseModel struct {
    // BuildSummaries the list of build summaries for a title.
    BuildSummaries []BuildSummaryModel `json:"BuildSummaries,omitempty"`
    // PageSize the page size on the response.
    PageSize int32 `json:"PageSize,omitempty"`
    // SkipToken the skip token for the paged response.
    SkipToken string `json:"SkipToken,omitempty"`
}

// ListCertificateSummariesRequest returns a list of multiplayer server game certificates for a title.
type ListCertificateSummariesRequestModel struct {
    // PageSize the page size for the request.
    PageSize int32 `json:"PageSize,omitempty"`
    // SkipToken the skip token for the paged request.
    SkipToken string `json:"SkipToken,omitempty"`
}

// ListCertificateSummariesResponse 
type ListCertificateSummariesResponseModel struct {
    // CertificateSummaries the list of game certificates.
    CertificateSummaries []CertificateSummaryModel `json:"CertificateSummaries,omitempty"`
    // PageSize the page size on the response.
    PageSize int32 `json:"PageSize,omitempty"`
    // SkipToken the skip token for the paged response.
    SkipToken string `json:"SkipToken,omitempty"`
}

// ListContainerImagesRequest returns a list of the container images that have been uploaded to the container registry for a title.
type ListContainerImagesRequestModel struct {
    // PageSize the page size for the request.
    PageSize int32 `json:"PageSize,omitempty"`
    // SkipToken the skip token for the paged request.
    SkipToken string `json:"SkipToken,omitempty"`
}

// ListContainerImagesResponse 
type ListContainerImagesResponseModel struct {
    // Images the list of container images.
    Images []string `json:"Images,omitempty"`
    // PageSize the page size on the response.
    PageSize int32 `json:"PageSize,omitempty"`
    // SkipToken the skip token for the paged response.
    SkipToken string `json:"SkipToken,omitempty"`
}

// ListContainerImageTagsRequest returns a list of the tags for a particular container image that exists in the container registry for a title.
type ListContainerImageTagsRequestModel struct {
    // ImageName the container images we want to list tags for.
    ImageName string `json:"ImageName,omitempty"`
}

// ListContainerImageTagsResponse 
type ListContainerImageTagsResponseModel struct {
    // Tags the list of tags for a particular container image.
    Tags []string `json:"Tags,omitempty"`
}

// ListMultiplayerServersRequest returns a list of multiplayer servers for a build in a specific region.
type ListMultiplayerServersRequestModel struct {
    // BuildId the guid string build ID of the multiplayer servers to list.
    BuildId string `json:"BuildId,omitempty"`
    // PageSize the page size for the request.
    PageSize int32 `json:"PageSize,omitempty"`
    // Region the region the multiplayer servers to list.
    Region AzureRegion `json:"Region,omitempty"`
    // SkipToken the skip token for the paged request.
    SkipToken string `json:"SkipToken,omitempty"`
}

// ListMultiplayerServersResponse 
type ListMultiplayerServersResponseModel struct {
    // MultiplayerServerSummaries the list of multiplayer server summary details.
    MultiplayerServerSummaries []MultiplayerServerSummaryModel `json:"MultiplayerServerSummaries,omitempty"`
    // PageSize the page size on the response.
    PageSize int32 `json:"PageSize,omitempty"`
    // SkipToken the skip token for the paged response.
    SkipToken string `json:"SkipToken,omitempty"`
}

// ListQosServersRequest returns a list of quality of service servers.
type ListQosServersRequestModel struct {
}

// ListQosServersResponse 
type ListQosServersResponseModel struct {
    // PageSize the page size on the response.
    PageSize int32 `json:"PageSize,omitempty"`
    // QosServers the list of QoS servers.
    QosServers []QosServerModel `json:"QosServers,omitempty"`
    // SkipToken the skip token for the paged response.
    SkipToken string `json:"SkipToken,omitempty"`
}

// ListVirtualMachineSummariesRequest returns a list of virtual machines for a title.
type ListVirtualMachineSummariesRequestModel struct {
    // BuildId the guid string build ID of the virtual machines to list.
    BuildId string `json:"BuildId,omitempty"`
    // PageSize the page size for the request.
    PageSize int32 `json:"PageSize,omitempty"`
    // Region the region of the virtual machines to list.
    Region AzureRegion `json:"Region,omitempty"`
    // SkipToken the skip token for the paged request.
    SkipToken string `json:"SkipToken,omitempty"`
}

// ListVirtualMachineSummariesResponse 
type ListVirtualMachineSummariesResponseModel struct {
    // PageSize the page size on the response.
    PageSize int32 `json:"PageSize,omitempty"`
    // SkipToken the skip token for the paged response.
    SkipToken string `json:"SkipToken,omitempty"`
    // VirtualMachines the list of virtual machine summaries.
    VirtualMachines []VirtualMachineSummaryModel `json:"VirtualMachines,omitempty"`
}

// MultiplayerServerSummary 
type MultiplayerServerSummaryModel struct {
    // ConnectedPlayers the connected players in the multiplayer server.
    ConnectedPlayers []ConnectedPlayerModel `json:"ConnectedPlayers,omitempty"`
    // LastStateTransitionTime the time (UTC) at which a change in the multiplayer server state was observed.
    LastStateTransitionTime time.Time `json:"LastStateTransitionTime,omitempty"`
    // Region the region the multiplayer server is located in.
    Region AzureRegion `json:"Region,omitempty"`
    // ServerId the string server ID of the multiplayer server generated by PlayFab.
    ServerId string `json:"ServerId,omitempty"`
    // SessionId the title generated guid string session ID of the multiplayer server.
    SessionId string `json:"SessionId,omitempty"`
    // State the state of the multiplayer server.
    State string `json:"State,omitempty"`
    // VmId the virtual machine ID that the multiplayer server is located on.
    VmId string `json:"VmId,omitempty"`
}

// Port 
type PortModel struct {
    // Name the name for the port.
    Name string `json:"Name,omitempty"`
    // Num the number for the port.
    Num int32 `json:"Num,omitempty"`
    // Protocol the protocol for the port.
    Protocol ProtocolType `json:"Protocol,omitempty"`
}

// ProtocolType 
type ProtocolType string
  
const (
     ProtocolTypeTCP ProtocolType = "TCP"
     ProtocolTypeUDP ProtocolType = "UDP"
      )
// QosServer 
type QosServerModel struct {
    // Region the region the QoS server is located in.
    Region AzureRegion `json:"Region,omitempty"`
    // ServerUrl the QoS server URL.
    ServerUrl string `json:"ServerUrl,omitempty"`
}

// RequestMultiplayerServerRequest requests a multiplayer server session from a particular build in any of the given preferred regions.
type RequestMultiplayerServerRequestModel struct {
    // BuildId the guid string build ID of the multiplayer server to request.
    BuildId string `json:"BuildId,omitempty"`
    // InitialPlayers initial list of players (potentially matchmade) allowed to connect to the game. This list is passed to the game server
// when requested (via GSDK) and can be used to validate players connecting to it.
    InitialPlayers []string `json:"InitialPlayers,omitempty"`
    // PreferredRegions the preferred regions to request a multiplayer server from. The Multiplayer Service will iterate through the regions in
// the specified order and allocate a server from the first one that has servers available.
    PreferredRegions []AzureRegion `json:"PreferredRegions,omitempty"`
    // SessionCookie data encoded as a string that is passed to the game server when requested. This can be used to to communicate
// information such as game mode or map through the request flow.
    SessionCookie string `json:"SessionCookie,omitempty"`
    // SessionId a guid string session ID created track the multiplayer server session over its life.
    SessionId string `json:"SessionId,omitempty"`
}

// RequestMultiplayerServerResponse 
type RequestMultiplayerServerResponseModel struct {
    // ConnectedPlayers the connected players in the multiplayer server.
    ConnectedPlayers []ConnectedPlayerModel `json:"ConnectedPlayers,omitempty"`
    // FQDN the fully qualified domain name of the virtual machine that is hosting this multiplayer server.
    FQDN string `json:"FQDN,omitempty"`
    // IPV4Address the IPv4 address of the virtual machine that is hosting this multiplayer server.
    IPV4Address string `json:"IPV4Address,omitempty"`
    // LastStateTransitionTime the time (UTC) at which a change in the multiplayer server state was observed.
    LastStateTransitionTime time.Time `json:"LastStateTransitionTime,omitempty"`
    // Ports the ports the multiplayer server uses.
    Ports []PortModel `json:"Ports,omitempty"`
    // Region the region the multiplayer server is located in.
    Region AzureRegion `json:"Region,omitempty"`
    // ServerId the string server ID of the multiplayer server generated by PlayFab.
    ServerId string `json:"ServerId,omitempty"`
    // SessionId the guid string session ID of the multiplayer server.
    SessionId string `json:"SessionId,omitempty"`
    // State the state of the multiplayer server.
    State string `json:"State,omitempty"`
    // VmId the virtual machine ID that the multiplayer server is located on.
    VmId string `json:"VmId,omitempty"`
}

// RolloverContainerRegistryCredentialsRequest gets new credentials to the container registry where game developers can upload custom container images to before
// creating a new build.
type RolloverContainerRegistryCredentialsRequestModel struct {
}

// RolloverContainerRegistryCredentialsResponse 
type RolloverContainerRegistryCredentialsResponseModel struct {
    // DnsName the url of the container registry.
    DnsName string `json:"DnsName,omitempty"`
    // Password the password for accessing the container registry.
    Password string `json:"Password,omitempty"`
    // Username the username for accessing the container registry.
    Username string `json:"Username,omitempty"`
}

// ShutdownMultiplayerServerRequest executes the shutdown callback from the GSDK and terminates the multiplayer server session. The callback in the GSDK
// will allow for graceful shutdown with a 15 minute timeoutIf graceful shutdown has not been completed before 15 minutes
// have elapsed, the multiplayer server session will be forcefully terminated on it's own.
type ShutdownMultiplayerServerRequestModel struct {
    // BuildId the guid string build ID of the multiplayer server to delete.
    BuildId string `json:"BuildId,omitempty"`
    // Region the region of the multiplayer server to shut down.
    Region AzureRegion `json:"Region,omitempty"`
    // SessionId a guid string session ID of the multiplayer server to shut down.
    SessionId string `json:"SessionId,omitempty"`
}

// TitleMultiplayerServerEnabledStatus 
type TitleMultiplayerServerEnabledStatus string
  
const (
     TitleMultiplayerServerEnabledStatusInitializing TitleMultiplayerServerEnabledStatus = "Initializing"
     TitleMultiplayerServerEnabledStatusEnabled TitleMultiplayerServerEnabledStatus = "Enabled"
     TitleMultiplayerServerEnabledStatusDisabled TitleMultiplayerServerEnabledStatus = "Disabled"
      )
// UpdateBuildRegionsRequest updates a multiplayer server build's regions.
type UpdateBuildRegionsRequestModel struct {
    // BuildId the guid string ID of the build we want to update regions for.
    BuildId string `json:"BuildId,omitempty"`
    // BuildRegions the updated region configuration that should be applied to the specified build.
    BuildRegions []BuildRegionParamsModel `json:"BuildRegions,omitempty"`
}

// UploadCertificateRequest uploads a multiplayer server game certificate.
type UploadCertificateRequestModel struct {
    // GameCertificate the game certificate to upload.
    GameCertificate CertificateModel `json:"GameCertificate,omitempty"`
}

// VirtualMachineSummary 
type VirtualMachineSummaryModel struct {
    // HealthStatus the virtual machine health status.
    HealthStatus string `json:"HealthStatus,omitempty"`
    // State the virtual machine state.
    State string `json:"State,omitempty"`
    // VmId the virtual machine ID.
    VmId string `json:"VmId,omitempty"`
}
