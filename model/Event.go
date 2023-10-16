package model

type Event struct {
	EventId          string
	SubmissionId     string
	Runtime          int
	CodeBucket       string
	FirstEventEpoch  int64
	LatestEventEpoch int64
	Retries          int
}
