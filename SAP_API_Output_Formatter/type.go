package sap_api_output_formatter

type ServiceAgent struct {
	ConnectionKey    string `json:"connection_key"`
	Result           bool   `json:"result"`
	RedisKey         string `json:"redis_key"`
	Filepath         string `json:"filepath"`
	APISchema        string `json:"api_schema"`
	ServiceAgentCode string `json:"service_agent_code"`
	Deleted          bool   `json:"deleted"`
}

type ServiceAgentCollection struct {
	ObjectID                          string `json:"ObjectID"`
	ServiceAgentID                    string `json:"ServiceAgentID"`
	ServiceAgentUUID                  string `json:"ServiceAgentUUID"`
	LifeCycleStatusCode               string `json:"LifeCycleStatusCode"`
	LifeCycleStatusCodeText           string `json:"LifeCycleStatusCodeText"`
	TitleCode                         string `json:"TitleCode"`
	TitleCodeText                     string `json:"TitleCodeText"`
	AcademicTitleCode                 string `json:"AcademicTitleCode"`
	AcademicTitleCodeText             string `json:"AcademicTitleCodeText"`
	BusinessPartnerFormattedName      string `json:"BusinessPartnerFormattedName"`
	FirstName                         string `json:"FirstName"`
	LastName                          string `json:"LastName"`
	MiddleName                        string `json:"MiddleName"`
	NickName                          string `json:"NickName"`
	GenderCode                        string `json:"GenderCode"`
	GenderCodeText                    string `json:"GenderCodeText"`
	LanguageCode                      string `json:"LanguageCode"`
	LanguageCodeText                  string `json:"LanguageCodeText"`
	NationalityCountryCode            string `json:"NationalityCountryCode"`
	NationalityCountryCodeText        string `json:"NationalityCountryCodeText"`
	BirthName                         string `json:"BirthName"`
	FormattedPostalAddressDescription string `json:"FormattedPostalAddressDescription"`
	CountryCode                       string `json:"CountryCode"`
	CountryCodeText                   string `json:"CountryCodeText"`
	RegionCode                        string `json:"RegionCode"`
	RegionCodeText                    string `json:"RegionCodeText"`
	AddressLine1                      string `json:"AddressLine1"`
	AddressLine2                      string `json:"AddressLine2"`
	HouseNumber                       string `json:"HouseNumber"`
	Street                            string `json:"Street"`
	AddressLine4                      string `json:"AddressLine4"`
	AddressLine5                      string `json:"AddressLine5"`
	District                          string `json:"District"`
	City                              string `json:"City"`
	DifferentCity                     string `json:"DifferentCity"`
	PostalCode                        string `json:"PostalCode"`
	County                            string `json:"County"`
	CompanyPostalCode                 string `json:"CompanyPostalCode"`
	POBox                             string `json:"POBox"`
	POBoxPostalCode                   string `json:"POBoxPostalCode"`
	POBoxCountryCode                  string `json:"POBoxCountryCode"`
	POBoxCountryCodeText              string `json:"POBoxCountryCodeText"`
	POBoxRegionCode                   string `json:"POBoxRegionCode"`
	POBoxRegionCodeText               string `json:"POBoxRegionCodeText"`
	POBoxCity                         string `json:"POBoxCity"`
	TimeZoneCode                      string `json:"TimeZoneCode"`
	TimeZoneCodeText                  string `json:"TimeZoneCodeText"`
	TaxJurisdictionCode               string `json:"TaxJurisdictionCode"`
	TaxJurisdictionCodeText           string `json:"TaxJurisdictionCodeText"`
	Building                          string `json:"Building"`
	Floor                             string `json:"Floor"`
	Room                              string `json:"Room"`
	InhouseMail                       string `json:"InhouseMail"`
	OfficePhoneNumber                 string `json:"OfficePhoneNumber"`
	MobilePhoneNumber                 string `json:"MobilePhoneNumber"`
	FaxNumber                         string `json:"FaxNumber"`
	Email                             string `json:"Email"`
	NormalisedOfficePhoneNumber       string `json:"NormalisedOfficePhoneNumber"`
	NormalisedMobilePhoneNumber       string `json:"NormalisedMobilePhoneNumber"`
	CreatedOn                         string `json:"CreatedOn"`
	CreatedBy                         string `json:"CreatedBy"`
	CreatedByIdentityUUID             string `json:"CreatedByIdentityUUID"`
	ChangedOn                         string `json:"ChangedOn"`
	ChangedBy                         string `json:"ChangedBy"`
	ChangedByIdentityUUID             string `json:"ChangedByIdentityUUID"`
	EntityLastChangedOn               string `json:"EntityLastChangedOn"`
	ETag                              string `json:"ETag"`
}
