package main

import (
	"fmt"
	"time"
)

func main() {
	// you can write to stdout for debugging purposes, e.g.
	fmt.Println("This is a debug message")

	portScan = make(map[string]arr)
}

type log_t struct {
	src_ip   string
	dst_ip   string
	dst_port int
	time     time.Time
}

type port struct {
	dst_port int
	time     time.Time
}

type arr []port
type portScan map[string]arr

// maintain arr of 10 elements
// account only last 1 second dst_ports

func process_log(log log_t) {

	// check for edge cases

	// process the log

	val, exists := portScan[log.src_ip+log.dst_ip]
	if exists {
		insertLog(log, val)
	} else {
		p := new(port)
		p.dst_port = log.dst_port
		p.time = time.Now()
		array := make([]port)
		array = append(arr, p)
		portScan[log.src_ip+log.dst_ip] = array
	}
}

func insertLog(log log_t, val arr) {

	if len(arr) < 10 { // less than 10
		p := new(port)
		p.dst_port = log.dst_port
		p.time = time.Now()
		// array := make([] port)
		// array = append(arr, p)
		// portScan[log.src_ip+log.dst_ip] = array

		// scan to make sure it is unique port
		for i := 0; i < len(arr); i++ {
			if log.dst_port == arr.dst_port {
				// found duplicate dst_port - unlink
				arr = arr[0:i] + arr[i+1:]
			}
		}
		arr = append(arr, p)
		portScan[log.src_ip+log.dst_ip] = arr
	}

	if len(arr) == 10 {
		// generate alert if arr[0].time - time.Now() <= 1 second
		t := time.Now()
		elapsed := t.Sub(arr[0])
		if elapsed <= 1000 {
			// generate the alert here
			gen_alert(log.src_ip, log.dst_ip)
		}
	}

	// shrink the slice if ports learned before 1 second
	i := 0
	for ; i < len(val); i++ {
		elapsed := t.Sub(arr[i])
		if elapsed > 1000 {
			// do nothing
		} else {
			break
		}
	}
	arr = arr[i:]
	portScan[log.src_ip+log.dst_ip] = arr

}
