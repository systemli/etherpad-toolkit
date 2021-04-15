package purge

import (
	"strings"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/systemli/etherpad-toolkit/pkg"
)

// Purger
type Purger struct {
	Etherpad *pkg.Etherpad
	DryRun   bool
}

// NewPurger returns a instance of Purger.
func NewPurger(ep *pkg.Etherpad, dryRun bool) *Purger {
	return &Purger{
		Etherpad: ep,
		DryRun:   dryRun,
	}
}

// PurgePads loops over a sorted map of pads and removes pads which are not edited for some times.
func (p *Purger) PurgePads(sorted map[string][]string, concurrency int) {
	var wg sync.WaitGroup

	for suffix, padIds := range sorted {
		wg.Add(1)
		go p.processPads(padIds, suffix, concurrency, &wg)
	}

	wg.Wait()
}

func (p *Purger) processPads(pads []string, suffix string, concurrency int, wg *sync.WaitGroup) {
	defer wg.Done()

	log.WithFields(log.Fields{"suffix": suffix, "count": len(pads), "concurrency": concurrency}).Info("start loop")
	start := time.Now()
	in := make(chan string)
	out := make(chan int)

	for x := 0; x < concurrency; x++ {
		go p.worker(in, out)
	}
	go func() {
		for _, pad := range pads {
			in <- pad
		}
		close(in)
	}()
	for n := range out {
		if n == 0 {
			break
		}
	}

	elapsed := time.Since(start)
	log.WithFields(log.Fields{"suffix": suffix, "took": elapsed, "processed": len(pads)}).Info("finished loop")
}

func (p *Purger) worker(pads chan string, out chan int) {
	for pad := range pads {
		log.WithField("pad", pad).Debug("Process Pad")

		revisions, err := p.Etherpad.GetRevisionsCount(pad)
		if err != nil {
			log.WithError(err).WithField("pad", pad).Error("failed to get last edited time")
			continue
		}

		lastEdited, err := p.Etherpad.GetLastEdited(pad)
		if err != nil {
			log.WithError(err).Error("")
			continue
		}

		deletable := lastEdited.Before(time.Now().Add(padDuration(pad))) || revisions == 0
		if !deletable {
			continue
		}

		log.WithFields(log.Fields{"pad": pad, "lastEdited": lastEdited, "revisions": revisions}).Info("Delete Pad")
		if p.DryRun {
			continue
		}
		err = p.Etherpad.DeletePad(pad)
		if err != nil {
			log.WithError(err).WithField("pad", pad).Error("failed to delete pad")
		}
	}
	out <- 0
}

// padDuration returns the time frame in which the pad should be edited.
func padDuration(padID string) time.Duration {
	if strings.HasSuffix(padID, "-keep") {
		return -365 * 24 * time.Hour
	} else if strings.HasSuffix(padID, "-temp") {
		return -24 * time.Hour
	} else {
		return -30 * 24 * time.Hour
	}
}
