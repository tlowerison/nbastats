package nbastats

import "github.com/tlowerison/nbastats/model"

type CommonPlayerInfoFields struct {
  // Required
  PlayerID string
}

func (c *NbaStatsClient) CommonPlayerInfo(fields *CommonPlayerInfoFields) (*model.Result, error) {
  // Validate
  err := playerID.Assert(fields.PlayerID, true); if err != nil { return nil, err }

  bytes, err := c.Get(model.FetchConfig{
    DataSource: "stats.nba.com",
    Endpoint:   "/commonplayerinfo",
    Fields:     &map[string]string{
      "PlayerID": playerID.FromVal(fields.PlayerID),
    },
  })
  if err != nil {
    return nil, err
  }

  var body nbastatsResponseBody
  return body.UnmarshalResult(bytes)
}
