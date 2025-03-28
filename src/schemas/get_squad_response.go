package schemas

type GetSquadResponse struct {
  Name string `json:"name,omitempty"`
  Description string `json:"description,omitempty"`
	Error error `json:"error,omitempty" gorm:"-"`
}
