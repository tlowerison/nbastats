package nbastats

import "github.com/tlowerison/nbastats/model"

type TeamDashboardFields struct {
  // Required
  TeamID         string
  // Optional
  DateFrom       *string
  DateTo         *string
  GameSegment    *string
  LastNGames     *int
  LeagueID       *string
  Location       *string
  MeasureType    *string
  Month          *int
  OpponentTeamID *string
  Outcome        *string
  PaceAdjust     *string
  Period         *int
  PerMode        *string
  PlusMinus      *string
  PORound        *int
  Rank           *string
  Season         *string
  SeasonSegment  *string
  SeasonType     *string
  ShotClockRange *string
  Split          *string
  VsConference   *string
  VsDivision     *string
}

func (c *NbaStatsClient) TeamDashboard(fields *TeamDashboardFields) (*model.Result, error) {
  if fields.DateFrom == nil       { fields.DateFrom       = &dateFrom.Default }
  if fields.DateTo == nil         { fields.DateTo         = &dateTo.Default }
  if fields.GameSegment == nil    { fields.GameSegment    = &gameSegment.Default }
  if fields.LastNGames == nil     { fields.LastNGames     = &lastNGames.Default }
  if fields.LeagueID == nil       { fields.LeagueID       = &leagueID.Default }
  if fields.Location == nil       { fields.Location       = &location.Default }
  if fields.MeasureType == nil    { fields.MeasureType    = &measureType.Default }
  if fields.Month == nil          { fields.Month          = &month.Default }
  if fields.OpponentTeamID == nil { fields.OpponentTeamID = &opponentTeamID.Default }
  if fields.Outcome == nil        { fields.Outcome        = &outcome.Default }
  if fields.PaceAdjust == nil     { fields.PaceAdjust     = &paceAdjust.Default }
  if fields.Period == nil         { fields.Period         = &period.Default }
  if fields.PerMode == nil        { fields.PerMode        = &perMode.Default }
  if fields.PlusMinus == nil      { fields.PlusMinus      = &plusMinus.Default }
  if fields.PORound == nil        { fields.PORound        = &pORound.Default }
  if fields.Rank == nil           { fields.Rank           = &rank.Default }
  if fields.Season == nil         { fields.Season         = &season.Default }
  if fields.SeasonSegment == nil  { fields.SeasonSegment  = &seasonSegment.Default }
  if fields.SeasonType == nil     { fields.SeasonType     = &seasonType.Default }
  if fields.ShotClockRange == nil { fields.ShotClockRange = &shotClockRange.Default }
  if fields.Split == nil          { fields.Split          = &split.Default }
  if fields.VsConference == nil   { fields.VsConference   = &vsConference.Default }
  if fields.VsDivision == nil     { fields.VsDivision     = &vsDivision.Default }

  // Validate
  err := teamID.Assert(fields.TeamID, true);                  if err != nil { return nil, err }
  err = dateFrom.Assert(*fields.DateFrom, false);             if err != nil { return nil, err }
  err = dateTo.Assert(*fields.DateTo, false);                 if err != nil { return nil, err }
  err = gameSegment.Assert(*fields.GameSegment, false);       if err != nil { return nil, err }
  err = lastNGames.Assert(*fields.LastNGames, false);         if err != nil { return nil, err }
  err = leagueID.Assert(*fields.LeagueID, false);             if err != nil { return nil, err }
  err = location.Assert(*fields.Location, false);             if err != nil { return nil, err }
  err = measureType.Assert(*fields.MeasureType, false);       if err != nil { return nil, err }
  err = month.Assert(*fields.Month, false);                   if err != nil { return nil, err }
  err = opponentTeamID.Assert(*fields.OpponentTeamID, false); if err != nil { return nil, err }
  err = outcome.Assert(*fields.Outcome, false);               if err != nil { return nil, err }
  err = paceAdjust.Assert(*fields.PaceAdjust, false);         if err != nil { return nil, err }
  err = period.Assert(*fields.Period, false);                 if err != nil { return nil, err }
  err = perMode.Assert(*fields.PerMode, false);               if err != nil { return nil, err }
  err = plusMinus.Assert(*fields.PlusMinus, false);           if err != nil { return nil, err }
  err = pORound.Assert(*fields.PORound, false);               if err != nil { return nil, err }
  err = rank.Assert(*fields.Rank, false);                     if err != nil { return nil, err }
  err = season.Assert(*fields.Season, false);                 if err != nil { return nil, err }
  err = seasonSegment.Assert(*fields.SeasonSegment, false);   if err != nil { return nil, err }
  err = seasonType.Assert(*fields.SeasonType, false);         if err != nil { return nil, err }
  err = shotClockRange.Assert(*fields.ShotClockRange, false); if err != nil { return nil, err }
  err = split.Assert(*fields.Split, false);                   if err != nil { return nil, err }
  err = vsConference.Assert(*fields.VsConference, false);     if err != nil { return nil, err }
  err = vsDivision.Assert(*fields.VsDivision, false);         if err != nil { return nil, err }

  bytes, err := c.Get(model.FetchConfig{
    DataSource: "stats.nba.com",
    Endpoint:   "/teamdashboardbygeneralsplits",
    Fields:     &map[string]string{
      "DateFrom":       dateFrom.FromPtr(fields.DateFrom),
      "DateTo":         dateTo.FromPtr(fields.DateTo),
      "GameSegment":    gameSegment.FromPtr(fields.GameSegment),
      "LastNGames":     lastNGames.FromPtr(fields.LastNGames),
      "LeagueID":       leagueID.FromPtr(fields.LeagueID),
      "Location":       location.FromPtr(fields.Location),
      "MeasureType":    measureType.FromPtr(fields.MeasureType),
      "Month":          month.FromPtr(fields.Month),
      "OpponentTeamID": opponentTeamID.FromPtr(fields.OpponentTeamID),
      "Outcome":        outcome.FromPtr(fields.Outcome),
      "PaceAdjust":     paceAdjust.FromPtr(fields.PaceAdjust),
      "Period":         period.FromPtr(fields.Period),
      "PerMode":        perMode.FromPtr(fields.PerMode),
      "PlusMinus":      plusMinus.FromPtr(fields.PlusMinus),
      "PORound":        pORound.FromPtr(fields.PORound),
      "Rank":           rank.FromPtr(fields.Rank),
      "Season":         season.FromPtr(fields.Season),
      "SeasonSegment":  seasonSegment.FromPtr(fields.SeasonSegment),
      "SeasonType":     seasonType.FromPtr(fields.SeasonType),
      "ShotClockRange": shotClockRange.FromPtr(fields.ShotClockRange),
      "Split":          split.FromPtr(fields.Split),
      "TeamID":         teamID.FromValAsInt(fields.TeamID),
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
