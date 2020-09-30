package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"sync"
)

var (
	wg sync.WaitGroup
)

func main() {
	consumer, err := sarama.NewConsumer([]string{"119.29.117.244:9092"}, nil)
	if err != nil {
		panic(err)
	}

	partitionList, err := consumer.Partitions("hello")
	if err != nil {
		panic(err)
	}
	defer consumer.Close()
	for partition := range partitionList {
		pc, err := consumer.ConsumePartition("hello", int32(partition), sarama.OffsetNewest)
		if err != nil {
			panic(err)
		}
		wg.Add(1)
		go func(sarama.PartitionConsumer) {
			defer wg.Done()
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d, Offset:%d, Key:%s, Value:%s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
		}(pc)
		// pc.AsyncClose()
	}
	wg.Wait()
}
