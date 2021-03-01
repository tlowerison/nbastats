package nbastats

import "github.com/tlowerison/nbastats/model"

type PlayByPlayFields struct {
  // Required
  GameID      string
  // Optional
  EndPeriod   *int
  StartPeriod *int
}

func (c *NbaStatsClient) PlayByPlay(fields *PlayByPlayFields) (*model.Result, error) {
  if fields.EndPeriod == nil   { fields.EndPeriod   = &EndPeriod.Default }
  if fields.StartPeriod == nil { fields.StartPeriod = &StartPeriod.Default }

  // Validate
  err := EndPeriod.Assert(*fields.EndPeriod);    if err != nil { return nil, err }
  err = GameID.Assert(fields.GameID);            if err != nil { return nil, err }
  err = StartPeriod.Assert(*fields.StartPeriod); if err != nil { return nil, err }

  bytes, err := c.Get(model.FetchConfig{
    DataSource: "stats.nba.com",
    Endpoint:   "/playbyplayv2",
    Fields:     &map[string]string{
      "EndPeriod":   EndPeriod.FromPtr(fields.EndPeriod),
      "GameID":      GameID.FromVal(fields.GameID),
      "StartPeriod": StartPeriod.FromPtr(fields.StartPeriod),
    },
  })
  if err != nil {
    return nil, err
  }

  var body nbastatsResponseBody
  return body.UnmarshalResult(bytes)
}
