package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"strings"
	"strconv"
	"os/exec"
)

type lvmCollector struct {
	vgFreeMetric *prometheus.Desc
	vgSizeMetric *prometheus.Desc
}

// LVM Collector contains VG size and VG free in MB
func newLvmCollector() *lvmCollector {
	return &lvmCollector{
		vgFreeMetric: prometheus.NewDesc("lvm_vg_free_bytes",
			"Shows LVM VG free size in MB",
			[]string{"vg_name"}, nil,
		),
		vgSizeMetric: prometheus.NewDesc("lvm_vg_bytes_total",
			"Shows LVM VG total size in MB",
			[]string{"vg_name"}, nil,
		),
	}
}

func (collector *lvmCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.vgFreeMetric
	ch <- collector.vgSizeMetric
}

// LVM Collect, call OS command and set values
func (collector *lvmCollector) Collect(ch chan<- prometheus.Metric) {
	out, err := exec.Command("/sbin/vgs", "--units", "B", "--separator", ",", "-o", "vg_name,vg_free,vg_size", "--noheadings").Output()
	if err != nil {
		log.Print(err)
	}
	lines := strings.Split(string(out),"\n")
	for _, line := range lines {
		values := strings.Split(line,",")
		if len(values)==3 {
			free_size, err := strconv.ParseFloat(strings.Trim(values[1],"B"), 64)
			if err!= nil {
				log.Print(err)
			} else {
				total_size, err := strconv.ParseFloat(strings.Trim(values[2],"B"), 64)
				if err!= nil {
					log.Print(err)
				} else {
					vg_name := strings.Trim(values[0], " ")
					ch <- prometheus.MustNewConstMetric(collector.vgFreeMetric, prometheus.GaugeValue, free_size, vg_name)
					ch <- prometheus.MustNewConstMetric(collector.vgSizeMetric, prometheus.GaugeValue, total_size, vg_name)
				}
			}
		}
	}

}
