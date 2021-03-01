# nbastats
A go HTTP client for fetching data from various NBA related data sources. Results are automatically formatted in .tsv format.

## Usage
```go
package main

import (
  "fmt"

  "github.com/tlowerison/nbastats"
)

func main() {
  client := nbastats.NewClient()
  data, err := client.TeamNews(&nbastats.TeamNewsFields{TeamID: "1610612763"})
  if err != nil {
    fmt.Println(err.Error())
  } else {
    fmt.Println(data)
  }
}
```

## Endpoints
### CommonAllPlayers
Basic data about all players in a league.

### CommonPlayerInfo
Basic data about a specific player.

### LeagueGameLog
Game logs for a specific league.

### LeaguePlayerStats
Player stats for a specific player.

### Lineups
Lineup data split by a variety of fields.

### PlayByPlay
Event data for a specific game.

### PlayerDashboard
Detailed data for a specific player.

### PlayerGameLog
Game logs for a specific player.

### Scoreboard
Live scores of ongoing games.

### TeamDashboard
Detailed data for a specific team.

### TeamGameLog
Game logs for a specific team.

### TeamNews
Recent news about players on a specific team.
