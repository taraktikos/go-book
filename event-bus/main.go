package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type DataEvent struct {
	Data  interface{}
	Topic string
}

type DataChannel chan DataEvent

type DataChannelSlice []DataChannel

type EventBus struct {
	subscribers map[string]DataChannelSlice
	rm          sync.RWMutex
}

var eb = &EventBus{
	subscribers: map[string]DataChannelSlice{},
}

func main() {
	ch1 := make(chan DataEvent)
	ch2 := make(chan DataEvent)
	ch3 := make(chan DataEvent)

	eb.Subscribe("topic1", ch1)
	eb.Subscribe("topic2", ch2)
	eb.Subscribe("topic2", ch3)

	go publishTo("topic1", "Hi topic 1")
	go publishTo("topic2", "Welcome topic2")

	for {
		select {
		case d := <-ch1:
			go printDataEvent("ch1", d)
		case d := <-ch2:
			go printDataEvent("ch2", d)
		case d := <-ch3:
			go printDataEvent("ch3", d)
		}
	}
}

func publishTo(topic string, data string) {
	for {
		eb.Publish(topic, data)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

func printDataEvent(ch string, data DataEvent) {
	fmt.Printf("Channel: %s; Topic: %s; DataEvent: %v\n", ch, data.Topic, data.Data)
}

func (eb *EventBus) Subscribe(topic string, ch DataChannel) {
	eb.rm.Lock()
	defer eb.rm.Unlock()

	eb.subscribers[topic] = append(eb.subscribers[topic], ch)
}

func (eb *EventBus) Publish(topic string, data interface{}) {
	eb.rm.RLock()
	defer eb.rm.RUnlock()

	if chans, ok := eb.subscribers[topic]; ok {
		go func(data DataEvent, dataChannelSlices DataChannelSlice) {
			for _, ch := range dataChannelSlices {
				ch <- data
			}
		}(DataEvent{data, topic}, chans)
	}
}
