package nbastats

import "github.com/tlowerison/nbastats/model"

type ScoreboardFields struct {
  // Required
  GameDate  string
  // Optional
  DayOffset *int
  LeagueID  *string
}

func (c *NbaStatsClient) Scoreboard(fields *ScoreboardFields) (*model.Result, error) {
  if fields.DayOffset == nil { fields.DayOffset = &DayOffset.Default }
  if fields.LeagueID == nil  { fields.LeagueID  = &LeagueID.Default }

  // Validate
  err := DayOffset.Assert(*fields.DayOffset); if err != nil { return nil, err }
  err = GameDate.Assert(fields.GameDate);     if err != nil { return nil, err }
  err = LeagueID.Assert(*fields.LeagueID);    if err != nil { return nil, err }

  bytes, err := c.Get(model.FetchConfig{
    DataSource: "stats.nba.com",
    Endpoint:   "/scoreboardV2",
    Fields:     &map[string]string{
      "DayOffset": DayOffset.FromPtr(fields.DayOffset),
      "GameDate":  GameDate.FromVal(fields.GameDate),
      "LeagueID":  LeagueID.FromPtr(fields.LeagueID),
    },
  })
  if err != nil {
    return nil, err
  }

  var body nbastatsResponseBody
  return body.UnmarshalResult(bytes)
}
