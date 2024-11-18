package dattoapi

import (
	"encoding/base64"

	"github.com/simonbuckner/goquadac"
)

type DattoApi struct {
	*goquadac.ApiHelper
}

func NewDattoApi(baseUrl string) *DattoApi {
	api := goquadac.NewApiHelper(baseUrl).SetDefaultHeader("accept", "application/json")
	return &DattoApi{
		ApiHelper: api,
	}
}

func (api *DattoApi) Authenticate(apiKey, apiSecret string) error {
	api.SetAuthHeader("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(apiKey+":"+apiSecret)))
	return nil
}

// SaaS Queries
func (api *DattoApi) GetSaasDomains() *SaasDomainQuery {
	return newSaaSDomainsQuery(api)
}

func (api *DattoApi) GetSaasSeatsForCustomerId(customerId int64) *SaasSeatsQuery {
	return newSaasSeatsQuery(api, customerId)
}

func (api *DattoApi) GetSaasApplicationsForCustomerId(customerId int64) *SaasApplicationsQuery {
	return newSaasApplicationsQuery(api, customerId)
}
