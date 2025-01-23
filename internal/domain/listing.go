package domain

import "time"

type Listing struct {
  ID          int
  UserID      int
  PropertyID  int
  Title       string
  Description string
  Status      int
  CreatedAt   time.Time
  UpdatedAt   time.Time
}
