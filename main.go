package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mukhametkaly/DAR_Internship/Account"
	"github.com/mukhametkaly/DAR_Internship/Contest"
	"github.com/mukhametkaly/DAR_Internship/Courses"
	"github.com/mukhametkaly/DAR_Internship/Intern"
	"github.com/mukhametkaly/DAR_Internship/Interview_Calendar"
	"github.com/mukhametkaly/DAR_Internship/Lecturer"
	"github.com/mukhametkaly/DAR_Internship/Questionnaire"
	"github.com/mukhametkaly/DAR_Internship/internship"
	"github.com/urfave/cli"
	"log"
	"net/http"
	"os"
	"strings"
)

//var(
//
//	conf Internship.MongoConfig
//
//    flags  = []cli.Flag{
//		&cli.StringFlag{
//			Name:        "Database",
//			Usage:       "Database name",
//			Destination: &conf.Database,
//		},
//
//		&cli.StringFlag{
//			Name:        "Host",
//			Usage:       "Database Hostname",
//			Destination: &conf.Host,
//		},
//
//		&cli.StringFlag{
//			Name:        "Port",
//			Usage:       "Database Port",
//			Destination: &conf.Port,
//		},
//	}
//
//)

func main() {

	app := cli.NewApp()
	app.Name = "Internship"
	app.Action = Start
	fmt.Println(app.Run(os.Args))
}

func Start(*cli.Context)  error{
	fs := noDirListing(http.FileServer(http.Dir("./public/static")))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	router:=mux.NewRouter()
	MongoConf:=Internship.MongoConfig{
		Host:     "localhost",
		Database: "example",
		Port:     "27017",
	}

	rConf := Account.RedisConfig{
		Addr:     "localhost:6379",
		Password: "",
		DB: 0,
	}

	RedisClient := Account.RedisConnection(rConf)
    Account.AuthRedisConnection(rConf)
	////// INTERNSHIP
	internshipcol,err:=Internship.NewInternshipCollection(MongoConf)
	if err!=nil{
		log.Fatal(err)
	}
	internshipendpoints:=Internship.NewEndpointsFactory(internshipcol)

	//////COURSES
	coursecol,err:=Courses.NewCourseCollection(MongoConf)
	if err!=nil{
		log.Fatal(err)
	}
	courseINintrnshp := Courses.NewCoursesInInternship(internshipcol, coursecol)
	courseendpoints:=Courses.NewEndpointsFactory(courseINintrnshp)

	/////QUESTIONNAIRE
	questionnaire,err:=Questionnaire.NewQuestionnaireCollection(MongoConf)
	if err!=nil{
		log.Fatal(err)
	}
	questionnaireINinternship := Questionnaire.NewQuestionnaireInInternship(internshipcol, questionnaire)
	questionnaireendpoints:=Questionnaire.NewEndpointsFactory(questionnaireINinternship)


    ///// CONTEST
	contest,err:=Contest.NewContestCollection(MongoConf)
	if err!=nil{
		log.Fatal(err)
	}
	contestToInternship := Contest.NewContestsinInternship(internshipcol, contest)
	contestendpoints:=Contest.NewEndpointsFactory(contestToInternship)

	/// CALENDAR
	calendar,err:= Interview_Calendar.NewInterviewCalendarCollection(MongoConf)
	if err!=nil{
		log.Fatal(err)
	}
	interviewCal := Interview_Calendar.NewCourseIntern(coursecol, calendar)
	calendarendpoints:= Interview_Calendar.NewEndpointsFactory(interviewCal)

	////INTERN
	intern,err:=Intern.NewInternCollection(MongoConf)
	if err!=nil{
		log.Fatal(err)
	}
	courseIntern := Intern.NewCourseIntern(coursecol, intern)
	internendpoints:=Intern.NewEndpointsFactory(courseIntern)

	///LECTURER
	lecturer, err := Lecturer.NewLecturerCollection(MongoConf)
	if (err != nil){
		log.Fatal(err)
	}
	courseLecturer := Lecturer.NewCourseLecturer(coursecol, lecturer)
	lecturerendpoints := Lecturer.NewEndpointsFactory(courseLecturer)




	router.Methods("POST").Path("/intern/").HandlerFunc(internendpoints.AddIntern())
	router.Methods("POST").Path("/intern/login").HandlerFunc(internendpoints.Authorization())
	router.Methods("POST").Path("/lecturer/login").HandlerFunc(lecturerendpoints.Authorization())
	cubrouter := router.PathPrefix("/lecturer").Subrouter()
	cubrouter.Use(Account.CustomAuth)
	cubrouter.Methods("GET").Path("/lecturer/{id}").HandlerFunc(lecturerendpoints.GetLecturer("id", RedisClient))




	router.Methods("GET").Path("/internship/").HandlerFunc(internshipendpoints.GetInternships(RedisClient))
	router.Methods("GET").Path("/internship/{id}").HandlerFunc(internshipendpoints.GetInternship("id", RedisClient))
	router.Methods("DELETE").Path("/internship/{id}").HandlerFunc(internshipendpoints.DeleteInternship("id", RedisClient))
	router.Methods("PUT").Path("/internship/{id}").HandlerFunc(internshipendpoints.UpdateInternship("id", RedisClient))
	router.Methods("POST").Path("/internship/").HandlerFunc(internshipendpoints.AddInternship(RedisClient))


	router.Methods("GET").Path("/course/").HandlerFunc(courseendpoints.GetCourses(RedisClient))
	router.Methods("GET").Path("/course/{id}").HandlerFunc(courseendpoints.GetCourse("id", RedisClient))
	router.Methods("DELETE").Path("/course/{id}").HandlerFunc(courseendpoints.DeleteCourse("id", RedisClient))
	router.Methods("PUT").Path("/course/{id}").HandlerFunc(courseendpoints.UpdateCourse("id", RedisClient))
	router.Methods("POST").Path("/course/").HandlerFunc(courseendpoints.AddCourse(RedisClient))
	router.Methods("GET").Path("/internship/{id}/courses").HandlerFunc(courseendpoints.GetCoursesFromInternship("id", RedisClient))


	router.Methods("GET").Path("/questionnaire/").HandlerFunc(questionnaireendpoints.GetQuestionnaires(RedisClient))
	router.Methods("GET").Path("/questionnaire/{id}").HandlerFunc(questionnaireendpoints.GetQuestionnaire("id"))
	router.Methods("DELETE").Path("/questionnaire/{id}").HandlerFunc(questionnaireendpoints.DeleteQuestionnaire("id", RedisClient))
	router.Methods("PUT").Path("/questionnaire/{id}").HandlerFunc(questionnaireendpoints.UpdateQuestionnaire("id", RedisClient))
	router.Methods("POST").Path("/questionnaire/").HandlerFunc(questionnaireendpoints.AddQuestionnaire(RedisClient))
	router.Methods("GET").Path("/internship/{id}/questionnaire/").HandlerFunc(questionnaireendpoints.GetQuestionnaireFromInternship("id", RedisClient))


	router.Methods("GET").Path("/contest/").HandlerFunc(contestendpoints.GetContests(RedisClient))
	router.Methods("GET").Path("/contest/{id}").HandlerFunc(contestendpoints.GetContest("id"))
	router.Methods("DELETE").Path("/contest/{id}").HandlerFunc(contestendpoints.DeleteContest("id", RedisClient))
	router.Methods("PUT").Path("/contest/{id}").HandlerFunc(contestendpoints.UpdateContest("id", RedisClient))
	router.Methods("POST").Path("/contest/").HandlerFunc(contestendpoints.AddContest(RedisClient))
	router.Methods("GET").Path("/internship/{id}/contest").HandlerFunc(contestendpoints.GetContestsFromInternship("id", RedisClient))


	router.Methods("GET").Path("/calendar/").HandlerFunc(calendarendpoints.GetInterviewCalendars(RedisClient))
	router.Methods("GET").Path("/calendar/{id}").HandlerFunc(calendarendpoints.GetInterviewCalendar("id", RedisClient))
	router.Methods("DELETE").Path("/calendar/{id}").HandlerFunc(calendarendpoints.DeleteInterviewCalendar("id", RedisClient))
	router.Methods("PUT").Path("/calendar/{id}").HandlerFunc(calendarendpoints.UpdateInterviewCalendar("id", RedisClient))
	router.Methods("POST").Path("/calendar/").HandlerFunc(calendarendpoints.AddInterviewCalendar(RedisClient))
	router.Methods("GET").Path("/course/{id}/calendar/").HandlerFunc(calendarendpoints.GetInternviewCalendarFromCourses("id", RedisClient))


	router.Methods("GET").Path("/intern/").HandlerFunc(internendpoints.GetInterns(RedisClient))
	router.Methods("GET").Path("/intern/{id}").HandlerFunc(internendpoints.GetIntern("id", RedisClient, coursecol))
	router.Methods("DELETE").Path("/intern/{id}").HandlerFunc(internendpoints.DeleteIntern("id", RedisClient))
	router.Methods("PUT").Path("/intern/{id}").HandlerFunc(internendpoints.UpdateIntern("id", RedisClient))

	router.Methods("GET").Path("/courses/{id}/interns/").HandlerFunc(internendpoints.GetInternsFromCourses("id", RedisClient))



	router.Methods("GET").Path("/lecturer/").HandlerFunc(lecturerendpoints.GetLecturers(RedisClient))
	//router.Methods("GET").Path("/lecturer/{id}").HandlerFunc(lecturerendpoints.GetLecturer("id"))
	router.Methods("DELETE").Path("/lecturer/{id}").HandlerFunc(lecturerendpoints.DeleteLecturer("id", RedisClient))
	router.Methods("PUT").Path("/lecturer/{id}").HandlerFunc(lecturerendpoints.UpdateLecturer("id", RedisClient))
	router.Methods("POST").Path("/lecturer/").HandlerFunc(lecturerendpoints.AddLecturer(RedisClient))
	router.Methods("GET").Path("/courses/{id}/lecturer/").HandlerFunc(lecturerendpoints.GetLecturerFromCourses("id", RedisClient))

	fmt.Println("Server is running")
	http.ListenAndServe(":8080",router)
	return nil


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