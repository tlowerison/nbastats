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
  if fields.Counter == nil    { fields.Counter    = &counter.Default }
  if fields.DateFrom == nil   { fields.DateFrom   = &dateFrom.Default }
  if fields.DateTo == nil     { fields.DateTo     = &dateTo.Default }
  if fields.Direction == nil  { fields.Direction  = &direction.Default }
  if fields.LeagueID == nil   { fields.LeagueID   = &leagueID.Default }
  if fields.Season == nil     { fields.Season     = &season.Default }
  if fields.SeasonType == nil { fields.SeasonType = &seasonType.Default }
  if fields.Sorter == nil     { fields.Sorter     = &sorter.Default }

  // Validate
  err := playerID.Assert(fields.PlayerID);     if err != nil { return nil, err }
  err = counter.Assert(*fields.Counter);       if err != nil { return nil, err }
  err = dateFrom.Assert(*fields.DateFrom);     if err != nil { return nil, err }
  err = dateTo.Assert(*fields.DateTo);         if err != nil { return nil, err }
  err = direction.Assert(*fields.Direction);   if err != nil { return nil, err }
  err = leagueID.Assert(*fields.LeagueID);     if err != nil { return nil, err }
  err = season.Assert(*fields.Season);         if err != nil { return nil, err }
  err = seasonType.Assert(*fields.SeasonType); if err != nil { return nil, err }
  err = sorter.Assert(*fields.Sorter);         if err != nil { return nil, err }

  bytes, err := c.Get(model.FetchConfig{
    DataSource: "stats.nba.com",
    Endpoint:   "/playergamelog",
    Fields:     &map[string]string{
      "Counter":      counter.FromPtr(fields.Counter),
      "DateFrom":     dateFrom.FromPtr(fields.DateFrom),
      "DateTo":       dateTo.FromPtr(fields.DateTo),
      "Direction":    direction.FromPtr(fields.Direction),
      "LeagueID":     leagueID.FromPtr(fields.LeagueID),
      "PlayerID":     playerID.FromVal(fields.PlayerID),
      "Season":       season.FromPtr(fields.Season),
      "SeasonType":   seasonType.FromPtr(fields.SeasonType),
      "Sorter":       sorter.FromPtr(fields.Sorter),
    },
  })
  if err != nil {
    return nil, err
  }

  var body nbastatsResponseBody
  return body.UnmarshalResult(bytes)
}
