package localization
 
// GetLanguageListRequest 
type GetLanguageListRequestModel struct {
}

// GetLanguageListResponse 
type GetLanguageListResponseModel struct {
    // LanguageList the list of allowed languages, in BCP47 two-letter format
    LanguageList []string `json:"LanguageList,omitempty"`
}
