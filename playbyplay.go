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
  if fields.EndPeriod == nil   { fields.EndPeriod   = &endPeriod.Default }
  if fields.StartPeriod == nil { fields.StartPeriod = &startPeriod.Default }

  // Validate
  err := endPeriod.Assert(*fields.EndPeriod);    if err != nil { return nil, err }
  err = gameID.Assert(fields.GameID);            if err != nil { return nil, err }
  err = startPeriod.Assert(*fields.StartPeriod); if err != nil { return nil, err }

  bytes, err := c.Get(model.FetchConfig{
    DataSource: "stats.nba.com",
    Endpoint:   "/playbyplayv2",
    Fields:     &map[string]string{
      "EndPeriod":   endPeriod.FromPtr(fields.EndPeriod),
      "GameID":      gameID.FromVal(fields.GameID),
      "StartPeriod": startPeriod.FromPtr(fields.StartPeriod),
    },
  })
  if err != nil {
    return nil, err
  }

  var body nbastatsResponseBody
  return body.UnmarshalResult(bytes)
}
