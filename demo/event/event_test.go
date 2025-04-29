package event

import (
	"testing"
	"time"

	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

type Job struct {
	ID        uint   `gorm:"primarykey"`
	Queue     string `gorm:"index"`
	Status    string `gorm:"index"` // pending, processing, completed, failed
	Data      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Queue struct {
	db *gorm.DB
}

func NewQueue(db *gorm.DB) *Queue {
	return &Queue{db: db}
}

func (q *Queue) Enqueue(queue, data string) error {
	job := &Job{
		Queue:     queue,
		Status:    "pending",
		Data:      data,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return q.db.Create(job).Error
}

func (q *Queue) Dequeue(queue string) (*Job, error) {
	var job Job
	err := q.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("queue = ? AND status = ?", queue, "pending").
			Order("created_at ASC").
			First(&job).Error; err != nil {
			return err
		}
		job.Status = "processing"
		job.UpdatedAt = time.Now()
		return tx.Save(&job).Error
	})
	if err != nil {
		return nil, err
	}
	return &job, nil
}

func (q *Queue) Complete(job *Job) error {
	job.Status = "completed"
	job.UpdatedAt = time.Now()
	return q.db.Save(job).Error
}

func (q *Queue) Fail(job *Job) error {
	job.Status = "failed"
	job.UpdatedAt = time.Now()
	return q.db.Save(job).Error
}

func TestEvent(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	require.NoError(t, err)

	// 创建表
	err = db.AutoMigrate(&Job{})
	require.NoError(t, err)

	// 创建队列
	queue := NewQueue(db)

	// 测试入队
	err = queue.Enqueue("default", "test job")
	require.NoError(t, err)

	// 测试出队
	job, err := queue.Dequeue("default")
	require.NoError(t, err)
	require.Equal(t, "test job", job.Data)

	// 测试完成
	err = queue.Complete(job)
	require.NoError(t, err)
}
