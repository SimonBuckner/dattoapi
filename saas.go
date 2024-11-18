package dattoapi

import "github.com/simonbuckner/goquadac"

type SaasDomainQuery struct {
	api   *DattoApi
	query *goquadac.ApiQuery
}

func newSaaSDomainsQuery(api *DattoApi) *SaasDomainQuery {
	return &SaasDomainQuery{
		api: api,
	}
}

func (q *SaasDomainQuery) Build() (*SaasDomainQuery, error) {
	endpoint := "v1/saas/domains"

	query := q.api.NewGetQuery(endpoint)

	q.query = query
	return q, nil
}

func (q *SaasDomainQuery) GetAll() ([]SaasDomain, error) {
	if q.query == nil {
		if _, err := q.Build(); err != nil {
			return nil, err
		}
	}
	var out []SaasDomain
	err := q.query.Get(&out)
	return out, err
}

type SaasSeatsQuery struct {
	api        *DattoApi
	query      *goquadac.ApiQuery
	customerId int64
}

func newSaasSeatsQuery(api *DattoApi, customerId int64) *SaasSeatsQuery {
	return &SaasSeatsQuery{
		api:        api,
		customerId: customerId,
	}
}

func (q *SaasSeatsQuery) Build() (*SaasSeatsQuery, error) {

	endpoint := "v1/saas/" + goquadac.I64toString(q.customerId) + "/seats"

	query := q.api.NewGetQuery(endpoint)

	q.query = query
	return q, nil
}

func (q *SaasSeatsQuery) GetAll() ([]SaasSeat, error) {
	if q.query == nil {
		if _, err := q.Build(); err != nil {
			return nil, err
		}
	}
	var out []SaasSeat
	err := q.query.Get(&out)
	return out, err
}

type SaasApplicationsQuery struct {
	api        *DattoApi
	query      *goquadac.ApiQuery
	customerId int64
}

func newSaasApplicationsQuery(api *DattoApi, customerId int64) *SaasApplicationsQuery {
	return &SaasApplicationsQuery{
		api:        api,
		customerId: customerId,
	}
}

func (q *SaasApplicationsQuery) Build() (*SaasApplicationsQuery, error) {

	endpoint := "v1/saas/" + goquadac.I64toString(q.customerId) + "/applications"

	query := q.api.NewGetQuery(endpoint)

	q.query = query
	return q, nil
}

func (q *SaasApplicationsQuery) Get() (*SaasBackup, error) {
	if q.query == nil {
		if _, err := q.Build(); err != nil {
			return nil, err
		}
	}
	var out SaasBackup
	err := q.query.Get(&out)
	return &out, err
}
