package scheduler

import (
	"github.com/go-co-op/gocron"
	"github.com/go-nunu/nunu-layout-advanced/pkg/log"
	"go.uber.org/zap"
	"time"
)

type Scheduler struct {
	scheduler *gocron.Scheduler
}

func NewScheduler(log *log.Logger) *Scheduler {
	gocron.SetPanicHandler(func(jobName string, recoverData interface{}) {
		log.Error("TaskServer Panic", zap.String("job", jobName), zap.Any("recover", recoverData))
	})
	return &Scheduler{
		scheduler: gocron.NewScheduler(time.FixedZone("PRC", 8*60*60)),
	}
}

func (s *Scheduler) Start() {
	s.scheduler.StartBlocking()
}

func (s *Scheduler) Stop() {
	s.scheduler.Stop()
}

func (s *Scheduler) Set() *gocron.Scheduler {
	return s.scheduler
}
