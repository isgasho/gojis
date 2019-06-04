package job

const (
	QueueBuffer = 10 // TODO: evaluate what could be a better value
)

type Queue struct {
	q chan PendingJob
}

func NewQueue() *Queue {
	q := new(Queue)
	q.q = make(chan PendingJob, QueueBuffer)
	return q
}

func (q *Queue) Enqueue(j PendingJob) {
	q.q <- j
}

func (q *Queue) Dequeue() (PendingJob, bool) {
	j, ok := <-q.q
	return j, ok
}
