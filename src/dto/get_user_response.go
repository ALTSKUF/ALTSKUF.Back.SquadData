package dto 

type GetUsersResponse struct {
  Users *[]User `json:"users,omitempty"`
  Error error `json:"error,omitempty"`
}
