package pagination

type PageBasedPagination[E any] struct {
	Page    int  `json:"page"`
	PerPage int  `json:"per_page"`
	Total   int  `json:"total"`
	Items   []E  `json:"items"`
	HasMore bool `json:"has_more"`
}

type OffsetLimitPagination[E any] struct {
	Offset  int  `json:"offset"`
	Limit   int  `json:"limit"`
	Items   []E  `json:"items"`
	HasMore bool `json:"has_more"`
}

type CursorBasedPagination[C any, E any] struct {
	Current    C    `json:"current"`
	NextCursor C    `json:"next_cursor"`
	Size       int  `json:"size"`
	Items      []E  `json:"items"`
	HasMore    bool `json:"has_more"`
}

type TokenBasedPagination[T any, E any] struct {
	NextToken T    `json:"next_token"`
	Size      int  `json:"size"`
	Items     []E  `json:"items"`
	HasMore   bool `json:"has_more"`
}

type TimeBasedPagination[E any] struct {
	Before  int64 `json:"before"`
	After   int64 `json:"after"`
	Size    int   `json:"size"`
	Items   []E   `json:"items"`
	HasMore bool  `json:"has_more"`
}
