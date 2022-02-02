package main

import (
	"github.com/google/uuid"
	"github.com/jasonlvhit/gocron"
	"goodwe_monitor/application"
	"goodwe_monitor/config"
	"goodwe_monitor/domain/reading"
	"goodwe_monitor/infrastructure/persistence"
	"log"
	"strconv"
	"strings"
	"time"
)

func task() {
    if application.Login() {
		readingResponse := application.GetCurrentReadings()

		powerFlow := readingResponse.Data.Powerflow
		//charts := readingResponse.Data.EnergeStatisticsCharts
		//inverter_info := readingResponse.Data.Inverter

		currentPvGeneration, _ := strconv.ParseInt(strings.Trim(powerFlow.Pv, " (W)"), 10, 32)
		currentLoad, _ := strconv.ParseInt(strings.Trim(powerFlow.Load, " (W)"), 10, 32)
		currentGridConsumption, _ := strconv.ParseInt(strings.Trim(powerFlow.Grid, " (W)"), 10, 32)

		/*total_consumption := charts.Sum // ??
		total_buy := charts.Buy
		total_buy_percent := charts.BuyPercent

		total_sell := charts.Sell
		total_sell_percent := charts.SellPercent
		total_self_use := charts.SelfUseOfPv

		today_generation := inverter_info[0].D.EDay*/

		readingId, _ := uuid.NewUUID()

		var currentReading reading.Reading = reading.Reading{
			ID:           readingId,
			PVGeneration: currentPvGeneration,
			Load:         currentLoad,
			Grid:         currentGridConsumption,
			Timestamp:    time.Now().Unix(),
		}

		conn, err := config.ConnectDB()
		if err != nil {
			log.Print(err)
		}
		defer conn.Close()

		repo := persistence.NewReadingRepositoryWithRDB(conn)
		repo.Save(&currentReading)
	}
}

func main() {
    gocron.Every(5).Seconds().Do(task)

    <- gocron.Start()
}
