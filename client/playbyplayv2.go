package client

import (
  "strconv"

  "github.com/tlowerison/nbastats/model"
)

type Playbyplayv2Fields struct {
  // Required
  GameID      string
  // Optional
  EndPeriod   *int
  StartPeriod *int
}

func (c *NbaStatsClient) Playbyplayv2(fields *Playbyplayv2Fields) ([]NbaStatsResultSet, error) {
  if fields.EndPeriod == nil   { fields.EndPeriod   = EndPeriod.Default }
  if fields.StartPeriod == nil { fields.StartPeriod = StartPeriod.Default }

  // Validate
  err := EndPeriod.Assert(*fields.EndPeriod);    if err != nil { return nil, err }
  err = GameID.Assert(fields.GameID);            if err != nil { return nil, err }
  err = StartPeriod.Assert(*fields.StartPeriod); if err != nil { return nil, err }

  bytes, err := c.Fetch(model.FetchConfig{
    Endpoint: "/playbyplayv2",
    Fields: &map[string]string{
      "EndPeriod":   strconv.Itoa(*fields.EndPeriod),
      "GameID":      fields.GameID,
      "StartPeriod": strconv.Itoa(*fields.StartPeriod),
    },
  })
  if err != nil {
    return nil, err
  }

  return bytesToResultSets(bytes)
}
