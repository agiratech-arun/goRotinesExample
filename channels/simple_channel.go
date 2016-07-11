// simple smaple

package simple_channel

import ("fmt"
        // "time"
       )


func example() {
  c := make(chan int)
  value := 5
  go gorotine(value,"gorotine",c)
  // gorotine(10,"normal")
  for {
    if (0==<-c) {
      fmt.Println(<-c)
      break
    }
  }
  var str string
  fmt.Scanln(&str)
  fmt.Println("Done")
}

func gorotine(a int, str string,c chan int) {
  for i:=1;i<=a;i++ {
    fmt.Println(i)
    c <- i
  }
  close(c)
}