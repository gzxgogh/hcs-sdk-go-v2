package model

type JobResponse struct {
	JobId string `json:"job_id"`
	Error Error  `json:"error,omitempty"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type JobErr struct {
	ErrorMsg  string `json:"error_msg"`
	ErrorCode string `json:"error_code"`
	RequestId string `json:"request_id"`
}

type JobInfo struct {
	JobId      string `json:"job_id"`
	JobType    string `json:"job_type"`
	BeginTime  string `json:"begin_time"`
	EndTime    string `json:"end_time"`
	Status     string `json:"status"`
	ErrorCode  string `json:"error_code"`
	FailReason string `json:"fail_reason"`
	Entities   struct {
		SubJobsTotal int          `json:"sub_jobs_total"`
		SubJobs      []SubJobInfo `json:"sub_jobs"`
	} `json:"entities"`
}

type SubJobInfo struct {
	JobId      string                 `json:"job_id"`
	JobType    string                 `json:"job_type"`
	BeginTime  string                 `json:"begin_time"`
	EndTime    string                 `json:"end_time"`
	Status     string                 `json:"status"`
	ErrorCode  string                 `json:"error_code"`
	FailReason string                 `json:"fail_reason"`
	Entities   map[string]interface{} `json:"entities"`
}
