package multiplayer

import "time"
                    
// AccessPolicy 
type AccessPolicy string
  
const (
     AccessPolicyPublic AccessPolicy = "Public"
     AccessPolicyFriends AccessPolicy = "Friends"
     AccessPolicyPrivate AccessPolicy = "Private"
      )
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
     AzureRegionSouthAfricaNorth AzureRegion = "SouthAfricaNorth"
     AzureRegionWestCentralUs AzureRegion = "WestCentralUs"
     AzureRegionKoreaCentral AzureRegion = "KoreaCentral"
     AzureRegionFranceCentral AzureRegion = "FranceCentral"
     AzureRegionWestUs2 AzureRegion = "WestUs2"
     AzureRegionCentralIndia AzureRegion = "CentralIndia"
     AzureRegionUaeNorth AzureRegion = "UaeNorth"
     AzureRegionUkSouth AzureRegion = "UkSouth"
      )
// AzureVmFamily 
type AzureVmFamily string
  
const (
     AzureVmFamilyA AzureVmFamily = "A"
     AzureVmFamilyAv2 AzureVmFamily = "Av2"
     AzureVmFamilyDv2 AzureVmFamily = "Dv2"
     AzureVmFamilyDv3 AzureVmFamily = "Dv3"
     AzureVmFamilyF AzureVmFamily = "F"
     AzureVmFamilyFsv2 AzureVmFamily = "Fsv2"
     AzureVmFamilyDasv4 AzureVmFamily = "Dasv4"
     AzureVmFamilyDav4 AzureVmFamily = "Dav4"
     AzureVmFamilyEav4 AzureVmFamily = "Eav4"
     AzureVmFamilyEasv4 AzureVmFamily = "Easv4"
     AzureVmFamilyEv4 AzureVmFamily = "Ev4"
     AzureVmFamilyEsv4 AzureVmFamily = "Esv4"
     AzureVmFamilyDsv3 AzureVmFamily = "Dsv3"
     AzureVmFamilyDsv2 AzureVmFamily = "Dsv2"
     AzureVmFamilyNCasT4_v3 AzureVmFamily = "NCasT4_v3"
     AzureVmFamilyDdv4 AzureVmFamily = "Ddv4"
     AzureVmFamilyDdsv4 AzureVmFamily = "Ddsv4"
     AzureVmFamilyHBv3 AzureVmFamily = "HBv3"
      )
// AzureVmSize 
type AzureVmSize string
  
const (
     AzureVmSizeStandard_A1 AzureVmSize = "Standard_A1"
     AzureVmSizeStandard_A2 AzureVmSize = "Standard_A2"
     AzureVmSizeStandard_A3 AzureVmSize = "Standard_A3"
     AzureVmSizeStandard_A4 AzureVmSize = "Standard_A4"
     AzureVmSizeStandard_A1_v2 AzureVmSize = "Standard_A1_v2"
     AzureVmSizeStandard_A2_v2 AzureVmSize = "Standard_A2_v2"
     AzureVmSizeStandard_A4_v2 AzureVmSize = "Standard_A4_v2"
     AzureVmSizeStandard_A8_v2 AzureVmSize = "Standard_A8_v2"
     AzureVmSizeStandard_D1_v2 AzureVmSize = "Standard_D1_v2"
     AzureVmSizeStandard_D2_v2 AzureVmSize = "Standard_D2_v2"
     AzureVmSizeStandard_D3_v2 AzureVmSize = "Standard_D3_v2"
     AzureVmSizeStandard_D4_v2 AzureVmSize = "Standard_D4_v2"
     AzureVmSizeStandard_D5_v2 AzureVmSize = "Standard_D5_v2"
     AzureVmSizeStandard_D2_v3 AzureVmSize = "Standard_D2_v3"
     AzureVmSizeStandard_D4_v3 AzureVmSize = "Standard_D4_v3"
     AzureVmSizeStandard_D8_v3 AzureVmSize = "Standard_D8_v3"
     AzureVmSizeStandard_D16_v3 AzureVmSize = "Standard_D16_v3"
     AzureVmSizeStandard_F1 AzureVmSize = "Standard_F1"
     AzureVmSizeStandard_F2 AzureVmSize = "Standard_F2"
     AzureVmSizeStandard_F4 AzureVmSize = "Standard_F4"
     AzureVmSizeStandard_F8 AzureVmSize = "Standard_F8"
     AzureVmSizeStandard_F16 AzureVmSize = "Standard_F16"
     AzureVmSizeStandard_F2s_v2 AzureVmSize = "Standard_F2s_v2"
     AzureVmSizeStandard_F4s_v2 AzureVmSize = "Standard_F4s_v2"
     AzureVmSizeStandard_F8s_v2 AzureVmSize = "Standard_F8s_v2"
     AzureVmSizeStandard_F16s_v2 AzureVmSize = "Standard_F16s_v2"
     AzureVmSizeStandard_D2as_v4 AzureVmSize = "Standard_D2as_v4"
     AzureVmSizeStandard_D4as_v4 AzureVmSize = "Standard_D4as_v4"
     AzureVmSizeStandard_D8as_v4 AzureVmSize = "Standard_D8as_v4"
     AzureVmSizeStandard_D16as_v4 AzureVmSize = "Standard_D16as_v4"
     AzureVmSizeStandard_D2a_v4 AzureVmSize = "Standard_D2a_v4"
     AzureVmSizeStandard_D4a_v4 AzureVmSize = "Standard_D4a_v4"
     AzureVmSizeStandard_D8a_v4 AzureVmSize = "Standard_D8a_v4"
     AzureVmSizeStandard_D16a_v4 AzureVmSize = "Standard_D16a_v4"
     AzureVmSizeStandard_E2a_v4 AzureVmSize = "Standard_E2a_v4"
     AzureVmSizeStandard_E4a_v4 AzureVmSize = "Standard_E4a_v4"
     AzureVmSizeStandard_E8a_v4 AzureVmSize = "Standard_E8a_v4"
     AzureVmSizeStandard_E16a_v4 AzureVmSize = "Standard_E16a_v4"
     AzureVmSizeStandard_E2as_v4 AzureVmSize = "Standard_E2as_v4"
     AzureVmSizeStandard_E4as_v4 AzureVmSize = "Standard_E4as_v4"
     AzureVmSizeStandard_E8as_v4 AzureVmSize = "Standard_E8as_v4"
     AzureVmSizeStandard_E16as_v4 AzureVmSize = "Standard_E16as_v4"
     AzureVmSizeStandard_D2s_v3 AzureVmSize = "Standard_D2s_v3"
     AzureVmSizeStandard_D4s_v3 AzureVmSize = "Standard_D4s_v3"
     AzureVmSizeStandard_D8s_v3 AzureVmSize = "Standard_D8s_v3"
     AzureVmSizeStandard_D16s_v3 AzureVmSize = "Standard_D16s_v3"
     AzureVmSizeStandard_DS1_v2 AzureVmSize = "Standard_DS1_v2"
     AzureVmSizeStandard_DS2_v2 AzureVmSize = "Standard_DS2_v2"
     AzureVmSizeStandard_DS3_v2 AzureVmSize = "Standard_DS3_v2"
     AzureVmSizeStandard_DS4_v2 AzureVmSize = "Standard_DS4_v2"
     AzureVmSizeStandard_DS5_v2 AzureVmSize = "Standard_DS5_v2"
     AzureVmSizeStandard_NC4as_T4_v3 AzureVmSize = "Standard_NC4as_T4_v3"
     AzureVmSizeStandard_D2d_v4 AzureVmSize = "Standard_D2d_v4"
     AzureVmSizeStandard_D4d_v4 AzureVmSize = "Standard_D4d_v4"
     AzureVmSizeStandard_D8d_v4 AzureVmSize = "Standard_D8d_v4"
     AzureVmSizeStandard_D16d_v4 AzureVmSize = "Standard_D16d_v4"
     AzureVmSizeStandard_D2ds_v4 AzureVmSize = "Standard_D2ds_v4"
     AzureVmSizeStandard_D4ds_v4 AzureVmSize = "Standard_D4ds_v4"
     AzureVmSizeStandard_D8ds_v4 AzureVmSize = "Standard_D8ds_v4"
     AzureVmSizeStandard_D16ds_v4 AzureVmSize = "Standard_D16ds_v4"
     AzureVmSizeStandard_HB120_16rs_v3 AzureVmSize = "Standard_HB120_16rs_v3"
     AzureVmSizeStandard_HB120_32rs_v3 AzureVmSize = "Standard_HB120_32rs_v3"
     AzureVmSizeStandard_HB120_64rs_v3 AzureVmSize = "Standard_HB120_64rs_v3"
     AzureVmSizeStandard_HB120_96rs_v3 AzureVmSize = "Standard_HB120_96rs_v3"
     AzureVmSizeStandard_HB120rs_v3 AzureVmSize = "Standard_HB120rs_v3"
      )
// BuildAliasDetailsResponse 
type BuildAliasDetailsResponseModel struct {
    // AliasId the guid string alias Id of the alias to be created or updated.
    AliasId string `json:"AliasId,omitempty"`
    // AliasName the alias name.
    AliasName string `json:"AliasName,omitempty"`
    // BuildSelectionCriteria array of build selection criteria.
    BuildSelectionCriteria []BuildSelectionCriterionModel `json:"BuildSelectionCriteria,omitempty"`
}

// BuildAliasParams 
type BuildAliasParamsModel struct {
    // AliasId the guid string alias ID to use for the request.
    AliasId string `json:"AliasId,omitempty"`
}

// BuildRegion 
type BuildRegionModel struct {
    // CurrentServerStats the current multiplayer server stats for the region.
    CurrentServerStats *CurrentServerStatsModel `json:"CurrentServerStats,omitempty"`
    // DynamicStandbySettings optional settings to control dynamic adjustment of standby target
    DynamicStandbySettings *DynamicStandbySettingsModel `json:"DynamicStandbySettings,omitempty"`
    // IsAssetReplicationComplete whether the game assets provided for the build have been replicated to this region.
    IsAssetReplicationComplete bool `json:"IsAssetReplicationComplete"`
    // MaxServers the maximum number of multiplayer servers for the region.
    MaxServers int32 `json:"MaxServers,omitempty"`
    // MultiplayerServerCountPerVm regional override for the number of multiplayer servers to host on a single VM of the build.
    MultiplayerServerCountPerVm int32 `json:"MultiplayerServerCountPerVm,omitempty"`
    // Region the build region.
    Region string `json:"Region,omitempty"`
    // ScheduledStandbySettings optional settings to set the standby target to specified values during the supplied schedules
    ScheduledStandbySettings *ScheduledStandbySettingsModel `json:"ScheduledStandbySettings,omitempty"`
    // StandbyServers the target number of standby multiplayer servers for the region.
    StandbyServers int32 `json:"StandbyServers,omitempty"`
    // Status the status of multiplayer servers in the build region. Valid values are - Unknown, Initialized, Deploying, Deployed,
// Unhealthy, Deleting, Deleted.
    Status string `json:"Status,omitempty"`
    // VmSize regional override for the VM size the build was created on.
    VmSize AzureVmSize `json:"VmSize,omitempty"`
}

// BuildRegionParams 
type BuildRegionParamsModel struct {
    // DynamicStandbySettings optional settings to control dynamic adjustment of standby target. If not specified, dynamic standby is disabled
    DynamicStandbySettings *DynamicStandbySettingsModel `json:"DynamicStandbySettings,omitempty"`
    // MaxServers the maximum number of multiplayer servers for the region.
    MaxServers int32 `json:"MaxServers,omitempty"`
    // MultiplayerServerCountPerVm regional override for the number of multiplayer servers to host on a single VM of the build.
    MultiplayerServerCountPerVm int32 `json:"MultiplayerServerCountPerVm,omitempty"`
    // Region the build region.
    Region string `json:"Region,omitempty"`
    // ScheduledStandbySettings optional settings to set the standby target to specified values during the supplied schedules
    ScheduledStandbySettings *ScheduledStandbySettingsModel `json:"ScheduledStandbySettings,omitempty"`
    // StandbyServers the number of standby multiplayer servers for the region.
    StandbyServers int32 `json:"StandbyServers,omitempty"`
    // VmSize regional override for the VM size the build was created on.
    VmSize AzureVmSize `json:"VmSize,omitempty"`
}

// BuildSelectionCriterion 
type BuildSelectionCriterionModel struct {
    // BuildWeightDistribution dictionary of build ids and their respective weights for distribution of allocation requests.
    BuildWeightDistribution map[string]uint32 `json:"BuildWeightDistribution,omitempty"`
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
    // RegionConfigurations the configuration and status for each region in the build.
    RegionConfigurations []BuildRegionModel `json:"RegionConfigurations,omitempty"`
}

// CancelAllMatchmakingTicketsForPlayerRequest cancels all tickets of which the player is a member in a given queue that are not cancelled or matched. This API is
// useful if you lose track of what tickets the player is a member of (if the title crashes for instance) and want to
// "reset". The Entity field is optional if the caller is a player and defaults to that player. Players may not cancel
// tickets for other people. The Entity field is required if the caller is a server (authenticated as the title).
type CancelAllMatchmakingTicketsForPlayerRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity key of the player whose tickets should be canceled.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // QueueName the name of the queue from which a player's tickets should be canceled.
    QueueName string `json:"QueueName,omitempty"`
}

// CancelAllMatchmakingTicketsForPlayerResult 
type CancelAllMatchmakingTicketsForPlayerResultModel struct {
}

// CancelAllServerBackfillTicketsForPlayerRequest cancels all backfill tickets of which the player is a member in a given queue that are not cancelled or matched. This
// API is useful if you lose track of what tickets the player is a member of (if the server crashes for instance) and want
// to "reset".
type CancelAllServerBackfillTicketsForPlayerRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity key of the player whose backfill tickets should be canceled.
    Entity* EntityKeyModel `json:"Entity,omitempty"`
    // QueueName the name of the queue from which a player's backfill tickets should be canceled.
    QueueName string `json:"QueueName,omitempty"`
}

// CancelAllServerBackfillTicketsForPlayerResult 
type CancelAllServerBackfillTicketsForPlayerResultModel struct {
}

// CancellationReason 
type CancellationReason string
  
const (
     CancellationReasonRequested CancellationReason = "Requested"
     CancellationReasonInternal CancellationReason = "Internal"
     CancellationReasonTimeout CancellationReason = "Timeout"
      )
// CancelMatchmakingTicketRequest only servers and ticket members can cancel a ticket. The ticket can be in five different states when it is cancelled. 1:
// the ticket is waiting for members to join it, and it has not started matching. If the ticket is cancelled at this stage,
// it will never match. 2: the ticket is matching. If the ticket is cancelled, it will stop matching. 3: the ticket is
// matched. A matched ticket cannot be cancelled. 4: the ticket is already cancelled and nothing happens. 5: the ticket is
// waiting for a server. If the ticket is cancelled, server allocation will be stopped. A server may still be allocated due
// to a race condition, but that will not be reflected in the ticket. There may be race conditions between the ticket
// getting matched and the client making a cancellation request. The client must handle the possibility that the cancel
// request fails if a match is found before the cancellation request is processed. We do not allow resubmitting a cancelled
// ticket because players must consent to enter matchmaking again. Create a new ticket instead.
type CancelMatchmakingTicketRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // QueueName the name of the queue the ticket is in.
    QueueName string `json:"QueueName,omitempty"`
    // TicketId the Id of the ticket to find a match for.
    TicketId string `json:"TicketId,omitempty"`
}

// CancelMatchmakingTicketResult 
type CancelMatchmakingTicketResultModel struct {
}

// CancelServerBackfillTicketRequest only servers can cancel a backfill ticket. The ticket can be in three different states when it is cancelled. 1: the
// ticket is matching. If the ticket is cancelled, it will stop matching. 2: the ticket is matched. A matched ticket cannot
// be cancelled. 3: the ticket is already cancelled and nothing happens. There may be race conditions between the ticket
// getting matched and the server making a cancellation request. The server must handle the possibility that the cancel
// request fails if a match is found before the cancellation request is processed. We do not allow resubmitting a cancelled
// ticket. Create a new ticket instead.
type CancelServerBackfillTicketRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // QueueName the name of the queue the ticket is in.
    QueueName string `json:"QueueName,omitempty"`
    // TicketId the Id of the ticket to find a match for.
    TicketId string `json:"TicketId,omitempty"`
}

// CancelServerBackfillTicketResult 
type CancelServerBackfillTicketResultModel struct {
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
     ContainerFlavorInvalid ContainerFlavor = "Invalid"
      )
// ContainerImageReference 
type ContainerImageReferenceModel struct {
    // ImageName the container image name.
    ImageName string `json:"ImageName,omitempty"`
    // Tag the container tag.
    Tag string `json:"Tag,omitempty"`
}

// CoreCapacity 
type CoreCapacityModel struct {
    // Available the available core capacity for the (Region, VmFamily)
    Available int32 `json:"Available,omitempty"`
    // Region the AzureRegion
    Region string `json:"Region,omitempty"`
    // Total the total core capacity for the (Region, VmFamily)
    Total int32 `json:"Total,omitempty"`
    // VmFamily the AzureVmFamily
    VmFamily AzureVmFamily `json:"VmFamily,omitempty"`
}

// CoreCapacityChange 
type CoreCapacityChangeModel struct {
    // NewCoreLimit new quota core limit for the given vm family/region.
    NewCoreLimit int32 `json:"NewCoreLimit,omitempty"`
    // Region region to change.
    Region string `json:"Region,omitempty"`
    // VmFamily virtual machine family to change.
    VmFamily AzureVmFamily `json:"VmFamily,omitempty"`
}

// CreateBuildAliasRequest creates a multiplayer server build alias and returns the created alias.
type CreateBuildAliasRequestModel struct {
    // AliasName the alias name.
    AliasName string `json:"AliasName,omitempty"`
    // BuildSelectionCriteria array of build selection criteria.
    BuildSelectionCriteria []BuildSelectionCriterionModel `json:"BuildSelectionCriteria,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
}

// CreateBuildWithCustomContainerRequest creates a multiplayer server build with a custom container and returns information about the build creation request.
type CreateBuildWithCustomContainerRequestModel struct {
    // AreAssetsReadonly when true, assets will not be copied for each server inside the VM. All serverswill run from the same set of assets, or
// will have the same assets mounted in the container.
    AreAssetsReadonly bool `json:"AreAssetsReadonly"`
    // BuildName the build name.
    BuildName string `json:"BuildName,omitempty"`
    // ContainerFlavor the flavor of container to create a build from.
    ContainerFlavor ContainerFlavor `json:"ContainerFlavor,omitempty"`
    // ContainerImageReference the container reference, consisting of the image name and tag.
    ContainerImageReference *ContainerImageReferenceModel `json:"ContainerImageReference,omitempty"`
    // ContainerRunCommand the container command to run when the multiplayer server has been allocated, including any arguments.
    ContainerRunCommand string `json:"ContainerRunCommand,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // GameAssetReferences the list of game assets related to the build.
    GameAssetReferences []AssetReferenceParamsModel `json:"GameAssetReferences,omitempty"`
    // GameCertificateReferences the game certificates for the build.
    GameCertificateReferences []GameCertificateReferenceParamsModel `json:"GameCertificateReferences,omitempty"`
    // LinuxInstrumentationConfiguration the Linux instrumentation configuration for the build.
    LinuxInstrumentationConfiguration *LinuxInstrumentationConfigurationModel `json:"LinuxInstrumentationConfiguration,omitempty"`
    // Metadata metadata to tag the build. The keys are case insensitive. The build metadata is made available to the server through
// Game Server SDK (GSDK).Constraints: Maximum number of keys: 30, Maximum key length: 50, Maximum value length: 100
    Metadata map[string]string `json:"Metadata,omitempty"`
    // MonitoringApplicationConfiguration the configuration for the monitoring application on the build
    MonitoringApplicationConfiguration *MonitoringApplicationConfigurationParamsModel `json:"MonitoringApplicationConfiguration,omitempty"`
    // MultiplayerServerCountPerVm the number of multiplayer servers to host on a single VM.
    MultiplayerServerCountPerVm int32 `json:"MultiplayerServerCountPerVm,omitempty"`
    // Ports the ports to map the build on.
    Ports []PortModel `json:"Ports,omitempty"`
    // RegionConfigurations the region configurations for the build.
    RegionConfigurations []BuildRegionParamsModel `json:"RegionConfigurations,omitempty"`
    // ServerResourceConstraints the resource constraints to apply to each server on the VM (EXPERIMENTAL API)
    ServerResourceConstraints *ServerResourceConstraintParamsModel `json:"ServerResourceConstraints,omitempty"`
    // UseStreamingForAssetDownloads when true, assets will be downloaded and uncompressed in memory, without the compressedversion being written first to
// disc.
    UseStreamingForAssetDownloads bool `json:"UseStreamingForAssetDownloads"`
    // VmSize the VM size to create the build on.
    VmSize AzureVmSize `json:"VmSize,omitempty"`
}

// CreateBuildWithCustomContainerResponse 
type CreateBuildWithCustomContainerResponseModel struct {
    // AreAssetsReadonly when true, assets will not be copied for each server inside the VM. All serverswill run from the same set of assets, or
// will have the same assets mounted in the container.
    AreAssetsReadonly bool `json:"AreAssetsReadonly"`
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
    CustomGameContainerImage *ContainerImageReferenceModel `json:"CustomGameContainerImage,omitempty"`
    // GameAssetReferences the game assets for the build.
    GameAssetReferences []AssetReferenceModel `json:"GameAssetReferences,omitempty"`
    // GameCertificateReferences the game certificates for the build.
    GameCertificateReferences []GameCertificateReferenceModel `json:"GameCertificateReferences,omitempty"`
    // LinuxInstrumentationConfiguration the Linux instrumentation configuration for this build.
    LinuxInstrumentationConfiguration *LinuxInstrumentationConfigurationModel `json:"LinuxInstrumentationConfiguration,omitempty"`
    // Metadata the metadata of the build.
    Metadata map[string]string `json:"Metadata,omitempty"`
    // MonitoringApplicationConfiguration the configuration for the monitoring application for the build
    MonitoringApplicationConfiguration *MonitoringApplicationConfigurationModel `json:"MonitoringApplicationConfiguration,omitempty"`
    // MultiplayerServerCountPerVm the number of multiplayer servers to host on a single VM of the build.
    MultiplayerServerCountPerVm int32 `json:"MultiplayerServerCountPerVm,omitempty"`
    // OsPlatform the OS platform used for running the game process.
    OsPlatform string `json:"OsPlatform,omitempty"`
    // Ports the ports the build is mapped on.
    Ports []PortModel `json:"Ports,omitempty"`
    // RegionConfigurations the region configuration for the build.
    RegionConfigurations []BuildRegionModel `json:"RegionConfigurations,omitempty"`
    // ServerResourceConstraints the resource constraints to apply to each server on the VM (EXPERIMENTAL API)
    ServerResourceConstraints *ServerResourceConstraintParamsModel `json:"ServerResourceConstraints,omitempty"`
    // ServerType the type of game server being hosted.
    ServerType string `json:"ServerType,omitempty"`
    // UseStreamingForAssetDownloads when true, assets will be downloaded and uncompressed in memory, without the compressedversion being written first to
// disc.
    UseStreamingForAssetDownloads bool `json:"UseStreamingForAssetDownloads"`
    // VmSize the VM size the build was created on.
    VmSize AzureVmSize `json:"VmSize,omitempty"`
}

// CreateBuildWithManagedContainerRequest creates a multiplayer server build with a managed container and returns information about the build creation request.
type CreateBuildWithManagedContainerRequestModel struct {
    // AreAssetsReadonly when true, assets will not be copied for each server inside the VM. All serverswill run from the same set of assets, or
// will have the same assets mounted in the container.
    AreAssetsReadonly bool `json:"AreAssetsReadonly"`
    // BuildName the build name.
    BuildName string `json:"BuildName,omitempty"`
    // ContainerFlavor the flavor of container to create a build from.
    ContainerFlavor ContainerFlavor `json:"ContainerFlavor,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // GameAssetReferences the list of game assets related to the build.
    GameAssetReferences []AssetReferenceParamsModel `json:"GameAssetReferences,omitempty"`
    // GameCertificateReferences the game certificates for the build.
    GameCertificateReferences []GameCertificateReferenceParamsModel `json:"GameCertificateReferences,omitempty"`
    // GameWorkingDirectory the directory containing the game executable. This would be the start path of the game assets that contain the main game
// server executable. If not provided, a best effort will be made to extract it from the start game command.
    GameWorkingDirectory string `json:"GameWorkingDirectory,omitempty"`
    // InstrumentationConfiguration the instrumentation configuration for the build.
    InstrumentationConfiguration *InstrumentationConfigurationModel `json:"InstrumentationConfiguration,omitempty"`
    // Metadata metadata to tag the build. The keys are case insensitive. The build metadata is made available to the server through
// Game Server SDK (GSDK).Constraints: Maximum number of keys: 30, Maximum key length: 50, Maximum value length: 100
    Metadata map[string]string `json:"Metadata,omitempty"`
    // MonitoringApplicationConfiguration the configuration for the monitoring application on the build
    MonitoringApplicationConfiguration *MonitoringApplicationConfigurationParamsModel `json:"MonitoringApplicationConfiguration,omitempty"`
    // MultiplayerServerCountPerVm the number of multiplayer servers to host on a single VM.
    MultiplayerServerCountPerVm int32 `json:"MultiplayerServerCountPerVm,omitempty"`
    // Ports the ports to map the build on.
    Ports []PortModel `json:"Ports,omitempty"`
    // RegionConfigurations the region configurations for the build.
    RegionConfigurations []BuildRegionParamsModel `json:"RegionConfigurations,omitempty"`
    // ServerResourceConstraints the resource constraints to apply to each server on the VM (EXPERIMENTAL API)
    ServerResourceConstraints *ServerResourceConstraintParamsModel `json:"ServerResourceConstraints,omitempty"`
    // StartMultiplayerServerCommand the command to run when the multiplayer server is started, including any arguments.
    StartMultiplayerServerCommand string `json:"StartMultiplayerServerCommand,omitempty"`
    // UseStreamingForAssetDownloads when true, assets will be downloaded and uncompressed in memory, without the compressedversion being written first to
// disc.
    UseStreamingForAssetDownloads bool `json:"UseStreamingForAssetDownloads"`
    // VmSize the VM size to create the build on.
    VmSize AzureVmSize `json:"VmSize,omitempty"`
    // WindowsCrashDumpConfiguration the crash dump configuration for the build.
    WindowsCrashDumpConfiguration *WindowsCrashDumpConfigurationModel `json:"WindowsCrashDumpConfiguration,omitempty"`
}

// CreateBuildWithManagedContainerResponse 
type CreateBuildWithManagedContainerResponseModel struct {
    // AreAssetsReadonly when true, assets will not be copied for each server inside the VM. All serverswill run from the same set of assets, or
// will have the same assets mounted in the container.
    AreAssetsReadonly bool `json:"AreAssetsReadonly"`
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
    // GameWorkingDirectory the directory containing the game executable. This would be the start path of the game assets that contain the main game
// server executable. If not provided, a best effort will be made to extract it from the start game command.
    GameWorkingDirectory string `json:"GameWorkingDirectory,omitempty"`
    // InstrumentationConfiguration the instrumentation configuration for this build.
    InstrumentationConfiguration *InstrumentationConfigurationModel `json:"InstrumentationConfiguration,omitempty"`
    // Metadata the metadata of the build.
    Metadata map[string]string `json:"Metadata,omitempty"`
    // MonitoringApplicationConfiguration the configuration for the monitoring application for the build
    MonitoringApplicationConfiguration *MonitoringApplicationConfigurationModel `json:"MonitoringApplicationConfiguration,omitempty"`
    // MultiplayerServerCountPerVm the number of multiplayer servers to host on a single VM of the build.
    MultiplayerServerCountPerVm int32 `json:"MultiplayerServerCountPerVm,omitempty"`
    // OsPlatform the OS platform used for running the game process.
    OsPlatform string `json:"OsPlatform,omitempty"`
    // Ports the ports the build is mapped on.
    Ports []PortModel `json:"Ports,omitempty"`
    // RegionConfigurations the region configuration for the build.
    RegionConfigurations []BuildRegionModel `json:"RegionConfigurations,omitempty"`
    // ServerResourceConstraints the resource constraints to apply to each server on the VM (EXPERIMENTAL API)
    ServerResourceConstraints *ServerResourceConstraintParamsModel `json:"ServerResourceConstraints,omitempty"`
    // ServerType the type of game server being hosted.
    ServerType string `json:"ServerType,omitempty"`
    // StartMultiplayerServerCommand the command to run when the multiplayer server has been allocated, including any arguments.
    StartMultiplayerServerCommand string `json:"StartMultiplayerServerCommand,omitempty"`
    // UseStreamingForAssetDownloads when true, assets will be downloaded and uncompressed in memory, without the compressedversion being written first to
// disc.
    UseStreamingForAssetDownloads bool `json:"UseStreamingForAssetDownloads"`
    // VmSize the VM size the build was created on.
    VmSize AzureVmSize `json:"VmSize,omitempty"`
}

// CreateBuildWithProcessBasedServerRequest creates a multiplayer server build with the game server running as a process and returns information about the build
// creation request.
type CreateBuildWithProcessBasedServerRequestModel struct {
    // AreAssetsReadonly when true, assets will not be copied for each server inside the VM. All serverswill run from the same set of assets, or
// will have the same assets mounted in the container.
    AreAssetsReadonly bool `json:"AreAssetsReadonly"`
    // BuildName the build name.
    BuildName string `json:"BuildName,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // GameAssetReferences the list of game assets related to the build.
    GameAssetReferences []AssetReferenceParamsModel `json:"GameAssetReferences,omitempty"`
    // GameCertificateReferences the game certificates for the build.
    GameCertificateReferences []GameCertificateReferenceParamsModel `json:"GameCertificateReferences,omitempty"`
    // GameWorkingDirectory the working directory for the game process. If this is not provided, the working directory will be set based on the
// mount path of the game server executable.
    GameWorkingDirectory string `json:"GameWorkingDirectory,omitempty"`
    // InstrumentationConfiguration the instrumentation configuration for the build.
    InstrumentationConfiguration *InstrumentationConfigurationModel `json:"InstrumentationConfiguration,omitempty"`
    // IsOSPreview indicates whether this build will be created using the OS Preview versionPreview OS is recommended for dev builds to
// detect any breaking changes before they are released to retail. Retail builds should set this value to false.
    IsOSPreview bool `json:"IsOSPreview"`
    // Metadata metadata to tag the build. The keys are case insensitive. The build metadata is made available to the server through
// Game Server SDK (GSDK).Constraints: Maximum number of keys: 30, Maximum key length: 50, Maximum value length: 100
    Metadata map[string]string `json:"Metadata,omitempty"`
    // MonitoringApplicationConfiguration the configuration for the monitoring application on the build
    MonitoringApplicationConfiguration *MonitoringApplicationConfigurationParamsModel `json:"MonitoringApplicationConfiguration,omitempty"`
    // MultiplayerServerCountPerVm the number of multiplayer servers to host on a single VM.
    MultiplayerServerCountPerVm int32 `json:"MultiplayerServerCountPerVm,omitempty"`
    // OsPlatform the OS platform used for running the game process.
    OsPlatform string `json:"OsPlatform,omitempty"`
    // Ports the ports to map the build on.
    Ports []PortModel `json:"Ports,omitempty"`
    // RegionConfigurations the region configurations for the build.
    RegionConfigurations []BuildRegionParamsModel `json:"RegionConfigurations,omitempty"`
    // StartMultiplayerServerCommand the command to run when the multiplayer server is started, including any arguments. The path to any executable should be
// relative to the root asset folder when unzipped.
    StartMultiplayerServerCommand string `json:"StartMultiplayerServerCommand,omitempty"`
    // UseStreamingForAssetDownloads when true, assets will be downloaded and uncompressed in memory, without the compressedversion being written first to
// disc.
    UseStreamingForAssetDownloads bool `json:"UseStreamingForAssetDownloads"`
    // VmSize the VM size to create the build on.
    VmSize AzureVmSize `json:"VmSize,omitempty"`
}

// CreateBuildWithProcessBasedServerResponse 
type CreateBuildWithProcessBasedServerResponseModel struct {
    // AreAssetsReadonly when true, assets will not be copied for each server inside the VM. All serverswill run from the same set of assets, or
// will have the same assets mounted in the container.
    AreAssetsReadonly bool `json:"AreAssetsReadonly"`
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
    // GameWorkingDirectory the working directory for the game process. If this is not provided, the working directory will be set based on the
// mount path of the game server executable.
    GameWorkingDirectory string `json:"GameWorkingDirectory,omitempty"`
    // InstrumentationConfiguration the instrumentation configuration for this build.
    InstrumentationConfiguration *InstrumentationConfigurationModel `json:"InstrumentationConfiguration,omitempty"`
    // IsOSPreview indicates whether this build will be created using the OS Preview versionPreview OS is recommended for dev builds to
// detect any breaking changes before they are released to retail. Retail builds should set this value to false.
    IsOSPreview bool `json:"IsOSPreview"`
    // Metadata the metadata of the build.
    Metadata map[string]string `json:"Metadata,omitempty"`
    // MonitoringApplicationConfiguration the configuration for the monitoring application for the build
    MonitoringApplicationConfiguration *MonitoringApplicationConfigurationModel `json:"MonitoringApplicationConfiguration,omitempty"`
    // MultiplayerServerCountPerVm the number of multiplayer servers to host on a single VM of the build.
    MultiplayerServerCountPerVm int32 `json:"MultiplayerServerCountPerVm,omitempty"`
    // OsPlatform the OS platform used for running the game process.
    OsPlatform string `json:"OsPlatform,omitempty"`
    // Ports the ports the build is mapped on.
    Ports []PortModel `json:"Ports,omitempty"`
    // RegionConfigurations the region configuration for the build.
    RegionConfigurations []BuildRegionModel `json:"RegionConfigurations,omitempty"`
    // ServerType the type of game server being hosted.
    ServerType string `json:"ServerType,omitempty"`
    // StartMultiplayerServerCommand the command to run when the multiplayer server is started, including any arguments. The path to any executable is
// relative to the root asset folder when unzipped.
    StartMultiplayerServerCommand string `json:"StartMultiplayerServerCommand,omitempty"`
    // UseStreamingForAssetDownloads when true, assets will be downloaded and uncompressed in memory, without the compressedversion being written first to
// disc.
    UseStreamingForAssetDownloads bool `json:"UseStreamingForAssetDownloads"`
    // VmSize the VM size the build was created on.
    VmSize AzureVmSize `json:"VmSize,omitempty"`
}

// CreateLobbyRequest request to create a lobby. A Server or client can create a lobby.
type CreateLobbyRequestModel struct {
    // AccessPolicy the policy indicating who is allowed to join the lobby, and the visibility to queries. May be 'Public', 'Friends' or
// 'Private'. Public means the lobby is both visible in queries and any player may join, including invited players. Friends
// means that users who are bidirectional friends of members in the lobby may search to find friend lobbies, to retrieve
// its connection string. Private means the lobby is not visible in queries, and a player must receive an invitation to
// join. Defaults to 'Public' on creation. Can only be changed by the lobby owner.
    AccessPolicy AccessPolicy `json:"AccessPolicy,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // LobbyData the private key-value pairs which are only visible to members of the lobby. At most 30 key-value pairs may be stored
// here, keys are limited to 30 characters and values to 1000. The total size of all lobbyData values may not exceed 4096
// bytes. Keys are case sensitive.
    LobbyData map[string]string `json:"LobbyData,omitempty"`
    // MaxPlayers the maximum number of players allowed in the lobby. The value must be between 2 and 128.
    MaxPlayers uint32 `json:"MaxPlayers,omitempty"`
    // Members the member initially added to the lobby. Client must specify exactly one member, which is the creator's entity and
// member data. Member PubSubConnectionHandle must be null or empty. Game servers must not specify any members.
    Members []MemberModel `json:"Members,omitempty"`
    // Owner the lobby owner. Must be the calling entity.
    Owner* EntityKeyModel `json:"Owner,omitempty"`
    // OwnerMigrationPolicy the policy for how a new owner is chosen. May be 'Automatic', 'Manual' or 'None'. Can only be specified by clients. If
// client-owned and 'Automatic' - The Lobby service will automatically assign another connected owner when the current
// owner leaves or disconnects. The useConnections property must be true. If client - owned and 'Manual' - Ownership is
// protected as long as the current owner is connected. If the current owner leaves or disconnects any member may set
// themselves as the current owner. The useConnections property must be true. If client-owned and 'None' - Any member can
// set ownership. The useConnections property can be either true or false.
    OwnerMigrationPolicy OwnerMigrationPolicy `json:"OwnerMigrationPolicy,omitempty"`
    // SearchData the public key-value pairs which allow queries to differentiate between lobbies. Queries will refer to these key-value
// pairs in their filter and order by clauses to retrieve lobbies fitting the specified criteria. At most 30 key-value
// pairs may be stored here. Keys are of the format string_key1, string_key2 ... string_key30 for string values, or
// number_key1, number_key2, ... number_key30 for numeric values.Numeric values are floats. Values can be at most 256
// characters long. The total size of all searchData values may not exceed 1024 bytes.
    SearchData map[string]string `json:"SearchData,omitempty"`
    // UseConnections a setting to control whether connections are used. Defaults to true. When true, notifications are sent to subscribed
// players, disconnect detection removes connectionHandles, only owner migration policies using connections are allowed,
// and lobbies must have at least one connected member to be searchable or be a server hosted lobby with a connected
// server. If false, then notifications are not sent, connections are not allowed, and lobbies do not need connections to
// be searchable.
    UseConnections bool `json:"UseConnections"`
}

// CreateLobbyResult 
type CreateLobbyResultModel struct {
    // ConnectionString a field which indicates which lobby the user will be joining.
    ConnectionString string `json:"ConnectionString,omitempty"`
    // LobbyId id to uniquely identify a lobby.
    LobbyId string `json:"LobbyId,omitempty"`
}

// CreateMatchmakingTicketRequest the client specifies the creator's attributes and optionally a list of other users to match with.
type CreateMatchmakingTicketRequestModel struct {
    // Creator the User who created this ticket.
    Creator* MatchmakingPlayerModel `json:"Creator,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // GiveUpAfterSeconds how long to attempt matching this ticket in seconds.
    GiveUpAfterSeconds int32 `json:"GiveUpAfterSeconds,omitempty"`
    // MembersToMatchWith a list of Entity Keys of other users to match with.
    MembersToMatchWith []EntityKeyModel `json:"MembersToMatchWith,omitempty"`
    // QueueName the Id of a match queue.
    QueueName string `json:"QueueName,omitempty"`
}

// CreateMatchmakingTicketResult 
type CreateMatchmakingTicketResultModel struct {
    // TicketId the Id of the ticket to find a match for.
    TicketId string `json:"TicketId,omitempty"`
}

// CreateRemoteUserRequest creates a remote user to log on to a VM for a multiplayer server build in a specific region. Returns user credential
// information necessary to log on.
type CreateRemoteUserRequestModel struct {
    // BuildId the guid string build ID of to create the remote user for.
    BuildId string `json:"BuildId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // ExpirationTime the expiration time for the remote user created. Defaults to expiring in one day if not specified.
    ExpirationTime time.Time `json:"ExpirationTime,omitempty"`
    // Region the region of virtual machine to create the remote user for.
    Region string `json:"Region,omitempty"`
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

// CreateServerBackfillTicketRequest the server specifies all the members, their teams and their attributes, and the server details if applicable.
type CreateServerBackfillTicketRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // GiveUpAfterSeconds how long to attempt matching this ticket in seconds.
    GiveUpAfterSeconds int32 `json:"GiveUpAfterSeconds,omitempty"`
    // Members the users who will be part of this ticket, along with their team assignments.
    Members []MatchmakingPlayerWithTeamAssignmentModel `json:"Members,omitempty"`
    // QueueName the Id of a match queue.
    QueueName string `json:"QueueName,omitempty"`
    // ServerDetails the details of the server the members are connected to.
    ServerDetails *ServerDetailsModel `json:"ServerDetails,omitempty"`
}

// CreateServerBackfillTicketResult 
type CreateServerBackfillTicketResultModel struct {
    // TicketId the Id of the ticket to find a match for.
    TicketId string `json:"TicketId,omitempty"`
}

// CreateServerMatchmakingTicketRequest the server specifies all the members and their attributes.
type CreateServerMatchmakingTicketRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // GiveUpAfterSeconds how long to attempt matching this ticket in seconds.
    GiveUpAfterSeconds int32 `json:"GiveUpAfterSeconds,omitempty"`
    // Members the users who will be part of this ticket.
    Members []MatchmakingPlayerModel `json:"Members,omitempty"`
    // QueueName the Id of a match queue.
    QueueName string `json:"QueueName,omitempty"`
}

// CreateTitleMultiplayerServersQuotaChangeRequest creates a request to change a title's multiplayer server quotas.
type CreateTitleMultiplayerServersQuotaChangeRequestModel struct {
    // ChangeDescription a brief description of the requested changes.
    ChangeDescription string `json:"ChangeDescription,omitempty"`
    // Changes changes to make to the titles cores quota.
    Changes []CoreCapacityChangeModel `json:"Changes,omitempty"`
    // ContactEmail email to be contacted by our team about this request. Only required when a request is not approved.
    ContactEmail string `json:"ContactEmail,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Notes additional information about this request that our team can use to better understand the requirements.
    Notes string `json:"Notes,omitempty"`
    // StartDate when these changes would need to be in effect. Only required when a request is not approved.
    StartDate time.Time `json:"StartDate,omitempty"`
}

// CreateTitleMultiplayerServersQuotaChangeResponse 
type CreateTitleMultiplayerServersQuotaChangeResponseModel struct {
    // RequestId id of the change request that was created.
    RequestId string `json:"RequestId,omitempty"`
    // WasApproved determines if the request was approved or not. When false, our team is reviewing and may respond within 2 business days.
    WasApproved bool `json:"WasApproved"`
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
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // FileName the filename of the asset to delete.
    FileName string `json:"FileName,omitempty"`
}

// DeleteBuildAliasRequest deletes a multiplayer server build alias.
type DeleteBuildAliasRequestModel struct {
    // AliasId the guid string alias ID of the alias to perform the action on.
    AliasId string `json:"AliasId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
}

// DeleteBuildRegionRequest removes a multiplayer server build's region.
type DeleteBuildRegionRequestModel struct {
    // BuildId the guid string ID of the build we want to update regions for.
    BuildId string `json:"BuildId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Region the build region to delete.
    Region string `json:"Region,omitempty"`
}

// DeleteBuildRequest deletes a multiplayer server build.
type DeleteBuildRequestModel struct {
    // BuildId the guid string build ID of the build to delete.
    BuildId string `json:"BuildId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
}

// DeleteCertificateRequest deletes a multiplayer server game certificate.
type DeleteCertificateRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Name the name of the certificate.
    Name string `json:"Name,omitempty"`
}

// DeleteContainerImageRequest removes the specified container image repository. After this operation, a 'docker pull' will fail for all the tags of
// the specified image. Morever, ListContainerImages will not return the specified image.
type DeleteContainerImageRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // ImageName the container image repository we want to delete.
    ImageName string `json:"ImageName,omitempty"`
}

// DeleteLobbyRequest request to delete a lobby. Only servers can delete lobbies.
type DeleteLobbyRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // LobbyId the id of the lobby.
    LobbyId string `json:"LobbyId,omitempty"`
}

// DeleteRemoteUserRequest deletes a remote user to log on to a VM for a multiplayer server build in a specific region. Returns user credential
// information necessary to log on.
type DeleteRemoteUserRequestModel struct {
    // BuildId the guid string build ID of the multiplayer server where the remote user is to delete.
    BuildId string `json:"BuildId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Region the region of the multiplayer server where the remote user is to delete.
    Region string `json:"Region,omitempty"`
    // Username the username of the remote user to delete.
    Username string `json:"Username,omitempty"`
    // VmId the virtual machine ID the multiplayer server is located on.
    VmId string `json:"VmId,omitempty"`
}

// DynamicStandbySettings 
type DynamicStandbySettingsModel struct {
    // DynamicFloorMultiplierThresholds list of auto standing by trigger values and corresponding standing by multiplier. Defaults to 1.5X at 50%, 3X at 25%,
// and 4X at 5%
    DynamicFloorMultiplierThresholds []DynamicStandbyThresholdModel `json:"DynamicFloorMultiplierThresholds,omitempty"`
    // IsEnabled when true, dynamic standby will be enabled
    IsEnabled bool `json:"IsEnabled"`
    // RampDownSeconds the time it takes to reduce target standing by to configured floor value after an increase. Defaults to 30 minutes
    RampDownSeconds int32 `json:"RampDownSeconds,omitempty"`
}

// DynamicStandbyThreshold 
type DynamicStandbyThresholdModel struct {
    // Multiplier when the trigger threshold is reached, multiply by this value
    Multiplier float64 `json:"Multiplier,omitempty"`
    // TriggerThresholdPercentage the multiplier will be applied when the actual standby divided by target standby floor is less than this value
    TriggerThresholdPercentage float64 `json:"TriggerThresholdPercentage,omitempty"`
}

// EmptyResponse 
type EmptyResponseModel struct {
}

// EnableMultiplayerServersForTitleRequest enables the multiplayer server feature for a title and returns the enabled status. The enabled status can be
// Initializing, Enabled, and Disabled. It can up to 20 minutes or more for the title to be enabled for the feature. On
// average, it can take up to 20 minutes for the title to be enabled for the feature.
type EnableMultiplayerServersForTitleRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
}

// EnableMultiplayerServersForTitleResponse 
type EnableMultiplayerServersForTitleResponseModel struct {
    // Status the enabled status for the multiplayer server features for the title.
    Status TitleMultiplayerServerEnabledStatus `json:"Status,omitempty"`
}

// EntityKey combined entity type and ID structure which uniquely identifies a single entity.
type EntityKeyModel struct {
    // Id unique ID of the entity.
    Id string `json:"Id,omitempty"`
    // Type entity type. See https://docs.microsoft.com/gaming/playfab/features/data/entities/available-built-in-entity-types
    Type string `json:"Type,omitempty"`
}

// ExternalFriendSources 
type ExternalFriendSources string
  
const (
     ExternalFriendSourcesNone ExternalFriendSources = "None"
     ExternalFriendSourcesSteam ExternalFriendSources = "Steam"
     ExternalFriendSourcesFacebook ExternalFriendSources = "Facebook"
     ExternalFriendSourcesXbox ExternalFriendSources = "Xbox"
     ExternalFriendSourcesPsn ExternalFriendSources = "Psn"
     ExternalFriendSourcesAll ExternalFriendSources = "All"
      )
// FindFriendLobbiesRequest request to find friends lobbies. Only a client can find friend lobbies.
type FindFriendLobbiesRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // ExcludeFacebookFriends controls whether this query should link to friends made on the Facebook network. Defaults to false
    ExcludeFacebookFriends bool `json:"ExcludeFacebookFriends"`
    // ExcludeSteamFriends controls whether this query should link to friends made on the Steam network. Defaults to false
    ExcludeSteamFriends bool `json:"ExcludeSteamFriends"`
    // ExternalPlatformFriends indicates which other platforms' friends this query should link to.
    ExternalPlatformFriends ExternalFriendSources `json:"ExternalPlatformFriends,omitempty"`
    // Filter oData style string that contains one or more filters. Only the following operators are supported: "and" (logical and),
// "eq" (equal), "ne" (not equals), "ge" (greater than or equal), "gt" (greater than), "le" (less than or equal), and "lt"
// (less than). The left-hand side of each OData logical expression should be either a search property key (e.g.
// string_key1, number_key3, etc) or one of the pre-defined search keys all of which must be prefixed by "lobby/":
// lobby/memberCount (number of players in a lobby), lobby/maxMemberCount (maximum number of players allowed in a lobby),
// lobby/membershipLock (must equal 'Unlocked' or 'Locked'), lobby/amOwner (required to equal "true"), lobby/amMember
// (required to equal "true").
    Filter string `json:"Filter,omitempty"`
    // OrderBy oData style string that contains sorting for this query in either ascending ("asc") or descending ("desc") order.
// OrderBy clauses are of the form "number_key1 asc" or the pre-defined search key "lobby/memberCount asc" and
// "lobby/maxMemberCount desc". To sort by closest, a moniker `distance{number_key1 = 5}` can be used to sort by distance
// from the given number. This field only supports either one sort clause or one distance clause.
    OrderBy string `json:"OrderBy,omitempty"`
    // Pagination request pagination information.
    Pagination *PaginationRequestModel `json:"Pagination,omitempty"`
    // XboxToken xbox token if Xbox friends should be included. Requires Xbox be configured on PlayFab.
    XboxToken string `json:"XboxToken,omitempty"`
}

// FindFriendLobbiesResult 
type FindFriendLobbiesResultModel struct {
    // Lobbies array of lobbies found that matched FindFriendLobbies request.
    Lobbies []FriendLobbySummaryModel `json:"Lobbies,omitempty"`
    // Pagination pagination response for FindFriendLobbies request.
    Pagination* PaginationResponseModel `json:"Pagination,omitempty"`
}

// FindLobbiesRequest request to find lobbies.
type FindLobbiesRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Filter oData style string that contains one or more filters. Only the following operators are supported: "and" (logical and),
// "eq" (equal), "ne" (not equals), "ge" (greater than or equal), "gt" (greater than), "le" (less than or equal), and "lt"
// (less than). The left-hand side of each OData logical expression should be either a search property key (e.g.
// string_key1, number_key3, etc) or one of the pre-defined search keys all of which must be prefixed by "lobby/":
// lobby/memberCount (number of players in a lobby), lobby/maxMemberCount (maximum number of players allowed in a lobby),
// lobby/membershipLock (must equal 'Unlocked' or 'Locked'), lobby/amOwner (required to equal "true"), lobby/amMember
// (required to equal "true").
    Filter string `json:"Filter,omitempty"`
    // OrderBy oData style string that contains sorting for this query in either ascending ("asc") or descending ("desc") order.
// OrderBy clauses are of the form "number_key1 asc" or the pre-defined search key "lobby/memberCount asc" and
// "lobby/maxMemberCount desc". To sort by closest, a moniker `distance{number_key1 = 5}` can be used to sort by distance
// from the given number. This field only supports either one sort clause or one distance clause.
    OrderBy string `json:"OrderBy,omitempty"`
    // Pagination request pagination information.
    Pagination *PaginationRequestModel `json:"Pagination,omitempty"`
}

// FindLobbiesResult 
type FindLobbiesResultModel struct {
    // Lobbies array of lobbies found that matched FindLobbies request.
    Lobbies []LobbySummaryModel `json:"Lobbies,omitempty"`
    // Pagination pagination response for FindLobbies request.
    Pagination* PaginationResponseModel `json:"Pagination,omitempty"`
}

// FriendLobbySummary 
type FriendLobbySummaryModel struct {
    // ConnectionString a string used to join the lobby.This field is populated by the Lobby service.Invites are performed by communicating this
// connectionString to other players.
    ConnectionString string `json:"ConnectionString,omitempty"`
    // CurrentPlayers the current number of players in the lobby.
    CurrentPlayers uint32 `json:"CurrentPlayers,omitempty"`
    // Friends friends in Lobby.
    Friends []EntityKeyModel `json:"Friends,omitempty"`
    // LobbyId id to uniquely identify a lobby.
    LobbyId string `json:"LobbyId,omitempty"`
    // MaxPlayers the maximum number of players allowed in the lobby.
    MaxPlayers uint32 `json:"MaxPlayers,omitempty"`
    // MembershipLock a setting indicating whether members are allowed to join this lobby. When Locked new members are prevented from joining.
    MembershipLock MembershipLock `json:"MembershipLock,omitempty"`
    // Owner the client or server entity which owns this lobby.
    Owner* EntityKeyModel `json:"Owner,omitempty"`
    // SearchData search data.
    SearchData map[string]string `json:"SearchData,omitempty"`
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

// GetAssetDownloadUrlRequest gets a URL that can be used to download the specified asset.
type GetAssetDownloadUrlRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // FileName the asset's file name to get the download URL for.
    FileName string `json:"FileName,omitempty"`
}

// GetAssetDownloadUrlResponse 
type GetAssetDownloadUrlResponseModel struct {
    // AssetDownloadUrl the asset's download URL.
    AssetDownloadUrl string `json:"AssetDownloadUrl,omitempty"`
    // FileName the asset's file name to get the download URL for.
    FileName string `json:"FileName,omitempty"`
}

// GetAssetUploadUrlRequest gets the URL to upload assets to.
type GetAssetUploadUrlRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
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

// GetBuildAliasRequest returns the details about a multiplayer server build alias.
type GetBuildAliasRequestModel struct {
    // AliasId the guid string alias ID of the alias to perform the action on.
    AliasId string `json:"AliasId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
}

// GetBuildRequest returns the details about a multiplayer server build.
type GetBuildRequestModel struct {
    // BuildId the guid string build ID of the build to get.
    BuildId string `json:"BuildId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
}

// GetBuildResponse 
type GetBuildResponseModel struct {
    // AreAssetsReadonly when true, assets will not be copied for each server inside the VM. All serverswill run from the same set of assets, or
// will have the same assets mounted in the container.
    AreAssetsReadonly bool `json:"AreAssetsReadonly"`
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
    CustomGameContainerImage *ContainerImageReferenceModel `json:"CustomGameContainerImage,omitempty"`
    // GameAssetReferences the game assets for the build.
    GameAssetReferences []AssetReferenceModel `json:"GameAssetReferences,omitempty"`
    // GameCertificateReferences the game certificates for the build.
    GameCertificateReferences []GameCertificateReferenceModel `json:"GameCertificateReferences,omitempty"`
    // InstrumentationConfiguration the instrumentation configuration of the build.
    InstrumentationConfiguration *InstrumentationConfigurationModel `json:"InstrumentationConfiguration,omitempty"`
    // Metadata metadata of the build. The keys are case insensitive. The build metadata is made available to the server through Game
// Server SDK (GSDK).
    Metadata map[string]string `json:"Metadata,omitempty"`
    // MultiplayerServerCountPerVm the number of multiplayer servers to hosted on a single VM of the build.
    MultiplayerServerCountPerVm int32 `json:"MultiplayerServerCountPerVm,omitempty"`
    // OsPlatform the OS platform used for running the game process.
    OsPlatform string `json:"OsPlatform,omitempty"`
    // Ports the ports the build is mapped on.
    Ports []PortModel `json:"Ports,omitempty"`
    // RegionConfigurations the region configuration for the build.
    RegionConfigurations []BuildRegionModel `json:"RegionConfigurations,omitempty"`
    // ServerResourceConstraints the resource constraints to apply to each server on the VM.
    ServerResourceConstraints *ServerResourceConstraintParamsModel `json:"ServerResourceConstraints,omitempty"`
    // ServerType the type of game server being hosted.
    ServerType string `json:"ServerType,omitempty"`
    // StartMultiplayerServerCommand the command to run when the multiplayer server has been allocated, including any arguments. This only applies to managed
// builds. If the build is a custom build, this field will be null.
    StartMultiplayerServerCommand string `json:"StartMultiplayerServerCommand,omitempty"`
    // VmSize the VM size the build was created on.
    VmSize AzureVmSize `json:"VmSize,omitempty"`
}

// GetContainerRegistryCredentialsRequest gets credentials to the container registry where game developers can upload custom container images to before creating a
// new build.
type GetContainerRegistryCredentialsRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
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

// GetLobbyRequest request to get a lobby.
type GetLobbyRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // LobbyId the id of the lobby.
    LobbyId string `json:"LobbyId,omitempty"`
}

// GetLobbyResult 
type GetLobbyResultModel struct {
    // Lobby the information pertaining to the requested lobby.
    Lobby* LobbyModel `json:"Lobby,omitempty"`
}

// GetMatchmakingTicketRequest the ticket includes the invited players, their attributes if they have joined, the ticket status, the match Id when
// applicable, etc. Only servers, the ticket creator and the invited players can get the ticket.
type GetMatchmakingTicketRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // EscapeObject determines whether the matchmaking attributes will be returned as an escaped JSON string or as an un-escaped JSON
// object.
    EscapeObject bool `json:"EscapeObject"`
    // QueueName the name of the queue to find a match for.
    QueueName string `json:"QueueName,omitempty"`
    // TicketId the Id of the ticket to find a match for.
    TicketId string `json:"TicketId,omitempty"`
}

// GetMatchmakingTicketResult 
type GetMatchmakingTicketResultModel struct {
    // CancellationReasonString the reason why the current ticket was canceled. This field is only set if the ticket is in canceled state.
    CancellationReasonString string `json:"CancellationReasonString,omitempty"`
    // ChangeNumber change number used for differentiating older matchmaking status updates from newer ones.
    ChangeNumber uint32 `json:"ChangeNumber,omitempty"`
    // Created the server date and time at which ticket was created.
    Created time.Time `json:"Created,omitempty"`
    // Creator the Creator's entity key.
    Creator* EntityKeyModel `json:"Creator,omitempty"`
    // GiveUpAfterSeconds how long to attempt matching this ticket in seconds.
    GiveUpAfterSeconds int32 `json:"GiveUpAfterSeconds,omitempty"`
    // MatchId the Id of a match.
    MatchId string `json:"MatchId,omitempty"`
    // Members a list of Users that have joined this ticket.
    Members []MatchmakingPlayerModel `json:"Members,omitempty"`
    // MembersToMatchWith a list of PlayFab Ids of Users to match with.
    MembersToMatchWith []EntityKeyModel `json:"MembersToMatchWith,omitempty"`
    // QueueName the Id of a match queue.
    QueueName string `json:"QueueName,omitempty"`
    // Status the current ticket status. Possible values are: WaitingForPlayers, WaitingForMatch, WaitingForServer, Canceled and
// Matched.
    Status string `json:"Status,omitempty"`
    // TicketId the Id of the ticket to find a match for.
    TicketId string `json:"TicketId,omitempty"`
}

// GetMatchRequest when matchmaking has successfully matched together a collection of tickets, it produces a 'match' with an Id. The match
// contains all of the players that were matched together, and their team assigments. Only servers and ticket members can
// get the match.
type GetMatchRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // EscapeObject determines whether the matchmaking attributes will be returned as an escaped JSON string or as an un-escaped JSON
// object.
    EscapeObject bool `json:"EscapeObject"`
    // MatchId the Id of a match.
    MatchId string `json:"MatchId,omitempty"`
    // QueueName the name of the queue to join.
    QueueName string `json:"QueueName,omitempty"`
    // ReturnMemberAttributes determines whether the matchmaking attributes for each user should be returned in the response for match request.
    ReturnMemberAttributes bool `json:"ReturnMemberAttributes"`
}

// GetMatchResult 
type GetMatchResultModel struct {
    // ArrangementString a string that is used by players that are matched together to join an arranged lobby.
    ArrangementString string `json:"ArrangementString,omitempty"`
    // MatchId the Id of a match.
    MatchId string `json:"MatchId,omitempty"`
    // Members a list of Users that are matched together, along with their team assignments.
    Members []MatchmakingPlayerWithTeamAssignmentModel `json:"Members,omitempty"`
    // RegionPreferences a list of regions that the match could be played in sorted by preference. This value is only set if the queue has a
// region selection rule.
    RegionPreferences []string `json:"RegionPreferences,omitempty"`
    // ServerDetails the details of the server that the match has been allocated to.
    ServerDetails *ServerDetailsModel `json:"ServerDetails,omitempty"`
}

// GetMultiplayerServerDetailsRequest gets multiplayer server session details for a build in a specific region.
type GetMultiplayerServerDetailsRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // SessionId the title generated guid string session ID of the multiplayer server to get details for. This is to keep track of
// multiplayer server sessions.
    SessionId string `json:"SessionId,omitempty"`
}

// GetMultiplayerServerDetailsResponse 
type GetMultiplayerServerDetailsResponseModel struct {
    // BuildId the identity of the build in which the server was allocated.
    BuildId string `json:"BuildId,omitempty"`
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
    Region string `json:"Region,omitempty"`
    // ServerId the string server ID of the multiplayer server generated by PlayFab.
    ServerId string `json:"ServerId,omitempty"`
    // SessionId the guid string session ID of the multiplayer server.
    SessionId string `json:"SessionId,omitempty"`
    // State the state of the multiplayer server.
    State string `json:"State,omitempty"`
    // VmId the virtual machine ID that the multiplayer server is located on.
    VmId string `json:"VmId,omitempty"`
}

// GetMultiplayerServerLogsRequest gets multiplayer server logs for a specific server id in a region. The logs are available only after a server has
// terminated.
type GetMultiplayerServerLogsRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // ServerId the server ID of multiplayer server to get logs for.
    ServerId string `json:"ServerId,omitempty"`
}

// GetMultiplayerServerLogsResponse 
type GetMultiplayerServerLogsResponseModel struct {
    // LogDownloadUrl uRL for logs download.
    LogDownloadUrl string `json:"LogDownloadUrl,omitempty"`
}

// GetMultiplayerSessionLogsBySessionIdRequest gets multiplayer server logs for a specific server id in a region. The logs are available only after a server has
// terminated.
type GetMultiplayerSessionLogsBySessionIdRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // SessionId the server ID of multiplayer server to get logs for.
    SessionId string `json:"SessionId,omitempty"`
}

// GetQueueStatisticsRequest returns the matchmaking statistics for a queue. These include the number of players matching and the statistics related
// to the time to match statistics in seconds (average and percentiles). Statistics are refreshed once every 5 minutes.
// Servers can access all statistics no matter what the ClientStatisticsVisibility is configured to. Clients can access
// statistics according to the ClientStatisticsVisibility. Client requests are forbidden if all visibility fields are
// false.
type GetQueueStatisticsRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // QueueName the name of the queue.
    QueueName string `json:"QueueName,omitempty"`
}

// GetQueueStatisticsResult 
type GetQueueStatisticsResultModel struct {
    // NumberOfPlayersMatching the current number of players in the matchmaking queue, who are waiting to be matched.
    NumberOfPlayersMatching uint32 `json:"NumberOfPlayersMatching,omitempty"`
    // TimeToMatchStatisticsInSeconds statistics representing the time (in seconds) it takes for tickets to find a match.
    TimeToMatchStatisticsInSeconds *StatisticsModel `json:"TimeToMatchStatisticsInSeconds,omitempty"`
}

// GetRemoteLoginEndpointRequest gets a remote login endpoint to a VM that is hosting a multiplayer server build in a specific region.
type GetRemoteLoginEndpointRequestModel struct {
    // BuildId the guid string build ID of the multiplayer server to get remote login information for.
    BuildId string `json:"BuildId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Region the region of the multiplayer server to get remote login information for.
    Region string `json:"Region,omitempty"`
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

// GetServerBackfillTicketRequest the ticket includes the players, their attributes, their teams, the ticket status, the match Id and the server details
// when applicable, etc. Only servers can get the ticket.
type GetServerBackfillTicketRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // EscapeObject determines whether the matchmaking attributes will be returned as an escaped JSON string or as an un-escaped JSON
// object.
    EscapeObject bool `json:"EscapeObject"`
    // QueueName the name of the queue to find a match for.
    QueueName string `json:"QueueName,omitempty"`
    // TicketId the Id of the ticket to find a match for.
    TicketId string `json:"TicketId,omitempty"`
}

// GetServerBackfillTicketResult 
type GetServerBackfillTicketResultModel struct {
    // CancellationReasonString the reason why the current ticket was canceled. This field is only set if the ticket is in canceled state.
    CancellationReasonString string `json:"CancellationReasonString,omitempty"`
    // Created the server date and time at which ticket was created.
    Created time.Time `json:"Created,omitempty"`
    // GiveUpAfterSeconds how long to attempt matching this ticket in seconds.
    GiveUpAfterSeconds int32 `json:"GiveUpAfterSeconds,omitempty"`
    // MatchId the Id of a match.
    MatchId string `json:"MatchId,omitempty"`
    // Members a list of Users that are part of this ticket, along with their team assignments.
    Members []MatchmakingPlayerWithTeamAssignmentModel `json:"Members,omitempty"`
    // QueueName the Id of a match queue.
    QueueName string `json:"QueueName,omitempty"`
    // ServerDetails the details of the server the members are connected to.
    ServerDetails* ServerDetailsModel `json:"ServerDetails,omitempty"`
    // Status the current ticket status. Possible values are: WaitingForMatch, Canceled and Matched.
    Status string `json:"Status,omitempty"`
    // TicketId the Id of the ticket to find a match for.
    TicketId string `json:"TicketId,omitempty"`
}

// GetTitleEnabledForMultiplayerServersStatusRequest gets the status of whether a title is enabled for the multiplayer server feature. The enabled status can be
// Initializing, Enabled, and Disabled.
type GetTitleEnabledForMultiplayerServersStatusRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
}

// GetTitleEnabledForMultiplayerServersStatusResponse 
type GetTitleEnabledForMultiplayerServersStatusResponseModel struct {
    // Status the enabled status for the multiplayer server features for the title.
    Status TitleMultiplayerServerEnabledStatus `json:"Status,omitempty"`
}

// GetTitleMultiplayerServersQuotaChangeRequest gets a title's server quota change request.
type GetTitleMultiplayerServersQuotaChangeRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // RequestId id of the change request to get.
    RequestId string `json:"RequestId,omitempty"`
}

// GetTitleMultiplayerServersQuotaChangeResponse 
type GetTitleMultiplayerServersQuotaChangeResponseModel struct {
    // Change the change request for this title.
    Change *QuotaChangeModel `json:"Change,omitempty"`
}

// GetTitleMultiplayerServersQuotasRequest gets the quotas for a title in relation to multiplayer servers.
type GetTitleMultiplayerServersQuotasRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
}

// GetTitleMultiplayerServersQuotasResponse 
type GetTitleMultiplayerServersQuotasResponseModel struct {
    // Quotas the various quotas for multiplayer servers for the title.
    Quotas *TitleMultiplayerServersQuotasModel `json:"Quotas,omitempty"`
}

// InstrumentationConfiguration 
type InstrumentationConfigurationModel struct {
    // IsEnabled designates whether windows instrumentation configuration will be enabled for this Build
    IsEnabled bool `json:"IsEnabled"`
    // ProcessesToMonitor this property is deprecated, use IsEnabled. The list of processes to be monitored on a VM for this build. Providing
// processes will turn on performance metrics collection for this build. Process names should not include extensions. If
// the game server process is: GameServer.exe; then, ProcessesToMonitor = [ GameServer ]
    ProcessesToMonitor []string `json:"ProcessesToMonitor,omitempty"`
}

// InviteToLobbyRequest request to invite a player to a lobby the caller is already a member of. Only a client can invite another player to a
// lobby.
type InviteToLobbyRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // InviteeEntity the entity invited to the lobby.
    InviteeEntity *EntityKeyModel `json:"InviteeEntity,omitempty"`
    // LobbyId the id of the lobby.
    LobbyId string `json:"LobbyId,omitempty"`
    // MemberEntity the member entity sending the invite. Must be a member of the lobby.
    MemberEntity *EntityKeyModel `json:"MemberEntity,omitempty"`
}

// JoinArrangedLobbyRequest request to join an arranged lobby. Only a client can join an arranged lobby.
type JoinArrangedLobbyRequestModel struct {
    // AccessPolicy the policy indicating who is allowed to join the lobby, and the visibility to queries. May be 'Public', 'Friends' or
// 'Private'. Public means the lobby is both visible in queries and any player may join, including invited players. Friends
// means that users who are bidirectional friends of members in the lobby may search to find friend lobbies, to retrieve
// its connection string. Private means the lobby is not visible in queries, and a player must receive an invitation to
// join. Defaults to 'Public' on creation. Can only be changed by the lobby owner.
    AccessPolicy AccessPolicy `json:"AccessPolicy,omitempty"`
    // ArrangementString a field which indicates which lobby the user will be joining. This field is opaque to everyone except the Lobby service
// and the creator of the arrangementString (Matchmaking). This string defines a unique identifier for the arranged lobby
// as well as the title and member the string is valid for. Arrangement strings have an expiration.
    ArrangementString string `json:"ArrangementString,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // MaxPlayers the maximum number of players allowed in the lobby. The value must be between 2 and 128.
    MaxPlayers uint32 `json:"MaxPlayers,omitempty"`
    // MemberData the private key-value pairs used by the member to communicate information to other members and the owner. Visible to all
// members of the lobby. At most 30 key-value pairs may be stored here, keys are limited to 30 characters and values to
// 1000. The total size of all memberData values may not exceed 4096 bytes. Keys are case sensitive.
    MemberData map[string]string `json:"MemberData,omitempty"`
    // MemberEntity the member entity who is joining the lobby. The first member to join will be the lobby owner.
    MemberEntity* EntityKeyModel `json:"MemberEntity,omitempty"`
    // OwnerMigrationPolicy the policy for how a new owner is chosen. May be 'Automatic', 'Manual' or 'None'. Can only be specified by clients. If
// client-owned and 'Automatic' - The Lobby service will automatically assign another connected owner when the current
// owner leaves or disconnects. The useConnections property must be true. If client - owned and 'Manual' - Ownership is
// protected as long as the current owner is connected. If the current owner leaves or disconnects any member may set
// themselves as the current owner. The useConnections property must be true. If client-owned and 'None' - Any member can
// set ownership. The useConnections property can be either true or false.
    OwnerMigrationPolicy OwnerMigrationPolicy `json:"OwnerMigrationPolicy,omitempty"`
    // UseConnections a setting to control whether connections are used. Defaults to true. When true, notifications are sent to subscribed
// players, disconnect detection removes connectionHandles, only owner migration policies using connections are allowed,
// and lobbies must have at least one connected member to be searchable or be a server hosted lobby with a connected
// server. If false, then notifications are not sent, connections are not allowed, and lobbies do not need connections to
// be searchable.
    UseConnections bool `json:"UseConnections"`
}

// JoinLobbyRequest request to join a lobby. Only a client can join a lobby.
type JoinLobbyRequestModel struct {
    // ConnectionString a field which indicates which lobby the user will be joining. This field is opaque to everyone except the Lobby service.
    ConnectionString string `json:"ConnectionString,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // MemberData the private key-value pairs used by the member to communicate information to other members and the owner. Visible to all
// members of the lobby. At most 30 key-value pairs may be stored here, keys are limited to 30 characters and values to
// 1000. The total size of all memberData values may not exceed 4096 bytes.Keys are case sensitive.
    MemberData map[string]string `json:"MemberData,omitempty"`
    // MemberEntity the member entity who is joining the lobby.
    MemberEntity *EntityKeyModel `json:"MemberEntity,omitempty"`
}

// JoinLobbyResult 
type JoinLobbyResultModel struct {
    // LobbyId successfully joined lobby's id.
    LobbyId string `json:"LobbyId,omitempty"`
}

// JoinMatchmakingTicketRequest add the player to a matchmaking ticket and specify all of its matchmaking attributes. Players can join a ticket if and
// only if their EntityKeys are already listed in the ticket's Members list. The matchmaking service automatically starts
// matching the ticket against other matchmaking tickets once all players have joined the ticket. It is not possible to
// join a ticket once it has started matching.
type JoinMatchmakingTicketRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Member the User who wants to join the ticket. Their Id must be listed in PlayFabIdsToMatchWith.
    Member* MatchmakingPlayerModel `json:"Member,omitempty"`
    // QueueName the name of the queue to join.
    QueueName string `json:"QueueName,omitempty"`
    // TicketId the Id of the ticket to find a match for.
    TicketId string `json:"TicketId,omitempty"`
}

// JoinMatchmakingTicketResult 
type JoinMatchmakingTicketResultModel struct {
}

// LeaveLobbyRequest request to leave a lobby. Only a client can leave a lobby.
type LeaveLobbyRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // LobbyId the id of the lobby.
    LobbyId string `json:"LobbyId,omitempty"`
    // MemberEntity the member entity leaving the lobby.
    MemberEntity *EntityKeyModel `json:"MemberEntity,omitempty"`
}

// LinuxInstrumentationConfiguration 
type LinuxInstrumentationConfigurationModel struct {
    // IsEnabled designates whether Linux instrumentation configuration will be enabled for this Build
    IsEnabled bool `json:"IsEnabled"`
}

// ListAssetSummariesRequest returns a list of multiplayer server game asset summaries for a title.
type ListAssetSummariesRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
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

// ListBuildAliasesRequest returns a list of summarized details of all multiplayer server builds for a title.
type ListBuildAliasesRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // PageSize the page size for the request.
    PageSize int32 `json:"PageSize,omitempty"`
    // SkipToken the skip token for the paged request.
    SkipToken string `json:"SkipToken,omitempty"`
}

// ListBuildAliasesResponse 
type ListBuildAliasesResponseModel struct {
    // BuildAliases the list of build aliases for the title
    BuildAliases []BuildAliasDetailsResponseModel `json:"BuildAliases,omitempty"`
    // PageSize the page size on the response.
    PageSize int32 `json:"PageSize,omitempty"`
    // SkipToken the skip token for the paged response.
    SkipToken string `json:"SkipToken,omitempty"`
}

// ListBuildSummariesRequest returns a list of summarized details of all multiplayer server builds for a title.
type ListBuildSummariesRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
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
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
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
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
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
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // ImageName the container images we want to list tags for.
    ImageName string `json:"ImageName,omitempty"`
}

// ListContainerImageTagsResponse 
type ListContainerImageTagsResponseModel struct {
    // Tags the list of tags for a particular container image.
    Tags []string `json:"Tags,omitempty"`
}

// ListMatchmakingTicketsForPlayerRequest if the caller is a title, the EntityKey in the request is required. If the caller is a player, then it is optional. If
// it is provided it must match the caller's entity.
type ListMatchmakingTicketsForPlayerRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity key for which to find the ticket Ids.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // QueueName the name of the queue to find a match for.
    QueueName string `json:"QueueName,omitempty"`
}

// ListMatchmakingTicketsForPlayerResult 
type ListMatchmakingTicketsForPlayerResultModel struct {
    // TicketIds the list of ticket Ids the user is a member of.
    TicketIds []string `json:"TicketIds,omitempty"`
}

// ListMultiplayerServersRequest returns a list of multiplayer servers for a build in a specific region.
type ListMultiplayerServersRequestModel struct {
    // BuildId the guid string build ID of the multiplayer servers to list.
    BuildId string `json:"BuildId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // PageSize the page size for the request.
    PageSize int32 `json:"PageSize,omitempty"`
    // Region the region the multiplayer servers to list.
    Region string `json:"Region,omitempty"`
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

// ListPartyQosServersRequest returns a list of quality of service servers for party.
type ListPartyQosServersRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
}

// ListPartyQosServersResponse 
type ListPartyQosServersResponseModel struct {
    // PageSize the page size on the response.
    PageSize int32 `json:"PageSize,omitempty"`
    // QosServers the list of QoS servers.
    QosServers []QosServerModel `json:"QosServers,omitempty"`
    // SkipToken the skip token for the paged response.
    SkipToken string `json:"SkipToken,omitempty"`
}

// ListQosServersForTitleRequest returns a list of quality of service servers for a title.
type ListQosServersForTitleRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // IncludeAllRegions indicates that the response should contain Qos servers for all regions, including those where there are no builds
// deployed for the title.
    IncludeAllRegions bool `json:"IncludeAllRegions"`
}

// ListQosServersForTitleResponse 
type ListQosServersForTitleResponseModel struct {
    // PageSize the page size on the response.
    PageSize int32 `json:"PageSize,omitempty"`
    // QosServers the list of QoS servers.
    QosServers []QosServerModel `json:"QosServers,omitempty"`
    // SkipToken the skip token for the paged response.
    SkipToken string `json:"SkipToken,omitempty"`
}

// ListServerBackfillTicketsForPlayerRequest list all server backfill ticket Ids the user is a member of.
type ListServerBackfillTicketsForPlayerRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity key for which to find the ticket Ids.
    Entity* EntityKeyModel `json:"Entity,omitempty"`
    // QueueName the name of the queue the tickets are in.
    QueueName string `json:"QueueName,omitempty"`
}

// ListServerBackfillTicketsForPlayerResult 
type ListServerBackfillTicketsForPlayerResultModel struct {
    // TicketIds the list of backfill ticket Ids the user is a member of.
    TicketIds []string `json:"TicketIds,omitempty"`
}

// ListTitleMultiplayerServersQuotaChangesRequest list all server quota change requests for a title.
type ListTitleMultiplayerServersQuotaChangesRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
}

// ListTitleMultiplayerServersQuotaChangesResponse 
type ListTitleMultiplayerServersQuotaChangesResponseModel struct {
    // Changes all change requests for this title.
    Changes []QuotaChangeModel `json:"Changes,omitempty"`
}

// ListVirtualMachineSummariesRequest returns a list of virtual machines for a title.
type ListVirtualMachineSummariesRequestModel struct {
    // BuildId the guid string build ID of the virtual machines to list.
    BuildId string `json:"BuildId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // PageSize the page size for the request.
    PageSize int32 `json:"PageSize,omitempty"`
    // Region the region of the virtual machines to list.
    Region string `json:"Region,omitempty"`
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

// Lobby 
type LobbyModel struct {
    // AccessPolicy a setting indicating who is allowed to join this lobby, as well as see it in queries.
    AccessPolicy AccessPolicy `json:"AccessPolicy,omitempty"`
    // ChangeNumber a number that increments once for each request that modifies the lobby.
    ChangeNumber uint32 `json:"ChangeNumber,omitempty"`
    // ConnectionString a string used to join the lobby. This field is populated by the Lobby service. Invites are performed by communicating
// this connectionString to other players.
    ConnectionString string `json:"ConnectionString,omitempty"`
    // LobbyData lobby data.
    LobbyData map[string]string `json:"LobbyData,omitempty"`
    // LobbyId id to uniquely identify a lobby.
    LobbyId string `json:"LobbyId,omitempty"`
    // MaxPlayers the maximum number of players allowed in the lobby.
    MaxPlayers uint32 `json:"MaxPlayers,omitempty"`
    // Members array of all lobby members.
    Members []MemberModel `json:"Members,omitempty"`
    // MembershipLock a setting indicating whether members are allowed to join this lobby. When Locked new members are prevented from joining.
    MembershipLock MembershipLock `json:"MembershipLock,omitempty"`
    // Owner the client or server entity which owns this lobby.
    Owner *EntityKeyModel `json:"Owner,omitempty"`
    // OwnerMigrationPolicy a setting indicating the owner migration policy. If server owned, this field is not present.
    OwnerMigrationPolicy OwnerMigrationPolicy `json:"OwnerMigrationPolicy,omitempty"`
    // PubSubConnectionHandle an opaque string stored on a SubscribeToLobbyResource call, which indicates the connection an owner or member has with
// PubSub.
    PubSubConnectionHandle string `json:"PubSubConnectionHandle,omitempty"`
    // SearchData search data.
    SearchData map[string]string `json:"SearchData,omitempty"`
    // UseConnections a flag which determines if connections are used. Defaults to true. Only set on create.
    UseConnections bool `json:"UseConnections"`
}

// LobbyEmptyResult 
type LobbyEmptyResultModel struct {
}

// LobbySummary 
type LobbySummaryModel struct {
    // ConnectionString a string used to join the lobby.This field is populated by the Lobby service.Invites are performed by communicating this
// connectionString to other players.
    ConnectionString string `json:"ConnectionString,omitempty"`
    // CurrentPlayers the current number of players in the lobby.
    CurrentPlayers uint32 `json:"CurrentPlayers,omitempty"`
    // LobbyId id to uniquely identify a lobby.
    LobbyId string `json:"LobbyId,omitempty"`
    // MaxPlayers the maximum number of players allowed in the lobby.
    MaxPlayers uint32 `json:"MaxPlayers,omitempty"`
    // MembershipLock a setting indicating whether members are allowed to join this lobby. When Locked new members are prevented from joining.
    MembershipLock MembershipLock `json:"MembershipLock,omitempty"`
    // Owner the client or server entity which owns this lobby.
    Owner* EntityKeyModel `json:"Owner,omitempty"`
    // SearchData search data.
    SearchData map[string]string `json:"SearchData,omitempty"`
}

// MatchmakingPlayer a user in a matchmaking ticket.
type MatchmakingPlayerModel struct {
    // Attributes the user's attributes custom to the title.
    Attributes *MatchmakingPlayerAttributesModel `json:"Attributes,omitempty"`
    // Entity the entity key of the matchmaking user.
    Entity* EntityKeyModel `json:"Entity,omitempty"`
}

// MatchmakingPlayerAttributes the matchmaking attributes for a user.
type MatchmakingPlayerAttributesModel struct {
    // DataObject a data object representing a user's attributes.
    DataObject interface{} `json:"DataObject,omitempty"`
    // EscapedDataObject an escaped data object representing a user's attributes.
    EscapedDataObject string `json:"EscapedDataObject,omitempty"`
}

// MatchmakingPlayerWithTeamAssignment a player in a created matchmaking match with a team assignment.
type MatchmakingPlayerWithTeamAssignmentModel struct {
    // Attributes the user's attributes custom to the title. These attributes will be null unless the request has ReturnMemberAttributes
// flag set to true.
    Attributes *MatchmakingPlayerAttributesModel `json:"Attributes,omitempty"`
    // Entity the entity key of the matchmaking user.
    Entity* EntityKeyModel `json:"Entity,omitempty"`
    // TeamId the Id of the team the User is assigned to.
    TeamId string `json:"TeamId,omitempty"`
}

// Member 
type MemberModel struct {
    // MemberData key-value pairs specific to member.
    MemberData map[string]string `json:"MemberData,omitempty"`
    // MemberEntity the member entity key.
    MemberEntity *EntityKeyModel `json:"MemberEntity,omitempty"`
    // PubSubConnectionHandle opaque string, stored on a Subscribe call, which indicates the connection an owner or member has with PubSub.
    PubSubConnectionHandle string `json:"PubSubConnectionHandle,omitempty"`
}

// MembershipLock 
type MembershipLock string
  
const (
     MembershipLockUnlocked MembershipLock = "Unlocked"
     MembershipLockLocked MembershipLock = "Locked"
      )
// MonitoringApplicationConfiguration 
type MonitoringApplicationConfigurationModel struct {
    // AssetReference asset which contains the monitoring application files and scripts.
    AssetReference* AssetReferenceModel `json:"AssetReference,omitempty"`
    // ExecutionScriptName execution script name, this will be the main executable for the monitoring application.
    ExecutionScriptName string `json:"ExecutionScriptName,omitempty"`
    // InstallationScriptName installation script name, this will be run before the ExecutionScript.
    InstallationScriptName string `json:"InstallationScriptName,omitempty"`
    // OnStartRuntimeInMinutes timespan the monitoring application will be kept alive when running from the start of the VM
    OnStartRuntimeInMinutes float64 `json:"OnStartRuntimeInMinutes,omitempty"`
}

// MonitoringApplicationConfigurationParams 
type MonitoringApplicationConfigurationParamsModel struct {
    // AssetReference asset which contains the monitoring application files and scripts.
    AssetReference* AssetReferenceParamsModel `json:"AssetReference,omitempty"`
    // ExecutionScriptName execution script name, this will be the main executable for the monitoring application.
    ExecutionScriptName string `json:"ExecutionScriptName,omitempty"`
    // InstallationScriptName installation script name, this will be run before the ExecutionScript.
    InstallationScriptName string `json:"InstallationScriptName,omitempty"`
    // OnStartRuntimeInMinutes timespan the monitoring application will be kept alive when running from the start of the VM
    OnStartRuntimeInMinutes float64 `json:"OnStartRuntimeInMinutes,omitempty"`
}

// MultiplayerServerSummary 
type MultiplayerServerSummaryModel struct {
    // ConnectedPlayers the connected players in the multiplayer server.
    ConnectedPlayers []ConnectedPlayerModel `json:"ConnectedPlayers,omitempty"`
    // LastStateTransitionTime the time (UTC) at which a change in the multiplayer server state was observed.
    LastStateTransitionTime time.Time `json:"LastStateTransitionTime,omitempty"`
    // Region the region the multiplayer server is located in.
    Region string `json:"Region,omitempty"`
    // ServerId the string server ID of the multiplayer server generated by PlayFab.
    ServerId string `json:"ServerId,omitempty"`
    // SessionId the title generated guid string session ID of the multiplayer server.
    SessionId string `json:"SessionId,omitempty"`
    // State the state of the multiplayer server.
    State string `json:"State,omitempty"`
    // VmId the virtual machine ID that the multiplayer server is located on.
    VmId string `json:"VmId,omitempty"`
}

// OsPlatform 
type OsPlatform string
  
const (
     OsPlatformWindows OsPlatform = "Windows"
     OsPlatformLinux OsPlatform = "Linux"
      )
// OwnerMigrationPolicy 
type OwnerMigrationPolicy string
  
const (
     OwnerMigrationPolicyNone OwnerMigrationPolicy = "None"
     OwnerMigrationPolicyAutomatic OwnerMigrationPolicy = "Automatic"
     OwnerMigrationPolicyManual OwnerMigrationPolicy = "Manual"
     OwnerMigrationPolicyServer OwnerMigrationPolicy = "Server"
      )
// PaginationRequest 
type PaginationRequestModel struct {
    // ContinuationToken continuation token returned as a result in a previous FindLobbies call. Cannot be specified by clients.
    ContinuationToken string `json:"ContinuationToken,omitempty"`
    // PageSizeRequested the number of lobbies that should be retrieved. Cannot be specified by servers, clients may specify any value up to 50
    PageSizeRequested uint32 `json:"PageSizeRequested,omitempty"`
}

// PaginationResponse 
type PaginationResponseModel struct {
    // ContinuationToken continuation token returned by server call. Not returned for clients
    ContinuationToken string `json:"ContinuationToken,omitempty"`
    // TotalMatchedLobbyCount the number of lobbies that matched the search request.
    TotalMatchedLobbyCount uint32 `json:"TotalMatchedLobbyCount,omitempty"`
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
    Region string `json:"Region,omitempty"`
    // ServerUrl the QoS server URL.
    ServerUrl string `json:"ServerUrl,omitempty"`
}

// QuotaChange 
type QuotaChangeModel struct {
    // ChangeDescription a brief description of the requested changes.
    ChangeDescription string `json:"ChangeDescription,omitempty"`
    // Changes requested changes to make to the titles cores quota.
    Changes []CoreCapacityChangeModel `json:"Changes,omitempty"`
    // IsPendingReview whether or not this request is pending a review.
    IsPendingReview bool `json:"IsPendingReview"`
    // Notes additional information about this request that our team can use to better understand the requirements.
    Notes string `json:"Notes,omitempty"`
    // RequestId id of the change request.
    RequestId string `json:"RequestId,omitempty"`
    // ReviewComments comments by our team when a request is reviewed.
    ReviewComments string `json:"ReviewComments,omitempty"`
    // WasApproved whether or not this request was approved.
    WasApproved bool `json:"WasApproved"`
}

// RemoveMemberFromLobbyRequest request to remove a member from a lobby. Owners may remove other members from a lobby. Members cannot remove themselves
// (use LeaveLobby instead).
type RemoveMemberFromLobbyRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // LobbyId the id of the lobby.
    LobbyId string `json:"LobbyId,omitempty"`
    // MemberEntity the member entity to be removed from the lobby.
    MemberEntity *EntityKeyModel `json:"MemberEntity,omitempty"`
    // PreventRejoin if true, removed member can never rejoin this lobby.
    PreventRejoin bool `json:"PreventRejoin"`
}

// RequestMultiplayerServerRequest requests a multiplayer server session from a particular build in any of the given preferred regions.
type RequestMultiplayerServerRequestModel struct {
    // BuildAliasParams the identifiers of the build alias to use for the request.
    BuildAliasParams *BuildAliasParamsModel `json:"BuildAliasParams,omitempty"`
    // BuildId the guid string build ID of the multiplayer server to request.
    BuildId string `json:"BuildId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // InitialPlayers initial list of players (potentially matchmade) allowed to connect to the game. This list is passed to the game server
// when requested (via GSDK) and can be used to validate players connecting to it.
    InitialPlayers []string `json:"InitialPlayers,omitempty"`
    // PreferredRegions the preferred regions to request a multiplayer server from. The Multiplayer Service will iterate through the regions in
// the specified order and allocate a server from the first one that has servers available.
    PreferredRegions []string `json:"PreferredRegions,omitempty"`
    // SessionCookie data encoded as a string that is passed to the game server when requested. This can be used to to communicate
// information such as game mode or map through the request flow.
    SessionCookie string `json:"SessionCookie,omitempty"`
    // SessionId a guid string session ID created track the multiplayer server session over its life.
    SessionId string `json:"SessionId,omitempty"`
}

// RequestMultiplayerServerResponse 
type RequestMultiplayerServerResponseModel struct {
    // BuildId the identity of the build in which the server was allocated.
    BuildId string `json:"BuildId,omitempty"`
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
    Region string `json:"Region,omitempty"`
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
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
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

// Schedule 
type ScheduleModel struct {
    // Description a short description about this schedule. For example, "Game launch on July 15th".
    Description string `json:"Description,omitempty"`
    // EndTime the date and time in UTC at which the schedule ends. If IsRecurringWeekly is true, this schedule will keep renewing for
// future weeks until disabled or removed.
    EndTime time.Time `json:"EndTime,omitempty"`
    // IsDisabled disables the schedule.
    IsDisabled bool `json:"IsDisabled"`
    // IsRecurringWeekly if true, the StartTime and EndTime will get renewed every week.
    IsRecurringWeekly bool `json:"IsRecurringWeekly"`
    // StartTime the date and time in UTC at which the schedule starts.
    StartTime time.Time `json:"StartTime,omitempty"`
    // TargetStandby the standby target to maintain for the duration of the schedule.
    TargetStandby int32 `json:"TargetStandby,omitempty"`
}

// ScheduledStandbySettings 
type ScheduledStandbySettingsModel struct {
    // IsEnabled when true, scheduled standby will be enabled
    IsEnabled bool `json:"IsEnabled"`
    // ScheduleList a list of non-overlapping schedules
    ScheduleList []ScheduleModel `json:"ScheduleList,omitempty"`
}

// ServerDetails 
type ServerDetailsModel struct {
    // Fqdn the fully qualified domain name of the virtual machine that is hosting this multiplayer server.
    Fqdn string `json:"Fqdn,omitempty"`
    // IPV4Address the IPv4 address of the virtual machine that is hosting this multiplayer server.
    IPV4Address string `json:"IPV4Address,omitempty"`
    // Ports the ports the multiplayer server uses.
    Ports []PortModel `json:"Ports,omitempty"`
    // Region the server's region.
    Region string `json:"Region,omitempty"`
}

// ServerResourceConstraintParams 
type ServerResourceConstraintParamsModel struct {
    // CpuLimit the maximum number of cores that each server is allowed to use.
    CpuLimit float64 `json:"CpuLimit,omitempty"`
    // MemoryLimitGB the maximum number of GiB of memory that each server is allowed to use. WARNING: After exceeding this limit, the server
// will be killed
    MemoryLimitGB float64 `json:"MemoryLimitGB,omitempty"`
}

// ServerType 
type ServerType string
  
const (
     ServerTypeContainer ServerType = "Container"
     ServerTypeProcess ServerType = "Process"
      )
// ShutdownMultiplayerServerRequest executes the shutdown callback from the GSDK and terminates the multiplayer server session. The callback in the GSDK
// will allow for graceful shutdown with a 15 minute timeoutIf graceful shutdown has not been completed before 15 minutes
// have elapsed, the multiplayer server session will be forcefully terminated on it's own.
type ShutdownMultiplayerServerRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // SessionId a guid string session ID of the multiplayer server to shut down.
    SessionId string `json:"SessionId,omitempty"`
}

// Statistics 
type StatisticsModel struct {
    // Average the average.
    Average float64 `json:"Average,omitempty"`
    // Percentile50 the 50th percentile.
    Percentile50 float64 `json:"Percentile50,omitempty"`
    // Percentile90 the 90th percentile.
    Percentile90 float64 `json:"Percentile90,omitempty"`
    // Percentile99 the 99th percentile.
    Percentile99 float64 `json:"Percentile99,omitempty"`
}

// SubscribeToLobbyResourceRequest request to subscribe to lobby resource notifications.
type SubscribeToLobbyResourceRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // EntityKey the entity performing the subscription.
    EntityKey* EntityKeyModel `json:"EntityKey,omitempty"`
    // PubSubConnectionHandle opaque string, given to a client upon creating a connection with PubSub.
    PubSubConnectionHandle string `json:"PubSubConnectionHandle,omitempty"`
    // ResourceId the name of the resource to subscribe to.
    ResourceId string `json:"ResourceId,omitempty"`
    // SubscriptionVersion version number for the subscription of this resource.
    SubscriptionVersion uint32 `json:"SubscriptionVersion,omitempty"`
    // Type subscription type.
    Type SubscriptionType `json:"Type,omitempty"`
}

// SubscribeToLobbyResourceResult 
type SubscribeToLobbyResourceResultModel struct {
    // Topic topic will be returned in all notifications that are the result of this subscription.
    Topic string `json:"Topic,omitempty"`
}

// SubscriptionType 
type SubscriptionType string
  
const (
     SubscriptionTypeLobbyChange SubscriptionType = "LobbyChange"
     SubscriptionTypeLobbyInvite SubscriptionType = "LobbyInvite"
      )
// TitleMultiplayerServerEnabledStatus 
type TitleMultiplayerServerEnabledStatus string
  
const (
     TitleMultiplayerServerEnabledStatusInitializing TitleMultiplayerServerEnabledStatus = "Initializing"
     TitleMultiplayerServerEnabledStatusEnabled TitleMultiplayerServerEnabledStatus = "Enabled"
     TitleMultiplayerServerEnabledStatusDisabled TitleMultiplayerServerEnabledStatus = "Disabled"
      )
// TitleMultiplayerServersQuotas 
type TitleMultiplayerServersQuotasModel struct {
    // CoreCapacities the core capacity for the various regions and VM Family
    CoreCapacities []CoreCapacityModel `json:"CoreCapacities,omitempty"`
}

// UnsubscribeFromLobbyResourceRequest request to unsubscribe from lobby notifications.
type UnsubscribeFromLobbyResourceRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // EntityKey the entity which performed the subscription.
    EntityKey* EntityKeyModel `json:"EntityKey,omitempty"`
    // PubSubConnectionHandle opaque string, given to a client upon creating a connection with PubSub.
    PubSubConnectionHandle string `json:"PubSubConnectionHandle,omitempty"`
    // ResourceId the name of the resource to unsubscribe from.
    ResourceId string `json:"ResourceId,omitempty"`
    // SubscriptionVersion version number passed for the subscription of this resource.
    SubscriptionVersion uint32 `json:"SubscriptionVersion,omitempty"`
    // Type subscription type.
    Type SubscriptionType `json:"Type,omitempty"`
}

// UntagContainerImageRequest removes the specified tag from the image. After this operation, a 'docker pull' will fail for the specified image and
// tag combination. Morever, ListContainerImageTags will not return the specified tag.
type UntagContainerImageRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // ImageName the container image which tag we want to remove.
    ImageName string `json:"ImageName,omitempty"`
    // Tag the tag we want to remove.
    Tag string `json:"Tag,omitempty"`
}

// UpdateBuildAliasRequest creates a multiplayer server build alias and returns the created alias.
type UpdateBuildAliasRequestModel struct {
    // AliasId the guid string alias Id of the alias to be updated.
    AliasId string `json:"AliasId,omitempty"`
    // AliasName the alias name.
    AliasName string `json:"AliasName,omitempty"`
    // BuildSelectionCriteria array of build selection criteria.
    BuildSelectionCriteria []BuildSelectionCriterionModel `json:"BuildSelectionCriteria,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
}

// UpdateBuildNameRequest updates a multiplayer server build's name.
type UpdateBuildNameRequestModel struct {
    // BuildId the guid string ID of the build we want to update the name of.
    BuildId string `json:"BuildId,omitempty"`
    // BuildName the build name.
    BuildName string `json:"BuildName,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
}

// UpdateBuildRegionRequest updates a multiplayer server build's region.
type UpdateBuildRegionRequestModel struct {
    // BuildId the guid string ID of the build we want to update regions for.
    BuildId string `json:"BuildId,omitempty"`
    // BuildRegion the updated region configuration that should be applied to the specified build.
    BuildRegion* BuildRegionParamsModel `json:"BuildRegion,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
}

// UpdateBuildRegionsRequest updates a multiplayer server build's regions.
type UpdateBuildRegionsRequestModel struct {
    // BuildId the guid string ID of the build we want to update regions for.
    BuildId string `json:"BuildId,omitempty"`
    // BuildRegions the updated region configuration that should be applied to the specified build.
    BuildRegions []BuildRegionParamsModel `json:"BuildRegions,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
}

// UpdateLobbyRequest request to update a lobby.
type UpdateLobbyRequestModel struct {
    // AccessPolicy the policy indicating who is allowed to join the lobby, and the visibility to queries. May be 'Public', 'Friends' or
// 'Private'. Public means the lobby is both visible in queries and any player may join, including invited players. Friends
// means that users who are bidirectional friends of members in the lobby may search to find friend lobbies, to retrieve
// its connection string. Private means the lobby is not visible in queries, and a player must receive an invitation to
// join. Defaults to 'Public' on creation. Can only be changed by the lobby owner.
    AccessPolicy AccessPolicy `json:"AccessPolicy,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // LobbyData the private key-value pairs which are only visible to members of the lobby. Optional. Sets or updates key-value pairs on
// the lobby. Only the current lobby owner can set lobby data. Keys may be an arbitrary string of at most 30 characters.
// The total size of all lobbyData values may not exceed 4096 bytes. Values are not individually limited. There can be up
// to 30 key-value pairs stored here. Keys are case sensitive.
    LobbyData map[string]string `json:"LobbyData,omitempty"`
    // LobbyDataToDelete the keys to delete from the lobby LobbyData. Optional. Behaves similar to searchDataToDelete, but applies to lobbyData.
    LobbyDataToDelete []string `json:"LobbyDataToDelete,omitempty"`
    // LobbyId the id of the lobby.
    LobbyId string `json:"LobbyId,omitempty"`
    // MaxPlayers the maximum number of players allowed in the lobby. Updates the maximum allowed number of players in the lobby. Only the
// current lobby owner can set this. If set, the value must be greater than or equal to the number of members currently in
// the lobby.
    MaxPlayers uint32 `json:"MaxPlayers,omitempty"`
    // MemberData the private key-value pairs used by the member to communicate information to other members and the owner. Optional. Sets
// or updates new key-value pairs on the caller's member data. New keys will be added with their values and existing keys
// will be updated with the new values. Visible to all members of the lobby. At most 30 key-value pairs may be stored here,
// keys are limited to 30 characters and values to 1000. The total size of all memberData values may not exceed 4096 bytes.
// Keys are case sensitive. Servers cannot specifiy this.
    MemberData map[string]string `json:"MemberData,omitempty"`
    // MemberDataToDelete the keys to delete from the lobby MemberData. Optional. Deletes key-value pairs on the caller's member data. All the
// specified keys will be removed from the caller's member data. Keys that do not exist are a no-op. If the key to delete
// exists in the memberData (same request) it will result in a bad request. Servers cannot specifiy this.
    MemberDataToDelete []string `json:"MemberDataToDelete,omitempty"`
    // MemberEntity the member entity whose data is being modified. Servers cannot specify this.
    MemberEntity *EntityKeyModel `json:"MemberEntity,omitempty"`
    // MembershipLock a setting indicating whether the lobby is locked. May be 'Unlocked' or 'Locked'. When Locked new members are not allowed
// to join. Defaults to 'Unlocked' on creation. Can only be changed by the lobby owner.
    MembershipLock MembershipLock `json:"MembershipLock,omitempty"`
    // Owner the lobby owner. Optional. Set to transfer ownership of the lobby. If client - owned and 'Automatic' - The Lobby service
// will automatically assign another connected owner when the current owner leaves or disconnects. useConnections must be
// true. If client - owned and 'Manual' - Ownership is protected as long as the current owner is connected. If the current
// owner leaves or disconnects any member may set themselves as the current owner. The useConnections property must be
// true. If client-owned and 'None' - Any member can set ownership. The useConnections property can be either true or
// false. For all client-owned lobbies when the owner leaves and a new owner can not be automatically selected - The owner
// field is set to null. For all client-owned lobbies when the owner disconnects and a new owner can not be automatically
// selected - The owner field remains unchanged and the current owner retains all owner abilities for the lobby. If
// server-owned (must be 'Server') - Any server can set ownership. The useConnections property must be true.
    Owner *EntityKeyModel `json:"Owner,omitempty"`
    // SearchData the public key-value pairs which allow queries to differentiate between lobbies. Optional. Sets or updates key-value
// pairs on the lobby for use with queries. Only the current lobby owner can set search data. New keys will be added with
// their values and existing keys will be updated with the new values. There can be up to 30 key-value pairs stored here.
// Keys are of the format string_key1, string_key2... string_key30 for string values, or number_key1, number_key2, ...
// number_key30 for numeric values. Numeric values are floats. Values can be at most 256 characters long. The total size of
// all searchData values may not exceed 1024 bytes.Keys are case sensitive.
    SearchData map[string]string `json:"SearchData,omitempty"`
    // SearchDataToDelete the keys to delete from the lobby SearchData. Optional. Deletes key-value pairs on the lobby. Only the current lobby
// owner can delete search data. All the specified keys will be removed from the search data. Keys that do not exist in the
// lobby are a no-op.If the key to delete exists in the searchData (same request) it will result in a bad request.
    SearchDataToDelete []string `json:"SearchDataToDelete,omitempty"`
}

// UploadCertificateRequest uploads a multiplayer server game certificate.
type UploadCertificateRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // GameCertificate the game certificate to upload.
    GameCertificate* CertificateModel `json:"GameCertificate,omitempty"`
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

// WindowsCrashDumpConfiguration 
type WindowsCrashDumpConfigurationModel struct {
    // CustomDumpFlags see https://docs.microsoft.com/en-us/windows/win32/wer/collecting-user-mode-dumps for valid values.
    CustomDumpFlags int32 `json:"CustomDumpFlags,omitempty"`
    // DumpType see https://docs.microsoft.com/en-us/windows/win32/wer/collecting-user-mode-dumps for valid values.
    DumpType int32 `json:"DumpType,omitempty"`
    // IsEnabled designates whether automatic crash dump capturing will be enabled for this Build.
    IsEnabled bool `json:"IsEnabled"`
}
