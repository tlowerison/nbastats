package nbastats

import "github.com/tlowerison/nbastats/model"

type CommonAllPlayersFields struct {
  // Optional
  IsOnlyCurrentSeason *string
  LeagueID            *string
  Season              *string
}

func (c *NbaStatsClient) CommonAllPlayers(fields *CommonAllPlayersFields) (*model.Result, error) {
  if fields.IsOnlyCurrentSeason == nil { fields.IsOnlyCurrentSeason = &IsOnlyCurrentSeason.Default }
  if fields.LeagueID == nil            { fields.LeagueID            = &LeagueID.Default }
  if fields.Season == nil              { fields.Season              = &Season.Default }

  // Validate
  err := IsOnlyCurrentSeason.Assert(*fields.IsOnlyCurrentSeason); if err != nil { return nil, err }
  err = LeagueID.Assert(*fields.LeagueID);                        if err != nil { return nil, err }
  err = Season.Assert(*fields.Season);                            if err != nil { return nil, err }

  bytes, err := c.Get(model.FetchConfig{
    DataSource: "stats.nba.com",
    Endpoint:   "/commonallplayers",
    Fields:     &map[string]string{
      "IsOnlyCurrentSeason": IsOnlyCurrentSeason.FromPtr(fields.IsOnlyCurrentSeason),
      "LeagueID":            LeagueID.FromPtr(fields.LeagueID),
      "Season":              Season.FromPtr(fields.Season),
    },
  })
  if err != nil {
    return nil, err
  }

  var body nbastatsResponseBody
  return body.UnmarshalResult(bytes)
}
