package presto

type StatementStats struct {
	State          string     `json:"state"`
	Queued         bool       `json:"queued"`
	Nodes          int        `json:"nodes"`
	TotalSplits    int        `json:"totalSplits"`
	QueuedSplits   int        `json:"queuedSplits"`
	RunningSplits  int        `json:"runningSplits"`
	CompleteSplits int        `json:"completeSplits"`
	UserMillis     int        `json:"userMillis"`
	CpuMillis      int        `json:"cpuMillis"`
	WallTimeMillis int        `json:"wallTimeMillis"`
	ProcessedRows  int        `json:"processedRows"`
	ProcessedBytes int        `json:"processedBytes"`
	RootState      StageStats `json:"rootState"`
}
