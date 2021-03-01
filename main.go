package nbastats

import (
  "encoding/json"

  "github.com/tlowerison/nbastats/model"
)

type NbaStatsClient struct {
  *model.Client
}

func NewClient(shouldLog bool) *NbaStatsClient {
  return &NbaStatsClient{&model.Client{
    DataSources: map[string]model.DataSource{
      "stats.nba.com": {
        BaseUrl: "https://stats.nba.com/stats",
        BaseHeaders: map[string]string{
          "User-Agent":         "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.16; rv:86.0) Gecko/20100101 Firefox/86.0",
          "Accept":             "application/json, text/plain, */*",
          "Accept-Language":    "en-US,en;q=0.5",
          "DNT":                "1",
          "Host":               "www.nba.com",
          "Origin":             "https://www.nba.com",
          "x-nba-stats-origin": "stats",
          "x-nba-stats-token":  "true",
          "Referer":            "https://www.nba.com/stats/",
          "Sec-GPC":            "1",
        },
      },
      "stats-prod.nba.com": {
        BaseUrl: "https://stats-prod.nba.com/wp-json/statscms/v1",
        BaseHeaders: map[string]string{
          "User-Agent":         "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.16; rv:86.0) Gecko/20100101 Firefox/86.0",
          "Accept":             "application/json, text/plain, */*",
          "Accept-Language":    "en-US,en;q=0.5",
          "DNT":                "1",
          "Host":               "stats-prod.nba.com",
          "Origin":             "https://www.nba.com",
          "x-nba-stats-origin": "stats",
          "x-nba-stats-token":  "true",
          "Referer":            "https://www.nba.com/",
          "Sec-GPC":            "1",
        },
      },
    },
    ShouldLog: shouldLog,
  }}
}

type nbastatsResponseBody struct {
  ResultSets []model.Result `json:"resultSets"`
}

func (body nbastatsResponseBody) UnmarshalResult(bytes []byte) (*model.Result, error) {
  err := json.Unmarshal(bytes, &body)
  if err != nil {
    return nil, err
  }
  if len(body.ResultSets) == 0 {
    return nil, nil
  }
  return &body.ResultSets[0], nil
}
