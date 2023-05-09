package economy

import "time"
                    
// AddInventoryItemsOperation 
type AddInventoryItemsOperationModel struct {
    // Amount the amount to add to the current item amount.
    Amount int32 `json:"Amount,omitempty"`
    // DurationInSeconds the duration to add to the current item expiration date.
    DurationInSeconds float64 `json:"DurationInSeconds,omitempty"`
    // Item the inventory item the operation applies to.
    Item *InventoryItemReferenceModel `json:"Item,omitempty"`
    // NewStackValues the values to apply to a stack newly created by this operation.
    NewStackValues *InitialValuesModel `json:"NewStackValues,omitempty"`
}

// AddInventoryItemsRequest given an entity type, entity identifier and container details, will add the specified inventory items.
type AddInventoryItemsRequestModel struct {
    // Amount the amount to add for the current item.
    Amount int32 `json:"Amount,omitempty"`
    // CollectionId the id of the entity's collection to perform this action on. (Default="default"). The number of inventory collections is
// unlimited.
    CollectionId string `json:"CollectionId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // DurationInSeconds the duration to add to the current item expiration date.
    DurationInSeconds float64 `json:"DurationInSeconds,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // ETag eTags are used for concurrency checking when updating resources. More information about using ETags can be found here:
// https://learn.microsoft.com/en-us/gaming/playfab/features/economy-v2/catalog/etags
    ETag string `json:"ETag,omitempty"`
    // IdempotencyId the Idempotency ID for this request. Idempotency IDs can be used to prevent operation replay in the medium term but will
// be garbage collected eventually.
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // Item the inventory item the request applies to.
    Item *InventoryItemReferenceModel `json:"Item,omitempty"`
    // NewStackValues the values to apply to a stack newly created by this request.
    NewStackValues *InitialValuesModel `json:"NewStackValues,omitempty"`
}

// AddInventoryItemsResponse 
type AddInventoryItemsResponseModel struct {
    // ETag eTags are used for concurrency checking when updating resources. More information about using ETags can be found here:
// https://learn.microsoft.com/en-us/gaming/playfab/features/economy-v2/catalog/etags
    ETag string `json:"ETag,omitempty"`
    // IdempotencyId the idempotency id used in the request.
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // TransactionIds the ids of transactions that occurred as a result of the request.
    TransactionIds []string `json:"TransactionIds,omitempty"`
}

// AlternateId 
type AlternateIdModel struct {
    // Type type of the alternate ID.
    Type string `json:"Type,omitempty"`
    // Value value of the alternate ID.
    Value string `json:"Value,omitempty"`
}

// CatalogAlternateId 
type CatalogAlternateIdModel struct {
    // Type type of the alternate ID.
    Type string `json:"Type,omitempty"`
    // Value value of the alternate ID.
    Value string `json:"Value,omitempty"`
}

// CatalogConfig 
type CatalogConfigModel struct {
    // AdminEntities a list of player entity keys that will have admin permissions. There is a maximum of 64 entities that can be added.
    AdminEntities []EntityKeyModel `json:"AdminEntities,omitempty"`
    // Catalog the set of configuration that only applies to catalog items.
    Catalog *CatalogSpecificConfigModel `json:"Catalog,omitempty"`
    // DeepLinkFormats a list of deep link formats. Up to 10 can be added.
    DeepLinkFormats []DeepLinkFormatModel `json:"DeepLinkFormats,omitempty"`
    // DisplayPropertyIndexInfos a list of display properties to index. Up to 5 mappings can be added per Display Property Type. More info on display
// properties can be found here:
// https://learn.microsoft.com/en-us/gaming/playfab/features/economy-v2/catalog/content-types-tags-and-properties#displayproperties
    DisplayPropertyIndexInfos []DisplayPropertyIndexInfoModel `json:"DisplayPropertyIndexInfos,omitempty"`
    // File the set of configuration that only applies to Files.
    File *FileConfigModel `json:"File,omitempty"`
    // Image the set of configuration that only applies to Images.
    Image *ImageConfigModel `json:"Image,omitempty"`
    // IsCatalogEnabled flag defining whether catalog is enabled.
    IsCatalogEnabled bool `json:"IsCatalogEnabled"`
    // Platforms a list of Platforms that can be applied to catalog items. Each platform can have a maximum character length of 40 and up
// to 128 platforms can be listed.
    Platforms []string `json:"Platforms,omitempty"`
    // ReviewerEntities a set of player entity keys that are allowed to review content. There is a maximum of 64 entities that can be added.
    ReviewerEntities []EntityKeyModel `json:"ReviewerEntities,omitempty"`
    // UserGeneratedContent the set of configuration that only applies to user generated contents.
    UserGeneratedContent *UserGeneratedContentSpecificConfigModel `json:"UserGeneratedContent,omitempty"`
}

// CatalogItem 
type CatalogItemModel struct {
    // AlternateIds the alternate IDs associated with this item. An alternate ID can be set to 'FriendlyId' or any of the supported
// marketplace names.
    AlternateIds []CatalogAlternateIdModel `json:"AlternateIds,omitempty"`
    // Contents the set of content/files associated with this item. Up to 100 files can be added to an item.
    Contents []ContentModel `json:"Contents,omitempty"`
    // ContentType the client-defined type of the item.
    ContentType string `json:"ContentType,omitempty"`
    // CreationDate the date and time when this item was created.
    CreationDate time.Time `json:"CreationDate,omitempty"`
    // CreatorEntity the ID of the creator of this catalog item.
    CreatorEntity *EntityKeyModel `json:"CreatorEntity,omitempty"`
    // DeepLinks the set of platform specific deep links for this item.
    DeepLinks []DeepLinkModel `json:"DeepLinks,omitempty"`
    // DefaultStackId the Stack Id that will be used as default for this item in Inventory when an explicit one is not provided. This
// DefaultStackId can be a static stack id or '{guid}', which will generate a unique stack id for the item. If null,
// Inventory's default stack id will be used.
    DefaultStackId string `json:"DefaultStackId,omitempty"`
    // Description a dictionary of localized descriptions. Key is language code and localized string is the value. The NEUTRAL locale is
// required. Descriptions have a 10000 character limit per country code.
    Description map[string]string `json:"Description,omitempty"`
    // DisplayProperties game specific properties for display purposes. This is an arbitrary JSON blob. The Display Properties field has a 10000
// byte limit per item.
    DisplayProperties interface{} `json:"DisplayProperties,omitempty"`
    // DisplayVersion the user provided version of the item for display purposes. Maximum character length of 50.
    DisplayVersion string `json:"DisplayVersion,omitempty"`
    // EndDate the date of when the item will cease to be available. If not provided then the product will be available indefinitely.
    EndDate time.Time `json:"EndDate,omitempty"`
    // ETag the current ETag value that can be used for optimistic concurrency in the If-None-Match header.
    ETag string `json:"ETag,omitempty"`
    // Id the unique ID of the item.
    Id string `json:"Id,omitempty"`
    // Images the images associated with this item. Images can be thumbnails or screenshots. Up to 100 images can be added to an item.
// Only .png, .jpg, .gif, and .bmp file types can be uploaded
    Images []ImageModel `json:"Images,omitempty"`
    // IsHidden indicates if the item is hidden.
    IsHidden bool `json:"IsHidden"`
    // ItemReferences the item references associated with this item. For example, the items in a Bundle/Store/Subscription. Every item can
// have up to 50 item references.
    ItemReferences []CatalogItemReferenceModel `json:"ItemReferences,omitempty"`
    // Keywords a dictionary of localized keywords. Key is language code and localized list of keywords is the value. Keywords have a 50
// character limit per keyword and up to 32 keywords can be added per country code.
    Keywords map[string]KeywordSetModel `json:"Keywords,omitempty"`
    // LastModifiedDate the date and time this item was last updated.
    LastModifiedDate time.Time `json:"LastModifiedDate,omitempty"`
    // Moderation the moderation state for this item.
    Moderation *ModerationStateModel `json:"Moderation,omitempty"`
    // Platforms the platforms supported by this item.
    Platforms []string `json:"Platforms,omitempty"`
    // PriceOptions the prices the item can be purchased for.
    PriceOptions *CatalogPriceOptionsModel `json:"PriceOptions,omitempty"`
    // Rating rating summary for this item.
    Rating *RatingModel `json:"Rating,omitempty"`
    // StartDate the date of when the item will be available. If not provided then the product will appear immediately.
    StartDate time.Time `json:"StartDate,omitempty"`
    // StoreDetails optional details for stores items.
    StoreDetails *StoreDetailsModel `json:"StoreDetails,omitempty"`
    // Tags the list of tags that are associated with this item. Up to 32 tags can be added to an item.
    Tags []string `json:"Tags,omitempty"`
    // Title a dictionary of localized titles. Key is language code and localized string is the value. The NEUTRAL locale is
// required. Titles have a 512 character limit per country code.
    Title map[string]string `json:"Title,omitempty"`
    // Type the high-level type of the item. The following item types are supported: bundle, catalogItem, currency, store, ugc.
    Type string `json:"Type,omitempty"`
}

// CatalogItemReference 
type CatalogItemReferenceModel struct {
    // Amount the amount of the catalog item.
    Amount int32 `json:"Amount,omitempty"`
    // Id the unique ID of the catalog item.
    Id string `json:"Id,omitempty"`
    // PriceOptions the prices the catalog item can be purchased for.
    PriceOptions *CatalogPriceOptionsModel `json:"PriceOptions,omitempty"`
}

// CatalogPrice 
type CatalogPriceModel struct {
    // Amounts the amounts of the catalog item price. Each price can have up to 15 item amounts.
    Amounts []CatalogPriceAmountModel `json:"Amounts,omitempty"`
    // UnitDurationInSeconds the per-unit duration this price can be used to purchase. The maximum duration is 100 years.
    UnitDurationInSeconds float64 `json:"UnitDurationInSeconds,omitempty"`
}

// CatalogPriceAmount 
type CatalogPriceAmountModel struct {
    // Amount the amount of the price.
    Amount int32 `json:"Amount,omitempty"`
    // ItemId the Item Id of the price.
    ItemId string `json:"ItemId,omitempty"`
}

// CatalogPriceAmountOverride 
type CatalogPriceAmountOverrideModel struct {
    // FixedValue the exact value that should be utilized in the override.
    FixedValue int32 `json:"FixedValue,omitempty"`
    // ItemId the id of the item this override should utilize.
    ItemId string `json:"ItemId,omitempty"`
    // Multiplier the multiplier that will be applied to the base Catalog value to determine what value should be utilized in the
// override.
    Multiplier float64 `json:"Multiplier,omitempty"`
}

// CatalogPriceOptions 
type CatalogPriceOptionsModel struct {
    // Prices prices of the catalog item. An item can have up to 15 prices
    Prices []CatalogPriceModel `json:"Prices,omitempty"`
}

// CatalogPriceOptionsOverride 
type CatalogPriceOptionsOverrideModel struct {
    // Prices the prices utilized in the override.
    Prices []CatalogPriceOverrideModel `json:"Prices,omitempty"`
}

// CatalogPriceOverride 
type CatalogPriceOverrideModel struct {
    // Amounts the currency amounts utilized in the override for a singular price.
    Amounts []CatalogPriceAmountOverrideModel `json:"Amounts,omitempty"`
}

// CatalogSpecificConfig 
type CatalogSpecificConfigModel struct {
    // ContentTypes the set of content types that will be used for validation. Each content type can have a maximum character length of 40
// and up to 128 types can be listed.
    ContentTypes []string `json:"ContentTypes,omitempty"`
    // Tags the set of tags that will be used for validation. Each tag can have a maximum character length of 32 and up to 1024 tags
// can be listed.
    Tags []string `json:"Tags,omitempty"`
}

// ConcernCategory 
type ConcernCategory string
  
const (
     ConcernCategoryNone ConcernCategory = "None"
     ConcernCategoryOffensiveContent ConcernCategory = "OffensiveContent"
     ConcernCategoryChildExploitation ConcernCategory = "ChildExploitation"
     ConcernCategoryMalwareOrVirus ConcernCategory = "MalwareOrVirus"
     ConcernCategoryPrivacyConcerns ConcernCategory = "PrivacyConcerns"
     ConcernCategoryMisleadingApp ConcernCategory = "MisleadingApp"
     ConcernCategoryPoorPerformance ConcernCategory = "PoorPerformance"
     ConcernCategoryReviewResponse ConcernCategory = "ReviewResponse"
     ConcernCategorySpamAdvertising ConcernCategory = "SpamAdvertising"
     ConcernCategoryProfanity ConcernCategory = "Profanity"
      )
// Content 
type ContentModel struct {
    // Id the content unique ID.
    Id string `json:"Id,omitempty"`
    // MaxClientVersion the maximum client version that this content is compatible with. Client Versions can be up to 3 segments separated by
// periods(.) and each segment can have a maximum value of 65535.
    MaxClientVersion string `json:"MaxClientVersion,omitempty"`
    // MinClientVersion the minimum client version that this content is compatible with. Client Versions can be up to 3 segments separated by
// periods(.) and each segment can have a maximum value of 65535.
    MinClientVersion string `json:"MinClientVersion,omitempty"`
    // Tags the list of tags that are associated with this content. Tags must be defined in the Catalog Config before being used in
// content.
    Tags []string `json:"Tags,omitempty"`
    // Type the client-defined type of the content. Content Types must be defined in the Catalog Config before being used.
    Type string `json:"Type,omitempty"`
    // Url the Azure CDN URL for retrieval of the catalog item binary content.
    Url string `json:"Url,omitempty"`
}

// ContentFeed 
type ContentFeedModel struct {
}

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
// CreateDraftItemRequest the item will not be published to the public catalog until the PublishItem API is called for the item.
type CreateDraftItemRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Item metadata describing the new catalog item to be created.
    Item *CatalogItemModel `json:"Item,omitempty"`
    // Publish whether the item should be published immediately. This value is optional, defaults to false.
    Publish bool `json:"Publish"`
}

// CreateDraftItemResponse 
type CreateDraftItemResponseModel struct {
    // Item updated metadata describing the catalog item just created.
    Item *CatalogItemModel `json:"Item,omitempty"`
}

// CreateUploadUrlsRequest upload URLs point to Azure Blobs; clients must follow the Microsoft Azure Storage Blob Service REST API pattern for
// uploading content. The response contains upload URLs and IDs for each file. The IDs and URLs returned must be added to
// the item metadata and committed using the CreateDraftItem or UpdateDraftItem Item APIs.
type CreateUploadUrlsRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Files description of the files to be uploaded by the client.
    Files []UploadInfoModel `json:"Files,omitempty"`
}

// CreateUploadUrlsResponse 
type CreateUploadUrlsResponseModel struct {
    // UploadUrls list of URLs metadata for the files to be uploaded by the client.
    UploadUrls []UploadUrlMetadataModel `json:"UploadUrls,omitempty"`
}

// DeepLink 
type DeepLinkModel struct {
    // Platform target platform for this deep link.
    Platform string `json:"Platform,omitempty"`
    // Url the deep link for this platform.
    Url string `json:"Url,omitempty"`
}

// DeepLinkFormat 
type DeepLinkFormatModel struct {
    // Format the format of the deep link to return. The format should contain '{id}' to represent where the item ID should be placed.
    Format string `json:"Format,omitempty"`
    // Platform the target platform for the deep link.
    Platform string `json:"Platform,omitempty"`
}

// DeleteEntityItemReviewsRequest 
type DeleteEntityItemReviewsRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
}

// DeleteEntityItemReviewsResponse 
type DeleteEntityItemReviewsResponseModel struct {
}

// DeleteInventoryCollectionRequest delete an Inventory Collection by the specified Id for an Entity
type DeleteInventoryCollectionRequestModel struct {
    // CollectionId the inventory collection id the request applies to.
    CollectionId string `json:"CollectionId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity the request is about. Set to the caller by default.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // ETag eTags are used for concurrency checking when updating resources. More information about using ETags can be found here:
// https://learn.microsoft.com/en-us/gaming/playfab/features/economy-v2/catalog/etags
    ETag string `json:"ETag,omitempty"`
}

// DeleteInventoryCollectionResponse 
type DeleteInventoryCollectionResponseModel struct {
}

// DeleteInventoryItemsOperation 
type DeleteInventoryItemsOperationModel struct {
    // Item the inventory item the operation applies to.
    Item *InventoryItemReferenceModel `json:"Item,omitempty"`
}

// DeleteInventoryItemsRequest given an entity type, entity identifier and container details, will delete the entity's inventory items
type DeleteInventoryItemsRequestModel struct {
    // CollectionId the id of the entity's collection to perform this action on. (Default="default"). The number of inventory collections is
// unlimited.
    CollectionId string `json:"CollectionId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // ETag eTags are used for concurrency checking when updating resources. More information about using ETags can be found here:
// https://learn.microsoft.com/en-us/gaming/playfab/features/economy-v2/catalog/etags
    ETag string `json:"ETag,omitempty"`
    // IdempotencyId the Idempotency ID for this request. Idempotency IDs can be used to prevent operation replay in the medium term but will
// be garbage collected eventually.
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // Item the inventory item the request applies to.
    Item *InventoryItemReferenceModel `json:"Item,omitempty"`
}

// DeleteInventoryItemsResponse 
type DeleteInventoryItemsResponseModel struct {
    // ETag eTags are used for concurrency checking when updating resources.
    ETag string `json:"ETag,omitempty"`
    // IdempotencyId the idempotency id used in the request.
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // TransactionIds the ids of transactions that occurred as a result of the request.
    TransactionIds []string `json:"TransactionIds,omitempty"`
}

// DeleteItemRequest 
type DeleteItemRequestModel struct {
    // AlternateId an alternate ID associated with this item.
    AlternateId *CatalogAlternateIdModel `json:"AlternateId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // Id the unique ID of the item.
    Id string `json:"Id,omitempty"`
}

// DeleteItemResponse 
type DeleteItemResponseModel struct {
}

// DisplayPropertyIndexInfo 
type DisplayPropertyIndexInfoModel struct {
    // Name the property name in the 'DisplayProperties' property to be indexed.
    Name string `json:"Name,omitempty"`
    // Type the type of the property to be indexed.
    Type DisplayPropertyType `json:"Type,omitempty"`
}

// DisplayPropertyType 
type DisplayPropertyType string
  
const (
     DisplayPropertyTypeNone DisplayPropertyType = "None"
     DisplayPropertyTypeQueryDateTime DisplayPropertyType = "QueryDateTime"
     DisplayPropertyTypeQueryDouble DisplayPropertyType = "QueryDouble"
     DisplayPropertyTypeQueryString DisplayPropertyType = "QueryString"
     DisplayPropertyTypeSearchString DisplayPropertyType = "SearchString"
      )
// EntityKey combined entity type and ID structure which uniquely identifies a single entity.
type EntityKeyModel struct {
    // Id unique ID of the entity.
    Id string `json:"Id,omitempty"`
    // Type entity type. See https://docs.microsoft.com/gaming/playfab/features/data/entities/available-built-in-entity-types
    Type string `json:"Type,omitempty"`
}

// ExecuteInventoryOperationsRequest execute a list of Inventory Operations for an Entity
type ExecuteInventoryOperationsRequestModel struct {
    // CollectionId the id of the entity's collection to perform this action on. (Default="default"). The number of inventory collections is
// unlimited.
    CollectionId string `json:"CollectionId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // ETag eTags are used for concurrency checking when updating resources. More information about using ETags can be found here:
// https://learn.microsoft.com/en-us/gaming/playfab/features/economy-v2/catalog/etags
    ETag string `json:"ETag,omitempty"`
    // IdempotencyId the Idempotency ID for this request. Idempotency IDs can be used to prevent operation replay in the medium term but will
// be garbage collected eventually.
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // Operations the operations to run transactionally. The operations will be executed in-order sequentially and will succeed or fail as
// a batch. Up to 10 operations can be added.
    Operations []InventoryOperationModel `json:"Operations,omitempty"`
}

// ExecuteInventoryOperationsResponse 
type ExecuteInventoryOperationsResponseModel struct {
    // ETag eTags are used for concurrency checking when updating resources. More information about using ETags can be found here:
// https://learn.microsoft.com/en-us/gaming/playfab/features/economy-v2/catalog/etags
    ETag string `json:"ETag,omitempty"`
    // IdempotencyId the idempotency id used in the request.
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // TransactionIds the ids of the transactions that occurred as a result of the request.
    TransactionIds []string `json:"TransactionIds,omitempty"`
}

// FileConfig 
type FileConfigModel struct {
    // ContentTypes the set of content types that will be used for validation. Each content type can have a maximum character length of 40
// and up to 128 types can be listed.
    ContentTypes []string `json:"ContentTypes,omitempty"`
    // Tags the set of tags that will be used for validation. Each tag can have a maximum character length of 32 and up to 1024 tags
// can be listed.
    Tags []string `json:"Tags,omitempty"`
}

// FilterOptions 
type FilterOptionsModel struct {
    // Filter the OData filter utilized. Mutually exclusive with 'IncludeAllItems'. More info about Filter Complexity limits can be
// found here: https://learn.microsoft.com/en-us/gaming/playfab/features/economy-v2/catalog/search#limits
    Filter string `json:"Filter,omitempty"`
    // IncludeAllItems the flag that overrides the filter and allows for returning all catalog items. Mutually exclusive with 'Filter'.
    IncludeAllItems bool `json:"IncludeAllItems"`
}

// GetCatalogConfigRequest 
type GetCatalogConfigRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
}

// GetCatalogConfigResponse 
type GetCatalogConfigResponseModel struct {
    // Config the catalog configuration.
    Config *CatalogConfigModel `json:"Config,omitempty"`
}

// GetDraftItemRequest 
type GetDraftItemRequestModel struct {
    // AlternateId an alternate ID associated with this item.
    AlternateId *CatalogAlternateIdModel `json:"AlternateId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // Id the unique ID of the item.
    Id string `json:"Id,omitempty"`
}

// GetDraftItemResponse 
type GetDraftItemResponseModel struct {
    // Item full metadata of the catalog item requested.
    Item *CatalogItemModel `json:"Item,omitempty"`
}

// GetDraftItemsRequest 
type GetDraftItemsRequestModel struct {
    // AlternateIds list of item alternate IDs.
    AlternateIds []CatalogAlternateIdModel `json:"AlternateIds,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // Ids list of Item Ids.
    Ids []string `json:"Ids,omitempty"`
}

// GetDraftItemsResponse 
type GetDraftItemsResponseModel struct {
    // ContinuationToken an opaque token used to retrieve the next page of items, if any are available.
    ContinuationToken string `json:"ContinuationToken,omitempty"`
    // Items a set of items created by the entity.
    Items []CatalogItemModel `json:"Items,omitempty"`
}

// GetEntityDraftItemsRequest 
type GetEntityDraftItemsRequestModel struct {
    // ContinuationToken an opaque token used to retrieve the next page of items created by the caller, if any are available. Should be null on
// initial request.
    ContinuationToken string `json:"ContinuationToken,omitempty"`
    // Count number of items to retrieve. This value is optional. Default value is 10.
    Count int32 `json:"Count,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // Filter oData Filter to refine the items returned. CatalogItem properties 'type' can be used in the filter. For example: "type
// eq 'ugc'"
    Filter string `json:"Filter,omitempty"`
}

// GetEntityDraftItemsResponse 
type GetEntityDraftItemsResponseModel struct {
    // ContinuationToken an opaque token used to retrieve the next page of items, if any are available.
    ContinuationToken string `json:"ContinuationToken,omitempty"`
    // Items a set of items created by the entity.
    Items []CatalogItemModel `json:"Items,omitempty"`
}

// GetEntityItemReviewRequest 
type GetEntityItemReviewRequestModel struct {
    // AlternateId an alternate ID associated with this item.
    AlternateId *CatalogAlternateIdModel `json:"AlternateId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // Id the unique ID of the item.
    Id string `json:"Id,omitempty"`
}

// GetEntityItemReviewResponse 
type GetEntityItemReviewResponseModel struct {
    // Review the review the entity submitted for the requested item.
    Review *ReviewModel `json:"Review,omitempty"`
}

// GetInventoryCollectionIdsRequest get a list of Inventory Collection Ids for the specified Entity
type GetInventoryCollectionIdsRequestModel struct {
    // ContinuationToken an opaque token used to retrieve the next page of collection ids, if any are available.
    ContinuationToken string `json:"ContinuationToken,omitempty"`
    // Count number of items to retrieve. This value is optional. The default value is 10
    Count int32 `json:"Count,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity the request is about. Set to the caller by default.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
}

// GetInventoryCollectionIdsResponse 
type GetInventoryCollectionIdsResponseModel struct {
    // CollectionIds the requested inventory collection ids.
    CollectionIds []string `json:"CollectionIds,omitempty"`
    // ContinuationToken an opaque token used to retrieve the next page of collection ids, if any are available.
    ContinuationToken string `json:"ContinuationToken,omitempty"`
}

// GetInventoryItemsRequest given an entity type, entity identifier and container details, will get the entity's inventory items.
type GetInventoryItemsRequestModel struct {
    // CollectionId the id of the entity's collection to perform this action on. (Default="default")
    CollectionId string `json:"CollectionId,omitempty"`
    // ContinuationToken an opaque token used to retrieve the next page of items in the inventory, if any are available. Should be null on
// initial request.
    ContinuationToken string `json:"ContinuationToken,omitempty"`
    // Count number of items to retrieve. This value is optional. Maximum page size is 50. The default value is 10
    Count int32 `json:"Count,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // Filter oData Filter to refine the items returned. InventoryItem properties 'type', 'id', and 'stackId' can be used in the
// filter. For example: "type eq 'currency'"
    Filter string `json:"Filter,omitempty"`
}

// GetInventoryItemsResponse 
type GetInventoryItemsResponseModel struct {
    // ContinuationToken an opaque token used to retrieve the next page of items, if any are available.
    ContinuationToken string `json:"ContinuationToken,omitempty"`
    // ETag eTags are used for concurrency checking when updating resources. More information about using ETags can be found here:
// https://learn.microsoft.com/en-us/gaming/playfab/features/economy-v2/catalog/etags
    ETag string `json:"ETag,omitempty"`
    // Items the requested inventory items.
    Items []InventoryItemModel `json:"Items,omitempty"`
}

// GetItemContainersRequest given an item, return a set of bundles and stores containing the item.
type GetItemContainersRequestModel struct {
    // AlternateId an alternate ID associated with this item.
    AlternateId *CatalogAlternateIdModel `json:"AlternateId,omitempty"`
    // ContinuationToken an opaque token used to retrieve the next page of items in the inventory, if any are available. Should be null on
// initial request.
    ContinuationToken string `json:"ContinuationToken,omitempty"`
    // Count number of items to retrieve. This value is optional. Default value is 10.
    Count int32 `json:"Count,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // Id the unique ID of the item.
    Id string `json:"Id,omitempty"`
}

// GetItemContainersResponse 
type GetItemContainersResponseModel struct {
    // Containers list of Bundles and Stores containing the requested items.
    Containers []CatalogItemModel `json:"Containers,omitempty"`
    // ContinuationToken an opaque token used to retrieve the next page of items, if any are available.
    ContinuationToken string `json:"ContinuationToken,omitempty"`
}

// GetItemModerationStateRequest 
type GetItemModerationStateRequestModel struct {
    // AlternateId an alternate ID associated with this item.
    AlternateId *CatalogAlternateIdModel `json:"AlternateId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Id the unique ID of the item.
    Id string `json:"Id,omitempty"`
}

// GetItemModerationStateResponse 
type GetItemModerationStateResponseModel struct {
    // State the current moderation state for the requested item.
    State *ModerationStateModel `json:"State,omitempty"`
}

// GetItemPublishStatusRequest 
type GetItemPublishStatusRequestModel struct {
    // AlternateId an alternate ID associated with this item.
    AlternateId *CatalogAlternateIdModel `json:"AlternateId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // Id the unique ID of the item.
    Id string `json:"Id,omitempty"`
}

// GetItemPublishStatusResponse 
type GetItemPublishStatusResponseModel struct {
    // Result high level status of the published item.
    Result PublishResult `json:"Result,omitempty"`
    // StatusMessage descriptive message about the current status of the publish.
    StatusMessage string `json:"StatusMessage,omitempty"`
}

// GetItemRequest 
type GetItemRequestModel struct {
    // AlternateId an alternate ID associated with this item.
    AlternateId *CatalogAlternateIdModel `json:"AlternateId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // Id the unique ID of the item.
    Id string `json:"Id,omitempty"`
}

// GetItemResponse get item result.
type GetItemResponseModel struct {
    // Item the item result.
    Item *CatalogItemModel `json:"Item,omitempty"`
}

// GetItemReviewsRequest 
type GetItemReviewsRequestModel struct {
    // AlternateId an alternate ID associated with this item.
    AlternateId *CatalogAlternateIdModel `json:"AlternateId,omitempty"`
    // ContinuationToken an opaque token used to retrieve the next page of items, if any are available.
    ContinuationToken string `json:"ContinuationToken,omitempty"`
    // Count number of items to retrieve. This value is optional. Default value is 10.
    Count int32 `json:"Count,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Id the unique ID of the item.
    Id string `json:"Id,omitempty"`
    // OrderBy an OData orderBy used to order the results of the query. Possible values are Helpfulness, Rating, and Submitted (For
// example: "Submitted desc")
    OrderBy string `json:"OrderBy,omitempty"`
}

// GetItemReviewsResponse 
type GetItemReviewsResponseModel struct {
    // ContinuationToken an opaque token used to retrieve the next page of items, if any are available.
    ContinuationToken string `json:"ContinuationToken,omitempty"`
    // Reviews the paginated set of results.
    Reviews []ReviewModel `json:"Reviews,omitempty"`
}

// GetItemReviewSummaryRequest 
type GetItemReviewSummaryRequestModel struct {
    // AlternateId an alternate ID associated with this item.
    AlternateId *CatalogAlternateIdModel `json:"AlternateId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Id the unique ID of the item.
    Id string `json:"Id,omitempty"`
}

// GetItemReviewSummaryResponse 
type GetItemReviewSummaryResponseModel struct {
    // LeastFavorableReview the least favorable review for this item.
    LeastFavorableReview *ReviewModel `json:"LeastFavorableReview,omitempty"`
    // MostFavorableReview the most favorable review for this item.
    MostFavorableReview *ReviewModel `json:"MostFavorableReview,omitempty"`
    // Rating the summary of ratings associated with this item.
    Rating *RatingModel `json:"Rating,omitempty"`
    // ReviewsCount the total number of reviews associated with this item.
    ReviewsCount int32 `json:"ReviewsCount,omitempty"`
}

// GetItemsRequest 
type GetItemsRequestModel struct {
    // AlternateIds list of item alternate IDs.
    AlternateIds []CatalogAlternateIdModel `json:"AlternateIds,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // Ids list of Item Ids.
    Ids []string `json:"Ids,omitempty"`
}

// GetItemsResponse 
type GetItemsResponseModel struct {
    // Items metadata of set of items.
    Items []CatalogItemModel `json:"Items,omitempty"`
}

// GetMicrosoftStoreAccessTokensRequest gets the access tokens for Microsoft Store authentication.
type GetMicrosoftStoreAccessTokensRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
}

// GetMicrosoftStoreAccessTokensResponse 
type GetMicrosoftStoreAccessTokensResponseModel struct {
    // CollectionsAccessToken the collections access token for calling https://onestore.microsoft.com/b2b/keys/create/collections to obtain a
// CollectionsIdKey for the user
    CollectionsAccessToken string `json:"CollectionsAccessToken,omitempty"`
    // CollectionsAccessTokenExpirationDate the date the collections access token expires
    CollectionsAccessTokenExpirationDate time.Time `json:"CollectionsAccessTokenExpirationDate,omitempty"`
}

// GetTransactionHistoryRequest get transaction history for specified entity and collection.
type GetTransactionHistoryRequestModel struct {
    // CollectionId the id of the entity's collection to perform this action on. (Default="default")
    CollectionId string `json:"CollectionId,omitempty"`
    // ContinuationToken an opaque token used to retrieve the next page of items, if any are available. Should be null on initial request.
    ContinuationToken string `json:"ContinuationToken,omitempty"`
    // Count number of items to retrieve. This value is optional. The default value is 10
    Count int32 `json:"Count,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // Filter an OData filter used to refine the TransactionHistory. Transaction property 'timestamp' can be used in the filter. For
// example: "timestamp ge 'timestamp ge'" By default, a 6 month timespan from the current date is used.
    Filter string `json:"Filter,omitempty"`
}

// GetTransactionHistoryResponse 
type GetTransactionHistoryResponseModel struct {
    // ContinuationToken an opaque token used to retrieve the next page of items, if any are available. Should be null on initial request.
    ContinuationToken string `json:"ContinuationToken,omitempty"`
    // Transactions the requested inventory transactions.
    Transactions []TransactionModel `json:"Transactions,omitempty"`
}

// GooglePlayProductPurchase 
type GooglePlayProductPurchaseModel struct {
    // ProductId the Product ID (SKU) of the InApp product purchased from the Google Play store.
    ProductId string `json:"ProductId,omitempty"`
    // Token the token provided to the player's device when the product was purchased
    Token string `json:"Token,omitempty"`
}

// HelpfulnessVote 
type HelpfulnessVote string
  
const (
     HelpfulnessVoteNone HelpfulnessVote = "None"
     HelpfulnessVoteUnHelpful HelpfulnessVote = "UnHelpful"
     HelpfulnessVoteHelpful HelpfulnessVote = "Helpful"
      )
// Image 
type ImageModel struct {
    // Id the image unique ID.
    Id string `json:"Id,omitempty"`
    // Tag the client-defined tag associated with this image. Tags must be defined in the Catalog Config before being used in
// images
    Tag string `json:"Tag,omitempty"`
    // Type images can be defined as either a "thumbnail" or "screenshot". There can only be one "thumbnail" image per item.
    Type string `json:"Type,omitempty"`
    // Url the URL for retrieval of the image.
    Url string `json:"Url,omitempty"`
}

// ImageConfig 
type ImageConfigModel struct {
    // Tags the set of tags that will be used for validation. Each tag can have a maximum character length of 32 and up to 1024 tags
// can be listed.
    Tags []string `json:"Tags,omitempty"`
}

// InitialValues 
type InitialValuesModel struct {
    // DisplayProperties game specific properties for display purposes. The Display Properties field has a 1000 byte limit.
    DisplayProperties interface{} `json:"DisplayProperties,omitempty"`
}

// InventoryItem 
type InventoryItemModel struct {
    // Amount the amount of the item.
    Amount int32 `json:"Amount,omitempty"`
    // DisplayProperties game specific properties for display purposes. This is an arbitrary JSON blob. The Display Properties field has a 1000
// byte limit.
    DisplayProperties interface{} `json:"DisplayProperties,omitempty"`
    // ExpirationDate only used for subscriptions. The date of when the item will expire in UTC.
    ExpirationDate time.Time `json:"ExpirationDate,omitempty"`
    // Id the id of the item. This should correspond to the item id in the catalog.
    Id string `json:"Id,omitempty"`
    // StackId the stack id of the item.
    StackId string `json:"StackId,omitempty"`
    // Type the type of the item. This should correspond to the item type in the catalog.
    Type string `json:"Type,omitempty"`
}

// InventoryItemReference 
type InventoryItemReferenceModel struct {
    // AlternateId the inventory item alternate id the request applies to.
    AlternateId *AlternateIdModel `json:"AlternateId,omitempty"`
    // Id the inventory item id the request applies to.
    Id string `json:"Id,omitempty"`
    // StackId the inventory stack id the request should redeem to. (Default="default")
    StackId string `json:"StackId,omitempty"`
}

// InventoryOperation 
type InventoryOperationModel struct {
    // Add the add operation.
    Add *AddInventoryItemsOperationModel `json:"Add,omitempty"`
    // Delete the delete operation.
    Delete *DeleteInventoryItemsOperationModel `json:"Delete,omitempty"`
    // Purchase the purchase operation.
    Purchase *PurchaseInventoryItemsOperationModel `json:"Purchase,omitempty"`
    // Subtract the subtract operation.
    Subtract *SubtractInventoryItemsOperationModel `json:"Subtract,omitempty"`
    // Transfer the transfer operation.
    Transfer *TransferInventoryItemsOperationModel `json:"Transfer,omitempty"`
    // Update the update operation.
    Update *UpdateInventoryItemsOperationModel `json:"Update,omitempty"`
}

// KeywordSet 
type KeywordSetModel struct {
    // Values a list of localized keywords.
    Values []string `json:"Values,omitempty"`
}

// ModerationState 
type ModerationStateModel struct {
    // LastModifiedDate the date and time this moderation state was last updated.
    LastModifiedDate time.Time `json:"LastModifiedDate,omitempty"`
    // Reason the current stated reason for the associated item being moderated.
    Reason string `json:"Reason,omitempty"`
    // Status the current moderation status for the associated item.
    Status ModerationStatus `json:"Status,omitempty"`
}

// ModerationStatus 
type ModerationStatus string
  
const (
     ModerationStatusUnknown ModerationStatus = "Unknown"
     ModerationStatusAwaitingModeration ModerationStatus = "AwaitingModeration"
     ModerationStatusApproved ModerationStatus = "Approved"
     ModerationStatusRejected ModerationStatus = "Rejected"
      )
// PayoutDetails 
type PayoutDetailsModel struct {
}

// PublishDraftItemRequest the call kicks off a workflow to publish the item to the public catalog. The Publish Status API should be used to
// monitor the publish job.
type PublishDraftItemRequestModel struct {
    // AlternateId an alternate ID associated with this item.
    AlternateId *CatalogAlternateIdModel `json:"AlternateId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // ETag eTag of the catalog item to published from the working catalog to the public catalog. Used for optimistic concurrency.
// If the provided ETag does not match the ETag in the current working catalog, the request will be rejected. If not
// provided, the current version of the document in the working catalog will be published.
    ETag string `json:"ETag,omitempty"`
    // Id the unique ID of the item.
    Id string `json:"Id,omitempty"`
}

// PublishDraftItemResponse 
type PublishDraftItemResponseModel struct {
}

// PublishResult 
type PublishResult string
  
const (
     PublishResultUnknown PublishResult = "Unknown"
     PublishResultPending PublishResult = "Pending"
     PublishResultSucceeded PublishResult = "Succeeded"
     PublishResultFailed PublishResult = "Failed"
     PublishResultCanceled PublishResult = "Canceled"
      )
// PurchaseInventoryItemsOperation 
type PurchaseInventoryItemsOperationModel struct {
    // Amount the amount to purchase.
    Amount int32 `json:"Amount,omitempty"`
    // DeleteEmptyStacks indicates whether stacks reduced to an amount of 0 during the operation should be deleted from the inventory. (Default =
// false)
    DeleteEmptyStacks bool `json:"DeleteEmptyStacks"`
    // DurationInSeconds the duration to purchase.
    DurationInSeconds float64 `json:"DurationInSeconds,omitempty"`
    // Item the inventory item the operation applies to.
    Item *InventoryItemReferenceModel `json:"Item,omitempty"`
    // NewStackValues the values to apply to a stack newly created by this operation.
    NewStackValues *InitialValuesModel `json:"NewStackValues,omitempty"`
    // PriceAmounts the per-item price the item is expected to be purchased at. This must match a value configured in the Catalog or
// specified Store.
    PriceAmounts []PurchasePriceAmountModel `json:"PriceAmounts,omitempty"`
    // StoreId the id of the Store to purchase the item from.
    StoreId string `json:"StoreId,omitempty"`
}

// PurchaseInventoryItemsRequest purchase a single item or bundle, paying the associated price.
type PurchaseInventoryItemsRequestModel struct {
    // Amount the amount to purchase.
    Amount int32 `json:"Amount,omitempty"`
    // CollectionId the id of the entity's collection to perform this action on. (Default="default"). The number of inventory collections is
// unlimited.
    CollectionId string `json:"CollectionId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // DeleteEmptyStacks indicates whether stacks reduced to an amount of 0 during the request should be deleted from the inventory.
// (Default=false)
    DeleteEmptyStacks bool `json:"DeleteEmptyStacks"`
    // DurationInSeconds the duration to purchase.
    DurationInSeconds float64 `json:"DurationInSeconds,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // ETag eTags are used for concurrency checking when updating resources. More information about using ETags can be found here:
// https://learn.microsoft.com/en-us/gaming/playfab/features/economy-v2/catalog/etags
    ETag string `json:"ETag,omitempty"`
    // IdempotencyId the Idempotency ID for this request. Idempotency IDs can be used to prevent operation replay in the medium term but will
// be garbage collected eventually.
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // Item the inventory item the request applies to.
    Item *InventoryItemReferenceModel `json:"Item,omitempty"`
    // NewStackValues the values to apply to a stack newly created by this request.
    NewStackValues *InitialValuesModel `json:"NewStackValues,omitempty"`
    // PriceAmounts the per-item price the item is expected to be purchased at. This must match a value configured in the Catalog or
// specified Store.
    PriceAmounts []PurchasePriceAmountModel `json:"PriceAmounts,omitempty"`
    // StoreId the id of the Store to purchase the item from.
    StoreId string `json:"StoreId,omitempty"`
}

// PurchaseInventoryItemsResponse 
type PurchaseInventoryItemsResponseModel struct {
    // ETag eTags are used for concurrency checking when updating resources. More information about using ETags can be found here:
// https://learn.microsoft.com/en-us/gaming/playfab/features/economy-v2/catalog/etags
    ETag string `json:"ETag,omitempty"`
    // IdempotencyId the idempotency id used in the request.
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // TransactionIds the ids of transactions that occurred as a result of the request.
    TransactionIds []string `json:"TransactionIds,omitempty"`
}

// PurchaseOverride 
type PurchaseOverrideModel struct {
}

// PurchasePriceAmount 
type PurchasePriceAmountModel struct {
    // Amount the amount of the inventory item to use in the purchase .
    Amount int32 `json:"Amount,omitempty"`
    // ItemId the inventory item id to use in the purchase .
    ItemId string `json:"ItemId,omitempty"`
    // StackId the inventory stack id the to use in the purchase. Set to "default" by default
    StackId string `json:"StackId,omitempty"`
}

// Rating 
type RatingModel struct {
    // Average the average rating for this item.
    Average float32 `json:"Average,omitempty"`
    // Count1Star the total count of 1 star ratings for this item.
    Count1Star int32 `json:"Count1Star,omitempty"`
    // Count2Star the total count of 2 star ratings for this item.
    Count2Star int32 `json:"Count2Star,omitempty"`
    // Count3Star the total count of 3 star ratings for this item.
    Count3Star int32 `json:"Count3Star,omitempty"`
    // Count4Star the total count of 4 star ratings for this item.
    Count4Star int32 `json:"Count4Star,omitempty"`
    // Count5Star the total count of 5 star ratings for this item.
    Count5Star int32 `json:"Count5Star,omitempty"`
    // TotalCount the total count of ratings for this item.
    TotalCount int32 `json:"TotalCount,omitempty"`
}

// RedeemAppleAppStoreInventoryItemsRequest redeem items from the Apple App Store.
type RedeemAppleAppStoreInventoryItemsRequestModel struct {
    // CollectionId the id of the entity's collection to perform this action on. (Default="default")
    CollectionId string `json:"CollectionId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // Receipt the receipt provided by the Apple marketplace upon successful purchase.
    Receipt string `json:"Receipt,omitempty"`
}

// RedeemAppleAppStoreInventoryItemsResponse 
type RedeemAppleAppStoreInventoryItemsResponseModel struct {
    // Failed the list of failed redemptions from the external marketplace.
    Failed []RedemptionFailureModel `json:"Failed,omitempty"`
    // Succeeded the list of successful redemptions from the external marketplace.
    Succeeded []RedemptionSuccessModel `json:"Succeeded,omitempty"`
    // TransactionIds the Transaction IDs associated with the inventory modifications
    TransactionIds []string `json:"TransactionIds,omitempty"`
}

// RedeemGooglePlayInventoryItemsRequest redeem items from the Google Play Store.
type RedeemGooglePlayInventoryItemsRequestModel struct {
    // CollectionId the id of the entity's collection to perform this action on. (Default="default")
    CollectionId string `json:"CollectionId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // Purchases the list of purchases to redeem
    Purchases []GooglePlayProductPurchaseModel `json:"Purchases,omitempty"`
}

// RedeemGooglePlayInventoryItemsResponse 
type RedeemGooglePlayInventoryItemsResponseModel struct {
    // Failed the list of failed redemptions from the external marketplace.
    Failed []RedemptionFailureModel `json:"Failed,omitempty"`
    // Succeeded the list of successful redemptions from the external marketplace.
    Succeeded []RedemptionSuccessModel `json:"Succeeded,omitempty"`
    // TransactionIds the Transaction IDs associated with the inventory modifications
    TransactionIds []string `json:"TransactionIds,omitempty"`
}

// RedeemMicrosoftStoreInventoryItemsRequest redeem items from the Microsoft Store.
type RedeemMicrosoftStoreInventoryItemsRequestModel struct {
    // CollectionId the id of the entity's collection to perform this action on. (Default="default")
    CollectionId string `json:"CollectionId,omitempty"`
    // CollectionsIdKey the OneStore Collections Id Key used for AAD authentication.
    CollectionsIdKey string `json:"CollectionsIdKey,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // XboxToken xbox Token used for delegated business partner authentication.
    XboxToken string `json:"XboxToken,omitempty"`
}

// RedeemMicrosoftStoreInventoryItemsResponse 
type RedeemMicrosoftStoreInventoryItemsResponseModel struct {
    // Failed the list of failed redemptions from the external marketplace.
    Failed []RedemptionFailureModel `json:"Failed,omitempty"`
    // Succeeded the list of successful redemptions from the external marketplace.
    Succeeded []RedemptionSuccessModel `json:"Succeeded,omitempty"`
    // TransactionIds the Transaction IDs associated with the inventory modifications
    TransactionIds []string `json:"TransactionIds,omitempty"`
}

// RedeemNintendoEShopInventoryItemsRequest redeem items from the Nintendo EShop.
type RedeemNintendoEShopInventoryItemsRequestModel struct {
    // CollectionId the id of the entity's collection to perform this action on. (Default="default")
    CollectionId string `json:"CollectionId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // NintendoServiceAccountIdToken the Nintendo provided token authorizing redemption
    NintendoServiceAccountIdToken string `json:"NintendoServiceAccountIdToken,omitempty"`
}

// RedeemNintendoEShopInventoryItemsResponse 
type RedeemNintendoEShopInventoryItemsResponseModel struct {
    // Failed the list of failed redemptions from the external marketplace.
    Failed []RedemptionFailureModel `json:"Failed,omitempty"`
    // Succeeded the list of successful redemptions from the external marketplace.
    Succeeded []RedemptionSuccessModel `json:"Succeeded,omitempty"`
    // TransactionIds the Transaction IDs associated with the inventory modifications
    TransactionIds []string `json:"TransactionIds,omitempty"`
}

// RedeemPlayStationStoreInventoryItemsRequest redeem items from the PlayStation Store.
type RedeemPlayStationStoreInventoryItemsRequestModel struct {
    // AuthorizationCode auth code returned by PlayStation :tm: Network OAuth system.
    AuthorizationCode string `json:"AuthorizationCode,omitempty"`
    // CollectionId the id of the entity's collection to perform this action on. (Default="default")
    CollectionId string `json:"CollectionId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // RedirectUri redirect URI supplied to PlayStation :tm: Network when requesting an auth code
    RedirectUri string `json:"RedirectUri,omitempty"`
    // ServiceLabel optional Service Label to pass into the request.
    ServiceLabel string `json:"ServiceLabel,omitempty"`
}

// RedeemPlayStationStoreInventoryItemsResponse 
type RedeemPlayStationStoreInventoryItemsResponseModel struct {
    // Failed the list of failed redemptions from the external marketplace.
    Failed []RedemptionFailureModel `json:"Failed,omitempty"`
    // Succeeded the list of successful redemptions from the external marketplace.
    Succeeded []RedemptionSuccessModel `json:"Succeeded,omitempty"`
    // TransactionIds the Transaction IDs associated with the inventory modifications
    TransactionIds []string `json:"TransactionIds,omitempty"`
}

// RedeemSteamInventoryItemsRequest redeem inventory items from Steam.
type RedeemSteamInventoryItemsRequestModel struct {
    // CollectionId the id of the entity's collection to perform this action on. (Default="default")
    CollectionId string `json:"CollectionId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
}

// RedeemSteamInventoryItemsResponse 
type RedeemSteamInventoryItemsResponseModel struct {
    // Failed the list of failed redemptions from the external marketplace.
    Failed []RedemptionFailureModel `json:"Failed,omitempty"`
    // Succeeded the list of successful redemptions from the external marketplace.
    Succeeded []RedemptionSuccessModel `json:"Succeeded,omitempty"`
    // TransactionIds the Transaction IDs associated with the inventory modifications
    TransactionIds []string `json:"TransactionIds,omitempty"`
}

// RedemptionFailure 
type RedemptionFailureModel struct {
    // FailureCode the marketplace failure code.
    FailureCode string `json:"FailureCode,omitempty"`
    // FailureDetails the marketplace error details explaining why the offer failed to redeem.
    FailureDetails string `json:"FailureDetails,omitempty"`
    // MarketplaceTransactionId the transaction id in the external marketplace.
    MarketplaceTransactionId string `json:"MarketplaceTransactionId,omitempty"`
    // OfferId the ID of the offer being redeemed.
    OfferId string `json:"OfferId,omitempty"`
}

// RedemptionSuccess 
type RedemptionSuccessModel struct {
    // MarketplaceTransactionId the transaction id in the external marketplace.
    MarketplaceTransactionId string `json:"MarketplaceTransactionId,omitempty"`
    // OfferId the ID of the offer being redeemed.
    OfferId string `json:"OfferId,omitempty"`
    // SuccessTimestamp the timestamp for when the redeem was completed.
    SuccessTimestamp time.Time `json:"SuccessTimestamp,omitempty"`
}

// ReportItemRequest 
type ReportItemRequestModel struct {
    // AlternateId an alternate ID associated with this item.
    AlternateId *CatalogAlternateIdModel `json:"AlternateId,omitempty"`
    // ConcernCategory category of concern for this report.
    ConcernCategory ConcernCategory `json:"ConcernCategory,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // Id the unique ID of the item.
    Id string `json:"Id,omitempty"`
    // Reason the string reason for this report.
    Reason string `json:"Reason,omitempty"`
}

// ReportItemResponse 
type ReportItemResponseModel struct {
}

// ReportItemReviewRequest submit a report for an inappropriate review, allowing the submitting user to specify their concern.
type ReportItemReviewRequestModel struct {
    // AlternateId an alternate ID of the item associated with the review.
    AlternateId *CatalogAlternateIdModel `json:"AlternateId,omitempty"`
    // ConcernCategory the reason this review is being reported.
    ConcernCategory ConcernCategory `json:"ConcernCategory,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // ItemId the string ID of the item associated with the review.
    ItemId string `json:"ItemId,omitempty"`
    // Reason the string reason for this report.
    Reason string `json:"Reason,omitempty"`
    // ReviewId the ID of the review to submit a report for.
    ReviewId string `json:"ReviewId,omitempty"`
}

// ReportItemReviewResponse 
type ReportItemReviewResponseModel struct {
}

// Review 
type ReviewModel struct {
    // HelpfulNegative the number of negative helpfulness votes for this review.
    HelpfulNegative int32 `json:"HelpfulNegative,omitempty"`
    // HelpfulPositive the number of positive helpfulness votes for this review.
    HelpfulPositive int32 `json:"HelpfulPositive,omitempty"`
    // IsInstalled indicates whether the review author has the item installed.
    IsInstalled bool `json:"IsInstalled"`
    // ItemId the ID of the item being reviewed.
    ItemId string `json:"ItemId,omitempty"`
    // ItemVersion the version of the item being reviewed.
    ItemVersion string `json:"ItemVersion,omitempty"`
    // Locale the locale for which this review was submitted in.
    Locale string `json:"Locale,omitempty"`
    // Rating star rating associated with this review.
    Rating int32 `json:"Rating,omitempty"`
    // ReviewerEntity the ID of the author of the review.
    ReviewerEntity *EntityKeyModel `json:"ReviewerEntity,omitempty"`
    // ReviewerId deprecated. Use ReviewerEntity instead. This property will be removed in a future release.
    ReviewerId string `json:"ReviewerId,omitempty"`
    // ReviewId the ID of the review.
    ReviewId string `json:"ReviewId,omitempty"`
    // ReviewText the full text of this review.
    ReviewText string `json:"ReviewText,omitempty"`
    // Submitted the date and time this review was last submitted.
    Submitted time.Time `json:"Submitted,omitempty"`
    // Title the title of this review.
    Title string `json:"Title,omitempty"`
}

// ReviewItemRequest 
type ReviewItemRequestModel struct {
    // AlternateId an alternate ID associated with this item.
    AlternateId *CatalogAlternateIdModel `json:"AlternateId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // Id the unique ID of the item.
    Id string `json:"Id,omitempty"`
    // Review the review to submit.
    Review *ReviewModel `json:"Review,omitempty"`
}

// ReviewItemResponse 
type ReviewItemResponseModel struct {
}

// ReviewTakedown 
type ReviewTakedownModel struct {
    // AlternateId an alternate ID associated with this item.
    AlternateId *CatalogAlternateIdModel `json:"AlternateId,omitempty"`
    // ItemId the ID of the item associated with the review to take down.
    ItemId string `json:"ItemId,omitempty"`
    // ReviewId the ID of the review to take down.
    ReviewId string `json:"ReviewId,omitempty"`
}

// ScanResult 
type ScanResultModel struct {
    // Url the URL of the item which failed the scan.
    Url string `json:"Url,omitempty"`
}

// SearchItemsRequest 
type SearchItemsRequestModel struct {
    // ContinuationToken an opaque token used to retrieve the next page of items, if any are available.
    ContinuationToken string `json:"ContinuationToken,omitempty"`
    // Count number of items to retrieve. This value is optional. Maximum page size is 50. Default value is 10.
    Count int32 `json:"Count,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // Filter an OData filter used to refine the search query (For example: "type eq 'ugc'"). More info about Filter Complexity limits
// can be found here: https://learn.microsoft.com/en-us/gaming/playfab/features/economy-v2/catalog/search#limits
    Filter string `json:"Filter,omitempty"`
    // OrderBy an OData orderBy used to order the results of the search query. For example: "rating/average asc"
    OrderBy string `json:"OrderBy,omitempty"`
    // Search the text to search for.
    Search string `json:"Search,omitempty"`
    // Select an OData select query option used to augment the search results. If not defined, the default search result metadata will
// be returned.
    Select string `json:"Select,omitempty"`
    // Store the store to restrict the search request to.
    Store *StoreReferenceModel `json:"Store,omitempty"`
}

// SearchItemsResponse 
type SearchItemsResponseModel struct {
    // ContinuationToken an opaque token used to retrieve the next page of items, if any are available.
    ContinuationToken string `json:"ContinuationToken,omitempty"`
    // Items the paginated set of results for the search query.
    Items []CatalogItemModel `json:"Items,omitempty"`
}

// SetItemModerationStateRequest 
type SetItemModerationStateRequestModel struct {
    // AlternateId an alternate ID associated with this item.
    AlternateId *CatalogAlternateIdModel `json:"AlternateId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Id the unique ID of the item.
    Id string `json:"Id,omitempty"`
    // Reason the reason for the moderation state change for the associated item.
    Reason string `json:"Reason,omitempty"`
    // Status the status to set for the associated item.
    Status ModerationStatus `json:"Status,omitempty"`
}

// SetItemModerationStateResponse 
type SetItemModerationStateResponseModel struct {
}

// StoreDetails 
type StoreDetailsModel struct {
    // FilterOptions the options for the filter in filter-based stores. These options are mutually exclusive with item references.
    FilterOptions *FilterOptionsModel `json:"FilterOptions,omitempty"`
    // PriceOptionsOverride the global prices utilized in the store. These options are mutually exclusive with price options in item references.
    PriceOptionsOverride *CatalogPriceOptionsOverrideModel `json:"PriceOptionsOverride,omitempty"`
}

// StoreReference 
type StoreReferenceModel struct {
    // AlternateId an alternate ID of the store.
    AlternateId *CatalogAlternateIdModel `json:"AlternateId,omitempty"`
    // Id the unique ID of the store.
    Id string `json:"Id,omitempty"`
}

// SubmitItemReviewVoteRequest 
type SubmitItemReviewVoteRequestModel struct {
    // AlternateId an alternate ID of the item associated with the review.
    AlternateId *CatalogAlternateIdModel `json:"AlternateId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // ItemId the string ID of the item associated with the review.
    ItemId string `json:"ItemId,omitempty"`
    // ReviewId the ID of the review to submit a helpfulness vote for.
    ReviewId string `json:"ReviewId,omitempty"`
    // Vote the helpfulness vote of the review.
    Vote HelpfulnessVote `json:"Vote,omitempty"`
}

// SubmitItemReviewVoteResponse 
type SubmitItemReviewVoteResponseModel struct {
}

// SubscriptionDetails 
type SubscriptionDetailsModel struct {
    // DurationInSeconds the length of time that the subscription will last in seconds.
    DurationInSeconds float64 `json:"DurationInSeconds,omitempty"`
}

// SubtractInventoryItemsOperation 
type SubtractInventoryItemsOperationModel struct {
    // Amount the amount to subtract from the current item amount.
    Amount int32 `json:"Amount,omitempty"`
    // DeleteEmptyStacks indicates whether stacks reduced to an amount of 0 during the request should be deleted from the inventory. (Default =
// false).
    DeleteEmptyStacks bool `json:"DeleteEmptyStacks"`
    // DurationInSeconds the duration to subtract from the current item expiration date.
    DurationInSeconds float64 `json:"DurationInSeconds,omitempty"`
    // Item the inventory item the operation applies to.
    Item *InventoryItemReferenceModel `json:"Item,omitempty"`
}

// SubtractInventoryItemsRequest given an entity type, entity identifier and container details, will subtract the specified inventory items.
type SubtractInventoryItemsRequestModel struct {
    // Amount the amount to subtract for the current item.
    Amount int32 `json:"Amount,omitempty"`
    // CollectionId the id of the entity's collection to perform this action on. (Default="default"). The number of inventory collections is
// unlimited.
    CollectionId string `json:"CollectionId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // DeleteEmptyStacks indicates whether stacks reduced to an amount of 0 during the request should be deleted from the inventory.
// (Default=false)
    DeleteEmptyStacks bool `json:"DeleteEmptyStacks"`
    // DurationInSeconds the duration to subtract from the current item expiration date.
    DurationInSeconds float64 `json:"DurationInSeconds,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // ETag eTags are used for concurrency checking when updating resources. More information about using ETags can be found here:
// https://learn.microsoft.com/en-us/gaming/playfab/features/economy-v2/catalog/etags
    ETag string `json:"ETag,omitempty"`
    // IdempotencyId the Idempotency ID for this request. Idempotency IDs can be used to prevent operation replay in the medium term but will
// be garbage collected eventually.
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // Item the inventory item the request applies to.
    Item *InventoryItemReferenceModel `json:"Item,omitempty"`
}

// SubtractInventoryItemsResponse 
type SubtractInventoryItemsResponseModel struct {
    // ETag eTags are used for concurrency checking when updating resources. More information about using ETags can be found here:
// https://learn.microsoft.com/en-us/gaming/playfab/features/economy-v2/catalog/etags
    ETag string `json:"ETag,omitempty"`
    // IdempotencyId the idempotency id used in the request.
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // TransactionIds the ids of transactions that occurred as a result of the request.
    TransactionIds []string `json:"TransactionIds,omitempty"`
}

// TakedownItemReviewsRequest submit a request to takedown one or more reviews, removing them from public view. Authors will still be able to see
// their reviews after being taken down.
type TakedownItemReviewsRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Reviews the set of reviews to take down.
    Reviews []ReviewTakedownModel `json:"Reviews,omitempty"`
}

// TakedownItemReviewsResponse 
type TakedownItemReviewsResponseModel struct {
}

// Transaction 
type TransactionModel struct {
    // ApiName the API call that caused this transaction.
    ApiName string `json:"ApiName,omitempty"`
    // ItemType the type of item that the the operation occurred on.
    ItemType string `json:"ItemType,omitempty"`
    // Operations the operations that occurred.
    Operations []TransactionOperationModel `json:"Operations,omitempty"`
    // OperationType the type of operation that was run.
    OperationType string `json:"OperationType,omitempty"`
    // PurchaseDetails additional details about the transaction. Null if it was not a purchase operation.
    PurchaseDetails *TransactionPurchaseDetailsModel `json:"PurchaseDetails,omitempty"`
    // RedeemDetails additional details about the transaction. Null if it was not a redeem operation.
    RedeemDetails *TransactionRedeemDetailsModel `json:"RedeemDetails,omitempty"`
    // Timestamp the time this transaction occurred in UTC.
    Timestamp time.Time `json:"Timestamp,omitempty"`
    // TransactionId the id of the transaction. This should be treated like an opaque token.
    TransactionId string `json:"TransactionId,omitempty"`
    // TransferDetails additional details about the transaction. Null if it was not a transfer operation.
    TransferDetails *TransactionTransferDetailsModel `json:"TransferDetails,omitempty"`
}

// TransactionOperation 
type TransactionOperationModel struct {
    // Amount the amount of items in this transaction.
    Amount int32 `json:"Amount,omitempty"`
    // DurationInSeconds the duration modified in this transaction.
    DurationInSeconds float64 `json:"DurationInSeconds,omitempty"`
    // ItemId the item id of the items in this transaction.
    ItemId string `json:"ItemId,omitempty"`
    // ItemType the type of item that the operation occurred on.
    ItemType string `json:"ItemType,omitempty"`
    // StackId the stack id of the items in this transaction.
    StackId string `json:"StackId,omitempty"`
    // Type the type of the operation that occurred.
    Type string `json:"Type,omitempty"`
}

// TransactionPurchaseDetails 
type TransactionPurchaseDetailsModel struct {
    // StoreId the id of the Store the item was purchased from or null.
    StoreId string `json:"StoreId,omitempty"`
}

// TransactionRedeemDetails 
type TransactionRedeemDetailsModel struct {
    // Marketplace the marketplace that the offer is being redeemed from.
    Marketplace string `json:"Marketplace,omitempty"`
    // MarketplaceTransactionId the transaction Id returned from the marketplace.
    MarketplaceTransactionId string `json:"MarketplaceTransactionId,omitempty"`
    // OfferId the offer Id of the item being redeemed.
    OfferId string `json:"OfferId,omitempty"`
}

// TransactionTransferDetails 
type TransactionTransferDetailsModel struct {
    // GivingCollectionId the collection id the items were transferred from or null if it was the current collection.
    GivingCollectionId string `json:"GivingCollectionId,omitempty"`
    // GivingEntity the entity the items were transferred from or null if it was the current entity.
    GivingEntity *EntityKeyModel `json:"GivingEntity,omitempty"`
    // ReceivingCollectionId the collection id the items were transferred to or null if it was the current collection.
    ReceivingCollectionId string `json:"ReceivingCollectionId,omitempty"`
    // ReceivingEntity the entity the items were transferred to or null if it was the current entity.
    ReceivingEntity *EntityKeyModel `json:"ReceivingEntity,omitempty"`
    // TransferId the id of the transfer that occurred.
    TransferId string `json:"TransferId,omitempty"`
}

// TransferInventoryItemsOperation 
type TransferInventoryItemsOperationModel struct {
    // Amount the amount to transfer.
    Amount int32 `json:"Amount,omitempty"`
    // DeleteEmptyStacks indicates whether stacks reduced to an amount of 0 during the operation should be deleted from the inventory. (Default =
// false)
    DeleteEmptyStacks bool `json:"DeleteEmptyStacks"`
    // GivingItem the inventory item the operation is transferring from.
    GivingItem *InventoryItemReferenceModel `json:"GivingItem,omitempty"`
    // NewStackValues the values to apply to a stack newly created by this operation.
    NewStackValues *InitialValuesModel `json:"NewStackValues,omitempty"`
    // ReceivingItem the inventory item the operation is transferring to.
    ReceivingItem *InventoryItemReferenceModel `json:"ReceivingItem,omitempty"`
}

// TransferInventoryItemsRequest transfer the specified inventory items of an entity's container Id to another entity's container Id.
type TransferInventoryItemsRequestModel struct {
    // Amount the amount to transfer .
    Amount int32 `json:"Amount,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // DeleteEmptyStacks indicates whether stacks reduced to an amount of 0 during the request should be deleted from the inventory. (Default =
// false)
    DeleteEmptyStacks bool `json:"DeleteEmptyStacks"`
    // GivingCollectionId the inventory collection id the request is transferring from. (Default="default")
    GivingCollectionId string `json:"GivingCollectionId,omitempty"`
    // GivingEntity the entity the request is transferring from. Set to the caller by default.
    GivingEntity *EntityKeyModel `json:"GivingEntity,omitempty"`
    // GivingETag eTags are used for concurrency checking when updating resources (before transferring from). More information about using
// ETags can be found here: https://learn.microsoft.com/en-us/gaming/playfab/features/economy-v2/catalog/etags
    GivingETag string `json:"GivingETag,omitempty"`
    // GivingItem the inventory item the request is transferring from.
    GivingItem *InventoryItemReferenceModel `json:"GivingItem,omitempty"`
    // IdempotencyId the idempotency id for the request.
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // NewStackValues the values to apply to a stack newly created by this request.
    NewStackValues *InitialValuesModel `json:"NewStackValues,omitempty"`
    // ReceivingCollectionId the inventory collection id the request is transferring to. (Default="default")
    ReceivingCollectionId string `json:"ReceivingCollectionId,omitempty"`
    // ReceivingEntity the entity the request is transferring to. Set to the caller by default.
    ReceivingEntity *EntityKeyModel `json:"ReceivingEntity,omitempty"`
    // ReceivingItem the inventory item the request is transferring to.
    ReceivingItem *InventoryItemReferenceModel `json:"ReceivingItem,omitempty"`
}

// TransferInventoryItemsResponse 
type TransferInventoryItemsResponseModel struct {
    // GivingETag eTags are used for concurrency checking when updating resources (after transferring from). More information about using
// ETags can be found here: https://learn.microsoft.com/en-us/gaming/playfab/features/economy-v2/catalog/etags
    GivingETag string `json:"GivingETag,omitempty"`
    // GivingTransactionIds the ids of transactions that occurred as a result of the request's giving action.
    GivingTransactionIds []string `json:"GivingTransactionIds,omitempty"`
    // IdempotencyId the idempotency id for the request.
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // ReceivingTransactionIds the ids of transactions that occurred as a result of the request's receiving action.
    ReceivingTransactionIds []string `json:"ReceivingTransactionIds,omitempty"`
}

// UpdateCatalogConfigRequest 
type UpdateCatalogConfigRequestModel struct {
    // Config the updated catalog configuration.
    Config *CatalogConfigModel `json:"Config,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
}

// UpdateCatalogConfigResponse 
type UpdateCatalogConfigResponseModel struct {
}

// UpdateDraftItemRequest 
type UpdateDraftItemRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Item updated metadata describing the catalog item to be updated.
    Item *CatalogItemModel `json:"Item,omitempty"`
    // Publish whether the item should be published immediately. This value is optional, defaults to false.
    Publish bool `json:"Publish"`
}

// UpdateDraftItemResponse 
type UpdateDraftItemResponseModel struct {
    // Item updated metadata describing the catalog item just updated.
    Item *CatalogItemModel `json:"Item,omitempty"`
}

// UpdateInventoryItemsOperation 
type UpdateInventoryItemsOperationModel struct {
    // Item the inventory item to update with the specified values.
    Item *InventoryItemModel `json:"Item,omitempty"`
}

// UpdateInventoryItemsRequest given an entity type, entity identifier and container details, will update the entity's inventory items
type UpdateInventoryItemsRequestModel struct {
    // CollectionId the id of the entity's collection to perform this action on. (Default="default"). The number of inventory collections is
// unlimited.
    CollectionId string `json:"CollectionId,omitempty"`
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // ETag eTags are used for concurrency checking when updating resources. More information about using ETags can be found here:
// https://learn.microsoft.com/en-us/gaming/playfab/features/economy-v2/catalog/etags
    ETag string `json:"ETag,omitempty"`
    // IdempotencyId the Idempotency ID for this request. Idempotency IDs can be used to prevent operation replay in the medium term but will
// be garbage collected eventually.
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // Item the inventory item to update with the specified values.
    Item *InventoryItemModel `json:"Item,omitempty"`
}

// UpdateInventoryItemsResponse 
type UpdateInventoryItemsResponseModel struct {
    // ETag eTags are used for concurrency checking when updating resources. More information about using ETags can be found here:
// https://learn.microsoft.com/en-us/gaming/playfab/features/economy-v2/catalog/etags
    ETag string `json:"ETag,omitempty"`
    // IdempotencyId the idempotency id used in the request.
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // TransactionIds the ids of transactions that occurred as a result of the request.
    TransactionIds []string `json:"TransactionIds,omitempty"`
}

// UploadInfo 
type UploadInfoModel struct {
    // FileName name of the file to be uploaded.
    FileName string `json:"FileName,omitempty"`
}

// UploadUrlMetadata 
type UploadUrlMetadataModel struct {
    // FileName name of the file for which this upload URL was requested.
    FileName string `json:"FileName,omitempty"`
    // Id unique ID for the binary content to be uploaded to the target URL.
    Id string `json:"Id,omitempty"`
    // Url uRL for the binary content to be uploaded to.
    Url string `json:"Url,omitempty"`
}

// UserGeneratedContentSpecificConfig 
type UserGeneratedContentSpecificConfigModel struct {
    // ContentTypes the set of content types that will be used for validation.
    ContentTypes []string `json:"ContentTypes,omitempty"`
    // Tags the set of tags that will be used for validation.
    Tags []string `json:"Tags,omitempty"`
}
