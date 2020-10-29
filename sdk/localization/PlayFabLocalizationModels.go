package localization
 
// GetLanguageListRequest 
type GetLanguageListRequestModel struct {
    // CustomTags the optional custom tags associated with the request (e.g. build number, external trace identifiers, etc.).
    CustomTags map[string]string `json:"CustomTags,omitempty"`
}

// GetLanguageListResponse 
type GetLanguageListResponseModel struct {
    // LanguageList the list of allowed languages, in BCP47 two-letter format
    LanguageList []string `json:"LanguageList,omitempty"`
}
