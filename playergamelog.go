package nbastats

import "github.com/tlowerison/nbastats/model"

type PlayerGameLogFields struct {
  // Required
  PlayerID   string
  // Optional
  Counter    *int
  DateFrom   *string
  DateTo     *string
  Direction  *string
  LeagueID   *string
  Season     *string
  SeasonType *string
  Sorter     *string
}

func (c *NbaStatsClient) PlayerGameLog(fields *PlayerGameLogFields) (*model.Result, error) {
  if fields.Counter == nil    { fields.Counter    = &Counter.Default }
  if fields.DateFrom == nil   { fields.DateFrom   = &DateFrom.Default }
  if fields.DateTo == nil     { fields.DateTo     = &DateTo.Default }
  if fields.Direction == nil  { fields.Direction  = &Direction.Default }
  if fields.LeagueID == nil   { fields.LeagueID   = &LeagueID.Default }
  if fields.Season == nil     { fields.Season     = &Season.Default }
  if fields.SeasonType == nil { fields.SeasonType = &SeasonType.Default }
  if fields.Sorter == nil     { fields.Sorter     = &Sorter.Default }

  // Validate
  err := PlayerID.Assert(fields.PlayerID);     if err != nil { return nil, err }
  err = Counter.Assert(*fields.Counter);       if err != nil { return nil, err }
  err = DateFrom.Assert(*fields.DateFrom);     if err != nil { return nil, err }
  err = DateTo.Assert(*fields.DateTo);         if err != nil { return nil, err }
  err = Direction.Assert(*fields.Direction);   if err != nil { return nil, err }
  err = LeagueID.Assert(*fields.LeagueID);     if err != nil { return nil, err }
  err = Season.Assert(*fields.Season);         if err != nil { return nil, err }
  err = SeasonType.Assert(*fields.SeasonType); if err != nil { return nil, err }
  err = Sorter.Assert(*fields.Sorter);         if err != nil { return nil, err }

  bytes, err := c.Get(model.FetchConfig{
    DataSource: "stats.nba.com",
    Endpoint:   "/playergamelog",
    Fields:     &map[string]string{
      "Counter":      Counter.FromPtr(fields.Counter),
      "DateFrom":     DateFrom.FromPtr(fields.DateFrom),
      "DateTo":       DateTo.FromPtr(fields.DateTo),
      "Direction":    Direction.FromPtr(fields.Direction),
      "LeagueID":     LeagueID.FromPtr(fields.LeagueID),
      "PlayerID":     PlayerID.FromVal(fields.PlayerID),
      "Season":       Season.FromPtr(fields.Season),
      "SeasonType":   SeasonType.FromPtr(fields.SeasonType),
      "Sorter":       Sorter.FromPtr(fields.Sorter),
    },
  })
  if err != nil {
    return nil, err
  }

  var body nbastatsResponseBody
  return body.UnmarshalResult(bytes)
}
