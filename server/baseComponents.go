package server

import (
	"github.com/curio-research/keystone/keystone/ecs"
)

type JobSchema struct {
	JobType        string
	JobId          string
	TickDataString string
	TickNumber     int
	TickUuid       string
}

var (
	JobTable = ecs.NewTable[JobSchema]()
)

func AddTickJob(w *ecs.GameWorld, tickNumber int, jobType string, jobData string, tickId string, uuid string) int {

	entity := JobTable.Add(w, JobSchema{
		JobType:        jobType,
		JobId:          tickId,
		TickDataString: jobData,
		TickNumber:     tickNumber,
		TickUuid:       uuid,
	})

	return entity
}

// registers default components keystone must operates on such as tick related
func RegisterDefaultComponents(w *ecs.GameWorld) {
	w.AddTables(JobTable)
}
