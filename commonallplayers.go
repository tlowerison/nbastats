package nbastats

import "github.com/tlowerison/nbastats/model"

type CommonAllPlayersFields struct {
  // Optional
  IsOnlyCurrentSeason *string
  LeagueID            *string
  Season              *string
}

func (c *NbaStatsClient) CommonAllPlayers(fields *CommonAllPlayersFields) (*model.Result, error) {
  if fields.IsOnlyCurrentSeason == nil { fields.IsOnlyCurrentSeason = &isOnlyCurrentSeason.Default }
  if fields.LeagueID == nil            { fields.LeagueID            = &leagueID.Default }
  if fields.Season == nil              { fields.Season              = &season.Default }

  // Validate
  err := isOnlyCurrentSeason.Assert(*fields.IsOnlyCurrentSeason, false); if err != nil { return nil, err }
  err = leagueID.Assert(*fields.LeagueID, false);                        if err != nil { return nil, err }
  err = season.Assert(*fields.Season, false);                            if err != nil { return nil, err }

  bytes, err := c.Get(model.FetchConfig{
    DataSource: "stats.nba.com",
    Endpoint:   "/commonallplayers",
    Fields:     &map[string]string{
      "IsOnlyCurrentSeason": isOnlyCurrentSeason.FromPtr(fields.IsOnlyCurrentSeason),
      "LeagueID":            leagueID.FromPtr(fields.LeagueID),
      "Season":              season.FromPtr(fields.Season),
    },
  })
  if err != nil {
    return nil, err
  }

  var body nbastatsResponseBody
  return body.UnmarshalResult(bytes)
}
