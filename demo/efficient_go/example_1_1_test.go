package efficientgo

import (
	"fmt"
	"testing"
	"time"
)

type Report struct{}

func (r Report) Error() any { return nil }

type ReportGetter interface {
	Get() []Report
}

func FailureRatio(reports ReportGetter) float64 {
	// 三次调用 reports.Get()，性能损耗
	if len(reports.Get()) == 0 {
		return 0
	}
	var sum float64
	for _, report := range reports.Get() {
		if report.Error() != nil {
			sum++
		}
	}
	return sum / float64(len(reports.Get()))
}

func TestTesting(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过")
	}
	time.Sleep(10 * time.Second)
	fmt.Println("ok")
}
