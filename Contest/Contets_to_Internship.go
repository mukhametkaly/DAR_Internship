package Contest

import (
	"Internship/internship"
	"errors"
)



type ContestsinInternship interface {
	CheckContest(contest *Contest)    (*Contest, error)
	AddContest (contest *Contest)    (*Contest, error)
	GetContests()                  ([]*Contest,error)
	GetContest(id int64)           (*Contest, error)
	UpdateContest(contest *Contest) (*Contest, error)
	DeleteContest(contest *Contest)            error
	GetContestsFromInternship (id int64)  ([]*Contest, error)
}

type ContestsinInternshipClass struct {
	conts ContestCollection
	internship Internship.InternshipCollection
}

func NewContestsinInternship(intcollection Internship.InternshipCollection, contestcollection ContestCollection )  ContestsinInternship {
	return &ContestsinInternshipClass{conts:contestcollection, internship: intcollection}
}

func(ContestIntrnshp *ContestsinInternshipClass) CheckContest (contest *Contest)  (*Contest, error)  {

	if contest.Title == "" {
		return nil, errors.New("No Title ")
	}
	if contest.InternshipID == 0 {
		return nil, errors.New("No Internship ID ")
	}
	if contest.URL == "" {
		return nil, errors.New("No URL ")
	}
	if contest.Message == "" {
		return nil, errors.New("No Message ")
	}
	if contest.StartTime == "" {
		return nil, errors.New("No start time ")
	}
	if contest.EndTime == "" {
		return nil, errors.New("No end time ")
	}
	_, err:=ContestIntrnshp.internship.GetInternship(contest.InternshipID)
	if err != nil {
		return nil, err
	}
	return contest, nil
}

func (ContestIntrnshp *ContestsinInternshipClass) GetContests() ([]*Contest,error) {
	contests, err:=ContestIntrnshp.conts.GetContests()
	if err != nil {
		return nil, err
	}
	return contests, err
}

func (ContestIntrnshp *ContestsinInternshipClass) GetContest(id int64)  (*Contest, error)  {

	course, err := ContestIntrnshp.conts.GetContest(id)
	if err!= nil {
		return nil, err
	}
	return course, nil
}

func (ContestIntrnshp *ContestsinInternshipClass)AddContest (contest *Contest)    (*Contest, error) {
	_, err := ContestIntrnshp.CheckContest(contest)
	if err != nil {
		return nil, err
	}
	_, err = ContestIntrnshp.conts.AddContest(contest)
	if err != nil {
		return nil, err
	}
	return contest, nil
}

func (ContestIntrnshp *ContestsinInternshipClass) UpdateContest(contest *Contest) (*Contest, error) {
	_, err := ContestIntrnshp.CheckContest(contest)
	if err != nil {
		return nil, err
	}
	_, err = ContestIntrnshp.conts.UpdateContest(contest)
	if err != nil {
		return nil, err
	}
	return contest, nil

}
func (ContestIntrnshp *ContestsinInternshipClass) DeleteContest(contest *Contest)  error {
	if contest.ContestID == 0 {
		return errors.New("NO ID")
	}
	err := ContestIntrnshp.conts.DeleteContest(contest)
	if err != nil {
		return err
	}
	return err

}
func (CrsIntrnship *ContestsinInternshipClass) GetContestsFromInternship (id int64)  ([]*Contest, error) {
	_, err := CrsIntrnship.internship.GetInternship(id)
	if err != nil {
		return nil,err
	}
	interns, err:= CrsIntrnship.conts.GetContestFromInternship(id)
	if err != nil {
		return nil, err
	}
	return interns, nil

}













