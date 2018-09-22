package main

import (
	"time"
	"os"
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"github.com/mikoim/go-loadavg"
	"runtime"
	"sync"
	"flag"
)


func main() {
	err := run()
	if err != nil {
		fmt.Printf("[ERR] %s\n",err)
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}

func run() (err error) {
	err = nil
	err = rpio.Open()
	defer rpio.Close()
	if err != nil {
		return
	}

	var (
		v = flag.Bool("v",false,"verbose mode")
	)
	flag.Parse()

	var pin rpio.Pin
	pin = rpio.Pin(21)
	pin.Output()
	pin.Low()

	var wait time.Duration = 1000
	m := new(sync.Mutex)

	go func() {
		for {

			var lv *loadavg.Loadavg
			lv, err = loadavg.Parse()
			if err != nil {
				return
			}

			var load float64 = lv.LoadAverage1
			load = 1000 - (load/float64(runtime.NumCPU())) * 500
			if load < 15 {
				load = 15
			}
			if *v {
				fmt.Printf("wait %f milliseconds\n",load)
			}

			m.Lock()
			wait = time.Duration(load) * time.Millisecond
			m.Unlock()
			time.Sleep(time.Millisecond * 1000)
		}
	}()

	for {
		m.Lock()
		waittime := wait
		m.Unlock()

		pin.Low()
		time.Sleep(waittime)

		pin.High()
		time.Sleep(waittime)
	}
}
