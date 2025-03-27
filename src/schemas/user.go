package schemas

type User struct {
  FullName string `json:"fullName"`
  Group string `json:"group"`
  Role string `json:"role"`
}
