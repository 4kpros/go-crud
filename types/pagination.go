package types

type Pagination struct {
	CurrentPage  int
	NextPage     int
	PreviousPage int
	TotalPages   int
	Count        int
	Limit        int
}
