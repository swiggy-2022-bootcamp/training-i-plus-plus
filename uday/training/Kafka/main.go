package main

import (
   appKafka "./kafka"
   "fmt"
   "time"
)

func main()  {

   fmt.Println("Okay...")
   go appKafka.StartKafka()

   fmt.Println("Kafka has been started...")

   time.Sleep(10 * time.Minute)

}
