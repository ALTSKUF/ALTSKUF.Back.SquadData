package schemas

type User struct {
  FullName string `json:"full_name"`
  Group string `json:"group"`
  Role string `json:"role"`
}
