package client

import (
  "encoding/json"

  "github.com/tlowerison/nbastats/model"
)

var (
  Counter = model.ArgInt{
    Name:   "Counter",
    Default: model.IntPtr(2160),
    Min:     model.IntPtr(1),
    Max:     model.IntPtr(10000),
  }
  DateFrom = model.ArgDateTime{Name: "DateFrom"}
  DayOffset = model.ArgInt{
    Name:   "DayOffset",
    Default: model.IntPtr(0),
    Min:     model.IntPtr(0),
  }
  DateTo = model.ArgDateTime{Name: "DateTo"}
  Direction = model.ArgEnum{
    Name:    "Direction",
    Default: model.StrPtr("DESC"),
    Options: []string{"DESC", "ASC"},
  }
  EndPeriod = model.ArgInt{
    Name:    "EndPeriod",
    Default: model.IntPtr(0),
    Min:     model.IntPtr(0),
    Max:     model.IntPtr(4),
  }
  GameDate = model.ArgDateTime{Name: "GameDate"}
  GameID = model.ArgString{
    Name: "GameID",
  }
  LeagueID = model.ArgEnum{
    Name:    "LeagueID",
    Default: model.StrPtr("00"),
    Options: []string{"00", "20"},
  }
  PlayerID = model.ArgString{
    Name: "PlayerID",
  }
  PlayerOrTeam = model.ArgEnum{
    Name:    "PlayerOrTeam",
    Options: []string{"T", "P"},
  }
  Season = model.ArgString{
    Name:    "Season",
    Default: model.StrPtr("2020-21"),
    Regexp:  model.StrPtr("\\d{4}-\\d{2}"),
  }
  SeasonType = model.ArgEnum{
    Name:    "SeasonType",
    Default: model.StrPtr("Regular+Season"),
    Options: []string{
      "Pre+Season",
      "Regular+Season",
      "Playoffs",
      "All+Star",
    },
  }
  Sorter = model.ArgEnum{
    Name: "Sorter",
    Default: model.StrPtr("DATE"),
    Options: []string{"DATE"},
  }
  StartPeriod = model.ArgInt{
    Name:    "StartPeriod",
    Default: model.IntPtr(0),
    Min:     model.IntPtr(0),
    Max:     model.IntPtr(4),
  }
  TeamID = model.ArgString{
    Name: "TeamID",
  }
)

type NbaStatsClient struct {
  *model.Client
}

func NewClient() *NbaStatsClient {
  return &NbaStatsClient{&model.Client{
    Name:    "nbastats",
    BaseUrl: "https://stats.nba.com/stats",
    BaseHeaders: map[string]string{
      "User-Agent":         "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.16; rv:86.0) Gecko/20100101 Firefox/86.0",
      "Accept":             "application/json, text/plain, */*",
      "Accept-Language":    "en-US,en;q=0.5",
      "Connection":         "keep-alive",
      "DNT":                "1",
      "Origin":             "https://www.nba.com",
      "x-nba-stats-origin": "stats",
      "x-nba-stats-token":  "true",
      "Referer":            "https://www.nba.com/stats/",
      "Sec-GPC":            "1",
    },
  }}
}

type NbaStatsResultSet struct {
  Headers []string      `json:"headers"`
  Name    string        `json:"name"`
  RowSet  []interface{} `json:"rowSet"`
}

type NbaStatsResponseBody struct {
  ResultSets []NbaStatsResultSet `json:"resultSets"`
}

func bytesToResultSets(bytes []byte) ([]NbaStatsResultSet, error) {
  var body NbaStatsResponseBody
  err := json.Unmarshal(bytes, &body)
  if err != nil {
    return nil, err
  }
  return body.ResultSets, nil
}
