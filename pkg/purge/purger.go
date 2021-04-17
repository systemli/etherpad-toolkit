package purge

import (
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/systemli/etherpad-toolkit/pkg"
	"github.com/systemli/etherpad-toolkit/pkg/helper"
)

type Purger struct {
	etherpad   *pkg.Etherpad
	expiration helper.PadExpiration
	dryRun     bool
}

// NewPurger returns a instance of Purger.
func NewPurger(ep *pkg.Etherpad, exp helper.PadExpiration, dryRun bool) *Purger {
	return &Purger{
		etherpad:   ep,
		expiration: exp,
		dryRun:     dryRun,
	}
}

// PurgePads loops over a sorted map of pads and removes pads which are not edited for some times.
func (p *Purger) PurgePads(concurrency int) {
	pads, err := p.etherpad.ListAllPads()
	if err != nil {
		log.WithError(err).Error("failed to list all pads")
		return
	}

	sorted := helper.GroupPadsByExpiration(pads, p.expiration)

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

		revisions, err := p.etherpad.GetRevisionsCount(pad)
		if err != nil {
			log.WithError(err).WithField("pad", pad).Error("failed to get last edited time")
			continue
		}

		lastEdited, err := p.etherpad.GetLastEdited(pad)
		if err != nil {
			log.WithError(err).Error("")
			continue
		}

		deletable := lastEdited.Before(time.Now().Add(p.expiration.GetDuration(pad))) || revisions == 0
		if !deletable {
			continue
		}

		log.WithFields(log.Fields{"pad": pad, "lastEdited": lastEdited, "revisions": revisions}).Info("Delete Pad")
		if p.dryRun {
			continue
		}
		err = p.etherpad.DeletePad(pad)
		if err != nil {
			log.WithError(err).WithField("pad", pad).Error("failed to delete pad")
		}
	}
	out <- 0
}
