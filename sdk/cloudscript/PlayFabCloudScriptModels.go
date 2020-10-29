package cloudscript

import "time"
                    
// AdCampaignAttributionModel 
type AdCampaignAttributionModelModel struct {
    // AttributedAt uTC time stamp of attribution
    AttributedAt time.Time `json:"AttributedAt,omitempty"`
    // CampaignId attribution campaign identifier
    CampaignId string `json:"CampaignId,omitempty"`
    // Platform attribution network name
    Platform string `json:"Platform,omitempty"`
}

// CloudScriptRevisionOption 
type CloudScriptRevisionOption string
  
const (
     CloudScriptRevisionOptionLive CloudScriptRevisionOption = "Live"
     CloudScriptRevisionOptionLatest CloudScriptRevisionOption = "Latest"
     CloudScriptRevisionOptionSpecific CloudScriptRevisionOption = "Specific"
      )
// ContactEmailInfoModel 
type ContactEmailInfoModelModel struct {
    // EmailAddress the email address
    EmailAddress string `json:"EmailAddress,omitempty"`
    // Name the name of the email info data
    Name string `json:"Name,omitempty"`
    // VerificationStatus the verification status of the email
    VerificationStatus EmailVerificationStatus `json:"VerificationStatus,omitempty"`
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
// EmailVerificationStatus 
type EmailVerificationStatus string
  
const (
     EmailVerificationStatusUnverified EmailVerificationStatus = "Unverified"
     EmailVerificationStatusPending EmailVerificationStatus = "Pending"
     EmailVerificationStatusConfirmed EmailVerificationStatus = "Confirmed"
      )
// EmptyResult 
type EmptyResultModel struct {
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

// ExecuteEntityCloudScriptRequest executes CloudScript with the entity profile that is defined in the request.
type ExecuteEntityCloudScriptRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // FunctionName the name of the CloudScript function to execute
    FunctionName string `json:"FunctionName,omitempty"`
    // FunctionParameter object that is passed in to the function as the first argument
    FunctionParameter interface{} `json:"FunctionParameter,omitempty"`
    // GeneratePlayStreamEvent generate a 'entity_executed_cloudscript' PlayStream event containing the results of the function execution and other
// contextual information. This event will show up in the PlayStream debugger console for the player in Game Manager.
    GeneratePlayStreamEvent bool `json:"GeneratePlayStreamEvent"`
    // RevisionSelection option for which revision of the CloudScript to execute. 'Latest' executes the most recently created revision, 'Live'
// executes the current live, published revision, and 'Specific' executes the specified revision. The default value is
// 'Specific', if the SpecificRevision parameter is specified, otherwise it is 'Live'.
    RevisionSelection CloudScriptRevisionOption `json:"RevisionSelection,omitempty"`
    // SpecificRevision the specific revision to execute, when RevisionSelection is set to 'Specific'
    SpecificRevision int32 `json:"SpecificRevision,omitempty"`
}

// ExecuteFunctionRequest executes an Azure Function with the profile of the entity that is defined in the request.
type ExecuteFunctionRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // FunctionName the name of the CloudScript function to execute
    FunctionName string `json:"FunctionName,omitempty"`
    // FunctionParameter object that is passed in to the function as the FunctionArgument field of the FunctionExecutionContext data structure
    FunctionParameter interface{} `json:"FunctionParameter,omitempty"`
    // GeneratePlayStreamEvent generate a 'entity_executed_cloudscript_function' PlayStream event containing the results of the function execution and
// other contextual information. This event will show up in the PlayStream debugger console for the player in Game Manager.
    GeneratePlayStreamEvent bool `json:"GeneratePlayStreamEvent"`
}

// ExecuteFunctionResult 
type ExecuteFunctionResultModel struct {
    // Error error from the CloudScript Azure Function.
    Error *FunctionExecutionErrorModel `json:"Error,omitempty"`
    // ExecutionTimeMilliseconds the amount of time the function took to execute
    ExecutionTimeMilliseconds int32 `json:"ExecutionTimeMilliseconds,omitempty"`
    // FunctionName the name of the function that executed
    FunctionName string `json:"FunctionName,omitempty"`
    // FunctionResult the object returned from the function, if any
    FunctionResult interface{} `json:"FunctionResult,omitempty"`
    // FunctionResultTooLarge flag indicating if the FunctionResult was too large and was subsequently dropped from this event.
    FunctionResultTooLarge bool `json:"FunctionResultTooLarge"`
}

// FunctionExecutionError 
type FunctionExecutionErrorModel struct {
    // Error error code, such as CloudScriptAzureFunctionsExecutionTimeLimitExceeded, CloudScriptAzureFunctionsArgumentSizeExceeded,
// CloudScriptAzureFunctionsReturnSizeExceeded or CloudScriptAzureFunctionsHTTPRequestError
    Error string `json:"Error,omitempty"`
    // Message details about the error
    Message string `json:"Message,omitempty"`
    // StackTrace point during the execution of the function at which the error occurred, if any
    StackTrace string `json:"StackTrace,omitempty"`
}

// FunctionModel 
type FunctionModelModel struct {
    // FunctionAddress the address of the function.
    FunctionAddress string `json:"FunctionAddress,omitempty"`
    // FunctionName the name the function was registered under.
    FunctionName string `json:"FunctionName,omitempty"`
    // TriggerType the trigger type for the function.
    TriggerType string `json:"TriggerType,omitempty"`
}

// HttpFunctionModel 
type HttpFunctionModelModel struct {
    // FunctionName the name the function was registered under.
    FunctionName string `json:"FunctionName,omitempty"`
    // FunctionUrl the URL of the function.
    FunctionUrl string `json:"FunctionUrl,omitempty"`
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

// ListFunctionsRequest a title can have many functions, ListHttpFunctions will return a list of all the currently registered HTTP triggered
// functions for a given title.
type ListFunctionsRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
}

// ListFunctionsResult 
type ListFunctionsResultModel struct {
    // Functions the list of functions that are currently registered for the title.
    Functions []FunctionModelModel `json:"Functions,omitempty"`
}

// ListHttpFunctionsResult 
type ListHttpFunctionsResultModel struct {
    // Functions the list of HTTP triggered functions that are currently registered for the title.
    Functions []HttpFunctionModelModel `json:"Functions,omitempty"`
}

// ListQueuedFunctionsResult 
type ListQueuedFunctionsResultModel struct {
    // Functions the list of Queue triggered functions that are currently registered for the title.
    Functions []QueuedFunctionModelModel `json:"Functions,omitempty"`
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

// NameIdentifier identifier by either name or ID. Note that a name may change due to renaming, or reused after being deleted. ID is
// immutable and unique.
type NameIdentifierModel struct {
    // Id id Identifier, if present
    Id string `json:"Id,omitempty"`
    // Name name Identifier, if present
    Name string `json:"Name,omitempty"`
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

// PlayStreamEventEnvelopeModel 
type PlayStreamEventEnvelopeModelModel struct {
    // EntityId the ID of the entity the event is about.
    EntityId string `json:"EntityId,omitempty"`
    // EntityType the type of the entity the event is about.
    EntityType string `json:"EntityType,omitempty"`
    // EventData data specific to this event.
    EventData string `json:"EventData,omitempty"`
    // EventName the name of the event.
    EventName string `json:"EventName,omitempty"`
    // EventNamespace the namespace of the event.
    EventNamespace string `json:"EventNamespace,omitempty"`
    // EventSettings settings for the event.
    EventSettings string `json:"EventSettings,omitempty"`
}

// PostFunctionResultForEntityTriggeredActionRequest 
type PostFunctionResultForEntityTriggeredActionRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity to perform this action on.
    Entity* EntityKeyModel `json:"Entity,omitempty"`
    // FunctionResult the result of the function execution.
    FunctionResult* ExecuteFunctionResultModel `json:"FunctionResult,omitempty"`
}

// PostFunctionResultForFunctionExecutionRequest 
type PostFunctionResultForFunctionExecutionRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity to perform this action on.
    Entity* EntityKeyModel `json:"Entity,omitempty"`
    // FunctionResult the result of the function execution.
    FunctionResult* ExecuteFunctionResultModel `json:"FunctionResult,omitempty"`
}

// PostFunctionResultForPlayerTriggeredActionRequest 
type PostFunctionResultForPlayerTriggeredActionRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // FunctionResult the result of the function execution.
    FunctionResult* ExecuteFunctionResultModel `json:"FunctionResult,omitempty"`
    // PlayerProfile the player profile the function was invoked with.
    PlayerProfile* PlayerProfileModelModel `json:"PlayerProfile,omitempty"`
    // PlayStreamEventEnvelope the triggering PlayStream event, if any, that caused the function to be invoked.
    PlayStreamEventEnvelope *PlayStreamEventEnvelopeModelModel `json:"PlayStreamEventEnvelope,omitempty"`
}

// PostFunctionResultForScheduledTaskRequest 
type PostFunctionResultForScheduledTaskRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity to perform this action on.
    Entity* EntityKeyModel `json:"Entity,omitempty"`
    // FunctionResult the result of the function execution
    FunctionResult* ExecuteFunctionResultModel `json:"FunctionResult,omitempty"`
    // ScheduledTaskId the id of the scheduled task that invoked the function.
    ScheduledTaskId* NameIdentifierModel `json:"ScheduledTaskId,omitempty"`
}

// PushNotificationPlatform 
type PushNotificationPlatform string
  
const (
     PushNotificationPlatformApplePushNotificationService PushNotificationPlatform = "ApplePushNotificationService"
     PushNotificationPlatformGoogleCloudMessaging PushNotificationPlatform = "GoogleCloudMessaging"
      )
// PushNotificationRegistrationModel 
type PushNotificationRegistrationModelModel struct {
    // NotificationEndpointARN notification configured endpoint
    NotificationEndpointARN string `json:"NotificationEndpointARN,omitempty"`
    // Platform push notification platform
    Platform PushNotificationPlatform `json:"Platform,omitempty"`
}

// QueuedFunctionModel 
type QueuedFunctionModelModel struct {
    // ConnectionString the connection string for the Azure Storage Account that hosts the queue.
    ConnectionString string `json:"ConnectionString,omitempty"`
    // FunctionName the name the function was registered under.
    FunctionName string `json:"FunctionName,omitempty"`
    // QueueName the name of the queue that triggers the Azure Function.
    QueueName string `json:"QueueName,omitempty"`
}

// RegisterHttpFunctionRequest 
type RegisterHttpFunctionRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // FunctionName the name of the function to register
    FunctionName string `json:"FunctionName,omitempty"`
    // FunctionUrl full URL for Azure Function that implements the function.
    FunctionUrl string `json:"FunctionUrl,omitempty"`
}

// RegisterQueuedFunctionRequest a title can have many functions, RegisterQueuedFunction associates a function name with a queue name and connection
// string.
type RegisterQueuedFunctionRequestModel struct {
    // ConnectionString a connection string for the storage account that hosts the queue for the Azure Function.
    ConnectionString string `json:"ConnectionString,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // FunctionName the name of the function to register
    FunctionName string `json:"FunctionName,omitempty"`
    // QueueName the name of the queue for the Azure Function.
    QueueName string `json:"QueueName,omitempty"`
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

// StatisticModel 
type StatisticModelModel struct {
    // Name statistic name
    Name string `json:"Name,omitempty"`
    // Value statistic value
    Value int32 `json:"Value,omitempty"`
    // Version statistic version (0 if not a versioned statistic)
    Version int32 `json:"Version,omitempty"`
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
// TagModel 
type TagModelModel struct {
    // TagValue full value of the tag, including namespace
    TagValue string `json:"TagValue,omitempty"`
}

// TriggerType 
type TriggerType string
  
const (
     TriggerTypeHTTP TriggerType = "HTTP"
     TriggerTypeQueue TriggerType = "Queue"
      )
// UnregisterFunctionRequest 
type UnregisterFunctionRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // FunctionName the name of the function to unregister
    FunctionName string `json:"FunctionName,omitempty"`
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
