package main

import (
	"container/heap"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Job struct {
	ID     string `json:"ID"`
	Type   string `json:"Type"`
	Status string `json:"Status"`
}

// statusMap - In-memory map for tracking/updating Job statuses
var statusMap = make(map[string]Job)

// Enqueue - Adds a job to the queue and returns the ID of the job
func Enqueue(c *gin.Context) {
	job := Job{}

	err := c.ShouldBind(&job)
	if err != nil {
		log.Println("error unmarshalling job:", err)
	}

	job.ID = uuid.New().String()
	job.Status = "QUEUED"

	queueItem1 := &Item{
		value:    job,
		priority: StringToPrio(job.Type),
	}

	heap.Push(Queue, queueItem1)
	Queue.update(queueItem1, queueItem1.value, queueItem1.priority)

	statusMap[job.ID] = job

	c.JSON(http.StatusCreated, gin.H{"job_id": job.ID})
}

// Dequeue - Removes a job from the queue and returns it. A Job is
// considered available for Dequeue if the job has not been concluded
// and has not dequeued already
func Dequeue(c *gin.Context) {
	job := heap.Pop(Queue).(*Item).value

	j := statusMap[job.ID]
	j.Status = "IN_PROGRESS"
	statusMap[job.ID] = j

	c.JSON(http.StatusOK, gin.H{"job": j})
}

// Conclude - Finish execution on the job and consider it done
func Conclude(c *gin.Context) {
	id := c.Param("job_id")
	fmt.Println("ID::", id)

	// SEND CALLS TO CANCEL/FINISH WORK TO APPROPRIATE SYSTEMS/SERVICES
	// THEN UPDATE STATUS MAP AND RETURN OR RETURN NOT FOUND

	if _, ok := statusMap[id]; ok {
		job := statusMap[id]
		job.Status = "CONCLUDED"
		statusMap[job.ID] = job

		c.JSON(http.StatusNoContent, gin.H{})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"job_id": id, "status": "job id not found"})
	}
}

// GetJob - Returns job information from given ID
func GetJob(c *gin.Context) {
	id := c.Param("job_id")

	if job, ok := statusMap[id]; ok {
		c.JSON(http.StatusOK, gin.H{"job": job})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "job not found"})
	}
}
