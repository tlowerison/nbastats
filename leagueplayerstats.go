package nbastats

import "github.com/tlowerison/nbastats/model"

type LeaguePlayerStatsFields struct {
  // Required
  PlayerID   string
  // Optional
  College          *string
  Conference       *string
  Country          *string
  DateFrom         *string
  DateTo           *string
  Division         *string
  DraftPick        *string
  DraftYear        *int
  GameSegment      *string
  Height           *string
  LastNGames       *int
  LeagueID         *string
  Location         *string
  MeasureType      *string
  Month            *int
  OpponentTeamID   *string
  Outcome          *string
  PORound          *int
  PaceAdjust       *string
  PerMode          *string
  Period           *int
  PlayerExperience *string
  PlayerPosition   *string
  PlusMinus        *string
  Rank             *string
  Season           *string
  SeasonSegment    *string
  SeasonType       *string
  ShotClockRange   *string
  StarterBench     *string
  TeamID           *string
  TwoWay           *int
  VsConference     *string
  VsDivision       *string
  Weight           *string
}

func (c *NbaStatsClient) LeaguePlayerStats(fields *LeaguePlayerStatsFields) (*model.Result, error) {
  if fields.College == nil          { fields.College          = &College.Default }
  if fields.Conference == nil       { fields.Conference       = &Conference.Default }
  if fields.Country == nil          { fields.Country          = &Country.Default }
  if fields.DateFrom == nil         { fields.DateFrom         = &DateFrom.Default }
  if fields.DateTo == nil           { fields.DateTo           = &DateTo.Default }
  if fields.Division == nil         { fields.Division         = &Division.Default }
  if fields.DraftPick == nil        { fields.DraftPick        = &DraftPick.Default }
  if fields.DraftYear == nil        { fields.DraftYear        = &DraftYear.Default }
  if fields.GameSegment == nil      { fields.GameSegment      = &GameSegment.Default }
  if fields.Height == nil           { fields.Height           = &Height.Default }
  if fields.LastNGames == nil       { fields.LastNGames       = &LastNGames.Default }
  if fields.LeagueID == nil         { fields.LeagueID         = &LeagueID.Default }
  if fields.Location == nil         { fields.Location         = &Location.Default }
  if fields.MeasureType == nil      { fields.MeasureType      = &MeasureType.Default }
  if fields.Month == nil            { fields.Month            = &Month.Default }
  if fields.OpponentTeamID == nil   { fields.OpponentTeamID   = &OpponentTeamID.Default }
  if fields.Outcome == nil          { fields.Outcome          = &Outcome.Default }
  if fields.PORound == nil          { fields.PORound          = &PORound.Default }
  if fields.PaceAdjust == nil       { fields.PaceAdjust       = &PaceAdjust.Default }
  if fields.PerMode == nil          { fields.PerMode          = &PerMode.Default }
  if fields.Period == nil           { fields.Period           = &Period.Default }
  if fields.PlayerExperience == nil { fields.PlayerExperience = &PlayerExperience.Default }
  if fields.PlayerPosition == nil   { fields.PlayerPosition   = &PlayerPosition.Default }
  if fields.PlusMinus == nil        { fields.PlusMinus        = &PlusMinus.Default }
  if fields.Rank == nil             { fields.Rank             = &Rank.Default }
  if fields.Season == nil           { fields.Season           = &Season.Default }
  if fields.SeasonSegment == nil    { fields.SeasonSegment    = &SeasonSegment.Default }
  if fields.SeasonType == nil       { fields.SeasonType       = &SeasonType.Default }
  if fields.ShotClockRange == nil   { fields.ShotClockRange   = &ShotClockRange.Default }
  if fields.StarterBench == nil     { fields.StarterBench     = &StarterBench.Default }
  if fields.TeamID == nil           { fields.TeamID           = &TeamID.Default }
  if fields.TwoWay == nil           { fields.TwoWay           = &TwoWay.Default }
  if fields.VsConference == nil     { fields.VsConference     = &VsConference.Default }
  if fields.VsDivision == nil       { fields.VsDivision       = &VsDivision.Default }
  if fields.Weight == nil           { fields.Weight           = &Weight.Default }

  // Validate
  err := PlayerID.Assert(fields.PlayerID);                 if err != nil { return nil, err }
  err = College.Assert(*fields.College);                   if err != nil { return nil, err }
  err = Conference.Assert(*fields.Conference);             if err != nil { return nil, err }
  err = Country.Assert(*fields.Country);                   if err != nil { return nil, err }
  err = DateFrom.Assert(*fields.DateFrom);                 if err != nil { return nil, err }
  err = DateTo.Assert(*fields.DateTo);                     if err != nil { return nil, err }
  err = Division.Assert(*fields.Division);                 if err != nil { return nil, err }
  err = DraftPick.Assert(*fields.DraftPick);               if err != nil { return nil, err }
  err = DraftYear.Assert(*fields.DraftYear);               if err != nil { return nil, err }
  err = GameSegment.Assert(*fields.GameSegment);           if err != nil { return nil, err }
  err = Height.Assert(*fields.Height);                     if err != nil { return nil, err }
  err = LastNGames.Assert(*fields.LastNGames);             if err != nil { return nil, err }
  err = LeagueID.Assert(*fields.LeagueID);                 if err != nil { return nil, err }
  err = Location.Assert(*fields.Location);                 if err != nil { return nil, err }
  err = MeasureType.Assert(*fields.MeasureType);           if err != nil { return nil, err }
  err = Month.Assert(*fields.Month);                       if err != nil { return nil, err }
  err = OpponentTeamID.Assert(*fields.OpponentTeamID);     if err != nil { return nil, err }
  err = Outcome.Assert(*fields.Outcome);                   if err != nil { return nil, err }
  err = PORound.Assert(*fields.PORound);                   if err != nil { return nil, err }
  err = PaceAdjust.Assert(*fields.PaceAdjust);             if err != nil { return nil, err }
  err = PerMode.Assert(*fields.PerMode);                   if err != nil { return nil, err }
  err = Period.Assert(*fields.Period);                     if err != nil { return nil, err }
  err = PlayerExperience.Assert(*fields.PlayerExperience); if err != nil { return nil, err }
  err = PlayerPosition.Assert(*fields.PlayerPosition);     if err != nil { return nil, err }
  err = PlusMinus.Assert(*fields.PlusMinus);               if err != nil { return nil, err }
  err = Rank.Assert(*fields.Rank);                         if err != nil { return nil, err }
  err = Season.Assert(*fields.Season);                     if err != nil { return nil, err }
  err = SeasonSegment.Assert(*fields.SeasonSegment);       if err != nil { return nil, err }
  err = SeasonType.Assert(*fields.SeasonType);             if err != nil { return nil, err }
  err = ShotClockRange.Assert(*fields.ShotClockRange);     if err != nil { return nil, err }
  err = StarterBench.Assert(*fields.StarterBench);         if err != nil { return nil, err }
  err = TeamID.Assert(*fields.TeamID);                     if err != nil { return nil, err }
  err = TwoWay.Assert(*fields.TwoWay);                     if err != nil { return nil, err }
  err = VsConference.Assert(*fields.VsConference);         if err != nil { return nil, err }
  err = VsDivision.Assert(*fields.VsDivision);             if err != nil { return nil, err }
  err = Weight.Assert(*fields.Weight);                     if err != nil { return nil, err }

  bytes, err := c.Get(model.FetchConfig{
    DataSource: "stats.nba.com",
    Endpoint:   "/leaguedashplayerstats",
    Fields:     &map[string]string{
      "College":          College.FromPtr(fields.College),
      "Conference":       Conference.FromPtr(fields.Conference),
      "Country":          Country.FromPtr(fields.Country),
      "DateFrom":         DateFrom.FromPtr(fields.DateFrom),
      "DateTo":           DateTo.FromPtr(fields.DateTo),
      "Division":         Division.FromPtr(fields.Division),
      "DraftPick":        DraftPick.FromPtr(fields.DraftPick),
      "DraftYear":        DraftYear.FromPtrAsStr(fields.DraftYear),
      "GameScope":        "",
      "GameSegment":      GameSegment.FromPtr(fields.GameSegment),
      "Height":           Height.FromPtr(fields.Height),
      "LastNGames":       LastNGames.FromPtr(fields.LastNGames),
      "LeagueID":         LeagueID.FromPtr(fields.LeagueID),
      "Location":         Location.FromPtr(fields.Location),
      "MeasureType":      MeasureType.FromPtr(fields.MeasureType),
      "Month":            Month.FromPtr(fields.Month),
      "OpponentTeamID":   OpponentTeamID.FromPtr(fields.OpponentTeamID),
      "Outcome":          Outcome.FromPtr(fields.Outcome),
      "PORound":          PORound.FromPtr(fields.PORound),
      "PaceAdjust":       PaceAdjust.FromPtr(fields.PaceAdjust),
      "PerMode":          PerMode.FromPtr(fields.PerMode),
      "Period":           Period.FromPtr(fields.Period),
      "PlayerExperience": PlayerExperience.FromPtr(fields.PlayerExperience),
      "PlayerPosition":   PlayerPosition.FromPtr(fields.PlayerPosition),
      "PlayerID":         PlayerID.FromVal(fields.PlayerID),
      "PlusMinus":        PlusMinus.FromPtr(fields.PlusMinus),
      "Rank":             Rank.FromPtr(fields.Rank),
      "Season":           Season.FromPtr(fields.Season),
      "SeasonSegment":    SeasonSegment.FromPtr(fields.SeasonSegment),
      "SeasonType":       SeasonType.FromPtr(fields.SeasonType),
      "ShotClockRange":   ShotClockRange.FromPtr(fields.ShotClockRange),
      "StarterBench":     StarterBench.FromPtr(fields.StarterBench),
      "TeamID":           TeamID.FromPtrAsInt(fields.TeamID),
      "TwoWay":           TwoWay.FromPtr(fields.TwoWay),
      "VsConference":     VsConference.FromPtr(fields.VsConference),
      "VsDivision":       VsDivision.FromPtr(fields.VsDivision),
      "Weight":           Weight.FromPtr(fields.Weight),
    },
  })
  if err != nil {
    return nil, err
  }

  var body nbastatsResponseBody
  return body.UnmarshalResult(bytes)
}
