package main

import (
	"Internship/internship"
	"fmt"
	"log"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
	"Internship/Courses"
	"Internship/Lecturer"
	"Internship/Intern"
	"Internship/Contest"
	"Internship/Questionnaire"
	"Internship/Interview_Calendar"

)

func main() {
	fs := noDirListing(http.FileServer(http.Dir("./public/static")))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	router:=mux.NewRouter()
	conf:=Internship.MongoConfig{
		Host:     "localhost",
		Database: "example",
		Port:     "27017",
	}

	////// INTERNSHIP
	internshipcol,err:=Internship.NewInternshipCollection(conf)
	if err!=nil{
		log.Fatal(err)
	}
	coursecol,err:=Courses.NewCourseCollection(conf)
	if err!=nil{
		log.Fatal(err)
	}
	questionnaire,err:=Questionnaire.NewQuestionnaireCollection(conf)
	if err!=nil{
		log.Fatal(err)
	}





	internshipendpoints:=Internship.NewEndpointsFactory(internshipcol)
	router.Methods("GET").Path("/internship/").HandlerFunc(internshipendpoints.GetInternships())
	router.Methods("GET").Path("/internship/{id}").HandlerFunc(internshipendpoints.GetInternship("id"))
	router.Methods("DELETE").Path("/internship/{id}").HandlerFunc(internshipendpoints.DeleteInternship("id"))
	router.Methods("PUT").Path("/internship/{id}").HandlerFunc(internshipendpoints.UpdateInternship("id"))
	router.Methods("POST").Path("/internship/").HandlerFunc(internshipendpoints.AddInternship())




	courseINintrnshp := Courses.NewCoursesInInternship(internshipcol, coursecol)
	courseendpoints:=Courses.NewEndpointsFactory(courseINintrnshp)
	router.Methods("GET").Path("/course/").HandlerFunc(courseendpoints.GetCourses())
	router.Methods("GET").Path("/course/{id}").HandlerFunc(courseendpoints.GetCourse("id"))
	router.Methods("DELETE").Path("/course/{id}").HandlerFunc(courseendpoints.DeleteCourse("id"))
	router.Methods("PUT").Path("/course/{id}").HandlerFunc(courseendpoints.UpdateCourse("id"))
	router.Methods("POST").Path("/course/").HandlerFunc(courseendpoints.AddCourse())
	router.Methods("GET").Path("/internship/{id}/courses").HandlerFunc(courseendpoints.GetCoursesFromInternship("id"))


	questionnaireINinternship := Questionnaire.NewQuestionnaireInInternship(internshipcol, questionnaire)
	questionnaireendpoints:=Questionnaire.NewEndpointsFactory(questionnaireINinternship)
	router.Methods("GET").Path("/questionnaire/").HandlerFunc(questionnaireendpoints.GetQuestionnaires())
	router.Methods("GET").Path("/questionnaire/{id}").HandlerFunc(questionnaireendpoints.GetQuestionnaire("id"))
	router.Methods("DELETE").Path("/questionnaire/{id}").HandlerFunc(questionnaireendpoints.DeleteQuestionnaire("id"))
	router.Methods("PUT").Path("/questionnaire/{id}").HandlerFunc(questionnaireendpoints.UpdateQuestionnaire("id"))
	router.Methods("POST").Path("/questionnaire/").HandlerFunc(questionnaireendpoints.AddQuestionnaire())
	router.Methods("GET").Path("/internship/{id}/questionnaire/").HandlerFunc(questionnaireendpoints.GetQuestionnaireFromInternship("id"))



	contest,err:=Contest.NewContestCollection(conf)
	if err!=nil{
		log.Fatal(err)
	}
	contestToInternship := Contest.NewContestsinInternship(internshipcol, contest)
	contestendpoints:=Contest.NewEndpointsFactory(contestToInternship)
	router.Methods("GET").Path("/contest/").HandlerFunc(contestendpoints.GetContests())
	router.Methods("GET").Path("/contest/{id}").HandlerFunc(contestendpoints.GetContest("id"))
	router.Methods("DELETE").Path("/contest/{id}").HandlerFunc(contestendpoints.DeleteContest("id"))
	router.Methods("PUT").Path("/contest/{id}").HandlerFunc(contestendpoints.UpdateContest("id"))
	router.Methods("POST").Path("/contest/").HandlerFunc(contestendpoints.AddContest())
	router.Methods("GET").Path("/internship/{id}/contest").HandlerFunc(contestendpoints.GetContestsFromInternship("id"))


	calendar,err:= Interview_Calendar.NewInterviewCalendarCollection(conf)
	if err!=nil{
		log.Fatal(err)
	}
	interviewCal := Interview_Calendar.NewCourseIntern(coursecol, calendar)
	calendarendpoints:= Interview_Calendar.NewEndpointsFactory(interviewCal)
	router.Methods("GET").Path("/calendar/").HandlerFunc(calendarendpoints.GetInterviewCalendars())
	router.Methods("GET").Path("/calendar/{id}").HandlerFunc(calendarendpoints.GetInterviewCalendar("id"))
	router.Methods("DELETE").Path("/calendar/{id}").HandlerFunc(calendarendpoints.DeleteInterviewCalendar("id"))
	router.Methods("PUT").Path("/calendar/{id}").HandlerFunc(calendarendpoints.UpdateInterviewCalendar("id"))
	router.Methods("POST").Path("/calendar/").HandlerFunc(calendarendpoints.AddInterviewCalendar())
	router.Methods("GET").Path("/course/{id}/calendar/").HandlerFunc(calendarendpoints.GetInternviewCalendarFromCourses("id"))




	intern,err:=Intern.NewInternCollection(conf)
	if err!=nil{
		log.Fatal(err)
	}
	courseIntern := Intern.NewCourseIntern(coursecol, intern)
	internendpoints:=Intern.NewEndpointsFactory(courseIntern)
	router.Methods("GET").Path("/intern/").HandlerFunc(internendpoints.GetInterns())
	router.Methods("GET").Path("/intern/{id}").HandlerFunc(internendpoints.GetIntern("id"))
	router.Methods("DELETE").Path("/intern/{id}").HandlerFunc(internendpoints.DeleteIntern("id"))
	router.Methods("PUT").Path("/intern/{id}").HandlerFunc(internendpoints.UpdateIntern("id"))
	router.Methods("POST").Path("/intern/").HandlerFunc(internendpoints.AddIntern())
	router.Methods("GET").Path("/courses/{id}/interns/").HandlerFunc(internendpoints.GetInternsFromCourses("id"))


    lecturer, err := Lecturer.NewLecturerCollection(conf)
    if (err != nil){
    	log.Fatal(err)
	}
    courseLecturer := Lecturer.NewCourseLecturer(coursecol, lecturer)
    lecturerendpoints := Lecturer.NewEndpointsFactory(courseLecturer)
	router.Methods("GET").Path("/lecturer/").HandlerFunc(lecturerendpoints.GetLecturers())
	router.Methods("GET").Path("/lecturer/{id}").HandlerFunc(lecturerendpoints.GetLecturer("id"))
	router.Methods("DELETE").Path("/lecturer/{id}").HandlerFunc(lecturerendpoints.DeleteLecturer("id"))
	router.Methods("PUT").Path("/lecturer/{id}").HandlerFunc(lecturerendpoints.UpdateLecturer("id"))
	router.Methods("POST").Path("/lecturer/").HandlerFunc(lecturerendpoints.AddLecturer())
	router.Methods("GET").Path("/courses/{id}/lecturer/").HandlerFunc(lecturerendpoints.GetLecturerFromCourses("id"))

	fmt.Println("Server is running")
	http.ListenAndServe(":8080",router)

}


func noDirListing(h http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") || r.URL.Path == "" {
			http.NotFound(w, r)
			return
		}
		h.ServeHTTP(w, r)
	})
}