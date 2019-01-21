package collector

import (
	"math/rand"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

// Metrics 指标结构体
type Metrics struct {
	metrics map[string]*prometheus.Desc
	mutex   sync.Mutex
}

// newGlobalMetric创建指标描述符
func newGlobalMetric(namespace string, metricName string, docString string, labels []string) *prometheus.Desc {
	return prometheus.NewDesc(namespace+"_"+metricName, docString, labels, nil)
}

// NewMetrics 功能：初始化指标信息，即Metrics结构体
func NewMetrics(namespace string) *Metrics {
	return &Metrics{
		metrics: map[string]*prometheus.Desc{
			"my_counter_metric": newGlobalMetric(namespace, "my_counter_metric", "The description of my_counter_metric", []string{"host"}),
			"my_gauge_metric":   newGlobalMetric(namespace, "my_gauge_metric", "The description of my_gauge_metric", []string{"host"}),
		},
	}
}

// Describe 功能：传递结构体中的指标描述符到channel
func (c *Metrics) Describe(ch chan<- *prometheus.Desc) {
	for _, m := range c.metrics {
		ch <- m
	}
}

// Collect 功能：抓取最新的数据，传递给channel
func (c *Metrics) Collect(ch chan<- prometheus.Metric) {
	c.mutex.Lock() // 加锁
	defer c.mutex.Unlock()

	mockCounterMetricData, mockGaugeMetricData := c.GenerateMockData()
	for host, currentValue := range mockCounterMetricData {
		ch <- prometheus.MustNewConstMetric(c.metrics["my_counter_metric"], prometheus.CounterValue, float64(currentValue), host)
	}
	for host, currentValue := range mockGaugeMetricData {
		ch <- prometheus.MustNewConstMetric(c.metrics["my_gauge_metric"], prometheus.GaugeValue, float64(currentValue), host)
	}
}

// GenerateMockData 功能：生成模拟数据
func (c *Metrics) GenerateMockData() (mockCounterMetricData map[string]int, mockGaugeMetricData map[string]int) {
	mockCounterMetricData = map[string]int{
		"yahoo.com":  int(rand.Int31n(1000)),
		"google.com": int(rand.Int31n(1000)),
	}
	mockGaugeMetricData = map[string]int{
		"yahoo.com":  int(rand.Int31n(10)),
		"google.com": int(rand.Int31n(10)),
	}
	return
}
