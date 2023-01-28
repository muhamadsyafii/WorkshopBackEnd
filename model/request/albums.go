package request

type Albums struct {
	Name string `json:"name" binding:"required"`
	Year int    `json:"year" binding:"required"`
	//CreatedAt time.Time `json:"created_at" binding:"required"`
	//UpdatedAt time.Time `json:"updated_at" binding:"required"`
}
