package internal

type Response struct {
	IsSuccess bool     `json:"isSuccess"`
	Result    []Result `json:"results"`
	Rows      int      `json:"rows"`
}

type Result struct {
	ID              int    `json:"ID"`
	RootID          int    `json:"rootID"`
	Publisher       string `json:"publisher"`
	DepartmentCode  string `json:"departmentCode"`
	DepartmentName  string `json:"departmentName"`
	DepartmentPhone string `json:"departmentPhone"`
	Executor        string `json:"executor"`
	ExecutorPhone   string `json:"executorPhone"`
	ExecutorEmail   string `json:"executorEmail"`
	DeductionType   string `json:"deductionType"`
	VpNum           string `json:"vpNum"`
}
