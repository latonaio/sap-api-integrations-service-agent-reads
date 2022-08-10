package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-service-agent-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToServiceAgentCollection(raw []byte, l *logger.Logger) ([]ServiceAgentCollection, error) {
	pm := &responses.ServiceAgentCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ServiceAgentCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	serviceAgentCollection := make([]ServiceAgentCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		serviceAgentCollection = append(serviceAgentCollection, ServiceAgentCollection{
			ObjectID:                          data.ObjectID,
			ServiceAgentID:                    data.ServiceAgentID,
			ServiceAgentUUID:                  data.ServiceAgentUUID,
			LifeCycleStatusCode:               data.LifeCycleStatusCode,
			LifeCycleStatusCodeText:           data.LifeCycleStatusCodeText,
			TitleCode:                         data.TitleCode,
			TitleCodeText:                     data.TitleCodeText,
			AcademicTitleCode:                 data.AcademicTitleCode,
			AcademicTitleCodeText:             data.AcademicTitleCodeText,
			BusinessPartnerFormattedName:      data.BusinessPartnerFormattedName,
			FirstName:                         data.FirstName,
			LastName:                          data.LastName,
			MiddleName:                        data.MiddleName,
			NickName:                          data.NickName,
			GenderCode:                        data.GenderCode,
			GenderCodeText:                    data.GenderCodeText,
			LanguageCode:                      data.LanguageCode,
			LanguageCodeText:                  data.LanguageCodeText,
			NationalityCountryCode:            data.NationalityCountryCode,
			NationalityCountryCodeText:        data.NationalityCountryCodeText,
			BirthName:                         data.BirthName,
			FormattedPostalAddressDescription: data.FormattedPostalAddressDescription,
			CountryCode:                       data.CountryCode,
			CountryCodeText:                   data.CountryCodeText,
			RegionCode:                        data.RegionCode,
			RegionCodeText:                    data.RegionCodeText,
			AddressLine1:                      data.AddressLine1,
			AddressLine2:                      data.AddressLine2,
			HouseNumber:                       data.HouseNumber,
			Street:                            data.Street,
			AddressLine4:                      data.AddressLine4,
			AddressLine5:                      data.AddressLine5,
			District:                          data.District,
			City:                              data.City,
			DifferentCity:                     data.DifferentCity,
			PostalCode:                        data.PostalCode,
			County:                            data.County,
			CompanyPostalCode:                 data.CompanyPostalCode,
			POBox:                             data.POBox,
			POBoxPostalCode:                   data.POBoxPostalCode,
			POBoxCountryCode:                  data.POBoxCountryCode,
			POBoxCountryCodeText:              data.POBoxCountryCodeText,
			POBoxRegionCode:                   data.POBoxRegionCode,
			POBoxRegionCodeText:               data.POBoxRegionCodeText,
			POBoxCity:                         data.POBoxCity,
			TimeZoneCode:                      data.TimeZoneCode,
			TimeZoneCodeText:                  data.TimeZoneCodeText,
			TaxJurisdictionCode:               data.TaxJurisdictionCode,
			TaxJurisdictionCodeText:           data.TaxJurisdictionCodeText,
			Building:                          data.Building,
			Floor:                             data.Floor,
			Room:                              data.Room,
			InhouseMail:                       data.InhouseMail,
			OfficePhoneNumber:                 data.OfficePhoneNumber,
			MobilePhoneNumber:                 data.MobilePhoneNumber,
			FaxNumber:                         data.FaxNumber,
			Email:                             data.Email,
			NormalisedOfficePhoneNumber:       data.NormalisedOfficePhoneNumber,
			NormalisedMobilePhoneNumber:       data.NormalisedMobilePhoneNumber,
			CreatedOn:                         data.CreatedOn,
			CreatedBy:                         data.CreatedBy,
			CreatedByIdentityUUID:             data.CreatedByIdentityUUID,
			ChangedOn:                         data.ChangedOn,
			ChangedBy:                         data.ChangedBy,
			ChangedByIdentityUUID:             data.ChangedByIdentityUUID,
			EntityLastChangedOn:               data.EntityLastChangedOn,
			ETag:                              data.ETag,
		})
	}

	return serviceAgentCollection, nil
}