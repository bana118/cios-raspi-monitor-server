package main

import (
	"fmt"

	"github.com/bana118/cios-raspi-monitor-server/pkg/example"
)

func main() {
	buckets := example.SampleGetBucket()
	fmt.Print(buckets)
}
