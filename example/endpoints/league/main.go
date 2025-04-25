package main

import (
	"context"
	"fmt"
	"github.com/pmurley/go-mlb/pkg/client"
	"github.com/pmurley/go-mlb/pkg/gen"
	"time"
)

/**
OUTPUT:
------------------------------------------------------------------------
American League: Season runs 2025-02-20 - 2025-10-31
National League: Season runs 2025-02-20 - 2025-10-31
Loaded 251 write-in players
Loaded 17 final vote players. Here are 5 of them for 2024.
1: Aaron Judge
2: Vladimir Guerrero
3: Ryan Mountcastle
4: Jose Altuve
5: Marcus Semien
*/

func main() {
	client, err := client.NewClient("https://statsapi.mlb.com")
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// GetLeagues
	params := &gen.LeagueParams{
		LeagueIds: new([]int32),
		Season:    new(string),
		SportId:   new(int32),
	}
	*params.SportId = 1
	*params.Season = "2025"

	// American and National League respectively
	*params.LeagueIds = []int32{103, 104}

	leagues, err := client.GetLeagues(ctx, params)
	if err != nil {
		panic(err)
	}

	for _, league := range leagues {
		fmt.Printf("%s: Season runs %s - %s\n",
			league.Name, league.SeasonDateInfo.SeasonStartDate, league.SeasonDateInfo.SeasonEndDate)
	}

	// GetAllStarWriteIns
	writeInParams := &gen.AllStarWriteInsParams{
		Season: new(string),
	}
	*writeInParams.Season = "2024"
	writeIns, err := client.GetAllStarWriteIns(ctx, 103, writeInParams)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Loaded %d write-in players\n", len(writeIns))

	// GetAllStarFinalVote
	fvParams := &gen.AllStarFinalVoteParams{
		Season: new(string),
	}
	*fvParams.Season = "2024"

	finalVote, err := client.GetAllStarFinalVote(ctx, 103, fvParams)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Loaded %d final vote players. Here are 5 of them for %s.\n",
		len(finalVote), *fvParams.Season)
	for i, player := range finalVote {
		if i > 4 {
			break
		}
		fmt.Printf("%d: %s %s\n", i+1, player.FirstName, player.LastName)
	}
}
