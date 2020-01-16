package main

import (
  "fmt"
  "sync"
  // "time"
)

/*func main() {
  var wg sync.WaitGroup
  for _, salutation := range []string{"hello", "greetings", "good day"} {
    wg.Add(1)
    go func() {
      defer wg.Done()
      fmt.Println(salutation)
    }()
  }
  wg.Wait()
  // good day x 3
}
*/

// its ok:

func main() {
  var wg sync.WaitGroup
  for _, salutation := range []string{"hello", "greetings", "good day"} {
    wg.Add(1)
    go func(salutation string) {
      defer wg.Done()
      fmt.Println(salutation)
    }(salutation)
  }
  wg.Wait()
}
