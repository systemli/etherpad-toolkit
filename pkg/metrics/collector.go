package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
	"github.com/systemli/etherpad-toolkit/pkg"
	"github.com/systemli/etherpad-toolkit/pkg/helper"
)

type PadCollector struct {
	etherpad     *pkg.Etherpad
	PadGaugeDesc *prometheus.Desc
}

func NewPadCollector(etherpad *pkg.Etherpad) *PadCollector {
	return &PadCollector{
		etherpad:     etherpad,
		PadGaugeDesc: prometheus.NewDesc("etherpad_toolkit_pads", "The current number of pads", []string{"suffix"}, nil),
	}
}

func (pc *PadCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- pc.PadGaugeDesc
}

func (pc *PadCollector) Collect(ch chan<- prometheus.Metric) {
	allPads, err := pc.etherpad.ListAllPads()
	if err != nil {
		log.WithError(err).Error("failed to list all allPads")
		return
	}

	sorted := helper.SortPads(allPads)

	for suffix, pads := range sorted {
		ch <- prometheus.MustNewConstMetric(
			pc.PadGaugeDesc,
			prometheus.GaugeValue,
			float64(len(pads)),
			suffix,
		)
	}
}
