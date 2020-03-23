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
	"./InterviewCalendar"


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





	internshipendpoints:=Internship.NewEndpointsFactory(internshipcol)
	router.Methods("GET").Path("/internship/").HandlerFunc(internshipendpoints.GetInternships())
	router.Methods("GET").Path("/internship/{id}").HandlerFunc(internshipendpoints.GetInternship("id"))
	router.Methods("DELETE").Path("/internship/{id}").HandlerFunc(internshipendpoints.DeleteInternship("id"))
	router.Methods("PUT").Path("/internship/{id}").HandlerFunc(internshipendpoints.UpdateInternship("id"))
	router.Methods("POST").Path("/internship/").HandlerFunc(internshipendpoints.AddInternship())




	course_in_intrnshp := Courses.New_Cours_in_Internship_Collection(internshipcol, coursecol)
	courseendpoints:=Courses.NewEndpointsFactory(course_in_intrnshp)
	router.Methods("GET").Path("/course/").HandlerFunc(courseendpoints.GetCourses())
	router.Methods("GET").Path("/course/{id}").HandlerFunc(courseendpoints.GetCourse("id"))
	router.Methods("DELETE").Path("/course/{id}").HandlerFunc(courseendpoints.DeleteCourse("id"))
	router.Methods("PUT").Path("/course/{id}").HandlerFunc(courseendpoints.UpdateCourse("id"))
	router.Methods("POST").Path("/course/").HandlerFunc(courseendpoints.AddCourse())


	questionnaire,err:=Questionnaire.NewQuestionnaireCollection(conf)
	if err!=nil{
		log.Fatal(err)
	}
	questionnaireendpoints:=Questionnaire.NewEndpointsFactory(questionnaire)
	router.Methods("GET").Path("/questionnaire/").HandlerFunc(questionnaireendpoints.GetQuestionnaires())
	router.Methods("GET").Path("/questionnaire/{id}").HandlerFunc(questionnaireendpoints.GetQuestionnaire("id"))
	router.Methods("DELETE").Path("/questionnaire/{id}").HandlerFunc(questionnaireendpoints.DeleteQuestionnaire("id"))
	router.Methods("PUT").Path("/questionnaire/{id}").HandlerFunc(questionnaireendpoints.UpdateQuestionnaire("id"))
	router.Methods("POST").Path("/questionnaire/").HandlerFunc(questionnaireendpoints.AddQuestionnaire())



	contest,err:=Contest.NewContestCollection(conf)
	if err!=nil{
		log.Fatal(err)
	}
	contestendpoints:=Contest.NewEndpointsFactory(contest)
	router.Methods("GET").Path("/questionnaire/").HandlerFunc(contestendpoints.GetContests())
	router.Methods("GET").Path("/questionnaire/{id}").HandlerFunc(contestendpoints.GetContest("id"))
	router.Methods("DELETE").Path("/questionnaire/{id}").HandlerFunc(contestendpoints.DeleteContest("id"))
	router.Methods("PUT").Path("/questionnaire/{id}").HandlerFunc(contestendpoints.UpdateContest("id"))
	router.Methods("POST").Path("/questionnaire/").HandlerFunc(contestendpoints.AddContest())


	calendar,err:=InterviewCalendar.NewInterviewCalendarCollection(conf)
	if err!=nil{
		log.Fatal(err)
	}
	calendarendpoints:=InterviewCalendar.NewEndpointsFactory(calendar)
	router.Methods("GET").Path("/questionnaire/").HandlerFunc(calendarendpoints.GetInterviewCalendars())
	router.Methods("GET").Path("/questionnaire/{id}").HandlerFunc(calendarendpoints.GetInterviewCalendar("id"))
	router.Methods("DELETE").Path("/questionnaire/{id}").HandlerFunc(calendarendpoints.DeleteInterviewCalendar("id"))
	router.Methods("PUT").Path("/questionnaire/{id}").HandlerFunc(calendarendpoints.UpdateInterviewCalendar("id"))
	router.Methods("POST").Path("/questionnaire/").HandlerFunc(calendarendpoints.AddInterviewCalendar())



	intern,err:=Intern.NewInternCollection(conf)
	if err!=nil{
		log.Fatal(err)
	}
	internendpoints:=Intern.NewEndpointsFactory(intern)
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