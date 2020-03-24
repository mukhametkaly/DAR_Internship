package main

import (
	"./internship"
	"fmt"
	"log"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
	"./Courses"
	"./Intern"
	"./Contest"
	"./Questionnaire"
	"./Interview_Calendar"


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


	questionnaireINinternship := Questionnaire.NewQuestionnaireInInternship(internshipcol, questionnaire)
	questionnaireendpoints:=Questionnaire.NewEndpointsFactory(questionnaireINinternship)
	router.Methods("GET").Path("/questionnaire/").HandlerFunc(questionnaireendpoints.GetQuestionnaires())
	router.Methods("GET").Path("/questionnaire/{id}").HandlerFunc(questionnaireendpoints.GetQuestionnaire("id"))
	router.Methods("DELETE").Path("/questionnaire/{id}").HandlerFunc(questionnaireendpoints.DeleteQuestionnaire("id"))
	router.Methods("PUT").Path("/questionnaire/{id}").HandlerFunc(questionnaireendpoints.UpdateQuestionnaire("id"))
	router.Methods("POST").Path("/questionnaire/").HandlerFunc(questionnaireendpoints.AddQuestionnaire())



	contest,err:=Contest.NewContestCollection(conf)
	if err!=nil{
		log.Fatal(err)
	}
	contestToInternship := Contest.NewCoursesinInternship(internshipcol, contest)
	contestendpoints:=Contest.NewEndpointsFactory(contestToInternship)
	router.Methods("GET").Path("/contest/").HandlerFunc(contestendpoints.GetContests())
	router.Methods("GET").Path("/contest/{id}").HandlerFunc(contestendpoints.GetContest("id"))
	router.Methods("DELETE").Path("/contest/{id}").HandlerFunc(contestendpoints.DeleteContest("id"))
	router.Methods("PUT").Path("/contest/{id}").HandlerFunc(contestendpoints.UpdateContest("id"))
	router.Methods("POST").Path("/contest/").HandlerFunc(contestendpoints.AddContest())


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



	intern,err:=Intern.NewInternCollection(conf)
	if err!=nil{
		log.Fatal(err)
	}
	courseIntern := Intern.NewCourseIntern(coursecol, intern)
	internendpoints:=Intern.NewEndpointsFactory(courseIntern)
	router.Methods("GET").Path("/questionnaire/").HandlerFunc(internendpoints.GetInterns())
	router.Methods("GET").Path("/questionnaire/{id}").HandlerFunc(internendpoints.GetIntern("id"))
	router.Methods("DELETE").Path("/questionnaire/{id}").HandlerFunc(internendpoints.DeleteIntern("id"))
	router.Methods("PUT").Path("/questionnaire/{id}").HandlerFunc(internendpoints.UpdateIntern("id"))
	router.Methods("POST").Path("/questionnaire/").HandlerFunc(internendpoints.AddIntern())





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