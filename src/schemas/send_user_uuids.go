package schemas

import (
  "github.com/google/uuid"
)

type SendUUIDS struct {
  UUIDS []uuid.UUID `json:"uuids"`
}
