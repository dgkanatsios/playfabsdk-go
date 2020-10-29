package server

import "time"
                    
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

// AddCharacterVirtualCurrencyRequest 
type AddCharacterVirtualCurrencyRequestModel struct {
    // Amount amount to be added to the character balance of the specified virtual currency. Maximum VC balance is Int32
// (2,147,483,647). Any increase over this value will be discarded.
    Amount int32 `json:"Amount,omitempty"`
    // CharacterId unique PlayFab assigned ID for a specific character owned by a user
    CharacterId string `json:"CharacterId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // PlayFabId playFab unique identifier of the user whose virtual currency balance is to be incremented.
    PlayFabId string `json:"PlayFabId,omitempty"`
    // VirtualCurrency name of the virtual currency which is to be incremented.
    VirtualCurrency string `json:"VirtualCurrency,omitempty"`
}

// AddFriendRequest 
type AddFriendRequestModel struct {
    // FriendEmail email address of the user being added.
    FriendEmail string `json:"FriendEmail,omitempty"`
    // FriendPlayFabId the PlayFab identifier of the user being added.
    FriendPlayFabId string `json:"FriendPlayFabId,omitempty"`
    // FriendTitleDisplayName title-specific display name of the user to being added.
    FriendTitleDisplayName string `json:"FriendTitleDisplayName,omitempty"`
    // FriendUsername the PlayFab username of the user being added
    FriendUsername string `json:"FriendUsername,omitempty"`
    // PlayFabId playFab identifier of the player to add a new friend.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// AddGenericIDRequest 
type AddGenericIDRequestModel struct {
    // GenericId generic service identifier to add to the player account.
    GenericId* GenericServiceIdModel `json:"GenericId,omitempty"`
    // PlayFabId playFabId of the user to link.
    PlayFabId string `json:"PlayFabId,omitempty"`
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

// AddSharedGroupMembersRequest 
type AddSharedGroupMembersRequestModel struct {
    // PlayFabIds an array of unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabIds []string `json:"PlayFabIds,omitempty"`
    // SharedGroupId unique identifier for the shared group.
    SharedGroupId string `json:"SharedGroupId,omitempty"`
}

// AddSharedGroupMembersResult 
type AddSharedGroupMembersResultModel struct {
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

// AdvancedPushPlatformMsg 
type AdvancedPushPlatformMsgModel struct {
    // GCMDataOnly stops GoogleCloudMessaging notifications from including both notification and data properties and instead only sends the
// data property.
    GCMDataOnly bool `json:"GCMDataOnly"`
    // Json the Json the platform should receive.
    Json string `json:"Json,omitempty"`
    // Platform the platform that should receive the Json.
    Platform PushNotificationPlatform `json:"Platform,omitempty"`
}

// AuthenticateSessionTicketRequest note that data returned may be Personally Identifying Information (PII), such as email address, and so care should be
// taken in how this data is stored and managed. Since this call will always return the relevant information for users who
// have accessed the title, the recommendation is to not store this data locally.
type AuthenticateSessionTicketRequestModel struct {
    // SessionTicket session ticket as issued by a PlayFab client login API.
    SessionTicket string `json:"SessionTicket,omitempty"`
}

// AuthenticateSessionTicketResult 
type AuthenticateSessionTicketResultModel struct {
    // IsSessionTicketExpired indicates if token was expired at request time.
    IsSessionTicketExpired bool `json:"IsSessionTicketExpired"`
    // UserInfo account info for the user whose session ticket was supplied.
    UserInfo *UserAccountInfoModel `json:"UserInfo,omitempty"`
}

// AwardSteamAchievementItem 
type AwardSteamAchievementItemModel struct {
    // AchievementName unique Steam achievement name.
    AchievementName string `json:"AchievementName,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
    // Result result of the award attempt (only valid on response, not on request).
    Result bool `json:"Result"`
}

// AwardSteamAchievementRequest 
type AwardSteamAchievementRequestModel struct {
    // Achievements array of achievements to grant and the users to whom they are to be granted.
    Achievements []AwardSteamAchievementItemModel `json:"Achievements,omitempty"`
}

// AwardSteamAchievementResult 
type AwardSteamAchievementResultModel struct {
    // AchievementResults array of achievements granted.
    AchievementResults []AwardSteamAchievementItemModel `json:"AchievementResults,omitempty"`
}

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

// CharacterInventory 
type CharacterInventoryModel struct {
    // CharacterId the id of this character.
    CharacterId string `json:"CharacterId,omitempty"`
    // Inventory the inventory of this character.
    Inventory []ItemInstanceModel `json:"Inventory,omitempty"`
}

// CharacterLeaderboardEntry 
type CharacterLeaderboardEntryModel struct {
    // CharacterId playFab unique identifier of the character that belongs to the user for this leaderboard entry.
    CharacterId string `json:"CharacterId,omitempty"`
    // CharacterName title-specific display name of the character for this leaderboard entry.
    CharacterName string `json:"CharacterName,omitempty"`
    // CharacterType name of the character class for this entry.
    CharacterType string `json:"CharacterType,omitempty"`
    // DisplayName title-specific display name of the user for this leaderboard entry.
    DisplayName string `json:"DisplayName,omitempty"`
    // PlayFabId playFab unique identifier of the user for this leaderboard entry.
    PlayFabId string `json:"PlayFabId,omitempty"`
    // Position user's overall position in the leaderboard.
    Position int32 `json:"Position,omitempty"`
    // StatValue specific value of the user's statistic.
    StatValue int32 `json:"StatValue,omitempty"`
}

// CharacterResult 
type CharacterResultModel struct {
    // CharacterId the id for this character on this player.
    CharacterId string `json:"CharacterId,omitempty"`
    // CharacterName the name of this character.
    CharacterName string `json:"CharacterName,omitempty"`
    // CharacterType the type-string that was given to this character on creation.
    CharacterType string `json:"CharacterType,omitempty"`
}

// CloudScriptRevisionOption 
type CloudScriptRevisionOption string
  
const (
     CloudScriptRevisionOptionLive CloudScriptRevisionOption = "Live"
     CloudScriptRevisionOptionLatest CloudScriptRevisionOption = "Latest"
     CloudScriptRevisionOptionSpecific CloudScriptRevisionOption = "Specific"
      )
// ConsumeItemRequest 
type ConsumeItemRequestModel struct {
    // CharacterId unique PlayFab assigned ID for a specific character owned by a user
    CharacterId string `json:"CharacterId,omitempty"`
    // ConsumeCount number of uses to consume from the item.
    ConsumeCount int32 `json:"ConsumeCount,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // ItemInstanceId unique instance identifier of the item to be consumed.
    ItemInstanceId string `json:"ItemInstanceId,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// ConsumeItemResult 
type ConsumeItemResultModel struct {
    // ItemInstanceId unique instance identifier of the item with uses consumed.
    ItemInstanceId string `json:"ItemInstanceId,omitempty"`
    // RemainingUses number of uses remaining on the item.
    RemainingUses int32 `json:"RemainingUses,omitempty"`
}

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
// CreateSharedGroupRequest if SharedGroupId is specified, the service will attempt to create a group with that identifier, and will return an error
// if it is already in use. If no SharedGroupId is specified, a random identifier will be assigned.
type CreateSharedGroupRequestModel struct {
    // SharedGroupId unique identifier for the shared group (a random identifier will be assigned, if one is not specified).
    SharedGroupId string `json:"SharedGroupId,omitempty"`
}

// CreateSharedGroupResult 
type CreateSharedGroupResultModel struct {
    // SharedGroupId unique identifier for the shared group.
    SharedGroupId string `json:"SharedGroupId,omitempty"`
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
// DeleteCharacterFromUserRequest this function will delete the specified character from the list allowed by the user, and will also delete any inventory
// or VC currently held by that character. It will NOT delete any statistics associated for this character, in order to
// preserve leaderboard integrity.
type DeleteCharacterFromUserRequestModel struct {
    // CharacterId unique PlayFab assigned ID for a specific character owned by a user
    CharacterId string `json:"CharacterId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
    // SaveCharacterInventory if true, the character's inventory will be transferred up to the owning user; otherwise, this request will purge those
// items.
    SaveCharacterInventory bool `json:"SaveCharacterInventory"`
}

// DeleteCharacterFromUserResult 
type DeleteCharacterFromUserResultModel struct {
}

// DeletePlayerRequest deletes all data associated with the player, including statistics, custom data, inventory, purchases, virtual currency
// balances, characters and shared group memberships. Removes the player from all leaderboards and player search indexes.
// Does not delete PlayStream event history associated with the player. Does not delete the publisher user account that
// created the player in the title nor associated data such as username, password, email address, account linkages, or
// friends list. Note, this API queues the player for deletion and returns immediately. It may take several minutes or more
// before all player data is fully deleted. Until the player data is fully deleted, attempts to recreate the player with
// the same user account in the same title will fail with the 'AccountDeleted' error. This API must be enabled for use as
// an option in the game manager website. It is disabled by default.
type DeletePlayerRequestModel struct {
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// DeletePlayerResult 
type DeletePlayerResultModel struct {
}

// DeletePushNotificationTemplateRequest represents the request to delete a push notification template.
type DeletePushNotificationTemplateRequestModel struct {
    // PushNotificationTemplateId id of the push notification template to be deleted.
    PushNotificationTemplateId string `json:"PushNotificationTemplateId,omitempty"`
}

// DeletePushNotificationTemplateResult 
type DeletePushNotificationTemplateResultModel struct {
}

// DeleteSharedGroupRequest 
type DeleteSharedGroupRequestModel struct {
    // SharedGroupId unique identifier for the shared group.
    SharedGroupId string `json:"SharedGroupId,omitempty"`
}

// DeregisterGameRequest 
type DeregisterGameRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // LobbyId unique identifier for the Game Server Instance that is being deregistered.
    LobbyId string `json:"LobbyId,omitempty"`
}

// DeregisterGameResponse 
type DeregisterGameResponseModel struct {
}

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

// EntityTokenResponse 
type EntityTokenResponseModel struct {
    // Entity the entity id and type.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // EntityToken the token used to set X-EntityToken for all entity based API calls.
    EntityToken string `json:"EntityToken,omitempty"`
    // TokenExpiration the time the token will expire, if it is an expiring token, in UTC.
    TokenExpiration time.Time `json:"TokenExpiration,omitempty"`
}

// EvaluateRandomResultTableRequest 
type EvaluateRandomResultTableRequestModel struct {
    // CatalogVersion specifies the catalog version that should be used to evaluate the Random Result Table. If unspecified, uses
// default/primary catalog.
    CatalogVersion string `json:"CatalogVersion,omitempty"`
    // TableId the unique identifier of the Random Result Table to use.
    TableId string `json:"TableId,omitempty"`
}

// EvaluateRandomResultTableResult note that if the Random Result Table contains no entries, or does not exist for the catalog specified (the Primary
// catalog if one is not specified), an InvalidDropTable error will be returned.
type EvaluateRandomResultTableResultModel struct {
    // ResultItemId unique identifier for the item returned from the Random Result Table evaluation, for the given catalog.
    ResultItemId string `json:"ResultItemId,omitempty"`
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

// ExecuteCloudScriptServerRequest 
type ExecuteCloudScriptServerRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // FunctionName the name of the CloudScript function to execute
    FunctionName string `json:"FunctionName,omitempty"`
    // FunctionParameter object that is passed in to the function as the first argument
    FunctionParameter interface{} `json:"FunctionParameter,omitempty"`
    // GeneratePlayStreamEvent generate a 'player_executed_cloudscript' PlayStream event containing the results of the function execution and other
// contextual information. This event will show up in the PlayStream debugger console for the player in Game Manager.
    GeneratePlayStreamEvent bool `json:"GeneratePlayStreamEvent"`
    // PlayFabId the unique user identifier for the player on whose behalf the script is being run
    PlayFabId string `json:"PlayFabId,omitempty"`
    // RevisionSelection option for which revision of the CloudScript to execute. 'Latest' executes the most recently created revision, 'Live'
// executes the current live, published revision, and 'Specific' executes the specified revision. The default value is
// 'Specific', if the SpeificRevision parameter is specified, otherwise it is 'Live'.
    RevisionSelection CloudScriptRevisionOption `json:"RevisionSelection,omitempty"`
    // SpecificRevision the specivic revision to execute, when RevisionSelection is set to 'Specific'
    SpecificRevision int32 `json:"SpecificRevision,omitempty"`
}

// FacebookInstantGamesPlayFabIdPair 
type FacebookInstantGamesPlayFabIdPairModel struct {
    // FacebookInstantGamesId unique Facebook Instant Games identifier for a user.
    FacebookInstantGamesId string `json:"FacebookInstantGamesId,omitempty"`
    // PlayFabId unique PlayFab identifier for a user, or null if no PlayFab account is linked to the Facebook Instant Games identifier.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// FacebookPlayFabIdPair 
type FacebookPlayFabIdPairModel struct {
    // FacebookId unique Facebook identifier for a user.
    FacebookId string `json:"FacebookId,omitempty"`
    // PlayFabId unique PlayFab identifier for a user, or null if no PlayFab account is linked to the Facebook identifier.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// FriendInfo 
type FriendInfoModel struct {
    // FacebookInfo available Facebook information (if the user and PlayFab friend are also connected in Facebook).
    FacebookInfo *UserFacebookInfoModel `json:"FacebookInfo,omitempty"`
    // FriendPlayFabId playFab unique identifier for this friend.
    FriendPlayFabId string `json:"FriendPlayFabId,omitempty"`
    // GameCenterInfo available Game Center information (if the user and PlayFab friend are also connected in Game Center).
    GameCenterInfo *UserGameCenterInfoModel `json:"GameCenterInfo,omitempty"`
    // Profile the profile of the user, if requested.
    Profile *PlayerProfileModelModel `json:"Profile,omitempty"`
    // PSNInfo available PSN information, if the user and PlayFab friend are both connected to PSN.
    PSNInfo *UserPsnInfoModel `json:"PSNInfo,omitempty"`
    // SteamInfo available Steam information (if the user and PlayFab friend are also connected in Steam).
    SteamInfo *UserSteamInfoModel `json:"SteamInfo,omitempty"`
    // Tags tags which have been associated with this friend.
    Tags []string `json:"Tags,omitempty"`
    // TitleDisplayName title-specific display name for this friend.
    TitleDisplayName string `json:"TitleDisplayName,omitempty"`
    // Username playFab unique username for this friend.
    Username string `json:"Username,omitempty"`
    // XboxInfo available Xbox information, if the user and PlayFab friend are both connected to Xbox Live.
    XboxInfo *UserXboxInfoModel `json:"XboxInfo,omitempty"`
}

// GameInstanceState 
type GameInstanceState string
  
const (
     GameInstanceStateOpen GameInstanceState = "Open"
     GameInstanceStateClosed GameInstanceState = "Closed"
      )
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
// GenericPlayFabIdPair 
type GenericPlayFabIdPairModel struct {
    // GenericId unique generic service identifier for a user.
    GenericId *GenericServiceIdModel `json:"GenericId,omitempty"`
    // PlayFabId unique PlayFab identifier for a user, or null if no PlayFab account is linked to the given generic identifier.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// GenericServiceId 
type GenericServiceIdModel struct {
    // ServiceName name of the service for which the player has a unique identifier.
    ServiceName string `json:"ServiceName,omitempty"`
    // UserId unique identifier of the player in that service.
    UserId string `json:"UserId,omitempty"`
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

// GetCharacterDataRequest data is stored as JSON key-value pairs. If the Keys parameter is provided, the data object returned will only contain
// the data specific to the indicated Keys. Otherwise, the full set of custom user data will be returned.
type GetCharacterDataRequestModel struct {
    // CharacterId unique PlayFab assigned ID for a specific character owned by a user
    CharacterId string `json:"CharacterId,omitempty"`
    // IfChangedFromDataVersion the version that currently exists according to the caller. The call will return the data for all of the keys if the
// version in the system is greater than this.
    IfChangedFromDataVersion uint32 `json:"IfChangedFromDataVersion,omitempty"`
    // Keys specific keys to search for in the custom user data.
    Keys []string `json:"Keys,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// GetCharacterDataResult 
type GetCharacterDataResultModel struct {
    // CharacterId unique PlayFab assigned ID for a specific character owned by a user
    CharacterId string `json:"CharacterId,omitempty"`
    // Data user specific data for this title.
    Data map[string]UserDataRecordModel `json:"Data,omitempty"`
    // DataVersion indicates the current version of the data that has been set. This is incremented with every set call for that type of
// data (read-only, internal, etc). This version can be provided in Get calls to find updated data.
    DataVersion uint32 `json:"DataVersion,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// GetCharacterInventoryRequest all items currently in the character inventory will be returned, irrespective of how they were acquired (via purchasing,
// grants, coupons, etc.). Items that are expired, fully consumed, or are no longer valid are not considered to be in the
// user's current inventory, and so will not be not included. Also returns their virtual currency balances.
type GetCharacterInventoryRequestModel struct {
    // CatalogVersion used to limit results to only those from a specific catalog version.
    CatalogVersion string `json:"CatalogVersion,omitempty"`
    // CharacterId unique PlayFab assigned ID for a specific character owned by a user
    CharacterId string `json:"CharacterId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// GetCharacterInventoryResult 
type GetCharacterInventoryResultModel struct {
    // CharacterId unique identifier of the character for this inventory.
    CharacterId string `json:"CharacterId,omitempty"`
    // Inventory array of inventory items belonging to the character.
    Inventory []ItemInstanceModel `json:"Inventory,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
    // VirtualCurrency array of virtual currency balance(s) belonging to the character.
    VirtualCurrency map[string]int32 `json:"VirtualCurrency,omitempty"`
    // VirtualCurrencyRechargeTimes array of remaining times and timestamps for virtual currencies.
    VirtualCurrencyRechargeTimes map[string]VirtualCurrencyRechargeTimeModel `json:"VirtualCurrencyRechargeTimes,omitempty"`
}

// GetCharacterLeaderboardRequest 
type GetCharacterLeaderboardRequestModel struct {
    // CharacterType optional character type on which to filter the leaderboard entries.
    CharacterType string `json:"CharacterType,omitempty"`
    // MaxResultsCount maximum number of entries to retrieve.
    MaxResultsCount int32 `json:"MaxResultsCount,omitempty"`
    // StartPosition first entry in the leaderboard to be retrieved.
    StartPosition int32 `json:"StartPosition,omitempty"`
    // StatisticName unique identifier for the title-specific statistic for the leaderboard.
    StatisticName string `json:"StatisticName,omitempty"`
}

// GetCharacterLeaderboardResult note that the Position of the character in the results is for the overall leaderboard.
type GetCharacterLeaderboardResultModel struct {
    // Leaderboard ordered list of leaderboard entries.
    Leaderboard []CharacterLeaderboardEntryModel `json:"Leaderboard,omitempty"`
}

// GetCharacterStatisticsRequest character statistics are similar to user statistics in that they are numeric values which may only be updated by a
// server operation, in order to minimize the opportunity for unauthorized changes. In addition to being available for use
// by the title, the statistics are used for all leaderboard operations in PlayFab.
type GetCharacterStatisticsRequestModel struct {
    // CharacterId unique PlayFab assigned ID for a specific character owned by a user
    CharacterId string `json:"CharacterId,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// GetCharacterStatisticsResult 
type GetCharacterStatisticsResultModel struct {
    // CharacterId unique identifier of the character for the statistics.
    CharacterId string `json:"CharacterId,omitempty"`
    // CharacterStatistics character statistics for the requested user.
    CharacterStatistics map[string]int32 `json:"CharacterStatistics,omitempty"`
    // PlayFabId playFab unique identifier of the user whose character statistics are being returned.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// GetContentDownloadUrlRequest 
type GetContentDownloadUrlRequestModel struct {
    // HttpMethod hTTP method to fetch item - GET or HEAD. Use HEAD when only fetching metadata. Default is GET.
    HttpMethod string `json:"HttpMethod,omitempty"`
    // Key key of the content item to fetch, usually formatted as a path, e.g. images/a.png
    Key string `json:"Key,omitempty"`
    // ThruCDN true to download through CDN. CDN provides higher download bandwidth and lower latency. However, if you want the latest,
// non-cached version of the content during development, set this to false. Default is true.
    ThruCDN bool `json:"ThruCDN"`
}

// GetContentDownloadUrlResult 
type GetContentDownloadUrlResultModel struct {
    // URL uRL for downloading content via HTTP GET or HEAD method. The URL will expire in approximately one hour.
    URL string `json:"URL,omitempty"`
}

// GetFriendLeaderboardRequest 
type GetFriendLeaderboardRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // IncludeFacebookFriends indicates whether Facebook friends should be included in the response. Default is true.
    IncludeFacebookFriends bool `json:"IncludeFacebookFriends"`
    // IncludeSteamFriends indicates whether Steam service friends should be included in the response. Default is true.
    IncludeSteamFriends bool `json:"IncludeSteamFriends"`
    // MaxResultsCount maximum number of entries to retrieve.
    MaxResultsCount int32 `json:"MaxResultsCount,omitempty"`
    // PlayFabId the player whose friend leaderboard to get
    PlayFabId string `json:"PlayFabId,omitempty"`
    // ProfileConstraints if non-null, this determines which properties of the resulting player profiles to return. For API calls from the client,
// only the allowed client profile properties for the title may be requested. These allowed properties are configured in
// the Game Manager "Client Profile Options" tab in the "Settings" section.
    ProfileConstraints *PlayerProfileViewConstraintsModel `json:"ProfileConstraints,omitempty"`
    // StartPosition position in the leaderboard to start this listing (defaults to the first entry).
    StartPosition int32 `json:"StartPosition,omitempty"`
    // StatisticName statistic used to rank friends for this leaderboard.
    StatisticName string `json:"StatisticName,omitempty"`
    // Version the version of the leaderboard to get.
    Version int32 `json:"Version,omitempty"`
    // XboxToken xbox token if Xbox friends should be included. Requires Xbox be configured on PlayFab.
    XboxToken string `json:"XboxToken,omitempty"`
}

// GetFriendsListRequest 
type GetFriendsListRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // IncludeFacebookFriends indicates whether Facebook friends should be included in the response. Default is true.
    IncludeFacebookFriends bool `json:"IncludeFacebookFriends"`
    // IncludeSteamFriends indicates whether Steam service friends should be included in the response. Default is true.
    IncludeSteamFriends bool `json:"IncludeSteamFriends"`
    // PlayFabId playFab identifier of the player whose friend list to get.
    PlayFabId string `json:"PlayFabId,omitempty"`
    // ProfileConstraints if non-null, this determines which properties of the resulting player profiles to return. For API calls from the client,
// only the allowed client profile properties for the title may be requested. These allowed properties are configured in
// the Game Manager "Client Profile Options" tab in the "Settings" section.
    ProfileConstraints *PlayerProfileViewConstraintsModel `json:"ProfileConstraints,omitempty"`
    // XboxToken xbox token if Xbox friends should be included. Requires Xbox be configured on PlayFab.
    XboxToken string `json:"XboxToken,omitempty"`
}

// GetFriendsListResult if any additional services are queried for the user's friends, those friends who also have a PlayFab account registered
// for the title will be returned in the results. For Facebook, user has to have logged into the title's Facebook app
// recently, and only friends who also plays this game will be included. For Xbox Live, user has to have logged into the
// Xbox Live recently, and only friends who also play this game will be included.
type GetFriendsListResultModel struct {
    // Friends array of friends found.
    Friends []FriendInfoModel `json:"Friends,omitempty"`
}

// GetLeaderboardAroundCharacterRequest 
type GetLeaderboardAroundCharacterRequestModel struct {
    // CharacterId unique PlayFab assigned ID for a specific character owned by a user
    CharacterId string `json:"CharacterId,omitempty"`
    // CharacterType optional character type on which to filter the leaderboard entries.
    CharacterType string `json:"CharacterType,omitempty"`
    // MaxResultsCount maximum number of entries to retrieve.
    MaxResultsCount int32 `json:"MaxResultsCount,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
    // StatisticName unique identifier for the title-specific statistic for the leaderboard.
    StatisticName string `json:"StatisticName,omitempty"`
}

// GetLeaderboardAroundCharacterResult note: When calling 'GetLeaderboardAround...' APIs, the position of the character defaults to 0 when the character does
// not have the corresponding statistic.
type GetLeaderboardAroundCharacterResultModel struct {
    // Leaderboard ordered list of leaderboard entries.
    Leaderboard []CharacterLeaderboardEntryModel `json:"Leaderboard,omitempty"`
}

// GetLeaderboardAroundUserRequest 
type GetLeaderboardAroundUserRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // MaxResultsCount maximum number of entries to retrieve.
    MaxResultsCount int32 `json:"MaxResultsCount,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
    // ProfileConstraints if non-null, this determines which properties of the resulting player profiles to return. For API calls from the client,
// only the allowed client profile properties for the title may be requested. These allowed properties are configured in
// the Game Manager "Client Profile Options" tab in the "Settings" section.
    ProfileConstraints *PlayerProfileViewConstraintsModel `json:"ProfileConstraints,omitempty"`
    // StatisticName unique identifier for the title-specific statistic for the leaderboard.
    StatisticName string `json:"StatisticName,omitempty"`
    // Version the version of the leaderboard to get.
    Version int32 `json:"Version,omitempty"`
}

// GetLeaderboardAroundUserResult note: When calling 'GetLeaderboardAround...' APIs, the position of the user defaults to 0 when the user does not have
// the corresponding statistic.
type GetLeaderboardAroundUserResultModel struct {
    // Leaderboard ordered listing of users and their positions in the requested leaderboard.
    Leaderboard []PlayerLeaderboardEntryModel `json:"Leaderboard,omitempty"`
    // NextReset the time the next scheduled reset will occur. Null if the leaderboard does not reset on a schedule.
    NextReset time.Time `json:"NextReset,omitempty"`
    // Version the version of the leaderboard returned.
    Version int32 `json:"Version,omitempty"`
}

// GetLeaderboardForUsersCharactersRequest 
type GetLeaderboardForUsersCharactersRequestModel struct {
    // MaxResultsCount maximum number of entries to retrieve.
    MaxResultsCount int32 `json:"MaxResultsCount,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
    // StatisticName unique identifier for the title-specific statistic for the leaderboard.
    StatisticName string `json:"StatisticName,omitempty"`
}

// GetLeaderboardForUsersCharactersResult nOTE: The position of the character in the results is relative to the other characters for that specific user. This mean
// the values will always be between 0 and one less than the number of characters returned regardless of the size of the
// actual leaderboard.
type GetLeaderboardForUsersCharactersResultModel struct {
    // Leaderboard ordered list of leaderboard entries.
    Leaderboard []CharacterLeaderboardEntryModel `json:"Leaderboard,omitempty"`
}

// GetLeaderboardRequest 
type GetLeaderboardRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // MaxResultsCount maximum number of entries to retrieve.
    MaxResultsCount int32 `json:"MaxResultsCount,omitempty"`
    // ProfileConstraints if non-null, this determines which properties of the resulting player profiles to return. For API calls from the client,
// only the allowed client profile properties for the title may be requested. These allowed properties are configured in
// the Game Manager "Client Profile Options" tab in the "Settings" section.
    ProfileConstraints *PlayerProfileViewConstraintsModel `json:"ProfileConstraints,omitempty"`
    // StartPosition first entry in the leaderboard to be retrieved.
    StartPosition int32 `json:"StartPosition,omitempty"`
    // StatisticName unique identifier for the title-specific statistic for the leaderboard.
    StatisticName string `json:"StatisticName,omitempty"`
    // Version the version of the leaderboard to get.
    Version int32 `json:"Version,omitempty"`
}

// GetLeaderboardResult note that the Position of the user in the results is for the overall leaderboard.
type GetLeaderboardResultModel struct {
    // Leaderboard ordered listing of users and their positions in the requested leaderboard.
    Leaderboard []PlayerLeaderboardEntryModel `json:"Leaderboard,omitempty"`
    // NextReset the time the next scheduled reset will occur. Null if the leaderboard does not reset on a schedule.
    NextReset time.Time `json:"NextReset,omitempty"`
    // Version the version of the leaderboard returned.
    Version int32 `json:"Version,omitempty"`
}

// GetPlayerCombinedInfoRequest 
type GetPlayerCombinedInfoRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // InfoRequestParameters flags for which pieces of info to return for the user.
    InfoRequestParameters* GetPlayerCombinedInfoRequestParamsModel `json:"InfoRequestParameters,omitempty"`
    // PlayFabId playFabId of the user whose data will be returned
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// GetPlayerCombinedInfoRequestParams 
type GetPlayerCombinedInfoRequestParamsModel struct {
    // GetCharacterInventories whether to get character inventories. Defaults to false.
    GetCharacterInventories bool `json:"GetCharacterInventories"`
    // GetCharacterList whether to get the list of characters. Defaults to false.
    GetCharacterList bool `json:"GetCharacterList"`
    // GetPlayerProfile whether to get player profile. Defaults to false. Has no effect for a new player.
    GetPlayerProfile bool `json:"GetPlayerProfile"`
    // GetPlayerStatistics whether to get player statistics. Defaults to false.
    GetPlayerStatistics bool `json:"GetPlayerStatistics"`
    // GetTitleData whether to get title data. Defaults to false.
    GetTitleData bool `json:"GetTitleData"`
    // GetUserAccountInfo whether to get the player's account Info. Defaults to false
    GetUserAccountInfo bool `json:"GetUserAccountInfo"`
    // GetUserData whether to get the player's custom data. Defaults to false
    GetUserData bool `json:"GetUserData"`
    // GetUserInventory whether to get the player's inventory. Defaults to false
    GetUserInventory bool `json:"GetUserInventory"`
    // GetUserReadOnlyData whether to get the player's read only data. Defaults to false
    GetUserReadOnlyData bool `json:"GetUserReadOnlyData"`
    // GetUserVirtualCurrency whether to get the player's virtual currency balances. Defaults to false
    GetUserVirtualCurrency bool `json:"GetUserVirtualCurrency"`
    // PlayerStatisticNames specific statistics to retrieve. Leave null to get all keys. Has no effect if GetPlayerStatistics is false
    PlayerStatisticNames []string `json:"PlayerStatisticNames,omitempty"`
    // ProfileConstraints specifies the properties to return from the player profile. Defaults to returning the player's display name.
    ProfileConstraints *PlayerProfileViewConstraintsModel `json:"ProfileConstraints,omitempty"`
    // TitleDataKeys specific keys to search for in the custom data. Leave null to get all keys. Has no effect if GetTitleData is false
    TitleDataKeys []string `json:"TitleDataKeys,omitempty"`
    // UserDataKeys specific keys to search for in the custom data. Leave null to get all keys. Has no effect if GetUserData is false
    UserDataKeys []string `json:"UserDataKeys,omitempty"`
    // UserReadOnlyDataKeys specific keys to search for in the custom data. Leave null to get all keys. Has no effect if GetUserReadOnlyData is
// false
    UserReadOnlyDataKeys []string `json:"UserReadOnlyDataKeys,omitempty"`
}

// GetPlayerCombinedInfoResult 
type GetPlayerCombinedInfoResultModel struct {
    // InfoResultPayload results for requested info.
    InfoResultPayload *GetPlayerCombinedInfoResultPayloadModel `json:"InfoResultPayload,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// GetPlayerCombinedInfoResultPayload 
type GetPlayerCombinedInfoResultPayloadModel struct {
    // AccountInfo account information for the user. This is always retrieved.
    AccountInfo *UserAccountInfoModel `json:"AccountInfo,omitempty"`
    // CharacterInventories inventories for each character for the user.
    CharacterInventories []CharacterInventoryModel `json:"CharacterInventories,omitempty"`
    // CharacterList list of characters for the user.
    CharacterList []CharacterResultModel `json:"CharacterList,omitempty"`
    // PlayerProfile the profile of the players. This profile is not guaranteed to be up-to-date. For a new player, this profile will not
// exist.
    PlayerProfile *PlayerProfileModelModel `json:"PlayerProfile,omitempty"`
    // PlayerStatistics list of statistics for this player.
    PlayerStatistics []StatisticValueModel `json:"PlayerStatistics,omitempty"`
    // TitleData title data for this title.
    TitleData map[string]string `json:"TitleData,omitempty"`
    // UserData user specific custom data.
    UserData map[string]UserDataRecordModel `json:"UserData,omitempty"`
    // UserDataVersion the version of the UserData that was returned.
    UserDataVersion uint32 `json:"UserDataVersion,omitempty"`
    // UserInventory array of inventory items in the user's current inventory.
    UserInventory []ItemInstanceModel `json:"UserInventory,omitempty"`
    // UserReadOnlyData user specific read-only data.
    UserReadOnlyData map[string]UserDataRecordModel `json:"UserReadOnlyData,omitempty"`
    // UserReadOnlyDataVersion the version of the Read-Only UserData that was returned.
    UserReadOnlyDataVersion uint32 `json:"UserReadOnlyDataVersion,omitempty"`
    // UserVirtualCurrency dictionary of virtual currency balance(s) belonging to the user.
    UserVirtualCurrency map[string]int32 `json:"UserVirtualCurrency,omitempty"`
    // UserVirtualCurrencyRechargeTimes dictionary of remaining times and timestamps for virtual currencies.
    UserVirtualCurrencyRechargeTimes map[string]VirtualCurrencyRechargeTimeModel `json:"UserVirtualCurrencyRechargeTimes,omitempty"`
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

// GetPlayerStatisticsRequest 
type GetPlayerStatisticsRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // PlayFabId user for whom statistics are being requested
    PlayFabId string `json:"PlayFabId,omitempty"`
    // StatisticNames statistics to return
    StatisticNames []string `json:"StatisticNames,omitempty"`
    // StatisticNameVersions statistics to return, if StatisticNames is not set (only statistics which have a version matching that provided will be
// returned)
    StatisticNameVersions []StatisticNameVersionModel `json:"StatisticNameVersions,omitempty"`
}

// GetPlayerStatisticsResult in addition to being available for use by the title, the statistics are used for all leaderboard operations in PlayFab.
type GetPlayerStatisticsResultModel struct {
    // PlayFabId playFab unique identifier of the user whose statistics are being returned
    PlayFabId string `json:"PlayFabId,omitempty"`
    // Statistics user statistics for the requested user.
    Statistics []StatisticValueModel `json:"Statistics,omitempty"`
}

// GetPlayerStatisticVersionsRequest 
type GetPlayerStatisticVersionsRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // StatisticName unique name of the statistic
    StatisticName string `json:"StatisticName,omitempty"`
}

// GetPlayerStatisticVersionsResult 
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

// GetPlayFabIDsFromFacebookIDsRequest 
type GetPlayFabIDsFromFacebookIDsRequestModel struct {
    // FacebookIDs array of unique Facebook identifiers for which the title needs to get PlayFab identifiers.
    FacebookIDs []string `json:"FacebookIDs,omitempty"`
}

// GetPlayFabIDsFromFacebookIDsResult for Facebook identifiers which have not been linked to PlayFab accounts, null will be returned.
type GetPlayFabIDsFromFacebookIDsResultModel struct {
    // Data mapping of Facebook identifiers to PlayFab identifiers.
    Data []FacebookPlayFabIdPairModel `json:"Data,omitempty"`
}

// GetPlayFabIDsFromFacebookInstantGamesIdsRequest 
type GetPlayFabIDsFromFacebookInstantGamesIdsRequestModel struct {
    // FacebookInstantGamesIds array of unique Facebook Instant Games identifiers for which the title needs to get PlayFab identifiers.
    FacebookInstantGamesIds []string `json:"FacebookInstantGamesIds,omitempty"`
}

// GetPlayFabIDsFromFacebookInstantGamesIdsResult for Facebook Instant Games identifiers which have not been linked to PlayFab accounts, null will be returned.
type GetPlayFabIDsFromFacebookInstantGamesIdsResultModel struct {
    // Data mapping of Facebook Instant Games identifiers to PlayFab identifiers.
    Data []FacebookInstantGamesPlayFabIdPairModel `json:"Data,omitempty"`
}

// GetPlayFabIDsFromGenericIDsRequest 
type GetPlayFabIDsFromGenericIDsRequestModel struct {
    // GenericIDs array of unique generic service identifiers for which the title needs to get PlayFab identifiers. Currently limited to a
// maximum of 10 in a single request.
    GenericIDs []GenericServiceIdModel `json:"GenericIDs,omitempty"`
}

// GetPlayFabIDsFromGenericIDsResult for generic service identifiers which have not been linked to PlayFab accounts, null will be returned.
type GetPlayFabIDsFromGenericIDsResultModel struct {
    // Data mapping of generic service identifiers to PlayFab identifiers.
    Data []GenericPlayFabIdPairModel `json:"Data,omitempty"`
}

// GetPlayFabIDsFromNintendoSwitchDeviceIdsRequest 
type GetPlayFabIDsFromNintendoSwitchDeviceIdsRequestModel struct {
    // NintendoSwitchDeviceIds array of unique Nintendo Switch Device identifiers for which the title needs to get PlayFab identifiers.
    NintendoSwitchDeviceIds []string `json:"NintendoSwitchDeviceIds,omitempty"`
}

// GetPlayFabIDsFromNintendoSwitchDeviceIdsResult for Nintendo Switch Device identifiers which have not been linked to PlayFab accounts, null will be returned.
type GetPlayFabIDsFromNintendoSwitchDeviceIdsResultModel struct {
    // Data mapping of Nintendo Switch Device identifiers to PlayFab identifiers.
    Data []NintendoSwitchPlayFabIdPairModel `json:"Data,omitempty"`
}

// GetPlayFabIDsFromPSNAccountIDsRequest 
type GetPlayFabIDsFromPSNAccountIDsRequestModel struct {
    // IssuerId id of the PSN issuer environment. If null, defaults to 256 (production)
    IssuerId int32 `json:"IssuerId,omitempty"`
    // PSNAccountIDs array of unique PlayStation Network identifiers for which the title needs to get PlayFab identifiers.
    PSNAccountIDs []string `json:"PSNAccountIDs,omitempty"`
}

// GetPlayFabIDsFromPSNAccountIDsResult for PlayStation Network identifiers which have not been linked to PlayFab accounts, null will be returned.
type GetPlayFabIDsFromPSNAccountIDsResultModel struct {
    // Data mapping of PlayStation Network identifiers to PlayFab identifiers.
    Data []PSNAccountPlayFabIdPairModel `json:"Data,omitempty"`
}

// GetPlayFabIDsFromSteamIDsRequest 
type GetPlayFabIDsFromSteamIDsRequestModel struct {
    // SteamStringIDs array of unique Steam identifiers (Steam profile IDs) for which the title needs to get PlayFab identifiers.
    SteamStringIDs []string `json:"SteamStringIDs,omitempty"`
}

// GetPlayFabIDsFromSteamIDsResult for Steam identifiers which have not been linked to PlayFab accounts, null will be returned.
type GetPlayFabIDsFromSteamIDsResultModel struct {
    // Data mapping of Steam identifiers to PlayFab identifiers.
    Data []SteamPlayFabIdPairModel `json:"Data,omitempty"`
}

// GetPlayFabIDsFromXboxLiveIDsRequest 
type GetPlayFabIDsFromXboxLiveIDsRequestModel struct {
    // Sandbox the ID of Xbox Live sandbox.
    Sandbox string `json:"Sandbox,omitempty"`
    // XboxLiveAccountIDs array of unique Xbox Live account identifiers for which the title needs to get PlayFab identifiers.
    XboxLiveAccountIDs []string `json:"XboxLiveAccountIDs,omitempty"`
}

// GetPlayFabIDsFromXboxLiveIDsResult for XboxLive identifiers which have not been linked to PlayFab accounts, null will be returned.
type GetPlayFabIDsFromXboxLiveIDsResultModel struct {
    // Data mapping of PlayStation Network identifiers to PlayFab identifiers.
    Data []XboxLiveAccountPlayFabIdPairModel `json:"Data,omitempty"`
}

// GetPublisherDataRequest this API is designed to return publisher-specific values which can be read, but not written to, by the client. This data
// is shared across all titles assigned to a particular publisher, and can be used for cross-game coordination. Only titles
// assigned to a publisher can use this API. For more information email helloplayfab@microsoft.com. Note that there may up
// to a minute delay in between updating title data and this API call returning the newest value.
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
    // CatalogVersion specifies the catalog version that should be used to retrieve the Random Result Tables. If unspecified, uses
// default/primary catalog.
    CatalogVersion string `json:"CatalogVersion,omitempty"`
    // TableIDs the unique identifier of the Random Result Table to use.
    TableIDs []string `json:"TableIDs,omitempty"`
}

// GetRandomResultTablesResult note that if a specified Random Result Table contains no entries, or does not exist in the catalog, an InvalidDropTable
// error will be returned.
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

// GetServerCustomIDsFromPlayFabIDsRequest 
type GetServerCustomIDsFromPlayFabIDsRequestModel struct {
    // PlayFabIDs array of unique PlayFab player identifiers for which the title needs to get server custom identifiers. Cannot contain
// more than 25 identifiers.
    PlayFabIDs []string `json:"PlayFabIDs,omitempty"`
}

// GetServerCustomIDsFromPlayFabIDsResult for a PlayFab account that isn't associated with a server custom identity, ServerCustomId will be null.
type GetServerCustomIDsFromPlayFabIDsResultModel struct {
    // Data mapping of server custom player identifiers to PlayFab identifiers.
    Data []ServerCustomIDPlayFabIDPairModel `json:"Data,omitempty"`
}

// GetSharedGroupDataRequest 
type GetSharedGroupDataRequestModel struct {
    // GetMembers if true, return the list of all members of the shared group.
    GetMembers bool `json:"GetMembers"`
    // Keys specific keys to retrieve from the shared group (if not specified, all keys will be returned, while an empty array
// indicates that no keys should be returned).
    Keys []string `json:"Keys,omitempty"`
    // SharedGroupId unique identifier for the shared group.
    SharedGroupId string `json:"SharedGroupId,omitempty"`
}

// GetSharedGroupDataResult 
type GetSharedGroupDataResultModel struct {
    // Data data for the requested keys.
    Data map[string]SharedGroupDataRecordModel `json:"Data,omitempty"`
    // Members list of PlayFabId identifiers for the members of this group, if requested.
    Members []string `json:"Members,omitempty"`
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

// GetStoreItemsServerRequest a store contains an array of references to items defined in one or more catalog versions of the game, along with the
// prices for the item, in both real world and virtual currencies. These prices act as an override to any prices defined in
// the catalog. In this way, the base definitions of the items may be defined in the catalog, with all associated
// properties, while the pricing can be set for each store, as needed. This allows for subsets of goods to be defined for
// different purposes (in order to simplify showing some, but not all catalog items to users, based upon different
// characteristics), along with unique prices. Note that all prices defined in the catalog and store definitions for the
// item are considered valid, and that a compromised client can be made to send a request for an item based upon any of
// these definitions. If no price is specified in the store for an item, the price set in the catalog should be displayed
// to the user.
type GetStoreItemsServerRequestModel struct {
    // CatalogVersion catalog version to store items from. Use default catalog version if null
    CatalogVersion string `json:"CatalogVersion,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // PlayFabId optional identifier for the player to use in requesting the store information - if used, segment overrides will be
// applied
    PlayFabId string `json:"PlayFabId,omitempty"`
    // StoreId unqiue identifier for the store which is being requested
    StoreId string `json:"StoreId,omitempty"`
}

// GetTimeRequest this query retrieves the current time from one of the servers in PlayFab. Please note that due to clock drift between
// servers, there is a potential variance of up to 5 seconds.
type GetTimeRequestModel struct {
}

// GetTimeResult time is always returned as Coordinated Universal Time (UTC).
type GetTimeResultModel struct {
    // Time current server time when the request was received, in UTC
    Time time.Time `json:"Time,omitempty"`
}

// GetTitleDataRequest this API is designed to return title specific values which can be read, but not written to, by the client. For example,
// a developer could choose to store values which modify the user experience, such as enemy spawn rates, weapon strengths,
// movement speeds, etc. This allows a developer to update the title without the need to create, test, and ship a new
// build. If an override label is specified in the request, the overrides are applied automatically and returned with the
// title data. Note that there may up to a minute delay in between updating title data and this API call returning the
// newest value.
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

// GetTitleNewsRequest 
type GetTitleNewsRequestModel struct {
    // Count limits the results to the last n entries. Defaults to 10 if not set.
    Count int32 `json:"Count,omitempty"`
}

// GetTitleNewsResult 
type GetTitleNewsResultModel struct {
    // News array of localized news items.
    News []TitleNewsItemModel `json:"News,omitempty"`
}

// GetUserAccountInfoRequest this API allows for access to details regarding a user in the PlayFab service, usually for purposes of customer support.
// Note that data returned may be Personally Identifying Information (PII), such as email address, and so care should be
// taken in how this data is stored and managed. Since this call will always return the relevant information for users who
// have accessed the title, the recommendation is to not store this data locally.
type GetUserAccountInfoRequestModel struct {
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// GetUserAccountInfoResult 
type GetUserAccountInfoResultModel struct {
    // UserInfo account details for the user whose information was requested.
    UserInfo *UserAccountInfoModel `json:"UserInfo,omitempty"`
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

// GrantCharacterToUserRequest grants a character to the user of the type and name specified in the request.
type GrantCharacterToUserRequestModel struct {
    // CharacterName non-unique display name of the character being granted (1-40 characters in length).
    CharacterName string `json:"CharacterName,omitempty"`
    // CharacterType type of the character being granted; statistics can be sliced based on this value.
    CharacterType string `json:"CharacterType,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// GrantCharacterToUserResult 
type GrantCharacterToUserResultModel struct {
    // CharacterId unique identifier tagged to this character.
    CharacterId string `json:"CharacterId,omitempty"`
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

// GrantItemsToCharacterRequest this function directly adds inventory items to the character's inventories. As a result of this operations, the user
// will not be charged any transaction fee, regardless of the inventory item catalog definition. Please note that the
// processing time for inventory grants and purchases increases fractionally the more items are in the inventory, and the
// more items are in the grant/purchase operation.
type GrantItemsToCharacterRequestModel struct {
    // Annotation string detailing any additional information concerning this operation.
    Annotation string `json:"Annotation,omitempty"`
    // CatalogVersion catalog version from which items are to be granted.
    CatalogVersion string `json:"CatalogVersion,omitempty"`
    // CharacterId unique PlayFab assigned ID for a specific character owned by a user
    CharacterId string `json:"CharacterId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // ItemIds array of itemIds to grant to the user.
    ItemIds []string `json:"ItemIds,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// GrantItemsToCharacterResult 
type GrantItemsToCharacterResultModel struct {
    // ItemGrantResults array of items granted to users.
    ItemGrantResults []GrantedItemInstanceModel `json:"ItemGrantResults,omitempty"`
}

// GrantItemsToUserRequest this function directly adds inventory items to the user's inventories. As a result of this operations, the user will not
// be charged any transaction fee, regardless of the inventory item catalog definition. Please note that the processing
// time for inventory grants and purchases increases fractionally the more items are in the inventory, and the more items
// are in the grant/purchase operation.
type GrantItemsToUserRequestModel struct {
    // Annotation string detailing any additional information concerning this operation.
    Annotation string `json:"Annotation,omitempty"`
    // CatalogVersion catalog version from which items are to be granted.
    CatalogVersion string `json:"CatalogVersion,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // ItemIds array of itemIds to grant to the user.
    ItemIds []string `json:"ItemIds,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// GrantItemsToUserResult please note that the order of the items in the response may not match the order of items in the request.
type GrantItemsToUserResultModel struct {
    // ItemGrantResults array of items granted to users.
    ItemGrantResults []GrantedItemInstanceModel `json:"ItemGrantResults,omitempty"`
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

// LinkPSNAccountRequest 
type LinkPSNAccountRequestModel struct {
    // AuthCode authentication code provided by the PlayStation Network.
    AuthCode string `json:"AuthCode,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // ForceLink if another user is already linked to the account, unlink the other user and re-link.
    ForceLink bool `json:"ForceLink"`
    // IssuerId id of the PSN issuer environment. If null, defaults to 256 (production)
    IssuerId int32 `json:"IssuerId,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
    // RedirectUri redirect URI supplied to PSN when requesting an auth code
    RedirectUri string `json:"RedirectUri,omitempty"`
}

// LinkPSNAccountResult 
type LinkPSNAccountResultModel struct {
}

// LinkServerCustomIdRequest 
type LinkServerCustomIdRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // ForceLink if another user is already linked to the custom ID, unlink the other user and re-link.
    ForceLink bool `json:"ForceLink"`
    // PlayFabId unique PlayFab identifier.
    PlayFabId string `json:"PlayFabId,omitempty"`
    // ServerCustomId unique server custom identifier for this player.
    ServerCustomId string `json:"ServerCustomId,omitempty"`
}

// LinkServerCustomIdResult 
type LinkServerCustomIdResultModel struct {
}

// LinkXboxAccountRequest 
type LinkXboxAccountRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // ForceLink if another user is already linked to the account, unlink the other user and re-link.
    ForceLink bool `json:"ForceLink"`
    // PlayFabId unique PlayFab identifier for a user, or null if no PlayFab account is linked to the Xbox Live identifier.
    PlayFabId string `json:"PlayFabId,omitempty"`
    // XboxToken token provided by the Xbox Live SDK/XDK method GetTokenAndSignatureAsync("POST", "https://playfabapi.com/", "").
    XboxToken string `json:"XboxToken,omitempty"`
}

// LinkXboxAccountResult 
type LinkXboxAccountResultModel struct {
}

// ListUsersCharactersRequest returns a list of every character that currently belongs to a user.
type ListUsersCharactersRequestModel struct {
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// ListUsersCharactersResult 
type ListUsersCharactersResultModel struct {
    // Characters the requested list of characters.
    Characters []CharacterResultModel `json:"Characters,omitempty"`
}

// LocalizedPushNotificationProperties contains the localized push notification content.
type LocalizedPushNotificationPropertiesModel struct {
    // Message message of the localized push notification template.
    Message string `json:"Message,omitempty"`
    // Subject subject of the localized push notification template.
    Subject string `json:"Subject,omitempty"`
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
// LoginWithServerCustomIdRequest 
type LoginWithServerCustomIdRequestModel struct {
    // CreateAccount automatically create a PlayFab account if one is not currently linked to this ID.
    CreateAccount bool `json:"CreateAccount"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // InfoRequestParameters flags for which pieces of info to return for the user.
    InfoRequestParameters *GetPlayerCombinedInfoRequestParamsModel `json:"InfoRequestParameters,omitempty"`
    // PlayerSecret player secret that is used to verify API request signatures (Enterprise Only).
    PlayerSecret string `json:"PlayerSecret,omitempty"`
    // ServerCustomId the backend server identifier for this player.
    ServerCustomId string `json:"ServerCustomId,omitempty"`
}

// LoginWithXboxIdRequest if this is the first time a user has signed in with the Xbox ID and CreateAccount is set to true, a new PlayFab account
// will be created and linked to the Xbox Live account. In this case, no email or username will be associated with the
// PlayFab account. Otherwise, if no PlayFab account is linked to the Xbox Live account, an error indicating this will be
// returned, so that the title can guide the user through creation of a PlayFab account.
type LoginWithXboxIdRequestModel struct {
    // CreateAccount automatically create a PlayFab account if one is not currently linked to this ID.
    CreateAccount bool `json:"CreateAccount"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // InfoRequestParameters flags for which pieces of info to return for the user.
    InfoRequestParameters *GetPlayerCombinedInfoRequestParamsModel `json:"InfoRequestParameters,omitempty"`
    // Sandbox the id of Xbox Live sandbox.
    Sandbox string `json:"Sandbox,omitempty"`
    // XboxId unique Xbox identifier for a user
    XboxId string `json:"XboxId,omitempty"`
}

// LoginWithXboxRequest if this is the first time a user has signed in with the Xbox Live account and CreateAccount is set to true, a new
// PlayFab account will be created and linked to the Xbox Live account. In this case, no email or username will be
// associated with the PlayFab account. Otherwise, if no PlayFab account is linked to the Xbox Live account, an error
// indicating this will be returned, so that the title can guide the user through creation of a PlayFab account.
type LoginWithXboxRequestModel struct {
    // CreateAccount automatically create a PlayFab account if one is not currently linked to this ID.
    CreateAccount bool `json:"CreateAccount"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // InfoRequestParameters flags for which pieces of info to return for the user.
    InfoRequestParameters *GetPlayerCombinedInfoRequestParamsModel `json:"InfoRequestParameters,omitempty"`
    // XboxToken token provided by the Xbox Live SDK/XDK method GetTokenAndSignatureAsync("POST", "https://playfabapi.com/", "").
    XboxToken string `json:"XboxToken,omitempty"`
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

// ModifyCharacterVirtualCurrencyResult 
type ModifyCharacterVirtualCurrencyResultModel struct {
    // Balance balance of the virtual currency after modification.
    Balance int32 `json:"Balance,omitempty"`
    // VirtualCurrency name of the virtual currency which was modified.
    VirtualCurrency string `json:"VirtualCurrency,omitempty"`
}

// ModifyItemUsesRequest this function can both add and remove uses of an inventory item. If the number of uses drops below zero, the item will
// be removed from active inventory.
type ModifyItemUsesRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // ItemInstanceId unique instance identifier of the item to be modified.
    ItemInstanceId string `json:"ItemInstanceId,omitempty"`
    // PlayFabId playFab unique identifier of the user whose item is being modified.
    PlayFabId string `json:"PlayFabId,omitempty"`
    // UsesToAdd number of uses to add to the item. Can be negative to remove uses.
    UsesToAdd int32 `json:"UsesToAdd,omitempty"`
}

// ModifyItemUsesResult 
type ModifyItemUsesResultModel struct {
    // ItemInstanceId unique instance identifier of the item with uses consumed.
    ItemInstanceId string `json:"ItemInstanceId,omitempty"`
    // RemainingUses number of uses remaining on the item.
    RemainingUses int32 `json:"RemainingUses,omitempty"`
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

// MoveItemToCharacterFromCharacterRequest transfers an item from a character to another character that is owned by the same user. This will remove the item from
// the character's inventory (until and unless it is moved back), and will enable the other character to make use of the
// item instead.
type MoveItemToCharacterFromCharacterRequestModel struct {
    // GivingCharacterId unique identifier of the character that currently has the item.
    GivingCharacterId string `json:"GivingCharacterId,omitempty"`
    // ItemInstanceId unique PlayFab assigned instance identifier of the item
    ItemInstanceId string `json:"ItemInstanceId,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
    // ReceivingCharacterId unique identifier of the character that will be receiving the item.
    ReceivingCharacterId string `json:"ReceivingCharacterId,omitempty"`
}

// MoveItemToCharacterFromCharacterResult 
type MoveItemToCharacterFromCharacterResultModel struct {
}

// MoveItemToCharacterFromUserRequest transfers an item from a user to a character she owns. This will remove the item from the user's inventory (until and
// unless it is moved back), and will enable the character to make use of the item instead.
type MoveItemToCharacterFromUserRequestModel struct {
    // CharacterId unique PlayFab assigned ID for a specific character owned by a user
    CharacterId string `json:"CharacterId,omitempty"`
    // ItemInstanceId unique PlayFab assigned instance identifier of the item
    ItemInstanceId string `json:"ItemInstanceId,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// MoveItemToCharacterFromUserResult 
type MoveItemToCharacterFromUserResultModel struct {
}

// MoveItemToUserFromCharacterRequest transfers an item from a character to the owning user. This will remove the item from the character's inventory (until
// and unless it is moved back), and will enable the user to make use of the item instead.
type MoveItemToUserFromCharacterRequestModel struct {
    // CharacterId unique PlayFab assigned ID for a specific character owned by a user
    CharacterId string `json:"CharacterId,omitempty"`
    // ItemInstanceId unique PlayFab assigned instance identifier of the item
    ItemInstanceId string `json:"ItemInstanceId,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// MoveItemToUserFromCharacterResult 
type MoveItemToUserFromCharacterResultModel struct {
}

// NintendoSwitchPlayFabIdPair 
type NintendoSwitchPlayFabIdPairModel struct {
    // NintendoSwitchDeviceId unique Nintendo Switch Device identifier for a user.
    NintendoSwitchDeviceId string `json:"NintendoSwitchDeviceId,omitempty"`
    // PlayFabId unique PlayFab identifier for a user, or null if no PlayFab account is linked to the Nintendo Switch Device identifier.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// NotifyMatchmakerPlayerLeftRequest 
type NotifyMatchmakerPlayerLeftRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // LobbyId unique identifier of the Game Instance the user is leaving.
    LobbyId string `json:"LobbyId,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// NotifyMatchmakerPlayerLeftResult 
type NotifyMatchmakerPlayerLeftResultModel struct {
    // PlayerState state of user leaving the Game Server Instance.
    PlayerState PlayerConnectionState `json:"PlayerState,omitempty"`
}

// PlayerConnectionState 
type PlayerConnectionState string
  
const (
     PlayerConnectionStateUnassigned PlayerConnectionState = "Unassigned"
     PlayerConnectionStateConnecting PlayerConnectionState = "Connecting"
     PlayerConnectionStateParticipating PlayerConnectionState = "Participating"
     PlayerConnectionStateParticipated PlayerConnectionState = "Participated"
      )
// PlayerLeaderboardEntry 
type PlayerLeaderboardEntryModel struct {
    // DisplayName title-specific display name of the user for this leaderboard entry.
    DisplayName string `json:"DisplayName,omitempty"`
    // PlayFabId playFab unique identifier of the user for this leaderboard entry.
    PlayFabId string `json:"PlayFabId,omitempty"`
    // Position user's overall position in the leaderboard.
    Position int32 `json:"Position,omitempty"`
    // Profile the profile of the user, if requested.
    Profile *PlayerProfileModelModel `json:"Profile,omitempty"`
    // StatValue specific value of the user's statistic.
    StatValue int32 `json:"StatValue,omitempty"`
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

// PlayerStatisticVersion 
type PlayerStatisticVersionModel struct {
    // ActivationTime time when the statistic version became active
    ActivationTime time.Time `json:"ActivationTime,omitempty"`
    // DeactivationTime time when the statistic version became inactive due to statistic version incrementing
    DeactivationTime time.Time `json:"DeactivationTime,omitempty"`
    // ScheduledActivationTime time at which the statistic version was scheduled to become active, based on the configured ResetInterval
    ScheduledActivationTime time.Time `json:"ScheduledActivationTime,omitempty"`
    // ScheduledDeactivationTime time at which the statistic version was scheduled to become inactive, based on the configured ResetInterval
    ScheduledDeactivationTime time.Time `json:"ScheduledDeactivationTime,omitempty"`
    // StatisticName name of the statistic when the version became active
    StatisticName string `json:"StatisticName,omitempty"`
    // Version version of the statistic
    Version uint32 `json:"Version,omitempty"`
}

// PSNAccountPlayFabIdPair 
type PSNAccountPlayFabIdPairModel struct {
    // PlayFabId unique PlayFab identifier for a user, or null if no PlayFab account is linked to the PlayStation Network identifier.
    PlayFabId string `json:"PlayFabId,omitempty"`
    // PSNAccountId unique PlayStation Network identifier for a user.
    PSNAccountId string `json:"PSNAccountId,omitempty"`
}

// PushNotificationPackage 
type PushNotificationPackageModel struct {
    // Badge numerical badge to display on App icon (iOS only)
    Badge int32 `json:"Badge,omitempty"`
    // CustomData this must be a JSON formatted object. For use with developer-created custom Push Notification plugins
    CustomData string `json:"CustomData,omitempty"`
    // Icon icon file to display with the message (Not supported for iOS)
    Icon string `json:"Icon,omitempty"`
    // Message content of the message (all platforms)
    Message string `json:"Message,omitempty"`
    // Sound sound file to play with the message (all platforms)
    Sound string `json:"Sound,omitempty"`
    // Title title/Subject of the message. Not supported for iOS
    Title string `json:"Title,omitempty"`
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

// RandomResultTableListing 
type RandomResultTableListingModel struct {
    // CatalogVersion catalog version this table is associated with
    CatalogVersion string `json:"CatalogVersion,omitempty"`
    // Nodes child nodes that indicate what kind of drop table item this actually is.
    Nodes []ResultTableNodeModel `json:"Nodes,omitempty"`
    // TableId unique name for this drop table
    TableId string `json:"TableId,omitempty"`
}

// RedeemCouponRequest coupon codes can be created for any item, or set of items, in the catalog for the title. This operation causes the
// coupon to be consumed, and the specific items to be awarded to the user. Attempting to re-use an already consumed code,
// or a code which has not yet been created in the service, will result in an error.
type RedeemCouponRequestModel struct {
    // CatalogVersion catalog version of the coupon.
    CatalogVersion string `json:"CatalogVersion,omitempty"`
    // CharacterId optional identifier for the Character that should receive the item. If null, item is added to the player
    CharacterId string `json:"CharacterId,omitempty"`
    // CouponCode generated coupon code to redeem.
    CouponCode string `json:"CouponCode,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// RedeemCouponResult 
type RedeemCouponResultModel struct {
    // GrantedItems items granted to the player as a result of redeeming the coupon.
    GrantedItems []ItemInstanceModel `json:"GrantedItems,omitempty"`
}

// RedeemMatchmakerTicketRequest this function is used by a Game Server Instance to validate with the PlayFab service that a user has been registered as
// connected to the server. The Ticket is provided to the client either as a result of a call to StartGame or Matchmake,
// each of which return a Ticket specific to the Game Server Instance. This function will fail in any case where the Ticket
// presented is not valid for the specific Game Server Instance making the call. Note that data returned may be Personally
// Identifying Information (PII), such as email address, and so care should be taken in how this data is stored and
// managed. Since this call will always return the relevant information for users who have accessed the title, the
// recommendation is to not store this data locally.
type RedeemMatchmakerTicketRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // LobbyId unique identifier of the Game Server Instance that is asking for validation of the authorization ticket.
    LobbyId string `json:"LobbyId,omitempty"`
    // Ticket server authorization ticket passed back from a call to Matchmake or StartGame.
    Ticket string `json:"Ticket,omitempty"`
}

// RedeemMatchmakerTicketResult 
type RedeemMatchmakerTicketResultModel struct {
    // Error error value if the ticket was not validated.
    Error string `json:"Error,omitempty"`
    // TicketIsValid boolean indicating whether the ticket was validated by the PlayFab service.
    TicketIsValid bool `json:"TicketIsValid"`
    // UserInfo user account information for the user validated.
    UserInfo *UserAccountInfoModel `json:"UserInfo,omitempty"`
}

// RefreshGameServerInstanceHeartbeatRequest 
type RefreshGameServerInstanceHeartbeatRequestModel struct {
    // LobbyId unique identifier of the Game Server Instance for which the heartbeat is updated.
    LobbyId string `json:"LobbyId,omitempty"`
}

// RefreshGameServerInstanceHeartbeatResult 
type RefreshGameServerInstanceHeartbeatResultModel struct {
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
// RegisterGameRequest 
type RegisterGameRequestModel struct {
    // Build unique identifier of the build running on the Game Server Instance.
    Build string `json:"Build,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // GameMode game Mode the Game Server instance is running. Note that this must be defined in the Game Modes tab in the PlayFab Game
// Manager, along with the Build ID (the same Game Mode can be defined for multiple Build IDs).
    GameMode string `json:"GameMode,omitempty"`
    // LobbyId previous lobby id if re-registering an existing game.
    LobbyId string `json:"LobbyId,omitempty"`
    // Region region in which the Game Server Instance is running. For matchmaking using non-AWS region names, set this to any AWS
// region and use Tags (below) to specify your custom region.
    Region Region `json:"Region,omitempty"`
    // ServerIPV4Address iPV4 address of the game server instance.
    ServerIPV4Address string `json:"ServerIPV4Address,omitempty"`
    // ServerIPV6Address iPV6 address (if any) of the game server instance.
    ServerIPV6Address string `json:"ServerIPV6Address,omitempty"`
    // ServerPort port number for communication with the Game Server Instance.
    ServerPort string `json:"ServerPort,omitempty"`
    // ServerPublicDNSName public DNS name (if any) of the server
    ServerPublicDNSName string `json:"ServerPublicDNSName,omitempty"`
    // Tags tags for the Game Server Instance
    Tags map[string]string `json:"Tags,omitempty"`
}

// RegisterGameResponse 
type RegisterGameResponseModel struct {
    // LobbyId unique identifier generated for the Game Server Instance that is registered. If LobbyId is specified in request and the
// game still exists in PlayFab, the LobbyId in request is returned. Otherwise a new lobby id will be returned.
    LobbyId string `json:"LobbyId,omitempty"`
}

// RemoveFriendRequest 
type RemoveFriendRequestModel struct {
    // FriendPlayFabId playFab identifier of the friend account which is to be removed.
    FriendPlayFabId string `json:"FriendPlayFabId,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// RemoveGenericIDRequest 
type RemoveGenericIDRequestModel struct {
    // GenericId generic service identifier to be removed from the player.
    GenericId* GenericServiceIdModel `json:"GenericId,omitempty"`
    // PlayFabId playFabId of the user to remove.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

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

// RemoveSharedGroupMembersRequest 
type RemoveSharedGroupMembersRequestModel struct {
    // PlayFabIds an array of unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabIds []string `json:"PlayFabIds,omitempty"`
    // SharedGroupId unique identifier for the shared group.
    SharedGroupId string `json:"SharedGroupId,omitempty"`
}

// RemoveSharedGroupMembersResult 
type RemoveSharedGroupMembersResultModel struct {
}

// ReportPlayerServerRequest 
type ReportPlayerServerRequestModel struct {
    // Comment optional additional comment by reporting player.
    Comment string `json:"Comment,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // ReporteeId unique PlayFab identifier of the reported player.
    ReporteeId string `json:"ReporteeId,omitempty"`
    // ReporterId playFabId of the reporting player.
    ReporterId string `json:"ReporterId,omitempty"`
}

// ReportPlayerServerResult players are currently limited to five reports per day. Attempts by a single user account to submit reports beyond five
// will result in Updated being returned as false.
type ReportPlayerServerResultModel struct {
    // SubmissionsRemaining the number of remaining reports which may be filed today by this reporting player.
    SubmissionsRemaining int32 `json:"SubmissionsRemaining,omitempty"`
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

// SavePushNotificationTemplateRequest represents the save push notification template request.
type SavePushNotificationTemplateRequestModel struct {
    // AndroidPayload android JSON for the notification template.
    AndroidPayload string `json:"AndroidPayload,omitempty"`
    // Id id of the push notification template.
    Id string `json:"Id,omitempty"`
    // IOSPayload iOS JSON for the notification template.
    IOSPayload string `json:"IOSPayload,omitempty"`
    // LocalizedPushNotificationTemplates dictionary of localized push notification templates with the language as the key.
    LocalizedPushNotificationTemplates map[string]LocalizedPushNotificationPropertiesModel `json:"LocalizedPushNotificationTemplates,omitempty"`
    // Name name of the push notification template.
    Name string `json:"Name,omitempty"`
}

// SavePushNotificationTemplateResult represents the save push notification template result.
type SavePushNotificationTemplateResultModel struct {
    // PushNotificationTemplateId id of the push notification template that was saved.
    PushNotificationTemplateId string `json:"PushNotificationTemplateId,omitempty"`
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

// SendCustomAccountRecoveryEmailRequest playFab accounts which have valid email address or username will be able to receive a password reset email using this
// API.The email sent must be an account recovery email template. The username or email can be passed in to send the email
type SendCustomAccountRecoveryEmailRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Email user email address attached to their account
    Email string `json:"Email,omitempty"`
    // EmailTemplateId the email template id of the account recovery email template to send.
    EmailTemplateId string `json:"EmailTemplateId,omitempty"`
    // Username the user's username requesting an account recovery.
    Username string `json:"Username,omitempty"`
}

// SendCustomAccountRecoveryEmailResult 
type SendCustomAccountRecoveryEmailResultModel struct {
}

// SendEmailFromTemplateRequest sends an email for only players that have contact emails associated with them. Takes in an email template ID
// specifyingthe email template to send.
type SendEmailFromTemplateRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // EmailTemplateId the email template id of the email template to send.
    EmailTemplateId string `json:"EmailTemplateId,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// SendEmailFromTemplateResult 
type SendEmailFromTemplateResultModel struct {
}

// SendPushNotificationFromTemplateRequest represents the request for sending a push notification template to a recipient.
type SendPushNotificationFromTemplateRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // PushNotificationTemplateId id of the push notification template.
    PushNotificationTemplateId string `json:"PushNotificationTemplateId,omitempty"`
    // Recipient playFabId of the push notification recipient.
    Recipient string `json:"Recipient,omitempty"`
}

// SendPushNotificationRequest 
type SendPushNotificationRequestModel struct {
    // AdvancedPlatformDelivery allows you to provide precisely formatted json to target devices. This is an advanced feature, allowing you to deliver
// to custom plugin logic, fields, or functionality not natively supported by PlayFab.
    AdvancedPlatformDelivery []AdvancedPushPlatformMsgModel `json:"AdvancedPlatformDelivery,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Message text of message to send.
    Message string `json:"Message,omitempty"`
    // Package defines all possible push attributes like message, title, icon, etc. Some parameters are device specific - please see
// the PushNotificationPackage documentation for details.
    Package *PushNotificationPackageModel `json:"Package,omitempty"`
    // Recipient playFabId of the recipient of the push notification.
    Recipient string `json:"Recipient,omitempty"`
    // Subject subject of message to send (may not be displayed in all platforms)
    Subject string `json:"Subject,omitempty"`
    // TargetPlatforms target Platforms that should receive the Message or Package. If omitted, we will send to all available platforms.
    TargetPlatforms []PushNotificationPlatform `json:"TargetPlatforms,omitempty"`
}

// SendPushNotificationResult 
type SendPushNotificationResultModel struct {
}

// ServerCustomIDPlayFabIDPair 
type ServerCustomIDPlayFabIDPairModel struct {
    // PlayFabId unique PlayFab identifier.
    PlayFabId string `json:"PlayFabId,omitempty"`
    // ServerCustomId unique server custom identifier for this player.
    ServerCustomId string `json:"ServerCustomId,omitempty"`
}

// ServerLoginResult 
type ServerLoginResultModel struct {
    // EntityToken if LoginTitlePlayerAccountEntity flag is set on the login request the title_player_account will also be logged in and
// returned.
    EntityToken *EntityTokenResponseModel `json:"EntityToken,omitempty"`
    // InfoResultPayload results for requested info.
    InfoResultPayload *GetPlayerCombinedInfoResultPayloadModel `json:"InfoResultPayload,omitempty"`
    // LastLoginTime the time of this user's previous login. If there was no previous login, then it's DateTime.MinValue
    LastLoginTime time.Time `json:"LastLoginTime,omitempty"`
    // NewlyCreated true if the account was newly created on this login.
    NewlyCreated bool `json:"NewlyCreated"`
    // PlayFabId player's unique PlayFabId.
    PlayFabId string `json:"PlayFabId,omitempty"`
    // SessionTicket unique token authorizing the user and game at the server level, for the current session.
    SessionTicket string `json:"SessionTicket,omitempty"`
    // SettingsForUser settings specific to this user.
    SettingsForUser *UserSettingsModel `json:"SettingsForUser,omitempty"`
    // TreatmentAssignment the experimentation treatments for this user at the time of login.
    TreatmentAssignment *TreatmentAssignmentModel `json:"TreatmentAssignment,omitempty"`
}

// SetFriendTagsRequest this operation is not additive. It will completely replace the tag list for the specified user. Please note that only
// users in the PlayFab friends list can be assigned tags. Attempting to set a tag on a friend only included in the friends
// list from a social site integration (such as Facebook or Steam) will return the AccountNotFound error.
type SetFriendTagsRequestModel struct {
    // FriendPlayFabId playFab identifier of the friend account to which the tag(s) should be applied.
    FriendPlayFabId string `json:"FriendPlayFabId,omitempty"`
    // PlayFabId playFab identifier of the player whose friend is to be updated.
    PlayFabId string `json:"PlayFabId,omitempty"`
    // Tags array of tags to set on the friend account.
    Tags []string `json:"Tags,omitempty"`
}

// SetGameServerInstanceDataRequest 
type SetGameServerInstanceDataRequestModel struct {
    // GameServerData custom data to set for the specified game server instance.
    GameServerData string `json:"GameServerData,omitempty"`
    // LobbyId unique identifier of the Game Instance to be updated, in decimal format.
    LobbyId string `json:"LobbyId,omitempty"`
}

// SetGameServerInstanceDataResult 
type SetGameServerInstanceDataResultModel struct {
}

// SetGameServerInstanceStateRequest 
type SetGameServerInstanceStateRequestModel struct {
    // LobbyId unique identifier of the Game Instance to be updated, in decimal format.
    LobbyId string `json:"LobbyId,omitempty"`
    // State state to set for the specified game server instance.
    State GameInstanceState `json:"State,omitempty"`
}

// SetGameServerInstanceStateResult 
type SetGameServerInstanceStateResultModel struct {
}

// SetGameServerInstanceTagsRequest 
type SetGameServerInstanceTagsRequestModel struct {
    // LobbyId unique identifier of the Game Server Instance to be updated.
    LobbyId string `json:"LobbyId,omitempty"`
    // Tags tags to set for the specified Game Server Instance. Note that this is the complete list of tags to be associated with
// the Game Server Instance.
    Tags map[string]string `json:"Tags,omitempty"`
}

// SetGameServerInstanceTagsResult 
type SetGameServerInstanceTagsResultModel struct {
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

// SetTitleDataRequest this API is designed to store title specific values which can be read, but not written to, by the client. For example, a
// developer could choose to store values which modify the user experience, such as enemy spawn rates, weapon strengths,
// movement speeds, etc. This allows a developer to update the title without the need to create, test, and ship a new
// build. This operation is additive. If a Key does not exist in the current dataset, it will be added with the specified
// Value. If it already exists, the Value for that key will be overwritten with the new Value.
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

// SharedGroupDataRecord 
type SharedGroupDataRecordModel struct {
    // LastUpdated timestamp for when this data was last updated.
    LastUpdated time.Time `json:"LastUpdated,omitempty"`
    // LastUpdatedBy playFabId of the user to last update this value.
    LastUpdatedBy string `json:"LastUpdatedBy,omitempty"`
    // Permission indicates whether this data can be read by all users (public) or only members of the group (private).
    Permission UserDataPermission `json:"Permission,omitempty"`
    // Value data stored for the specified group data key.
    Value string `json:"Value,omitempty"`
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
// StatisticModel 
type StatisticModelModel struct {
    // Name statistic name
    Name string `json:"Name,omitempty"`
    // Value statistic value
    Value int32 `json:"Value,omitempty"`
    // Version statistic version (0 if not a versioned statistic)
    Version int32 `json:"Version,omitempty"`
}

// StatisticNameVersion 
type StatisticNameVersionModel struct {
    // StatisticName unique name of the statistic
    StatisticName string `json:"StatisticName,omitempty"`
    // Version the version of the statistic to be returned
    Version uint32 `json:"Version,omitempty"`
}

// StatisticUpdate 
type StatisticUpdateModel struct {
    // StatisticName unique name of the statistic
    StatisticName string `json:"StatisticName,omitempty"`
    // Value statistic value for the player
    Value int32 `json:"Value,omitempty"`
    // Version for updates to an existing statistic value for a player, the version of the statistic when it was loaded. Null when
// setting the statistic value for the first time.
    Version uint32 `json:"Version,omitempty"`
}

// StatisticValue 
type StatisticValueModel struct {
    // StatisticName unique name of the statistic
    StatisticName string `json:"StatisticName,omitempty"`
    // Value statistic value for the player
    Value int32 `json:"Value,omitempty"`
    // Version for updates to an existing statistic value for a player, the version of the statistic when it was loaded
    Version uint32 `json:"Version,omitempty"`
}

// SteamPlayFabIdPair 
type SteamPlayFabIdPairModel struct {
    // PlayFabId unique PlayFab identifier for a user, or null if no PlayFab account is linked to the Steam identifier.
    PlayFabId string `json:"PlayFabId,omitempty"`
    // SteamStringId unique Steam identifier for a user.
    SteamStringId string `json:"SteamStringId,omitempty"`
}

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
// SubtractCharacterVirtualCurrencyRequest 
type SubtractCharacterVirtualCurrencyRequestModel struct {
    // Amount amount to be subtracted from the user balance of the specified virtual currency.
    Amount int32 `json:"Amount,omitempty"`
    // CharacterId unique PlayFab assigned ID for a specific character owned by a user
    CharacterId string `json:"CharacterId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
    // VirtualCurrency name of the virtual currency which is to be decremented.
    VirtualCurrency string `json:"VirtualCurrency,omitempty"`
}

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

// TitleActivationStatus 
type TitleActivationStatus string
  
const (
     TitleActivationStatusNone TitleActivationStatus = "None"
     TitleActivationStatusActivatedTitleKey TitleActivationStatus = "ActivatedTitleKey"
     TitleActivationStatusPendingSteam TitleActivationStatus = "PendingSteam"
     TitleActivationStatusActivatedSteam TitleActivationStatus = "ActivatedSteam"
     TitleActivationStatusRevokedSteam TitleActivationStatus = "RevokedSteam"
      )
// TitleNewsItem 
type TitleNewsItemModel struct {
    // Body news item body.
    Body string `json:"Body,omitempty"`
    // NewsId unique identifier of news item.
    NewsId string `json:"NewsId,omitempty"`
    // Timestamp date and time when the news item was posted.
    Timestamp time.Time `json:"Timestamp,omitempty"`
    // Title title of the news item.
    Title string `json:"Title,omitempty"`
}

// TreatmentAssignment 
type TreatmentAssignmentModel struct {
    // Variables list of the experiment variables.
    Variables []VariableModel `json:"Variables,omitempty"`
    // Variants list of the experiment variants.
    Variants []string `json:"Variants,omitempty"`
}

// UnlinkPSNAccountRequest 
type UnlinkPSNAccountRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// UnlinkPSNAccountResult 
type UnlinkPSNAccountResultModel struct {
}

// UnlinkServerCustomIdRequest 
type UnlinkServerCustomIdRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // PlayFabId unique PlayFab identifier.
    PlayFabId string `json:"PlayFabId,omitempty"`
    // ServerCustomId unique server custom identifier for this player.
    ServerCustomId string `json:"ServerCustomId,omitempty"`
}

// UnlinkServerCustomIdResult 
type UnlinkServerCustomIdResultModel struct {
}

// UnlinkXboxAccountRequest 
type UnlinkXboxAccountRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // PlayFabId unique PlayFab identifier for a user, or null if no PlayFab account is linked to the Xbox Live identifier.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// UnlinkXboxAccountResult 
type UnlinkXboxAccountResultModel struct {
}

// UnlockContainerInstanceRequest specify the container and optionally the catalogVersion for the container to open
type UnlockContainerInstanceRequestModel struct {
    // CatalogVersion specifies the catalog version that should be used to determine container contents. If unspecified, uses catalog
// associated with the item instance.
    CatalogVersion string `json:"CatalogVersion,omitempty"`
    // CharacterId unique PlayFab assigned ID for a specific character owned by a user
    CharacterId string `json:"CharacterId,omitempty"`
    // ContainerItemInstanceId itemInstanceId of the container to unlock.
    ContainerItemInstanceId string `json:"ContainerItemInstanceId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // KeyItemInstanceId itemInstanceId of the key that will be consumed by unlocking this container. If the container requires a key, this
// parameter is required.
    KeyItemInstanceId string `json:"KeyItemInstanceId,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// UnlockContainerItemRequest specify the type of container to open and optionally the catalogVersion for the container to open
type UnlockContainerItemRequestModel struct {
    // CatalogVersion specifies the catalog version that should be used to determine container contents. If unspecified, uses default/primary
// catalog.
    CatalogVersion string `json:"CatalogVersion,omitempty"`
    // CharacterId unique PlayFab assigned ID for a specific character owned by a user
    CharacterId string `json:"CharacterId,omitempty"`
    // ContainerItemId catalog ItemId of the container type to unlock.
    ContainerItemId string `json:"ContainerItemId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// UnlockContainerItemResult the items and vc found within the container. These will be added and stacked in your inventory as appropriate.
type UnlockContainerItemResultModel struct {
    // GrantedItems items granted to the player as a result of unlocking the container.
    GrantedItems []ItemInstanceModel `json:"GrantedItems,omitempty"`
    // UnlockedItemInstanceId unique instance identifier of the container unlocked.
    UnlockedItemInstanceId string `json:"UnlockedItemInstanceId,omitempty"`
    // UnlockedWithItemInstanceId unique instance identifier of the key used to unlock the container, if applicable.
    UnlockedWithItemInstanceId string `json:"UnlockedWithItemInstanceId,omitempty"`
    // VirtualCurrency virtual currency granted to the player as a result of unlocking the container.
    VirtualCurrency map[string]uint32 `json:"VirtualCurrency,omitempty"`
}

// UpdateAvatarUrlRequest 
type UpdateAvatarUrlRequestModel struct {
    // ImageUrl uRL of the avatar image. If empty, it removes the existing avatar URL.
    ImageUrl string `json:"ImageUrl,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
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

// UpdateCharacterDataRequest this function performs an additive update of the arbitrary JSON object containing the custom data for the user. In
// updating the custom data object, keys which already exist in the object will have their values overwritten, while keys
// with null values will be removed. No other key-value pairs will be changed apart from those specified in the call.
type UpdateCharacterDataRequestModel struct {
    // CharacterId unique PlayFab assigned ID for a specific character owned by a user
    CharacterId string `json:"CharacterId,omitempty"`
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

// UpdateCharacterDataResult 
type UpdateCharacterDataResultModel struct {
    // DataVersion indicates the current version of the data that has been set. This is incremented with every set call for that type of
// data (read-only, internal, etc). This version can be provided in Get calls to find updated data.
    DataVersion uint32 `json:"DataVersion,omitempty"`
}

// UpdateCharacterStatisticsRequest character statistics are similar to user statistics in that they are numeric values which may only be updated by a
// server operation, in order to minimize the opportunity for unauthorized changes. In addition to being available for use
// by the title, the statistics are used for all leaderboard operations in PlayFab.
type UpdateCharacterStatisticsRequestModel struct {
    // CharacterId unique PlayFab assigned ID for a specific character owned by a user
    CharacterId string `json:"CharacterId,omitempty"`
    // CharacterStatistics statistics to be updated with the provided values.
    CharacterStatistics map[string]int32 `json:"CharacterStatistics,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// UpdateCharacterStatisticsResult 
type UpdateCharacterStatisticsResultModel struct {
}

// UpdatePlayerStatisticsRequest this operation is additive. Statistics not currently defined will be added, while those already defined will be updated
// with the given values. All other user statistics will remain unchanged.
type UpdatePlayerStatisticsRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // ForceUpdate indicates whether the statistics provided should be set, regardless of the aggregation method set on the statistic.
// Default is false.
    ForceUpdate bool `json:"ForceUpdate"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
    // Statistics statistics to be updated with the provided values
    Statistics []StatisticUpdateModel `json:"Statistics,omitempty"`
}

// UpdatePlayerStatisticsResult 
type UpdatePlayerStatisticsResultModel struct {
}

// UpdateSharedGroupDataRequest note that in the case of multiple calls to write to the same shared group data keys, the last write received by the
// PlayFab service will determine the value available to subsequent read operations. For scenarios requiring coordination
// of data updates, it is recommended that titles make use of user data with read permission set to public, or a
// combination of user data and shared group data.
type UpdateSharedGroupDataRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Data key-value pairs to be written to the custom data. Note that keys are trimmed of whitespace, are limited in size, and may
// not begin with a '!' character or be null.
    Data map[string]string `json:"Data,omitempty"`
    // KeysToRemove optional list of Data-keys to remove from UserData. Some SDKs cannot insert null-values into Data due to language
// constraints. Use this to delete the keys directly.
    KeysToRemove []string `json:"KeysToRemove,omitempty"`
    // Permission permission to be applied to all user data keys in this request.
    Permission UserDataPermission `json:"Permission,omitempty"`
    // SharedGroupId unique identifier for the shared group.
    SharedGroupId string `json:"SharedGroupId,omitempty"`
}

// UpdateSharedGroupDataResult 
type UpdateSharedGroupDataResultModel struct {
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

// UpdateUserInventoryItemDataRequest this function performs an additive update of the arbitrary JSON object containing the custom data for the item instance
// which belongs to the specified user. In updating the custom data object, keys which already exist in the object will
// have their values overwritten, while keys with null values will be removed. No other key-value pairs will be changed
// apart from those specified in the call.
type UpdateUserInventoryItemDataRequestModel struct {
    // CharacterId unique PlayFab assigned ID for a specific character owned by a user
    CharacterId string `json:"CharacterId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Data key-value pairs to be written to the custom data. Note that keys are trimmed of whitespace, are limited in size, and may
// not begin with a '!' character or be null.
    Data map[string]string `json:"Data,omitempty"`
    // ItemInstanceId unique PlayFab assigned instance identifier of the item
    ItemInstanceId string `json:"ItemInstanceId,omitempty"`
    // KeysToRemove optional list of Data-keys to remove from UserData. Some SDKs cannot insert null-values into Data due to language
// constraints. Use this to delete the keys directly.
    KeysToRemove []string `json:"KeysToRemove,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
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

// UserSettings 
type UserSettingsModel struct {
    // GatherDeviceInfo boolean for whether this player is eligible for gathering device info.
    GatherDeviceInfo bool `json:"GatherDeviceInfo"`
    // GatherFocusInfo boolean for whether this player should report OnFocus play-time tracking.
    GatherFocusInfo bool `json:"GatherFocusInfo"`
    // NeedsAttribution boolean for whether this player is eligible for ad tracking.
    NeedsAttribution bool `json:"NeedsAttribution"`
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

// Variable 
type VariableModel struct {
    // Name name of the variable.
    Name string `json:"Name,omitempty"`
    // Value value of the variable.
    Value string `json:"Value,omitempty"`
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

// WriteEventResponse 
type WriteEventResponseModel struct {
    // EventId the unique identifier of the event. The values of this identifier consist of ASCII characters and are not constrained to
// any particular format.
    EventId string `json:"EventId,omitempty"`
}

// WriteServerCharacterEventRequest this API is designed to write a multitude of different event types into PlayStream. It supports a flexible JSON schema,
// which allowsfor arbitrary key-value pairs to describe any character-based event. The created event will be locked to the
// authenticated title.
type WriteServerCharacterEventRequestModel struct {
    // Body custom event properties. Each property consists of a name (string) and a value (JSON object).
    Body map[string]interface{} `json:"Body,omitempty"`
    // CharacterId unique PlayFab assigned ID for a specific character owned by a user
    CharacterId string `json:"CharacterId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // EventName the name of the event, within the namespace scoped to the title. The naming convention is up to the caller, but it
// commonly follows the subject_verb_object pattern (e.g. player_logged_in).
    EventName string `json:"EventName,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
    // Timestamp the time (in UTC) associated with this event. The value defaults to the current time.
    Timestamp time.Time `json:"Timestamp,omitempty"`
}

// WriteServerPlayerEventRequest this API is designed to write a multitude of different event types into PlayStream. It supports a flexible JSON schema,
// which allowsfor arbitrary key-value pairs to describe any player-based event. The created event will be locked to the
// authenticated title.
type WriteServerPlayerEventRequestModel struct {
    // Body custom data properties associated with the event. Each property consists of a name (string) and a value (JSON object).
    Body map[string]interface{} `json:"Body,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // EventName the name of the event, within the namespace scoped to the title. The naming convention is up to the caller, but it
// commonly follows the subject_verb_object pattern (e.g. player_logged_in).
    EventName string `json:"EventName,omitempty"`
    // PlayFabId unique PlayFab assigned ID of the user on whom the operation will be performed.
    PlayFabId string `json:"PlayFabId,omitempty"`
    // Timestamp the time (in UTC) associated with this event. The value defaults to the current time.
    Timestamp time.Time `json:"Timestamp,omitempty"`
}

// WriteTitleEventRequest this API is designed to write a multitude of different event types into PlayStream. It supports a flexible JSON schema,
// which allowsfor arbitrary key-value pairs to describe any title-based event. The created event will be locked to the
// authenticated title.
type WriteTitleEventRequestModel struct {
    // Body custom event properties. Each property consists of a name (string) and a value (JSON object).
    Body map[string]interface{} `json:"Body,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // EventName the name of the event, within the namespace scoped to the title. The naming convention is up to the caller, but it
// commonly follows the subject_verb_object pattern (e.g. player_logged_in).
    EventName string `json:"EventName,omitempty"`
    // Timestamp the time (in UTC) associated with this event. The value defaults to the current time.
    Timestamp time.Time `json:"Timestamp,omitempty"`
}

// XboxLiveAccountPlayFabIdPair 
type XboxLiveAccountPlayFabIdPairModel struct {
    // PlayFabId unique PlayFab identifier for a user, or null if no PlayFab account is linked to the Xbox Live identifier.
    PlayFabId string `json:"PlayFabId,omitempty"`
    // XboxLiveAccountId unique Xbox Live identifier for a user.
    XboxLiveAccountId string `json:"XboxLiveAccountId,omitempty"`
}
