package dto 

type GetUsersResponse struct {
  Users *[]User `json:"users"`
  Error error `json:"error,omitempty"`
}
