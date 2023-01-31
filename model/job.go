package model

type JobResponse struct {
	JobId string `json:"job_id"`
}

type JobInfo struct {
	JobId      string      `json:"job_id"`
	JobType    string      `json:"job_type"`
	BeginTime  string      `json:"begin_time"`
	EndTime    string      `json:"end_time"`
	Status     string      `json:"status"`
	ErrorCode  interface{} `json:"error_code"`
	FailReason interface{} `json:"fail_reason"`
	Entities   interface{} `json:"entities"`
}
