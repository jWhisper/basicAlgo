package schedule

import (
	"context"

	"github.com/marusama/cyclicbarrier"
	"golang.org/x/sync/semaphore"
)

//H20制造工厂
type H2O struct {
	semaH *semaphore.Weighted
	semaO *semaphore.Weighted

	b cyclicbarrier.CyclicBarrier
}

func NewH2O() *H2O {
	return &H2O{
		semaH: semaphore.NewWeighted(2),
		semaO: semaphore.NewWeighted(1),
		b:     cyclicbarrier.New(3),
	}
}

func (h2o *H2O) produceH(produceHFunc func() error) error {
	h2o.semaH.Acquire(context.Background(), 1)
	defer h2o.semaH.Release(1)
	if err := produceHFunc(); err != nil {
		return err
	}
	//在栅栏前排队等候
	h2o.b.Await(context.Background())
	return nil
}

func (h2o *H2O) produceO(produceOFunc func() error) error {
	h2o.semaO.Acquire(context.Background(), 1)
	defer h2o.semaO.Release(1)
	if err := produceOFunc(); err != nil {
		return err
	}
	//在栅栏前排队等候
	h2o.b.Await(context.Background())
	return nil
}
