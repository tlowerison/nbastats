package nbastats

import "github.com/tlowerison/nbastats/model"

type LineupsFields struct {
  // Optional
  Conference     *string
  DateFrom       *string
  DateTo         *string
  Division       *string
  GameID         *string
  GameSegment    *string
  GroupQuantity  *int
  LastNGames     *int
  LeagueID       *string
  Location       *string
  MeasureType    *string
  Month          *int
  OpponentTeamID *string
  Outcome        *string
  PORound        *int
  PaceAdjust     *string
  PerMode        *string
  Period         *int
  PlusMinus      *string
  Rank           *string
  Season         *string
  SeasonSegment  *string
  SeasonType     *string
  ShotClockRange *string
  TeamID         *string
  VsConference   *string
  VsDivision     *string
}

func (c *NbaStatsClient) Lineups(fields *LineupsFields) (*model.Result, error) {
  if fields.Conference == nil     { fields.Conference     = &conference.Default }
  if fields.DateFrom == nil       { fields.DateFrom       = &dateFrom.Default }
  if fields.DateTo == nil         { fields.DateTo         = &dateTo.Default }
  if fields.Division == nil       { fields.Division       = &division.Default }
  if fields.GameID == nil         { fields.GameID         = &gameID.Default }
  if fields.GameSegment == nil    { fields.GameSegment    = &gameSegment.Default }
  if fields.GroupQuantity == nil  { fields.GroupQuantity  = &groupQuantity.Default }
  if fields.LastNGames == nil     { fields.LastNGames     = &lastNGames.Default }
  if fields.LeagueID == nil       { fields.LeagueID       = &leagueID.Default }
  if fields.Location == nil       { fields.Location       = &location.Default }
  if fields.MeasureType == nil    { fields.MeasureType    = &measureType.Default }
  if fields.Month == nil          { fields.Month          = &month.Default }
  if fields.OpponentTeamID == nil { fields.OpponentTeamID = &opponentTeamID.Default }
  if fields.Outcome == nil        { fields.Outcome        = &outcome.Default }
  if fields.PORound == nil        { fields.PORound        = &pORound.Default }
  if fields.PaceAdjust == nil     { fields.PaceAdjust     = &paceAdjust.Default }
  if fields.PerMode == nil        { fields.PerMode        = &perMode.Default }
  if fields.Period == nil         { fields.Period         = &period.Default }
  if fields.PlusMinus == nil      { fields.PlusMinus      = &plusMinus.Default }
  if fields.Rank == nil           { fields.Rank           = &rank.Default }
  if fields.Season == nil         { fields.Season         = &season.Default }
  if fields.SeasonSegment == nil  { fields.SeasonSegment  = &seasonSegment.Default }
  if fields.SeasonType == nil     { fields.SeasonType     = &seasonType.Default }
  if fields.ShotClockRange == nil { fields.ShotClockRange = &shotClockRange.Default }
  if fields.TeamID == nil         { fields.TeamID         = &teamID.Default }
  if fields.VsConference == nil   { fields.VsConference   = &vsConference.Default }
  if fields.VsDivision == nil     { fields.VsDivision     = &vsDivision.Default }

  // Validate
  err := conference.Assert(*fields.Conference, false);        if err != nil { return nil, err }
  err = dateFrom.Assert(*fields.DateFrom, false);             if err != nil { return nil, err }
  err = dateTo.Assert(*fields.DateTo, false);                 if err != nil { return nil, err }
  err = division.Assert(*fields.Division, false);             if err != nil { return nil, err }
  err = gameID.Assert(*fields.GameID, false);                 if err != nil { return nil, err }
  err = gameSegment.Assert(*fields.GameSegment, false);       if err != nil { return nil, err }
  err = groupQuantity.Assert(*fields.GroupQuantity, false);   if err != nil { return nil, err }
  err = lastNGames.Assert(*fields.LastNGames, false);         if err != nil { return nil, err }
  err = leagueID.Assert(*fields.LeagueID, false);             if err != nil { return nil, err }
  err = location.Assert(*fields.Location, false);             if err != nil { return nil, err }
  err = measureType.Assert(*fields.MeasureType, false);       if err != nil { return nil, err }
  err = month.Assert(*fields.Month, false);                   if err != nil { return nil, err }
  err = opponentTeamID.Assert(*fields.OpponentTeamID, false); if err != nil { return nil, err }
  err = outcome.Assert(*fields.Outcome, false);               if err != nil { return nil, err }
  err = pORound.Assert(*fields.PORound, false);               if err != nil { return nil, err }
  err = paceAdjust.Assert(*fields.PaceAdjust, false);         if err != nil { return nil, err }
  err = perMode.Assert(*fields.PerMode, false);               if err != nil { return nil, err }
  err = period.Assert(*fields.Period, false);                 if err != nil { return nil, err }
  err = plusMinus.Assert(*fields.PlusMinus, false);           if err != nil { return nil, err }
  err = rank.Assert(*fields.Rank, false);                     if err != nil { return nil, err }
  err = season.Assert(*fields.Season, false);                 if err != nil { return nil, err }
  err = seasonSegment.Assert(*fields.SeasonSegment, false);   if err != nil { return nil, err }
  err = seasonType.Assert(*fields.SeasonType, false);         if err != nil { return nil, err }
  err = shotClockRange.Assert(*fields.ShotClockRange, false); if err != nil { return nil, err }
  err = teamID.Assert(*fields.TeamID, false);                 if err != nil { return nil, err }
  err = vsConference.Assert(*fields.VsConference, false);     if err != nil { return nil, err }
  err = vsDivision.Assert(*fields.VsDivision, false);         if err != nil { return nil, err }

  bytes, err := c.Get(model.FetchConfig{
    DataSource: "stats.nba.com",
    Endpoint:   "/leaguedashlineups",
    Fields:     &map[string]string{
      "Conference":     conference.FromPtr(fields.Conference),
      "DateFrom":       dateFrom.FromPtr(fields.DateFrom),
      "DateTo":         dateTo.FromPtr(fields.DateTo),
      "Division":       division.FromPtr(fields.Division),
      "GameID":         gameID.FromPtr(fields.GameID),
      "GameSegment":    gameSegment.FromPtr(fields.GameSegment),
      "GroupQuantity":  groupQuantity.FromPtr(fields.GroupQuantity),
      "LastNGames":     lastNGames.FromPtr(fields.LastNGames),
      "LeagueID":       leagueID.FromPtr(fields.LeagueID),
      "Location":       location.FromPtr(fields.Location),
      "MeasureType":    measureType.FromPtr(fields.MeasureType),
      "Month":          month.FromPtr(fields.Month),
      "OpponentTeamID": opponentTeamID.FromPtr(fields.OpponentTeamID),
      "Outcome":        outcome.FromPtr(fields.Outcome),
      "PORound":        pORound.FromPtr(fields.PORound),
      "PaceAdjust":     paceAdjust.FromPtr(fields.PaceAdjust),
      "PerMode":        perMode.FromPtr(fields.PerMode),
      "Period":         period.FromPtr(fields.Period),
      "PlusMinus":      plusMinus.FromPtr(fields.PlusMinus),
      "Rank":           rank.FromPtr(fields.Rank),
      "Season":         season.FromPtr(fields.Season),
      "SeasonSegment":  seasonSegment.FromPtr(fields.SeasonSegment),
      "SeasonType":     seasonType.FromPtr(fields.SeasonType),
      "ShotClockRange": shotClockRange.FromPtr(fields.ShotClockRange),
      "TeamID":         teamID.FromPtr(fields.TeamID),
      "VsConference":   vsConference.FromPtr(fields.VsConference),
      "VsDivision":     vsDivision.FromPtr(fields.VsDivision),
    },
  })
  if err != nil {
    return nil, err
  }

  var body nbastatsResponseBody
  return body.UnmarshalResult(bytes)
}
