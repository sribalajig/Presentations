package request

type Request struct {
	Filters          *[]Filter
	PaginationOption *PaginationOption
	SortingOption    *SortingOption
}
