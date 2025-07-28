package server

import (
	"context"
	"github.com/go-nunu/nunu-layout-advanced/pkg/log"
	"github.com/go-nunu/nunu-layout-advanced/pkg/scheduler"
)

type TaskServer struct {
	log       *log.Logger
	scheduler *scheduler.Scheduler
}

func NewTaskServer(
	log *log.Logger,
	scheduler *scheduler.Scheduler,
) *TaskServer {
	return &TaskServer{
		log:       log,
		scheduler: scheduler,
	}
}
func (t *TaskServer) Start(ctx context.Context) error {
	// 周期性任务直接写死
	// 定时任务 通过scheduler 进行注册
	// 临时任务 通过scheduler 进行注册 并写入最大运行次数

	t.scheduler.Start()
	return nil
}
func (t *TaskServer) Stop(ctx context.Context) error {
	t.scheduler.Stop()
	t.log.Info("TaskServer stop...")
	return nil
}
