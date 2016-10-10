package response

type PaginatedResponse struct {
	TotalRecords int         `json:"totalRecords"`
	TotalPages   int         `json:"totalPages"`
	Results      interface{} `json:"results"`
}
