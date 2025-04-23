package main

import (
	"context"
	"fmt"
	"github.com/pmurley/go-mlb/pkg/client"
	"github.com/pmurley/go-mlb/pkg/gen"
	log "github.com/sirupsen/logrus"
	"time"
)

func main() {
	client, err := client.NewClient("https://statsapi.mlb.com")
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Println("call func")
	var params gen.ScheduleParams
	params.SportId = new([]int32)
	*params.SportId = []int32{1}
	schedule, err := client.GetSchedule(ctx, &params)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("test result")
	fmt.Printf("Total Games: %d\n", schedule.TotalGames)
	for _, date := range schedule.Dates {
		fmt.Printf("Date: %s\n", date.Date)
		for _, game := range date.Games {
			fmt.Printf("  GamePk: %d, Home Team: %s, Away Team: %s\n", game.GamePk, game.Teams.Home.Team.Name, game.Teams.Away.Team.Name)
			fmt.Println("    " + game.Status.DetailedState)
		}
	}

}
