package main

import (
	"context"
	"log"
	"sync"
	"testdb/database"
	"testdb/utils"
)

func main() {
	err := utils.Init(".")
	if err != nil {
		log.Fatal(err)
	}
	pgDb := database.OpenPg()
	redis := database.OpenRedis()
	defer pgDb.Close()
	defer redis.Close()
	var wg sync.WaitGroup

	//6 goroutines to process messages
	for i := 0; i < 6; i++ {
		wg.Add(1)
		latestOffset := database.GetLatestOffset(i)
		go func(partition int, latestOffset int64) {
			consumer := database.GetKafkaConsumer(partition)
			defer consumer.Close()
			defer wg.Done()
			for {
				m, err := consumer.ReadMessage(context.Background())
				if err != nil {
					log.Fatal(err)
					break
				}
				utils.ProcessMessage(pgDb, redis, m.Value)
				if m.Offset >= latestOffset {
					log.Printf("Partition %d has finished processing at offset %d", partition, m.Offset)
					break
				}
			}
		}(i, latestOffset)
	}

	wg.Wait()
}
