package domain

import "time"

type House struct {
  ID              int
  BuiltIN         int
  Address         string
  Developer       string
  CreatedAt       time.Time
  LastFlatUpdate time.Time
}
