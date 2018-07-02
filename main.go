package main

import "runtime"
import "fmt"
import "time"

/*
#include "collatz.h"
#cgo CFLAGS: -march=native -O3
*/
import "C"


var pl = fmt.Println
var pf = fmt.Printf

// COLLATZ/CONTROL
func control(start uint64, ret uint64, blockSize uint64, ch chan uint64) {
  start = start
  ret = uint64(C.collatz(C.ulong(start), C.ulong(blockSize), C.ulong(ret)))

  ch<-ret
}

// MAIN
func main() {
  var start uint64 = 3
  var end uint64 =  400000000
  var threads int = runtime.NumCPU()
  var blockSize uint64 = 2000000
  ch := make(chan uint64, threads)

  _ = runtime.GOMAXPROCS(threads)
  for i:=0; i<threads; i++ {
    go control(start, uint64(2), blockSize, ch)
    time.Sleep(100 * time.Millisecond)
    start += blockSize
  }


  // define and execute status-timer thread
  status := func() {
    begin := time.Now()
    delay := 1 * time.Second
    rate := uint64(0)
    last := uint64(0)

    for {
      time.Sleep(delay)
      pf("\r                                               ")
      pf("\rCur: %d\tRate: %d/min          ", start, rate)

      if time.Since(begin) > (10 * time.Second) {
        begin = time.Now()
        rate = ((start - last) * 6)
        last = start
      }
    }
  }
  go status()

  // spawn control threads to compute collatz sequences
  tmpRec := uint64(2)
  for start < end {
    select {
      case ret := <-ch:
        // if returned record is larger than last (tmpRec) record
        // set tmpRec to new record value
        if ret > tmpRec {
          tmpRec = ret
        }
        start += blockSize
        go control(start, tmpRec, blockSize, ch)
        continue
    }
  }
}


