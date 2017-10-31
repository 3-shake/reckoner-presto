package presto

type StageStats struct {
	StageID        string       `json:"stageId"`
	State          string       `json:"state"`
	Done           string       `json:"done"`
	Nodes          string       `json:"nodes"`
	TotalSplits    int          `json:"totalSplits"`
	QueuedSplits   int          `json:"queuedSplits"`
	RunningSplits  int          `json:"runningSplits"`
	CompleteSplits int          `json:"completeSplits"`
	UserMillis     int          `json:"userMillis"`
	CpuMillis      int          `json:"cpuMillis"`
	WallTimeMillis int          `json:"wallTimeMillis"`
	ProcessedRows  int          `json:"processedRows"`
	ProcessedBytes int          `json:"processedBytes"`
	SubStages      []StageStats `json:"subStages"`
}
