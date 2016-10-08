package request

type PresentationRequest struct {
	Filters          []Filter
	PaginationOption PaginationOption
	SortingOption    SortingOption
}
