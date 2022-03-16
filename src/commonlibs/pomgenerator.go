package commonlibs

import (
	"fmt"
	"math/rand"
	"time"
)

func GeneratePomMetrics(metricsVO *MetricsVO) (err error) {
	// fmt.Printf("hello world\n")
	// fmt.Printf("%+v\n", metricsVO.Metrics)
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < len(metricsVO.Metrics); i++ {
		a := 950 + (rand.Intn(100) - 50)
		b := 1000 + (rand.Intn(10) - 5)
		fmt.Println("csacam_"+metricsVO.Metrics[i].MetricID+
			" {authorization_boundary=\"foo\"} ", int(100*(float32(a)/float32(b))))
		fmt.Println("csacam_"+metricsVO.Metrics[i].MetricID+"_measureA"+
			" {authorization_boundary=\"foo\"} ", a)
		fmt.Println("csacam_"+metricsVO.Metrics[i].MetricID+"_measureB"+
			" {authorization_boundary=\"foo\"} ", b)
	}

	return
}
