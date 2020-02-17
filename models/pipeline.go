package models

// No native imp// lementation of enums in Go
type PipelineResult struct {
	// By convention, 1 will be success, 0 will be failure
	Status uint

	// If there is a failure, we guarantee that that the Message field is filled
	// and the reason is surfaced through Message.
	Message string
}

type Pipeline interface {
	New(*MongoDB) Pipeline
	Run(map[string]string) *PipelineResult
}
