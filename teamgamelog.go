package nbastats

import "github.com/tlowerison/nbastats/model"

type TeamGameLogFields struct {
  // Required
  TeamID     string
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

func (c *NbaStatsClient) TeamGameLog(fields *TeamGameLogFields) (*model.Result, error) {
  if fields.Counter == nil    { fields.Counter    = &counter.Default }
  if fields.DateFrom == nil   { fields.DateFrom   = &dateFrom.Default }
  if fields.DateTo == nil     { fields.DateTo     = &dateTo.Default }
  if fields.Direction == nil  { fields.Direction  = &direction.Default }
  if fields.LeagueID == nil   { fields.LeagueID   = &leagueID.Default }
  if fields.Season == nil     { fields.Season     = &season.Default }
  if fields.SeasonType == nil { fields.SeasonType = &seasonType.Default }
  if fields.Sorter == nil     { fields.Sorter     = &sorter.Default }

  // Validate
  err := teamID.Assert(fields.TeamID, true);          if err != nil { return nil, err }
  err = counter.Assert(*fields.Counter, false);       if err != nil { return nil, err }
  err = dateFrom.Assert(*fields.DateFrom, false);     if err != nil { return nil, err }
  err = dateTo.Assert(*fields.DateTo, false);         if err != nil { return nil, err }
  err = direction.Assert(*fields.Direction, false);   if err != nil { return nil, err }
  err = leagueID.Assert(*fields.LeagueID, false);     if err != nil { return nil, err }
  err = season.Assert(*fields.Season, false);         if err != nil { return nil, err }
  err = seasonType.Assert(*fields.SeasonType, false); if err != nil { return nil, err }
  err = sorter.Assert(*fields.Sorter, false);         if err != nil { return nil, err }

  bytes, err := c.Get(model.FetchConfig{
    DataSource: "stats.nba.com",
    Endpoint:   "/teamgamelog",
    Fields:     &map[string]string{
      "Counter":      counter.FromPtr(fields.Counter),
      "DateFrom":     dateFrom.FromPtr(fields.DateFrom),
      "DateTo":       dateTo.FromPtr(fields.DateTo),
      "Direction":    direction.FromPtr(fields.Direction),
      "LeagueID":     leagueID.FromPtr(fields.LeagueID),
      "Season":       season.FromPtr(fields.Season),
      "SeasonType":   seasonType.FromPtr(fields.SeasonType),
      "Sorter":       sorter.FromPtr(fields.Sorter),
      "TeamID":       teamID.FromValAsInt(fields.TeamID),
    },
  })
  if err != nil {
    return nil, err
  }

  var body nbastatsResponseBody
  return body.UnmarshalResult(bytes)
}
