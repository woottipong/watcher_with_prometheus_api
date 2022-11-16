package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"strconv"
	"time"

	"arbitrage_watcher/internal/config"
	"arbitrage_watcher/internal/model"
	"arbitrage_watcher/internal/networking"

	"github.com/juunini/simple-go-line-notify/notify"
)

func main() {
	config := config.Load("config.yaml")
	watchingInterval := config.Interval * time.Second

	// Watching
	latestSwap := make(map[int]float64)
	for {
		// Query
		t := time.Now().In(time.UTC)
		url := fmt.Sprintf("%s/query?query={strategy=\"JUNO-bJUNO\"}&time=%s", config.PrometheusApi, t.Format(time.RFC3339))
		resp, err := networking.GetContent(url)
		if err != nil {
			log.Fatal("failed!", err.Error())
		}
		res := model.Response{}
		json.Unmarshal(resp, &res)

		swapResult, _ := strconv.ParseFloat(res.Data.Results[0].Value[1], 64)
		expected := config.Pairs[0]
		if swapResult >= expected.ExpectedProfitForward { // ProfitForward
			if (math.Abs(math.Abs(latestSwap[0])-math.Abs(swapResult)) > expected.ExpectedProfileStep) || latestSwap[0] == 0 {
				// line notify
				message := fmt.Sprintf("🟢 %s → %s (%.2f%%)",
					config.Pairs[0].Token.Input, config.Pairs[0].Token.Output, swapResult)
				if err := notify.SendText(config.Webhooks.LineNotifyToken, message); err != nil {
					log.Println(err)
				}
				log.Println(message)
				latestSwap[0] = swapResult
			}
		} else if swapResult <= expected.ExpectedProfitBackward { // ProfitBackward
			if (math.Abs(math.Abs(latestSwap[0])-math.Abs(swapResult)) > expected.ExpectedProfileStep) || latestSwap[0] == 0 {
				// line notify
				message := fmt.Sprintf("🔴 %s → %s (%.2f%%)",
					config.Pairs[0].Token.Input, config.Pairs[0].Token.Output, swapResult)
				if err := notify.SendText(config.Webhooks.LineNotifyToken, message); err != nil {
					log.Println(err)
				}
				log.Println(message)
				latestSwap[0] = swapResult
			}
		} else {
			latestSwap[0] = 0
		}

		// Interval
		if watchingInterval == 0 {
			return
		}
		time.Sleep(watchingInterval)
	}
}
