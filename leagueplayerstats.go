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
  if fields.College == nil          { fields.College          = &college.Default }
  if fields.Conference == nil       { fields.Conference       = &conference.Default }
  if fields.Country == nil          { fields.Country          = &country.Default }
  if fields.DateFrom == nil         { fields.DateFrom         = &dateFrom.Default }
  if fields.DateTo == nil           { fields.DateTo           = &dateTo.Default }
  if fields.Division == nil         { fields.Division         = &division.Default }
  if fields.DraftPick == nil        { fields.DraftPick        = &draftPick.Default }
  if fields.DraftYear == nil        { fields.DraftYear        = &draftYear.Default }
  if fields.GameSegment == nil      { fields.GameSegment      = &gameSegment.Default }
  if fields.Height == nil           { fields.Height           = &height.Default }
  if fields.LastNGames == nil       { fields.LastNGames       = &lastNGames.Default }
  if fields.LeagueID == nil         { fields.LeagueID         = &leagueID.Default }
  if fields.Location == nil         { fields.Location         = &location.Default }
  if fields.MeasureType == nil      { fields.MeasureType      = &measureType.Default }
  if fields.Month == nil            { fields.Month            = &month.Default }
  if fields.OpponentTeamID == nil   { fields.OpponentTeamID   = &opponentTeamID.Default }
  if fields.Outcome == nil          { fields.Outcome          = &outcome.Default }
  if fields.PORound == nil          { fields.PORound          = &pORound.Default }
  if fields.PaceAdjust == nil       { fields.PaceAdjust       = &paceAdjust.Default }
  if fields.PerMode == nil          { fields.PerMode          = &perMode.Default }
  if fields.Period == nil           { fields.Period           = &period.Default }
  if fields.PlayerExperience == nil { fields.PlayerExperience = &playerExperience.Default }
  if fields.PlayerPosition == nil   { fields.PlayerPosition   = &playerPosition.Default }
  if fields.PlusMinus == nil        { fields.PlusMinus        = &plusMinus.Default }
  if fields.Rank == nil             { fields.Rank             = &rank.Default }
  if fields.Season == nil           { fields.Season           = &season.Default }
  if fields.SeasonSegment == nil    { fields.SeasonSegment    = &seasonSegment.Default }
  if fields.SeasonType == nil       { fields.SeasonType       = &seasonType.Default }
  if fields.ShotClockRange == nil   { fields.ShotClockRange   = &shotClockRange.Default }
  if fields.StarterBench == nil     { fields.StarterBench     = &starterBench.Default }
  if fields.TeamID == nil           { fields.TeamID           = &teamID.Default }
  if fields.TwoWay == nil           { fields.TwoWay           = &twoWay.Default }
  if fields.VsConference == nil     { fields.VsConference     = &vsConference.Default }
  if fields.VsDivision == nil       { fields.VsDivision       = &vsDivision.Default }
  if fields.Weight == nil           { fields.Weight           = &weight.Default }

  // Validate
  err := playerID.Assert(fields.PlayerID);                 if err != nil { return nil, err }
  err = college.Assert(*fields.College);                   if err != nil { return nil, err }
  err = conference.Assert(*fields.Conference);             if err != nil { return nil, err }
  err = country.Assert(*fields.Country);                   if err != nil { return nil, err }
  err = dateFrom.Assert(*fields.DateFrom);                 if err != nil { return nil, err }
  err = dateTo.Assert(*fields.DateTo);                     if err != nil { return nil, err }
  err = division.Assert(*fields.Division);                 if err != nil { return nil, err }
  err = draftPick.Assert(*fields.DraftPick);               if err != nil { return nil, err }
  err = draftYear.Assert(*fields.DraftYear);               if err != nil { return nil, err }
  err = gameSegment.Assert(*fields.GameSegment);           if err != nil { return nil, err }
  err = height.Assert(*fields.Height);                     if err != nil { return nil, err }
  err = lastNGames.Assert(*fields.LastNGames);             if err != nil { return nil, err }
  err = leagueID.Assert(*fields.LeagueID);                 if err != nil { return nil, err }
  err = location.Assert(*fields.Location);                 if err != nil { return nil, err }
  err = measureType.Assert(*fields.MeasureType);           if err != nil { return nil, err }
  err = month.Assert(*fields.Month);                       if err != nil { return nil, err }
  err = opponentTeamID.Assert(*fields.OpponentTeamID);     if err != nil { return nil, err }
  err = outcome.Assert(*fields.Outcome);                   if err != nil { return nil, err }
  err = pORound.Assert(*fields.PORound);                   if err != nil { return nil, err }
  err = paceAdjust.Assert(*fields.PaceAdjust);             if err != nil { return nil, err }
  err = perMode.Assert(*fields.PerMode);                   if err != nil { return nil, err }
  err = period.Assert(*fields.Period);                     if err != nil { return nil, err }
  err = playerExperience.Assert(*fields.PlayerExperience); if err != nil { return nil, err }
  err = playerPosition.Assert(*fields.PlayerPosition);     if err != nil { return nil, err }
  err = plusMinus.Assert(*fields.PlusMinus);               if err != nil { return nil, err }
  err = rank.Assert(*fields.Rank);                         if err != nil { return nil, err }
  err = season.Assert(*fields.Season);                     if err != nil { return nil, err }
  err = seasonSegment.Assert(*fields.SeasonSegment);       if err != nil { return nil, err }
  err = seasonType.Assert(*fields.SeasonType);             if err != nil { return nil, err }
  err = shotClockRange.Assert(*fields.ShotClockRange);     if err != nil { return nil, err }
  err = starterBench.Assert(*fields.StarterBench);         if err != nil { return nil, err }
  err = teamID.Assert(*fields.TeamID);                     if err != nil { return nil, err }
  err = twoWay.Assert(*fields.TwoWay);                     if err != nil { return nil, err }
  err = vsConference.Assert(*fields.VsConference);         if err != nil { return nil, err }
  err = vsDivision.Assert(*fields.VsDivision);             if err != nil { return nil, err }
  err = weight.Assert(*fields.Weight);                     if err != nil { return nil, err }

  bytes, err := c.Get(model.FetchConfig{
    DataSource: "stats.nba.com",
    Endpoint:   "/leaguedashplayerstats",
    Fields:     &map[string]string{
      "College":          college.FromPtr(fields.College),
      "Conference":       conference.FromPtr(fields.Conference),
      "Country":          country.FromPtr(fields.Country),
      "DateFrom":         dateFrom.FromPtr(fields.DateFrom),
      "DateTo":           dateTo.FromPtr(fields.DateTo),
      "Division":         division.FromPtr(fields.Division),
      "DraftPick":        draftPick.FromPtr(fields.DraftPick),
      "DraftYear":        draftYear.FromPtrAsStr(fields.DraftYear),
      "GameScope":        "",
      "GameSegment":      gameSegment.FromPtr(fields.GameSegment),
      "Height":           height.FromPtr(fields.Height),
      "LastNGames":       lastNGames.FromPtr(fields.LastNGames),
      "LeagueID":         leagueID.FromPtr(fields.LeagueID),
      "Location":         location.FromPtr(fields.Location),
      "MeasureType":      measureType.FromPtr(fields.MeasureType),
      "Month":            month.FromPtr(fields.Month),
      "OpponentTeamID":   opponentTeamID.FromPtr(fields.OpponentTeamID),
      "Outcome":          outcome.FromPtr(fields.Outcome),
      "PORound":          pORound.FromPtr(fields.PORound),
      "PaceAdjust":       paceAdjust.FromPtr(fields.PaceAdjust),
      "PerMode":          perMode.FromPtr(fields.PerMode),
      "Period":           period.FromPtr(fields.Period),
      "PlayerExperience": playerExperience.FromPtr(fields.PlayerExperience),
      "PlayerPosition":   playerPosition.FromPtr(fields.PlayerPosition),
      "PlayerID":         playerID.FromVal(fields.PlayerID),
      "PlusMinus":        plusMinus.FromPtr(fields.PlusMinus),
      "Rank":             rank.FromPtr(fields.Rank),
      "Season":           season.FromPtr(fields.Season),
      "SeasonSegment":    seasonSegment.FromPtr(fields.SeasonSegment),
      "SeasonType":       seasonType.FromPtr(fields.SeasonType),
      "ShotClockRange":   shotClockRange.FromPtr(fields.ShotClockRange),
      "StarterBench":     starterBench.FromPtr(fields.StarterBench),
      "TeamID":           teamID.FromPtrAsInt(fields.TeamID),
      "TwoWay":           twoWay.FromPtr(fields.TwoWay),
      "VsConference":     vsConference.FromPtr(fields.VsConference),
      "VsDivision":       vsDivision.FromPtr(fields.VsDivision),
      "Weight":           weight.FromPtr(fields.Weight),
    },
  })
  if err != nil {
    return nil, err
  }

  var body nbastatsResponseBody
  return body.UnmarshalResult(bytes)
}
