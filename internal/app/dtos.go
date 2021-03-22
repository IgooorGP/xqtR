package app

type JobsFile struct {
	Jobs []Job `mapstructure:"jobs"`
}

type Job struct {
	Title      string    `mapstructure:"title"` // job name
	Steps      []JobStep `mapstructure:"steps"`
	NumWorkers int       `mapstructure:"num_workers"`
}

type JobStep struct {
	Name string `mapstructure:"name"`
	Run  string `mapstructure:"run"`
}

func NewJobsFile() *JobsFile {
	return &JobsFile{}
}
