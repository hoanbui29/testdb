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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	//1000 goroutines to process messages
	for i := 0; i < 6; i++ {
		wg.Add(1)
		go func(partition int) {
			consumer := database.GetKafkaConsumer(partition)
			defer consumer.Close()
			defer wg.Done()
			for {
				m, err := consumer.ReadMessage(ctx)
				if err != nil {
					log.Fatal(err)
					break
				}
				utils.ProcessMessage(pgDb, redis, m.Value)
			}
		}(i)
	}
	wg.Wait()
}
