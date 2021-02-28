package client

import (
  "strconv"

  "github.com/tlowerison/nbastats/model"
)

type LeaguegamelogFields struct {
  // Required
  PlayerOrTeam string
  // Optional
  Counter      *int
  DateFrom     *string
  DateTo       *string
  Direction    *string
  LeagueID     *string
  Season       *string
  SeasonType   *string
  Sorter       *string
}

func (c *NbaStatsClient) Leaguegamelog(fields *LeaguegamelogFields) ([]NbaStatsResultSet, error) {
  if fields.Counter == nil    { fields.Counter    = Counter.Default }
  if fields.DateFrom == nil   { fields.DateFrom   = model.StrPtr("") }
  if fields.DateTo == nil     { fields.DateTo     = model.StrPtr("") }
  if fields.Direction == nil  { fields.Direction  = Direction.Default }
  if fields.LeagueID == nil   { fields.LeagueID   = LeagueID.Default }
  if fields.Season == nil     { fields.Season     = Season.Default }
  if fields.SeasonType == nil { fields.SeasonType = SeasonType.Default }
  if fields.Sorter == nil     { fields.Sorter     = Sorter.Default }

  // Validate
  err := PlayerOrTeam.Assert(fields.PlayerOrTeam); if err != nil { return nil, err }
  err = Counter.Assert(*fields.Counter);           if err != nil { return nil, err }
  err = DateFrom.Assert(*fields.DateFrom);         if err != nil { return nil, err }
  err = DateTo.Assert(*fields.DateTo);             if err != nil { return nil, err }
  err = Direction.Assert(*fields.Direction);       if err != nil { return nil, err }
  err = LeagueID.Assert(*fields.LeagueID);         if err != nil { return nil, err }
  err = Season.Assert(*fields.Season);             if err != nil { return nil, err }
  err = SeasonType.Assert(*fields.SeasonType);     if err != nil { return nil, err }
  err = Sorter.Assert(*fields.Sorter);             if err != nil { return nil, err }

  bytes, err := c.Fetch(model.FetchConfig{
    Endpoint: "/leaguegamelog",
    Fields: &map[string]string{
      "Counter":      strconv.Itoa(*fields.Counter),
      "DateFrom":     *fields.DateFrom,
      "DateTo":       *fields.DateTo,
      "Direction":    *fields.Direction,
      "LeagueID":     *fields.LeagueID,
      "PlayerOrTeam": fields.PlayerOrTeam,
      "Season":       *fields.Season,
      "SeasonType":   *fields.SeasonType,
      "Sorter":       *fields.Sorter,
    },
  })
  if err != nil {
    return nil, err
  }

  return bytesToResultSets(bytes)
}
