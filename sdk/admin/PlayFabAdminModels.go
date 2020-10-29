package admin

import "time"
                    
// AbortTaskInstanceRequest if the task instance has already completed, there will be no-op.
type AbortTaskInstanceRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // TaskInstanceId iD of a task instance that is being aborted.
    TaskInstanceId string `json:"TaskInstanceId,omitempty"`
}

// ActionsOnPlayersInSegmentTaskParameter 
type ActionsOnPlayersInSegmentTaskParameterModel struct {
    // ActionId iD of the action to perform on each player in segment.
    ActionId string `json:"ActionId,omitempty"`
    // SegmentId iD of the segment to perform actions on.
    SegmentId string `json:"SegmentId,omitempty"`
}

// ActionsOnPlayersInSegmentTaskSummary 
type ActionsOnPlayersInSegmentTaskSummaryModel struct {
    // CompletedAt uTC timestamp when the task completed.
    CompletedAt time.Time `json:"CompletedAt,omitempty"`
    // ErrorMessage error message for last processing attempt, if an error occured.
    ErrorMessage string `json:"ErrorMessage,omitempty"`
    // ErrorWasFatal flag indicating if the error was fatal, if false job will be retried.
    ErrorWasFatal bool `json:"ErrorWasFatal"`
    // EstimatedSecondsRemaining estimated time remaining in seconds.
    EstimatedSecondsRemaining float64 `json:"EstimatedSecondsRemaining,omitempty"`
    // PercentComplete progress represented as percentage.
    PercentComplete float64 `json:"PercentComplete,omitempty"`
    // ScheduledByUserId if manually scheduled, ID of user who scheduled the task.
    ScheduledByUserId string `json:"ScheduledByUserId,omitempty"`
    // StartedAt uTC timestamp when the task started.
    StartedAt time.Time `json:"StartedAt,omitempty"`
    // Status current status of the task instance.
    Status TaskInstanceStatus `json:"Status,omitempty"`
    // TaskIdentifier identifier of the task this instance belongs to.
    TaskIdentifier *NameIdentifierModel `json:"TaskIdentifier,omitempty"`
    // TaskInstanceId iD of the task instance.
    TaskInstanceId string `json:"TaskInstanceId,omitempty"`
    // TotalPlayersInSegment total players in segment when task was started.
    TotalPlayersInSegment int32 `json:"TotalPlayersInSegment,omitempty"`
    // TotalPlayersProcessed total number of players that have had the actions applied to.
    TotalPlayersProcessed int32 `json:"TotalPlayersProcessed,omitempty"`
}

// AdCampaignAttribution 
type AdCampaignAttributionModel struct {
    // AttributedAt uTC time stamp of attribution
    AttributedAt time.Time `json:"AttributedAt,omitempty"`
    // CampaignId attribution campaign identifier
    CampaignId string `json:"CampaignId,omitempty"`
    // Platform attribution network name
    Platform string `json:"Platform,omitempty"`
}

// AdCampaignAttributionModel 
type AdCampaignAttributionModelModel struct {
    // AttributedAt uTC time stamp of attribution
    AttributedAt time.Time `json:"AttributedAt,omitempty"`
    // CampaignId attribution campaign identifier
    CampaignId string `json:"CampaignId,omitempty"`
    // Platform attribution network name
    Platform string `json:"Platform,omitempty"`
}

// AddLocalizedNewsRequest 
type AddLocalizedNewsRequestModel struct {
    // Body localized body text of the news.
    Body string `json:"Body,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Language language of the news item.
    Language string `json:"Language,omitempty"`
    // NewsId unique id of the updated news item.
    NewsId string `json:"NewsId,omitempty"`
    // Title localized title (headline) of the news item.
    Title string `json:"Title,omitempty"`
}

// AddLocalizedNewsResult 
type AddLocalizedNewsResultModel struct {
}

// AddNewsRequest 
type AddNewsRequestModel struct {
    // Body default body text of the news.
    Body string `json:"Body,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Timestamp time this news was published. If not set, defaults to now.
    Timestamp time.Time `json:"Timestamp,omitempty"`
    // Title default title (headline) of the news item.
    Title string `json:"Title,omitempty"`
}

// AddNewsResult 
type AddNewsResultModel struct {
    // NewsId unique id of the new news item
    NewsId string `json:"NewsId,omitempty"`
}

// AddPlayerTagRequest this API will trigger a player_tag_added event and add a tag with the given TagName and PlayFabID to the corresponding
// player profile. TagName can be used for segmentation and it is limited to 256 characters. Also there is a limit on the
// number of tags a title can have.
type AddPlayerTagRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
    // TagName unique tag for player profile.
    TagName string `json:"TagName,omitempty"`
}

// AddPlayerTagResult 
type AddPlayerTagResultModel struct {
}

// AddServerBuildRequest 
type AddServerBuildRequestModel struct {
    // ActiveRegions server host regions in which this build should be running and available
    ActiveRegions []Region `json:"ActiveRegions,omitempty"`
    // BuildId unique identifier for the build executable
    BuildId string `json:"BuildId,omitempty"`
    // CommandLineTemplate appended to the end of the command line when starting game servers
    CommandLineTemplate string `json:"CommandLineTemplate,omitempty"`
    // Comment developer comment(s) for this build
    Comment string `json:"Comment,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // ExecutablePath path to the game server executable. Defaults to gameserver.exe
    ExecutablePath string `json:"ExecutablePath,omitempty"`
    // MaxGamesPerHost maximum number of game server instances that can run on a single host machine
    MaxGamesPerHost int32 `json:"MaxGamesPerHost,omitempty"`
    // MinFreeGameSlots minimum capacity of additional game server instances that can be started before the autoscaling service starts new host
// machines (given the number of current running host machines and game server instances)
    MinFreeGameSlots int32 `json:"MinFreeGameSlots,omitempty"`
}

// AddServerBuildResult 
type AddServerBuildResultModel struct {
    // ActiveRegions array of regions where this build can used, when it is active
    ActiveRegions []Region `json:"ActiveRegions,omitempty"`
    // BuildId unique identifier for this build executable
    BuildId string `json:"BuildId,omitempty"`
    // CommandLineTemplate appended to the end of the command line when starting game servers
    CommandLineTemplate string `json:"CommandLineTemplate,omitempty"`
    // Comment developer comment(s) for this build
    Comment string `json:"Comment,omitempty"`
    // ExecutablePath path to the game server executable. Defaults to gameserver.exe
    ExecutablePath string `json:"ExecutablePath,omitempty"`
    // MaxGamesPerHost maximum number of game server instances that can run on a single host machine
    MaxGamesPerHost int32 `json:"MaxGamesPerHost,omitempty"`
    // MinFreeGameSlots minimum capacity of additional game server instances that can be started before the autoscaling service starts new host
// machines (given the number of current running host machines and game server instances)
    MinFreeGameSlots int32 `json:"MinFreeGameSlots,omitempty"`
    // Status the current status of the build validation and processing steps
    Status GameBuildStatus `json:"Status,omitempty"`
    // Timestamp time this build was last modified (or uploaded, if this build has never been modified)
    Timestamp time.Time `json:"Timestamp,omitempty"`
    // TitleId unique identifier for the title, found in the Settings > Game Properties section of the PlayFab developer site when a
// title has been selected.
    TitleId string `json:"TitleId,omitempty"`
}

// AddUserVirtualCurrencyRequest 
type AddUserVirtualCurrencyRequestModel struct {
    // Amount amount to be added to the user balance of the specified virtual currency. Maximum VC balance is Int32 (2,147,483,647).
// Any increase over this value will be discarded.
    Amount int32 `json:"Amount,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // PlayFabId playFab unique identifier of the user whose virtual currency balance is to be increased.
    PlayFabId string `json:"PlayFabId,omitempty"`
    // VirtualCurrency name of the virtual currency which is to be incremented.
    VirtualCurrency string `json:"VirtualCurrency,omitempty"`
}

// AddVirtualCurrencyTypesRequest this operation is additive. Any new currencies defined in the array will be added to the set of those available for the
// title, while any CurrencyCode identifiers matching existing ones in the game will be overwritten with the new values.
type AddVirtualCurrencyTypesRequestModel struct {
    // VirtualCurrencies list of virtual currencies and their initial deposits (the amount a user is granted when signing in for the first time)
// to the title
    VirtualCurrencies []VirtualCurrencyDataModel `json:"VirtualCurrencies,omitempty"`
}

// ApiCondition 
type ApiConditionModel struct {
    // HasSignatureOrEncryption require that API calls contain an RSA encrypted payload or signed headers.
    HasSignatureOrEncryption Conditionals `json:"HasSignatureOrEncryption,omitempty"`
}

// AuthTokenType 
type AuthTokenType string
  
const (
     AuthTokenTypeEmail AuthTokenType = "Email"
      )
// BanInfo contains information for a ban.
type BanInfoModel struct {
    // Active the active state of this ban. Expired bans may still have this value set to true but they will have no effect.
    Active bool `json:"Active"`
    // BanId the unique Ban Id associated with this ban.
    BanId string `json:"BanId,omitempty"`
    // Created the time when this ban was applied.
    Created time.Time `json:"Created,omitempty"`
    // Expires the time when this ban expires. Permanent bans do not have expiration date.
    Expires time.Time `json:"Expires,omitempty"`
    // IPAddress the IP address on which the ban was applied. May affect multiple players.
    IPAddress string `json:"IPAddress,omitempty"`
    // MACAddress the MAC address on which the ban was applied. May affect multiple players.
    MACAddress string `json:"MACAddress,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
    // Reason the reason why this ban was applied.
    Reason string `json:"Reason,omitempty"`
}

// BanRequest represents a single ban request.
type BanRequestModel struct {
    // DurationInHours the duration in hours for the ban. Leave this blank for a permanent ban.
    DurationInHours uint32 `json:"DurationInHours,omitempty"`
    // IPAddress iP address to be banned. May affect multiple players.
    IPAddress string `json:"IPAddress,omitempty"`
    // MACAddress mAC address to be banned. May affect multiple players.
    MACAddress string `json:"MACAddress,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
    // Reason the reason for this ban. Maximum 140 characters.
    Reason string `json:"Reason,omitempty"`
}

// BanUsersRequest the existence of each user will not be verified. When banning by IP or MAC address, multiple players may be affected, so
// use this feature with caution. Returns information about the new bans.
type BanUsersRequestModel struct {
    // Bans list of ban requests to be applied. Maximum 100.
    Bans []BanRequestModel `json:"Bans,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
}

// BanUsersResult 
type BanUsersResultModel struct {
    // BanData information on the bans that were applied
    BanData []BanInfoModel `json:"BanData,omitempty"`
}

// BlankResult 
type BlankResultModel struct {
}

// CatalogItem a purchasable item from the item catalog
type CatalogItemModel struct {
    // Bundle defines the bundle properties for the item - bundles are items which contain other items, including random drop tables
// and virtual currencies
    Bundle *CatalogItemBundleInfoModel `json:"Bundle,omitempty"`
    // CanBecomeCharacter if true, then an item instance of this type can be used to grant a character to a user.
    CanBecomeCharacter bool `json:"CanBecomeCharacter"`
    // CatalogVersion catalog version for this item
    CatalogVersion string `json:"CatalogVersion,omitempty"`
    // Consumable defines the consumable properties (number of uses, timeout) for the item
    Consumable *CatalogItemConsumableInfoModel `json:"Consumable,omitempty"`
    // Container defines the container properties for the item - what items it contains, including random drop tables and virtual
// currencies, and what item (if any) is required to open it via the UnlockContainerItem API
    Container *CatalogItemContainerInfoModel `json:"Container,omitempty"`
    // CustomData game specific custom data
    CustomData string `json:"CustomData,omitempty"`
    // Description text description of item, to show in-game
    Description string `json:"Description,omitempty"`
    // DisplayName text name for the item, to show in-game
    DisplayName string `json:"DisplayName,omitempty"`
    // InitialLimitedEditionCount if the item has IsLImitedEdition set to true, and this is the first time this ItemId has been defined as a limited
// edition item, this value determines the total number of instances to allocate for the title. Once this limit has been
// reached, no more instances of this ItemId can be created, and attempts to purchase or grant it will return a Result of
// false for that ItemId. If the item has already been defined to have a limited edition count, or if this value is less
// than zero, it will be ignored.
    InitialLimitedEditionCount int32 `json:"InitialLimitedEditionCount,omitempty"`
    // IsLimitedEdition bETA: If true, then only a fixed number can ever be granted.
    IsLimitedEdition bool `json:"IsLimitedEdition"`
    // IsStackable if true, then only one item instance of this type will exist and its remaininguses will be incremented instead.
// RemainingUses will cap out at Int32.Max (2,147,483,647). All subsequent increases will be discarded
    IsStackable bool `json:"IsStackable"`
    // IsTradable if true, then an item instance of this type can be traded between players using the trading APIs
    IsTradable bool `json:"IsTradable"`
    // ItemClass class to which the item belongs
    ItemClass string `json:"ItemClass,omitempty"`
    // ItemId unique identifier for this item
    ItemId string `json:"ItemId,omitempty"`
    // ItemImageUrl uRL to the item image. For Facebook purchase to display the image on the item purchase page, this must be set to an HTTP
// URL.
    ItemImageUrl string `json:"ItemImageUrl,omitempty"`
    // RealCurrencyPrices override prices for this item for specific currencies
    RealCurrencyPrices map[string]uint32 `json:"RealCurrencyPrices,omitempty"`
    // Tags list of item tags
    Tags []string `json:"Tags,omitempty"`
    // VirtualCurrencyPrices price of this item in virtual currencies and "RM" (the base Real Money purchase price, in USD pennies)
    VirtualCurrencyPrices map[string]uint32 `json:"VirtualCurrencyPrices,omitempty"`
}

// CatalogItemBundleInfo 
type CatalogItemBundleInfoModel struct {
    // BundledItems unique ItemId values for all items which will be added to the player inventory when the bundle is added
    BundledItems []string `json:"BundledItems,omitempty"`
    // BundledResultTables unique TableId values for all RandomResultTable objects which are part of the bundle (random tables will be resolved and
// add the relevant items to the player inventory when the bundle is added)
    BundledResultTables []string `json:"BundledResultTables,omitempty"`
    // BundledVirtualCurrencies virtual currency types and balances which will be added to the player inventory when the bundle is added
    BundledVirtualCurrencies map[string]uint32 `json:"BundledVirtualCurrencies,omitempty"`
}

// CatalogItemConsumableInfo 
type CatalogItemConsumableInfoModel struct {
    // UsageCount number of times this object can be used, after which it will be removed from the player inventory
    UsageCount uint32 `json:"UsageCount,omitempty"`
    // UsagePeriod duration in seconds for how long the item will remain in the player inventory - once elapsed, the item will be removed
// (recommended minimum value is 5 seconds, as lower values can cause the item to expire before operations depending on
// this item's details have completed)
    UsagePeriod uint32 `json:"UsagePeriod,omitempty"`
    // UsagePeriodGroup all inventory item instances in the player inventory sharing a non-null UsagePeriodGroup have their UsagePeriod values
// added together, and share the result - when that period has elapsed, all the items in the group will be removed
    UsagePeriodGroup string `json:"UsagePeriodGroup,omitempty"`
}

// CatalogItemContainerInfo containers are inventory items that can hold other items defined in the catalog, as well as virtual currency, which is
// added to the player inventory when the container is unlocked, using the UnlockContainerItem API. The items can be
// anything defined in the catalog, as well as RandomResultTable objects which will be resolved when the container is
// unlocked. Containers and their keys should be defined as Consumable (having a limited number of uses) in their catalog
// defintiions, unless the intent is for the player to be able to re-use them infinitely.
type CatalogItemContainerInfoModel struct {
    // ItemContents unique ItemId values for all items which will be added to the player inventory, once the container has been unlocked
    ItemContents []string `json:"ItemContents,omitempty"`
    // KeyItemId itemId for the catalog item used to unlock the container, if any (if not specified, a call to UnlockContainerItem will
// open the container, adding the contents to the player inventory and currency balances)
    KeyItemId string `json:"KeyItemId,omitempty"`
    // ResultTableContents unique TableId values for all RandomResultTable objects which are part of the container (once unlocked, random tables
// will be resolved and add the relevant items to the player inventory)
    ResultTableContents []string `json:"ResultTableContents,omitempty"`
    // VirtualCurrencyContents virtual currency types and balances which will be added to the player inventory when the container is unlocked
    VirtualCurrencyContents map[string]uint32 `json:"VirtualCurrencyContents,omitempty"`
}

// CheckLimitedEditionItemAvailabilityRequest this returns the total number of these items available.
type CheckLimitedEditionItemAvailabilityRequestModel struct {
    // CatalogVersion which catalog is being updated. If null, uses the default catalog.
    CatalogVersion string `json:"CatalogVersion,omitempty"`
    // ItemId the item to check for.
    ItemId string `json:"ItemId,omitempty"`
}

// CheckLimitedEditionItemAvailabilityResult 
type CheckLimitedEditionItemAvailabilityResultModel struct {
    // Amount the amount of the specified resource remaining.
    Amount int32 `json:"Amount,omitempty"`
}

// CloudScriptFile 
type CloudScriptFileModel struct {
    // FileContents contents of the Cloud Script javascript. Must be string-escaped javascript.
    FileContents string `json:"FileContents,omitempty"`
    // Filename name of the javascript file. These names are not used internally by the server, they are only for developer
// organizational purposes.
    Filename string `json:"Filename,omitempty"`
}

// CloudScriptTaskParameter 
type CloudScriptTaskParameterModel struct {
    // Argument argument to pass to the CloudScript function.
    Argument interface{} `json:"Argument,omitempty"`
    // FunctionName name of the CloudScript function to execute.
    FunctionName string `json:"FunctionName,omitempty"`
}

// CloudScriptTaskSummary 
type CloudScriptTaskSummaryModel struct {
    // CompletedAt uTC timestamp when the task completed.
    CompletedAt time.Time `json:"CompletedAt,omitempty"`
    // EstimatedSecondsRemaining estimated time remaining in seconds.
    EstimatedSecondsRemaining float64 `json:"EstimatedSecondsRemaining,omitempty"`
    // PercentComplete progress represented as percentage.
    PercentComplete float64 `json:"PercentComplete,omitempty"`
    // Result result of CloudScript execution
    Result *ExecuteCloudScriptResultModel `json:"Result,omitempty"`
    // ScheduledByUserId if manually scheduled, ID of user who scheduled the task.
    ScheduledByUserId string `json:"ScheduledByUserId,omitempty"`
    // StartedAt uTC timestamp when the task started.
    StartedAt time.Time `json:"StartedAt,omitempty"`
    // Status current status of the task instance.
    Status TaskInstanceStatus `json:"Status,omitempty"`
    // TaskIdentifier identifier of the task this instance belongs to.
    TaskIdentifier *NameIdentifierModel `json:"TaskIdentifier,omitempty"`
    // TaskInstanceId iD of the task instance.
    TaskInstanceId string `json:"TaskInstanceId,omitempty"`
}

// CloudScriptVersionStatus 
type CloudScriptVersionStatusModel struct {
    // LatestRevision most recent revision for this Cloud Script version
    LatestRevision int32 `json:"LatestRevision,omitempty"`
    // PublishedRevision published code revision for this Cloud Script version
    PublishedRevision int32 `json:"PublishedRevision,omitempty"`
    // Version version number
    Version int32 `json:"Version,omitempty"`
}

// Conditionals 
type Conditionals string
  
const (
     ConditionalsAny Conditionals = "Any"
     ConditionalsTrue Conditionals = "True"
     ConditionalsFalse Conditionals = "False"
      )
// ContactEmailInfo 
type ContactEmailInfoModel struct {
    // EmailAddress the email address
    EmailAddress string `json:"EmailAddress,omitempty"`
    // Name the name of the email info data
    Name string `json:"Name,omitempty"`
    // VerificationStatus the verification status of the email
    VerificationStatus EmailVerificationStatus `json:"VerificationStatus,omitempty"`
}

// ContactEmailInfoModel 
type ContactEmailInfoModelModel struct {
    // EmailAddress the email address
    EmailAddress string `json:"EmailAddress,omitempty"`
    // Name the name of the email info data
    Name string `json:"Name,omitempty"`
    // VerificationStatus the verification status of the email
    VerificationStatus EmailVerificationStatus `json:"VerificationStatus,omitempty"`
}

// ContentInfo 
type ContentInfoModel struct {
    // Key key of the content
    Key string `json:"Key,omitempty"`
    // LastModified last modified time
    LastModified time.Time `json:"LastModified,omitempty"`
    // Size size of the content in bytes
    Size uint32 `json:"Size,omitempty"`
}

// ContinentCode 
type ContinentCode string
  
const (
     ContinentCodeAF ContinentCode = "AF"
     ContinentCodeAN ContinentCode = "AN"
     ContinentCodeAS ContinentCode = "AS"
     ContinentCodeEU ContinentCode = "EU"
     ContinentCodeNA ContinentCode = "NA"
     ContinentCodeOC ContinentCode = "OC"
     ContinentCodeSA ContinentCode = "SA"
      )
// CountryCode 
type CountryCode string
  
const (
     CountryCodeAF CountryCode = "AF"
     CountryCodeAX CountryCode = "AX"
     CountryCodeAL CountryCode = "AL"
     CountryCodeDZ CountryCode = "DZ"
     CountryCodeAS CountryCode = "AS"
     CountryCodeAD CountryCode = "AD"
     CountryCodeAO CountryCode = "AO"
     CountryCodeAI CountryCode = "AI"
     CountryCodeAQ CountryCode = "AQ"
     CountryCodeAG CountryCode = "AG"
     CountryCodeAR CountryCode = "AR"
     CountryCodeAM CountryCode = "AM"
     CountryCodeAW CountryCode = "AW"
     CountryCodeAU CountryCode = "AU"
     CountryCodeAT CountryCode = "AT"
     CountryCodeAZ CountryCode = "AZ"
     CountryCodeBS CountryCode = "BS"
     CountryCodeBH CountryCode = "BH"
     CountryCodeBD CountryCode = "BD"
     CountryCodeBB CountryCode = "BB"
     CountryCodeBY CountryCode = "BY"
     CountryCodeBE CountryCode = "BE"
     CountryCodeBZ CountryCode = "BZ"
     CountryCodeBJ CountryCode = "BJ"
     CountryCodeBM CountryCode = "BM"
     CountryCodeBT CountryCode = "BT"
     CountryCodeBO CountryCode = "BO"
     CountryCodeBQ CountryCode = "BQ"
     CountryCodeBA CountryCode = "BA"
     CountryCodeBW CountryCode = "BW"
     CountryCodeBV CountryCode = "BV"
     CountryCodeBR CountryCode = "BR"
     CountryCodeIO CountryCode = "IO"
     CountryCodeBN CountryCode = "BN"
     CountryCodeBG CountryCode = "BG"
     CountryCodeBF CountryCode = "BF"
     CountryCodeBI CountryCode = "BI"
     CountryCodeKH CountryCode = "KH"
     CountryCodeCM CountryCode = "CM"
     CountryCodeCA CountryCode = "CA"
     CountryCodeCV CountryCode = "CV"
     CountryCodeKY CountryCode = "KY"
     CountryCodeCF CountryCode = "CF"
     CountryCodeTD CountryCode = "TD"
     CountryCodeCL CountryCode = "CL"
     CountryCodeCN CountryCode = "CN"
     CountryCodeCX CountryCode = "CX"
     CountryCodeCC CountryCode = "CC"
     CountryCodeCO CountryCode = "CO"
     CountryCodeKM CountryCode = "KM"
     CountryCodeCG CountryCode = "CG"
     CountryCodeCD CountryCode = "CD"
     CountryCodeCK CountryCode = "CK"
     CountryCodeCR CountryCode = "CR"
     CountryCodeCI CountryCode = "CI"
     CountryCodeHR CountryCode = "HR"
     CountryCodeCU CountryCode = "CU"
     CountryCodeCW CountryCode = "CW"
     CountryCodeCY CountryCode = "CY"
     CountryCodeCZ CountryCode = "CZ"
     CountryCodeDK CountryCode = "DK"
     CountryCodeDJ CountryCode = "DJ"
     CountryCodeDM CountryCode = "DM"
     CountryCodeDO CountryCode = "DO"
     CountryCodeEC CountryCode = "EC"
     CountryCodeEG CountryCode = "EG"
     CountryCodeSV CountryCode = "SV"
     CountryCodeGQ CountryCode = "GQ"
     CountryCodeER CountryCode = "ER"
     CountryCodeEE CountryCode = "EE"
     CountryCodeET CountryCode = "ET"
     CountryCodeFK CountryCode = "FK"
     CountryCodeFO CountryCode = "FO"
     CountryCodeFJ CountryCode = "FJ"
     CountryCodeFI CountryCode = "FI"
     CountryCodeFR CountryCode = "FR"
     CountryCodeGF CountryCode = "GF"
     CountryCodePF CountryCode = "PF"
     CountryCodeTF CountryCode = "TF"
     CountryCodeGA CountryCode = "GA"
     CountryCodeGM CountryCode = "GM"
     CountryCodeGE CountryCode = "GE"
     CountryCodeDE CountryCode = "DE"
     CountryCodeGH CountryCode = "GH"
     CountryCodeGI CountryCode = "GI"
     CountryCodeGR CountryCode = "GR"
     CountryCodeGL CountryCode = "GL"
     CountryCodeGD CountryCode = "GD"
     CountryCodeGP CountryCode = "GP"
     CountryCodeGU CountryCode = "GU"
     CountryCodeGT CountryCode = "GT"
     CountryCodeGG CountryCode = "GG"
     CountryCodeGN CountryCode = "GN"
     CountryCodeGW CountryCode = "GW"
     CountryCodeGY CountryCode = "GY"
     CountryCodeHT CountryCode = "HT"
     CountryCodeHM CountryCode = "HM"
     CountryCodeVA CountryCode = "VA"
     CountryCodeHN CountryCode = "HN"
     CountryCodeHK CountryCode = "HK"
     CountryCodeHU CountryCode = "HU"
     CountryCodeIS CountryCode = "IS"
     CountryCodeIN CountryCode = "IN"
     CountryCodeID CountryCode = "ID"
     CountryCodeIR CountryCode = "IR"
     CountryCodeIQ CountryCode = "IQ"
     CountryCodeIE CountryCode = "IE"
     CountryCodeIM CountryCode = "IM"
     CountryCodeIL CountryCode = "IL"
     CountryCodeIT CountryCode = "IT"
     CountryCodeJM CountryCode = "JM"
     CountryCodeJP CountryCode = "JP"
     CountryCodeJE CountryCode = "JE"
     CountryCodeJO CountryCode = "JO"
     CountryCodeKZ CountryCode = "KZ"
     CountryCodeKE CountryCode = "KE"
     CountryCodeKI CountryCode = "KI"
     CountryCodeKP CountryCode = "KP"
     CountryCodeKR CountryCode = "KR"
     CountryCodeKW CountryCode = "KW"
     CountryCodeKG CountryCode = "KG"
     CountryCodeLA CountryCode = "LA"
     CountryCodeLV CountryCode = "LV"
     CountryCodeLB CountryCode = "LB"
     CountryCodeLS CountryCode = "LS"
     CountryCodeLR CountryCode = "LR"
     CountryCodeLY CountryCode = "LY"
     CountryCodeLI CountryCode = "LI"
     CountryCodeLT CountryCode = "LT"
     CountryCodeLU CountryCode = "LU"
     CountryCodeMO CountryCode = "MO"
     CountryCodeMK CountryCode = "MK"
     CountryCodeMG CountryCode = "MG"
     CountryCodeMW CountryCode = "MW"
     CountryCodeMY CountryCode = "MY"
     CountryCodeMV CountryCode = "MV"
     CountryCodeML CountryCode = "ML"
     CountryCodeMT CountryCode = "MT"
     CountryCodeMH CountryCode = "MH"
     CountryCodeMQ CountryCode = "MQ"
     CountryCodeMR CountryCode = "MR"
     CountryCodeMU CountryCode = "MU"
     CountryCodeYT CountryCode = "YT"
     CountryCodeMX CountryCode = "MX"
     CountryCodeFM CountryCode = "FM"
     CountryCodeMD CountryCode = "MD"
     CountryCodeMC CountryCode = "MC"
     CountryCodeMN CountryCode = "MN"
     CountryCodeME CountryCode = "ME"
     CountryCodeMS CountryCode = "MS"
     CountryCodeMA CountryCode = "MA"
     CountryCodeMZ CountryCode = "MZ"
     CountryCodeMM CountryCode = "MM"
     CountryCodeNA CountryCode = "NA"
     CountryCodeNR CountryCode = "NR"
     CountryCodeNP CountryCode = "NP"
     CountryCodeNL CountryCode = "NL"
     CountryCodeNC CountryCode = "NC"
     CountryCodeNZ CountryCode = "NZ"
     CountryCodeNI CountryCode = "NI"
     CountryCodeNE CountryCode = "NE"
     CountryCodeNG CountryCode = "NG"
     CountryCodeNU CountryCode = "NU"
     CountryCodeNF CountryCode = "NF"
     CountryCodeMP CountryCode = "MP"
     CountryCodeNO CountryCode = "NO"
     CountryCodeOM CountryCode = "OM"
     CountryCodePK CountryCode = "PK"
     CountryCodePW CountryCode = "PW"
     CountryCodePS CountryCode = "PS"
     CountryCodePA CountryCode = "PA"
     CountryCodePG CountryCode = "PG"
     CountryCodePY CountryCode = "PY"
     CountryCodePE CountryCode = "PE"
     CountryCodePH CountryCode = "PH"
     CountryCodePN CountryCode = "PN"
     CountryCodePL CountryCode = "PL"
     CountryCodePT CountryCode = "PT"
     CountryCodePR CountryCode = "PR"
     CountryCodeQA CountryCode = "QA"
     CountryCodeRE CountryCode = "RE"
     CountryCodeRO CountryCode = "RO"
     CountryCodeRU CountryCode = "RU"
     CountryCodeRW CountryCode = "RW"
     CountryCodeBL CountryCode = "BL"
     CountryCodeSH CountryCode = "SH"
     CountryCodeKN CountryCode = "KN"
     CountryCodeLC CountryCode = "LC"
     CountryCodeMF CountryCode = "MF"
     CountryCodePM CountryCode = "PM"
     CountryCodeVC CountryCode = "VC"
     CountryCodeWS CountryCode = "WS"
     CountryCodeSM CountryCode = "SM"
     CountryCodeST CountryCode = "ST"
     CountryCodeSA CountryCode = "SA"
     CountryCodeSN CountryCode = "SN"
     CountryCodeRS CountryCode = "RS"
     CountryCodeSC CountryCode = "SC"
     CountryCodeSL CountryCode = "SL"
     CountryCodeSG CountryCode = "SG"
     CountryCodeSX CountryCode = "SX"
     CountryCodeSK CountryCode = "SK"
     CountryCodeSI CountryCode = "SI"
     CountryCodeSB CountryCode = "SB"
     CountryCodeSO CountryCode = "SO"
     CountryCodeZA CountryCode = "ZA"
     CountryCodeGS CountryCode = "GS"
     CountryCodeSS CountryCode = "SS"
     CountryCodeES CountryCode = "ES"
     CountryCodeLK CountryCode = "LK"
     CountryCodeSD CountryCode = "SD"
     CountryCodeSR CountryCode = "SR"
     CountryCodeSJ CountryCode = "SJ"
     CountryCodeSZ CountryCode = "SZ"
     CountryCodeSE CountryCode = "SE"
     CountryCodeCH CountryCode = "CH"
     CountryCodeSY CountryCode = "SY"
     CountryCodeTW CountryCode = "TW"
     CountryCodeTJ CountryCode = "TJ"
     CountryCodeTZ CountryCode = "TZ"
     CountryCodeTH CountryCode = "TH"
     CountryCodeTL CountryCode = "TL"
     CountryCodeTG CountryCode = "TG"
     CountryCodeTK CountryCode = "TK"
     CountryCodeTO CountryCode = "TO"
     CountryCodeTT CountryCode = "TT"
     CountryCodeTN CountryCode = "TN"
     CountryCodeTR CountryCode = "TR"
     CountryCodeTM CountryCode = "TM"
     CountryCodeTC CountryCode = "TC"
     CountryCodeTV CountryCode = "TV"
     CountryCodeUG CountryCode = "UG"
     CountryCodeUA CountryCode = "UA"
     CountryCodeAE CountryCode = "AE"
     CountryCodeGB CountryCode = "GB"
     CountryCodeUS CountryCode = "US"
     CountryCodeUM CountryCode = "UM"
     CountryCodeUY CountryCode = "UY"
     CountryCodeUZ CountryCode = "UZ"
     CountryCodeVU CountryCode = "VU"
     CountryCodeVE CountryCode = "VE"
     CountryCodeVN CountryCode = "VN"
     CountryCodeVG CountryCode = "VG"
     CountryCodeVI CountryCode = "VI"
     CountryCodeWF CountryCode = "WF"
     CountryCodeEH CountryCode = "EH"
     CountryCodeYE CountryCode = "YE"
     CountryCodeZM CountryCode = "ZM"
     CountryCodeZW CountryCode = "ZW"
      )
// CreateActionsOnPlayerSegmentTaskRequest task name is unique within a title. Using a task name that's already taken will cause a name conflict error. Too many
// create-task requests within a short time will cause a create conflict error.
type CreateActionsOnPlayerSegmentTaskRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Description description the task
    Description string `json:"Description,omitempty"`
    // IsActive whether the schedule is active. Inactive schedule will not trigger task execution.
    IsActive bool `json:"IsActive"`
    // Name name of the task. This is a unique identifier for tasks in the title.
    Name string `json:"Name,omitempty"`
    // Parameter task details related to segment and action
    Parameter* ActionsOnPlayersInSegmentTaskParameterModel `json:"Parameter,omitempty"`
    // Schedule cron expression for the run schedule of the task. The expression should be in UTC.
    Schedule string `json:"Schedule,omitempty"`
}

// CreateCloudScriptTaskRequest task name is unique within a title. Using a task name that's already taken will cause a name conflict error. Too many
// create-task requests within a short time will cause a create conflict error.
type CreateCloudScriptTaskRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Description description the task
    Description string `json:"Description,omitempty"`
    // IsActive whether the schedule is active. Inactive schedule will not trigger task execution.
    IsActive bool `json:"IsActive"`
    // Name name of the task. This is a unique identifier for tasks in the title.
    Name string `json:"Name,omitempty"`
    // Parameter task details related to CloudScript
    Parameter* CloudScriptTaskParameterModel `json:"Parameter,omitempty"`
    // Schedule cron expression for the run schedule of the task. The expression should be in UTC.
    Schedule string `json:"Schedule,omitempty"`
}

// CreateInsightsScheduledScalingTaskRequest task name is unique within a title. Using a task name that's already taken will cause a name conflict error. Too many
// create-task requests within a short time will cause a create conflict error.
type CreateInsightsScheduledScalingTaskRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Description description the task
    Description string `json:"Description,omitempty"`
    // IsActive whether the schedule is active. Inactive schedule will not trigger task execution.
    IsActive bool `json:"IsActive"`
    // Name name of the task. This is a unique identifier for tasks in the title.
    Name string `json:"Name,omitempty"`
    // Parameter task details related to Insights Scaling
    Parameter* InsightsScalingTaskParameterModel `json:"Parameter,omitempty"`
    // Schedule cron expression for the run schedule of the task. The expression should be in UTC.
    Schedule string `json:"Schedule,omitempty"`
}

// CreateOpenIdConnectionRequest 
type CreateOpenIdConnectionRequestModel struct {
    // ClientId the client ID given by the ID provider.
    ClientId string `json:"ClientId,omitempty"`
    // ClientSecret the client secret given by the ID provider.
    ClientSecret string `json:"ClientSecret,omitempty"`
    // ConnectionId a name for the connection that identifies it within the title.
    ConnectionId string `json:"ConnectionId,omitempty"`
    // IssuerDiscoveryUrl the discovery document URL to read issuer information from. This must be the absolute URL to the JSON OpenId
// Configuration document and must be accessible from the internet. If you don't know it, try your issuer URL followed by
// "/.well-known/openid-configuration". For example, if the issuer is https://example.com, try
// https://example.com/.well-known/openid-configuration
    IssuerDiscoveryUrl string `json:"IssuerDiscoveryUrl,omitempty"`
    // IssuerInformation manually specified information for an OpenID Connect issuer.
    IssuerInformation *OpenIdIssuerInformationModel `json:"IssuerInformation,omitempty"`
}

// CreatePlayerSharedSecretRequest player Shared Secret Keys are used for the call to Client/GetTitlePublicKey, which exchanges the shared secret for an
// RSA CSP blob to be used to encrypt the payload of account creation requests when that API requires a signature header.
type CreatePlayerSharedSecretRequestModel struct {
    // FriendlyName friendly name for this key
    FriendlyName string `json:"FriendlyName,omitempty"`
}

// CreatePlayerSharedSecretResult 
type CreatePlayerSharedSecretResultModel struct {
    // SecretKey the player shared secret to use when calling Client/GetTitlePublicKey
    SecretKey string `json:"SecretKey,omitempty"`
}

// CreatePlayerStatisticDefinitionRequest statistics are numeric values, with each statistic in the title also generating a leaderboard. The ResetInterval enables
// automatically resetting leaderboards on a specified interval. Upon reset, the statistic updates to a new version with no
// values (effectively removing all players from the leaderboard). The previous version's statistic values are also
// archived for retrieval, if needed (see GetPlayerStatisticVersions). Statistics not created via a call to
// CreatePlayerStatisticDefinition by default have a VersionChangeInterval of Never, meaning they do not reset on a
// schedule, but they can be set to do so via a call to UpdatePlayerStatisticDefinition. Once a statistic has been reset
// (sometimes referred to as versioned or incremented), the now-previous version can still be written to for up a short,
// pre-defined period (currently 10 seconds), to prevent issues with levels completing around the time of the reset. Also,
// once reset, the historical statistics for players in the title may be retrieved using the URL specified in the version
// information (GetPlayerStatisticVersions). The AggregationMethod determines what action is taken when a new statistic
// value is submitted - always update with the new value (Last), use the highest of the old and new values (Max), use the
// smallest (Min), or add them together (Sum).
type CreatePlayerStatisticDefinitionRequestModel struct {
    // AggregationMethod the aggregation method to use in updating the statistic (defaults to last)
    AggregationMethod StatisticAggregationMethod `json:"AggregationMethod,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // StatisticName unique name of the statistic
    StatisticName string `json:"StatisticName,omitempty"`
    // VersionChangeInterval interval at which the values of the statistic for all players are reset (resets begin at the next interval boundary)
    VersionChangeInterval StatisticResetIntervalOption `json:"VersionChangeInterval,omitempty"`
}

// CreatePlayerStatisticDefinitionResult 
type CreatePlayerStatisticDefinitionResultModel struct {
    // Statistic created statistic definition
    Statistic *PlayerStatisticDefinitionModel `json:"Statistic,omitempty"`
}

// CreateTaskResult 
type CreateTaskResultModel struct {
    // TaskId iD of the task
    TaskId string `json:"TaskId,omitempty"`
}

// Currency 
type Currency string
  
const (
     CurrencyAED Currency = "AED"
     CurrencyAFN Currency = "AFN"
     CurrencyALL Currency = "ALL"
     CurrencyAMD Currency = "AMD"
     CurrencyANG Currency = "ANG"
     CurrencyAOA Currency = "AOA"
     CurrencyARS Currency = "ARS"
     CurrencyAUD Currency = "AUD"
     CurrencyAWG Currency = "AWG"
     CurrencyAZN Currency = "AZN"
     CurrencyBAM Currency = "BAM"
     CurrencyBBD Currency = "BBD"
     CurrencyBDT Currency = "BDT"
     CurrencyBGN Currency = "BGN"
     CurrencyBHD Currency = "BHD"
     CurrencyBIF Currency = "BIF"
     CurrencyBMD Currency = "BMD"
     CurrencyBND Currency = "BND"
     CurrencyBOB Currency = "BOB"
     CurrencyBRL Currency = "BRL"
     CurrencyBSD Currency = "BSD"
     CurrencyBTN Currency = "BTN"
     CurrencyBWP Currency = "BWP"
     CurrencyBYR Currency = "BYR"
     CurrencyBZD Currency = "BZD"
     CurrencyCAD Currency = "CAD"
     CurrencyCDF Currency = "CDF"
     CurrencyCHF Currency = "CHF"
     CurrencyCLP Currency = "CLP"
     CurrencyCNY Currency = "CNY"
     CurrencyCOP Currency = "COP"
     CurrencyCRC Currency = "CRC"
     CurrencyCUC Currency = "CUC"
     CurrencyCUP Currency = "CUP"
     CurrencyCVE Currency = "CVE"
     CurrencyCZK Currency = "CZK"
     CurrencyDJF Currency = "DJF"
     CurrencyDKK Currency = "DKK"
     CurrencyDOP Currency = "DOP"
     CurrencyDZD Currency = "DZD"
     CurrencyEGP Currency = "EGP"
     CurrencyERN Currency = "ERN"
     CurrencyETB Currency = "ETB"
     CurrencyEUR Currency = "EUR"
     CurrencyFJD Currency = "FJD"
     CurrencyFKP Currency = "FKP"
     CurrencyGBP Currency = "GBP"
     CurrencyGEL Currency = "GEL"
     CurrencyGGP Currency = "GGP"
     CurrencyGHS Currency = "GHS"
     CurrencyGIP Currency = "GIP"
     CurrencyGMD Currency = "GMD"
     CurrencyGNF Currency = "GNF"
     CurrencyGTQ Currency = "GTQ"
     CurrencyGYD Currency = "GYD"
     CurrencyHKD Currency = "HKD"
     CurrencyHNL Currency = "HNL"
     CurrencyHRK Currency = "HRK"
     CurrencyHTG Currency = "HTG"
     CurrencyHUF Currency = "HUF"
     CurrencyIDR Currency = "IDR"
     CurrencyILS Currency = "ILS"
     CurrencyIMP Currency = "IMP"
     CurrencyINR Currency = "INR"
     CurrencyIQD Currency = "IQD"
     CurrencyIRR Currency = "IRR"
     CurrencyISK Currency = "ISK"
     CurrencyJEP Currency = "JEP"
     CurrencyJMD Currency = "JMD"
     CurrencyJOD Currency = "JOD"
     CurrencyJPY Currency = "JPY"
     CurrencyKES Currency = "KES"
     CurrencyKGS Currency = "KGS"
     CurrencyKHR Currency = "KHR"
     CurrencyKMF Currency = "KMF"
     CurrencyKPW Currency = "KPW"
     CurrencyKRW Currency = "KRW"
     CurrencyKWD Currency = "KWD"
     CurrencyKYD Currency = "KYD"
     CurrencyKZT Currency = "KZT"
     CurrencyLAK Currency = "LAK"
     CurrencyLBP Currency = "LBP"
     CurrencyLKR Currency = "LKR"
     CurrencyLRD Currency = "LRD"
     CurrencyLSL Currency = "LSL"
     CurrencyLYD Currency = "LYD"
     CurrencyMAD Currency = "MAD"
     CurrencyMDL Currency = "MDL"
     CurrencyMGA Currency = "MGA"
     CurrencyMKD Currency = "MKD"
     CurrencyMMK Currency = "MMK"
     CurrencyMNT Currency = "MNT"
     CurrencyMOP Currency = "MOP"
     CurrencyMRO Currency = "MRO"
     CurrencyMUR Currency = "MUR"
     CurrencyMVR Currency = "MVR"
     CurrencyMWK Currency = "MWK"
     CurrencyMXN Currency = "MXN"
     CurrencyMYR Currency = "MYR"
     CurrencyMZN Currency = "MZN"
     CurrencyNAD Currency = "NAD"
     CurrencyNGN Currency = "NGN"
     CurrencyNIO Currency = "NIO"
     CurrencyNOK Currency = "NOK"
     CurrencyNPR Currency = "NPR"
     CurrencyNZD Currency = "NZD"
     CurrencyOMR Currency = "OMR"
     CurrencyPAB Currency = "PAB"
     CurrencyPEN Currency = "PEN"
     CurrencyPGK Currency = "PGK"
     CurrencyPHP Currency = "PHP"
     CurrencyPKR Currency = "PKR"
     CurrencyPLN Currency = "PLN"
     CurrencyPYG Currency = "PYG"
     CurrencyQAR Currency = "QAR"
     CurrencyRON Currency = "RON"
     CurrencyRSD Currency = "RSD"
     CurrencyRUB Currency = "RUB"
     CurrencyRWF Currency = "RWF"
     CurrencySAR Currency = "SAR"
     CurrencySBD Currency = "SBD"
     CurrencySCR Currency = "SCR"
     CurrencySDG Currency = "SDG"
     CurrencySEK Currency = "SEK"
     CurrencySGD Currency = "SGD"
     CurrencySHP Currency = "SHP"
     CurrencySLL Currency = "SLL"
     CurrencySOS Currency = "SOS"
     CurrencySPL Currency = "SPL"
     CurrencySRD Currency = "SRD"
     CurrencySTD Currency = "STD"
     CurrencySVC Currency = "SVC"
     CurrencySYP Currency = "SYP"
     CurrencySZL Currency = "SZL"
     CurrencyTHB Currency = "THB"
     CurrencyTJS Currency = "TJS"
     CurrencyTMT Currency = "TMT"
     CurrencyTND Currency = "TND"
     CurrencyTOP Currency = "TOP"
     CurrencyTRY Currency = "TRY"
     CurrencyTTD Currency = "TTD"
     CurrencyTVD Currency = "TVD"
     CurrencyTWD Currency = "TWD"
     CurrencyTZS Currency = "TZS"
     CurrencyUAH Currency = "UAH"
     CurrencyUGX Currency = "UGX"
     CurrencyUSD Currency = "USD"
     CurrencyUYU Currency = "UYU"
     CurrencyUZS Currency = "UZS"
     CurrencyVEF Currency = "VEF"
     CurrencyVND Currency = "VND"
     CurrencyVUV Currency = "VUV"
     CurrencyWST Currency = "WST"
     CurrencyXAF Currency = "XAF"
     CurrencyXCD Currency = "XCD"
     CurrencyXDR Currency = "XDR"
     CurrencyXOF Currency = "XOF"
     CurrencyXPF Currency = "XPF"
     CurrencyYER Currency = "YER"
     CurrencyZAR Currency = "ZAR"
     CurrencyZMW Currency = "ZMW"
     CurrencyZWD Currency = "ZWD"
      )
// DeleteContentRequest 
type DeleteContentRequestModel struct {
    // Key key of the content item to be deleted
    Key string `json:"Key,omitempty"`
}

// DeleteMasterPlayerAccountRequest deletes all data associated with the master player account, including data from all titles the player has played, such
// as statistics, custom data, inventory, purchases, virtual currency balances, characters, group memberships, publisher
// data, credential data, account linkages, friends list and PlayStream event history. Removes the player from all
// leaderboards and player search indexes. Note, this API queues the player for deletion and returns a receipt immediately.
// Record the receipt ID for future reference. It may take some time before all player data is fully deleted. Upon
// completion of the deletion, an email will be sent to the notification email address configured for the title confirming
// the deletion. Until the player data is fully deleted, attempts to recreate the player with the same user account in the
// same title will fail with the 'AccountDeleted' error. It is highly recommended to know the impact of the deletion by
// calling GetPlayedTitleList, before calling this API.
type DeleteMasterPlayerAccountRequestModel struct {
    // MetaData developer created string to identify a user without PlayFab ID
    MetaData string `json:"MetaData,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// DeleteMasterPlayerAccountResult 
type DeleteMasterPlayerAccountResultModel struct {
    // JobReceiptId a notification email with this job receipt Id will be sent to the title notification email address when deletion is
// complete.
    JobReceiptId string `json:"JobReceiptId,omitempty"`
    // TitleIds list of titles from which the player's data will be deleted.
    TitleIds []string `json:"TitleIds,omitempty"`
}

// DeleteOpenIdConnectionRequest 
type DeleteOpenIdConnectionRequestModel struct {
    // ConnectionId unique name of the connection
    ConnectionId string `json:"ConnectionId,omitempty"`
}

// DeletePlayerRequest deletes all data associated with the player, including statistics, custom data, inventory, purchases, virtual currency
// balances, characters and shared group memberships. Removes the player from all leaderboards and player search indexes.
// Does not delete PlayStream event history associated with the player. Does not delete the publisher user account that
// created the player in the title nor associated data such as username, password, email address, account linkages, or
// friends list. Note, this API queues the player for deletion and returns immediately. It may take several minutes or more
// before all player data is fully deleted. Until the player data is fully deleted, attempts to recreate the player with
// the same user account in the same title will fail with the 'AccountDeleted' error.
type DeletePlayerRequestModel struct {
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// DeletePlayerResult 
type DeletePlayerResultModel struct {
}

// DeletePlayerSharedSecretRequest player Shared Secret Keys are used for the call to Client/GetTitlePublicKey, which exchanges the shared secret for an
// RSA CSP blob to be used to encrypt the payload of account creation requests when that API requires a signature header.
type DeletePlayerSharedSecretRequestModel struct {
    // SecretKey the shared secret key to delete
    SecretKey string `json:"SecretKey,omitempty"`
}

// DeletePlayerSharedSecretResult 
type DeletePlayerSharedSecretResultModel struct {
}

// DeleteStoreRequest this non-reversible operation will permanently delete the requested store.
type DeleteStoreRequestModel struct {
    // CatalogVersion catalog version of the store to delete. If null, uses the default catalog.
    CatalogVersion string `json:"CatalogVersion,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // StoreId unqiue identifier for the store which is to be deleted
    StoreId string `json:"StoreId,omitempty"`
}

// DeleteStoreResult 
type DeleteStoreResultModel struct {
}

// DeleteTaskRequest after a task is deleted, for tracking purposes, the task instances belonging to this task will still remain. They will
// become orphaned and does not belongs to any task. Executions of any in-progress task instances will continue. If the
// task specified does not exist, the deletion is considered a success.
type DeleteTaskRequestModel struct {
    // Identifier specify either the task ID or the name of task to be deleted.
    Identifier *NameIdentifierModel `json:"Identifier,omitempty"`
}

// DeleteTitleDataOverrideRequest will delete all the title data associated with the given override label.
type DeleteTitleDataOverrideRequestModel struct {
    // OverrideLabel name of the override.
    OverrideLabel string `json:"OverrideLabel,omitempty"`
}

// DeleteTitleDataOverrideResult 
type DeleteTitleDataOverrideResultModel struct {
}

// DeleteTitleRequest deletes all data associated with the title, including catalog, virtual currencies, leaderboard statistics, Cloud Script
// revisions, segment definitions, event rules, tasks, add-ons, secret keys, data encryption keys, and permission policies.
// Removes the title from its studio and removes all associated developer roles and permissions. Does not delete PlayStream
// event history associated with the title. Note, this API queues the title for deletion and returns immediately. It may
// take several hours or more before all title data is fully deleted. All player accounts in the title must be deleted
// before deleting the title. If any player accounts exist, the API will return a 'TitleContainsUserAccounts' error. Until
// the title data is fully deleted, attempts to call APIs with the title will fail with the 'TitleDeleted' error.
type DeleteTitleRequestModel struct {
}

// DeleteTitleResult 
type DeleteTitleResultModel struct {
}

// EffectType 
type EffectType string
  
const (
     EffectTypeAllow EffectType = "Allow"
     EffectTypeDeny EffectType = "Deny"
      )
// EmailVerificationStatus 
type EmailVerificationStatus string
  
const (
     EmailVerificationStatusUnverified EmailVerificationStatus = "Unverified"
     EmailVerificationStatusPending EmailVerificationStatus = "Pending"
     EmailVerificationStatusConfirmed EmailVerificationStatus = "Confirmed"
      )
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

// ExecuteCloudScriptResult 
type ExecuteCloudScriptResultModel struct {
    // APIRequestsIssued number of PlayFab API requests issued by the CloudScript function
    APIRequestsIssued int32 `json:"APIRequestsIssued,omitempty"`
    // Error information about the error, if any, that occurred during execution
    Error *ScriptExecutionErrorModel `json:"Error,omitempty"`
    // ExecutionTimeSeconds 
    ExecutionTimeSeconds float64 `json:"ExecutionTimeSeconds,omitempty"`
    // FunctionName the name of the function that executed
    FunctionName string `json:"FunctionName,omitempty"`
    // FunctionResult the object returned from the CloudScript function, if any
    FunctionResult interface{} `json:"FunctionResult,omitempty"`
    // FunctionResultTooLarge flag indicating if the FunctionResult was too large and was subsequently dropped from this event. This only occurs if
// the total event size is larger than 350KB.
    FunctionResultTooLarge bool `json:"FunctionResultTooLarge"`
    // HttpRequestsIssued number of external HTTP requests issued by the CloudScript function
    HttpRequestsIssued int32 `json:"HttpRequestsIssued,omitempty"`
    // Logs entries logged during the function execution. These include both entries logged in the function code using log.info()
// and log.error() and error entries for API and HTTP request failures.
    Logs []LogStatementModel `json:"Logs,omitempty"`
    // LogsTooLarge flag indicating if the logs were too large and were subsequently dropped from this event. This only occurs if the total
// event size is larger than 350KB after the FunctionResult was removed.
    LogsTooLarge bool `json:"LogsTooLarge"`
    // MemoryConsumedBytes 
    MemoryConsumedBytes uint32 `json:"MemoryConsumedBytes,omitempty"`
    // ProcessorTimeSeconds processor time consumed while executing the function. This does not include time spent waiting on API calls or HTTP
// requests.
    ProcessorTimeSeconds float64 `json:"ProcessorTimeSeconds,omitempty"`
    // Revision the revision of the CloudScript that executed
    Revision int32 `json:"Revision,omitempty"`
}

// ExportMasterPlayerDataRequest exports all data associated with the master player account, including data from all titles the player has played, such
// as statistics, custom data, inventory, purchases, virtual currency balances, characters, group memberships, publisher
// data, credential data, account linkages, friends list and PlayStream event history. Note, this API queues the player for
// export and returns a receipt immediately. Record the receipt ID for future reference. It may take some time before the
// export is available for download. Upon completion of the export, an email containing the URL to download the export dump
// will be sent to the notification email address configured for the title.
type ExportMasterPlayerDataRequestModel struct {
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// ExportMasterPlayerDataResult 
type ExportMasterPlayerDataResultModel struct {
    // JobReceiptId an email with this job receipt Id containing the export download link will be sent to the title notification email
// address when the export is complete.
    JobReceiptId string `json:"JobReceiptId,omitempty"`
}

// GameBuildStatus 
type GameBuildStatus string
  
const (
     GameBuildStatusAvailable GameBuildStatus = "Available"
     GameBuildStatusValidating GameBuildStatus = "Validating"
     GameBuildStatusInvalidBuildPackage GameBuildStatus = "InvalidBuildPackage"
     GameBuildStatusProcessing GameBuildStatus = "Processing"
     GameBuildStatusFailedToProcess GameBuildStatus = "FailedToProcess"
      )
// GameModeInfo 
type GameModeInfoModel struct {
    // Gamemode specific game mode type
    Gamemode string `json:"Gamemode,omitempty"`
    // MaxPlayerCount maximum user count a specific Game Server Instance can support
    MaxPlayerCount uint32 `json:"MaxPlayerCount,omitempty"`
    // MinPlayerCount minimum user count required for this Game Server Instance to continue (usually 1)
    MinPlayerCount uint32 `json:"MinPlayerCount,omitempty"`
    // StartOpen whether to start as an open session, meaning that players can matchmake into it (defaults to true)
    StartOpen bool `json:"StartOpen"`
}

// GenericErrorCodes 
type GenericErrorCodes string
  
const (
     GenericErrorCodesSuccess GenericErrorCodes = "Success"
     GenericErrorCodesUnkownError GenericErrorCodes = "UnkownError"
     GenericErrorCodesInvalidParams GenericErrorCodes = "InvalidParams"
     GenericErrorCodesAccountNotFound GenericErrorCodes = "AccountNotFound"
     GenericErrorCodesAccountBanned GenericErrorCodes = "AccountBanned"
     GenericErrorCodesInvalidUsernameOrPassword GenericErrorCodes = "InvalidUsernameOrPassword"
     GenericErrorCodesInvalidTitleId GenericErrorCodes = "InvalidTitleId"
     GenericErrorCodesInvalidEmailAddress GenericErrorCodes = "InvalidEmailAddress"
     GenericErrorCodesEmailAddressNotAvailable GenericErrorCodes = "EmailAddressNotAvailable"
     GenericErrorCodesInvalidUsername GenericErrorCodes = "InvalidUsername"
     GenericErrorCodesInvalidPassword GenericErrorCodes = "InvalidPassword"
     GenericErrorCodesUsernameNotAvailable GenericErrorCodes = "UsernameNotAvailable"
     GenericErrorCodesInvalidSteamTicket GenericErrorCodes = "InvalidSteamTicket"
     GenericErrorCodesAccountAlreadyLinked GenericErrorCodes = "AccountAlreadyLinked"
     GenericErrorCodesLinkedAccountAlreadyClaimed GenericErrorCodes = "LinkedAccountAlreadyClaimed"
     GenericErrorCodesInvalidFacebookToken GenericErrorCodes = "InvalidFacebookToken"
     GenericErrorCodesAccountNotLinked GenericErrorCodes = "AccountNotLinked"
     GenericErrorCodesFailedByPaymentProvider GenericErrorCodes = "FailedByPaymentProvider"
     GenericErrorCodesCouponCodeNotFound GenericErrorCodes = "CouponCodeNotFound"
     GenericErrorCodesInvalidContainerItem GenericErrorCodes = "InvalidContainerItem"
     GenericErrorCodesContainerNotOwned GenericErrorCodes = "ContainerNotOwned"
     GenericErrorCodesKeyNotOwned GenericErrorCodes = "KeyNotOwned"
     GenericErrorCodesInvalidItemIdInTable GenericErrorCodes = "InvalidItemIdInTable"
     GenericErrorCodesInvalidReceipt GenericErrorCodes = "InvalidReceipt"
     GenericErrorCodesReceiptAlreadyUsed GenericErrorCodes = "ReceiptAlreadyUsed"
     GenericErrorCodesReceiptCancelled GenericErrorCodes = "ReceiptCancelled"
     GenericErrorCodesGameNotFound GenericErrorCodes = "GameNotFound"
     GenericErrorCodesGameModeNotFound GenericErrorCodes = "GameModeNotFound"
     GenericErrorCodesInvalidGoogleToken GenericErrorCodes = "InvalidGoogleToken"
     GenericErrorCodesUserIsNotPartOfDeveloper GenericErrorCodes = "UserIsNotPartOfDeveloper"
     GenericErrorCodesInvalidTitleForDeveloper GenericErrorCodes = "InvalidTitleForDeveloper"
     GenericErrorCodesTitleNameConflicts GenericErrorCodes = "TitleNameConflicts"
     GenericErrorCodesUserisNotValid GenericErrorCodes = "UserisNotValid"
     GenericErrorCodesValueAlreadyExists GenericErrorCodes = "ValueAlreadyExists"
     GenericErrorCodesBuildNotFound GenericErrorCodes = "BuildNotFound"
     GenericErrorCodesPlayerNotInGame GenericErrorCodes = "PlayerNotInGame"
     GenericErrorCodesInvalidTicket GenericErrorCodes = "InvalidTicket"
     GenericErrorCodesInvalidDeveloper GenericErrorCodes = "InvalidDeveloper"
     GenericErrorCodesInvalidOrderInfo GenericErrorCodes = "InvalidOrderInfo"
     GenericErrorCodesRegistrationIncomplete GenericErrorCodes = "RegistrationIncomplete"
     GenericErrorCodesInvalidPlatform GenericErrorCodes = "InvalidPlatform"
     GenericErrorCodesUnknownError GenericErrorCodes = "UnknownError"
     GenericErrorCodesSteamApplicationNotOwned GenericErrorCodes = "SteamApplicationNotOwned"
     GenericErrorCodesWrongSteamAccount GenericErrorCodes = "WrongSteamAccount"
     GenericErrorCodesTitleNotActivated GenericErrorCodes = "TitleNotActivated"
     GenericErrorCodesRegistrationSessionNotFound GenericErrorCodes = "RegistrationSessionNotFound"
     GenericErrorCodesNoSuchMod GenericErrorCodes = "NoSuchMod"
     GenericErrorCodesFileNotFound GenericErrorCodes = "FileNotFound"
     GenericErrorCodesDuplicateEmail GenericErrorCodes = "DuplicateEmail"
     GenericErrorCodesItemNotFound GenericErrorCodes = "ItemNotFound"
     GenericErrorCodesItemNotOwned GenericErrorCodes = "ItemNotOwned"
     GenericErrorCodesItemNotRecycleable GenericErrorCodes = "ItemNotRecycleable"
     GenericErrorCodesItemNotAffordable GenericErrorCodes = "ItemNotAffordable"
     GenericErrorCodesInvalidVirtualCurrency GenericErrorCodes = "InvalidVirtualCurrency"
     GenericErrorCodesWrongVirtualCurrency GenericErrorCodes = "WrongVirtualCurrency"
     GenericErrorCodesWrongPrice GenericErrorCodes = "WrongPrice"
     GenericErrorCodesNonPositiveValue GenericErrorCodes = "NonPositiveValue"
     GenericErrorCodesInvalidRegion GenericErrorCodes = "InvalidRegion"
     GenericErrorCodesRegionAtCapacity GenericErrorCodes = "RegionAtCapacity"
     GenericErrorCodesServerFailedToStart GenericErrorCodes = "ServerFailedToStart"
     GenericErrorCodesNameNotAvailable GenericErrorCodes = "NameNotAvailable"
     GenericErrorCodesInsufficientFunds GenericErrorCodes = "InsufficientFunds"
     GenericErrorCodesInvalidDeviceID GenericErrorCodes = "InvalidDeviceID"
     GenericErrorCodesInvalidPushNotificationToken GenericErrorCodes = "InvalidPushNotificationToken"
     GenericErrorCodesNoRemainingUses GenericErrorCodes = "NoRemainingUses"
     GenericErrorCodesInvalidPaymentProvider GenericErrorCodes = "InvalidPaymentProvider"
     GenericErrorCodesPurchaseInitializationFailure GenericErrorCodes = "PurchaseInitializationFailure"
     GenericErrorCodesDuplicateUsername GenericErrorCodes = "DuplicateUsername"
     GenericErrorCodesInvalidBuyerInfo GenericErrorCodes = "InvalidBuyerInfo"
     GenericErrorCodesNoGameModeParamsSet GenericErrorCodes = "NoGameModeParamsSet"
     GenericErrorCodesBodyTooLarge GenericErrorCodes = "BodyTooLarge"
     GenericErrorCodesReservedWordInBody GenericErrorCodes = "ReservedWordInBody"
     GenericErrorCodesInvalidTypeInBody GenericErrorCodes = "InvalidTypeInBody"
     GenericErrorCodesInvalidRequest GenericErrorCodes = "InvalidRequest"
     GenericErrorCodesReservedEventName GenericErrorCodes = "ReservedEventName"
     GenericErrorCodesInvalidUserStatistics GenericErrorCodes = "InvalidUserStatistics"
     GenericErrorCodesNotAuthenticated GenericErrorCodes = "NotAuthenticated"
     GenericErrorCodesStreamAlreadyExists GenericErrorCodes = "StreamAlreadyExists"
     GenericErrorCodesErrorCreatingStream GenericErrorCodes = "ErrorCreatingStream"
     GenericErrorCodesStreamNotFound GenericErrorCodes = "StreamNotFound"
     GenericErrorCodesInvalidAccount GenericErrorCodes = "InvalidAccount"
     GenericErrorCodesPurchaseDoesNotExist GenericErrorCodes = "PurchaseDoesNotExist"
     GenericErrorCodesInvalidPurchaseTransactionStatus GenericErrorCodes = "InvalidPurchaseTransactionStatus"
     GenericErrorCodesAPINotEnabledForGameClientAccess GenericErrorCodes = "APINotEnabledForGameClientAccess"
     GenericErrorCodesNoPushNotificationARNForTitle GenericErrorCodes = "NoPushNotificationARNForTitle"
     GenericErrorCodesBuildAlreadyExists GenericErrorCodes = "BuildAlreadyExists"
     GenericErrorCodesBuildPackageDoesNotExist GenericErrorCodes = "BuildPackageDoesNotExist"
     GenericErrorCodesCustomAnalyticsEventsNotEnabledForTitle GenericErrorCodes = "CustomAnalyticsEventsNotEnabledForTitle"
     GenericErrorCodesInvalidSharedGroupId GenericErrorCodes = "InvalidSharedGroupId"
     GenericErrorCodesNotAuthorized GenericErrorCodes = "NotAuthorized"
     GenericErrorCodesMissingTitleGoogleProperties GenericErrorCodes = "MissingTitleGoogleProperties"
     GenericErrorCodesInvalidItemProperties GenericErrorCodes = "InvalidItemProperties"
     GenericErrorCodesInvalidPSNAuthCode GenericErrorCodes = "InvalidPSNAuthCode"
     GenericErrorCodesInvalidItemId GenericErrorCodes = "InvalidItemId"
     GenericErrorCodesPushNotEnabledForAccount GenericErrorCodes = "PushNotEnabledForAccount"
     GenericErrorCodesPushServiceError GenericErrorCodes = "PushServiceError"
     GenericErrorCodesReceiptDoesNotContainInAppItems GenericErrorCodes = "ReceiptDoesNotContainInAppItems"
     GenericErrorCodesReceiptContainsMultipleInAppItems GenericErrorCodes = "ReceiptContainsMultipleInAppItems"
     GenericErrorCodesInvalidBundleID GenericErrorCodes = "InvalidBundleID"
     GenericErrorCodesJavascriptException GenericErrorCodes = "JavascriptException"
     GenericErrorCodesInvalidSessionTicket GenericErrorCodes = "InvalidSessionTicket"
     GenericErrorCodesUnableToConnectToDatabase GenericErrorCodes = "UnableToConnectToDatabase"
     GenericErrorCodesInternalServerError GenericErrorCodes = "InternalServerError"
     GenericErrorCodesInvalidReportDate GenericErrorCodes = "InvalidReportDate"
     GenericErrorCodesReportNotAvailable GenericErrorCodes = "ReportNotAvailable"
     GenericErrorCodesDatabaseThroughputExceeded GenericErrorCodes = "DatabaseThroughputExceeded"
     GenericErrorCodesInvalidGameTicket GenericErrorCodes = "InvalidGameTicket"
     GenericErrorCodesExpiredGameTicket GenericErrorCodes = "ExpiredGameTicket"
     GenericErrorCodesGameTicketDoesNotMatchLobby GenericErrorCodes = "GameTicketDoesNotMatchLobby"
     GenericErrorCodesLinkedDeviceAlreadyClaimed GenericErrorCodes = "LinkedDeviceAlreadyClaimed"
     GenericErrorCodesDeviceAlreadyLinked GenericErrorCodes = "DeviceAlreadyLinked"
     GenericErrorCodesDeviceNotLinked GenericErrorCodes = "DeviceNotLinked"
     GenericErrorCodesPartialFailure GenericErrorCodes = "PartialFailure"
     GenericErrorCodesPublisherNotSet GenericErrorCodes = "PublisherNotSet"
     GenericErrorCodesServiceUnavailable GenericErrorCodes = "ServiceUnavailable"
     GenericErrorCodesVersionNotFound GenericErrorCodes = "VersionNotFound"
     GenericErrorCodesRevisionNotFound GenericErrorCodes = "RevisionNotFound"
     GenericErrorCodesInvalidPublisherId GenericErrorCodes = "InvalidPublisherId"
     GenericErrorCodesDownstreamServiceUnavailable GenericErrorCodes = "DownstreamServiceUnavailable"
     GenericErrorCodesAPINotIncludedInTitleUsageTier GenericErrorCodes = "APINotIncludedInTitleUsageTier"
     GenericErrorCodesDAULimitExceeded GenericErrorCodes = "DAULimitExceeded"
     GenericErrorCodesAPIRequestLimitExceeded GenericErrorCodes = "APIRequestLimitExceeded"
     GenericErrorCodesInvalidAPIEndpoint GenericErrorCodes = "InvalidAPIEndpoint"
     GenericErrorCodesBuildNotAvailable GenericErrorCodes = "BuildNotAvailable"
     GenericErrorCodesConcurrentEditError GenericErrorCodes = "ConcurrentEditError"
     GenericErrorCodesContentNotFound GenericErrorCodes = "ContentNotFound"
     GenericErrorCodesCharacterNotFound GenericErrorCodes = "CharacterNotFound"
     GenericErrorCodesCloudScriptNotFound GenericErrorCodes = "CloudScriptNotFound"
     GenericErrorCodesContentQuotaExceeded GenericErrorCodes = "ContentQuotaExceeded"
     GenericErrorCodesInvalidCharacterStatistics GenericErrorCodes = "InvalidCharacterStatistics"
     GenericErrorCodesPhotonNotEnabledForTitle GenericErrorCodes = "PhotonNotEnabledForTitle"
     GenericErrorCodesPhotonApplicationNotFound GenericErrorCodes = "PhotonApplicationNotFound"
     GenericErrorCodesPhotonApplicationNotAssociatedWithTitle GenericErrorCodes = "PhotonApplicationNotAssociatedWithTitle"
     GenericErrorCodesInvalidEmailOrPassword GenericErrorCodes = "InvalidEmailOrPassword"
     GenericErrorCodesFacebookAPIError GenericErrorCodes = "FacebookAPIError"
     GenericErrorCodesInvalidContentType GenericErrorCodes = "InvalidContentType"
     GenericErrorCodesKeyLengthExceeded GenericErrorCodes = "KeyLengthExceeded"
     GenericErrorCodesDataLengthExceeded GenericErrorCodes = "DataLengthExceeded"
     GenericErrorCodesTooManyKeys GenericErrorCodes = "TooManyKeys"
     GenericErrorCodesFreeTierCannotHaveVirtualCurrency GenericErrorCodes = "FreeTierCannotHaveVirtualCurrency"
     GenericErrorCodesMissingAmazonSharedKey GenericErrorCodes = "MissingAmazonSharedKey"
     GenericErrorCodesAmazonValidationError GenericErrorCodes = "AmazonValidationError"
     GenericErrorCodesInvalidPSNIssuerId GenericErrorCodes = "InvalidPSNIssuerId"
     GenericErrorCodesPSNInaccessible GenericErrorCodes = "PSNInaccessible"
     GenericErrorCodesExpiredAuthToken GenericErrorCodes = "ExpiredAuthToken"
     GenericErrorCodesFailedToGetEntitlements GenericErrorCodes = "FailedToGetEntitlements"
     GenericErrorCodesFailedToConsumeEntitlement GenericErrorCodes = "FailedToConsumeEntitlement"
     GenericErrorCodesTradeAcceptingUserNotAllowed GenericErrorCodes = "TradeAcceptingUserNotAllowed"
     GenericErrorCodesTradeInventoryItemIsAssignedToCharacter GenericErrorCodes = "TradeInventoryItemIsAssignedToCharacter"
     GenericErrorCodesTradeInventoryItemIsBundle GenericErrorCodes = "TradeInventoryItemIsBundle"
     GenericErrorCodesTradeStatusNotValidForCancelling GenericErrorCodes = "TradeStatusNotValidForCancelling"
     GenericErrorCodesTradeStatusNotValidForAccepting GenericErrorCodes = "TradeStatusNotValidForAccepting"
     GenericErrorCodesTradeDoesNotExist GenericErrorCodes = "TradeDoesNotExist"
     GenericErrorCodesTradeCancelled GenericErrorCodes = "TradeCancelled"
     GenericErrorCodesTradeAlreadyFilled GenericErrorCodes = "TradeAlreadyFilled"
     GenericErrorCodesTradeWaitForStatusTimeout GenericErrorCodes = "TradeWaitForStatusTimeout"
     GenericErrorCodesTradeInventoryItemExpired GenericErrorCodes = "TradeInventoryItemExpired"
     GenericErrorCodesTradeMissingOfferedAndAcceptedItems GenericErrorCodes = "TradeMissingOfferedAndAcceptedItems"
     GenericErrorCodesTradeAcceptedItemIsBundle GenericErrorCodes = "TradeAcceptedItemIsBundle"
     GenericErrorCodesTradeAcceptedItemIsStackable GenericErrorCodes = "TradeAcceptedItemIsStackable"
     GenericErrorCodesTradeInventoryItemInvalidStatus GenericErrorCodes = "TradeInventoryItemInvalidStatus"
     GenericErrorCodesTradeAcceptedCatalogItemInvalid GenericErrorCodes = "TradeAcceptedCatalogItemInvalid"
     GenericErrorCodesTradeAllowedUsersInvalid GenericErrorCodes = "TradeAllowedUsersInvalid"
     GenericErrorCodesTradeInventoryItemDoesNotExist GenericErrorCodes = "TradeInventoryItemDoesNotExist"
     GenericErrorCodesTradeInventoryItemIsConsumed GenericErrorCodes = "TradeInventoryItemIsConsumed"
     GenericErrorCodesTradeInventoryItemIsStackable GenericErrorCodes = "TradeInventoryItemIsStackable"
     GenericErrorCodesTradeAcceptedItemsMismatch GenericErrorCodes = "TradeAcceptedItemsMismatch"
     GenericErrorCodesInvalidKongregateToken GenericErrorCodes = "InvalidKongregateToken"
     GenericErrorCodesFeatureNotConfiguredForTitle GenericErrorCodes = "FeatureNotConfiguredForTitle"
     GenericErrorCodesNoMatchingCatalogItemForReceipt GenericErrorCodes = "NoMatchingCatalogItemForReceipt"
     GenericErrorCodesInvalidCurrencyCode GenericErrorCodes = "InvalidCurrencyCode"
     GenericErrorCodesNoRealMoneyPriceForCatalogItem GenericErrorCodes = "NoRealMoneyPriceForCatalogItem"
     GenericErrorCodesTradeInventoryItemIsNotTradable GenericErrorCodes = "TradeInventoryItemIsNotTradable"
     GenericErrorCodesTradeAcceptedCatalogItemIsNotTradable GenericErrorCodes = "TradeAcceptedCatalogItemIsNotTradable"
     GenericErrorCodesUsersAlreadyFriends GenericErrorCodes = "UsersAlreadyFriends"
     GenericErrorCodesLinkedIdentifierAlreadyClaimed GenericErrorCodes = "LinkedIdentifierAlreadyClaimed"
     GenericErrorCodesCustomIdNotLinked GenericErrorCodes = "CustomIdNotLinked"
     GenericErrorCodesTotalDataSizeExceeded GenericErrorCodes = "TotalDataSizeExceeded"
     GenericErrorCodesDeleteKeyConflict GenericErrorCodes = "DeleteKeyConflict"
     GenericErrorCodesInvalidXboxLiveToken GenericErrorCodes = "InvalidXboxLiveToken"
     GenericErrorCodesExpiredXboxLiveToken GenericErrorCodes = "ExpiredXboxLiveToken"
     GenericErrorCodesResettableStatisticVersionRequired GenericErrorCodes = "ResettableStatisticVersionRequired"
     GenericErrorCodesNotAuthorizedByTitle GenericErrorCodes = "NotAuthorizedByTitle"
     GenericErrorCodesNoPartnerEnabled GenericErrorCodes = "NoPartnerEnabled"
     GenericErrorCodesInvalidPartnerResponse GenericErrorCodes = "InvalidPartnerResponse"
     GenericErrorCodesAPINotEnabledForGameServerAccess GenericErrorCodes = "APINotEnabledForGameServerAccess"
     GenericErrorCodesStatisticNotFound GenericErrorCodes = "StatisticNotFound"
     GenericErrorCodesStatisticNameConflict GenericErrorCodes = "StatisticNameConflict"
     GenericErrorCodesStatisticVersionClosedForWrites GenericErrorCodes = "StatisticVersionClosedForWrites"
     GenericErrorCodesStatisticVersionInvalid GenericErrorCodes = "StatisticVersionInvalid"
     GenericErrorCodesAPIClientRequestRateLimitExceeded GenericErrorCodes = "APIClientRequestRateLimitExceeded"
     GenericErrorCodesInvalidJSONContent GenericErrorCodes = "InvalidJSONContent"
     GenericErrorCodesInvalidDropTable GenericErrorCodes = "InvalidDropTable"
     GenericErrorCodesStatisticVersionAlreadyIncrementedForScheduledInterval GenericErrorCodes = "StatisticVersionAlreadyIncrementedForScheduledInterval"
     GenericErrorCodesStatisticCountLimitExceeded GenericErrorCodes = "StatisticCountLimitExceeded"
     GenericErrorCodesStatisticVersionIncrementRateExceeded GenericErrorCodes = "StatisticVersionIncrementRateExceeded"
     GenericErrorCodesContainerKeyInvalid GenericErrorCodes = "ContainerKeyInvalid"
     GenericErrorCodesCloudScriptExecutionTimeLimitExceeded GenericErrorCodes = "CloudScriptExecutionTimeLimitExceeded"
     GenericErrorCodesNoWritePermissionsForEvent GenericErrorCodes = "NoWritePermissionsForEvent"
     GenericErrorCodesCloudScriptFunctionArgumentSizeExceeded GenericErrorCodes = "CloudScriptFunctionArgumentSizeExceeded"
     GenericErrorCodesCloudScriptAPIRequestCountExceeded GenericErrorCodes = "CloudScriptAPIRequestCountExceeded"
     GenericErrorCodesCloudScriptAPIRequestError GenericErrorCodes = "CloudScriptAPIRequestError"
     GenericErrorCodesCloudScriptHTTPRequestError GenericErrorCodes = "CloudScriptHTTPRequestError"
     GenericErrorCodesInsufficientGuildRole GenericErrorCodes = "InsufficientGuildRole"
     GenericErrorCodesGuildNotFound GenericErrorCodes = "GuildNotFound"
     GenericErrorCodesOverLimit GenericErrorCodes = "OverLimit"
     GenericErrorCodesEventNotFound GenericErrorCodes = "EventNotFound"
     GenericErrorCodesInvalidEventField GenericErrorCodes = "InvalidEventField"
     GenericErrorCodesInvalidEventName GenericErrorCodes = "InvalidEventName"
     GenericErrorCodesCatalogNotConfigured GenericErrorCodes = "CatalogNotConfigured"
     GenericErrorCodesOperationNotSupportedForPlatform GenericErrorCodes = "OperationNotSupportedForPlatform"
     GenericErrorCodesSegmentNotFound GenericErrorCodes = "SegmentNotFound"
     GenericErrorCodesStoreNotFound GenericErrorCodes = "StoreNotFound"
     GenericErrorCodesInvalidStatisticName GenericErrorCodes = "InvalidStatisticName"
     GenericErrorCodesTitleNotQualifiedForLimit GenericErrorCodes = "TitleNotQualifiedForLimit"
     GenericErrorCodesInvalidServiceLimitLevel GenericErrorCodes = "InvalidServiceLimitLevel"
     GenericErrorCodesServiceLimitLevelInTransition GenericErrorCodes = "ServiceLimitLevelInTransition"
     GenericErrorCodesCouponAlreadyRedeemed GenericErrorCodes = "CouponAlreadyRedeemed"
     GenericErrorCodesGameServerBuildSizeLimitExceeded GenericErrorCodes = "GameServerBuildSizeLimitExceeded"
     GenericErrorCodesGameServerBuildCountLimitExceeded GenericErrorCodes = "GameServerBuildCountLimitExceeded"
     GenericErrorCodesVirtualCurrencyCountLimitExceeded GenericErrorCodes = "VirtualCurrencyCountLimitExceeded"
     GenericErrorCodesVirtualCurrencyCodeExists GenericErrorCodes = "VirtualCurrencyCodeExists"
     GenericErrorCodesTitleNewsItemCountLimitExceeded GenericErrorCodes = "TitleNewsItemCountLimitExceeded"
     GenericErrorCodesInvalidTwitchToken GenericErrorCodes = "InvalidTwitchToken"
     GenericErrorCodesTwitchResponseError GenericErrorCodes = "TwitchResponseError"
     GenericErrorCodesProfaneDisplayName GenericErrorCodes = "ProfaneDisplayName"
     GenericErrorCodesUserAlreadyAdded GenericErrorCodes = "UserAlreadyAdded"
     GenericErrorCodesInvalidVirtualCurrencyCode GenericErrorCodes = "InvalidVirtualCurrencyCode"
     GenericErrorCodesVirtualCurrencyCannotBeDeleted GenericErrorCodes = "VirtualCurrencyCannotBeDeleted"
     GenericErrorCodesIdentifierAlreadyClaimed GenericErrorCodes = "IdentifierAlreadyClaimed"
     GenericErrorCodesIdentifierNotLinked GenericErrorCodes = "IdentifierNotLinked"
     GenericErrorCodesInvalidContinuationToken GenericErrorCodes = "InvalidContinuationToken"
     GenericErrorCodesExpiredContinuationToken GenericErrorCodes = "ExpiredContinuationToken"
     GenericErrorCodesInvalidSegment GenericErrorCodes = "InvalidSegment"
     GenericErrorCodesInvalidSessionId GenericErrorCodes = "InvalidSessionId"
     GenericErrorCodesSessionLogNotFound GenericErrorCodes = "SessionLogNotFound"
     GenericErrorCodesInvalidSearchTerm GenericErrorCodes = "InvalidSearchTerm"
     GenericErrorCodesTwoFactorAuthenticationTokenRequired GenericErrorCodes = "TwoFactorAuthenticationTokenRequired"
     GenericErrorCodesGameServerHostCountLimitExceeded GenericErrorCodes = "GameServerHostCountLimitExceeded"
     GenericErrorCodesPlayerTagCountLimitExceeded GenericErrorCodes = "PlayerTagCountLimitExceeded"
     GenericErrorCodesRequestAlreadyRunning GenericErrorCodes = "RequestAlreadyRunning"
     GenericErrorCodesActionGroupNotFound GenericErrorCodes = "ActionGroupNotFound"
     GenericErrorCodesMaximumSegmentBulkActionJobsRunning GenericErrorCodes = "MaximumSegmentBulkActionJobsRunning"
     GenericErrorCodesNoActionsOnPlayersInSegmentJob GenericErrorCodes = "NoActionsOnPlayersInSegmentJob"
     GenericErrorCodesDuplicateStatisticName GenericErrorCodes = "DuplicateStatisticName"
     GenericErrorCodesScheduledTaskNameConflict GenericErrorCodes = "ScheduledTaskNameConflict"
     GenericErrorCodesScheduledTaskCreateConflict GenericErrorCodes = "ScheduledTaskCreateConflict"
     GenericErrorCodesInvalidScheduledTaskName GenericErrorCodes = "InvalidScheduledTaskName"
     GenericErrorCodesInvalidTaskSchedule GenericErrorCodes = "InvalidTaskSchedule"
     GenericErrorCodesSteamNotEnabledForTitle GenericErrorCodes = "SteamNotEnabledForTitle"
     GenericErrorCodesLimitNotAnUpgradeOption GenericErrorCodes = "LimitNotAnUpgradeOption"
     GenericErrorCodesNoSecretKeyEnabledForCloudScript GenericErrorCodes = "NoSecretKeyEnabledForCloudScript"
     GenericErrorCodesTaskNotFound GenericErrorCodes = "TaskNotFound"
     GenericErrorCodesTaskInstanceNotFound GenericErrorCodes = "TaskInstanceNotFound"
     GenericErrorCodesInvalidIdentityProviderId GenericErrorCodes = "InvalidIdentityProviderId"
     GenericErrorCodesMisconfiguredIdentityProvider GenericErrorCodes = "MisconfiguredIdentityProvider"
     GenericErrorCodesInvalidScheduledTaskType GenericErrorCodes = "InvalidScheduledTaskType"
     GenericErrorCodesBillingInformationRequired GenericErrorCodes = "BillingInformationRequired"
     GenericErrorCodesLimitedEditionItemUnavailable GenericErrorCodes = "LimitedEditionItemUnavailable"
     GenericErrorCodesInvalidAdPlacementAndReward GenericErrorCodes = "InvalidAdPlacementAndReward"
     GenericErrorCodesAllAdPlacementViewsAlreadyConsumed GenericErrorCodes = "AllAdPlacementViewsAlreadyConsumed"
     GenericErrorCodesGoogleOAuthNotConfiguredForTitle GenericErrorCodes = "GoogleOAuthNotConfiguredForTitle"
     GenericErrorCodesGoogleOAuthError GenericErrorCodes = "GoogleOAuthError"
     GenericErrorCodesUserNotFriend GenericErrorCodes = "UserNotFriend"
     GenericErrorCodesInvalidSignature GenericErrorCodes = "InvalidSignature"
     GenericErrorCodesInvalidPublicKey GenericErrorCodes = "InvalidPublicKey"
     GenericErrorCodesGoogleOAuthNoIdTokenIncludedInResponse GenericErrorCodes = "GoogleOAuthNoIdTokenIncludedInResponse"
     GenericErrorCodesStatisticUpdateInProgress GenericErrorCodes = "StatisticUpdateInProgress"
     GenericErrorCodesLeaderboardVersionNotAvailable GenericErrorCodes = "LeaderboardVersionNotAvailable"
     GenericErrorCodesStatisticAlreadyHasPrizeTable GenericErrorCodes = "StatisticAlreadyHasPrizeTable"
     GenericErrorCodesPrizeTableHasOverlappingRanks GenericErrorCodes = "PrizeTableHasOverlappingRanks"
     GenericErrorCodesPrizeTableHasMissingRanks GenericErrorCodes = "PrizeTableHasMissingRanks"
     GenericErrorCodesPrizeTableRankStartsAtZero GenericErrorCodes = "PrizeTableRankStartsAtZero"
     GenericErrorCodesInvalidStatistic GenericErrorCodes = "InvalidStatistic"
     GenericErrorCodesExpressionParseFailure GenericErrorCodes = "ExpressionParseFailure"
     GenericErrorCodesExpressionInvokeFailure GenericErrorCodes = "ExpressionInvokeFailure"
     GenericErrorCodesExpressionTooLong GenericErrorCodes = "ExpressionTooLong"
     GenericErrorCodesDataUpdateRateExceeded GenericErrorCodes = "DataUpdateRateExceeded"
     GenericErrorCodesRestrictedEmailDomain GenericErrorCodes = "RestrictedEmailDomain"
     GenericErrorCodesEncryptionKeyDisabled GenericErrorCodes = "EncryptionKeyDisabled"
     GenericErrorCodesEncryptionKeyMissing GenericErrorCodes = "EncryptionKeyMissing"
     GenericErrorCodesEncryptionKeyBroken GenericErrorCodes = "EncryptionKeyBroken"
     GenericErrorCodesNoSharedSecretKeyConfigured GenericErrorCodes = "NoSharedSecretKeyConfigured"
     GenericErrorCodesSecretKeyNotFound GenericErrorCodes = "SecretKeyNotFound"
     GenericErrorCodesPlayerSecretAlreadyConfigured GenericErrorCodes = "PlayerSecretAlreadyConfigured"
     GenericErrorCodesAPIRequestsDisabledForTitle GenericErrorCodes = "APIRequestsDisabledForTitle"
     GenericErrorCodesInvalidSharedSecretKey GenericErrorCodes = "InvalidSharedSecretKey"
     GenericErrorCodesPrizeTableHasNoRanks GenericErrorCodes = "PrizeTableHasNoRanks"
     GenericErrorCodesProfileDoesNotExist GenericErrorCodes = "ProfileDoesNotExist"
     GenericErrorCodesContentS3OriginBucketNotConfigured GenericErrorCodes = "ContentS3OriginBucketNotConfigured"
     GenericErrorCodesInvalidEnvironmentForReceipt GenericErrorCodes = "InvalidEnvironmentForReceipt"
     GenericErrorCodesEncryptedRequestNotAllowed GenericErrorCodes = "EncryptedRequestNotAllowed"
     GenericErrorCodesSignedRequestNotAllowed GenericErrorCodes = "SignedRequestNotAllowed"
     GenericErrorCodesRequestViewConstraintParamsNotAllowed GenericErrorCodes = "RequestViewConstraintParamsNotAllowed"
     GenericErrorCodesBadPartnerConfiguration GenericErrorCodes = "BadPartnerConfiguration"
     GenericErrorCodesXboxBPCertificateFailure GenericErrorCodes = "XboxBPCertificateFailure"
     GenericErrorCodesXboxXASSExchangeFailure GenericErrorCodes = "XboxXASSExchangeFailure"
     GenericErrorCodesInvalidEntityId GenericErrorCodes = "InvalidEntityId"
     GenericErrorCodesStatisticValueAggregationOverflow GenericErrorCodes = "StatisticValueAggregationOverflow"
     GenericErrorCodesEmailMessageFromAddressIsMissing GenericErrorCodes = "EmailMessageFromAddressIsMissing"
     GenericErrorCodesEmailMessageToAddressIsMissing GenericErrorCodes = "EmailMessageToAddressIsMissing"
     GenericErrorCodesSmtpServerAuthenticationError GenericErrorCodes = "SmtpServerAuthenticationError"
     GenericErrorCodesSmtpServerLimitExceeded GenericErrorCodes = "SmtpServerLimitExceeded"
     GenericErrorCodesSmtpServerInsufficientStorage GenericErrorCodes = "SmtpServerInsufficientStorage"
     GenericErrorCodesSmtpServerCommunicationError GenericErrorCodes = "SmtpServerCommunicationError"
     GenericErrorCodesSmtpServerGeneralFailure GenericErrorCodes = "SmtpServerGeneralFailure"
     GenericErrorCodesEmailClientTimeout GenericErrorCodes = "EmailClientTimeout"
     GenericErrorCodesEmailClientCanceledTask GenericErrorCodes = "EmailClientCanceledTask"
     GenericErrorCodesEmailTemplateMissing GenericErrorCodes = "EmailTemplateMissing"
     GenericErrorCodesInvalidHostForTitleId GenericErrorCodes = "InvalidHostForTitleId"
     GenericErrorCodesEmailConfirmationTokenDoesNotExist GenericErrorCodes = "EmailConfirmationTokenDoesNotExist"
     GenericErrorCodesEmailConfirmationTokenExpired GenericErrorCodes = "EmailConfirmationTokenExpired"
     GenericErrorCodesAccountDeleted GenericErrorCodes = "AccountDeleted"
     GenericErrorCodesPlayerSecretNotConfigured GenericErrorCodes = "PlayerSecretNotConfigured"
     GenericErrorCodesInvalidSignatureTime GenericErrorCodes = "InvalidSignatureTime"
     GenericErrorCodesNoContactEmailAddressFound GenericErrorCodes = "NoContactEmailAddressFound"
     GenericErrorCodesInvalidAuthToken GenericErrorCodes = "InvalidAuthToken"
     GenericErrorCodesAuthTokenDoesNotExist GenericErrorCodes = "AuthTokenDoesNotExist"
     GenericErrorCodesAuthTokenExpired GenericErrorCodes = "AuthTokenExpired"
     GenericErrorCodesAuthTokenAlreadyUsedToResetPassword GenericErrorCodes = "AuthTokenAlreadyUsedToResetPassword"
     GenericErrorCodesMembershipNameTooLong GenericErrorCodes = "MembershipNameTooLong"
     GenericErrorCodesMembershipNotFound GenericErrorCodes = "MembershipNotFound"
     GenericErrorCodesGoogleServiceAccountInvalid GenericErrorCodes = "GoogleServiceAccountInvalid"
     GenericErrorCodesGoogleServiceAccountParseFailure GenericErrorCodes = "GoogleServiceAccountParseFailure"
     GenericErrorCodesEntityTokenMissing GenericErrorCodes = "EntityTokenMissing"
     GenericErrorCodesEntityTokenInvalid GenericErrorCodes = "EntityTokenInvalid"
     GenericErrorCodesEntityTokenExpired GenericErrorCodes = "EntityTokenExpired"
     GenericErrorCodesEntityTokenRevoked GenericErrorCodes = "EntityTokenRevoked"
     GenericErrorCodesInvalidProductForSubscription GenericErrorCodes = "InvalidProductForSubscription"
     GenericErrorCodesXboxInaccessible GenericErrorCodes = "XboxInaccessible"
     GenericErrorCodesSubscriptionAlreadyTaken GenericErrorCodes = "SubscriptionAlreadyTaken"
     GenericErrorCodesSmtpAddonNotEnabled GenericErrorCodes = "SmtpAddonNotEnabled"
     GenericErrorCodesAPIConcurrentRequestLimitExceeded GenericErrorCodes = "APIConcurrentRequestLimitExceeded"
     GenericErrorCodesXboxRejectedXSTSExchangeRequest GenericErrorCodes = "XboxRejectedXSTSExchangeRequest"
     GenericErrorCodesVariableNotDefined GenericErrorCodes = "VariableNotDefined"
     GenericErrorCodesTemplateVersionNotDefined GenericErrorCodes = "TemplateVersionNotDefined"
     GenericErrorCodesFileTooLarge GenericErrorCodes = "FileTooLarge"
     GenericErrorCodesTitleDeleted GenericErrorCodes = "TitleDeleted"
     GenericErrorCodesTitleContainsUserAccounts GenericErrorCodes = "TitleContainsUserAccounts"
     GenericErrorCodesTitleDeletionPlayerCleanupFailure GenericErrorCodes = "TitleDeletionPlayerCleanupFailure"
     GenericErrorCodesEntityFileOperationPending GenericErrorCodes = "EntityFileOperationPending"
     GenericErrorCodesNoEntityFileOperationPending GenericErrorCodes = "NoEntityFileOperationPending"
     GenericErrorCodesEntityProfileVersionMismatch GenericErrorCodes = "EntityProfileVersionMismatch"
     GenericErrorCodesTemplateVersionTooOld GenericErrorCodes = "TemplateVersionTooOld"
     GenericErrorCodesMembershipDefinitionInUse GenericErrorCodes = "MembershipDefinitionInUse"
     GenericErrorCodesPaymentPageNotConfigured GenericErrorCodes = "PaymentPageNotConfigured"
     GenericErrorCodesFailedLoginAttemptRateLimitExceeded GenericErrorCodes = "FailedLoginAttemptRateLimitExceeded"
     GenericErrorCodesEntityBlockedByGroup GenericErrorCodes = "EntityBlockedByGroup"
     GenericErrorCodesRoleDoesNotExist GenericErrorCodes = "RoleDoesNotExist"
     GenericErrorCodesEntityIsAlreadyMember GenericErrorCodes = "EntityIsAlreadyMember"
     GenericErrorCodesDuplicateRoleId GenericErrorCodes = "DuplicateRoleId"
     GenericErrorCodesGroupInvitationNotFound GenericErrorCodes = "GroupInvitationNotFound"
     GenericErrorCodesGroupApplicationNotFound GenericErrorCodes = "GroupApplicationNotFound"
     GenericErrorCodesOutstandingInvitationAcceptedInstead GenericErrorCodes = "OutstandingInvitationAcceptedInstead"
     GenericErrorCodesOutstandingApplicationAcceptedInstead GenericErrorCodes = "OutstandingApplicationAcceptedInstead"
     GenericErrorCodesRoleIsGroupDefaultMember GenericErrorCodes = "RoleIsGroupDefaultMember"
     GenericErrorCodesRoleIsGroupAdmin GenericErrorCodes = "RoleIsGroupAdmin"
     GenericErrorCodesRoleNameNotAvailable GenericErrorCodes = "RoleNameNotAvailable"
     GenericErrorCodesGroupNameNotAvailable GenericErrorCodes = "GroupNameNotAvailable"
     GenericErrorCodesEmailReportAlreadySent GenericErrorCodes = "EmailReportAlreadySent"
     GenericErrorCodesEmailReportRecipientBlacklisted GenericErrorCodes = "EmailReportRecipientBlacklisted"
     GenericErrorCodesEventNamespaceNotAllowed GenericErrorCodes = "EventNamespaceNotAllowed"
     GenericErrorCodesEventEntityNotAllowed GenericErrorCodes = "EventEntityNotAllowed"
     GenericErrorCodesInvalidEntityType GenericErrorCodes = "InvalidEntityType"
     GenericErrorCodesNullTokenResultFromAad GenericErrorCodes = "NullTokenResultFromAad"
     GenericErrorCodesInvalidTokenResultFromAad GenericErrorCodes = "InvalidTokenResultFromAad"
     GenericErrorCodesNoValidCertificateForAad GenericErrorCodes = "NoValidCertificateForAad"
     GenericErrorCodesInvalidCertificateForAad GenericErrorCodes = "InvalidCertificateForAad"
     GenericErrorCodesDuplicateDropTableId GenericErrorCodes = "DuplicateDropTableId"
     GenericErrorCodesMultiplayerServerError GenericErrorCodes = "MultiplayerServerError"
     GenericErrorCodesMultiplayerServerTooManyRequests GenericErrorCodes = "MultiplayerServerTooManyRequests"
     GenericErrorCodesMultiplayerServerNoContent GenericErrorCodes = "MultiplayerServerNoContent"
     GenericErrorCodesMultiplayerServerBadRequest GenericErrorCodes = "MultiplayerServerBadRequest"
     GenericErrorCodesMultiplayerServerUnauthorized GenericErrorCodes = "MultiplayerServerUnauthorized"
     GenericErrorCodesMultiplayerServerForbidden GenericErrorCodes = "MultiplayerServerForbidden"
     GenericErrorCodesMultiplayerServerNotFound GenericErrorCodes = "MultiplayerServerNotFound"
     GenericErrorCodesMultiplayerServerConflict GenericErrorCodes = "MultiplayerServerConflict"
     GenericErrorCodesMultiplayerServerInternalServerError GenericErrorCodes = "MultiplayerServerInternalServerError"
     GenericErrorCodesMultiplayerServerUnavailable GenericErrorCodes = "MultiplayerServerUnavailable"
     GenericErrorCodesExplicitContentDetected GenericErrorCodes = "ExplicitContentDetected"
     GenericErrorCodesPIIContentDetected GenericErrorCodes = "PIIContentDetected"
     GenericErrorCodesInvalidScheduledTaskParameter GenericErrorCodes = "InvalidScheduledTaskParameter"
     GenericErrorCodesPerEntityEventRateLimitExceeded GenericErrorCodes = "PerEntityEventRateLimitExceeded"
     GenericErrorCodesTitleDefaultLanguageNotSet GenericErrorCodes = "TitleDefaultLanguageNotSet"
     GenericErrorCodesEmailTemplateMissingDefaultVersion GenericErrorCodes = "EmailTemplateMissingDefaultVersion"
     GenericErrorCodesFacebookInstantGamesIdNotLinked GenericErrorCodes = "FacebookInstantGamesIdNotLinked"
     GenericErrorCodesInvalidFacebookInstantGamesSignature GenericErrorCodes = "InvalidFacebookInstantGamesSignature"
     GenericErrorCodesFacebookInstantGamesAuthNotConfiguredForTitle GenericErrorCodes = "FacebookInstantGamesAuthNotConfiguredForTitle"
     GenericErrorCodesEntityProfileConstraintValidationFailed GenericErrorCodes = "EntityProfileConstraintValidationFailed"
     GenericErrorCodesTelemetryIngestionKeyPending GenericErrorCodes = "TelemetryIngestionKeyPending"
     GenericErrorCodesTelemetryIngestionKeyNotFound GenericErrorCodes = "TelemetryIngestionKeyNotFound"
     GenericErrorCodesStatisticChildNameInvalid GenericErrorCodes = "StatisticChildNameInvalid"
     GenericErrorCodesDataIntegrityError GenericErrorCodes = "DataIntegrityError"
     GenericErrorCodesVirtualCurrencyCannotBeSetToOlderVersion GenericErrorCodes = "VirtualCurrencyCannotBeSetToOlderVersion"
     GenericErrorCodesVirtualCurrencyMustBeWithinIntegerRange GenericErrorCodes = "VirtualCurrencyMustBeWithinIntegerRange"
     GenericErrorCodesEmailTemplateInvalidSyntax GenericErrorCodes = "EmailTemplateInvalidSyntax"
     GenericErrorCodesEmailTemplateMissingCallback GenericErrorCodes = "EmailTemplateMissingCallback"
     GenericErrorCodesPushNotificationTemplateInvalidPayload GenericErrorCodes = "PushNotificationTemplateInvalidPayload"
     GenericErrorCodesInvalidLocalizedPushNotificationLanguage GenericErrorCodes = "InvalidLocalizedPushNotificationLanguage"
     GenericErrorCodesMissingLocalizedPushNotificationMessage GenericErrorCodes = "MissingLocalizedPushNotificationMessage"
     GenericErrorCodesPushNotificationTemplateMissingPlatformPayload GenericErrorCodes = "PushNotificationTemplateMissingPlatformPayload"
     GenericErrorCodesPushNotificationTemplatePayloadContainsInvalidJson GenericErrorCodes = "PushNotificationTemplatePayloadContainsInvalidJson"
     GenericErrorCodesPushNotificationTemplateContainsInvalidIosPayload GenericErrorCodes = "PushNotificationTemplateContainsInvalidIosPayload"
     GenericErrorCodesPushNotificationTemplateContainsInvalidAndroidPayload GenericErrorCodes = "PushNotificationTemplateContainsInvalidAndroidPayload"
     GenericErrorCodesPushNotificationTemplateIosPayloadMissingNotificationBody GenericErrorCodes = "PushNotificationTemplateIosPayloadMissingNotificationBody"
     GenericErrorCodesPushNotificationTemplateAndroidPayloadMissingNotificationBody GenericErrorCodes = "PushNotificationTemplateAndroidPayloadMissingNotificationBody"
     GenericErrorCodesPushNotificationTemplateNotFound GenericErrorCodes = "PushNotificationTemplateNotFound"
     GenericErrorCodesPushNotificationTemplateMissingDefaultVersion GenericErrorCodes = "PushNotificationTemplateMissingDefaultVersion"
     GenericErrorCodesPushNotificationTemplateInvalidSyntax GenericErrorCodes = "PushNotificationTemplateInvalidSyntax"
     GenericErrorCodesPushNotificationTemplateNoCustomPayloadForV1 GenericErrorCodes = "PushNotificationTemplateNoCustomPayloadForV1"
     GenericErrorCodesNoLeaderboardForStatistic GenericErrorCodes = "NoLeaderboardForStatistic"
     GenericErrorCodesTitleNewsMissingDefaultLanguage GenericErrorCodes = "TitleNewsMissingDefaultLanguage"
     GenericErrorCodesTitleNewsNotFound GenericErrorCodes = "TitleNewsNotFound"
     GenericErrorCodesTitleNewsDuplicateLanguage GenericErrorCodes = "TitleNewsDuplicateLanguage"
     GenericErrorCodesTitleNewsMissingTitleOrBody GenericErrorCodes = "TitleNewsMissingTitleOrBody"
     GenericErrorCodesTitleNewsInvalidLanguage GenericErrorCodes = "TitleNewsInvalidLanguage"
     GenericErrorCodesEmailRecipientBlacklisted GenericErrorCodes = "EmailRecipientBlacklisted"
     GenericErrorCodesInvalidGameCenterAuthRequest GenericErrorCodes = "InvalidGameCenterAuthRequest"
     GenericErrorCodesGameCenterAuthenticationFailed GenericErrorCodes = "GameCenterAuthenticationFailed"
     GenericErrorCodesCannotEnablePartiesForTitle GenericErrorCodes = "CannotEnablePartiesForTitle"
     GenericErrorCodesPartyError GenericErrorCodes = "PartyError"
     GenericErrorCodesPartyRequests GenericErrorCodes = "PartyRequests"
     GenericErrorCodesPartyNoContent GenericErrorCodes = "PartyNoContent"
     GenericErrorCodesPartyBadRequest GenericErrorCodes = "PartyBadRequest"
     GenericErrorCodesPartyUnauthorized GenericErrorCodes = "PartyUnauthorized"
     GenericErrorCodesPartyForbidden GenericErrorCodes = "PartyForbidden"
     GenericErrorCodesPartyNotFound GenericErrorCodes = "PartyNotFound"
     GenericErrorCodesPartyConflict GenericErrorCodes = "PartyConflict"
     GenericErrorCodesPartyInternalServerError GenericErrorCodes = "PartyInternalServerError"
     GenericErrorCodesPartyUnavailable GenericErrorCodes = "PartyUnavailable"
     GenericErrorCodesPartyTooManyRequests GenericErrorCodes = "PartyTooManyRequests"
     GenericErrorCodesPushNotificationTemplateMissingName GenericErrorCodes = "PushNotificationTemplateMissingName"
     GenericErrorCodesCannotEnableMultiplayerServersForTitle GenericErrorCodes = "CannotEnableMultiplayerServersForTitle"
     GenericErrorCodesWriteAttemptedDuringExport GenericErrorCodes = "WriteAttemptedDuringExport"
     GenericErrorCodesMultiplayerServerTitleQuotaCoresExceeded GenericErrorCodes = "MultiplayerServerTitleQuotaCoresExceeded"
     GenericErrorCodesAutomationRuleNotFound GenericErrorCodes = "AutomationRuleNotFound"
     GenericErrorCodesEntityAPIKeyLimitExceeded GenericErrorCodes = "EntityAPIKeyLimitExceeded"
     GenericErrorCodesEntityAPIKeyNotFound GenericErrorCodes = "EntityAPIKeyNotFound"
     GenericErrorCodesEntityAPIKeyOrSecretInvalid GenericErrorCodes = "EntityAPIKeyOrSecretInvalid"
     GenericErrorCodesEconomyServiceUnavailable GenericErrorCodes = "EconomyServiceUnavailable"
     GenericErrorCodesEconomyServiceInternalError GenericErrorCodes = "EconomyServiceInternalError"
     GenericErrorCodesQueryRateLimitExceeded GenericErrorCodes = "QueryRateLimitExceeded"
     GenericErrorCodesEntityAPIKeyCreationDisabledForEntity GenericErrorCodes = "EntityAPIKeyCreationDisabledForEntity"
     GenericErrorCodesForbiddenByEntityPolicy GenericErrorCodes = "ForbiddenByEntityPolicy"
     GenericErrorCodesUpdateInventoryRateLimitExceeded GenericErrorCodes = "UpdateInventoryRateLimitExceeded"
     GenericErrorCodesStudioCreationRateLimited GenericErrorCodes = "StudioCreationRateLimited"
     GenericErrorCodesStudioCreationInProgress GenericErrorCodes = "StudioCreationInProgress"
     GenericErrorCodesDuplicateStudioName GenericErrorCodes = "DuplicateStudioName"
     GenericErrorCodesStudioNotFound GenericErrorCodes = "StudioNotFound"
     GenericErrorCodesStudioDeleted GenericErrorCodes = "StudioDeleted"
     GenericErrorCodesStudioDeactivated GenericErrorCodes = "StudioDeactivated"
     GenericErrorCodesStudioActivated GenericErrorCodes = "StudioActivated"
     GenericErrorCodesTitleCreationRateLimited GenericErrorCodes = "TitleCreationRateLimited"
     GenericErrorCodesTitleCreationInProgress GenericErrorCodes = "TitleCreationInProgress"
     GenericErrorCodesDuplicateTitleName GenericErrorCodes = "DuplicateTitleName"
     GenericErrorCodesTitleActivationRateLimited GenericErrorCodes = "TitleActivationRateLimited"
     GenericErrorCodesTitleActivationInProgress GenericErrorCodes = "TitleActivationInProgress"
     GenericErrorCodesTitleDeactivated GenericErrorCodes = "TitleDeactivated"
     GenericErrorCodesTitleActivated GenericErrorCodes = "TitleActivated"
     GenericErrorCodesCloudScriptAzureFunctionsExecutionTimeLimitExceeded GenericErrorCodes = "CloudScriptAzureFunctionsExecutionTimeLimitExceeded"
     GenericErrorCodesCloudScriptAzureFunctionsArgumentSizeExceeded GenericErrorCodes = "CloudScriptAzureFunctionsArgumentSizeExceeded"
     GenericErrorCodesCloudScriptAzureFunctionsReturnSizeExceeded GenericErrorCodes = "CloudScriptAzureFunctionsReturnSizeExceeded"
     GenericErrorCodesCloudScriptAzureFunctionsHTTPRequestError GenericErrorCodes = "CloudScriptAzureFunctionsHTTPRequestError"
     GenericErrorCodesVirtualCurrencyBetaGetError GenericErrorCodes = "VirtualCurrencyBetaGetError"
     GenericErrorCodesVirtualCurrencyBetaCreateError GenericErrorCodes = "VirtualCurrencyBetaCreateError"
     GenericErrorCodesVirtualCurrencyBetaInitialDepositSaveError GenericErrorCodes = "VirtualCurrencyBetaInitialDepositSaveError"
     GenericErrorCodesVirtualCurrencyBetaSaveError GenericErrorCodes = "VirtualCurrencyBetaSaveError"
     GenericErrorCodesVirtualCurrencyBetaDeleteError GenericErrorCodes = "VirtualCurrencyBetaDeleteError"
     GenericErrorCodesVirtualCurrencyBetaRestoreError GenericErrorCodes = "VirtualCurrencyBetaRestoreError"
     GenericErrorCodesVirtualCurrencyBetaSaveConflict GenericErrorCodes = "VirtualCurrencyBetaSaveConflict"
     GenericErrorCodesVirtualCurrencyBetaUpdateError GenericErrorCodes = "VirtualCurrencyBetaUpdateError"
     GenericErrorCodesInsightsManagementDatabaseNotFound GenericErrorCodes = "InsightsManagementDatabaseNotFound"
     GenericErrorCodesInsightsManagementOperationNotFound GenericErrorCodes = "InsightsManagementOperationNotFound"
     GenericErrorCodesInsightsManagementErrorPendingOperationExists GenericErrorCodes = "InsightsManagementErrorPendingOperationExists"
     GenericErrorCodesInsightsManagementSetPerformanceLevelInvalidParameter GenericErrorCodes = "InsightsManagementSetPerformanceLevelInvalidParameter"
     GenericErrorCodesInsightsManagementSetStorageRetentionInvalidParameter GenericErrorCodes = "InsightsManagementSetStorageRetentionInvalidParameter"
     GenericErrorCodesInsightsManagementGetStorageUsageInvalidParameter GenericErrorCodes = "InsightsManagementGetStorageUsageInvalidParameter"
     GenericErrorCodesInsightsManagementGetOperationStatusInvalidParameter GenericErrorCodes = "InsightsManagementGetOperationStatusInvalidParameter"
     GenericErrorCodesDuplicatePurchaseTransactionId GenericErrorCodes = "DuplicatePurchaseTransactionId"
     GenericErrorCodesEvaluationModePlayerCountExceeded GenericErrorCodes = "EvaluationModePlayerCountExceeded"
     GenericErrorCodesGetPlayersInSegmentRateLimitExceeded GenericErrorCodes = "GetPlayersInSegmentRateLimitExceeded"
     GenericErrorCodesCloudScriptFunctionNameSizeExceeded GenericErrorCodes = "CloudScriptFunctionNameSizeExceeded"
     GenericErrorCodesInsightsManagementTitleInEvaluationMode GenericErrorCodes = "InsightsManagementTitleInEvaluationMode"
     GenericErrorCodesCloudScriptAzureFunctionsQueueRequestError GenericErrorCodes = "CloudScriptAzureFunctionsQueueRequestError"
     GenericErrorCodesEvaluationModeTitleCountExceeded GenericErrorCodes = "EvaluationModeTitleCountExceeded"
     GenericErrorCodesInsightsManagementTitleNotInFlight GenericErrorCodes = "InsightsManagementTitleNotInFlight"
     GenericErrorCodesLimitNotFound GenericErrorCodes = "LimitNotFound"
     GenericErrorCodesLimitNotAvailableViaAPI GenericErrorCodes = "LimitNotAvailableViaAPI"
     GenericErrorCodesInsightsManagementSetStorageRetentionBelowMinimum GenericErrorCodes = "InsightsManagementSetStorageRetentionBelowMinimum"
     GenericErrorCodesInsightsManagementSetStorageRetentionAboveMaximum GenericErrorCodes = "InsightsManagementSetStorageRetentionAboveMaximum"
     GenericErrorCodesAppleNotEnabledForTitle GenericErrorCodes = "AppleNotEnabledForTitle"
     GenericErrorCodesInsightsManagementNewActiveEventExportLimitInvalid GenericErrorCodes = "InsightsManagementNewActiveEventExportLimitInvalid"
     GenericErrorCodesInsightsManagementSetPerformanceRateLimited GenericErrorCodes = "InsightsManagementSetPerformanceRateLimited"
     GenericErrorCodesPartyRequestsThrottledFromRateLimiter GenericErrorCodes = "PartyRequestsThrottledFromRateLimiter"
     GenericErrorCodesXboxServiceTooManyRequests GenericErrorCodes = "XboxServiceTooManyRequests"
     GenericErrorCodesNintendoSwitchNotEnabledForTitle GenericErrorCodes = "NintendoSwitchNotEnabledForTitle"
     GenericErrorCodesRequestMultiplayerServersThrottledFromRateLimiter GenericErrorCodes = "RequestMultiplayerServersThrottledFromRateLimiter"
     GenericErrorCodesTitleDataOverrideNotFound GenericErrorCodes = "TitleDataOverrideNotFound"
     GenericErrorCodesDuplicateKeys GenericErrorCodes = "DuplicateKeys"
     GenericErrorCodesWasNotCreatedWithCloudRoot GenericErrorCodes = "WasNotCreatedWithCloudRoot"
     GenericErrorCodesLegacyMultiplayerServersDeprecated GenericErrorCodes = "LegacyMultiplayerServersDeprecated"
     GenericErrorCodesMatchmakingEntityInvalid GenericErrorCodes = "MatchmakingEntityInvalid"
     GenericErrorCodesMatchmakingPlayerAttributesInvalid GenericErrorCodes = "MatchmakingPlayerAttributesInvalid"
     GenericErrorCodesMatchmakingQueueNotFound GenericErrorCodes = "MatchmakingQueueNotFound"
     GenericErrorCodesMatchmakingMatchNotFound GenericErrorCodes = "MatchmakingMatchNotFound"
     GenericErrorCodesMatchmakingTicketNotFound GenericErrorCodes = "MatchmakingTicketNotFound"
     GenericErrorCodesMatchmakingAlreadyJoinedTicket GenericErrorCodes = "MatchmakingAlreadyJoinedTicket"
     GenericErrorCodesMatchmakingTicketAlreadyCompleted GenericErrorCodes = "MatchmakingTicketAlreadyCompleted"
     GenericErrorCodesMatchmakingQueueConfigInvalid GenericErrorCodes = "MatchmakingQueueConfigInvalid"
     GenericErrorCodesMatchmakingMemberProfileInvalid GenericErrorCodes = "MatchmakingMemberProfileInvalid"
     GenericErrorCodesNintendoSwitchDeviceIdNotLinked GenericErrorCodes = "NintendoSwitchDeviceIdNotLinked"
     GenericErrorCodesMatchmakingNotEnabled GenericErrorCodes = "MatchmakingNotEnabled"
     GenericErrorCodesMatchmakingPlayerAttributesTooLarge GenericErrorCodes = "MatchmakingPlayerAttributesTooLarge"
     GenericErrorCodesMatchmakingNumberOfPlayersInTicketTooLarge GenericErrorCodes = "MatchmakingNumberOfPlayersInTicketTooLarge"
     GenericErrorCodesMatchmakingAttributeInvalid GenericErrorCodes = "MatchmakingAttributeInvalid"
     GenericErrorCodesMatchmakingPlayerHasNotJoinedTicket GenericErrorCodes = "MatchmakingPlayerHasNotJoinedTicket"
     GenericErrorCodesMatchmakingRateLimitExceeded GenericErrorCodes = "MatchmakingRateLimitExceeded"
     GenericErrorCodesMatchmakingTicketMembershipLimitExceeded GenericErrorCodes = "MatchmakingTicketMembershipLimitExceeded"
     GenericErrorCodesMatchmakingUnauthorized GenericErrorCodes = "MatchmakingUnauthorized"
     GenericErrorCodesMatchmakingQueueLimitExceeded GenericErrorCodes = "MatchmakingQueueLimitExceeded"
     GenericErrorCodesMatchmakingRequestTypeMismatch GenericErrorCodes = "MatchmakingRequestTypeMismatch"
     GenericErrorCodesMatchmakingBadRequest GenericErrorCodes = "MatchmakingBadRequest"
     GenericErrorCodesTitleConfigNotFound GenericErrorCodes = "TitleConfigNotFound"
     GenericErrorCodesTitleConfigUpdateConflict GenericErrorCodes = "TitleConfigUpdateConflict"
     GenericErrorCodesTitleConfigSerializationError GenericErrorCodes = "TitleConfigSerializationError"
     GenericErrorCodesCatalogEntityInvalid GenericErrorCodes = "CatalogEntityInvalid"
     GenericErrorCodesCatalogTitleIdMissing GenericErrorCodes = "CatalogTitleIdMissing"
     GenericErrorCodesCatalogPlayerIdMissing GenericErrorCodes = "CatalogPlayerIdMissing"
     GenericErrorCodesCatalogClientIdentityInvalid GenericErrorCodes = "CatalogClientIdentityInvalid"
     GenericErrorCodesCatalogOneOrMoreFilesInvalid GenericErrorCodes = "CatalogOneOrMoreFilesInvalid"
     GenericErrorCodesCatalogItemMetadataInvalid GenericErrorCodes = "CatalogItemMetadataInvalid"
     GenericErrorCodesCatalogItemIdInvalid GenericErrorCodes = "CatalogItemIdInvalid"
     GenericErrorCodesCatalogSearchParameterInvalid GenericErrorCodes = "CatalogSearchParameterInvalid"
     GenericErrorCodesCatalogFeatureDisabled GenericErrorCodes = "CatalogFeatureDisabled"
     GenericErrorCodesCatalogConfigInvalid GenericErrorCodes = "CatalogConfigInvalid"
     GenericErrorCodesCatalogItemTypeInvalid GenericErrorCodes = "CatalogItemTypeInvalid"
     GenericErrorCodesCatalogBadRequest GenericErrorCodes = "CatalogBadRequest"
     GenericErrorCodesCatalogTooManyRequests GenericErrorCodes = "CatalogTooManyRequests"
     GenericErrorCodesExportInvalidStatusUpdate GenericErrorCodes = "ExportInvalidStatusUpdate"
     GenericErrorCodesExportInvalidPrefix GenericErrorCodes = "ExportInvalidPrefix"
     GenericErrorCodesExportBlobContainerDoesNotExist GenericErrorCodes = "ExportBlobContainerDoesNotExist"
     GenericErrorCodesExportNotFound GenericErrorCodes = "ExportNotFound"
     GenericErrorCodesExportCouldNotUpdate GenericErrorCodes = "ExportCouldNotUpdate"
     GenericErrorCodesExportInvalidStorageType GenericErrorCodes = "ExportInvalidStorageType"
     GenericErrorCodesExportAmazonBucketDoesNotExist GenericErrorCodes = "ExportAmazonBucketDoesNotExist"
     GenericErrorCodesExportInvalidBlobStorage GenericErrorCodes = "ExportInvalidBlobStorage"
     GenericErrorCodesExportKustoException GenericErrorCodes = "ExportKustoException"
     GenericErrorCodesExportKustoConnectionFailed GenericErrorCodes = "ExportKustoConnectionFailed"
     GenericErrorCodesExportUnknownError GenericErrorCodes = "ExportUnknownError"
     GenericErrorCodesExportCantEditPendingExport GenericErrorCodes = "ExportCantEditPendingExport"
     GenericErrorCodesExportLimitExports GenericErrorCodes = "ExportLimitExports"
     GenericErrorCodesExportLimitEvents GenericErrorCodes = "ExportLimitEvents"
     GenericErrorCodesExportInvalidPartitionStatusModification GenericErrorCodes = "ExportInvalidPartitionStatusModification"
     GenericErrorCodesExportCouldNotCreate GenericErrorCodes = "ExportCouldNotCreate"
     GenericErrorCodesExportNoBackingDatabaseFound GenericErrorCodes = "ExportNoBackingDatabaseFound"
     GenericErrorCodesExportCouldNotDelete GenericErrorCodes = "ExportCouldNotDelete"
     GenericErrorCodesExportCannotDetermineEventQuery GenericErrorCodes = "ExportCannotDetermineEventQuery"
     GenericErrorCodesExportInvalidQuerySchemaModification GenericErrorCodes = "ExportInvalidQuerySchemaModification"
     GenericErrorCodesExportQuerySchemaMissingRequiredColumns GenericErrorCodes = "ExportQuerySchemaMissingRequiredColumns"
     GenericErrorCodesExportCannotParseQuery GenericErrorCodes = "ExportCannotParseQuery"
     GenericErrorCodesExportControlCommandsNotAllowed GenericErrorCodes = "ExportControlCommandsNotAllowed"
     GenericErrorCodesExportQueryMissingTableReference GenericErrorCodes = "ExportQueryMissingTableReference"
     GenericErrorCodesTitleNotEnabledForParty GenericErrorCodes = "TitleNotEnabledForParty"
     GenericErrorCodesPartyVersionNotFound GenericErrorCodes = "PartyVersionNotFound"
     GenericErrorCodesMultiplayerServerBuildReferencedByMatchmakingQueue GenericErrorCodes = "MultiplayerServerBuildReferencedByMatchmakingQueue"
     GenericErrorCodesExperimentationExperimentStopped GenericErrorCodes = "ExperimentationExperimentStopped"
     GenericErrorCodesExperimentationExperimentRunning GenericErrorCodes = "ExperimentationExperimentRunning"
     GenericErrorCodesExperimentationExperimentNotFound GenericErrorCodes = "ExperimentationExperimentNotFound"
     GenericErrorCodesExperimentationExperimentNeverStarted GenericErrorCodes = "ExperimentationExperimentNeverStarted"
     GenericErrorCodesExperimentationExperimentDeleted GenericErrorCodes = "ExperimentationExperimentDeleted"
     GenericErrorCodesExperimentationClientTimeout GenericErrorCodes = "ExperimentationClientTimeout"
     GenericErrorCodesExperimentationInvalidVariantConfiguration GenericErrorCodes = "ExperimentationInvalidVariantConfiguration"
     GenericErrorCodesExperimentationInvalidVariableConfiguration GenericErrorCodes = "ExperimentationInvalidVariableConfiguration"
     GenericErrorCodesExperimentInvalidId GenericErrorCodes = "ExperimentInvalidId"
     GenericErrorCodesExperimentationNoScorecard GenericErrorCodes = "ExperimentationNoScorecard"
     GenericErrorCodesExperimentationTreatmentAssignmentFailed GenericErrorCodes = "ExperimentationTreatmentAssignmentFailed"
     GenericErrorCodesExperimentationTreatmentAssignmentDisabled GenericErrorCodes = "ExperimentationTreatmentAssignmentDisabled"
     GenericErrorCodesExperimentationInvalidDuration GenericErrorCodes = "ExperimentationInvalidDuration"
     GenericErrorCodesExperimentationMaxExperimentsReached GenericErrorCodes = "ExperimentationMaxExperimentsReached"
     GenericErrorCodesExperimentationExperimentSchedulingInProgress GenericErrorCodes = "ExperimentationExperimentSchedulingInProgress"
     GenericErrorCodesExperimentationInvalidEndDate GenericErrorCodes = "ExperimentationInvalidEndDate"
     GenericErrorCodesExperimentationInvalidStartDate GenericErrorCodes = "ExperimentationInvalidStartDate"
     GenericErrorCodesExperimentationMaxDurationExceeded GenericErrorCodes = "ExperimentationMaxDurationExceeded"
     GenericErrorCodesExperimentationExclusionGroupNotFound GenericErrorCodes = "ExperimentationExclusionGroupNotFound"
     GenericErrorCodesExperimentationExclusionGroupInsufficientCapacity GenericErrorCodes = "ExperimentationExclusionGroupInsufficientCapacity"
     GenericErrorCodesExperimentationExclusionGroupCannotDelete GenericErrorCodes = "ExperimentationExclusionGroupCannotDelete"
     GenericErrorCodesExperimentationExclusionGroupInvalidTrafficAllocation GenericErrorCodes = "ExperimentationExclusionGroupInvalidTrafficAllocation"
     GenericErrorCodesMaxActionDepthExceeded GenericErrorCodes = "MaxActionDepthExceeded"
     GenericErrorCodesTitleNotOnUpdatedPricingPlan GenericErrorCodes = "TitleNotOnUpdatedPricingPlan"
     GenericErrorCodesSegmentManagementTitleNotInFlight GenericErrorCodes = "SegmentManagementTitleNotInFlight"
     GenericErrorCodesSegmentManagementNoExpressionTree GenericErrorCodes = "SegmentManagementNoExpressionTree"
     GenericErrorCodesSegmentManagementTriggerActionCountOverLimit GenericErrorCodes = "SegmentManagementTriggerActionCountOverLimit"
     GenericErrorCodesSegmentManagementSegmentCountOverLimit GenericErrorCodes = "SegmentManagementSegmentCountOverLimit"
     GenericErrorCodesSegmentManagementInvalidSegmentId GenericErrorCodes = "SegmentManagementInvalidSegmentId"
     GenericErrorCodesSegmentManagementInvalidInput GenericErrorCodes = "SegmentManagementInvalidInput"
     GenericErrorCodesSegmentManagementInvalidSegmentName GenericErrorCodes = "SegmentManagementInvalidSegmentName"
     GenericErrorCodesDeleteSegmentRateLimitExceeded GenericErrorCodes = "DeleteSegmentRateLimitExceeded"
     GenericErrorCodesCreateSegmentRateLimitExceeded GenericErrorCodes = "CreateSegmentRateLimitExceeded"
     GenericErrorCodesUpdateSegmentRateLimitExceeded GenericErrorCodes = "UpdateSegmentRateLimitExceeded"
     GenericErrorCodesGetSegmentsRateLimitExceeded GenericErrorCodes = "GetSegmentsRateLimitExceeded"
     GenericErrorCodesSnapshotNotFound GenericErrorCodes = "SnapshotNotFound"
      )
// GetActionsOnPlayersInSegmentTaskInstanceResult 
type GetActionsOnPlayersInSegmentTaskInstanceResultModel struct {
    // Parameter parameter of this task instance
    Parameter *ActionsOnPlayersInSegmentTaskParameterModel `json:"Parameter,omitempty"`
    // Summary status summary of the actions-on-players-in-segment task instance
    Summary *ActionsOnPlayersInSegmentTaskSummaryModel `json:"Summary,omitempty"`
}

// GetAllSegmentsRequest request has no paramaters.
type GetAllSegmentsRequestModel struct {
}

// GetAllSegmentsResult 
type GetAllSegmentsResultModel struct {
    // Segments array of segments for this title.
    Segments []GetSegmentResultModel `json:"Segments,omitempty"`
}

// GetCatalogItemsRequest 
type GetCatalogItemsRequestModel struct {
    // CatalogVersion which catalog is being requested. If null, uses the default catalog.
    CatalogVersion string `json:"CatalogVersion,omitempty"`
}

// GetCatalogItemsResult 
type GetCatalogItemsResultModel struct {
    // Catalog array of items which can be purchased.
    Catalog []CatalogItemModel `json:"Catalog,omitempty"`
}

// GetCloudScriptRevisionRequest 
type GetCloudScriptRevisionRequestModel struct {
    // Revision revision number. If left null, defaults to the latest revision
    Revision int32 `json:"Revision,omitempty"`
    // Version version number. If left null, defaults to the latest version
    Version int32 `json:"Version,omitempty"`
}

// GetCloudScriptRevisionResult 
type GetCloudScriptRevisionResultModel struct {
    // CreatedAt time this revision was created
    CreatedAt time.Time `json:"CreatedAt,omitempty"`
    // Files list of Cloud Script files in this revision.
    Files []CloudScriptFileModel `json:"Files,omitempty"`
    // IsPublished true if this is the currently published revision
    IsPublished bool `json:"IsPublished"`
    // Revision revision number.
    Revision int32 `json:"Revision,omitempty"`
    // Version version number.
    Version int32 `json:"Version,omitempty"`
}

// GetCloudScriptTaskInstanceResult 
type GetCloudScriptTaskInstanceResultModel struct {
    // Parameter parameter of this task instance
    Parameter *CloudScriptTaskParameterModel `json:"Parameter,omitempty"`
    // Summary status summary of the CloudScript task instance
    Summary *CloudScriptTaskSummaryModel `json:"Summary,omitempty"`
}

// GetCloudScriptVersionsRequest 
type GetCloudScriptVersionsRequestModel struct {
}

// GetCloudScriptVersionsResult 
type GetCloudScriptVersionsResultModel struct {
    // Versions list of versions
    Versions []CloudScriptVersionStatusModel `json:"Versions,omitempty"`
}

// GetContentListRequest 
type GetContentListRequestModel struct {
    // Prefix limits the response to keys that begin with the specified prefix. You can use prefixes to list contents under a folder,
// or for a specified version, etc.
    Prefix string `json:"Prefix,omitempty"`
}

// GetContentListResult 
type GetContentListResultModel struct {
    // Contents list of content items.
    Contents []ContentInfoModel `json:"Contents,omitempty"`
    // ItemCount number of content items returned. We currently have a maximum of 1000 items limit.
    ItemCount int32 `json:"ItemCount,omitempty"`
    // TotalSize the total size of listed contents in bytes.
    TotalSize uint32 `json:"TotalSize,omitempty"`
}

// GetContentUploadUrlRequest 
type GetContentUploadUrlRequestModel struct {
    // ContentType a standard MIME type describing the format of the contents. The same MIME type has to be set in the header when
// uploading the content. If not specified, the MIME type is 'binary/octet-stream' by default.
    ContentType string `json:"ContentType,omitempty"`
    // Key key of the content item to upload, usually formatted as a path, e.g. images/a.png
    Key string `json:"Key,omitempty"`
}

// GetContentUploadUrlResult 
type GetContentUploadUrlResultModel struct {
    // URL uRL for uploading content via HTTP PUT method. The URL requires the 'x-ms-blob-type' header to have the value
// 'BlockBlob'. The URL will expire in approximately one hour.
    URL string `json:"URL,omitempty"`
}

// GetDataReportRequest gets the download URL for the requested report data (in CSV form). The reports available through this API call are those
// available in the Game Manager, in the Analytics->Reports tab.
type GetDataReportRequestModel struct {
    // Day reporting year (UTC)
    Day int32 `json:"Day,omitempty"`
    // Month reporting month (UTC)
    Month int32 `json:"Month,omitempty"`
    // ReportName report name
    ReportName string `json:"ReportName,omitempty"`
    // Year reporting year (UTC)
    Year int32 `json:"Year,omitempty"`
}

// GetDataReportResult 
type GetDataReportResultModel struct {
    // DownloadUrl the URL where the requested report can be downloaded. This can be any PlayFab generated reports. The full list of
// reports can be found at: https://docs.microsoft.com/en-us/gaming/playfab/features/analytics/reports/quickstart.
    DownloadUrl string `json:"DownloadUrl,omitempty"`
}

// GetMatchmakerGameInfoRequest 
type GetMatchmakerGameInfoRequestModel struct {
    // LobbyId unique identifier of the lobby for which info is being requested
    LobbyId string `json:"LobbyId,omitempty"`
}

// GetMatchmakerGameInfoResult 
type GetMatchmakerGameInfoResultModel struct {
    // BuildVersion version identifier of the game server executable binary being run
    BuildVersion string `json:"BuildVersion,omitempty"`
    // EndTime time when Game Server Instance is currently scheduled to end
    EndTime time.Time `json:"EndTime,omitempty"`
    // LobbyId unique identifier of the lobby
    LobbyId string `json:"LobbyId,omitempty"`
    // Mode game mode for this Game Server Instance
    Mode string `json:"Mode,omitempty"`
    // Players array of unique PlayFab identifiers for users currently connected to this Game Server Instance
    Players []string `json:"Players,omitempty"`
    // Region region in which the Game Server Instance is running
    Region Region `json:"Region,omitempty"`
    // ServerIPV4Address iPV4 address of the server
    ServerIPV4Address string `json:"ServerIPV4Address,omitempty"`
    // ServerIPV6Address iPV6 address of the server
    ServerIPV6Address string `json:"ServerIPV6Address,omitempty"`
    // ServerPort communication port for this Game Server Instance
    ServerPort uint32 `json:"ServerPort,omitempty"`
    // ServerPublicDNSName public DNS name (if any) of the server
    ServerPublicDNSName string `json:"ServerPublicDNSName,omitempty"`
    // StartTime time when the Game Server Instance was created
    StartTime time.Time `json:"StartTime,omitempty"`
    // TitleId unique identifier of the Game Server Instance for this lobby
    TitleId string `json:"TitleId,omitempty"`
}

// GetMatchmakerGameModesRequest these details are used by the PlayFab matchmaking service to determine if an existing Game Server Instance has room for
// additional users, and by the PlayFab game server management service to determine when a new Game Server Host should be
// created in order to prevent excess load on existing Hosts.
type GetMatchmakerGameModesRequestModel struct {
    // BuildVersion previously uploaded build version for which game modes are being requested
    BuildVersion string `json:"BuildVersion,omitempty"`
}

// GetMatchmakerGameModesResult 
type GetMatchmakerGameModesResultModel struct {
    // GameModes array of game modes available for the specified build
    GameModes []GameModeInfoModel `json:"GameModes,omitempty"`
}

// GetPlayedTitleListRequest useful for identifying titles of which the player's data will be deleted by DeleteMasterPlayer.
type GetPlayedTitleListRequestModel struct {
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// GetPlayedTitleListResult 
type GetPlayedTitleListResultModel struct {
    // TitleIds list of titles the player has played
    TitleIds []string `json:"TitleIds,omitempty"`
}

// GetPlayerIdFromAuthTokenRequest gets a player ID from an auth token. The token expires after 30 minutes and cannot be used to look up a player when
// expired.
type GetPlayerIdFromAuthTokenRequestModel struct {
    // Token the auth token of the player requesting the password reset.
    Token string `json:"Token,omitempty"`
    // TokenType the type of auth token of the player requesting the password reset.
    TokenType AuthTokenType `json:"TokenType,omitempty"`
}

// GetPlayerIdFromAuthTokenResult 
type GetPlayerIdFromAuthTokenResultModel struct {
    // PlayFabId the player ID from the token passed in
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// GetPlayerProfileRequest this API allows for access to details regarding a user in the PlayFab service, usually for purposes of customer support.
// Note that data returned may be Personally Identifying Information (PII), such as email address, and so care should be
// taken in how this data is stored and managed. Since this call will always return the relevant information for users who
// have accessed the title, the recommendation is to not store this data locally.
type GetPlayerProfileRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
    // ProfileConstraints if non-null, this determines which properties of the resulting player profiles to return. For API calls from the client,
// only the allowed client profile properties for the title may be requested. These allowed properties are configured in
// the Game Manager "Client Profile Options" tab in the "Settings" section.
    ProfileConstraints *PlayerProfileViewConstraintsModel `json:"ProfileConstraints,omitempty"`
}

// GetPlayerProfileResult 
type GetPlayerProfileResultModel struct {
    // PlayerProfile the profile of the player. This profile is not guaranteed to be up-to-date. For a new player, this profile will not
// exist.
    PlayerProfile *PlayerProfileModelModel `json:"PlayerProfile,omitempty"`
}

// GetPlayerSegmentsResult 
type GetPlayerSegmentsResultModel struct {
    // Segments array of segments the requested player currently belongs to.
    Segments []GetSegmentResultModel `json:"Segments,omitempty"`
}

// GetPlayerSharedSecretsRequest player Shared Secret Keys are used for the call to Client/GetTitlePublicKey, which exchanges the shared secret for an
// RSA CSP blob to be used to encrypt the payload of account creation requests when that API requires a signature header.
type GetPlayerSharedSecretsRequestModel struct {
}

// GetPlayerSharedSecretsResult 
type GetPlayerSharedSecretsResultModel struct {
    // SharedSecrets the player shared secret to use when calling Client/GetTitlePublicKey
    SharedSecrets []SharedSecretModel `json:"SharedSecrets,omitempty"`
}

// GetPlayersInSegmentRequest initial request must contain at least a Segment ID. Subsequent requests must contain the Segment ID as well as the
// Continuation Token. Failure to send the Continuation Token will result in a new player segment list being generated.
// Each time the Continuation Token is passed in the length of the Total Seconds to Live is refreshed. If too much time
// passes between requests to the point that a subsequent request is past the Total Seconds to Live an error will be
// returned and paging will be terminated. This API is resource intensive and should not be used in scenarios which might
// generate high request volumes. Only one request to this API at a time should be made per title. Concurrent requests to
// the API may be rejected with the APIConcurrentRequestLimitExceeded error.
type GetPlayersInSegmentRequestModel struct {
    // ContinuationToken continuation token if retrieving subsequent pages of results.
    ContinuationToken string `json:"ContinuationToken,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // MaxBatchSize maximum number of profiles to load. Default is 1,000. Maximum is 10,000.
    MaxBatchSize uint32 `json:"MaxBatchSize,omitempty"`
    // SecondsToLive number of seconds to keep the continuation token active. After token expiration it is not possible to continue paging
// results. Default is 300 (5 minutes). Maximum is 1,800 (30 minutes).
    SecondsToLive uint32 `json:"SecondsToLive,omitempty"`
    // SegmentId unique identifier for this segment.
    SegmentId string `json:"SegmentId,omitempty"`
}

// GetPlayersInSegmentResult 
type GetPlayersInSegmentResultModel struct {
    // ContinuationToken continuation token to use to retrieve subsequent pages of results. If token returns null there are no more results.
    ContinuationToken string `json:"ContinuationToken,omitempty"`
    // PlayerProfiles array of player profiles in this segment.
    PlayerProfiles []PlayerProfileModel `json:"PlayerProfiles,omitempty"`
    // ProfilesInSegment count of profiles matching this segment.
    ProfilesInSegment int32 `json:"ProfilesInSegment,omitempty"`
}

// GetPlayersSegmentsRequest 
type GetPlayersSegmentsRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// GetPlayerStatisticDefinitionsRequest 
type GetPlayerStatisticDefinitionsRequestModel struct {
}

// GetPlayerStatisticDefinitionsResult statistics are numeric values, with each statistic in the title also generating a leaderboard. The ResetInterval defines
// the period of time at which the leaderboard for the statistic will automatically reset. Upon reset, the statistic
// updates to a new version with no values (effectively removing all players from the leaderboard). The previous version's
// statistic values are also archived for retrieval, if needed (see GetPlayerStatisticVersions). Statistics not created via
// a call to CreatePlayerStatisticDefinition by default have a VersionChangeInterval of Never, meaning they do not reset on
// a schedule, but they can be set to do so via a call to UpdatePlayerStatisticDefinition. Once a statistic has been reset
// (sometimes referred to as versioned or incremented), the previous version can still be written to for up a short,
// pre-defined period (currently 10 seconds), to prevent issues with levels completing around the time of the reset. Also,
// once reset, the historical statistics for players in the title may be retrieved using the URL specified in the version
// information (GetPlayerStatisticVersions). The AggregationMethod defines what action is taken when a new statistic value
// is submitted - always update with the new value (Last), use the highest of the old and new values (Max), use the
// smallest (Min), or add them together (Sum).
type GetPlayerStatisticDefinitionsResultModel struct {
    // Statistics the player statistic definitions for the title
    Statistics []PlayerStatisticDefinitionModel `json:"Statistics,omitempty"`
}

// GetPlayerStatisticVersionsRequest 
type GetPlayerStatisticVersionsRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // StatisticName unique name of the statistic
    StatisticName string `json:"StatisticName,omitempty"`
}

// GetPlayerStatisticVersionsResult statistics are numeric values, with each statistic in the title also generating a leaderboard. The information returned
// in the results defines the state of a specific version of a statistic, including when it was or will become the
// currently active version, when it will (or did) become a previous version, and its archival state if it is no longer the
// active version. For a statistic which has been reset, once the archival status is Complete, the full set of statistics
// for all players in the leaderboard for that version may be retrieved via the ArchiveDownloadUrl. Statistics which have
// not been reset (incremented/versioned) will only have a single version which is not scheduled to reset.
type GetPlayerStatisticVersionsResultModel struct {
    // StatisticVersions version change history of the statistic
    StatisticVersions []PlayerStatisticVersionModel `json:"StatisticVersions,omitempty"`
}

// GetPlayerTagsRequest this API will return a list of canonical tags which includes both namespace and tag's name. If namespace is not
// provided, the result is a list of all canonical tags. TagName can be used for segmentation and Namespace is limited to
// 128 characters.
type GetPlayerTagsRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Namespace optional namespace to filter results by
    Namespace string `json:"Namespace,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// GetPlayerTagsResult 
type GetPlayerTagsResultModel struct {
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
    // Tags canonical tags (including namespace and tag's name) for the requested user
    Tags []string `json:"Tags,omitempty"`
}

// GetPolicyRequest views the requested policy. Today, the only supported policy is 'ApiPolicy'.
type GetPolicyRequestModel struct {
    // PolicyName the name of the policy to read. Only supported name is 'ApiPolicy'.
    PolicyName string `json:"PolicyName,omitempty"`
}

// GetPolicyResponse 
type GetPolicyResponseModel struct {
    // PolicyName the name of the policy read.
    PolicyName string `json:"PolicyName,omitempty"`
    // Statements the statements in the requested policy.
    Statements []PermissionStatementModel `json:"Statements,omitempty"`
}

// GetPublisherDataRequest this API is designed to return publisher-specific values which can be read, but not written to, by the client. This data
// is shared across all titles assigned to a particular publisher, and can be used for cross-game coordination. Only titles
// assigned to a publisher can use this API. For more information email helloplayfab@microsoft.com. This AdminAPI call for
// getting title data guarantees no delay in between update and retrieval of newly set data.
type GetPublisherDataRequestModel struct {
    // Keys array of keys to get back data from the Publisher data blob, set by the admin tools
    Keys []string `json:"Keys,omitempty"`
}

// GetPublisherDataResult 
type GetPublisherDataResultModel struct {
    // Data a dictionary object of key / value pairs
    Data map[string]string `json:"Data,omitempty"`
}

// GetRandomResultTablesRequest 
type GetRandomResultTablesRequestModel struct {
    // CatalogVersion catalog version to fetch tables from. Use default catalog version if null
    CatalogVersion string `json:"CatalogVersion,omitempty"`
}

// GetRandomResultTablesResult 
type GetRandomResultTablesResultModel struct {
    // Tables array of random result tables currently available
    Tables map[string]RandomResultTableListingModel `json:"Tables,omitempty"`
}

// GetSegmentResult 
type GetSegmentResultModel struct {
    // ABTestParent identifier of the segments AB Test, if it is attached to one.
    ABTestParent string `json:"ABTestParent,omitempty"`
    // Id unique identifier for this segment.
    Id string `json:"Id,omitempty"`
    // Name segment name.
    Name string `json:"Name,omitempty"`
}

// GetServerBuildInfoRequest 
type GetServerBuildInfoRequestModel struct {
    // BuildId unique identifier of the previously uploaded build executable for which information is being requested
    BuildId string `json:"BuildId,omitempty"`
}

// GetServerBuildInfoResult information about a particular server build
type GetServerBuildInfoResultModel struct {
    // ActiveRegions array of regions where this build can used, when it is active
    ActiveRegions []Region `json:"ActiveRegions,omitempty"`
    // BuildId unique identifier for this build executable
    BuildId string `json:"BuildId,omitempty"`
    // Comment developer comment(s) for this build
    Comment string `json:"Comment,omitempty"`
    // ErrorMessage error message, if any, about this build
    ErrorMessage string `json:"ErrorMessage,omitempty"`
    // MaxGamesPerHost maximum number of game server instances that can run on a single host machine
    MaxGamesPerHost int32 `json:"MaxGamesPerHost,omitempty"`
    // MinFreeGameSlots minimum capacity of additional game server instances that can be started before the autoscaling service starts new host
// machines (given the number of current running host machines and game server instances)
    MinFreeGameSlots int32 `json:"MinFreeGameSlots,omitempty"`
    // Status the current status of the build validation and processing steps
    Status GameBuildStatus `json:"Status,omitempty"`
    // Timestamp time this build was last modified (or uploaded, if this build has never been modified)
    Timestamp time.Time `json:"Timestamp,omitempty"`
    // TitleId unique identifier for the title, found in the Settings > Game Properties section of the PlayFab developer site when a
// title has been selected.
    TitleId string `json:"TitleId,omitempty"`
}

// GetServerBuildUploadURLRequest 
type GetServerBuildUploadURLRequestModel struct {
    // BuildId unique identifier of the game server build to upload
    BuildId string `json:"BuildId,omitempty"`
}

// GetServerBuildUploadURLResult 
type GetServerBuildUploadURLResultModel struct {
    // URL pre-authorized URL for uploading the game server build package
    URL string `json:"URL,omitempty"`
}

// GetStoreItemsRequest a store contains an array of references to items defined in the catalog, along with the prices for the item, in both
// real world and virtual currencies. These prices act as an override to any prices defined in the catalog. In this way,
// the base definitions of the items may be defined in the catalog, with all associated properties, while the pricing can
// be set for each store, as needed. This allows for subsets of goods to be defined for different purposes (in order to
// simplify showing some, but not all catalog items to users, based upon different characteristics), along with unique
// prices. Note that all prices defined in the catalog and store definitions for the item are considered valid, and that a
// compromised client can be made to send a request for an item based upon any of these definitions. If no price is
// specified in the store for an item, the price set in the catalog should be displayed to the user.
type GetStoreItemsRequestModel struct {
    // CatalogVersion catalog version to store items from. Use default catalog version if null
    CatalogVersion string `json:"CatalogVersion,omitempty"`
    // StoreId unqiue identifier for the store which is being requested.
    StoreId string `json:"StoreId,omitempty"`
}

// GetStoreItemsResult 
type GetStoreItemsResultModel struct {
    // CatalogVersion the base catalog that this store is a part of.
    CatalogVersion string `json:"CatalogVersion,omitempty"`
    // MarketingData additional data about the store.
    MarketingData *StoreMarketingModelModel `json:"MarketingData,omitempty"`
    // Source how the store was last updated (Admin or a third party).
    Source SourceType `json:"Source,omitempty"`
    // Store array of items which can be purchased from this store.
    Store []StoreItemModel `json:"Store,omitempty"`
    // StoreId the ID of this store.
    StoreId string `json:"StoreId,omitempty"`
}

// GetTaskInstanceRequest the result includes detail information that's specific to a CloudScript task. Only CloudScript tasks configured as "Run
// Cloud Script function once" will be retrieved. To get a list of task instances by task, status, or time range, use
// GetTaskInstances.
type GetTaskInstanceRequestModel struct {
    // TaskInstanceId iD of the requested task instance.
    TaskInstanceId string `json:"TaskInstanceId,omitempty"`
}

// GetTaskInstancesRequest only the most recent 100 task instances are returned, ordered by start time descending. The results are generic basic
// information for task instances. To get detail information specific to each task type, use Get*TaskInstance based on its
// corresponding task type.
type GetTaskInstancesRequestModel struct {
    // StartedAtRangeFrom optional range-from filter for task instances' StartedAt timestamp.
    StartedAtRangeFrom time.Time `json:"StartedAtRangeFrom,omitempty"`
    // StartedAtRangeTo optional range-to filter for task instances' StartedAt timestamp.
    StartedAtRangeTo time.Time `json:"StartedAtRangeTo,omitempty"`
    // StatusFilter optional filter for task instances that are of a specific status.
    StatusFilter TaskInstanceStatus `json:"StatusFilter,omitempty"`
    // TaskIdentifier name or ID of the task whose instances are being queried. If not specified, return all task instances that satisfy
// conditions set by other filters.
    TaskIdentifier *NameIdentifierModel `json:"TaskIdentifier,omitempty"`
}

// GetTaskInstancesResult 
type GetTaskInstancesResultModel struct {
    // Summaries basic status summaries of the queried task instances. Empty If no task instances meets the filter criteria. To get
// detailed status summary, use Get*TaskInstance API according to task type (e.g.
// GetActionsOnPlayersInSegmentTaskInstance).
    Summaries []TaskInstanceBasicSummaryModel `json:"Summaries,omitempty"`
}

// GetTasksRequest 
type GetTasksRequestModel struct {
    // Identifier provide either the task ID or the task name to get a specific task. If not specified, return all defined tasks.
    Identifier *NameIdentifierModel `json:"Identifier,omitempty"`
}

// GetTasksResult 
type GetTasksResultModel struct {
    // Tasks result tasks. Empty if there is no task found.
    Tasks []ScheduledTaskModel `json:"Tasks,omitempty"`
}

// GetTitleDataRequest this API method is designed to return title specific values which can be read by the client. For example, a developer
// could choose to store values which modify the user experience, such as enemy spawn rates, weapon strengths, movement
// speeds, etc. This allows a developer to update the title without the need to create, test, and ship a new build. If an
// override label is specified in the request, the overrides are applied automatically and returned with the title data.
// Note that due to caching, there may up to a minute delay in between updating title data and a query returning the newest
// value.
type GetTitleDataRequestModel struct {
    // Keys specific keys to search for in the title data (leave null to get all keys)
    Keys []string `json:"Keys,omitempty"`
    // OverrideLabel optional field that specifies the name of an override. This value is ignored when used by the game client; otherwise,
// the overrides are applied automatically to the title data.
    OverrideLabel string `json:"OverrideLabel,omitempty"`
}

// GetTitleDataResult 
type GetTitleDataResultModel struct {
    // Data a dictionary object of key / value pairs
    Data map[string]string `json:"Data,omitempty"`
}

// GetUserBansRequest get all bans for a user, including inactive and expired bans.
type GetUserBansRequestModel struct {
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// GetUserBansResult 
type GetUserBansResultModel struct {
    // BanData information about the bans
    BanData []BanInfoModel `json:"BanData,omitempty"`
}

// GetUserDataRequest data is stored as JSON key-value pairs. If the Keys parameter is provided, the data object returned will only contain
// the data specific to the indicated Keys. Otherwise, the full set of custom user data will be returned.
type GetUserDataRequestModel struct {
    // IfChangedFromDataVersion the version that currently exists according to the caller. The call will return the data for all of the keys if the
// version in the system is greater than this.
    IfChangedFromDataVersion uint32 `json:"IfChangedFromDataVersion,omitempty"`
    // Keys specific keys to search for in the custom user data.
    Keys []string `json:"Keys,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// GetUserDataResult 
type GetUserDataResultModel struct {
    // Data user specific data for this title.
    Data map[string]UserDataRecordModel `json:"Data,omitempty"`
    // DataVersion indicates the current version of the data that has been set. This is incremented with every set call for that type of
// data (read-only, internal, etc). This version can be provided in Get calls to find updated data.
    DataVersion uint32 `json:"DataVersion,omitempty"`
    // PlayFabId playFab unique identifier of the user whose custom data is being returned.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// GetUserInventoryRequest all items currently in the user inventory will be returned, irrespective of how they were acquired (via purchasing,
// grants, coupons, etc.). Items that are expired, fully consumed, or are no longer valid are not considered to be in the
// user's current inventory, and so will not be not included.
type GetUserInventoryRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// GetUserInventoryResult 
type GetUserInventoryResultModel struct {
    // Inventory array of inventory items belonging to the user.
    Inventory []ItemInstanceModel `json:"Inventory,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
    // VirtualCurrency array of virtual currency balance(s) belonging to the user.
    VirtualCurrency map[string]int32 `json:"VirtualCurrency,omitempty"`
    // VirtualCurrencyRechargeTimes array of remaining times and timestamps for virtual currencies.
    VirtualCurrencyRechargeTimes map[string]VirtualCurrencyRechargeTimeModel `json:"VirtualCurrencyRechargeTimes,omitempty"`
}

// GrantedItemInstance result of granting an item to a user. Note, to retrieve additional information for an item such as Tags, Description
// that are the same across all instances of the item, a call to GetCatalogItems is required. The ItemID of can be matched
// to a catalog entry, which contains the additional information. Also note that Custom Data is only set when the User's
// specific instance has updated the CustomData via a call to UpdateUserInventoryItemCustomData. Other fields such as
// UnitPrice and UnitCurrency are only set when the item was granted via a purchase.
type GrantedItemInstanceModel struct {
    // Annotation game specific comment associated with this instance when it was added to the user inventory.
    Annotation string `json:"Annotation,omitempty"`
    // BundleContents array of unique items that were awarded when this catalog item was purchased.
    BundleContents []string `json:"BundleContents,omitempty"`
    // BundleParent unique identifier for the parent inventory item, as defined in the catalog, for object which were added from a bundle or
// container.
    BundleParent string `json:"BundleParent,omitempty"`
    // CatalogVersion catalog version for the inventory item, when this instance was created.
    CatalogVersion string `json:"CatalogVersion,omitempty"`
    // CharacterId unique PlayFab assigned ID for a specific character owned by a user
    CharacterId string `json:"CharacterId,omitempty"`
    // CustomData a set of custom key-value pairs on the instance of the inventory item, which is not to be confused with the catalog
// item's custom data.
    CustomData map[string]string `json:"CustomData,omitempty"`
    // DisplayName catalogItem.DisplayName at the time this item was purchased.
    DisplayName string `json:"DisplayName,omitempty"`
    // Expiration timestamp for when this instance will expire.
    Expiration time.Time `json:"Expiration,omitempty"`
    // ItemClass class name for the inventory item, as defined in the catalog.
    ItemClass string `json:"ItemClass,omitempty"`
    // ItemId unique identifier for the inventory item, as defined in the catalog.
    ItemId string `json:"ItemId,omitempty"`
    // ItemInstanceId unique item identifier for this specific instance of the item.
    ItemInstanceId string `json:"ItemInstanceId,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
    // PurchaseDate timestamp for when this instance was purchased.
    PurchaseDate time.Time `json:"PurchaseDate,omitempty"`
    // RemainingUses total number of remaining uses, if this is a consumable item.
    RemainingUses int32 `json:"RemainingUses,omitempty"`
    // Result result of this operation.
    Result bool `json:"Result"`
    // UnitCurrency currency type for the cost of the catalog item. Not available when granting items.
    UnitCurrency string `json:"UnitCurrency,omitempty"`
    // UnitPrice cost of the catalog item in the given currency. Not available when granting items.
    UnitPrice uint32 `json:"UnitPrice,omitempty"`
    // UsesIncrementedBy the number of uses that were added or removed to this item in this call.
    UsesIncrementedBy int32 `json:"UsesIncrementedBy,omitempty"`
}

// GrantItemsToUsersRequest this function directly adds inventory items to user inventories. As a result of this operations, the user will not be
// charged any transaction fee, regardless of the inventory item catalog definition. Please note that the processing time
// for inventory grants and purchases increases fractionally the more items are in the inventory, and the more items are in
// the grant/purchase operation.
type GrantItemsToUsersRequestModel struct {
    // CatalogVersion catalog version from which items are to be granted.
    CatalogVersion string `json:"CatalogVersion,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // ItemGrants array of items to grant and the users to whom the items are to be granted.
    ItemGrants []ItemGrantModel `json:"ItemGrants,omitempty"`
}

// GrantItemsToUsersResult please note that the order of the items in the response may not match the order of items in the request.
type GrantItemsToUsersResultModel struct {
    // ItemGrantResults array of items granted to users.
    ItemGrantResults []GrantedItemInstanceModel `json:"ItemGrantResults,omitempty"`
}

// IncrementLimitedEditionItemAvailabilityRequest this operation will increment the global counter for the number of these items available. This number cannot be
// decremented, except by actual grants.
type IncrementLimitedEditionItemAvailabilityRequestModel struct {
    // Amount amount to increase availability by.
    Amount int32 `json:"Amount,omitempty"`
    // CatalogVersion which catalog is being updated. If null, uses the default catalog.
    CatalogVersion string `json:"CatalogVersion,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // ItemId the item which needs more availability.
    ItemId string `json:"ItemId,omitempty"`
}

// IncrementLimitedEditionItemAvailabilityResult 
type IncrementLimitedEditionItemAvailabilityResultModel struct {
}

// IncrementPlayerStatisticVersionRequest statistics are numeric values, with each statistic in the title also generating a leaderboard. When this call is made on
// a given statistic, this forces a reset of that statistic. Upon reset, the statistic updates to a new version with no
// values (effectively removing all players from the leaderboard). The previous version's statistic values are also
// archived for retrieval, if needed (see GetPlayerStatisticVersions). Statistics not created via a call to
// CreatePlayerStatisticDefinition by default have a VersionChangeInterval of Never, meaning they do not reset on a
// schedule, but they can be set to do so via a call to UpdatePlayerStatisticDefinition. Once a statistic has been reset
// (sometimes referred to as versioned or incremented), the now-previous version can still be written to for up a short,
// pre-defined period (currently 10 seconds), to prevent issues with levels completing around the time of the reset. Also,
// once reset, the historical statistics for players in the title may be retrieved using the URL specified in the version
// information (GetPlayerStatisticVersions).
type IncrementPlayerStatisticVersionRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // StatisticName unique name of the statistic
    StatisticName string `json:"StatisticName,omitempty"`
}

// IncrementPlayerStatisticVersionResult 
type IncrementPlayerStatisticVersionResultModel struct {
    // StatisticVersion version change history of the statistic
    StatisticVersion *PlayerStatisticVersionModel `json:"StatisticVersion,omitempty"`
}

// InsightsScalingTaskParameter 
type InsightsScalingTaskParameterModel struct {
    // Level insights Performance Level to scale to.
    Level int32 `json:"Level,omitempty"`
}

// ItemGrant 
type ItemGrantModel struct {
    // Annotation string detailing any additional information concerning this operation.
    Annotation string `json:"Annotation,omitempty"`
    // CharacterId unique PlayFab assigned ID for a specific character owned by a user
    CharacterId string `json:"CharacterId,omitempty"`
    // Data key-value pairs to be written to the custom data. Note that keys are trimmed of whitespace, are limited in size, and may
// not begin with a '!' character or be null.
    Data map[string]string `json:"Data,omitempty"`
    // ItemId unique identifier of the catalog item to be granted to the user.
    ItemId string `json:"ItemId,omitempty"`
    // KeysToRemove optional list of Data-keys to remove from UserData. Some SDKs cannot insert null-values into Data due to language
// constraints. Use this to delete the keys directly.
    KeysToRemove []string `json:"KeysToRemove,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// ItemInstance a unique instance of an item in a user's inventory. Note, to retrieve additional information for an item such as Tags,
// Description that are the same across all instances of the item, a call to GetCatalogItems is required. The ItemID of can
// be matched to a catalog entry, which contains the additional information. Also note that Custom Data is only set when
// the User's specific instance has updated the CustomData via a call to UpdateUserInventoryItemCustomData. Other fields
// such as UnitPrice and UnitCurrency are only set when the item was granted via a purchase.
type ItemInstanceModel struct {
    // Annotation game specific comment associated with this instance when it was added to the user inventory.
    Annotation string `json:"Annotation,omitempty"`
    // BundleContents array of unique items that were awarded when this catalog item was purchased.
    BundleContents []string `json:"BundleContents,omitempty"`
    // BundleParent unique identifier for the parent inventory item, as defined in the catalog, for object which were added from a bundle or
// container.
    BundleParent string `json:"BundleParent,omitempty"`
    // CatalogVersion catalog version for the inventory item, when this instance was created.
    CatalogVersion string `json:"CatalogVersion,omitempty"`
    // CustomData a set of custom key-value pairs on the instance of the inventory item, which is not to be confused with the catalog
// item's custom data.
    CustomData map[string]string `json:"CustomData,omitempty"`
    // DisplayName catalogItem.DisplayName at the time this item was purchased.
    DisplayName string `json:"DisplayName,omitempty"`
    // Expiration timestamp for when this instance will expire.
    Expiration time.Time `json:"Expiration,omitempty"`
    // ItemClass class name for the inventory item, as defined in the catalog.
    ItemClass string `json:"ItemClass,omitempty"`
    // ItemId unique identifier for the inventory item, as defined in the catalog.
    ItemId string `json:"ItemId,omitempty"`
    // ItemInstanceId unique item identifier for this specific instance of the item.
    ItemInstanceId string `json:"ItemInstanceId,omitempty"`
    // PurchaseDate timestamp for when this instance was purchased.
    PurchaseDate time.Time `json:"PurchaseDate,omitempty"`
    // RemainingUses total number of remaining uses, if this is a consumable item.
    RemainingUses int32 `json:"RemainingUses,omitempty"`
    // UnitCurrency currency type for the cost of the catalog item. Not available when granting items.
    UnitCurrency string `json:"UnitCurrency,omitempty"`
    // UnitPrice cost of the catalog item in the given currency. Not available when granting items.
    UnitPrice uint32 `json:"UnitPrice,omitempty"`
    // UsesIncrementedBy the number of uses that were added or removed to this item in this call.
    UsesIncrementedBy int32 `json:"UsesIncrementedBy,omitempty"`
}

// LinkedPlatformAccountModel 
type LinkedPlatformAccountModelModel struct {
    // Email linked account email of the user on the platform, if available
    Email string `json:"Email,omitempty"`
    // Platform authentication platform
    Platform LoginIdentityProvider `json:"Platform,omitempty"`
    // PlatformUserId unique account identifier of the user on the platform
    PlatformUserId string `json:"PlatformUserId,omitempty"`
    // Username linked account username of the user on the platform, if available
    Username string `json:"Username,omitempty"`
}

// ListBuildsRequest 
type ListBuildsRequestModel struct {
}

// ListBuildsResult 
type ListBuildsResultModel struct {
    // Builds array of uploaded game server builds
    Builds []GetServerBuildInfoResultModel `json:"Builds,omitempty"`
}

// ListOpenIdConnectionRequest 
type ListOpenIdConnectionRequestModel struct {
}

// ListOpenIdConnectionResponse 
type ListOpenIdConnectionResponseModel struct {
    // Connections the list of Open ID Connections
    Connections []OpenIdConnectionModel `json:"Connections,omitempty"`
}

// ListVirtualCurrencyTypesRequest 
type ListVirtualCurrencyTypesRequestModel struct {
}

// ListVirtualCurrencyTypesResult 
type ListVirtualCurrencyTypesResultModel struct {
    // VirtualCurrencies list of virtual currency names defined for this title
    VirtualCurrencies []VirtualCurrencyDataModel `json:"VirtualCurrencies,omitempty"`
}

// LocationModel 
type LocationModelModel struct {
    // City city name.
    City string `json:"City,omitempty"`
    // ContinentCode the two-character continent code for this location
    ContinentCode ContinentCode `json:"ContinentCode,omitempty"`
    // CountryCode the two-character ISO 3166-1 country code for the country associated with the location
    CountryCode CountryCode `json:"CountryCode,omitempty"`
    // Latitude latitude coordinate of the geographic location.
    Latitude float64 `json:"Latitude,omitempty"`
    // Longitude longitude coordinate of the geographic location.
    Longitude float64 `json:"Longitude,omitempty"`
}

// LoginIdentityProvider 
type LoginIdentityProvider string
  
const (
     LoginIdentityProviderUnknown LoginIdentityProvider = "Unknown"
     LoginIdentityProviderPlayFab LoginIdentityProvider = "PlayFab"
     LoginIdentityProviderCustom LoginIdentityProvider = "Custom"
     LoginIdentityProviderGameCenter LoginIdentityProvider = "GameCenter"
     LoginIdentityProviderGooglePlay LoginIdentityProvider = "GooglePlay"
     LoginIdentityProviderSteam LoginIdentityProvider = "Steam"
     LoginIdentityProviderXBoxLive LoginIdentityProvider = "XBoxLive"
     LoginIdentityProviderPSN LoginIdentityProvider = "PSN"
     LoginIdentityProviderKongregate LoginIdentityProvider = "Kongregate"
     LoginIdentityProviderFacebook LoginIdentityProvider = "Facebook"
     LoginIdentityProviderIOSDevice LoginIdentityProvider = "IOSDevice"
     LoginIdentityProviderAndroidDevice LoginIdentityProvider = "AndroidDevice"
     LoginIdentityProviderTwitch LoginIdentityProvider = "Twitch"
     LoginIdentityProviderWindowsHello LoginIdentityProvider = "WindowsHello"
     LoginIdentityProviderGameServer LoginIdentityProvider = "GameServer"
     LoginIdentityProviderCustomServer LoginIdentityProvider = "CustomServer"
     LoginIdentityProviderNintendoSwitch LoginIdentityProvider = "NintendoSwitch"
     LoginIdentityProviderFacebookInstantGames LoginIdentityProvider = "FacebookInstantGames"
     LoginIdentityProviderOpenIdConnect LoginIdentityProvider = "OpenIdConnect"
     LoginIdentityProviderApple LoginIdentityProvider = "Apple"
     LoginIdentityProviderNintendoSwitchAccount LoginIdentityProvider = "NintendoSwitchAccount"
      )
// LogStatement 
type LogStatementModel struct {
    // Data optional object accompanying the message as contextual information
    Data interface{} `json:"Data,omitempty"`
    // Level 'Debug', 'Info', or 'Error'
    Level string `json:"Level,omitempty"`
    // Message 
    Message string `json:"Message,omitempty"`
}

// LookupUserAccountInfoRequest this API allows for access to details regarding a user in the PlayFab service, usually for purposes of customer support.
// Note that data returned may be Personally Identifying Information (PII), such as email address, and so care should be
// taken in how this data is stored and managed. Since this call will always return the relevant information for users who
// have accessed the title, the recommendation is to not store this data locally.
type LookupUserAccountInfoRequestModel struct {
    // Email user email address attached to their account
    Email string `json:"Email,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
    // TitleDisplayName title specific username to match against existing user accounts
    TitleDisplayName string `json:"TitleDisplayName,omitempty"`
    // Username playFab username for the account (3-20 characters)
    Username string `json:"Username,omitempty"`
}

// LookupUserAccountInfoResult 
type LookupUserAccountInfoResultModel struct {
    // UserInfo user info for the user matching the request
    UserInfo *UserAccountInfoModel `json:"UserInfo,omitempty"`
}

// MembershipModel 
type MembershipModelModel struct {
    // IsActive whether this membership is active. That is, whether the MembershipExpiration time has been reached.
    IsActive bool `json:"IsActive"`
    // MembershipExpiration the time this membership expires
    MembershipExpiration time.Time `json:"MembershipExpiration,omitempty"`
    // MembershipId the id of the membership
    MembershipId string `json:"MembershipId,omitempty"`
    // OverrideExpiration membership expirations can be explicitly overridden (via game manager or the admin api). If this membership has been
// overridden, this will be the new expiration time.
    OverrideExpiration time.Time `json:"OverrideExpiration,omitempty"`
    // Subscriptions the list of subscriptions that this player has for this membership
    Subscriptions []SubscriptionModelModel `json:"Subscriptions,omitempty"`
}

// ModifyMatchmakerGameModesRequest these details are used by the PlayFab matchmaking service to determine if an existing Game Server Instance has room for
// additional users, and by the PlayFab game server management service to determine when a new Game Server Host should be
// created in order to prevent excess load on existing Hosts. This operation is not additive. Using it will cause the game
// mode definition for the game server executable in question to be created from scratch. If there is an existing game
// server mode definition for the given BuildVersion, it will be deleted and replaced with the data specified in this call.
type ModifyMatchmakerGameModesRequestModel struct {
    // BuildVersion previously uploaded build version for which game modes are being specified
    BuildVersion string `json:"BuildVersion,omitempty"`
    // GameModes array of game modes (Note: this will replace all game modes for the indicated build version)
    GameModes []GameModeInfoModel `json:"GameModes,omitempty"`
}

// ModifyMatchmakerGameModesResult 
type ModifyMatchmakerGameModesResultModel struct {
}

// ModifyServerBuildRequest 
type ModifyServerBuildRequestModel struct {
    // ActiveRegions array of regions where this build can used, when it is active
    ActiveRegions []Region `json:"ActiveRegions,omitempty"`
    // BuildId unique identifier of the previously uploaded build executable to be updated
    BuildId string `json:"BuildId,omitempty"`
    // CommandLineTemplate appended to the end of the command line when starting game servers
    CommandLineTemplate string `json:"CommandLineTemplate,omitempty"`
    // Comment developer comment(s) for this build
    Comment string `json:"Comment,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // ExecutablePath path to the game server executable. Defaults to gameserver.exe
    ExecutablePath string `json:"ExecutablePath,omitempty"`
    // MaxGamesPerHost maximum number of game server instances that can run on a single host machine
    MaxGamesPerHost int32 `json:"MaxGamesPerHost,omitempty"`
    // MinFreeGameSlots minimum capacity of additional game server instances that can be started before the autoscaling service starts new host
// machines (given the number of current running host machines and game server instances)
    MinFreeGameSlots int32 `json:"MinFreeGameSlots,omitempty"`
    // Timestamp new timestamp
    Timestamp time.Time `json:"Timestamp,omitempty"`
}

// ModifyServerBuildResult 
type ModifyServerBuildResultModel struct {
    // ActiveRegions array of regions where this build can used, when it is active
    ActiveRegions []Region `json:"ActiveRegions,omitempty"`
    // BuildId unique identifier for this build executable
    BuildId string `json:"BuildId,omitempty"`
    // CommandLineTemplate appended to the end of the command line when starting game servers
    CommandLineTemplate string `json:"CommandLineTemplate,omitempty"`
    // Comment developer comment(s) for this build
    Comment string `json:"Comment,omitempty"`
    // ExecutablePath path to the game server executable. Defaults to gameserver.exe
    ExecutablePath string `json:"ExecutablePath,omitempty"`
    // MaxGamesPerHost maximum number of game server instances that can run on a single host machine
    MaxGamesPerHost int32 `json:"MaxGamesPerHost,omitempty"`
    // MinFreeGameSlots minimum capacity of additional game server instances that can be started before the autoscaling service starts new host
// machines (given the number of current running host machines and game server instances)
    MinFreeGameSlots int32 `json:"MinFreeGameSlots,omitempty"`
    // Status the current status of the build validation and processing steps
    Status GameBuildStatus `json:"Status,omitempty"`
    // Timestamp time this build was last modified (or uploaded, if this build has never been modified)
    Timestamp time.Time `json:"Timestamp,omitempty"`
    // TitleId unique identifier for the title, found in the Settings > Game Properties section of the PlayFab developer site when a
// title has been selected.
    TitleId string `json:"TitleId,omitempty"`
}

// ModifyUserVirtualCurrencyResult 
type ModifyUserVirtualCurrencyResultModel struct {
    // Balance balance of the virtual currency after modification.
    Balance int32 `json:"Balance,omitempty"`
    // BalanceChange amount added or subtracted from the user's virtual currency. Maximum VC balance is Int32 (2,147,483,647). Any increase
// over this value will be discarded.
    BalanceChange int32 `json:"BalanceChange,omitempty"`
    // PlayFabId user currency was subtracted from.
    PlayFabId string `json:"PlayFabId,omitempty"`
    // VirtualCurrency name of the virtual currency which was modified.
    VirtualCurrency string `json:"VirtualCurrency,omitempty"`
}

// NameIdentifier identifier by either name or ID. Note that a name may change due to renaming, or reused after being deleted. ID is
// immutable and unique.
type NameIdentifierModel struct {
    // Id id Identifier, if present
    Id string `json:"Id,omitempty"`
    // Name name Identifier, if present
    Name string `json:"Name,omitempty"`
}

// OpenIdConnection 
type OpenIdConnectionModel struct {
    // ClientId the client ID given by the ID provider.
    ClientId string `json:"ClientId,omitempty"`
    // ClientSecret the client secret given by the ID provider.
    ClientSecret string `json:"ClientSecret,omitempty"`
    // ConnectionId a name for the connection to identify it within the title.
    ConnectionId string `json:"ConnectionId,omitempty"`
    // DiscoverConfiguration shows if data about the connection will be loaded from the issuer's discovery document
    DiscoverConfiguration bool `json:"DiscoverConfiguration"`
    // IssuerInformation information for an OpenID Connect provider.
    IssuerInformation *OpenIdIssuerInformationModel `json:"IssuerInformation,omitempty"`
}

// OpenIdIssuerInformation 
type OpenIdIssuerInformationModel struct {
    // AuthorizationUrl authorization endpoint URL to direct users to for signin.
    AuthorizationUrl string `json:"AuthorizationUrl,omitempty"`
    // Issuer the URL of the issuer of the tokens. This must match the exact URL of the issuer field in tokens.
    Issuer string `json:"Issuer,omitempty"`
    // JsonWebKeySet jSON Web Key Set for validating the signature of tokens.
    JsonWebKeySet interface{} `json:"JsonWebKeySet,omitempty"`
    // TokenUrl token endpoint URL for code verification.
    TokenUrl string `json:"TokenUrl,omitempty"`
}

// PermissionStatement 
type PermissionStatementModel struct {
    // Action the action this statement effects. The only supported action is 'Execute'.
    Action string `json:"Action,omitempty"`
    // ApiConditions additional conditions to be applied for API Resources.
    ApiConditions *ApiConditionModel `json:"ApiConditions,omitempty"`
    // Comment a comment about the statement. Intended solely for bookkeeping and debugging.
    Comment string `json:"Comment,omitempty"`
    // Effect the effect this statement will have. It could be either Allow or Deny
    Effect EffectType `json:"Effect,omitempty"`
    // Principal the principal this statement will effect. The only supported principal is '*'.
    Principal string `json:"Principal,omitempty"`
    // Resource the resource this statements effects. The only supported resources look like 'pfrn:api--*' for all apis, or
// 'pfrn:api--/Client/ConfirmPurchase' for specific apis.
    Resource string `json:"Resource,omitempty"`
}

// PlayerLinkedAccount 
type PlayerLinkedAccountModel struct {
    // Email linked account's email
    Email string `json:"Email,omitempty"`
    // Platform authentication platform
    Platform LoginIdentityProvider `json:"Platform,omitempty"`
    // PlatformUserId platform user identifier
    PlatformUserId string `json:"PlatformUserId,omitempty"`
    // Username linked account's username
    Username string `json:"Username,omitempty"`
}

// PlayerLocation 
type PlayerLocationModel struct {
    // City city of the player's geographic location.
    City string `json:"City,omitempty"`
    // ContinentCode the two-character continent code for this location
    ContinentCode ContinentCode `json:"ContinentCode,omitempty"`
    // CountryCode the two-character ISO 3166-1 country code for the country associated with the location
    CountryCode CountryCode `json:"CountryCode,omitempty"`
    // Latitude latitude coordinate of the player's geographic location.
    Latitude float64 `json:"Latitude,omitempty"`
    // Longitude longitude coordinate of the player's geographic location.
    Longitude float64 `json:"Longitude,omitempty"`
}

// PlayerProfile 
type PlayerProfileModel struct {
    // AdCampaignAttributions array of ad campaigns player has been attributed to
    AdCampaignAttributions []AdCampaignAttributionModel `json:"AdCampaignAttributions,omitempty"`
    // AvatarUrl image URL of the player's avatar.
    AvatarUrl string `json:"AvatarUrl,omitempty"`
    // BannedUntil banned until UTC Date. If permanent ban this is set for 20 years after the original ban date.
    BannedUntil time.Time `json:"BannedUntil,omitempty"`
    // ContactEmailAddresses array of contact email addresses associated with the player
    ContactEmailAddresses []ContactEmailInfoModel `json:"ContactEmailAddresses,omitempty"`
    // Created player record created
    Created time.Time `json:"Created,omitempty"`
    // DisplayName player Display Name
    DisplayName string `json:"DisplayName,omitempty"`
    // LastLogin last login
    LastLogin time.Time `json:"LastLogin,omitempty"`
    // LinkedAccounts array of third party accounts linked to this player
    LinkedAccounts []PlayerLinkedAccountModel `json:"LinkedAccounts,omitempty"`
    // Locations dictionary of player's locations by type.
    Locations map[string]PlayerLocationModel `json:"Locations,omitempty"`
    // Origination player account origination
    Origination LoginIdentityProvider `json:"Origination,omitempty"`
    // PlayerExperimentVariants list of player variants for experimentation
    PlayerExperimentVariants []string `json:"PlayerExperimentVariants,omitempty"`
    // PlayerId playFab Player ID
    PlayerId string `json:"PlayerId,omitempty"`
    // PlayerStatistics array of player statistics
    PlayerStatistics []PlayerStatisticModel `json:"PlayerStatistics,omitempty"`
    // PublisherId publisher this player belongs to
    PublisherId string `json:"PublisherId,omitempty"`
    // PushNotificationRegistrations array of configured push notification end points
    PushNotificationRegistrations []PushNotificationRegistrationModel `json:"PushNotificationRegistrations,omitempty"`
    // Statistics dictionary of player's statistics using only the latest version's value
    Statistics map[string]int32 `json:"Statistics,omitempty"`
    // Tags list of player's tags for segmentation.
    Tags []string `json:"Tags,omitempty"`
    // TitleId title ID this profile applies to
    TitleId string `json:"TitleId,omitempty"`
    // TotalValueToDateInUSD a sum of player's total purchases in USD across all currencies.
    TotalValueToDateInUSD uint32 `json:"TotalValueToDateInUSD,omitempty"`
    // ValuesToDate dictionary of player's total purchases by currency.
    ValuesToDate map[string]uint32 `json:"ValuesToDate,omitempty"`
    // VirtualCurrencyBalances dictionary of player's virtual currency balances
    VirtualCurrencyBalances map[string]int32 `json:"VirtualCurrencyBalances,omitempty"`
}

// PlayerProfileModel 
type PlayerProfileModelModel struct {
    // AdCampaignAttributions list of advertising campaigns the player has been attributed to
    AdCampaignAttributions []AdCampaignAttributionModelModel `json:"AdCampaignAttributions,omitempty"`
    // AvatarUrl uRL of the player's avatar image
    AvatarUrl string `json:"AvatarUrl,omitempty"`
    // BannedUntil if the player is currently banned, the UTC Date when the ban expires
    BannedUntil time.Time `json:"BannedUntil,omitempty"`
    // ContactEmailAddresses list of all contact email info associated with the player account
    ContactEmailAddresses []ContactEmailInfoModelModel `json:"ContactEmailAddresses,omitempty"`
    // Created player record created
    Created time.Time `json:"Created,omitempty"`
    // DisplayName player display name
    DisplayName string `json:"DisplayName,omitempty"`
    // ExperimentVariants list of experiment variants for the player.
    ExperimentVariants []string `json:"ExperimentVariants,omitempty"`
    // LastLogin uTC time when the player most recently logged in to the title
    LastLogin time.Time `json:"LastLogin,omitempty"`
    // LinkedAccounts list of all authentication systems linked to this player account
    LinkedAccounts []LinkedPlatformAccountModelModel `json:"LinkedAccounts,omitempty"`
    // Locations list of geographic locations from which the player has logged in to the title
    Locations []LocationModelModel `json:"Locations,omitempty"`
    // Memberships list of memberships for the player, along with whether are expired.
    Memberships []MembershipModelModel `json:"Memberships,omitempty"`
    // Origination player account origination
    Origination LoginIdentityProvider `json:"Origination,omitempty"`
    // PlayerId playFab player account unique identifier
    PlayerId string `json:"PlayerId,omitempty"`
    // PublisherId publisher this player belongs to
    PublisherId string `json:"PublisherId,omitempty"`
    // PushNotificationRegistrations list of configured end points registered for sending the player push notifications
    PushNotificationRegistrations []PushNotificationRegistrationModelModel `json:"PushNotificationRegistrations,omitempty"`
    // Statistics list of leaderboard statistic values for the player
    Statistics []StatisticModelModel `json:"Statistics,omitempty"`
    // Tags list of player's tags for segmentation
    Tags []TagModelModel `json:"Tags,omitempty"`
    // TitleId title ID this player profile applies to
    TitleId string `json:"TitleId,omitempty"`
    // TotalValueToDateInUSD sum of the player's purchases made with real-money currencies, converted to US dollars equivalent and represented as a
// whole number of cents (1/100 USD). For example, 999 indicates nine dollars and ninety-nine cents.
    TotalValueToDateInUSD uint32 `json:"TotalValueToDateInUSD,omitempty"`
    // ValuesToDate list of the player's lifetime purchase totals, summed by real-money currency
    ValuesToDate []ValueToDateModelModel `json:"ValuesToDate,omitempty"`
}

// PlayerProfileViewConstraints 
type PlayerProfileViewConstraintsModel struct {
    // ShowAvatarUrl whether to show player's avatar URL. Defaults to false
    ShowAvatarUrl bool `json:"ShowAvatarUrl"`
    // ShowBannedUntil whether to show the banned until time. Defaults to false
    ShowBannedUntil bool `json:"ShowBannedUntil"`
    // ShowCampaignAttributions whether to show campaign attributions. Defaults to false
    ShowCampaignAttributions bool `json:"ShowCampaignAttributions"`
    // ShowContactEmailAddresses whether to show contact email addresses. Defaults to false
    ShowContactEmailAddresses bool `json:"ShowContactEmailAddresses"`
    // ShowCreated whether to show the created date. Defaults to false
    ShowCreated bool `json:"ShowCreated"`
    // ShowDisplayName whether to show the display name. Defaults to false
    ShowDisplayName bool `json:"ShowDisplayName"`
    // ShowExperimentVariants whether to show player's experiment variants. Defaults to false
    ShowExperimentVariants bool `json:"ShowExperimentVariants"`
    // ShowLastLogin whether to show the last login time. Defaults to false
    ShowLastLogin bool `json:"ShowLastLogin"`
    // ShowLinkedAccounts whether to show the linked accounts. Defaults to false
    ShowLinkedAccounts bool `json:"ShowLinkedAccounts"`
    // ShowLocations whether to show player's locations. Defaults to false
    ShowLocations bool `json:"ShowLocations"`
    // ShowMemberships whether to show player's membership information. Defaults to false
    ShowMemberships bool `json:"ShowMemberships"`
    // ShowOrigination whether to show origination. Defaults to false
    ShowOrigination bool `json:"ShowOrigination"`
    // ShowPushNotificationRegistrations whether to show push notification registrations. Defaults to false
    ShowPushNotificationRegistrations bool `json:"ShowPushNotificationRegistrations"`
    // ShowStatistics reserved for future development
    ShowStatistics bool `json:"ShowStatistics"`
    // ShowTags whether to show tags. Defaults to false
    ShowTags bool `json:"ShowTags"`
    // ShowTotalValueToDateInUsd whether to show the total value to date in usd. Defaults to false
    ShowTotalValueToDateInUsd bool `json:"ShowTotalValueToDateInUsd"`
    // ShowValuesToDate whether to show the values to date. Defaults to false
    ShowValuesToDate bool `json:"ShowValuesToDate"`
}

// PlayerStatistic 
type PlayerStatisticModel struct {
    // Id statistic ID
    Id string `json:"Id,omitempty"`
    // Name statistic name
    Name string `json:"Name,omitempty"`
    // StatisticValue current statistic value
    StatisticValue int32 `json:"StatisticValue,omitempty"`
    // StatisticVersion statistic version (0 if not a versioned statistic)
    StatisticVersion int32 `json:"StatisticVersion,omitempty"`
}

// PlayerStatisticDefinition 
type PlayerStatisticDefinitionModel struct {
    // AggregationMethod the aggregation method to use in updating the statistic (defaults to last)
    AggregationMethod StatisticAggregationMethod `json:"AggregationMethod,omitempty"`
    // CurrentVersion current active version of the statistic, incremented each time the statistic resets
    CurrentVersion uint32 `json:"CurrentVersion,omitempty"`
    // StatisticName unique name of the statistic
    StatisticName string `json:"StatisticName,omitempty"`
    // VersionChangeInterval interval at which the values of the statistic for all players are reset automatically
    VersionChangeInterval StatisticResetIntervalOption `json:"VersionChangeInterval,omitempty"`
}

// PlayerStatisticVersion 
type PlayerStatisticVersionModel struct {
    // ActivationTime time when the statistic version became active
    ActivationTime time.Time `json:"ActivationTime,omitempty"`
    // ArchiveDownloadUrl uRL for the downloadable archive of player statistic values, if available
    ArchiveDownloadUrl string `json:"ArchiveDownloadUrl,omitempty"`
    // DeactivationTime time when the statistic version became inactive due to statistic version incrementing
    DeactivationTime time.Time `json:"DeactivationTime,omitempty"`
    // ScheduledActivationTime time at which the statistic version was scheduled to become active, based on the configured ResetInterval
    ScheduledActivationTime time.Time `json:"ScheduledActivationTime,omitempty"`
    // ScheduledDeactivationTime time at which the statistic version was scheduled to become inactive, based on the configured ResetInterval
    ScheduledDeactivationTime time.Time `json:"ScheduledDeactivationTime,omitempty"`
    // StatisticName name of the statistic when the version became active
    StatisticName string `json:"StatisticName,omitempty"`
    // Status status of the statistic version
    Status StatisticVersionStatus `json:"Status,omitempty"`
    // Version version of the statistic
    Version uint32 `json:"Version,omitempty"`
}

// PushNotificationPlatform 
type PushNotificationPlatform string
  
const (
     PushNotificationPlatformApplePushNotificationService PushNotificationPlatform = "ApplePushNotificationService"
     PushNotificationPlatformGoogleCloudMessaging PushNotificationPlatform = "GoogleCloudMessaging"
      )
// PushNotificationRegistration 
type PushNotificationRegistrationModel struct {
    // NotificationEndpointARN notification configured endpoint
    NotificationEndpointARN string `json:"NotificationEndpointARN,omitempty"`
    // Platform push notification platform
    Platform PushNotificationPlatform `json:"Platform,omitempty"`
}

// PushNotificationRegistrationModel 
type PushNotificationRegistrationModelModel struct {
    // NotificationEndpointARN notification configured endpoint
    NotificationEndpointARN string `json:"NotificationEndpointARN,omitempty"`
    // Platform push notification platform
    Platform PushNotificationPlatform `json:"Platform,omitempty"`
}

// PushSetupPlatform 
type PushSetupPlatform string
  
const (
     PushSetupPlatformGCM PushSetupPlatform = "GCM"
     PushSetupPlatformAPNS PushSetupPlatform = "APNS"
     PushSetupPlatformAPNS_SANDBOX PushSetupPlatform = "APNS_SANDBOX"
      )
// RandomResultTable 
type RandomResultTableModel struct {
    // Nodes child nodes that indicate what kind of drop table item this actually is.
    Nodes []ResultTableNodeModel `json:"Nodes,omitempty"`
    // TableId unique name for this drop table
    TableId string `json:"TableId,omitempty"`
}

// RandomResultTableListing 
type RandomResultTableListingModel struct {
    // CatalogVersion catalog version this table is associated with
    CatalogVersion string `json:"CatalogVersion,omitempty"`
    // Nodes child nodes that indicate what kind of drop table item this actually is.
    Nodes []ResultTableNodeModel `json:"Nodes,omitempty"`
    // TableId unique name for this drop table
    TableId string `json:"TableId,omitempty"`
}

// RefundPurchaseRequest 
type RefundPurchaseRequestModel struct {
    // OrderId unique order ID for the purchase in question.
    OrderId string `json:"OrderId,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
    // Reason the Reason parameter should correspond with the payment providers reason field, if they require one such as Facebook. In
// the case of Facebook this must match one of their refund or dispute resolution enums (See:
// https://developers.facebook.com/docs/payments/implementation-guide/handling-disputes-refunds)
    Reason string `json:"Reason,omitempty"`
}

// RefundPurchaseResponse 
type RefundPurchaseResponseModel struct {
    // PurchaseStatus the order's updated purchase status.
    PurchaseStatus string `json:"PurchaseStatus,omitempty"`
}

// Region 
type Region string
  
const (
     RegionUSCentral Region = "USCentral"
     RegionUSEast Region = "USEast"
     RegionEUWest Region = "EUWest"
     RegionSingapore Region = "Singapore"
     RegionJapan Region = "Japan"
     RegionBrazil Region = "Brazil"
     RegionAustralia Region = "Australia"
      )
// RemovePlayerTagRequest this API will trigger a player_tag_removed event and remove a tag with the given TagName and PlayFabID from the
// corresponding player profile. TagName can be used for segmentation and it is limited to 256 characters
type RemovePlayerTagRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
    // TagName unique tag for player profile.
    TagName string `json:"TagName,omitempty"`
}

// RemovePlayerTagResult 
type RemovePlayerTagResultModel struct {
}

// RemoveServerBuildRequest 
type RemoveServerBuildRequestModel struct {
    // BuildId unique identifier of the previously uploaded build executable to be removed
    BuildId string `json:"BuildId,omitempty"`
}

// RemoveServerBuildResult 
type RemoveServerBuildResultModel struct {
}

// RemoveVirtualCurrencyTypesRequest virtual currencies to be removed cannot have entries in any catalog nor store for the title. Note that this operation
// will not remove player balances for the removed currencies; if a deleted currency is recreated at any point, user
// balances will be in an undefined state.
type RemoveVirtualCurrencyTypesRequestModel struct {
    // VirtualCurrencies list of virtual currencies to delete
    VirtualCurrencies []VirtualCurrencyDataModel `json:"VirtualCurrencies,omitempty"`
}

// ResetCharacterStatisticsRequest note that this action cannot be un-done. All statistics for this character will be deleted, removing the user from all
// leaderboards for the game.
type ResetCharacterStatisticsRequestModel struct {
    // CharacterId unique PlayFab assigned ID for a specific character owned by a user
    CharacterId string `json:"CharacterId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// ResetCharacterStatisticsResult 
type ResetCharacterStatisticsResultModel struct {
}

// ResetPasswordRequest resets a player's password taking in a new password based and validating the user based off of a token sent to the
// playerto their email. The token expires after 30 minutes.
type ResetPasswordRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Password the new password for the player.
    Password string `json:"Password,omitempty"`
    // Token the token of the player requesting the password reset.
    Token string `json:"Token,omitempty"`
}

// ResetPasswordResult 
type ResetPasswordResultModel struct {
}

// ResetUserStatisticsRequest note that this action cannot be un-done. All statistics for this user will be deleted, removing the user from all
// leaderboards for the game.
type ResetUserStatisticsRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// ResetUserStatisticsResult 
type ResetUserStatisticsResultModel struct {
}

// ResolutionOutcome 
type ResolutionOutcome string
  
const (
     ResolutionOutcomeRevoke ResolutionOutcome = "Revoke"
     ResolutionOutcomeReinstate ResolutionOutcome = "Reinstate"
     ResolutionOutcomeManual ResolutionOutcome = "Manual"
      )
// ResolvePurchaseDisputeRequest 
type ResolvePurchaseDisputeRequestModel struct {
    // OrderId unique order ID for the purchase in question.
    OrderId string `json:"OrderId,omitempty"`
    // Outcome enum for the desired purchase result state after notifying the payment provider. Valid values are Revoke, Reinstate and
// Manual. Manual will cause no change to the order state.
    Outcome ResolutionOutcome `json:"Outcome,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
    // Reason the Reason parameter should correspond with the payment providers reason field, if they require one such as Facebook. In
// the case of Facebook this must match one of their refund or dispute resolution enums (See:
// https://developers.facebook.com/docs/payments/implementation-guide/handling-disputes-refunds)
    Reason string `json:"Reason,omitempty"`
}

// ResolvePurchaseDisputeResponse 
type ResolvePurchaseDisputeResponseModel struct {
    // PurchaseStatus the order's updated purchase status.
    PurchaseStatus string `json:"PurchaseStatus,omitempty"`
}

// ResultTableNode 
type ResultTableNodeModel struct {
    // ResultItem either an ItemId, or the TableId of another random result table
    ResultItem string `json:"ResultItem,omitempty"`
    // ResultItemType whether this entry in the table is an item or a link to another table
    ResultItemType ResultTableNodeType `json:"ResultItemType,omitempty"`
    // Weight how likely this is to be rolled - larger numbers add more weight
    Weight int32 `json:"Weight,omitempty"`
}

// ResultTableNodeType 
type ResultTableNodeType string
  
const (
     ResultTableNodeTypeItemId ResultTableNodeType = "ItemId"
     ResultTableNodeTypeTableId ResultTableNodeType = "TableId"
      )
// RevokeAllBansForUserRequest setting the active state of all non-expired bans for a user to Inactive. Expired bans with an Active state will be
// ignored, however. Returns information about applied updates only.
type RevokeAllBansForUserRequestModel struct {
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// RevokeAllBansForUserResult 
type RevokeAllBansForUserResultModel struct {
    // BanData information on the bans that were revoked.
    BanData []BanInfoModel `json:"BanData,omitempty"`
}

// RevokeBansRequest setting the active state of all bans requested to Inactive regardless of whether that ban has already expired. BanIds
// that do not exist will be skipped. Returns information about applied updates only.
type RevokeBansRequestModel struct {
    // BanIds ids of the bans to be revoked. Maximum 100.
    BanIds []string `json:"BanIds,omitempty"`
}

// RevokeBansResult 
type RevokeBansResultModel struct {
    // BanData information on the bans that were revoked
    BanData []BanInfoModel `json:"BanData,omitempty"`
}

// RevokeInventoryItem 
type RevokeInventoryItemModel struct {
    // CharacterId unique PlayFab assigned ID for a specific character owned by a user
    CharacterId string `json:"CharacterId,omitempty"`
    // ItemInstanceId unique PlayFab assigned instance identifier of the item
    ItemInstanceId string `json:"ItemInstanceId,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// RevokeInventoryItemRequest in cases where the inventory item in question is a "crate", and the items it contained have already been dispensed, this
// will not revoke access or otherwise remove the items which were dispensed.
type RevokeInventoryItemRequestModel struct {
    // CharacterId unique PlayFab assigned ID for a specific character owned by a user
    CharacterId string `json:"CharacterId,omitempty"`
    // ItemInstanceId unique PlayFab assigned instance identifier of the item
    ItemInstanceId string `json:"ItemInstanceId,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// RevokeInventoryItemsRequest in cases where the inventory item in question is a "crate", and the items it contained have already been dispensed, this
// will not revoke access or otherwise remove the items which were dispensed.
type RevokeInventoryItemsRequestModel struct {
    // Items array of player items to revoke, between 1 and 25 items.
    Items []RevokeInventoryItemModel `json:"Items,omitempty"`
}

// RevokeInventoryItemsResult 
type RevokeInventoryItemsResultModel struct {
    // Errors collection of any errors that occurred during processing.
    Errors []RevokeItemErrorModel `json:"Errors,omitempty"`
}

// RevokeInventoryResult 
type RevokeInventoryResultModel struct {
}

// RevokeItemError 
type RevokeItemErrorModel struct {
    // Error specific error that was encountered.
    Error GenericErrorCodes `json:"Error,omitempty"`
    // Item item information that failed to be revoked.
    Item *RevokeInventoryItemModel `json:"Item,omitempty"`
}

// RunTaskRequest the returned task instance ID can be used to query for task execution status.
type RunTaskRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Identifier provide either the task ID or the task name to run a task.
    Identifier *NameIdentifierModel `json:"Identifier,omitempty"`
}

// RunTaskResult 
type RunTaskResultModel struct {
    // TaskInstanceId iD of the task instance that is started. This can be used in Get*TaskInstance (e.g. GetCloudScriptTaskInstance) API call
// to retrieve status for the task instance.
    TaskInstanceId string `json:"TaskInstanceId,omitempty"`
}

// ScheduledTask 
type ScheduledTaskModel struct {
    // Description description the task
    Description string `json:"Description,omitempty"`
    // IsActive whether the schedule is active. Inactive schedule will not trigger task execution.
    IsActive bool `json:"IsActive"`
    // LastRunTime uTC time of last run
    LastRunTime time.Time `json:"LastRunTime,omitempty"`
    // Name name of the task. This is a unique identifier for tasks in the title.
    Name string `json:"Name,omitempty"`
    // NextRunTime uTC time of next run
    NextRunTime time.Time `json:"NextRunTime,omitempty"`
    // Parameter task parameter. Different types of task have different parameter structure. See each task type's create API
// documentation for the details.
    Parameter interface{} `json:"Parameter,omitempty"`
    // Schedule cron expression for the run schedule of the task. The expression should be in UTC.
    Schedule string `json:"Schedule,omitempty"`
    // TaskId iD of the task
    TaskId string `json:"TaskId,omitempty"`
    // Type task type.
    Type ScheduledTaskType `json:"Type,omitempty"`
}

// ScheduledTaskType 
type ScheduledTaskType string
  
const (
     ScheduledTaskTypeCloudScript ScheduledTaskType = "CloudScript"
     ScheduledTaskTypeActionsOnPlayerSegment ScheduledTaskType = "ActionsOnPlayerSegment"
     ScheduledTaskTypeCloudScriptAzureFunctions ScheduledTaskType = "CloudScriptAzureFunctions"
     ScheduledTaskTypeInsightsScheduledScaling ScheduledTaskType = "InsightsScheduledScaling"
      )
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

// SendAccountRecoveryEmailRequest if the account in question is a "temporary" account (for example, one that was created via a call to
// LoginFromIOSDeviceID), thisfunction will have no effect. Only PlayFab accounts which have valid email addresses will be
// able to receive a password reset email using this API.
type SendAccountRecoveryEmailRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Email user email address attached to their account
    Email string `json:"Email,omitempty"`
    // EmailTemplateId the email template id of the account recovery email template to send.
    EmailTemplateId string `json:"EmailTemplateId,omitempty"`
}

// SendAccountRecoveryEmailResult 
type SendAccountRecoveryEmailResultModel struct {
}

// SetPlayerSecretRequest aPIs that require signatures require that the player have a configured Player Secret Key that is used to sign all
// requests. Players that don't have a secret will be blocked from making API calls until it is configured. To create a
// signature header add a SHA256 hashed string containing UTF8 encoded JSON body as it will be sent to the server, the
// current time in UTC formatted to ISO 8601, and the players secret formatted as 'body.date.secret'. Place the resulting
// hash into the header X-PlayFab-Signature, along with a header X-PlayFab-Timestamp of the same UTC timestamp used in the
// signature.
type SetPlayerSecretRequestModel struct {
    // PlayerSecret player secret that is used to verify API request signatures (Enterprise Only).
    PlayerSecret string `json:"PlayerSecret,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// SetPlayerSecretResult 
type SetPlayerSecretResultModel struct {
}

// SetPublishedRevisionRequest 
type SetPublishedRevisionRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Revision revision to make the current published revision
    Revision int32 `json:"Revision,omitempty"`
    // Version version number
    Version int32 `json:"Version,omitempty"`
}

// SetPublishedRevisionResult 
type SetPublishedRevisionResultModel struct {
}

// SetPublisherDataRequest this API is designed to store publisher-specific values which can be read, but not written to, by the client. This data
// is shared across all titles assigned to a particular publisher, and can be used for cross-game coordination. Only titles
// assigned to a publisher can use this API. This operation is additive. If a Key does not exist in the current dataset, it
// will be added with the specified Value. If it already exists, the Value for that key will be overwritten with the new
// Value. For more information email helloplayfab@microsoft.com
type SetPublisherDataRequestModel struct {
    // Key key we want to set a value on (note, this is additive - will only replace an existing key's value if they are the same
// name.) Keys are trimmed of whitespace. Keys may not begin with the '!' character.
    Key string `json:"Key,omitempty"`
    // Value new value to set. Set to null to remove a value
    Value string `json:"Value,omitempty"`
}

// SetPublisherDataResult 
type SetPublisherDataResultModel struct {
}

// SetTitleDataAndOverridesRequest will set the given key values in the specified override or the primary title data based on whether the label is provided
// or not.
type SetTitleDataAndOverridesRequestModel struct {
    // KeyValues list of titleData key-value pairs to set/delete. Use an empty value to delete an existing key; use a non-empty value to
// create/update a key.
    KeyValues []TitleDataKeyValueModel `json:"KeyValues,omitempty"`
    // OverrideLabel name of the override.
    OverrideLabel string `json:"OverrideLabel,omitempty"`
}

// SetTitleDataAndOverridesResult 
type SetTitleDataAndOverridesResultModel struct {
}

// SetTitleDataRequest this API method is designed to store title specific values which can be read by the client. For example, a developer
// could choose to store values which modify the user experience, such as enemy spawn rates, weapon strengths, movement
// speeds, etc. This allows a developer to update the title without the need to create, test, and ship a new build. This
// operation is additive. If a Key does not exist in the current dataset, it will be added with the specified Value. If it
// already exists, the Value for that key will be overwritten with the new Value.
type SetTitleDataRequestModel struct {
    // Key key we want to set a value on (note, this is additive - will only replace an existing key's value if they are the same
// name.) Keys are trimmed of whitespace. Keys may not begin with the '!' character.
    Key string `json:"Key,omitempty"`
    // Value new value to set. Set to null to remove a value
    Value string `json:"Value,omitempty"`
}

// SetTitleDataResult 
type SetTitleDataResultModel struct {
}

// SetupPushNotificationRequest when using the Apple Push Notification service (APNS) or the development version (APNS_SANDBOX), the APNS Private Key
// should be used as the Credential in this call. With Google Cloud Messaging (GCM), the Android API Key should be used.
// The current ARN (if one exists) can be overwritten by setting the OverwriteOldARN boolean to true.
type SetupPushNotificationRequestModel struct {
    // Credential credential is the Private Key for APNS/APNS_SANDBOX, and the API Key for GCM
    Credential string `json:"Credential,omitempty"`
    // Key for APNS, this is the PlatformPrincipal (SSL Certificate)
    Key string `json:"Key,omitempty"`
    // Name name of the application sending the message (application names must be made up of only uppercase and lowercase ASCII
// letters, numbers, underscores, hyphens, and periods, and must be between 1 and 256 characters long)
    Name string `json:"Name,omitempty"`
    // OverwriteOldARN replace any existing ARN with the newly generated one. If this is set to false, an error will be returned if
// notifications have already setup for this platform.
    OverwriteOldARN bool `json:"OverwriteOldARN"`
    // Platform supported notification platforms are Apple Push Notification Service (APNS and APNS_SANDBOX) for iOS and Google Cloud
// Messaging (GCM) for Android
    Platform PushSetupPlatform `json:"Platform,omitempty"`
}

// SetupPushNotificationResult 
type SetupPushNotificationResultModel struct {
    // ARN amazon Resource Name for the created notification topic.
    ARN string `json:"ARN,omitempty"`
}

// SharedSecret 
type SharedSecretModel struct {
    // Disabled flag to indicate if this key is disabled
    Disabled bool `json:"Disabled"`
    // FriendlyName friendly name for this key
    FriendlyName string `json:"FriendlyName,omitempty"`
    // SecretKey the player shared secret to use when calling Client/GetTitlePublicKey
    SecretKey string `json:"SecretKey,omitempty"`
}

// SourceType 
type SourceType string
  
const (
     SourceTypeAdmin SourceType = "Admin"
     SourceTypeBackEnd SourceType = "BackEnd"
     SourceTypeGameClient SourceType = "GameClient"
     SourceTypeGameServer SourceType = "GameServer"
     SourceTypePartner SourceType = "Partner"
     SourceTypeCustom SourceType = "Custom"
     SourceTypeAPI SourceType = "API"
      )
// StatisticAggregationMethod 
type StatisticAggregationMethod string
  
const (
     StatisticAggregationMethodLast StatisticAggregationMethod = "Last"
     StatisticAggregationMethodMin StatisticAggregationMethod = "Min"
     StatisticAggregationMethodMax StatisticAggregationMethod = "Max"
     StatisticAggregationMethodSum StatisticAggregationMethod = "Sum"
      )
// StatisticModel 
type StatisticModelModel struct {
    // Name statistic name
    Name string `json:"Name,omitempty"`
    // Value statistic value
    Value int32 `json:"Value,omitempty"`
    // Version statistic version (0 if not a versioned statistic)
    Version int32 `json:"Version,omitempty"`
}

// StatisticResetIntervalOption 
type StatisticResetIntervalOption string
  
const (
     StatisticResetIntervalOptionNever StatisticResetIntervalOption = "Never"
     StatisticResetIntervalOptionHour StatisticResetIntervalOption = "Hour"
     StatisticResetIntervalOptionDay StatisticResetIntervalOption = "Day"
     StatisticResetIntervalOptionWeek StatisticResetIntervalOption = "Week"
     StatisticResetIntervalOptionMonth StatisticResetIntervalOption = "Month"
      )
// StatisticVersionArchivalStatus 
type StatisticVersionArchivalStatus string
  
const (
     StatisticVersionArchivalStatusNotScheduled StatisticVersionArchivalStatus = "NotScheduled"
     StatisticVersionArchivalStatusScheduled StatisticVersionArchivalStatus = "Scheduled"
     StatisticVersionArchivalStatusQueued StatisticVersionArchivalStatus = "Queued"
     StatisticVersionArchivalStatusInProgress StatisticVersionArchivalStatus = "InProgress"
     StatisticVersionArchivalStatusComplete StatisticVersionArchivalStatus = "Complete"
      )
// StatisticVersionStatus 
type StatisticVersionStatus string
  
const (
     StatisticVersionStatusActive StatisticVersionStatus = "Active"
     StatisticVersionStatusSnapshotPending StatisticVersionStatus = "SnapshotPending"
     StatisticVersionStatusSnapshot StatisticVersionStatus = "Snapshot"
     StatisticVersionStatusArchivalPending StatisticVersionStatus = "ArchivalPending"
     StatisticVersionStatusArchived StatisticVersionStatus = "Archived"
      )
// StoreItem a store entry that list a catalog item at a particular price
type StoreItemModel struct {
    // CustomData store specific custom data. The data only exists as part of this store; it is not transferred to item instances
    CustomData interface{} `json:"CustomData,omitempty"`
    // DisplayPosition intended display position for this item. Note that 0 is the first position
    DisplayPosition uint32 `json:"DisplayPosition,omitempty"`
    // ItemId unique identifier of the item as it exists in the catalog - note that this must exactly match the ItemId from the
// catalog
    ItemId string `json:"ItemId,omitempty"`
    // RealCurrencyPrices override prices for this item for specific currencies
    RealCurrencyPrices map[string]uint32 `json:"RealCurrencyPrices,omitempty"`
    // VirtualCurrencyPrices override prices for this item in virtual currencies and "RM" (the base Real Money purchase price, in USD pennies)
    VirtualCurrencyPrices map[string]uint32 `json:"VirtualCurrencyPrices,omitempty"`
}

// StoreMarketingModel marketing data about a specific store
type StoreMarketingModelModel struct {
    // Description tagline for a store.
    Description string `json:"Description,omitempty"`
    // DisplayName display name of a store as it will appear to users.
    DisplayName string `json:"DisplayName,omitempty"`
    // Metadata custom data about a store.
    Metadata interface{} `json:"Metadata,omitempty"`
}

// SubscriptionModel 
type SubscriptionModelModel struct {
    // Expiration when this subscription expires.
    Expiration time.Time `json:"Expiration,omitempty"`
    // InitialSubscriptionTime the time the subscription was orignially purchased
    InitialSubscriptionTime time.Time `json:"InitialSubscriptionTime,omitempty"`
    // IsActive whether this subscription is currently active. That is, if Expiration > now.
    IsActive bool `json:"IsActive"`
    // Status the status of this subscription, according to the subscription provider.
    Status SubscriptionProviderStatus `json:"Status,omitempty"`
    // SubscriptionId the id for this subscription
    SubscriptionId string `json:"SubscriptionId,omitempty"`
    // SubscriptionItemId the item id for this subscription from the primary catalog
    SubscriptionItemId string `json:"SubscriptionItemId,omitempty"`
    // SubscriptionProvider the provider for this subscription. Apple or Google Play are supported today.
    SubscriptionProvider string `json:"SubscriptionProvider,omitempty"`
}

// SubscriptionProviderStatus 
type SubscriptionProviderStatus string
  
const (
     SubscriptionProviderStatusNoError SubscriptionProviderStatus = "NoError"
     SubscriptionProviderStatusCancelled SubscriptionProviderStatus = "Cancelled"
     SubscriptionProviderStatusUnknownError SubscriptionProviderStatus = "UnknownError"
     SubscriptionProviderStatusBillingError SubscriptionProviderStatus = "BillingError"
     SubscriptionProviderStatusProductUnavailable SubscriptionProviderStatus = "ProductUnavailable"
     SubscriptionProviderStatusCustomerDidNotAcceptPriceChange SubscriptionProviderStatus = "CustomerDidNotAcceptPriceChange"
     SubscriptionProviderStatusFreeTrial SubscriptionProviderStatus = "FreeTrial"
     SubscriptionProviderStatusPaymentPending SubscriptionProviderStatus = "PaymentPending"
      )
// SubtractUserVirtualCurrencyRequest 
type SubtractUserVirtualCurrencyRequestModel struct {
    // Amount amount to be subtracted from the user balance of the specified virtual currency.
    Amount int32 `json:"Amount,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // PlayFabId playFab unique identifier of the user whose virtual currency balance is to be decreased.
    PlayFabId string `json:"PlayFabId,omitempty"`
    // VirtualCurrency name of the virtual currency which is to be decremented.
    VirtualCurrency string `json:"VirtualCurrency,omitempty"`
}

// TagModel 
type TagModelModel struct {
    // TagValue full value of the tag, including namespace
    TagValue string `json:"TagValue,omitempty"`
}

// TaskInstanceBasicSummary 
type TaskInstanceBasicSummaryModel struct {
    // CompletedAt uTC timestamp when the task completed.
    CompletedAt time.Time `json:"CompletedAt,omitempty"`
    // ErrorMessage error message for last processing attempt, if an error occured.
    ErrorMessage string `json:"ErrorMessage,omitempty"`
    // EstimatedSecondsRemaining estimated time remaining in seconds.
    EstimatedSecondsRemaining float64 `json:"EstimatedSecondsRemaining,omitempty"`
    // PercentComplete progress represented as percentage.
    PercentComplete float64 `json:"PercentComplete,omitempty"`
    // ScheduledByUserId if manually scheduled, ID of user who scheduled the task.
    ScheduledByUserId string `json:"ScheduledByUserId,omitempty"`
    // StartedAt uTC timestamp when the task started.
    StartedAt time.Time `json:"StartedAt,omitempty"`
    // Status current status of the task instance.
    Status TaskInstanceStatus `json:"Status,omitempty"`
    // TaskIdentifier identifier of the task this instance belongs to.
    TaskIdentifier *NameIdentifierModel `json:"TaskIdentifier,omitempty"`
    // TaskInstanceId iD of the task instance.
    TaskInstanceId string `json:"TaskInstanceId,omitempty"`
    // Type type of the task.
    Type ScheduledTaskType `json:"Type,omitempty"`
}

// TaskInstanceStatus 
type TaskInstanceStatus string
  
const (
     TaskInstanceStatusSucceeded TaskInstanceStatus = "Succeeded"
     TaskInstanceStatusStarting TaskInstanceStatus = "Starting"
     TaskInstanceStatusInProgress TaskInstanceStatus = "InProgress"
     TaskInstanceStatusFailed TaskInstanceStatus = "Failed"
     TaskInstanceStatusAborted TaskInstanceStatus = "Aborted"
     TaskInstanceStatusStalled TaskInstanceStatus = "Stalled"
      )
// TitleActivationStatus 
type TitleActivationStatus string
  
const (
     TitleActivationStatusNone TitleActivationStatus = "None"
     TitleActivationStatusActivatedTitleKey TitleActivationStatus = "ActivatedTitleKey"
     TitleActivationStatusPendingSteam TitleActivationStatus = "PendingSteam"
     TitleActivationStatusActivatedSteam TitleActivationStatus = "ActivatedSteam"
     TitleActivationStatusRevokedSteam TitleActivationStatus = "RevokedSteam"
      )
// TitleDataKeyValue 
type TitleDataKeyValueModel struct {
    // Key key we want to set a value on (note, this is additive - will only replace an existing key's value if they are the same
// name.) Keys are trimmed of whitespace. Keys may not begin with the '!' character.
    Key string `json:"Key,omitempty"`
    // Value new value to set. Set to null to remove a value
    Value string `json:"Value,omitempty"`
}

// UpdateBanRequest represents a single update ban request.
type UpdateBanRequestModel struct {
    // Active the updated active state for the ban. Null for no change.
    Active bool `json:"Active"`
    // BanId the id of the ban to be updated.
    BanId string `json:"BanId,omitempty"`
    // Expires the updated expiration date for the ban. Null for no change.
    Expires time.Time `json:"Expires,omitempty"`
    // IPAddress the updated IP address for the ban. Null for no change.
    IPAddress string `json:"IPAddress,omitempty"`
    // MACAddress the updated MAC address for the ban. Null for no change.
    MACAddress string `json:"MACAddress,omitempty"`
    // Permanent whether to make this ban permanent. Set to true to make this ban permanent. This will not modify Active state.
    Permanent bool `json:"Permanent"`
    // Reason the updated reason for the ban to be updated. Maximum 140 characters. Null for no change.
    Reason string `json:"Reason,omitempty"`
}

// UpdateBansRequest for each ban, only updates the values that are set. Leave any value to null for no change. If a ban could not be found,
// the rest are still applied. Returns information about applied updates only.
type UpdateBansRequestModel struct {
    // Bans list of bans to be updated. Maximum 100.
    Bans []UpdateBanRequestModel `json:"Bans,omitempty"`
}

// UpdateBansResult 
type UpdateBansResultModel struct {
    // BanData information on the bans that were updated
    BanData []BanInfoModel `json:"BanData,omitempty"`
}

// UpdateCatalogItemsRequest this operation is not additive. Using it will cause the indicated catalog version to be created from scratch. If there
// is an existing catalog with the version number in question, it will be deleted and replaced with only the items
// specified in this call.
type UpdateCatalogItemsRequestModel struct {
    // Catalog array of catalog items to be submitted. Note that while CatalogItem has a parameter for CatalogVersion, it is not
// required and ignored in this call.
    Catalog []CatalogItemModel `json:"Catalog,omitempty"`
    // CatalogVersion which catalog is being updated. If null, uses the default catalog.
    CatalogVersion string `json:"CatalogVersion,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // SetAsDefaultCatalog should this catalog be set as the default catalog. Defaults to true. If there is currently no default catalog, this will
// always set it.
    SetAsDefaultCatalog bool `json:"SetAsDefaultCatalog"`
}

// UpdateCatalogItemsResult 
type UpdateCatalogItemsResultModel struct {
}

// UpdateCloudScriptRequest 
type UpdateCloudScriptRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // DeveloperPlayFabId playFab user ID of the developer initiating the request.
    DeveloperPlayFabId string `json:"DeveloperPlayFabId,omitempty"`
    // Files list of Cloud Script files to upload to create the new revision. Must have at least one file.
    Files []CloudScriptFileModel `json:"Files,omitempty"`
    // Publish immediately publish the new revision
    Publish bool `json:"Publish"`
}

// UpdateCloudScriptResult 
type UpdateCloudScriptResultModel struct {
    // Revision new revision number created
    Revision int32 `json:"Revision,omitempty"`
    // Version cloud Script version updated
    Version int32 `json:"Version,omitempty"`
}

// UpdateOpenIdConnectionRequest 
type UpdateOpenIdConnectionRequestModel struct {
    // ClientId the client ID given by the ID provider.
    ClientId string `json:"ClientId,omitempty"`
    // ClientSecret the client secret given by the ID provider.
    ClientSecret string `json:"ClientSecret,omitempty"`
    // ConnectionId a name for the connection that identifies it within the title.
    ConnectionId string `json:"ConnectionId,omitempty"`
    // IssuerDiscoveryUrl the issuer URL or discovery document URL to read issuer information from
    IssuerDiscoveryUrl string `json:"IssuerDiscoveryUrl,omitempty"`
    // IssuerInformation manually specified information for an OpenID Connect issuer.
    IssuerInformation *OpenIdIssuerInformationModel `json:"IssuerInformation,omitempty"`
}

// UpdatePlayerSharedSecretRequest player Shared Secret Keys are used for the call to Client/GetTitlePublicKey, which exchanges the shared secret for an
// RSA CSP blob to be used to encrypt the payload of account creation requests when that API requires a signature header.
type UpdatePlayerSharedSecretRequestModel struct {
    // Disabled disable or Enable this key
    Disabled bool `json:"Disabled"`
    // FriendlyName friendly name for this key
    FriendlyName string `json:"FriendlyName,omitempty"`
    // SecretKey the shared secret key to update
    SecretKey string `json:"SecretKey,omitempty"`
}

// UpdatePlayerSharedSecretResult 
type UpdatePlayerSharedSecretResultModel struct {
}

// UpdatePlayerStatisticDefinitionRequest statistics are numeric values, with each statistic in the title also generating a leaderboard. The ResetInterval enables
// automatically resetting leaderboards on a specified interval. Upon reset, the statistic updates to a new version with no
// values (effectively removing all players from the leaderboard). The previous version's statistic values are also
// archived for retrieval, if needed (see GetPlayerStatisticVersions). Statistics not created via a call to
// CreatePlayerStatisticDefinition by default have a VersionChangeInterval of Never, meaning they do not reset on a
// schedule, but they can be set to do so via a call to UpdatePlayerStatisticDefinition. Once a statistic has been reset
// (sometimes referred to as versioned or incremented), the now-previous version can still be written to for up a short,
// pre-defined period (currently 10 seconds), to prevent issues with levels completing around the time of the reset. Also,
// once reset, the historical statistics for players in the title may be retrieved using the URL specified in the version
// information (GetPlayerStatisticVersions). The AggregationMethod determines what action is taken when a new statistic
// value is submitted - always update with the new value (Last), use the highest of the old and new values (Max), use the
// smallest (Min), or add them together (Sum).
type UpdatePlayerStatisticDefinitionRequestModel struct {
    // AggregationMethod the aggregation method to use in updating the statistic (defaults to last)
    AggregationMethod StatisticAggregationMethod `json:"AggregationMethod,omitempty"`
    // StatisticName unique name of the statistic
    StatisticName string `json:"StatisticName,omitempty"`
    // VersionChangeInterval interval at which the values of the statistic for all players are reset (changes are effective at the next occurance of
// the new interval boundary)
    VersionChangeInterval StatisticResetIntervalOption `json:"VersionChangeInterval,omitempty"`
}

// UpdatePlayerStatisticDefinitionResult 
type UpdatePlayerStatisticDefinitionResultModel struct {
    // Statistic updated statistic definition
    Statistic *PlayerStatisticDefinitionModel `json:"Statistic,omitempty"`
}

// UpdatePolicyRequest updates permissions for your title. Policies affect what is allowed to happen on your title. Your policy is a collection
// of statements that, together, govern particular area for your title. Today, the only allowed policy is called
// 'ApiPolicy' and it governs what api calls are allowed.
type UpdatePolicyRequestModel struct {
    // OverwritePolicy whether to overwrite or append to the existing policy.
    OverwritePolicy bool `json:"OverwritePolicy"`
    // PolicyName the name of the policy being updated. Only supported name is 'ApiPolicy'
    PolicyName string `json:"PolicyName,omitempty"`
    // Statements the new statements to include in the policy.
    Statements []PermissionStatementModel `json:"Statements,omitempty"`
}

// UpdatePolicyResponse 
type UpdatePolicyResponseModel struct {
    // PolicyName the name of the policy that was updated.
    PolicyName string `json:"PolicyName,omitempty"`
    // Statements the statements included in the new version of the policy.
    Statements []PermissionStatementModel `json:"Statements,omitempty"`
}

// UpdateRandomResultTablesRequest this operation is additive. Tables with TableId values not currently defined will be added, while those with TableId
// values matching Tables currently in the catalog will be overwritten with the given values.
type UpdateRandomResultTablesRequestModel struct {
    // CatalogVersion which catalog is being updated. If null, update the current default catalog version
    CatalogVersion string `json:"CatalogVersion,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Tables array of random result tables to make available (Note: specifying an existing TableId will result in overwriting that
// table, while any others will be added to the available set)
    Tables []RandomResultTableModel `json:"Tables,omitempty"`
}

// UpdateRandomResultTablesResult 
type UpdateRandomResultTablesResultModel struct {
}

// UpdateStoreItemsRequest this operation is not additive. Using it will cause the indicated virtual store to be created from scratch. If there is
// an existing store with the same storeId, it will be deleted and replaced with only the items specified in this call. A
// store contains an array of references to items defined inthe catalog, along with the prices for the item, in both real
// world and virtual currencies. These prices act as an override to any prices defined in the catalog. In this way, the
// base definitions of the items may be defined in the catalog, with all associated properties, while the pricing can be
// set for each store, as needed. This allows for subsets of goods to be defined for different purposes (in order to
// simplify showing some, but not all catalog items to users, based upon different characteristics), along with unique
// prices. Note that all prices defined in the catalog and store definitions for the item are considered valid, and that a
// compromised client can be made to send a request for an item based upon any of these definitions. If no price is
// specified in the store for an item, the price set in the catalog should be displayed to the user.
type UpdateStoreItemsRequestModel struct {
    // CatalogVersion catalog version of the store to update. If null, uses the default catalog.
    CatalogVersion string `json:"CatalogVersion,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // MarketingData additional data about the store
    MarketingData *StoreMarketingModelModel `json:"MarketingData,omitempty"`
    // Store array of store items - references to catalog items, with specific pricing - to be added
    Store []StoreItemModel `json:"Store,omitempty"`
    // StoreId unique identifier for the store which is to be updated
    StoreId string `json:"StoreId,omitempty"`
}

// UpdateStoreItemsResult 
type UpdateStoreItemsResultModel struct {
}

// UpdateTaskRequest note that when calling this API, all properties of the task have to be provided, including properties that you do not
// want to change. Parameters not specified would be set to default value. If the task name in the update request is new, a
// task rename operation will be executed before updating other fields of the task. WARNING: Renaming of a task may break
// logics where the task name is used as an identifier.
type UpdateTaskRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Description description the task
    Description string `json:"Description,omitempty"`
    // Identifier specify either the task ID or the name of the task to be updated.
    Identifier *NameIdentifierModel `json:"Identifier,omitempty"`
    // IsActive whether the schedule is active. Inactive schedule will not trigger task execution.
    IsActive bool `json:"IsActive"`
    // Name name of the task. This is a unique identifier for tasks in the title.
    Name string `json:"Name,omitempty"`
    // Parameter parameter object specific to the task type. See each task type's create API documentation for details.
    Parameter interface{} `json:"Parameter,omitempty"`
    // Schedule cron expression for the run schedule of the task. The expression should be in UTC.
    Schedule string `json:"Schedule,omitempty"`
    // Type task type.
    Type ScheduledTaskType `json:"Type,omitempty"`
}

// UpdateUserDataRequest this function performs an additive update of the arbitrary JSON object containing the custom data for the user. In
// updating the custom data object, keys which already exist in the object will have their values overwritten, while keys
// with null values will be removed. No other key-value pairs will be changed apart from those specified in the call.
type UpdateUserDataRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Data key-value pairs to be written to the custom data. Note that keys are trimmed of whitespace, are limited in size, and may
// not begin with a '!' character or be null.
    Data map[string]string `json:"Data,omitempty"`
    // KeysToRemove optional list of Data-keys to remove from UserData. Some SDKs cannot insert null-values into Data due to language
// constraints. Use this to delete the keys directly.
    KeysToRemove []string `json:"KeysToRemove,omitempty"`
    // Permission permission to be applied to all user data keys written in this request. Defaults to "private" if not set.
    Permission UserDataPermission `json:"Permission,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// UpdateUserDataResult 
type UpdateUserDataResultModel struct {
    // DataVersion indicates the current version of the data that has been set. This is incremented with every set call for that type of
// data (read-only, internal, etc). This version can be provided in Get calls to find updated data.
    DataVersion uint32 `json:"DataVersion,omitempty"`
}

// UpdateUserInternalDataRequest this function performs an additive update of the arbitrary JSON object containing the custom data for the user. In
// updating the custom data object, keys which already exist in the object will have their values overwritten, keys with
// null values will be removed. No other key-value pairs will be changed apart from those specified in the call.
type UpdateUserInternalDataRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Data key-value pairs to be written to the custom data. Note that keys are trimmed of whitespace, are limited in size, and may
// not begin with a '!' character or be null.
    Data map[string]string `json:"Data,omitempty"`
    // KeysToRemove optional list of Data-keys to remove from UserData. Some SDKs cannot insert null-values into Data due to language
// constraints. Use this to delete the keys directly.
    KeysToRemove []string `json:"KeysToRemove,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// UpdateUserTitleDisplayNameRequest in addition to the PlayFab username, titles can make use of a DisplayName which is also a unique identifier, but
// specific to the title. This allows for unique names which more closely match the theme or genre of a title, for example.
// This API enables changing that name, whether due to a customer request, an offensive name choice, etc.
type UpdateUserTitleDisplayNameRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // DisplayName new title display name for the user - must be between 3 and 25 characters
    DisplayName string `json:"DisplayName,omitempty"`
    // PlayFabId playFab unique identifier of the user whose title specific display name is to be changed
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// UpdateUserTitleDisplayNameResult 
type UpdateUserTitleDisplayNameResultModel struct {
    // DisplayName current title display name for the user (this will be the original display name if the rename attempt failed)
    DisplayName string `json:"DisplayName,omitempty"`
}

// UserAccountInfo 
type UserAccountInfoModel struct {
    // AndroidDeviceInfo user Android device information, if an Android device has been linked
    AndroidDeviceInfo *UserAndroidDeviceInfoModel `json:"AndroidDeviceInfo,omitempty"`
    // AppleAccountInfo sign in with Apple account information, if an Apple account has been linked
    AppleAccountInfo *UserAppleIdInfoModel `json:"AppleAccountInfo,omitempty"`
    // Created timestamp indicating when the user account was created
    Created time.Time `json:"Created,omitempty"`
    // CustomIdInfo custom ID information, if a custom ID has been assigned
    CustomIdInfo *UserCustomIdInfoModel `json:"CustomIdInfo,omitempty"`
    // FacebookInfo user Facebook information, if a Facebook account has been linked
    FacebookInfo *UserFacebookInfoModel `json:"FacebookInfo,omitempty"`
    // FacebookInstantGamesIdInfo facebook Instant Games account information, if a Facebook Instant Games account has been linked
    FacebookInstantGamesIdInfo *UserFacebookInstantGamesIdInfoModel `json:"FacebookInstantGamesIdInfo,omitempty"`
    // GameCenterInfo user Gamecenter information, if a Gamecenter account has been linked
    GameCenterInfo *UserGameCenterInfoModel `json:"GameCenterInfo,omitempty"`
    // GoogleInfo user Google account information, if a Google account has been linked
    GoogleInfo *UserGoogleInfoModel `json:"GoogleInfo,omitempty"`
    // IosDeviceInfo user iOS device information, if an iOS device has been linked
    IosDeviceInfo *UserIosDeviceInfoModel `json:"IosDeviceInfo,omitempty"`
    // KongregateInfo user Kongregate account information, if a Kongregate account has been linked
    KongregateInfo *UserKongregateInfoModel `json:"KongregateInfo,omitempty"`
    // NintendoSwitchAccountInfo nintendo Switch account information, if a Nintendo Switch account has been linked
    NintendoSwitchAccountInfo *UserNintendoSwitchAccountIdInfoModel `json:"NintendoSwitchAccountInfo,omitempty"`
    // NintendoSwitchDeviceIdInfo nintendo Switch device information, if a Nintendo Switch device has been linked
    NintendoSwitchDeviceIdInfo *UserNintendoSwitchDeviceIdInfoModel `json:"NintendoSwitchDeviceIdInfo,omitempty"`
    // OpenIdInfo openID Connect information, if any OpenID Connect accounts have been linked
    OpenIdInfo []UserOpenIdInfoModel `json:"OpenIdInfo,omitempty"`
    // PlayFabId unique identifier for the user account
    PlayFabId string `json:"PlayFabId,omitempty"`
    // PrivateInfo personal information for the user which is considered more sensitive
    PrivateInfo *UserPrivateAccountInfoModel `json:"PrivateInfo,omitempty"`
    // PsnInfo user PSN account information, if a PSN account has been linked
    PsnInfo *UserPsnInfoModel `json:"PsnInfo,omitempty"`
    // SteamInfo user Steam information, if a Steam account has been linked
    SteamInfo *UserSteamInfoModel `json:"SteamInfo,omitempty"`
    // TitleInfo title-specific information for the user account
    TitleInfo *UserTitleInfoModel `json:"TitleInfo,omitempty"`
    // TwitchInfo user Twitch account information, if a Twitch account has been linked
    TwitchInfo *UserTwitchInfoModel `json:"TwitchInfo,omitempty"`
    // Username user account name in the PlayFab service
    Username string `json:"Username,omitempty"`
    // WindowsHelloInfo windows Hello account information, if a Windows Hello account has been linked
    WindowsHelloInfo *UserWindowsHelloInfoModel `json:"WindowsHelloInfo,omitempty"`
    // XboxInfo user XBox account information, if a XBox account has been linked
    XboxInfo *UserXboxInfoModel `json:"XboxInfo,omitempty"`
}

// UserAndroidDeviceInfo 
type UserAndroidDeviceInfoModel struct {
    // AndroidDeviceId android device ID
    AndroidDeviceId string `json:"AndroidDeviceId,omitempty"`
}

// UserAppleIdInfo 
type UserAppleIdInfoModel struct {
    // AppleSubjectId apple subject ID
    AppleSubjectId string `json:"AppleSubjectId,omitempty"`
}

// UserCustomIdInfo 
type UserCustomIdInfoModel struct {
    // CustomId custom ID
    CustomId string `json:"CustomId,omitempty"`
}

// UserDataPermission indicates whether a given data key is private (readable only by the player) or public (readable by all players). When a
// player makes a GetUserData request about another player, only keys marked Public will be returned.
type UserDataPermission string
  
const (
     UserDataPermissionPrivate UserDataPermission = "Private"
     UserDataPermissionPublic UserDataPermission = "Public"
      )
// UserDataRecord 
type UserDataRecordModel struct {
    // LastUpdated timestamp for when this data was last updated.
    LastUpdated time.Time `json:"LastUpdated,omitempty"`
    // Permission indicates whether this data can be read by all users (public) or only the user (private). This is used for GetUserData
// requests being made by one player about another player.
    Permission UserDataPermission `json:"Permission,omitempty"`
    // Value data stored for the specified user data key.
    Value string `json:"Value,omitempty"`
}

// UserFacebookInfo 
type UserFacebookInfoModel struct {
    // FacebookId facebook identifier
    FacebookId string `json:"FacebookId,omitempty"`
    // FullName facebook full name
    FullName string `json:"FullName,omitempty"`
}

// UserFacebookInstantGamesIdInfo 
type UserFacebookInstantGamesIdInfoModel struct {
    // FacebookInstantGamesId facebook Instant Games ID
    FacebookInstantGamesId string `json:"FacebookInstantGamesId,omitempty"`
}

// UserGameCenterInfo 
type UserGameCenterInfoModel struct {
    // GameCenterId gamecenter identifier
    GameCenterId string `json:"GameCenterId,omitempty"`
}

// UserGoogleInfo 
type UserGoogleInfoModel struct {
    // GoogleEmail email address of the Google account
    GoogleEmail string `json:"GoogleEmail,omitempty"`
    // GoogleGender gender information of the Google account
    GoogleGender string `json:"GoogleGender,omitempty"`
    // GoogleId google ID
    GoogleId string `json:"GoogleId,omitempty"`
    // GoogleLocale locale of the Google account
    GoogleLocale string `json:"GoogleLocale,omitempty"`
    // GoogleName name of the Google account user
    GoogleName string `json:"GoogleName,omitempty"`
}

// UserIosDeviceInfo 
type UserIosDeviceInfoModel struct {
    // IosDeviceId iOS device ID
    IosDeviceId string `json:"IosDeviceId,omitempty"`
}

// UserKongregateInfo 
type UserKongregateInfoModel struct {
    // KongregateId kongregate ID
    KongregateId string `json:"KongregateId,omitempty"`
    // KongregateName kongregate Username
    KongregateName string `json:"KongregateName,omitempty"`
}

// UserNintendoSwitchAccountIdInfo 
type UserNintendoSwitchAccountIdInfoModel struct {
    // NintendoSwitchAccountSubjectId nintendo Switch account subject ID
    NintendoSwitchAccountSubjectId string `json:"NintendoSwitchAccountSubjectId,omitempty"`
}

// UserNintendoSwitchDeviceIdInfo 
type UserNintendoSwitchDeviceIdInfoModel struct {
    // NintendoSwitchDeviceId nintendo Switch Device ID
    NintendoSwitchDeviceId string `json:"NintendoSwitchDeviceId,omitempty"`
}

// UserOpenIdInfo 
type UserOpenIdInfoModel struct {
    // ConnectionId openID Connection ID
    ConnectionId string `json:"ConnectionId,omitempty"`
    // Issuer openID Issuer
    Issuer string `json:"Issuer,omitempty"`
    // Subject openID Subject
    Subject string `json:"Subject,omitempty"`
}

// UserOrigination 
type UserOrigination string
  
const (
     UserOriginationOrganic UserOrigination = "Organic"
     UserOriginationSteam UserOrigination = "Steam"
     UserOriginationGoogle UserOrigination = "Google"
     UserOriginationAmazon UserOrigination = "Amazon"
     UserOriginationFacebook UserOrigination = "Facebook"
     UserOriginationKongregate UserOrigination = "Kongregate"
     UserOriginationGamersFirst UserOrigination = "GamersFirst"
     UserOriginationUnknown UserOrigination = "Unknown"
     UserOriginationIOS UserOrigination = "IOS"
     UserOriginationLoadTest UserOrigination = "LoadTest"
     UserOriginationAndroid UserOrigination = "Android"
     UserOriginationPSN UserOrigination = "PSN"
     UserOriginationGameCenter UserOrigination = "GameCenter"
     UserOriginationCustomId UserOrigination = "CustomId"
     UserOriginationXboxLive UserOrigination = "XboxLive"
     UserOriginationParse UserOrigination = "Parse"
     UserOriginationTwitch UserOrigination = "Twitch"
     UserOriginationWindowsHello UserOrigination = "WindowsHello"
     UserOriginationServerCustomId UserOrigination = "ServerCustomId"
     UserOriginationNintendoSwitchDeviceId UserOrigination = "NintendoSwitchDeviceId"
     UserOriginationFacebookInstantGamesId UserOrigination = "FacebookInstantGamesId"
     UserOriginationOpenIdConnect UserOrigination = "OpenIdConnect"
     UserOriginationApple UserOrigination = "Apple"
     UserOriginationNintendoSwitchAccount UserOrigination = "NintendoSwitchAccount"
      )
// UserPrivateAccountInfo 
type UserPrivateAccountInfoModel struct {
    // Email user email address
    Email string `json:"Email,omitempty"`
}

// UserPsnInfo 
type UserPsnInfoModel struct {
    // PsnAccountId pSN account ID
    PsnAccountId string `json:"PsnAccountId,omitempty"`
    // PsnOnlineId pSN online ID
    PsnOnlineId string `json:"PsnOnlineId,omitempty"`
}

// UserSteamInfo 
type UserSteamInfoModel struct {
    // SteamActivationStatus what stage of game ownership the user is listed as being in, from Steam
    SteamActivationStatus TitleActivationStatus `json:"SteamActivationStatus,omitempty"`
    // SteamCountry the country in which the player resides, from Steam data
    SteamCountry string `json:"SteamCountry,omitempty"`
    // SteamCurrency currency type set in the user Steam account
    SteamCurrency Currency `json:"SteamCurrency,omitempty"`
    // SteamId steam identifier
    SteamId string `json:"SteamId,omitempty"`
    // SteamName steam display name
    SteamName string `json:"SteamName,omitempty"`
}

// UserTitleInfo 
type UserTitleInfoModel struct {
    // AvatarUrl uRL to the player's avatar.
    AvatarUrl string `json:"AvatarUrl,omitempty"`
    // Created timestamp indicating when the user was first associated with this game (this can differ significantly from when the user
// first registered with PlayFab)
    Created time.Time `json:"Created,omitempty"`
    // DisplayName name of the user, as it is displayed in-game
    DisplayName string `json:"DisplayName,omitempty"`
    // FirstLogin timestamp indicating when the user first signed into this game (this can differ from the Created timestamp, as other
// events, such as issuing a beta key to the user, can associate the title to the user)
    FirstLogin time.Time `json:"FirstLogin,omitempty"`
    // isBanned boolean indicating whether or not the user is currently banned for a title
    IsBanned bool `json:"isBanned"`
    // LastLogin timestamp for the last user login for this title
    LastLogin time.Time `json:"LastLogin,omitempty"`
    // Origination source by which the user first joined the game, if known
    Origination UserOrigination `json:"Origination,omitempty"`
    // TitlePlayerAccount title player account entity for this user
    TitlePlayerAccount *EntityKeyModel `json:"TitlePlayerAccount,omitempty"`
}

// UserTwitchInfo 
type UserTwitchInfoModel struct {
    // TwitchId twitch ID
    TwitchId string `json:"TwitchId,omitempty"`
    // TwitchUserName twitch Username
    TwitchUserName string `json:"TwitchUserName,omitempty"`
}

// UserWindowsHelloInfo 
type UserWindowsHelloInfoModel struct {
    // WindowsHelloDeviceName windows Hello Device Name
    WindowsHelloDeviceName string `json:"WindowsHelloDeviceName,omitempty"`
    // WindowsHelloPublicKeyHash windows Hello Public Key Hash
    WindowsHelloPublicKeyHash string `json:"WindowsHelloPublicKeyHash,omitempty"`
}

// UserXboxInfo 
type UserXboxInfoModel struct {
    // XboxUserId xBox user ID
    XboxUserId string `json:"XboxUserId,omitempty"`
}

// ValueToDateModel 
type ValueToDateModelModel struct {
    // Currency iSO 4217 code of the currency used in the purchases
    Currency string `json:"Currency,omitempty"`
    // TotalValue total value of the purchases in a whole number of 1/100 monetary units. For example, 999 indicates nine dollars and
// ninety-nine cents when Currency is 'USD')
    TotalValue uint32 `json:"TotalValue,omitempty"`
    // TotalValueAsDecimal total value of the purchases in a string representation of decimal monetary units. For example, '9.99' indicates nine
// dollars and ninety-nine cents when Currency is 'USD'.
    TotalValueAsDecimal string `json:"TotalValueAsDecimal,omitempty"`
}

// VirtualCurrencyData 
type VirtualCurrencyDataModel struct {
    // CurrencyCode unique two-character identifier for this currency type (e.g.: "CC")
    CurrencyCode string `json:"CurrencyCode,omitempty"`
    // DisplayName friendly name to show in the developer portal, reports, etc.
    DisplayName string `json:"DisplayName,omitempty"`
    // InitialDeposit amount to automatically grant users upon first login to the title
    InitialDeposit int32 `json:"InitialDeposit,omitempty"`
    // RechargeMax maximum amount to which the currency will recharge (cannot exceed MaxAmount, but can be less)
    RechargeMax int32 `json:"RechargeMax,omitempty"`
    // RechargeRate rate at which the currency automatically be added to over time, in units per day (24 hours)
    RechargeRate int32 `json:"RechargeRate,omitempty"`
}

// VirtualCurrencyRechargeTime 
type VirtualCurrencyRechargeTimeModel struct {
    // RechargeMax maximum value to which the regenerating currency will automatically increment. Note that it can exceed this value
// through use of the AddUserVirtualCurrency API call. However, it will not regenerate automatically until it has fallen
// below this value.
    RechargeMax int32 `json:"RechargeMax,omitempty"`
    // RechargeTime server timestamp in UTC indicating the next time the virtual currency will be incremented.
    RechargeTime time.Time `json:"RechargeTime,omitempty"`
    // SecondsToRecharge time remaining (in seconds) before the next recharge increment of the virtual currency.
    SecondsToRecharge int32 `json:"SecondsToRecharge,omitempty"`
}
