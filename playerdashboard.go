package nbastats

import "github.com/tlowerison/nbastats/model"

type PlayerDashboardFields struct {
  // Required
  PlayerID   string
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
  Split          *string
  VsConference   *string
  VsDivision     *string
}

func (c *NbaStatsClient) PlayerDashboard(fields *PlayerDashboardFields) (*model.Result, error) {
  if fields.DateFrom == nil       { fields.DateFrom       = &DateFrom.Default }
  if fields.DateTo == nil         { fields.DateTo         = &DateTo.Default }
  if fields.GameSegment == nil    { fields.GameSegment    = &GameSegment.Default }
  if fields.LastNGames == nil     { fields.LastNGames     = &LastNGames.Default }
  if fields.LeagueID == nil       { fields.LeagueID       = &LeagueID.Default }
  if fields.Location == nil       { fields.Location       = &Location.Default }
  if fields.MeasureType == nil    { fields.MeasureType    = &MeasureType.Default }
  if fields.Month == nil          { fields.Month          = &Month.Default }
  if fields.OpponentTeamID == nil { fields.OpponentTeamID = &OpponentTeamID.Default }
  if fields.Outcome == nil        { fields.Outcome        = &Outcome.Default }
  if fields.PORound == nil        { fields.PORound        = &PORound.Default }
  if fields.PaceAdjust == nil     { fields.PaceAdjust     = &PaceAdjust.Default }
  if fields.PerMode == nil        { fields.PerMode        = &PerMode.Default }
  if fields.Period == nil         { fields.Period         = &Period.Default }
  if fields.PlusMinus == nil      { fields.PlusMinus      = &PlusMinus.Default }
  if fields.Rank == nil           { fields.Rank           = &Rank.Default }
  if fields.Season == nil         { fields.Season         = &Season.Default }
  if fields.SeasonSegment == nil  { fields.SeasonSegment  = &SeasonSegment.Default }
  if fields.SeasonType == nil     { fields.SeasonType     = &SeasonType.Default }
  if fields.ShotClockRange == nil { fields.ShotClockRange = &ShotClockRange.Default }
  if fields.Split == nil          { fields.Split          = &Split.Default }
  if fields.VsConference == nil   { fields.VsConference   = &VsConference.Default }
  if fields.VsDivision == nil     { fields.VsDivision     = &VsDivision.Default }

  // Validate
  err := PlayerID.Assert(fields.PlayerID);             if err != nil { return nil, err }
  err = DateFrom.Assert(*fields.DateFrom);             if err != nil { return nil, err }
  err = DateTo.Assert(*fields.DateTo);                 if err != nil { return nil, err }
  err = GameSegment.Assert(*fields.GameSegment);       if err != nil { return nil, err }
  err = LastNGames.Assert(*fields.LastNGames);         if err != nil { return nil, err }
  err = LeagueID.Assert(*fields.LeagueID);             if err != nil { return nil, err }
  err = Location.Assert(*fields.Location);             if err != nil { return nil, err }
  err = MeasureType.Assert(*fields.MeasureType);       if err != nil { return nil, err }
  err = Month.Assert(*fields.Month);                   if err != nil { return nil, err }
  err = OpponentTeamID.Assert(*fields.OpponentTeamID); if err != nil { return nil, err }
  err = Outcome.Assert(*fields.Outcome);               if err != nil { return nil, err }
  err = PORound.Assert(*fields.PORound);               if err != nil { return nil, err }
  err = PaceAdjust.Assert(*fields.PaceAdjust);         if err != nil { return nil, err }
  err = PerMode.Assert(*fields.PerMode);               if err != nil { return nil, err }
  err = Period.Assert(*fields.Period);                 if err != nil { return nil, err }
  err = PlusMinus.Assert(*fields.PlusMinus);           if err != nil { return nil, err }
  err = Rank.Assert(*fields.Rank);                     if err != nil { return nil, err }
  err = Season.Assert(*fields.Season);                 if err != nil { return nil, err }
  err = SeasonSegment.Assert(*fields.SeasonSegment);   if err != nil { return nil, err }
  err = SeasonType.Assert(*fields.SeasonType);         if err != nil { return nil, err }
  err = ShotClockRange.Assert(*fields.ShotClockRange); if err != nil { return nil, err }
  err = Split.Assert(*fields.Split);                   if err != nil { return nil, err }
  err = VsConference.Assert(*fields.VsConference);     if err != nil { return nil, err }
  err = VsDivision.Assert(*fields.VsDivision);         if err != nil { return nil, err }

  bytes, err := c.Get(model.FetchConfig{
    DataSource: "stats.nba.com",
    Endpoint:   "/playerdashboardbygeneralsplits",
    Fields:     &map[string]string{
      "DateFrom":       DateFrom.FromPtr(fields.DateFrom),
      "DateTo":         DateTo.FromPtr(fields.DateTo),
      "GameSegment":    GameSegment.FromPtr(fields.GameSegment),
      "LastNGames":     LastNGames.FromPtr(fields.LastNGames),
      "LeagueID":       LeagueID.FromPtr(fields.LeagueID),
      "Location":       Location.FromPtr(fields.Location),
      "MeasureType":    MeasureType.FromPtr(fields.MeasureType),
      "Month":          Month.FromPtr(fields.Month),
      "OpponentTeamID": OpponentTeamID.FromPtr(fields.OpponentTeamID),
      "Outcome":        Outcome.FromPtr(fields.Outcome),
      "PORound":        PORound.FromPtr(fields.PORound),
      "PaceAdjust":     PaceAdjust.FromPtr(fields.PaceAdjust),
      "PerMode":        PerMode.FromPtr(fields.PerMode),
      "Period":         Period.FromPtr(fields.Period),
      "PlayerID":       PlayerID.FromVal(fields.PlayerID),
      "PlusMinus":      PlusMinus.FromPtr(fields.PlusMinus),
      "Rank":           Rank.FromPtr(fields.Rank),
      "Season":         Season.FromPtr(fields.Season),
      "SeasonSegment":  SeasonSegment.FromPtr(fields.SeasonSegment),
      "SeasonType":     SeasonType.FromPtr(fields.SeasonType),
      "ShotClockRange": ShotClockRange.FromPtr(fields.ShotClockRange),
      "Split":          Split.FromPtr(fields.Split),
      "VsConference":   VsConference.FromPtr(fields.VsConference),
      "VsDivision":     VsDivision.FromPtr(fields.VsDivision),
    },
  })
  if err != nil {
    return nil, err
  }

  var body nbastatsResponseBody
  return body.UnmarshalResult(bytes)
}
