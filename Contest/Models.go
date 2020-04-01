package Contest

type ContestCollection interface {
	AddContest(contest *Contest) (*Contest, error)
	GetContests() ([]*Contest,error)
	GetContest(id int64) (*Contest, error)
	UpdateContest(contest *Contest) (*Contest, error)
	DeleteContest(contest *Contest) error
	GetContestFromInternship (id int64)  ([]*Contest, error)
}

type Contest struct {

	ContestID int64 `json:"contest_id"`
	Title string `json:"title,omitempty"`
	URL string `json:"url"`
	Message string `json:"message"`
	StartTime string `json:"starttime,omitempty"`
	EndTime string `json:"endtime,omitempty"`
	InternshipID int64 `json:"internship_id"`
}

