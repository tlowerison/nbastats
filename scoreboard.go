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
  if fields.DayOffset == nil { fields.DayOffset = &dayOffset.Default }
  if fields.LeagueID == nil  { fields.LeagueID  = &leagueID.Default }

  // Validate
  err := gameDate.Assert(fields.GameDate, true);    if err != nil { return nil, err }
  err = dayOffset.Assert(*fields.DayOffset, false); if err != nil { return nil, err }
  err = leagueID.Assert(*fields.LeagueID, false);   if err != nil { return nil, err }

  bytes, err := c.Get(model.FetchConfig{
    DataSource: "stats.nba.com",
    Endpoint:   "/scoreboardV2",
    Fields:     &map[string]string{
      "DayOffset": dayOffset.FromPtr(fields.DayOffset),
      "GameDate":  gameDate.FromVal(fields.GameDate),
      "LeagueID":  leagueID.FromPtr(fields.LeagueID),
    },
  })
  if err != nil {
    return nil, err
  }

  var body nbastatsResponseBody
  return body.UnmarshalResult(bytes)
}
