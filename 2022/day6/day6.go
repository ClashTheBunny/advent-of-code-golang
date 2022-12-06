package main

import (
	"fmt"
	"log"
	"os"
	// "strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("2022/day6/day6.txt")
	if err != nil {
		log.Fatal(err)
	}
	data_stream := strings.TrimSpace(string(data))

  _ = data_stream

  result := 0

  for i, _ := range data_stream {
    if i > 3{
      dataMap := map[string]bool{
        string(data_stream[i]): true,
        string(data_stream[i-1]): true,
        string(data_stream[i-2]): true,
        string(data_stream[i-3]): true,
      }
      if len(dataMap) == 4 {
        result = i + 1
        break
      }
    }
  }

  fmt.Printf("%v\n", result)

  for i, _ := range data_stream {
    if i > 13{
      dataMap := map[string]bool{
        string(data_stream[i]): true,
        string(data_stream[i-1]): true,
        string(data_stream[i-2]): true,
        string(data_stream[i-3]): true,
        string(data_stream[i-4]): true,
        string(data_stream[i-5]): true,
        string(data_stream[i-6]): true,
        string(data_stream[i-7]): true,
        string(data_stream[i-8]): true,
        string(data_stream[i-9]): true,
        string(data_stream[i-10]): true,
        string(data_stream[i-11]): true,
        string(data_stream[i-12]): true,
        string(data_stream[i-13]): true,
      }
      if len(dataMap) == 14 {
        result = i + 1
        break
      }
    }
  }

  fmt.Printf("%v\n", result)
}
