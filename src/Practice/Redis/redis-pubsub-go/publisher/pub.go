package main

import (
 "bufio"
 "context"
 "fmt"
 "github.com/go-redis/redis/v8"
 "os"
 "time"
)

func main() {
 ctx := context.Background()

 // Create a new Redis client
 rdb := redis.NewClient(&redis.Options{
  Addr:     "localhost:6379", // Redis server address
  Password: "",               // no password set
  DB:       0,                // default DB
 })

 scanner := bufio.NewScanner(os.Stdin)
 fmt.Println("Enter messages to send to the subscriber. Type 'exit' to quit.")

 // Loop to accept input from the terminal
 for scanner.Scan() {
  message := scanner.Text()
  if message == "exit" {
   break
  }
  ist, err := time.LoadLocation("Asia/Kolkata")
  if err != nil {
   fmt.Println("Error loading IST timezone:", err)
   return
  }
  currentTime := time.Now().In(ist).Format("Jan 02, 2006 at 3:04pm (MST)")
  // Publish the input message to the 'updates' channel
  err = rdb.Publish(ctx, "updates", message).Err()
  if err != nil {
   fmt.Printf("[%s] - Error publishing message:%s\n", currentTime, err.Error())
   continue
  }

  fmt.Printf("[%s] - Message published: [%s]\n", currentTime, message)
 }
}
