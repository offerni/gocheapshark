package deal

type Service interface {
	List(opts DealListOpts) (*DealListResponse, error)
	Lookup(opts DealLookupOpts) (*DealLookupResponse, error)
}
