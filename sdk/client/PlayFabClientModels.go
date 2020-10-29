package client

import "time"
                    
// AcceptTradeRequest 
type AcceptTradeRequestModel struct {
    // AcceptedInventoryInstanceIds items from the accepting player's inventory in exchange for the offered items in the trade. In the case of a gift, this
// will be null.
    AcceptedInventoryInstanceIds []string `json:"AcceptedInventoryInstanceIds,omitempty"`
    // OfferingPlayerId player who opened the trade.
    OfferingPlayerId string `json:"OfferingPlayerId,omitempty"`
    // TradeId trade identifier.
    TradeId string `json:"TradeId,omitempty"`
}

// AcceptTradeResponse 
type AcceptTradeResponseModel struct {
    // Trade details about trade which was just accepted.
    Trade *TradeInfoModel `json:"Trade,omitempty"`
}

// AdActivity 
type AdActivity string
  
const (
     AdActivityOpened AdActivity = "Opened"
     AdActivityClosed AdActivity = "Closed"
     AdActivityStart AdActivity = "Start"
     AdActivityEnd AdActivity = "End"
      )
// AdCampaignAttributionModel 
type AdCampaignAttributionModelModel struct {
    // AttributedAt uTC time stamp of attribution
    AttributedAt time.Time `json:"AttributedAt,omitempty"`
    // CampaignId attribution campaign identifier
    CampaignId string `json:"CampaignId,omitempty"`
    // Platform attribution network name
    Platform string `json:"Platform,omitempty"`
}

// AddFriendRequest 
type AddFriendRequestModel struct {
    // FriendEmail email address of the user to attempt to add to the local user's friend list.
    FriendEmail string `json:"FriendEmail,omitempty"`
    // FriendPlayFabId playFab identifier of the user to attempt to add to the local user's friend list.
    FriendPlayFabId string `json:"FriendPlayFabId,omitempty"`
    // FriendTitleDisplayName title-specific display name of the user to attempt to add to the local user's friend list.
    FriendTitleDisplayName string `json:"FriendTitleDisplayName,omitempty"`
    // FriendUsername playFab username of the user to attempt to add to the local user's friend list.
    FriendUsername string `json:"FriendUsername,omitempty"`
}

// AddFriendResult 
type AddFriendResultModel struct {
    // Created true if the friend request was processed successfully.
    Created bool `json:"Created"`
}

// AddGenericIDRequest 
type AddGenericIDRequestModel struct {
    // GenericId generic service identifier to add to the player account.
    GenericId* GenericServiceIdModel `json:"GenericId,omitempty"`
}

// AddGenericIDResult 
type AddGenericIDResultModel struct {
}

// AddOrUpdateContactEmailRequest this API adds a contact email to the player's profile. If the player's profile already contains a contact email, it will
// update the contact email to the email address specified.
type AddOrUpdateContactEmailRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // EmailAddress the new contact email to associate with the player.
    EmailAddress string `json:"EmailAddress,omitempty"`
}

// AddOrUpdateContactEmailResult 
type AddOrUpdateContactEmailResultModel struct {
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

// AddUsernamePasswordRequest 
type AddUsernamePasswordRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Email user email address attached to their account
    Email string `json:"Email,omitempty"`
    // Password password for the PlayFab account (6-100 characters)
    Password string `json:"Password,omitempty"`
    // Username playFab username for the account (3-20 characters)
    Username string `json:"Username,omitempty"`
}

// AddUsernamePasswordResult each account must have a unique username and email address in the PlayFab service. Once created, the account may be
// associated with additional accounts (Steam, Facebook, Game Center, etc.), allowing for added social network lists and
// achievements systems. This can also be used to provide a recovery method if the user loses their original means of
// access.
type AddUsernamePasswordResultModel struct {
    // Username playFab unique user name.
    Username string `json:"Username,omitempty"`
}

// AddUserVirtualCurrencyRequest this API must be enabled for use as an option in the game manager website. It is disabled by default.
type AddUserVirtualCurrencyRequestModel struct {
    // Amount amount to be added to the user balance of the specified virtual currency.
    Amount int32 `json:"Amount,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // VirtualCurrency name of the virtual currency which is to be incremented.
    VirtualCurrency string `json:"VirtualCurrency,omitempty"`
}

// AdPlacementDetails a single ad placement details including placement and reward information
type AdPlacementDetailsModel struct {
    // PlacementId placement unique ID
    PlacementId string `json:"PlacementId,omitempty"`
    // PlacementName placement name
    PlacementName string `json:"PlacementName,omitempty"`
    // PlacementViewsRemaining if placement has viewing limits indicates how many views are left
    PlacementViewsRemaining int32 `json:"PlacementViewsRemaining,omitempty"`
    // PlacementViewsResetMinutes if placement has viewing limits indicates when they will next reset
    PlacementViewsResetMinutes float64 `json:"PlacementViewsResetMinutes,omitempty"`
    // RewardAssetUrl optional URL to a reward asset
    RewardAssetUrl string `json:"RewardAssetUrl,omitempty"`
    // RewardDescription reward description
    RewardDescription string `json:"RewardDescription,omitempty"`
    // RewardId reward unique ID
    RewardId string `json:"RewardId,omitempty"`
    // RewardName reward name
    RewardName string `json:"RewardName,omitempty"`
}

// AdRewardItemGranted details for each item granted
type AdRewardItemGrantedModel struct {
    // CatalogId catalog ID
    CatalogId string `json:"CatalogId,omitempty"`
    // DisplayName catalog item display name
    DisplayName string `json:"DisplayName,omitempty"`
    // InstanceId inventory instance ID
    InstanceId string `json:"InstanceId,omitempty"`
    // ItemId item ID
    ItemId string `json:"ItemId,omitempty"`
}

// AdRewardResults details on what was granted to the player
type AdRewardResultsModel struct {
    // GrantedItems array of the items granted to the player
    GrantedItems []AdRewardItemGrantedModel `json:"GrantedItems,omitempty"`
    // GrantedVirtualCurrencies dictionary of virtual currencies that were granted to the player
    GrantedVirtualCurrencies map[string]int32 `json:"GrantedVirtualCurrencies,omitempty"`
    // IncrementedStatistics dictionary of statistics that were modified for the player
    IncrementedStatistics map[string]int32 `json:"IncrementedStatistics,omitempty"`
}

// AndroidDevicePushNotificationRegistrationRequest more information can be found on configuring your game for the Google Cloud Messaging service in the Google developer
// documentation, here: http://developer.android.com/google/gcm/client.html. The steps to configure and send Push
// Notifications is described in the PlayFab tutorials, here:
// https://docs.microsoft.com/gaming/playfab/features/engagement/push-notifications/quickstart.
type AndroidDevicePushNotificationRegistrationRequestModel struct {
    // ConfirmationMessage message to display when confirming push notification.
    ConfirmationMessage string `json:"ConfirmationMessage,omitempty"`
    // DeviceToken registration ID provided by the Google Cloud Messaging service when the title registered to receive push notifications
// (see the GCM documentation, here: http://developer.android.com/google/gcm/client.html).
    DeviceToken string `json:"DeviceToken,omitempty"`
    // SendPushNotificationConfirmation if true, send a test push message immediately after sucessful registration. Defaults to false.
    SendPushNotificationConfirmation bool `json:"SendPushNotificationConfirmation"`
}

// AndroidDevicePushNotificationRegistrationResult 
type AndroidDevicePushNotificationRegistrationResultModel struct {
}

// AttributeInstallRequest if you have an ad attribution partner enabled, this will post an install to their service to track the device. It uses
// the given device id to match based on clicks on ads.
type AttributeInstallRequestModel struct {
    // Adid the adid for this device.
    Adid string `json:"Adid,omitempty"`
    // Idfa the IdentifierForAdvertisers for iOS Devices.
    Idfa string `json:"Idfa,omitempty"`
}

// AttributeInstallResult 
type AttributeInstallResultModel struct {
}

// CancelTradeRequest 
type CancelTradeRequestModel struct {
    // TradeId trade identifier.
    TradeId string `json:"TradeId,omitempty"`
}

// CancelTradeResponse 
type CancelTradeResponseModel struct {
    // Trade details about trade which was just canceled.
    Trade *TradeInfoModel `json:"Trade,omitempty"`
}

// CartItem 
type CartItemModel struct {
    // Description description of the catalog item.
    Description string `json:"Description,omitempty"`
    // DisplayName display name for the catalog item.
    DisplayName string `json:"DisplayName,omitempty"`
    // ItemClass class name to which catalog item belongs.
    ItemClass string `json:"ItemClass,omitempty"`
    // ItemId unique identifier for the catalog item.
    ItemId string `json:"ItemId,omitempty"`
    // ItemInstanceId unique instance identifier for this catalog item.
    ItemInstanceId string `json:"ItemInstanceId,omitempty"`
    // RealCurrencyPrices cost of the catalog item for each applicable real world currency.
    RealCurrencyPrices map[string]uint32 `json:"RealCurrencyPrices,omitempty"`
    // VCAmount amount of each applicable virtual currency which will be received as a result of purchasing this catalog item.
    VCAmount map[string]uint32 `json:"VCAmount,omitempty"`
    // VirtualCurrencyPrices cost of the catalog item for each applicable virtual currency.
    VirtualCurrencyPrices map[string]uint32 `json:"VirtualCurrencyPrices,omitempty"`
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
// CollectionFilter collection filter to include and/or exclude collections with certain key-value pairs. The filter generates a collection
// set defined by Includes rules and then remove collections that matches the Excludes rules. A collection is considered
// matching a rule if the rule describes a subset of the collection.
type CollectionFilterModel struct {
    // Excludes list of Exclude rules, with any of which if a collection matches, it is excluded by the filter.
    Excludes []Container_Dictionary_String_StringModel `json:"Excludes,omitempty"`
    // Includes list of Include rules, with any of which if a collection matches, it is included by the filter, unless it is excluded by
// one of the Exclude rule
    Includes []Container_Dictionary_String_StringModel `json:"Includes,omitempty"`
}

// ConfirmPurchaseRequest the final step in the purchasing process, this API finalizes the purchase with the payment provider, where applicable,
// adding virtual goods to the player inventory (including random drop table resolution and recursive addition of bundled
// items) and adjusting virtual currency balances for funds used or added. Note that this is a pull operation, and should
// be polled regularly when a purchase is in progress. Please note that the processing time for inventory grants and
// purchases increases fractionally the more items are in the inventory, and the more items are in the grant/purchase
// operation.
type ConfirmPurchaseRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // OrderId purchase order identifier returned from StartPurchase.
    OrderId string `json:"OrderId,omitempty"`
}

// ConfirmPurchaseResult when the FailedByPaymentProvider error is returned, it's important to check the ProviderErrorCode, ProviderErrorMessage,
// and ProviderErrorDetails to understand the specific reason the payment was rejected, as in some rare cases, this may
// mean that the provider hasn't completed some operation required to finalize the purchase.
type ConfirmPurchaseResultModel struct {
    // Items array of items purchased.
    Items []ItemInstanceModel `json:"Items,omitempty"`
    // OrderId purchase order identifier.
    OrderId string `json:"OrderId,omitempty"`
    // PurchaseDate date and time of the purchase.
    PurchaseDate time.Time `json:"PurchaseDate,omitempty"`
}

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
}

// ConsumeItemResult 
type ConsumeItemResultModel struct {
    // ItemInstanceId unique instance identifier of the item with uses consumed.
    ItemInstanceId string `json:"ItemInstanceId,omitempty"`
    // RemainingUses number of uses remaining on the item.
    RemainingUses int32 `json:"RemainingUses,omitempty"`
}

// ConsumeMicrosoftStoreEntitlementsRequest 
type ConsumeMicrosoftStoreEntitlementsRequestModel struct {
    // CatalogVersion catalog version to use
    CatalogVersion string `json:"CatalogVersion,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // MarketplaceSpecificData marketplace specific payload containing details to fetch in app purchase transactions
    MarketplaceSpecificData* MicrosoftStorePayloadModel `json:"MarketplaceSpecificData,omitempty"`
}

// ConsumeMicrosoftStoreEntitlementsResponse 
type ConsumeMicrosoftStoreEntitlementsResponseModel struct {
    // Items details for the items purchased.
    Items []ItemInstanceModel `json:"Items,omitempty"`
}

// ConsumePSNEntitlementsRequest 
type ConsumePSNEntitlementsRequestModel struct {
    // CatalogVersion which catalog to match granted entitlements against. If null, defaults to title default catalog
    CatalogVersion string `json:"CatalogVersion,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // ServiceLabel id of the PSN service label to consume entitlements from
    ServiceLabel int32 `json:"ServiceLabel,omitempty"`
}

// ConsumePSNEntitlementsResult 
type ConsumePSNEntitlementsResultModel struct {
    // ItemsGranted array of items granted to the player as a result of consuming entitlements.
    ItemsGranted []ItemInstanceModel `json:"ItemsGranted,omitempty"`
}

// ConsumeXboxEntitlementsRequest 
type ConsumeXboxEntitlementsRequestModel struct {
    // CatalogVersion catalog version to use
    CatalogVersion string `json:"CatalogVersion,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // XboxToken token provided by the Xbox Live SDK/XDK method GetTokenAndSignatureAsync("POST", "https://playfabapi.com/", "").
    XboxToken string `json:"XboxToken,omitempty"`
}

// ConsumeXboxEntitlementsResult 
type ConsumeXboxEntitlementsResultModel struct {
    // Items details for the items purchased.
    Items []ItemInstanceModel `json:"Items,omitempty"`
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

// Container_Dictionary_String_String a data container
type Container_Dictionary_String_StringModel struct {
    // Data content of data
    Data map[string]string `json:"Data,omitempty"`
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
// CurrentGamesRequest 
type CurrentGamesRequestModel struct {
    // BuildVersion build to match against.
    BuildVersion string `json:"BuildVersion,omitempty"`
    // GameMode game mode to look for.
    GameMode string `json:"GameMode,omitempty"`
    // Region region to check for Game Server Instances.
    Region Region `json:"Region,omitempty"`
    // StatisticName statistic name to find statistic-based matches.
    StatisticName string `json:"StatisticName,omitempty"`
    // TagFilter filter to include and/or exclude Game Server Instances associated with certain tags.
    TagFilter *CollectionFilterModel `json:"TagFilter,omitempty"`
}

// CurrentGamesResult 
type CurrentGamesResultModel struct {
    // GameCount number of games running
    GameCount int32 `json:"GameCount,omitempty"`
    // Games array of games found
    Games []GameInfoModel `json:"Games,omitempty"`
    // PlayerCount total number of players across all servers
    PlayerCount int32 `json:"PlayerCount,omitempty"`
}

// DeviceInfoRequest any arbitrary information collected by the device
type DeviceInfoRequestModel struct {
    // Info information posted to the PlayStream Event. Currently arbitrary, and specific to the environment sending it.
    Info map[string]interface{} `json:"Info,omitempty"`
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

// ExecuteCloudScriptRequest 
type ExecuteCloudScriptRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // FunctionName the name of the CloudScript function to execute
    FunctionName string `json:"FunctionName,omitempty"`
    // FunctionParameter object that is passed in to the function as the first argument
    FunctionParameter interface{} `json:"FunctionParameter,omitempty"`
    // GeneratePlayStreamEvent generate a 'player_executed_cloudscript' PlayStream event containing the results of the function execution and other
// contextual information. This event will show up in the PlayStream debugger console for the player in Game Manager.
    GeneratePlayStreamEvent bool `json:"GeneratePlayStreamEvent"`
    // RevisionSelection option for which revision of the CloudScript to execute. 'Latest' executes the most recently created revision, 'Live'
// executes the current live, published revision, and 'Specific' executes the specified revision. The default value is
// 'Specific', if the SpeificRevision parameter is specified, otherwise it is 'Live'.
    RevisionSelection CloudScriptRevisionOption `json:"RevisionSelection,omitempty"`
    // SpecificRevision the specivic revision to execute, when RevisionSelection is set to 'Specific'
    SpecificRevision int32 `json:"SpecificRevision,omitempty"`
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

// GameCenterPlayFabIdPair 
type GameCenterPlayFabIdPairModel struct {
    // GameCenterId unique Game Center identifier for a user.
    GameCenterId string `json:"GameCenterId,omitempty"`
    // PlayFabId unique PlayFab identifier for a user, or null if no PlayFab account is linked to the Game Center identifier.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// GameInfo 
type GameInfoModel struct {
    // BuildVersion build version this server is running
    BuildVersion string `json:"BuildVersion,omitempty"`
    // GameMode game mode this server is running
    GameMode string `json:"GameMode,omitempty"`
    // GameServerData game session custom data
    GameServerData string `json:"GameServerData,omitempty"`
    // GameServerStateEnum game specific string denoting server configuration
    GameServerStateEnum GameInstanceState `json:"GameServerStateEnum,omitempty"`
    // LastHeartbeat last heartbeat of the game server instance, used in external game server provider mode
    LastHeartbeat time.Time `json:"LastHeartbeat,omitempty"`
    // LobbyID unique lobby identifier for this game server
    LobbyID string `json:"LobbyID,omitempty"`
    // MaxPlayers maximum players this server can support
    MaxPlayers int32 `json:"MaxPlayers,omitempty"`
    // PlayerUserIds array of current player IDs on this server
    PlayerUserIds []string `json:"PlayerUserIds,omitempty"`
    // Region region to which this server is associated
    Region Region `json:"Region,omitempty"`
    // RunTime duration in seconds this server has been running
    RunTime uint32 `json:"RunTime,omitempty"`
    // ServerIPV4Address iPV4 address of the server
    ServerIPV4Address string `json:"ServerIPV4Address,omitempty"`
    // ServerIPV6Address iPV6 address of the server
    ServerIPV6Address string `json:"ServerIPV6Address,omitempty"`
    // ServerPort port number to use for non-http communications with the server
    ServerPort int32 `json:"ServerPort,omitempty"`
    // ServerPublicDNSName public DNS name (if any) of the server
    ServerPublicDNSName string `json:"ServerPublicDNSName,omitempty"`
    // StatisticName stastic used to match this game in player statistic matchmaking
    StatisticName string `json:"StatisticName,omitempty"`
    // Tags game session tags
    Tags map[string]string `json:"Tags,omitempty"`
}

// GameInstanceState 
type GameInstanceState string
  
const (
     GameInstanceStateOpen GameInstanceState = "Open"
     GameInstanceStateClosed GameInstanceState = "Closed"
      )
// GameServerRegionsRequest 
type GameServerRegionsRequestModel struct {
    // BuildVersion version of game server for which stats are being requested
    BuildVersion string `json:"BuildVersion,omitempty"`
    // TitleId unique identifier for the title, found in the Settings > Game Properties section of the PlayFab developer site when a
// title has been selected.
    TitleId string `json:"TitleId,omitempty"`
}

// GameServerRegionsResult 
type GameServerRegionsResultModel struct {
    // Regions array of regions found matching the request parameters
    Regions []RegionInfoModel `json:"Regions,omitempty"`
}

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

// GetAccountInfoRequest 
type GetAccountInfoRequestModel struct {
    // Email user email address for the account to find (if no Username is specified).
    Email string `json:"Email,omitempty"`
    // PlayFabId unique PlayFab identifier of the user whose info is being requested. Optional, defaults to the authenticated user if no
// other lookup identifier set.
    PlayFabId string `json:"PlayFabId,omitempty"`
    // TitleDisplayName title-specific username for the account to find (if no Email is set). Note that if the non-unique Title Display Names
// option is enabled for the title, attempts to look up users by Title Display Name will always return AccountNotFound.
    TitleDisplayName string `json:"TitleDisplayName,omitempty"`
    // Username playFab Username for the account to find (if no PlayFabId is specified).
    Username string `json:"Username,omitempty"`
}

// GetAccountInfoResult this API retrieves details regarding the player in the PlayFab service. Note that when this call is used to retrieve
// data about another player (not the one signed into the local client), some data, such as Personally Identifying
// Information (PII), will be omitted for privacy reasons or to comply with the requirements of the platform belongs to.
// The user account returned will be based on the identifier provided in priority order: PlayFabId, Username, Email, then
// TitleDisplayName. If no identifier is specified, the currently signed in user's information will be returned.
type GetAccountInfoResultModel struct {
    // AccountInfo account information for the local user.
    AccountInfo *UserAccountInfoModel `json:"AccountInfo,omitempty"`
}

// GetAdPlacementsRequest using an AppId to return a list of valid ad placements for a player.
type GetAdPlacementsRequestModel struct {
    // AppId the current AppId to use
    AppId string `json:"AppId,omitempty"`
    // Identifier using the name or unique identifier, filter the result for get a specific placement.
    Identifier *NameIdentifierModel `json:"Identifier,omitempty"`
}

// GetAdPlacementsResult array of AdPlacementDetails
type GetAdPlacementsResultModel struct {
    // AdPlacements array of results
    AdPlacements []AdPlacementDetailsModel `json:"AdPlacements,omitempty"`
}

// GetCatalogItemsRequest 
type GetCatalogItemsRequestModel struct {
    // CatalogVersion which catalog is being requested. If null, uses the default catalog.
    CatalogVersion string `json:"CatalogVersion,omitempty"`
}

// GetCatalogItemsResult if CatalogVersion is not specified, only inventory items associated with the most recent version of the catalog will be
// returned.
type GetCatalogItemsResultModel struct {
    // Catalog array of items which can be purchased.
    Catalog []CatalogItemModel `json:"Catalog,omitempty"`
}

// GetCharacterDataRequest data is stored as JSON key-value pairs. If the Keys parameter is provided, the data object returned will only contain
// the data specific to the indicated Keys. Otherwise, the full set of custom character data will be returned.
type GetCharacterDataRequestModel struct {
    // CharacterId unique PlayFab assigned ID for a specific character owned by a user
    CharacterId string `json:"CharacterId,omitempty"`
    // IfChangedFromDataVersion the version that currently exists according to the caller. The call will return the data for all of the keys if the
// version in the system is greater than this.
    IfChangedFromDataVersion uint32 `json:"IfChangedFromDataVersion,omitempty"`
    // Keys specific keys to search for in the custom user data.
    Keys []string `json:"Keys,omitempty"`
    // PlayFabId unique PlayFab identifier of the user to load data for. Optional, defaults to yourself if not set.
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
}

// GetCharacterInventoryResult 
type GetCharacterInventoryResultModel struct {
    // CharacterId unique identifier of the character for this inventory.
    CharacterId string `json:"CharacterId,omitempty"`
    // Inventory array of inventory items belonging to the character.
    Inventory []ItemInstanceModel `json:"Inventory,omitempty"`
    // VirtualCurrency array of virtual currency balance(s) belonging to the character.
    VirtualCurrency map[string]int32 `json:"VirtualCurrency,omitempty"`
    // VirtualCurrencyRechargeTimes array of remaining times and timestamps for virtual currencies.
    VirtualCurrencyRechargeTimes map[string]VirtualCurrencyRechargeTimeModel `json:"VirtualCurrencyRechargeTimes,omitempty"`
}

// GetCharacterLeaderboardRequest 
type GetCharacterLeaderboardRequestModel struct {
    // CharacterType optional character type on which to filter the leaderboard entries.
    CharacterType string `json:"CharacterType,omitempty"`
    // MaxResultsCount maximum number of entries to retrieve. Default 10, maximum 100.
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

// GetCharacterStatisticsRequest 
type GetCharacterStatisticsRequestModel struct {
    // CharacterId unique PlayFab assigned ID for a specific character owned by a user
    CharacterId string `json:"CharacterId,omitempty"`
}

// GetCharacterStatisticsResult in addition to being available for use by the title, the statistics are used for all leaderboard operations in PlayFab.
type GetCharacterStatisticsResultModel struct {
    // CharacterStatistics the requested character statistics.
    CharacterStatistics map[string]int32 `json:"CharacterStatistics,omitempty"`
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

// GetFriendLeaderboardAroundPlayerRequest 
type GetFriendLeaderboardAroundPlayerRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // IncludeFacebookFriends indicates whether Facebook friends should be included in the response. Default is true.
    IncludeFacebookFriends bool `json:"IncludeFacebookFriends"`
    // IncludeSteamFriends indicates whether Steam service friends should be included in the response. Default is true.
    IncludeSteamFriends bool `json:"IncludeSteamFriends"`
    // MaxResultsCount maximum number of entries to retrieve. Default 10, maximum 100.
    MaxResultsCount int32 `json:"MaxResultsCount,omitempty"`
    // PlayFabId playFab unique identifier of the user to center the leaderboard around. If null will center on the logged in user.
    PlayFabId string `json:"PlayFabId,omitempty"`
    // ProfileConstraints if non-null, this determines which properties of the resulting player profiles to return. For API calls from the client,
// only the allowed client profile properties for the title may be requested. These allowed properties are configured in
// the Game Manager "Client Profile Options" tab in the "Settings" section.
    ProfileConstraints *PlayerProfileViewConstraintsModel `json:"ProfileConstraints,omitempty"`
    // StatisticName statistic used to rank players for this leaderboard.
    StatisticName string `json:"StatisticName,omitempty"`
    // Version the version of the leaderboard to get.
    Version int32 `json:"Version,omitempty"`
    // XboxToken xbox token if Xbox friends should be included. Requires Xbox be configured on PlayFab.
    XboxToken string `json:"XboxToken,omitempty"`
}

// GetFriendLeaderboardAroundPlayerResult note: When calling 'GetLeaderboardAround...' APIs, the position of the user defaults to 0 when the user does not have
// the corresponding statistic.If Facebook friends are included, make sure the access token from previous LoginWithFacebook
// call is still valid and not expired. If Xbox Live friends are included, make sure the access token from the previous
// LoginWithXbox call is still valid and not expired.
type GetFriendLeaderboardAroundPlayerResultModel struct {
    // Leaderboard ordered listing of users and their positions in the requested leaderboard.
    Leaderboard []PlayerLeaderboardEntryModel `json:"Leaderboard,omitempty"`
    // NextReset the time the next scheduled reset will occur. Null if the leaderboard does not reset on a schedule.
    NextReset time.Time `json:"NextReset,omitempty"`
    // Version the version of the leaderboard returned.
    Version int32 `json:"Version,omitempty"`
}

// GetFriendLeaderboardRequest 
type GetFriendLeaderboardRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // IncludeFacebookFriends indicates whether Facebook friends should be included in the response. Default is true.
    IncludeFacebookFriends bool `json:"IncludeFacebookFriends"`
    // IncludeSteamFriends indicates whether Steam service friends should be included in the response. Default is true.
    IncludeSteamFriends bool `json:"IncludeSteamFriends"`
    // MaxResultsCount maximum number of entries to retrieve. Default 10, maximum 100.
    MaxResultsCount int32 `json:"MaxResultsCount,omitempty"`
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
    // CharacterId unique PlayFab assigned ID for a specific character on which to center the leaderboard.
    CharacterId string `json:"CharacterId,omitempty"`
    // CharacterType optional character type on which to filter the leaderboard entries.
    CharacterType string `json:"CharacterType,omitempty"`
    // MaxResultsCount maximum number of entries to retrieve. Default 10, maximum 100.
    MaxResultsCount int32 `json:"MaxResultsCount,omitempty"`
    // StatisticName unique identifier for the title-specific statistic for the leaderboard.
    StatisticName string `json:"StatisticName,omitempty"`
}

// GetLeaderboardAroundCharacterResult note: When calling 'GetLeaderboardAround...' APIs, the position of the character defaults to 0 when the character does
// not have the corresponding statistic.
type GetLeaderboardAroundCharacterResultModel struct {
    // Leaderboard ordered list of leaderboard entries.
    Leaderboard []CharacterLeaderboardEntryModel `json:"Leaderboard,omitempty"`
}

// GetLeaderboardAroundPlayerRequest 
type GetLeaderboardAroundPlayerRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // MaxResultsCount maximum number of entries to retrieve. Default 10, maximum 100.
    MaxResultsCount int32 `json:"MaxResultsCount,omitempty"`
    // PlayFabId playFab unique identifier of the user to center the leaderboard around. If null will center on the logged in user.
    PlayFabId string `json:"PlayFabId,omitempty"`
    // ProfileConstraints if non-null, this determines which properties of the resulting player profiles to return. For API calls from the client,
// only the allowed client profile properties for the title may be requested. These allowed properties are configured in
// the Game Manager "Client Profile Options" tab in the "Settings" section.
    ProfileConstraints *PlayerProfileViewConstraintsModel `json:"ProfileConstraints,omitempty"`
    // StatisticName statistic used to rank players for this leaderboard.
    StatisticName string `json:"StatisticName,omitempty"`
    // Version the version of the leaderboard to get.
    Version int32 `json:"Version,omitempty"`
}

// GetLeaderboardAroundPlayerResult note: When calling 'GetLeaderboardAround...' APIs, the position of the user defaults to 0 when the user does not have
// the corresponding statistic.
type GetLeaderboardAroundPlayerResultModel struct {
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
    // MaxResultsCount maximum number of entries to retrieve. Default 10, maximum 100.
    MaxResultsCount int32 `json:"MaxResultsCount,omitempty"`
    // ProfileConstraints if non-null, this determines which properties of the resulting player profiles to return. For API calls from the client,
// only the allowed client profile properties for the title may be requested. These allowed properties are configured in
// the Game Manager "Client Profile Options" tab in the "Settings" section.
    ProfileConstraints *PlayerProfileViewConstraintsModel `json:"ProfileConstraints,omitempty"`
    // StartPosition position in the leaderboard to start this listing (defaults to the first entry).
    StartPosition int32 `json:"StartPosition,omitempty"`
    // StatisticName statistic used to rank players for this leaderboard.
    StatisticName string `json:"StatisticName,omitempty"`
    // Version the version of the leaderboard to get.
    Version int32 `json:"Version,omitempty"`
}

// GetLeaderboardResult note: the user's Position is relative to the overall leaderboard.
type GetLeaderboardResultModel struct {
    // Leaderboard ordered listing of users and their positions in the requested leaderboard.
    Leaderboard []PlayerLeaderboardEntryModel `json:"Leaderboard,omitempty"`
    // NextReset the time the next scheduled reset will occur. Null if the leaderboard does not reset on a schedule.
    NextReset time.Time `json:"NextReset,omitempty"`
    // Version the version of the leaderboard returned.
    Version int32 `json:"Version,omitempty"`
}

// GetPaymentTokenRequest 
type GetPaymentTokenRequestModel struct {
    // TokenProvider the name of service to provide the payment token. Allowed Values are: xsolla
    TokenProvider string `json:"TokenProvider,omitempty"`
}

// GetPaymentTokenResult 
type GetPaymentTokenResultModel struct {
    // OrderId playFab's purchase order identifier.
    OrderId string `json:"OrderId,omitempty"`
    // ProviderToken the token from provider.
    ProviderToken string `json:"ProviderToken,omitempty"`
}

// GetPhotonAuthenticationTokenRequest 
type GetPhotonAuthenticationTokenRequestModel struct {
    // PhotonApplicationId the Photon applicationId for the game you wish to log into.
    PhotonApplicationId string `json:"PhotonApplicationId,omitempty"`
}

// GetPhotonAuthenticationTokenResult 
type GetPhotonAuthenticationTokenResultModel struct {
    // PhotonCustomAuthenticationToken the Photon authentication token for this game-session.
    PhotonCustomAuthenticationToken string `json:"PhotonCustomAuthenticationToken,omitempty"`
}

// GetPlayerCombinedInfoRequest 
type GetPlayerCombinedInfoRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // InfoRequestParameters flags for which pieces of info to return for the user.
    InfoRequestParameters* GetPlayerCombinedInfoRequestParamsModel `json:"InfoRequestParameters,omitempty"`
    // PlayFabId playFabId of the user whose data will be returned. If not filled included, we return the data for the calling player.
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

// GetPlayerCombinedInfoResult returns whatever info is requested in the response for the user. If no user is explicitly requested this defaults to the
// authenticated user. If the user is the same as the requester, PII (like email address, facebook id) is returned if
// available. Otherwise, only public information is returned. All parameters default to false.
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

// GetPlayerSegmentsRequest 
type GetPlayerSegmentsRequestModel struct {
}

// GetPlayerSegmentsResult 
type GetPlayerSegmentsResultModel struct {
    // Segments array of segments the requested player currently belongs to.
    Segments []GetSegmentResultModel `json:"Segments,omitempty"`
}

// GetPlayerStatisticsRequest 
type GetPlayerStatisticsRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // StatisticNames statistics to return (current version will be returned for each)
    StatisticNames []string `json:"StatisticNames,omitempty"`
    // StatisticNameVersions statistics to return, if StatisticNames is not set (only statistics which have a version matching that provided will be
// returned)
    StatisticNameVersions []StatisticNameVersionModel `json:"StatisticNameVersions,omitempty"`
}

// GetPlayerStatisticsResult in addition to being available for use by the title, the statistics are used for all leaderboard operations in PlayFab.
type GetPlayerStatisticsResultModel struct {
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

// GetPlayerTradesRequest 
type GetPlayerTradesRequestModel struct {
    // StatusFilter returns only trades with the given status. If null, returns all trades.
    StatusFilter TradeStatus `json:"StatusFilter,omitempty"`
}

// GetPlayerTradesResponse 
type GetPlayerTradesResponseModel struct {
    // AcceptedTrades history of trades which this player has accepted.
    AcceptedTrades []TradeInfoModel `json:"AcceptedTrades,omitempty"`
    // OpenedTrades the trades for this player which are currently available to be accepted.
    OpenedTrades []TradeInfoModel `json:"OpenedTrades,omitempty"`
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

// GetPlayFabIDsFromFacebookInstantGamesIdsResult for Facebook Instant Game identifiers which have not been linked to PlayFab accounts, null will be returned.
type GetPlayFabIDsFromFacebookInstantGamesIdsResultModel struct {
    // Data mapping of Facebook Instant Games identifiers to PlayFab identifiers.
    Data []FacebookInstantGamesPlayFabIdPairModel `json:"Data,omitempty"`
}

// GetPlayFabIDsFromGameCenterIDsRequest 
type GetPlayFabIDsFromGameCenterIDsRequestModel struct {
    // GameCenterIDs array of unique Game Center identifiers (the Player Identifier) for which the title needs to get PlayFab identifiers.
    GameCenterIDs []string `json:"GameCenterIDs,omitempty"`
}

// GetPlayFabIDsFromGameCenterIDsResult for Game Center identifiers which have not been linked to PlayFab accounts, null will be returned.
type GetPlayFabIDsFromGameCenterIDsResultModel struct {
    // Data mapping of Game Center identifiers to PlayFab identifiers.
    Data []GameCenterPlayFabIdPairModel `json:"Data,omitempty"`
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

// GetPlayFabIDsFromGoogleIDsRequest 
type GetPlayFabIDsFromGoogleIDsRequestModel struct {
    // GoogleIDs array of unique Google identifiers (Google+ user IDs) for which the title needs to get PlayFab identifiers.
    GoogleIDs []string `json:"GoogleIDs,omitempty"`
}

// GetPlayFabIDsFromGoogleIDsResult for Google identifiers which have not been linked to PlayFab accounts, null will be returned.
type GetPlayFabIDsFromGoogleIDsResultModel struct {
    // Data mapping of Google identifiers to PlayFab identifiers.
    Data []GooglePlayFabIdPairModel `json:"Data,omitempty"`
}

// GetPlayFabIDsFromKongregateIDsRequest 
type GetPlayFabIDsFromKongregateIDsRequestModel struct {
    // KongregateIDs array of unique Kongregate identifiers (Kongregate's user_id) for which the title needs to get PlayFab identifiers.
    KongregateIDs []string `json:"KongregateIDs,omitempty"`
}

// GetPlayFabIDsFromKongregateIDsResult for Kongregate identifiers which have not been linked to PlayFab accounts, null will be returned.
type GetPlayFabIDsFromKongregateIDsResultModel struct {
    // Data mapping of Kongregate identifiers to PlayFab identifiers.
    Data []KongregatePlayFabIdPairModel `json:"Data,omitempty"`
}

// GetPlayFabIDsFromNintendoSwitchDeviceIdsRequest 
type GetPlayFabIDsFromNintendoSwitchDeviceIdsRequestModel struct {
    // NintendoSwitchDeviceIds array of unique Nintendo Switch Device identifiers for which the title needs to get PlayFab identifiers.
    NintendoSwitchDeviceIds []string `json:"NintendoSwitchDeviceIds,omitempty"`
}

// GetPlayFabIDsFromNintendoSwitchDeviceIdsResult for Nintendo Switch identifiers which have not been linked to PlayFab accounts, null will be returned.
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

// GetPlayFabIDsFromTwitchIDsRequest 
type GetPlayFabIDsFromTwitchIDsRequestModel struct {
    // TwitchIds array of unique Twitch identifiers (Twitch's _id) for which the title needs to get PlayFab identifiers.
    TwitchIds []string `json:"TwitchIds,omitempty"`
}

// GetPlayFabIDsFromTwitchIDsResult for Twitch identifiers which have not been linked to PlayFab accounts, null will be returned.
type GetPlayFabIDsFromTwitchIDsResultModel struct {
    // Data mapping of Twitch identifiers to PlayFab identifiers.
    Data []TwitchPlayFabIdPairModel `json:"Data,omitempty"`
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

// GetPurchaseRequest 
type GetPurchaseRequestModel struct {
    // OrderId purchase order identifier.
    OrderId string `json:"OrderId,omitempty"`
}

// GetPurchaseResult 
type GetPurchaseResultModel struct {
    // OrderId purchase order identifier.
    OrderId string `json:"OrderId,omitempty"`
    // PaymentProvider payment provider used for transaction (If not VC)
    PaymentProvider string `json:"PaymentProvider,omitempty"`
    // PurchaseDate date and time of the purchase.
    PurchaseDate time.Time `json:"PurchaseDate,omitempty"`
    // TransactionId provider transaction ID (If not VC)
    TransactionId string `json:"TransactionId,omitempty"`
    // TransactionStatus playFab transaction status
    TransactionStatus string `json:"TransactionStatus,omitempty"`
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

// GetStoreItemsRequest a store contains an array of references to items defined in one or more catalog versions of the game, along with the
// prices for the item, in both real world and virtual currencies. These prices act as an override to any prices defined in
// the catalog. In this way, the base definitions of the items may be defined in the catalog, with all associated
// properties, while the pricing can be set for each store, as needed. This allows for subsets of goods to be defined for
// different purposes (in order to simplify showing some, but not all catalog items to users, based upon different
// characteristics), along with unique prices. Note that all prices defined in the catalog and store definitions for the
// item are considered valid, and that a compromised client can be made to send a request for an item based upon any of
// these definitions. If no price is specified in the store for an item, the price set in the catalog should be displayed
// to the user.
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
// build. If the player belongs to an experiment variant that uses title data overrides, the overrides are applied
// automatically and returned with the title data. Note that there may up to a minute delay in between updating title data
// and this API call returning the newest value.
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
    // News array of news items.
    News []TitleNewsItemModel `json:"News,omitempty"`
}

// GetTitlePublicKeyRequest an RSA CSP blob to be used to encrypt the payload of account creation requests when that API requires a signature
// header. For example if Client/LoginWithCustomId requires signature headers but the player does not have an account yet
// follow these steps: 1) Call Client/GetTitlePublicKey with one of the title's shared secrets. 2) Convert the Base64
// encoded CSP blob to a byte array and create an RSA signing object. 3) Encrypt the UTF8 encoded JSON body of the
// registration request and place the Base64 encoded result into the EncryptedRequest and with the TitleId field, all other
// fields can be left empty when performing the API request. 4) Client receives authentication token as normal. Future
// requests to LoginWithCustomId will require the X-PlayFab-Signature header.
type GetTitlePublicKeyRequestModel struct {
    // TitleId unique identifier for the title, found in the Settings > Game Properties section of the PlayFab developer site when a
// title has been selected.
    TitleId string `json:"TitleId,omitempty"`
    // TitleSharedSecret the shared secret key for this title
    TitleSharedSecret string `json:"TitleSharedSecret,omitempty"`
}

// GetTitlePublicKeyResult 
type GetTitlePublicKeyResultModel struct {
    // RSAPublicKey base64 encoded RSA CSP byte array blob containing the title's public RSA key
    RSAPublicKey string `json:"RSAPublicKey,omitempty"`
}

// GetTradeStatusRequest 
type GetTradeStatusRequestModel struct {
    // OfferingPlayerId player who opened trade.
    OfferingPlayerId string `json:"OfferingPlayerId,omitempty"`
    // TradeId trade identifier as returned by OpenTradeOffer.
    TradeId string `json:"TradeId,omitempty"`
}

// GetTradeStatusResponse 
type GetTradeStatusResponseModel struct {
    // Trade information about the requested trade.
    Trade *TradeInfoModel `json:"Trade,omitempty"`
}

// GetUserDataRequest data is stored as JSON key-value pairs. Every time the data is updated via any source, the version counter is
// incremented. If the Version parameter is provided, then this call will only return data if the current version on the
// system is greater than the value provided. If the Keys parameter is provided, the data object returned will only contain
// the data specific to the indicated Keys. Otherwise, the full set of custom user data will be returned.
type GetUserDataRequestModel struct {
    // IfChangedFromDataVersion the version that currently exists according to the caller. The call will return the data for all of the keys if the
// version in the system is greater than this.
    IfChangedFromDataVersion uint32 `json:"IfChangedFromDataVersion,omitempty"`
    // Keys list of unique keys to load from.
    Keys []string `json:"Keys,omitempty"`
    // PlayFabId unique PlayFab identifier of the user to load data for. Optional, defaults to yourself if not set. When specified to a
// PlayFab id of another player, then this will only return public keys for that account.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// GetUserDataResult 
type GetUserDataResultModel struct {
    // Data user specific data for this title.
    Data map[string]UserDataRecordModel `json:"Data,omitempty"`
    // DataVersion indicates the current version of the data that has been set. This is incremented with every set call for that type of
// data (read-only, internal, etc). This version can be provided in Get calls to find updated data.
    DataVersion uint32 `json:"DataVersion,omitempty"`
}

// GetUserInventoryRequest 
type GetUserInventoryRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
}

// GetUserInventoryResult all items currently in the user inventory will be returned, irrespective of how they were acquired (via purchasing,
// grants, coupons, etc.). Items that are expired, fully consumed, or are no longer valid are not considered to be in the
// user's current inventory, and so will not be not included.
type GetUserInventoryResultModel struct {
    // Inventory array of inventory items belonging to the user.
    Inventory []ItemInstanceModel `json:"Inventory,omitempty"`
    // VirtualCurrency array of virtual currency balance(s) belonging to the user.
    VirtualCurrency map[string]int32 `json:"VirtualCurrency,omitempty"`
    // VirtualCurrencyRechargeTimes array of remaining times and timestamps for virtual currencies.
    VirtualCurrencyRechargeTimes map[string]VirtualCurrencyRechargeTimeModel `json:"VirtualCurrencyRechargeTimes,omitempty"`
}

// GetWindowsHelloChallengeRequest requires the SHA256 hash of the user's public key.
type GetWindowsHelloChallengeRequestModel struct {
    // PublicKeyHint sHA256 hash of the PublicKey generated by Windows Hello.
    PublicKeyHint string `json:"PublicKeyHint,omitempty"`
    // TitleId unique identifier for the title, found in the Settings > Game Properties section of the PlayFab developer site when a
// title has been selected.
    TitleId string `json:"TitleId,omitempty"`
}

// GetWindowsHelloChallengeResponse 
type GetWindowsHelloChallengeResponseModel struct {
    // Challenge server generated challenge to be signed by the user.
    Challenge string `json:"Challenge,omitempty"`
}

// GooglePlayFabIdPair 
type GooglePlayFabIdPairModel struct {
    // GoogleId unique Google identifier for a user.
    GoogleId string `json:"GoogleId,omitempty"`
    // PlayFabId unique PlayFab identifier for a user, or null if no PlayFab account is linked to the Google identifier.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// GrantCharacterToUserRequest grants a character to the user of the type specified by the item ID. The user must already have an instance of this item
// in their inventory in order to allow character creation. This item can come from a purchase or grant, which must be done
// before calling to create the character.
type GrantCharacterToUserRequestModel struct {
    // CatalogVersion catalog version from which items are to be granted.
    CatalogVersion string `json:"CatalogVersion,omitempty"`
    // CharacterName non-unique display name of the character being granted (1-40 characters in length).
    CharacterName string `json:"CharacterName,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // ItemId catalog item identifier of the item in the user's inventory that corresponds to the character in the catalog to be
// created.
    ItemId string `json:"ItemId,omitempty"`
}

// GrantCharacterToUserResult 
type GrantCharacterToUserResultModel struct {
    // CharacterId unique identifier tagged to this character.
    CharacterId string `json:"CharacterId,omitempty"`
    // CharacterType type of character that was created.
    CharacterType string `json:"CharacterType,omitempty"`
    // Result indicates whether this character was created successfully.
    Result bool `json:"Result"`
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

// ItemPurchaseRequest 
type ItemPurchaseRequestModel struct {
    // Annotation title-specific text concerning this purchase.
    Annotation string `json:"Annotation,omitempty"`
    // ItemId unique ItemId of the item to purchase.
    ItemId string `json:"ItemId,omitempty"`
    // Quantity how many of this item to purchase. Min 1, maximum 25.
    Quantity uint32 `json:"Quantity,omitempty"`
    // UpgradeFromItems items to be upgraded as a result of this purchase (upgraded items are hidden, as they are "replaced" by the new items).
    UpgradeFromItems []string `json:"UpgradeFromItems,omitempty"`
}

// KongregatePlayFabIdPair 
type KongregatePlayFabIdPairModel struct {
    // KongregateId unique Kongregate identifier for a user.
    KongregateId string `json:"KongregateId,omitempty"`
    // PlayFabId unique PlayFab identifier for a user, or null if no PlayFab account is linked to the Kongregate identifier.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// LinkAndroidDeviceIDRequest 
type LinkAndroidDeviceIDRequestModel struct {
    // AndroidDevice specific model of the user's device.
    AndroidDevice string `json:"AndroidDevice,omitempty"`
    // AndroidDeviceId android device identifier for the user's device.
    AndroidDeviceId string `json:"AndroidDeviceId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // ForceLink if another user is already linked to the device, unlink the other user and re-link.
    ForceLink bool `json:"ForceLink"`
    // OS specific Operating System version for the user's device.
    OS string `json:"OS,omitempty"`
}

// LinkAndroidDeviceIDResult 
type LinkAndroidDeviceIDResultModel struct {
}

// LinkAppleRequest 
type LinkAppleRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // ForceLink if another user is already linked to a specific Apple account, unlink the other user and re-link.
    ForceLink bool `json:"ForceLink"`
    // IdentityToken the JSON Web token (JWT) returned by Apple after login. Represented as the identityToken field in the authorization
// credential payload. Used to validate the request and find the user ID (Apple subject) to link with.
    IdentityToken string `json:"IdentityToken,omitempty"`
}

// LinkCustomIDRequest 
type LinkCustomIDRequestModel struct {
    // CustomId custom unique identifier for the user, generated by the title.
    CustomId string `json:"CustomId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // ForceLink if another user is already linked to the custom ID, unlink the other user and re-link.
    ForceLink bool `json:"ForceLink"`
}

// LinkCustomIDResult 
type LinkCustomIDResultModel struct {
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

// LinkFacebookAccountRequest facebook sign-in is accomplished using the Facebook User Access Token. More information on the Token can be found in the
// Facebook developer documentation (https://developers.facebook.com/docs/facebook-login/access-tokens/). In Unity, for
// example, the Token is available as AccessToken in the Facebook SDK ScriptableObject FB. Note that titles should never
// re-use the same Facebook applications between PlayFab Title IDs, as Facebook provides unique user IDs per application
// and doing so can result in issues with the Facebook ID for the user in their PlayFab account information. If you must
// re-use an application in a new PlayFab Title ID, please be sure to first unlink all accounts from Facebook, or delete
// all users in the first Title ID.
type LinkFacebookAccountRequestModel struct {
    // AccessToken unique identifier from Facebook for the user.
    AccessToken string `json:"AccessToken,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // ForceLink if another user is already linked to the account, unlink the other user and re-link.
    ForceLink bool `json:"ForceLink"`
}

// LinkFacebookAccountResult 
type LinkFacebookAccountResultModel struct {
}

// LinkFacebookInstantGamesIdRequest 
type LinkFacebookInstantGamesIdRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // FacebookInstantGamesSignature facebook Instant Games signature for the user.
    FacebookInstantGamesSignature string `json:"FacebookInstantGamesSignature,omitempty"`
    // ForceLink if another user is already linked to the Facebook Instant Games ID, unlink the other user and re-link.
    ForceLink bool `json:"ForceLink"`
}

// LinkFacebookInstantGamesIdResult 
type LinkFacebookInstantGamesIdResultModel struct {
}

// LinkGameCenterAccountRequest 
type LinkGameCenterAccountRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // ForceLink if another user is already linked to the account, unlink the other user and re-link.
    ForceLink bool `json:"ForceLink"`
    // GameCenterId game Center identifier for the player account to be linked.
    GameCenterId string `json:"GameCenterId,omitempty"`
    // PublicKeyUrl the URL for the public encryption key that will be used to verify the signature.
    PublicKeyUrl string `json:"PublicKeyUrl,omitempty"`
    // Salt a random value used to compute the hash and keep it randomized.
    Salt string `json:"Salt,omitempty"`
    // Signature the verification signature of the authentication payload.
    Signature string `json:"Signature,omitempty"`
    // Timestamp the integer representation of date and time that the signature was created on. PlayFab will reject authentication
// signatures not within 10 minutes of the server's current time.
    Timestamp string `json:"Timestamp,omitempty"`
}

// LinkGameCenterAccountResult 
type LinkGameCenterAccountResultModel struct {
}

// LinkGoogleAccountRequest google sign-in is accomplished by obtaining a Google OAuth 2.0 credential using the Google sign-in for Android APIs on
// the device and passing it to this API.
type LinkGoogleAccountRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // ForceLink if another user is already linked to the account, unlink the other user and re-link.
    ForceLink bool `json:"ForceLink"`
    // ServerAuthCode server authentication code obtained on the client by calling getServerAuthCode()
// (https://developers.google.com/identity/sign-in/android/offline-access) from Google Play for the user.
    ServerAuthCode string `json:"ServerAuthCode,omitempty"`
}

// LinkGoogleAccountResult 
type LinkGoogleAccountResultModel struct {
}

// LinkIOSDeviceIDRequest 
type LinkIOSDeviceIDRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // DeviceId vendor-specific iOS identifier for the user's device.
    DeviceId string `json:"DeviceId,omitempty"`
    // DeviceModel specific model of the user's device.
    DeviceModel string `json:"DeviceModel,omitempty"`
    // ForceLink if another user is already linked to the device, unlink the other user and re-link.
    ForceLink bool `json:"ForceLink"`
    // OS specific Operating System version for the user's device.
    OS string `json:"OS,omitempty"`
}

// LinkIOSDeviceIDResult 
type LinkIOSDeviceIDResultModel struct {
}

// LinkKongregateAccountRequest 
type LinkKongregateAccountRequestModel struct {
    // AuthTicket valid session auth ticket issued by Kongregate
    AuthTicket string `json:"AuthTicket,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // ForceLink if another user is already linked to the account, unlink the other user and re-link.
    ForceLink bool `json:"ForceLink"`
    // KongregateId numeric user ID assigned by Kongregate
    KongregateId string `json:"KongregateId,omitempty"`
}

// LinkKongregateAccountResult 
type LinkKongregateAccountResultModel struct {
}

// LinkNintendoServiceAccountRequest 
type LinkNintendoServiceAccountRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // ForceLink if another user is already linked to a specific Nintendo Switch account, unlink the other user and re-link.
    ForceLink bool `json:"ForceLink"`
    // IdentityToken the JSON Web token (JWT) returned by Nintendo after login. Used to validate the request and find the user ID (Nintendo
// Switch subject) to link with.
    IdentityToken string `json:"IdentityToken,omitempty"`
}

// LinkNintendoSwitchDeviceIdRequest 
type LinkNintendoSwitchDeviceIdRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // ForceLink if another user is already linked to the Nintendo Switch Device ID, unlink the other user and re-link.
    ForceLink bool `json:"ForceLink"`
    // NintendoSwitchDeviceId nintendo Switch unique identifier for the user's device.
    NintendoSwitchDeviceId string `json:"NintendoSwitchDeviceId,omitempty"`
}

// LinkNintendoSwitchDeviceIdResult 
type LinkNintendoSwitchDeviceIdResultModel struct {
}

// LinkOpenIdConnectRequest 
type LinkOpenIdConnectRequestModel struct {
    // ConnectionId a name that identifies which configured OpenID Connect provider relationship to use. Maximum 100 characters.
    ConnectionId string `json:"ConnectionId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // ForceLink if another user is already linked to a specific OpenId Connect user, unlink the other user and re-link.
    ForceLink bool `json:"ForceLink"`
    // IdToken the JSON Web token (JWT) returned by the identity provider after login. Represented as the id_token field in the
// identity provider's response. Used to validate the request and find the user ID (OpenID Connect subject) to link with.
    IdToken string `json:"IdToken,omitempty"`
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
    // RedirectUri redirect URI supplied to PSN when requesting an auth code
    RedirectUri string `json:"RedirectUri,omitempty"`
}

// LinkPSNAccountResult 
type LinkPSNAccountResultModel struct {
}

// LinkSteamAccountRequest steam authentication is accomplished with the Steam Session Ticket. More information on the Ticket can be found in the
// Steamworks SDK, here: https://partner.steamgames.com/documentation/auth (requires sign-in). NOTE: For Steam
// authentication to work, the title must be configured with the Steam Application ID and Publisher Key in the PlayFab Game
// Manager (under Properties). Information on creating a Publisher Key (referred to as the Secret Key in PlayFab) for your
// title can be found here: https://partner.steamgames.com/documentation/webapi#publisherkey.
type LinkSteamAccountRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // ForceLink if another user is already linked to the account, unlink the other user and re-link.
    ForceLink bool `json:"ForceLink"`
    // SteamTicket authentication token for the user, returned as a byte array from Steam, and converted to a string (for example, the byte
// 0x08 should become "08").
    SteamTicket string `json:"SteamTicket,omitempty"`
}

// LinkSteamAccountResult 
type LinkSteamAccountResultModel struct {
}

// LinkTwitchAccountRequest 
type LinkTwitchAccountRequestModel struct {
    // AccessToken valid token issued by Twitch
    AccessToken string `json:"AccessToken,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // ForceLink if another user is already linked to the account, unlink the other user and re-link.
    ForceLink bool `json:"ForceLink"`
}

// LinkTwitchAccountResult 
type LinkTwitchAccountResultModel struct {
}

// LinkWindowsHelloAccountRequest publicKey must be generated using the Windows Hello Passport service.
type LinkWindowsHelloAccountRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // DeviceName device name.
    DeviceName string `json:"DeviceName,omitempty"`
    // ForceLink if another user is already linked to the account, unlink the other user and re-link.
    ForceLink bool `json:"ForceLink"`
    // PublicKey publicKey generated by Windows Hello.
    PublicKey string `json:"PublicKey,omitempty"`
    // UserName player's user named used by Windows Hello.
    UserName string `json:"UserName,omitempty"`
}

// LinkWindowsHelloAccountResponse 
type LinkWindowsHelloAccountResponseModel struct {
}

// LinkXboxAccountRequest 
type LinkXboxAccountRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // ForceLink if another user is already linked to the account, unlink the other user and re-link.
    ForceLink bool `json:"ForceLink"`
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
// LoginResult 
type LoginResultModel struct {
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

// LoginWithAndroidDeviceIDRequest on Android devices, the recommendation is to use the Settings.Secure.ANDROID_ID as the AndroidDeviceId, as described in
// this blog post (http://android-developers.blogspot.com/2011/03/identifying-app-installations.html). More information on
// this identifier can be found in the Android documentation
// (http://developer.android.com/reference/android/provider/Settings.Secure.html). If this is the first time a user has
// signed in with the Android device and CreateAccount is set to true, a new PlayFab account will be created and linked to
// the Android device ID. In this case, no email or username will be associated with the PlayFab account. Otherwise, if no
// PlayFab account is linked to the Android device, an error indicating this will be returned, so that the title can guide
// the user through creation of a PlayFab account. Please note that while multiple devices of this type can be linked to a
// single user account, only the one most recently used to login (or most recently linked) will be reflected in the user's
// account information. We will be updating to show all linked devices in a future release.
type LoginWithAndroidDeviceIDRequestModel struct {
    // AndroidDevice specific model of the user's device.
    AndroidDevice string `json:"AndroidDevice,omitempty"`
    // AndroidDeviceId android device identifier for the user's device.
    AndroidDeviceId string `json:"AndroidDeviceId,omitempty"`
    // CreateAccount automatically create a PlayFab account if one is not currently linked to this ID.
    CreateAccount bool `json:"CreateAccount"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // EncryptedRequest base64 encoded body that is encrypted with the Title's public RSA key (Enterprise Only).
    EncryptedRequest string `json:"EncryptedRequest,omitempty"`
    // InfoRequestParameters flags for which pieces of info to return for the user.
    InfoRequestParameters *GetPlayerCombinedInfoRequestParamsModel `json:"InfoRequestParameters,omitempty"`
    // OS specific Operating System version for the user's device.
    OS string `json:"OS,omitempty"`
    // PlayerSecret player secret that is used to verify API request signatures (Enterprise Only).
    PlayerSecret string `json:"PlayerSecret,omitempty"`
    // TitleId unique identifier for the title, found in the Settings > Game Properties section of the PlayFab developer site when a
// title has been selected.
    TitleId string `json:"TitleId,omitempty"`
}

// LoginWithAppleRequest 
type LoginWithAppleRequestModel struct {
    // CreateAccount automatically create a PlayFab account if one is not currently linked to this ID.
    CreateAccount bool `json:"CreateAccount"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // EncryptedRequest base64 encoded body that is encrypted with the Title's public RSA key (Enterprise Only).
    EncryptedRequest string `json:"EncryptedRequest,omitempty"`
    // IdentityToken the JSON Web token (JWT) returned by Apple after login. Represented as the identityToken field in the authorization
// credential payload.
    IdentityToken string `json:"IdentityToken,omitempty"`
    // InfoRequestParameters flags for which pieces of info to return for the user.
    InfoRequestParameters *GetPlayerCombinedInfoRequestParamsModel `json:"InfoRequestParameters,omitempty"`
    // PlayerSecret player secret that is used to verify API request signatures (Enterprise Only).
    PlayerSecret string `json:"PlayerSecret,omitempty"`
    // TitleId unique identifier for the title, found in the Settings > Game Properties section of the PlayFab developer site when a
// title has been selected.
    TitleId string `json:"TitleId,omitempty"`
}

// LoginWithCustomIDRequest it is highly recommended that developers ensure that it is extremely unlikely that a customer could generate an ID which
// is already in use by another customer. If this is the first time a user has signed in with the Custom ID and
// CreateAccount is set to true, a new PlayFab account will be created and linked to the Custom ID. In this case, no email
// or username will be associated with the PlayFab account. Otherwise, if no PlayFab account is linked to the Custom ID, an
// error indicating this will be returned, so that the title can guide the user through creation of a PlayFab account.
type LoginWithCustomIDRequestModel struct {
    // CreateAccount automatically create a PlayFab account if one is not currently linked to this ID.
    CreateAccount bool `json:"CreateAccount"`
    // CustomId custom unique identifier for the user, generated by the title.
    CustomId string `json:"CustomId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // EncryptedRequest base64 encoded body that is encrypted with the Title's public RSA key (Enterprise Only).
    EncryptedRequest string `json:"EncryptedRequest,omitempty"`
    // InfoRequestParameters flags for which pieces of info to return for the user.
    InfoRequestParameters *GetPlayerCombinedInfoRequestParamsModel `json:"InfoRequestParameters,omitempty"`
    // PlayerSecret player secret that is used to verify API request signatures (Enterprise Only).
    PlayerSecret string `json:"PlayerSecret,omitempty"`
    // TitleId unique identifier for the title, found in the Settings > Game Properties section of the PlayFab developer site when a
// title has been selected.
    TitleId string `json:"TitleId,omitempty"`
}

// LoginWithEmailAddressRequest email address and password lengths are provided for information purposes. The server will validate that data passed in
// conforms to the field definition and report errors appropriately. It is recommended that developers not perform this
// validation locally, so that future updates do not require client updates.
type LoginWithEmailAddressRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Email email address for the account.
    Email string `json:"Email,omitempty"`
    // InfoRequestParameters flags for which pieces of info to return for the user.
    InfoRequestParameters *GetPlayerCombinedInfoRequestParamsModel `json:"InfoRequestParameters,omitempty"`
    // Password password for the PlayFab account (6-100 characters)
    Password string `json:"Password,omitempty"`
    // TitleId unique identifier for the title, found in the Settings > Game Properties section of the PlayFab developer site when a
// title has been selected.
    TitleId string `json:"TitleId,omitempty"`
}

// LoginWithFacebookInstantGamesIdRequest 
type LoginWithFacebookInstantGamesIdRequestModel struct {
    // CreateAccount automatically create a PlayFab account if one is not currently linked to this ID.
    CreateAccount bool `json:"CreateAccount"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // EncryptedRequest base64 encoded body that is encrypted with the Title's public RSA key (Enterprise Only).
    EncryptedRequest string `json:"EncryptedRequest,omitempty"`
    // FacebookInstantGamesSignature facebook Instant Games signature for the user.
    FacebookInstantGamesSignature string `json:"FacebookInstantGamesSignature,omitempty"`
    // InfoRequestParameters flags for which pieces of info to return for the user.
    InfoRequestParameters *GetPlayerCombinedInfoRequestParamsModel `json:"InfoRequestParameters,omitempty"`
    // PlayerSecret player secret that is used to verify API request signatures (Enterprise Only).
    PlayerSecret string `json:"PlayerSecret,omitempty"`
    // TitleId unique identifier for the title, found in the Settings > Game Properties section of the PlayFab developer site when a
// title has been selected.
    TitleId string `json:"TitleId,omitempty"`
}

// LoginWithFacebookRequest facebook sign-in is accomplished using the Facebook User Access Token. More information on the Token can be found in the
// Facebook developer documentation (https://developers.facebook.com/docs/facebook-login/access-tokens/). In Unity, for
// example, the Token is available as AccessToken in the Facebook SDK ScriptableObject FB. If this is the first time a user
// has signed in with the Facebook account and CreateAccount is set to true, a new PlayFab account will be created and
// linked to the provided account's Facebook ID. In this case, no email or username will be associated with the PlayFab
// account. Otherwise, if no PlayFab account is linked to the Facebook account, an error indicating this will be returned,
// so that the title can guide the user through creation of a PlayFab account. Note that titles should never re-use the
// same Facebook applications between PlayFab Title IDs, as Facebook provides unique user IDs per application and doing so
// can result in issues with the Facebook ID for the user in their PlayFab account information. If you must re-use an
// application in a new PlayFab Title ID, please be sure to first unlink all accounts from Facebook, or delete all users in
// the first Title ID.
type LoginWithFacebookRequestModel struct {
    // AccessToken unique identifier from Facebook for the user.
    AccessToken string `json:"AccessToken,omitempty"`
    // CreateAccount automatically create a PlayFab account if one is not currently linked to this ID.
    CreateAccount bool `json:"CreateAccount"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // EncryptedRequest base64 encoded body that is encrypted with the Title's public RSA key (Enterprise Only).
    EncryptedRequest string `json:"EncryptedRequest,omitempty"`
    // InfoRequestParameters flags for which pieces of info to return for the user.
    InfoRequestParameters *GetPlayerCombinedInfoRequestParamsModel `json:"InfoRequestParameters,omitempty"`
    // PlayerSecret player secret that is used to verify API request signatures (Enterprise Only).
    PlayerSecret string `json:"PlayerSecret,omitempty"`
    // TitleId unique identifier for the title, found in the Settings > Game Properties section of the PlayFab developer site when a
// title has been selected.
    TitleId string `json:"TitleId,omitempty"`
}

// LoginWithGameCenterRequest the Game Center player identifier
// (https://developer.apple.com/library/ios/documentation/Accounts/Reference/ACAccountClassRef/index.html#//apple_ref/occ/instp/ACAccount/identifier)
// is a generated string which is stored on the local device. As with device identifiers, care must be taken to never
// expose a player's Game Center identifier to end users, as that could result in a user's account being compromised. If
// this is the first time a user has signed in with Game Center and CreateAccount is set to true, a new PlayFab account
// will be created and linked to the Game Center identifier. In this case, no email or username will be associated with the
// PlayFab account. Otherwise, if no PlayFab account is linked to the Game Center account, an error indicating this will be
// returned, so that the title can guide the user through creation of a PlayFab account.
type LoginWithGameCenterRequestModel struct {
    // CreateAccount automatically create a PlayFab account if one is not currently linked to this ID.
    CreateAccount bool `json:"CreateAccount"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // EncryptedRequest base64 encoded body that is encrypted with the Title's public RSA key (Enterprise Only).
    EncryptedRequest string `json:"EncryptedRequest,omitempty"`
    // InfoRequestParameters flags for which pieces of info to return for the user.
    InfoRequestParameters *GetPlayerCombinedInfoRequestParamsModel `json:"InfoRequestParameters,omitempty"`
    // PlayerId unique Game Center player id.
    PlayerId string `json:"PlayerId,omitempty"`
    // PlayerSecret player secret that is used to verify API request signatures (Enterprise Only).
    PlayerSecret string `json:"PlayerSecret,omitempty"`
    // PublicKeyUrl the URL for the public encryption key that will be used to verify the signature.
    PublicKeyUrl string `json:"PublicKeyUrl,omitempty"`
    // Salt a random value used to compute the hash and keep it randomized.
    Salt string `json:"Salt,omitempty"`
    // Signature the verification signature of the authentication payload.
    Signature string `json:"Signature,omitempty"`
    // Timestamp the integer representation of date and time that the signature was created on. PlayFab will reject authentication
// signatures not within 10 minutes of the server's current time.
    Timestamp string `json:"Timestamp,omitempty"`
    // TitleId unique identifier for the title, found in the Settings > Game Properties section of the PlayFab developer site when a
// title has been selected.
    TitleId string `json:"TitleId,omitempty"`
}

// LoginWithGoogleAccountRequest google sign-in is accomplished by obtaining a Google OAuth 2.0 credential using the Google sign-in for Android APIs on
// the device and passing it to this API. If this is the first time a user has signed in with the Google account and
// CreateAccount is set to true, a new PlayFab account will be created and linked to the Google account. Otherwise, if no
// PlayFab account is linked to the Google account, an error indicating this will be returned, so that the title can guide
// the user through creation of a PlayFab account. The current (recommended) method for obtaining a Google account
// credential in an Android application is to call GoogleSignInAccount.getServerAuthCode() and send the auth code as the
// ServerAuthCode parameter of this API. Before doing this, you must create an OAuth 2.0 web application client ID in the
// Google API Console and configure its client ID and secret in the PlayFab Game Manager Google Add-on for your title. This
// method does not require prompting of the user for additional Google account permissions, resulting in a user experience
// with the least possible friction. For more information about obtaining the server auth code, see
// https://developers.google.com/identity/sign-in/android/offline-access. The previous (deprecated) method was to obtain an
// OAuth access token by calling GetAccessToken() on the client and passing it as the AccessToken parameter to this API.
// for the with the Google OAuth 2.0 Access Token. More information on this change can be found in the Google developer
// documentation (https://android-developers.googleblog.com/2016/01/play-games-permissions-are-changing-in.html).
type LoginWithGoogleAccountRequestModel struct {
    // CreateAccount automatically create a PlayFab account if one is not currently linked to this ID.
    CreateAccount bool `json:"CreateAccount"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // EncryptedRequest base64 encoded body that is encrypted with the Title's public RSA key (Enterprise Only).
    EncryptedRequest string `json:"EncryptedRequest,omitempty"`
    // InfoRequestParameters flags for which pieces of info to return for the user.
    InfoRequestParameters *GetPlayerCombinedInfoRequestParamsModel `json:"InfoRequestParameters,omitempty"`
    // PlayerSecret player secret that is used to verify API request signatures (Enterprise Only).
    PlayerSecret string `json:"PlayerSecret,omitempty"`
    // ServerAuthCode oAuth 2.0 server authentication code obtained on the client by calling the getServerAuthCode()
// (https://developers.google.com/identity/sign-in/android/offline-access) Google client API.
    ServerAuthCode string `json:"ServerAuthCode,omitempty"`
    // TitleId unique identifier for the title, found in the Settings > Game Properties section of the PlayFab developer site when a
// title has been selected.
    TitleId string `json:"TitleId,omitempty"`
}

// LoginWithIOSDeviceIDRequest on iOS devices, the identifierForVendor
// (https://developer.apple.com/library/ios/documentation/UIKit/Reference/UIDevice_Class/index.html#//apple_ref/occ/instp/UIDevice/identifierForVendor)
// must be used as the DeviceId, as the UIDevice uniqueIdentifier has been deprecated as of iOS 5, and use of the
// advertisingIdentifier for this purpose will result in failure of Apple's certification process. If this is the first
// time a user has signed in with the iOS device and CreateAccount is set to true, a new PlayFab account will be created
// and linked to the vendor-specific iOS device ID. In this case, no email or username will be associated with the PlayFab
// account. Otherwise, if no PlayFab account is linked to the iOS device, an error indicating this will be returned, so
// that the title can guide the user through creation of a PlayFab account. Please note that while multiple devices of this
// type can be linked to a single user account, only the one most recently used to login (or most recently linked) will be
// reflected in the user's account information. We will be updating to show all linked devices in a future release.
type LoginWithIOSDeviceIDRequestModel struct {
    // CreateAccount automatically create a PlayFab account if one is not currently linked to this ID.
    CreateAccount bool `json:"CreateAccount"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // DeviceId vendor-specific iOS identifier for the user's device.
    DeviceId string `json:"DeviceId,omitempty"`
    // DeviceModel specific model of the user's device.
    DeviceModel string `json:"DeviceModel,omitempty"`
    // EncryptedRequest base64 encoded body that is encrypted with the Title's public RSA key (Enterprise Only).
    EncryptedRequest string `json:"EncryptedRequest,omitempty"`
    // InfoRequestParameters flags for which pieces of info to return for the user.
    InfoRequestParameters *GetPlayerCombinedInfoRequestParamsModel `json:"InfoRequestParameters,omitempty"`
    // OS specific Operating System version for the user's device.
    OS string `json:"OS,omitempty"`
    // PlayerSecret player secret that is used to verify API request signatures (Enterprise Only).
    PlayerSecret string `json:"PlayerSecret,omitempty"`
    // TitleId unique identifier for the title, found in the Settings > Game Properties section of the PlayFab developer site when a
// title has been selected.
    TitleId string `json:"TitleId,omitempty"`
}

// LoginWithKongregateRequest more details regarding Kongregate and their game authentication system can be found at
// http://developers.kongregate.com/docs/virtual-goods/authentication. Developers must provide the Kongregate user ID and
// auth token that are generated using the Kongregate client library. PlayFab will combine these identifiers with the
// title's unique Kongregate app ID to log the player into the Kongregate system. If CreateAccount is set to true and there
// is not already a user matched to this Kongregate ID, then PlayFab will create a new account for this user and link the
// ID. In this case, no email or username will be associated with the PlayFab account. If there is already a different
// PlayFab user linked with this account, then an error will be returned.
type LoginWithKongregateRequestModel struct {
    // AuthTicket token issued by Kongregate's client API for the user.
    AuthTicket string `json:"AuthTicket,omitempty"`
    // CreateAccount automatically create a PlayFab account if one is not currently linked to this ID.
    CreateAccount bool `json:"CreateAccount"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // EncryptedRequest base64 encoded body that is encrypted with the Title's public RSA key (Enterprise Only).
    EncryptedRequest string `json:"EncryptedRequest,omitempty"`
    // InfoRequestParameters flags for which pieces of info to return for the user.
    InfoRequestParameters *GetPlayerCombinedInfoRequestParamsModel `json:"InfoRequestParameters,omitempty"`
    // KongregateId numeric user ID assigned by Kongregate
    KongregateId string `json:"KongregateId,omitempty"`
    // PlayerSecret player secret that is used to verify API request signatures (Enterprise Only).
    PlayerSecret string `json:"PlayerSecret,omitempty"`
    // TitleId unique identifier for the title, found in the Settings > Game Properties section of the PlayFab developer site when a
// title has been selected.
    TitleId string `json:"TitleId,omitempty"`
}

// LoginWithNintendoServiceAccountRequest 
type LoginWithNintendoServiceAccountRequestModel struct {
    // CreateAccount automatically create a PlayFab account if one is not currently linked to this ID.
    CreateAccount bool `json:"CreateAccount"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // EncryptedRequest base64 encoded body that is encrypted with the Title's public RSA key (Enterprise Only).
    EncryptedRequest string `json:"EncryptedRequest,omitempty"`
    // IdentityToken the JSON Web token (JWT) returned by Nintendo after login.
    IdentityToken string `json:"IdentityToken,omitempty"`
    // InfoRequestParameters flags for which pieces of info to return for the user.
    InfoRequestParameters *GetPlayerCombinedInfoRequestParamsModel `json:"InfoRequestParameters,omitempty"`
    // PlayerSecret player secret that is used to verify API request signatures (Enterprise Only).
    PlayerSecret string `json:"PlayerSecret,omitempty"`
    // TitleId unique identifier for the title, found in the Settings > Game Properties section of the PlayFab developer site when a
// title has been selected.
    TitleId string `json:"TitleId,omitempty"`
}

// LoginWithNintendoSwitchDeviceIdRequest 
type LoginWithNintendoSwitchDeviceIdRequestModel struct {
    // CreateAccount automatically create a PlayFab account if one is not currently linked to this ID.
    CreateAccount bool `json:"CreateAccount"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // EncryptedRequest base64 encoded body that is encrypted with the Title's public RSA key (Enterprise Only).
    EncryptedRequest string `json:"EncryptedRequest,omitempty"`
    // InfoRequestParameters flags for which pieces of info to return for the user.
    InfoRequestParameters *GetPlayerCombinedInfoRequestParamsModel `json:"InfoRequestParameters,omitempty"`
    // NintendoSwitchDeviceId nintendo Switch unique identifier for the user's device.
    NintendoSwitchDeviceId string `json:"NintendoSwitchDeviceId,omitempty"`
    // PlayerSecret player secret that is used to verify API request signatures (Enterprise Only).
    PlayerSecret string `json:"PlayerSecret,omitempty"`
    // TitleId unique identifier for the title, found in the Settings > Game Properties section of the PlayFab developer site when a
// title has been selected.
    TitleId string `json:"TitleId,omitempty"`
}

// LoginWithOpenIdConnectRequest 
type LoginWithOpenIdConnectRequestModel struct {
    // ConnectionId a name that identifies which configured OpenID Connect provider relationship to use. Maximum 100 characters.
    ConnectionId string `json:"ConnectionId,omitempty"`
    // CreateAccount automatically create a PlayFab account if one is not currently linked to this ID.
    CreateAccount bool `json:"CreateAccount"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // EncryptedRequest base64 encoded body that is encrypted with the Title's public RSA key (Enterprise Only).
    EncryptedRequest string `json:"EncryptedRequest,omitempty"`
    // IdToken the JSON Web token (JWT) returned by the identity provider after login. Represented as the id_token field in the
// identity provider's response.
    IdToken string `json:"IdToken,omitempty"`
    // InfoRequestParameters flags for which pieces of info to return for the user.
    InfoRequestParameters *GetPlayerCombinedInfoRequestParamsModel `json:"InfoRequestParameters,omitempty"`
    // PlayerSecret player secret that is used to verify API request signatures (Enterprise Only).
    PlayerSecret string `json:"PlayerSecret,omitempty"`
    // TitleId unique identifier for the title, found in the Settings > Game Properties section of the PlayFab developer site when a
// title has been selected.
    TitleId string `json:"TitleId,omitempty"`
}

// LoginWithPlayFabRequest username and password lengths are provided for information purposes. The server will validate that data passed in
// conforms to the field definition and report errors appropriately. It is recommended that developers not perform this
// validation locally, so that future updates to the username or password do not require client updates.
type LoginWithPlayFabRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // InfoRequestParameters flags for which pieces of info to return for the user.
    InfoRequestParameters *GetPlayerCombinedInfoRequestParamsModel `json:"InfoRequestParameters,omitempty"`
    // Password password for the PlayFab account (6-100 characters)
    Password string `json:"Password,omitempty"`
    // TitleId unique identifier for the title, found in the Settings > Game Properties section of the PlayFab developer site when a
// title has been selected.
    TitleId string `json:"TitleId,omitempty"`
    // Username playFab username for the account.
    Username string `json:"Username,omitempty"`
}

// LoginWithPSNRequest if this is the first time a user has signed in with the PlayStation Network account and CreateAccount is set to true, a
// new PlayFab account will be created and linked to the PSN account. In this case, no email or username will be associated
// with the PlayFab account. Otherwise, if no PlayFab account is linked to the PSN account, an error indicating this will
// be returned, so that the title can guide the user through creation of a PlayFab account.
type LoginWithPSNRequestModel struct {
    // AuthCode auth code provided by the PSN OAuth provider.
    AuthCode string `json:"AuthCode,omitempty"`
    // CreateAccount automatically create a PlayFab account if one is not currently linked to this ID.
    CreateAccount bool `json:"CreateAccount"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // EncryptedRequest base64 encoded body that is encrypted with the Title's public RSA key (Enterprise Only).
    EncryptedRequest string `json:"EncryptedRequest,omitempty"`
    // InfoRequestParameters flags for which pieces of info to return for the user.
    InfoRequestParameters *GetPlayerCombinedInfoRequestParamsModel `json:"InfoRequestParameters,omitempty"`
    // IssuerId id of the PSN issuer environment. If null, defaults to 256 (production)
    IssuerId int32 `json:"IssuerId,omitempty"`
    // PlayerSecret player secret that is used to verify API request signatures (Enterprise Only).
    PlayerSecret string `json:"PlayerSecret,omitempty"`
    // RedirectUri redirect URI supplied to PSN when requesting an auth code
    RedirectUri string `json:"RedirectUri,omitempty"`
    // TitleId unique identifier for the title, found in the Settings > Game Properties section of the PlayFab developer site when a
// title has been selected.
    TitleId string `json:"TitleId,omitempty"`
}

// LoginWithSteamRequest steam sign-in is accomplished with the Steam Session Ticket. More information on the Ticket can be found in the
// Steamworks SDK, here: https://partner.steamgames.com/documentation/auth (requires sign-in). NOTE: For Steam
// authentication to work, the title must be configured with the Steam Application ID and Web API Key in the PlayFab Game
// Manager (under Steam in the Add-ons Marketplace). You can obtain a Web API Key from the Permissions page of any Group
// associated with your App ID in the Steamworks site. If this is the first time a user has signed in with the Steam
// account and CreateAccount is set to true, a new PlayFab account will be created and linked to the provided account's
// Steam ID. In this case, no email or username will be associated with the PlayFab account. Otherwise, if no PlayFab
// account is linked to the Steam account, an error indicating this will be returned, so that the title can guide the user
// through creation of a PlayFab account.
type LoginWithSteamRequestModel struct {
    // CreateAccount automatically create a PlayFab account if one is not currently linked to this ID.
    CreateAccount bool `json:"CreateAccount"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // EncryptedRequest base64 encoded body that is encrypted with the Title's public RSA key (Enterprise Only).
    EncryptedRequest string `json:"EncryptedRequest,omitempty"`
    // InfoRequestParameters flags for which pieces of info to return for the user.
    InfoRequestParameters *GetPlayerCombinedInfoRequestParamsModel `json:"InfoRequestParameters,omitempty"`
    // PlayerSecret player secret that is used to verify API request signatures (Enterprise Only).
    PlayerSecret string `json:"PlayerSecret,omitempty"`
    // SteamTicket authentication token for the user, returned as a byte array from Steam, and converted to a string (for example, the byte
// 0x08 should become "08").
    SteamTicket string `json:"SteamTicket,omitempty"`
    // TitleId unique identifier for the title, found in the Settings > Game Properties section of the PlayFab developer site when a
// title has been selected.
    TitleId string `json:"TitleId,omitempty"`
}

// LoginWithTwitchRequest more details regarding Twitch and their authentication system can be found at
// https://github.com/justintv/Twitch-API/blob/master/authentication.md. Developers must provide the Twitch access token
// that is generated using one of the Twitch authentication flows. PlayFab will use the title's unique Twitch Client ID to
// authenticate the token and log in to the PlayFab system. If CreateAccount is set to true and there is not already a user
// matched to the Twitch username that generated the token, then PlayFab will create a new account for this user and link
// the ID. In this case, no email or username will be associated with the PlayFab account. If there is already a different
// PlayFab user linked with this account, then an error will be returned.
type LoginWithTwitchRequestModel struct {
    // AccessToken token issued by Twitch's API for the user.
    AccessToken string `json:"AccessToken,omitempty"`
    // CreateAccount automatically create a PlayFab account if one is not currently linked to this ID.
    CreateAccount bool `json:"CreateAccount"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // EncryptedRequest base64 encoded body that is encrypted with the Title's public RSA key (Enterprise Only).
    EncryptedRequest string `json:"EncryptedRequest,omitempty"`
    // InfoRequestParameters flags for which pieces of info to return for the user.
    InfoRequestParameters *GetPlayerCombinedInfoRequestParamsModel `json:"InfoRequestParameters,omitempty"`
    // PlayerSecret player secret that is used to verify API request signatures (Enterprise Only).
    PlayerSecret string `json:"PlayerSecret,omitempty"`
    // TitleId unique identifier for the title, found in the Settings > Game Properties section of the PlayFab developer site when a
// title has been selected.
    TitleId string `json:"TitleId,omitempty"`
}

// LoginWithWindowsHelloRequest requires both the SHA256 hash of the user's public key as well as the signed response from GetWindowsHelloChallenge
type LoginWithWindowsHelloRequestModel struct {
    // ChallengeSignature the signed response from the user for the Challenge.
    ChallengeSignature string `json:"ChallengeSignature,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // InfoRequestParameters flags for which pieces of info to return for the user.
    InfoRequestParameters *GetPlayerCombinedInfoRequestParamsModel `json:"InfoRequestParameters,omitempty"`
    // PublicKeyHint sHA256 hash of the PublicKey generated by Windows Hello.
    PublicKeyHint string `json:"PublicKeyHint,omitempty"`
    // TitleId unique identifier for the title, found in the Settings > Game Properties section of the PlayFab developer site when a
// title has been selected.
    TitleId string `json:"TitleId,omitempty"`
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
    // EncryptedRequest base64 encoded body that is encrypted with the Title's public RSA key (Enterprise Only).
    EncryptedRequest string `json:"EncryptedRequest,omitempty"`
    // InfoRequestParameters flags for which pieces of info to return for the user.
    InfoRequestParameters *GetPlayerCombinedInfoRequestParamsModel `json:"InfoRequestParameters,omitempty"`
    // PlayerSecret player secret that is used to verify API request signatures (Enterprise Only).
    PlayerSecret string `json:"PlayerSecret,omitempty"`
    // TitleId unique identifier for the title, found in the Settings > Game Properties section of the PlayFab developer site when a
// title has been selected.
    TitleId string `json:"TitleId,omitempty"`
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

// MatchmakeRequest 
type MatchmakeRequestModel struct {
    // BuildVersion build version to match against. [Note: Required if LobbyId is not specified]
    BuildVersion string `json:"BuildVersion,omitempty"`
    // CharacterId character to use for stats based matching. Leave null to use account stats.
    CharacterId string `json:"CharacterId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // GameMode game mode to match make against. [Note: Required if LobbyId is not specified]
    GameMode string `json:"GameMode,omitempty"`
    // LobbyId lobby identifier to match make against. This is used to select a specific Game Server Instance.
    LobbyId string `json:"LobbyId,omitempty"`
    // Region region to match make against. [Note: Required if LobbyId is not specified]
    Region Region `json:"Region,omitempty"`
    // StartNewIfNoneFound start a game session if one with an open slot is not found. Defaults to true.
    StartNewIfNoneFound bool `json:"StartNewIfNoneFound"`
    // StatisticName player statistic to use in finding a match. May be null for no stat-based matching.
    StatisticName string `json:"StatisticName,omitempty"`
    // TagFilter filter to include and/or exclude Game Server Instances associated with certain Tags
    TagFilter *CollectionFilterModel `json:"TagFilter,omitempty"`
}

// MatchmakeResult 
type MatchmakeResultModel struct {
    // Expires timestamp for when the server will expire, if applicable
    Expires string `json:"Expires,omitempty"`
    // LobbyID unique lobby identifier of the server matched
    LobbyID string `json:"LobbyID,omitempty"`
    // PollWaitTimeMS time in milliseconds the application is configured to wait on matchmaking results
    PollWaitTimeMS int32 `json:"PollWaitTimeMS,omitempty"`
    // ServerIPV4Address iPV4 address of the server
    ServerIPV4Address string `json:"ServerIPV4Address,omitempty"`
    // ServerIPV6Address iPV6 address of the server
    ServerIPV6Address string `json:"ServerIPV6Address,omitempty"`
    // ServerPort port number to use for non-http communications with the server
    ServerPort int32 `json:"ServerPort,omitempty"`
    // ServerPublicDNSName public DNS name (if any) of the server
    ServerPublicDNSName string `json:"ServerPublicDNSName,omitempty"`
    // Status result of match making process
    Status MatchmakeStatus `json:"Status,omitempty"`
    // Ticket server authorization ticket (used by RedeemMatchmakerTicket to validate user insertion into the game)
    Ticket string `json:"Ticket,omitempty"`
}

// MatchmakeStatus 
type MatchmakeStatus string
  
const (
     MatchmakeStatusComplete MatchmakeStatus = "Complete"
     MatchmakeStatusWaiting MatchmakeStatus = "Waiting"
     MatchmakeStatusGameNotFound MatchmakeStatus = "GameNotFound"
     MatchmakeStatusNoAvailableSlots MatchmakeStatus = "NoAvailableSlots"
     MatchmakeStatusSessionClosed MatchmakeStatus = "SessionClosed"
      )
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

// MicrosoftStorePayload 
type MicrosoftStorePayloadModel struct {
    // CollectionsMsIdKey microsoft store ID key. This is optional. Alternatively you can use XboxToken
    CollectionsMsIdKey string `json:"CollectionsMsIdKey,omitempty"`
    // UserId if collectionsMsIdKey is provided, this will verify the user id in the collectionsMsIdKey is the same.
    UserId string `json:"UserId,omitempty"`
    // XboxToken token provided by the Xbox Live SDK/XDK method GetTokenAndSignatureAsync("POST", "https://playfabapi.com/", ""). This is
// optional. Alternatively can use CollectionsMsIdKey
    XboxToken string `json:"XboxToken,omitempty"`
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

// NintendoSwitchPlayFabIdPair 
type NintendoSwitchPlayFabIdPairModel struct {
    // NintendoSwitchDeviceId unique Nintendo Switch Device identifier for a user.
    NintendoSwitchDeviceId string `json:"NintendoSwitchDeviceId,omitempty"`
    // PlayFabId unique PlayFab identifier for a user, or null if no PlayFab account is linked to the Nintendo Switch Device identifier.
    PlayFabId string `json:"PlayFabId,omitempty"`
}

// OpenTradeRequest 
type OpenTradeRequestModel struct {
    // AllowedPlayerIds players who are allowed to accept the trade. If null, the trade may be accepted by any player. If empty, the trade may
// not be accepted by any player.
    AllowedPlayerIds []string `json:"AllowedPlayerIds,omitempty"`
    // OfferedInventoryInstanceIds player inventory items offered for trade. If not set, the trade is effectively a gift request
    OfferedInventoryInstanceIds []string `json:"OfferedInventoryInstanceIds,omitempty"`
    // RequestedCatalogItemIds catalog items accepted for the trade. If not set, the trade is effectively a gift.
    RequestedCatalogItemIds []string `json:"RequestedCatalogItemIds,omitempty"`
}

// OpenTradeResponse 
type OpenTradeResponseModel struct {
    // Trade the information about the trade that was just opened.
    Trade *TradeInfoModel `json:"Trade,omitempty"`
}

// PayForPurchaseRequest this is the second step in the purchasing process, initiating the purchase transaction with the payment provider (if
// applicable). For payment provider scenarios, the title should next present the user with the payment provider'sinterface
// for payment. Once the player has completed the payment with the provider, the title should call ConfirmPurchase
// tofinalize the process and add the appropriate items to the player inventory.
type PayForPurchaseRequestModel struct {
    // Currency currency to use to fund the purchase.
    Currency string `json:"Currency,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // OrderId purchase order identifier returned from StartPurchase.
    OrderId string `json:"OrderId,omitempty"`
    // ProviderName payment provider to use to fund the purchase.
    ProviderName string `json:"ProviderName,omitempty"`
    // ProviderTransactionId payment provider transaction identifier. Required for Facebook Payments.
    ProviderTransactionId string `json:"ProviderTransactionId,omitempty"`
}

// PayForPurchaseResult for web-based payment providers, this operation returns the URL to which the user should be directed inorder to approve
// the purchase. Items added to the user inventory as a result of this operation will be marked as unconfirmed.
type PayForPurchaseResultModel struct {
    // CreditApplied local credit applied to the transaction (provider specific).
    CreditApplied uint32 `json:"CreditApplied,omitempty"`
    // OrderId purchase order identifier.
    OrderId string `json:"OrderId,omitempty"`
    // ProviderData provider used for the transaction.
    ProviderData string `json:"ProviderData,omitempty"`
    // ProviderToken a token generated by the provider to authenticate the request (provider-specific).
    ProviderToken string `json:"ProviderToken,omitempty"`
    // PurchaseConfirmationPageURL uRL to the purchase provider page that details the purchase.
    PurchaseConfirmationPageURL string `json:"PurchaseConfirmationPageURL,omitempty"`
    // PurchaseCurrency currency for the transaction, may be a virtual currency or real money.
    PurchaseCurrency string `json:"PurchaseCurrency,omitempty"`
    // PurchasePrice cost of the transaction.
    PurchasePrice uint32 `json:"PurchasePrice,omitempty"`
    // Status status of the transaction.
    Status TransactionStatus `json:"Status,omitempty"`
    // VCAmount virtual currencies granted by the transaction, if any.
    VCAmount map[string]int32 `json:"VCAmount,omitempty"`
    // VirtualCurrency current virtual currency balances for the user.
    VirtualCurrency map[string]int32 `json:"VirtualCurrency,omitempty"`
}

// PaymentOption 
type PaymentOptionModel struct {
    // Currency specific currency to use to fund the purchase.
    Currency string `json:"Currency,omitempty"`
    // Price amount of the specified currency needed for the purchase.
    Price uint32 `json:"Price,omitempty"`
    // ProviderName name of the purchase provider for this option.
    ProviderName string `json:"ProviderName,omitempty"`
    // StoreCredit amount of existing credit the user has with the provider.
    StoreCredit uint32 `json:"StoreCredit,omitempty"`
}

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

// PurchaseItemRequest please note that the processing time for inventory grants and purchases increases fractionally the more items are in the
// inventory, and the more items are in the grant/purchase operation (with each item in a bundle being a distinct add).
type PurchaseItemRequestModel struct {
    // CatalogVersion catalog version for the items to be purchased (defaults to most recent version.
    CatalogVersion string `json:"CatalogVersion,omitempty"`
    // CharacterId unique PlayFab assigned ID for a specific character owned by a user
    CharacterId string `json:"CharacterId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // ItemId unique identifier of the item to purchase.
    ItemId string `json:"ItemId,omitempty"`
    // Price price the client expects to pay for the item (in case a new catalog or store was uploaded, with new prices).
    Price int32 `json:"Price,omitempty"`
    // StoreId store to buy this item through. If not set, prices default to those in the catalog.
    StoreId string `json:"StoreId,omitempty"`
    // VirtualCurrency virtual currency to use to purchase the item.
    VirtualCurrency string `json:"VirtualCurrency,omitempty"`
}

// PurchaseItemResult 
type PurchaseItemResultModel struct {
    // Items details for the items purchased.
    Items []ItemInstanceModel `json:"Items,omitempty"`
}

// PurchaseReceiptFulfillment 
type PurchaseReceiptFulfillmentModel struct {
    // FulfilledItems items granted to the player in fulfillment of the validated receipt.
    FulfilledItems []ItemInstanceModel `json:"FulfilledItems,omitempty"`
    // RecordedPriceSource source of the payment price information for the recorded purchase transaction. A value of 'Request' indicates that the
// price specified in the request was used, whereas a value of 'Catalog' indicates that the real-money price of the catalog
// item matching the product ID in the validated receipt transaction and the currency specified in the request (defaulting
// to USD) was used.
    RecordedPriceSource string `json:"RecordedPriceSource,omitempty"`
    // RecordedTransactionCurrency currency used to purchase the items (ISO 4217 currency code).
    RecordedTransactionCurrency string `json:"RecordedTransactionCurrency,omitempty"`
    // RecordedTransactionTotal amount of the stated currency paid for the items, in centesimal units
    RecordedTransactionTotal uint32 `json:"RecordedTransactionTotal,omitempty"`
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

// RedeemCouponRequest coupon codes can be created for any item, or set of items, in the catalog for the title. This operation causes the
// coupon to be consumed, and the specific items to be awarded to the user. Attempting to re-use an already consumed code,
// or a code which has not yet been created in the service, will result in an error.
type RedeemCouponRequestModel struct {
    // CatalogVersion catalog version of the coupon. If null, uses the default catalog
    CatalogVersion string `json:"CatalogVersion,omitempty"`
    // CharacterId optional identifier for the Character that should receive the item. If null, item is added to the player
    CharacterId string `json:"CharacterId,omitempty"`
    // CouponCode generated coupon code to redeem.
    CouponCode string `json:"CouponCode,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
}

// RedeemCouponResult 
type RedeemCouponResultModel struct {
    // GrantedItems items granted to the player as a result of redeeming the coupon.
    GrantedItems []ItemInstanceModel `json:"GrantedItems,omitempty"`
}

// RefreshPSNAuthTokenRequest 
type RefreshPSNAuthTokenRequestModel struct {
    // AuthCode auth code returned by PSN OAuth system.
    AuthCode string `json:"AuthCode,omitempty"`
    // IssuerId id of the PSN issuer environment. If null, defaults to 256 (production)
    IssuerId int32 `json:"IssuerId,omitempty"`
    // RedirectUri redirect URI supplied to PSN when requesting an auth code
    RedirectUri string `json:"RedirectUri,omitempty"`
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
// RegionInfo 
type RegionInfoModel struct {
    // Available indicates whether the server specified is available in this region
    Available bool `json:"Available"`
    // Name name of the region
    Name string `json:"Name,omitempty"`
    // PingUrl url to ping to get roundtrip time
    PingUrl string `json:"PingUrl,omitempty"`
    // Region unique identifier for the region
    Region Region `json:"Region,omitempty"`
}

// RegisterForIOSPushNotificationRequest the steps to configure and send Push Notifications is described in the PlayFab tutorials, here:
// https://docs.microsoft.com/gaming/playfab/features/engagement/push-notifications/quickstart
type RegisterForIOSPushNotificationRequestModel struct {
    // ConfirmationMessage message to display when confirming push notification.
    ConfirmationMessage string `json:"ConfirmationMessage,omitempty"`
    // DeviceToken unique token generated by the Apple Push Notification service when the title registered to receive push notifications.
    DeviceToken string `json:"DeviceToken,omitempty"`
    // SendPushNotificationConfirmation if true, send a test push message immediately after sucessful registration. Defaults to false.
    SendPushNotificationConfirmation bool `json:"SendPushNotificationConfirmation"`
}

// RegisterForIOSPushNotificationResult 
type RegisterForIOSPushNotificationResultModel struct {
}

// RegisterPlayFabUserRequest 
type RegisterPlayFabUserRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // DisplayName an optional parameter for setting the display name for this title (3-25 characters).
    DisplayName string `json:"DisplayName,omitempty"`
    // Email user email address attached to their account
    Email string `json:"Email,omitempty"`
    // EncryptedRequest base64 encoded body that is encrypted with the Title's public RSA key (Enterprise Only).
    EncryptedRequest string `json:"EncryptedRequest,omitempty"`
    // InfoRequestParameters flags for which pieces of info to return for the user.
    InfoRequestParameters *GetPlayerCombinedInfoRequestParamsModel `json:"InfoRequestParameters,omitempty"`
    // Password password for the PlayFab account (6-100 characters)
    Password string `json:"Password,omitempty"`
    // PlayerSecret player secret that is used to verify API request signatures (Enterprise Only).
    PlayerSecret string `json:"PlayerSecret,omitempty"`
    // RequireBothUsernameAndEmail an optional parameter that specifies whether both the username and email parameters are required. If true, both
// parameters are required; if false, the user must supply either the username or email parameter. The default value is
// true.
    RequireBothUsernameAndEmail bool `json:"RequireBothUsernameAndEmail"`
    // TitleId unique identifier for the title, found in the Settings > Game Properties section of the PlayFab developer site when a
// title has been selected.
    TitleId string `json:"TitleId,omitempty"`
    // Username playFab username for the account (3-20 characters)
    Username string `json:"Username,omitempty"`
}

// RegisterPlayFabUserResult each account must have a unique email address in the PlayFab service. Once created, the account may be associated with
// additional accounts (Steam, Facebook, Game Center, etc.), allowing for added social network lists and achievements
// systems.
type RegisterPlayFabUserResultModel struct {
    // EntityToken if LoginTitlePlayerAccountEntity flag is set on the login request the title_player_account will also be logged in and
// returned.
    EntityToken *EntityTokenResponseModel `json:"EntityToken,omitempty"`
    // PlayFabId playFab unique identifier for this newly created account.
    PlayFabId string `json:"PlayFabId,omitempty"`
    // SessionTicket unique token identifying the user and game at the server level, for the current session.
    SessionTicket string `json:"SessionTicket,omitempty"`
    // SettingsForUser settings specific to this user.
    SettingsForUser *UserSettingsModel `json:"SettingsForUser,omitempty"`
    // Username playFab unique user name.
    Username string `json:"Username,omitempty"`
}

// RegisterWithWindowsHelloRequest publicKey must be generated using the Windows Hello Passport service.
type RegisterWithWindowsHelloRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // DeviceName device name.
    DeviceName string `json:"DeviceName,omitempty"`
    // EncryptedRequest base64 encoded body that is encrypted with the Title's public RSA key (Enterprise Only).
    EncryptedRequest string `json:"EncryptedRequest,omitempty"`
    // InfoRequestParameters flags for which pieces of info to return for the user.
    InfoRequestParameters *GetPlayerCombinedInfoRequestParamsModel `json:"InfoRequestParameters,omitempty"`
    // PlayerSecret player secret that is used to verify API request signatures (Enterprise Only).
    PlayerSecret string `json:"PlayerSecret,omitempty"`
    // PublicKey publicKey generated by Windows Hello.
    PublicKey string `json:"PublicKey,omitempty"`
    // TitleId unique identifier for the title, found in the Settings > Game Properties section of the PlayFab developer site when a
// title has been selected.
    TitleId string `json:"TitleId,omitempty"`
    // UserName player's user name used by Windows Hello.
    UserName string `json:"UserName,omitempty"`
}

// RemoveContactEmailRequest this API removes an existing contact email from the player's profile.
type RemoveContactEmailRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
}

// RemoveContactEmailResult 
type RemoveContactEmailResultModel struct {
}

// RemoveFriendRequest 
type RemoveFriendRequestModel struct {
    // FriendPlayFabId playFab identifier of the friend account which is to be removed.
    FriendPlayFabId string `json:"FriendPlayFabId,omitempty"`
}

// RemoveFriendResult 
type RemoveFriendResultModel struct {
}

// RemoveGenericIDRequest 
type RemoveGenericIDRequestModel struct {
    // GenericId generic service identifier to be removed from the player.
    GenericId* GenericServiceIdModel `json:"GenericId,omitempty"`
}

// RemoveGenericIDResult 
type RemoveGenericIDResultModel struct {
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

// ReportAdActivityRequest report ad activity
type ReportAdActivityRequestModel struct {
    // Activity type of activity, may be Opened, Closed, Start or End
    Activity AdActivity `json:"Activity,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // PlacementId unique ID of the placement to report for
    PlacementId string `json:"PlacementId,omitempty"`
    // RewardId unique ID of the reward the player was offered
    RewardId string `json:"RewardId,omitempty"`
}

// ReportAdActivityResult report ad activity response has no body
type ReportAdActivityResultModel struct {
}

// ReportPlayerClientRequest 
type ReportPlayerClientRequestModel struct {
    // Comment optional additional comment by reporting player.
    Comment string `json:"Comment,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // ReporteeId unique PlayFab identifier of the reported player.
    ReporteeId string `json:"ReporteeId,omitempty"`
}

// ReportPlayerClientResult players are currently limited to five reports per day. Attempts by a single user account to submit reports beyond five
// will result in Updated being returned as false.
type ReportPlayerClientResultModel struct {
    // SubmissionsRemaining the number of remaining reports which may be filed today.
    SubmissionsRemaining int32 `json:"SubmissionsRemaining,omitempty"`
}

// RestoreIOSPurchasesRequest the title should obtain a refresh receipt via restoreCompletedTransactions in the SKPaymentQueue of the Apple StoreKit
// and pass that in to this call. The resultant receipt contains new receipt instances for all non-consumable goods
// previously purchased by the user. This API call iterates through every purchase in the receipt and restores the items if
// they still exist in the catalog and can be validated.
type RestoreIOSPurchasesRequestModel struct {
    // CatalogVersion catalog version of the restored items. If null, defaults to primary catalog.
    CatalogVersion string `json:"CatalogVersion,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // ReceiptData base64 encoded receipt data, passed back by the App Store as a result of a successful purchase.
    ReceiptData string `json:"ReceiptData,omitempty"`
}

// RestoreIOSPurchasesResult once verified, the valid items will be restored into the user's inventory.
type RestoreIOSPurchasesResultModel struct {
    // Fulfillments fulfilled inventory items and recorded purchases in fulfillment of the validated receipt transactions.
    Fulfillments []PurchaseReceiptFulfillmentModel `json:"Fulfillments,omitempty"`
}

// RewardAdActivityRequest details on which placement and reward to perform a grant on
type RewardAdActivityRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // PlacementId placement unique ID
    PlacementId string `json:"PlacementId,omitempty"`
    // RewardId reward unique ID
    RewardId string `json:"RewardId,omitempty"`
}

// RewardAdActivityResult result for rewarding an ad activity
type RewardAdActivityResultModel struct {
    // AdActivityEventId playStream Event ID that was generated by this reward (all subsequent events are associated with this event identifier)
    AdActivityEventId string `json:"AdActivityEventId,omitempty"`
    // DebugResults debug results from the grants
    DebugResults []string `json:"DebugResults,omitempty"`
    // PlacementId id of the placement the reward was for
    PlacementId string `json:"PlacementId,omitempty"`
    // PlacementName name of the placement the reward was for
    PlacementName string `json:"PlacementName,omitempty"`
    // PlacementViewsRemaining if placement has viewing limits indicates how many views are left
    PlacementViewsRemaining int32 `json:"PlacementViewsRemaining,omitempty"`
    // PlacementViewsResetMinutes if placement has viewing limits indicates when they will next reset
    PlacementViewsResetMinutes float64 `json:"PlacementViewsResetMinutes,omitempty"`
    // RewardResults reward results
    RewardResults *AdRewardResultsModel `json:"RewardResults,omitempty"`
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
    // TitleId unique identifier for the title, found in the Settings > Game Properties section of the PlayFab developer site when a
// title has been selected.
    TitleId string `json:"TitleId,omitempty"`
}

// SendAccountRecoveryEmailResult 
type SendAccountRecoveryEmailResultModel struct {
}

// SetFriendTagsRequest this operation is not additive. It will completely replace the tag list for the specified user. Please note that only
// users in the PlayFab friends list can be assigned tags. Attempting to set a tag on a friend only included in the friends
// list from a social site integration (such as Facebook or Steam) will return the AccountNotFound error.
type SetFriendTagsRequestModel struct {
    // FriendPlayFabId playFab identifier of the friend account to which the tag(s) should be applied.
    FriendPlayFabId string `json:"FriendPlayFabId,omitempty"`
    // Tags array of tags to set on the friend account.
    Tags []string `json:"Tags,omitempty"`
}

// SetFriendTagsResult 
type SetFriendTagsResultModel struct {
}

// SetPlayerSecretRequest aPIs that require signatures require that the player have a configured Player Secret Key that is used to sign all
// requests. Players that don't have a secret will be blocked from making API calls until it is configured. To create a
// signature header add a SHA256 hashed string containing UTF8 encoded JSON body as it will be sent to the server, the
// current time in UTC formatted to ISO 8601, and the players secret formatted as 'body.date.secret'. Place the resulting
// hash into the header X-PlayFab-Signature, along with a header X-PlayFab-Timestamp of the same UTC timestamp used in the
// signature.
type SetPlayerSecretRequestModel struct {
    // EncryptedRequest base64 encoded body that is encrypted with the Title's public RSA key (Enterprise Only).
    EncryptedRequest string `json:"EncryptedRequest,omitempty"`
    // PlayerSecret player secret that is used to verify API request signatures (Enterprise Only).
    PlayerSecret string `json:"PlayerSecret,omitempty"`
}

// SetPlayerSecretResult 
type SetPlayerSecretResultModel struct {
}

// SharedGroupDataRecord 
type SharedGroupDataRecordModel struct {
    // LastUpdated timestamp for when this data was last updated.
    LastUpdated time.Time `json:"LastUpdated,omitempty"`
    // LastUpdatedBy unique PlayFab identifier of the user to last update this value.
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
// StartGameRequest this API must be enabled for use as an option in the game manager website. It is disabled by default.
type StartGameRequestModel struct {
    // BuildVersion version information for the build of the game server which is to be started
    BuildVersion string `json:"BuildVersion,omitempty"`
    // CharacterId character to use for stats based matching. Leave null to use account stats
    CharacterId string `json:"CharacterId,omitempty"`
    // CustomCommandLineData custom command line argument when starting game server process
    CustomCommandLineData string `json:"CustomCommandLineData,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // GameMode the title-defined game mode this server is to be running (defaults to 0 if there is only one mode)
    GameMode string `json:"GameMode,omitempty"`
    // Region the region to associate this server with for match filtering
    Region Region `json:"Region,omitempty"`
    // StatisticName player statistic for others to use in finding this game. May be null for no stat-based matching
    StatisticName string `json:"StatisticName,omitempty"`
}

// StartGameResult 
type StartGameResultModel struct {
    // Expires timestamp for when the server should expire, if applicable
    Expires string `json:"Expires,omitempty"`
    // LobbyID unique identifier for the lobby of the server started
    LobbyID string `json:"LobbyID,omitempty"`
    // Password password required to log into the server
    Password string `json:"Password,omitempty"`
    // ServerIPV4Address server IPV4 address
    ServerIPV4Address string `json:"ServerIPV4Address,omitempty"`
    // ServerIPV6Address server IPV6 address
    ServerIPV6Address string `json:"ServerIPV6Address,omitempty"`
    // ServerPort port on the server to be used for communication
    ServerPort int32 `json:"ServerPort,omitempty"`
    // ServerPublicDNSName server public DNS name
    ServerPublicDNSName string `json:"ServerPublicDNSName,omitempty"`
    // Ticket unique identifier for the server
    Ticket string `json:"Ticket,omitempty"`
}

// StartPurchaseRequest this is the first step in the purchasing process. For security purposes, once the order (or "cart") has been created,
// additional inventory objects may no longer be added. In addition, inventory objects will be locked to the current
// prices, regardless of any subsequent changes at the catalog level which may occur during the next two steps.
type StartPurchaseRequestModel struct {
    // CatalogVersion catalog version for the items to be purchased. Defaults to most recent catalog.
    CatalogVersion string `json:"CatalogVersion,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Items array of items to purchase.
    Items []ItemPurchaseRequestModel `json:"Items,omitempty"`
    // StoreId store through which to purchase items. If not set, prices will be pulled from the catalog itself.
    StoreId string `json:"StoreId,omitempty"`
}

// StartPurchaseResult 
type StartPurchaseResultModel struct {
    // Contents cart items to be purchased.
    Contents []CartItemModel `json:"Contents,omitempty"`
    // OrderId purchase order identifier.
    OrderId string `json:"OrderId,omitempty"`
    // PaymentOptions available methods by which the user can pay.
    PaymentOptions []PaymentOptionModel `json:"PaymentOptions,omitempty"`
    // VirtualCurrencyBalances current virtual currency totals for the user.
    VirtualCurrencyBalances map[string]int32 `json:"VirtualCurrencyBalances,omitempty"`
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
// SubtractUserVirtualCurrencyRequest this API must be enabled for use as an option in the game manager website. It is disabled by default.
type SubtractUserVirtualCurrencyRequestModel struct {
    // Amount amount to be subtracted from the user balance of the specified virtual currency.
    Amount int32 `json:"Amount,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
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
    // Body news item text.
    Body string `json:"Body,omitempty"`
    // NewsId unique identifier of news item.
    NewsId string `json:"NewsId,omitempty"`
    // Timestamp date and time when the news item was posted.
    Timestamp time.Time `json:"Timestamp,omitempty"`
    // Title title of the news item.
    Title string `json:"Title,omitempty"`
}

// TradeInfo 
type TradeInfoModel struct {
    // AcceptedInventoryInstanceIds item instances from the accepting player that are used to fulfill the trade. If null, no one has accepted the trade.
    AcceptedInventoryInstanceIds []string `json:"AcceptedInventoryInstanceIds,omitempty"`
    // AcceptedPlayerId the PlayFab ID of the player who accepted the trade. If null, no one has accepted the trade.
    AcceptedPlayerId string `json:"AcceptedPlayerId,omitempty"`
    // AllowedPlayerIds an optional list of players allowed to complete this trade. If null, anybody can complete the trade.
    AllowedPlayerIds []string `json:"AllowedPlayerIds,omitempty"`
    // CancelledAt if set, The UTC time when this trade was canceled.
    CancelledAt time.Time `json:"CancelledAt,omitempty"`
    // FilledAt if set, The UTC time when this trade was fulfilled.
    FilledAt time.Time `json:"FilledAt,omitempty"`
    // InvalidatedAt if set, The UTC time when this trade was made invalid.
    InvalidatedAt time.Time `json:"InvalidatedAt,omitempty"`
    // OfferedCatalogItemIds the catalogItem Ids of the item instances being offered.
    OfferedCatalogItemIds []string `json:"OfferedCatalogItemIds,omitempty"`
    // OfferedInventoryInstanceIds the itemInstance Ids that are being offered.
    OfferedInventoryInstanceIds []string `json:"OfferedInventoryInstanceIds,omitempty"`
    // OfferingPlayerId the PlayFabId for the offering player.
    OfferingPlayerId string `json:"OfferingPlayerId,omitempty"`
    // OpenedAt the UTC time when this trade was created.
    OpenedAt time.Time `json:"OpenedAt,omitempty"`
    // RequestedCatalogItemIds the catalogItem Ids requested in exchange.
    RequestedCatalogItemIds []string `json:"RequestedCatalogItemIds,omitempty"`
    // Status describes the current state of this trade.
    Status TradeStatus `json:"Status,omitempty"`
    // TradeId the identifier for this trade.
    TradeId string `json:"TradeId,omitempty"`
}

// TradeStatus 
type TradeStatus string
  
const (
     TradeStatusInvalid TradeStatus = "Invalid"
     TradeStatusOpening TradeStatus = "Opening"
     TradeStatusOpen TradeStatus = "Open"
     TradeStatusAccepting TradeStatus = "Accepting"
     TradeStatusAccepted TradeStatus = "Accepted"
     TradeStatusFilled TradeStatus = "Filled"
     TradeStatusCancelled TradeStatus = "Cancelled"
      )
// TransactionStatus 
type TransactionStatus string
  
const (
     TransactionStatusCreateCart TransactionStatus = "CreateCart"
     TransactionStatusInit TransactionStatus = "Init"
     TransactionStatusApproved TransactionStatus = "Approved"
     TransactionStatusSucceeded TransactionStatus = "Succeeded"
     TransactionStatusFailedByProvider TransactionStatus = "FailedByProvider"
     TransactionStatusDisputePending TransactionStatus = "DisputePending"
     TransactionStatusRefundPending TransactionStatus = "RefundPending"
     TransactionStatusRefunded TransactionStatus = "Refunded"
     TransactionStatusRefundFailed TransactionStatus = "RefundFailed"
     TransactionStatusChargedBack TransactionStatus = "ChargedBack"
     TransactionStatusFailedByUber TransactionStatus = "FailedByUber"
     TransactionStatusFailedByPlayFab TransactionStatus = "FailedByPlayFab"
     TransactionStatusRevoked TransactionStatus = "Revoked"
     TransactionStatusTradePending TransactionStatus = "TradePending"
     TransactionStatusTraded TransactionStatus = "Traded"
     TransactionStatusUpgraded TransactionStatus = "Upgraded"
     TransactionStatusStackPending TransactionStatus = "StackPending"
     TransactionStatusStacked TransactionStatus = "Stacked"
     TransactionStatusOther TransactionStatus = "Other"
     TransactionStatusFailed TransactionStatus = "Failed"
      )
// TreatmentAssignment 
type TreatmentAssignmentModel struct {
    // Variables list of the experiment variables.
    Variables []VariableModel `json:"Variables,omitempty"`
    // Variants list of the experiment variants.
    Variants []string `json:"Variants,omitempty"`
}

// TwitchPlayFabIdPair 
type TwitchPlayFabIdPairModel struct {
    // PlayFabId unique PlayFab identifier for a user, or null if no PlayFab account is linked to the Twitch identifier.
    PlayFabId string `json:"PlayFabId,omitempty"`
    // TwitchId unique Twitch identifier for a user.
    TwitchId string `json:"TwitchId,omitempty"`
}

// UnlinkAndroidDeviceIDRequest 
type UnlinkAndroidDeviceIDRequestModel struct {
    // AndroidDeviceId android device identifier for the user's device. If not specified, the most recently signed in Android Device ID will be
// used.
    AndroidDeviceId string `json:"AndroidDeviceId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
}

// UnlinkAndroidDeviceIDResult 
type UnlinkAndroidDeviceIDResultModel struct {
}

// UnlinkAppleRequest 
type UnlinkAppleRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
}

// UnlinkCustomIDRequest 
type UnlinkCustomIDRequestModel struct {
    // CustomId custom unique identifier for the user, generated by the title. If not specified, the most recently signed in Custom ID
// will be used.
    CustomId string `json:"CustomId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
}

// UnlinkCustomIDResult 
type UnlinkCustomIDResultModel struct {
}

// UnlinkFacebookAccountRequest 
type UnlinkFacebookAccountRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
}

// UnlinkFacebookAccountResult 
type UnlinkFacebookAccountResultModel struct {
}

// UnlinkFacebookInstantGamesIdRequest 
type UnlinkFacebookInstantGamesIdRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // FacebookInstantGamesId facebook Instant Games identifier for the user. If not specified, the most recently signed in ID will be used.
    FacebookInstantGamesId string `json:"FacebookInstantGamesId,omitempty"`
}

// UnlinkFacebookInstantGamesIdResult 
type UnlinkFacebookInstantGamesIdResultModel struct {
}

// UnlinkGameCenterAccountRequest 
type UnlinkGameCenterAccountRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
}

// UnlinkGameCenterAccountResult 
type UnlinkGameCenterAccountResultModel struct {
}

// UnlinkGoogleAccountRequest 
type UnlinkGoogleAccountRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
}

// UnlinkGoogleAccountResult 
type UnlinkGoogleAccountResultModel struct {
}

// UnlinkIOSDeviceIDRequest 
type UnlinkIOSDeviceIDRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // DeviceId vendor-specific iOS identifier for the user's device. If not specified, the most recently signed in iOS Device ID will
// be used.
    DeviceId string `json:"DeviceId,omitempty"`
}

// UnlinkIOSDeviceIDResult 
type UnlinkIOSDeviceIDResultModel struct {
}

// UnlinkKongregateAccountRequest 
type UnlinkKongregateAccountRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
}

// UnlinkKongregateAccountResult 
type UnlinkKongregateAccountResultModel struct {
}

// UnlinkNintendoServiceAccountRequest 
type UnlinkNintendoServiceAccountRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
}

// UnlinkNintendoSwitchDeviceIdRequest 
type UnlinkNintendoSwitchDeviceIdRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // NintendoSwitchDeviceId nintendo Switch Device identifier for the user. If not specified, the most recently signed in device ID will be used.
    NintendoSwitchDeviceId string `json:"NintendoSwitchDeviceId,omitempty"`
}

// UnlinkNintendoSwitchDeviceIdResult 
type UnlinkNintendoSwitchDeviceIdResultModel struct {
}

// UnlinkOpenIdConnectRequest 
type UnlinkOpenIdConnectRequestModel struct {
    // ConnectionId a name that identifies which configured OpenID Connect provider relationship to use. Maximum 100 characters.
    ConnectionId string `json:"ConnectionId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
}

// UnlinkPSNAccountRequest 
type UnlinkPSNAccountRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
}

// UnlinkPSNAccountResult 
type UnlinkPSNAccountResultModel struct {
}

// UnlinkSteamAccountRequest 
type UnlinkSteamAccountRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
}

// UnlinkSteamAccountResult 
type UnlinkSteamAccountResultModel struct {
}

// UnlinkTwitchAccountRequest 
type UnlinkTwitchAccountRequestModel struct {
    // AccessToken valid token issued by Twitch. Used to specify which twitch account to unlink from the profile. By default it uses the
// one that is present on the profile.
    AccessToken string `json:"AccessToken,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
}

// UnlinkTwitchAccountResult 
type UnlinkTwitchAccountResultModel struct {
}

// UnlinkWindowsHelloAccountRequest must include the Public Key Hint
type UnlinkWindowsHelloAccountRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // PublicKeyHint sHA256 hash of the PublicKey generated by Windows Hello.
    PublicKeyHint string `json:"PublicKeyHint,omitempty"`
}

// UnlinkWindowsHelloAccountResponse 
type UnlinkWindowsHelloAccountResponseModel struct {
}

// UnlinkXboxAccountRequest 
type UnlinkXboxAccountRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
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
}

// UnlockContainerItemResult the items and vc found within the container. These will be added and stacked in the appropriate inventory.
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
}

// UpdateCharacterDataRequest this function performs an additive update of the arbitrary strings containing the custom data for the character. In
// updating the custom data object, keys which already exist in the object will have their values overwritten, while keys
// with null values will be removed. New keys will be added, with the given values. No other key-value pairs will be
// changed apart from those specified in the call.
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
}

// UpdateCharacterDataResult 
type UpdateCharacterDataResultModel struct {
    // DataVersion indicates the current version of the data that has been set. This is incremented with every set call for that type of
// data (read-only, internal, etc). This version can be provided in Get calls to find updated data.
    DataVersion uint32 `json:"DataVersion,omitempty"`
}

// UpdateCharacterStatisticsRequest enable this option with the 'Allow Client to Post Player Statistics' option in PlayFab GameManager for your title.
// However, this is not best practice, as this data will no longer be safely controlled by the server. This operation is
// additive. Character Statistics not currently defined will be added, while those already defined will be updated with the
// given values. All other user statistics will remain unchanged. Character statistics are used by the
// character-leaderboard apis, and accessible for custom game-logic.
type UpdateCharacterStatisticsRequestModel struct {
    // CharacterId unique PlayFab assigned ID for a specific character owned by a user
    CharacterId string `json:"CharacterId,omitempty"`
    // CharacterStatistics statistics to be updated with the provided values, in the Key(string), Value(int) pattern.
    CharacterStatistics map[string]int32 `json:"CharacterStatistics,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
}

// UpdateCharacterStatisticsResult 
type UpdateCharacterStatisticsResultModel struct {
}

// UpdatePlayerStatisticsRequest enable this option with the 'Allow Client to Post Player Statistics' option in PlayFab GameManager for your title.
// However, this is not best practice, as this data will no longer be safely controlled by the server. This operation is
// additive. Statistics not currently defined will be added, while those already defined will be updated with the given
// values. All other user statistics will remain unchanged. Note that if the statistic is intended to have a reset period,
// the UpdatePlayerStatisticDefinition API call can be used to define that reset period. Once a statistic has been
// versioned (reset), the now-previous version can still be written to for up a short, pre-defined period (currently 10
// seconds), using the Version parameter in this call.
type UpdatePlayerStatisticsRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
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

// UpdateUserDataRequest this function performs an additive update of the arbitrary strings containing the custom data for the user. In updating
// the custom data object, keys which already exist in the object will have their values overwritten, while keys with null
// values will be removed. New keys will be added, with the given values. No other key-value pairs will be changed apart
// from those specified in the call.
type UpdateUserDataRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Data key-value pairs to be written to the custom data. Note that keys are trimmed of whitespace, are limited in size, and may
// not begin with a '!' character or be null.
    Data map[string]string `json:"Data,omitempty"`
    // KeysToRemove optional list of Data-keys to remove from UserData. Some SDKs cannot insert null-values into Data due to language
// constraints. Use this to delete the keys directly.
    KeysToRemove []string `json:"KeysToRemove,omitempty"`
    // Permission permission to be applied to all user data keys written in this request. Defaults to "private" if not set. This is used
// for requests by one player for information about another player; those requests will only return Public keys.
    Permission UserDataPermission `json:"Permission,omitempty"`
}

// UpdateUserDataResult 
type UpdateUserDataResultModel struct {
    // DataVersion indicates the current version of the data that has been set. This is incremented with every set call for that type of
// data (read-only, internal, etc). This version can be provided in Get calls to find updated data.
    DataVersion uint32 `json:"DataVersion,omitempty"`
}

// UpdateUserTitleDisplayNameRequest in addition to the PlayFab username, titles can make use of a DisplayName which is also a unique identifier, but
// specific to the title. This allows for unique names which more closely match the theme or genre of a title, for example.
type UpdateUserTitleDisplayNameRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // DisplayName new title display name for the user - must be between 3 and 25 characters.
    DisplayName string `json:"DisplayName,omitempty"`
}

// UpdateUserTitleDisplayNameResult 
type UpdateUserTitleDisplayNameResultModel struct {
    // DisplayName current title display name for the user (this will be the original display name if the rename attempt failed).
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

// ValidateAmazonReceiptRequest 
type ValidateAmazonReceiptRequestModel struct {
    // CatalogVersion catalog version of the fulfilled items. If null, defaults to the primary catalog.
    CatalogVersion string `json:"CatalogVersion,omitempty"`
    // CurrencyCode currency used to pay for the purchase (ISO 4217 currency code).
    CurrencyCode string `json:"CurrencyCode,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // PurchasePrice amount of the stated currency paid, in centesimal units.
    PurchasePrice int32 `json:"PurchasePrice,omitempty"`
    // ReceiptId receiptId returned by the Amazon App Store in-app purchase API
    ReceiptId string `json:"ReceiptId,omitempty"`
    // UserId amazonId of the user making the purchase as returned by the Amazon App Store in-app purchase API
    UserId string `json:"UserId,omitempty"`
}

// ValidateAmazonReceiptResult once verified, the catalog item matching the Amazon item name will be added to the user's inventory.
type ValidateAmazonReceiptResultModel struct {
    // Fulfillments fulfilled inventory items and recorded purchases in fulfillment of the validated receipt transactions.
    Fulfillments []PurchaseReceiptFulfillmentModel `json:"Fulfillments,omitempty"`
}

// ValidateGooglePlayPurchaseRequest the packageName and productId are defined in the GooglePlay store. The productId must match the ItemId of the inventory
// item in the PlayFab catalog for the title. This enables the PlayFab service to securely validate that the purchase is
// for the correct item, in order to prevent uses from passing valid receipts as being for more expensive items (passing a
// receipt for a 99-cent purchase as being for a $19.99 purchase, for example). Each receipt may be validated only once to
// avoid granting the same item over and over from a single purchase.
type ValidateGooglePlayPurchaseRequestModel struct {
    // CatalogVersion catalog version of the fulfilled items. If null, defaults to the primary catalog.
    CatalogVersion string `json:"CatalogVersion,omitempty"`
    // CurrencyCode currency used to pay for the purchase (ISO 4217 currency code).
    CurrencyCode string `json:"CurrencyCode,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // PurchasePrice amount of the stated currency paid, in centesimal units.
    PurchasePrice uint32 `json:"PurchasePrice,omitempty"`
    // ReceiptJson original JSON string returned by the Google Play IAB API.
    ReceiptJson string `json:"ReceiptJson,omitempty"`
    // Signature signature returned by the Google Play IAB API.
    Signature string `json:"Signature,omitempty"`
}

// ValidateGooglePlayPurchaseResult once verified, the catalog item (ItemId) matching the GooglePlay store item (productId) will be added to the user's
// inventory.
type ValidateGooglePlayPurchaseResultModel struct {
    // Fulfillments fulfilled inventory items and recorded purchases in fulfillment of the validated receipt transactions.
    Fulfillments []PurchaseReceiptFulfillmentModel `json:"Fulfillments,omitempty"`
}

// ValidateIOSReceiptRequest the CurrencyCode and PurchasePrice must match the price which was set up for the item in the Apple store. In addition,
// The ItemId of the inventory in the PlayFab Catalog must match the Product ID as it was set up in the Apple store. This
// enables the PlayFab service to securely validate that the purchase is for the correct item, in order to prevent uses
// from passing valid receipts as being for more expensive items (passing a receipt for a 99-cent purchase as being for a
// $19.99 purchase, for example).
type ValidateIOSReceiptRequestModel struct {
    // CatalogVersion catalog version of the fulfilled items. If null, defaults to the primary catalog.
    CatalogVersion string `json:"CatalogVersion,omitempty"`
    // CurrencyCode currency used to pay for the purchase (ISO 4217 currency code).
    CurrencyCode string `json:"CurrencyCode,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // PurchasePrice amount of the stated currency paid, in centesimal units.
    PurchasePrice int32 `json:"PurchasePrice,omitempty"`
    // ReceiptData base64 encoded receipt data, passed back by the App Store as a result of a successful purchase.
    ReceiptData string `json:"ReceiptData,omitempty"`
}

// ValidateIOSReceiptResult once verified, the catalog item matching the iTunes item name will be added to the user's inventory.
type ValidateIOSReceiptResultModel struct {
    // Fulfillments fulfilled inventory items and recorded purchases in fulfillment of the validated receipt transactions.
    Fulfillments []PurchaseReceiptFulfillmentModel `json:"Fulfillments,omitempty"`
}

// ValidateWindowsReceiptRequest 
type ValidateWindowsReceiptRequestModel struct {
    // CatalogVersion catalog version of the fulfilled items. If null, defaults to the primary catalog.
    CatalogVersion string `json:"CatalogVersion,omitempty"`
    // CurrencyCode currency used to pay for the purchase (ISO 4217 currency code).
    CurrencyCode string `json:"CurrencyCode,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // PurchasePrice amount of the stated currency paid, in centesimal units.
    PurchasePrice uint32 `json:"PurchasePrice,omitempty"`
    // Receipt xML Receipt returned by the Windows App Store in-app purchase API
    Receipt string `json:"Receipt,omitempty"`
}

// ValidateWindowsReceiptResult once verified, the catalog item matching the Product name will be added to the user's inventory.
type ValidateWindowsReceiptResultModel struct {
    // Fulfillments fulfilled inventory items and recorded purchases in fulfillment of the validated receipt transactions.
    Fulfillments []PurchaseReceiptFulfillmentModel `json:"Fulfillments,omitempty"`
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

// WriteClientCharacterEventRequest this API is designed to write a multitude of different client-defined events into PlayStream. It supports a flexible
// JSON schema, which allowsfor arbitrary key-value pairs to describe any character-based event. The created event will be
// locked to the authenticated title and player.
type WriteClientCharacterEventRequestModel struct {
    // Body custom event properties. Each property consists of a name (string) and a value (JSON object).
    Body map[string]interface{} `json:"Body,omitempty"`
    // CharacterId unique PlayFab assigned ID for a specific character owned by a user
    CharacterId string `json:"CharacterId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // EventName the name of the event, within the namespace scoped to the title. The naming convention is up to the caller, but it
// commonly follows the subject_verb_object pattern (e.g. player_logged_in).
    EventName string `json:"EventName,omitempty"`
    // Timestamp the time (in UTC) associated with this event. The value defaults to the current time.
    Timestamp time.Time `json:"Timestamp,omitempty"`
}

// WriteClientPlayerEventRequest this API is designed to write a multitude of different event types into PlayStream. It supports a flexible JSON schema,
// which allowsfor arbitrary key-value pairs to describe any player-based event. The created event will be locked to the
// authenticated title and player.
type WriteClientPlayerEventRequestModel struct {
    // Body custom data properties associated with the event. Each property consists of a name (string) and a value (JSON object).
    Body map[string]interface{} `json:"Body,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // EventName the name of the event, within the namespace scoped to the title. The naming convention is up to the caller, but it
// commonly follows the subject_verb_object pattern (e.g. player_logged_in).
    EventName string `json:"EventName,omitempty"`
    // Timestamp the time (in UTC) associated with this event. The value defaults to the current time.
    Timestamp time.Time `json:"Timestamp,omitempty"`
}

// WriteEventResponse 
type WriteEventResponseModel struct {
    // EventId the unique identifier of the event. The values of this identifier consist of ASCII characters and are not constrained to
// any particular format.
    EventId string `json:"EventId,omitempty"`
}

// WriteTitleEventRequest this API is designed to write a multitude of different client-defined events into PlayStream. It supports a flexible
// JSON schema, which allowsfor arbitrary key-value pairs to describe any title-based event. The created event will be
// locked to the authenticated title.
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
