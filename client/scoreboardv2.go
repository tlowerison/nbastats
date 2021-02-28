package client

import (
  "strconv"

  "github.com/tlowerison/nbastats/model"
)

type Scoreboardv2Fields struct {
  // Required
  GameDate  string
  // Optional
  DayOffset *int
  LeagueID  *string
}

func (c *NbaStatsClient) Scoreboardv2(fields *Scoreboardv2Fields) ([]NbaStatsResultSet, error) {
  if fields.DayOffset == nil { fields.DayOffset = model.IntPtr(0) }
  if fields.LeagueID == nil  { fields.LeagueID  = LeagueID.Default }

  // Validate
  err := DayOffset.Assert(*fields.DayOffset); if err != nil { return nil, err }
  err = GameDate.Assert(fields.GameDate);     if err != nil { return nil, err }
  err = LeagueID.Assert(*fields.LeagueID);    if err != nil { return nil, err }

  bytes, err := c.Fetch(model.FetchConfig{
    Endpoint: "/scoreboardV2",
    Fields: &map[string]string{
      "DayOffset": strconv.Itoa(*fields.DayOffset),
      "GameDate":  fields.GameDate,
      "LeagueID":  *fields.LeagueID,
    },
  })
  if err != nil {
    return nil, err
  }

  return bytesToResultSets(bytes)
}
