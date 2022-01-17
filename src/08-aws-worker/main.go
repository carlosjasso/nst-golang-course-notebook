package main

import (
	"context"
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

func process(ctx context.Context, delivery amqp.Delivery) error {
	return nil
	// delivery.Ack(true) // In case message wants to be acked
}

func deliver(ctx context.Context, prefetchCount int, queueName string, f func(context.Context, amqp.Delivery) error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("!ERROR: ", err)
		}

		time.Sleep(1 * time.Second)
		deliver(ctx, prefetchCount, queueName, f)
	}()
	conn, err := amqp.Dial("amqp://<user>:<pass>@<host>:<port>")
	if err != nil {
		fmt.Println("!ERROR: ", err)
		return
	}

	defer conn.Close()

	connErrors := make(chan *amqp.Error, 1)

	conn.NotifyClose(connErrors)

	channel, err := conn.Channel()
	if err != nil {
		fmt.Println("!ERROR: ", err)
		return
	}

	defer channel.Close()

	channelErrors := make(chan *amqp.Error, 1)
	channel.NotifyClose(channelErrors)

	if err := channel.Qos(prefetchCount, 0, false); err != nil {
		return
	}

	queue, err := channel.QueueDeclare("image", true, false, false, false, nil)
	if err != nil {
		fmt.Println("!ERROR: ", err)
		return
	}

	deliveries, err := channel.Consume(queue.Name, "", false, false, false, false, nil)
	if err != nil {
		fmt.Println("!ERROR: ", err)
		return
	}

	closed := false
	for !closed {
		select {
		case <-connErrors:
			_, ok := <-connErrors
			if ok {
				close(connErrors)
			}
			time.Sleep(1 * time.Second)
			closed = true
		case <-channelErrors:
			_, ok := <-connErrors
			if ok {
				close(connErrors)
			}
			time.Sleep(1 * time.Second)
			closed = true
		case delivery := <-deliveries:
			if err := f(ctx, delivery); err != nil {
				fmt.Println("!FATAL ERRORL: ", err)
			}
		}
	}
}

func main() {
	ctx := context.Background()
	ch := make(chan bool)
	for i := 0; i < 3; i++ {
		go func() {
			deliver(ctx, 3, "image", process)
		}()
	}
	<-ch
}
