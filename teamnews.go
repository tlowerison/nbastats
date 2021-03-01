package nbastats

import (
  "encoding/json"
  "fmt"

  "github.com/tlowerison/nbastats/model"
)

type TeamNewsFields struct {
  // Required
  TeamID string
}

var (
  headers = []string{
    "Date",
    "FirstName",
    "Headline",
    "Injured",
    "InjuredStatus",
    "InjuryDetail",
    "InjuryLocation",
    "InjurySide",
    "InjuryType",
    "LastName",
    "LastUpdate",
    "ListItemCaption",
    "ListItemDescription",
    "ListItemPubDate",
    "PlayerCode",
    "PlayerID",
    "Position",
    "Priority",
    "RotoId",
    "Team",
    "TeamCode",
    "UpdateId",
  }
  teamIDToName = map[string]string{
    "1610612737": "hawks",
    "1610612738": "celtics",
    "1610612751": "nets",
    "1610612766": "hornets",
    "1610612741": "bulls",
    "1610612739": "cavaliers",
    "1610612742": "mavericks",
    "1610612743": "nuggets",
    "1610612765": "pistons",
    "1610612744": "warriors",
    "1610612745": "rockets",
    "1610612754": "pacers",
    "1610612746": "clippers",
    "1610612747": "lakers",
    "1610612763": "grizzlies",
    "1610612748": "heat",
    "1610612749": "bucks",
    "1610612750": "timberwolves",
    "1610612740": "pelicans",
    "1610612752": "knicks",
    "1610612760": "thunder",
    "1610612753": "magic",
    "1610612755": "76ers",
    "1610612756": "suns",
    "1610612757": "blazers",
    "1610612758": "kings",
    "1610612759": "spurs",
    "1610612761": "raptors",
    "1610612762": "jazz",
    "1610612764": "wizards",
  }
)

type rotowireListItem struct {
  Date                string `json:"Date"`
  FirstName           string `json:"FirstName"`
  Headline            string `json:"Headline"`
  Injured             string `json:"Injured"`
  InjuredStatus       string `json:"Injured_Status"`
  InjuryDetail        string `json:"Injury_Detail"`
  InjuryLocation      string `json:"Injury_Location"`
  InjurySide          string `json:"Injury_Side"`
  InjuryType          string `json:"Injury_Type"`
  LastName            string `json:"LastName"`
  LastUpdate          string `json:"lastUpdate"`
  ListItemCaption     string `json:"ListItemCaption"`
  ListItemDescription string `json:"ListItemDescription"`
  ListItemPubDate     string `json:"ListItemPubDate"`
  PlayerCode          string `json:"player_code"`
  PlayerID            string `json:"PlayerID"`
  Position            string `json:"Position"`
  Priority            string `json:"Priority"`
  RotoId              string `json:"RotoId"`
  Team                string `json:"Team"`
  TeamCode            string `json:"TeamCode"`
  UpdateId            string `json:"UpdateId"`
}

type rotowireResponseBody struct {
  ListItems []rotowireListItem `json:"ListItems"`
}

func (body rotowireResponseBody) UnmarshalResult(bytes []byte) (*model.Result, error) {
  err := json.Unmarshal(bytes, &body)
  if err != nil {
    return nil, err
  }
  rows := make([][]interface{}, len(body.ListItems))
  for i, listItem := range body.ListItems {
    rows[i] = []interface{}{
      listItem.Date,
      listItem.FirstName,
      listItem.Headline,
      listItem.Injured,
      listItem.InjuredStatus,
      listItem.InjuryDetail,
      listItem.InjuryLocation,
      listItem.InjurySide,
      listItem.InjuryType,
      listItem.LastName,
      listItem.LastUpdate,
      listItem.ListItemCaption,
      listItem.ListItemDescription,
      listItem.ListItemPubDate,
      listItem.PlayerCode,
      listItem.PlayerID,
      listItem.Position,
      listItem.Priority,
      listItem.RotoId,
      listItem.Team,
      listItem.TeamCode,
      listItem.UpdateId,
    }
  }
  return &model.Result{Headers: headers, Rows: rows}, nil
}

func (c *NbaStatsClient) TeamNews(fields *TeamNewsFields) (*model.Result, error) {
  // Validate
  err := TeamID.Assert(fields.TeamID); if err != nil { return nil, err }
  if _, ok := teamIDToName[fields.TeamID]; !ok {
    return nil, fmt.Errorf("Invalid TeamID, received: %s", fields.TeamID)
  }

  teamName := teamIDToName[fields.TeamID]

  bytes, err := c.Get(model.FetchConfig{
    DataSource: "stats-prod.nba.com",
    Endpoint:   "/rotowire/player",
    Fields:     &map[string]string{"Team": teamName},
  })
  if err != nil {
    return nil, err
  }

  var body rotowireResponseBody
  return body.UnmarshalResult(bytes)
}
