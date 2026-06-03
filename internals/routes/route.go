package routes

import (
	"net/http"

	"UniCore/internals/handlers"
	"UniCore/internals/middlewares"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// Public routes
	router.HandleFunc("/api/auth/register", handlers.Register).Methods("POST")
	router.HandleFunc("/api/auth/login", handlers.Login).Methods("POST")
	router.HandleFunc("/api/upload", handlers.UploadFile).Methods("POST")

	// Serve uploaded files
	router.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	// Protected routes
	protected := router.PathPrefix("/api").Subrouter()
	protected.Use(middlewares.AuthMiddleware)

	protected.HandleFunc("/auth/me", handlers.GetMe).Methods("GET")

	// Cafeteria
	protected.HandleFunc("/cafeteria", handlers.CreateCafeSchedule).Methods("POST")
	protected.HandleFunc("/cafeteria", handlers.GetCafeSchedule).Methods("GET")
	protected.HandleFunc("/cafeteria/{id}", handlers.GetSingleSchedule).Methods("GET")
	protected.HandleFunc("/cafeteria/{id}", handlers.UpdateCafeSchedule).Methods("PUT")
	protected.HandleFunc("/cafeteria/{id}", handlers.DeleteCafeSchedule).Methods("DELETE")

	// Colleges
	protected.HandleFunc("/colleges", handlers.CreateCollage).Methods("POST")
	protected.HandleFunc("/colleges", handlers.GetCollage).Methods("GET")
	protected.HandleFunc("/colleges/{id}", handlers.GetSingleCollage).Methods("GET")
	protected.HandleFunc("/colleges/{id}", handlers.UpdateCollage).Methods("PUT")
	protected.HandleFunc("/colleges/{id}", handlers.DeleteCollage).Methods("DELETE")

	// Courses
	protected.HandleFunc("/courses", handlers.CreateCourse).Methods("POST")
	protected.HandleFunc("/courses", handlers.GetCourse).Methods("GET")
	protected.HandleFunc("/courses/{id}", handlers.GetSingleCourse).Methods("GET")
	protected.HandleFunc("/courses/{id}", handlers.UpdateCourse).Methods("PUT")
	protected.HandleFunc("/courses/{id}", handlers.DeleteCourse).Methods("DELETE")

	// Curriculums
	protected.HandleFunc("/curriculums", handlers.CreateCurriculum).Methods("POST")
	protected.HandleFunc("/curriculums", handlers.GetCurriculum).Methods("GET")
	protected.HandleFunc("/curriculums/{id}", handlers.GetSingleCurriculum).Methods("GET")
	protected.HandleFunc("/curriculums/{id}", handlers.UpdateCurriculum).Methods("PUT")
	protected.HandleFunc("/curriculums/{id}", handlers.DeleteCurriculum).Methods("DELETE")

	// Departments
	protected.HandleFunc("/departments", handlers.CreateDepartment).Methods("POST")
	protected.HandleFunc("/departments", handlers.GetDepartment).Methods("GET")
	protected.HandleFunc("/departments/{id}", handlers.GetSingleDepartment).Methods("GET")
	protected.HandleFunc("/departments/{id}", handlers.UpdateDepartment).Methods("PUT")
	protected.HandleFunc("/departments/{id}", handlers.DeleteDepartment).Methods("DELETE")

	// Events
	protected.HandleFunc("/events", handlers.CreateEvent).Methods("POST")
	protected.HandleFunc("/events", handlers.GetEvent).Methods("GET")
	protected.HandleFunc("/events/{id}", handlers.GetSingleEvent).Methods("GET")
	protected.HandleFunc("/events/{id}", handlers.UpdateEvent).Methods("PUT")
	protected.HandleFunc("/events/{id}", handlers.DeleteEvent).Methods("DELETE")

	// Galleries
	protected.HandleFunc("/galleries", handlers.CreateGallery).Methods("POST")
	protected.HandleFunc("/galleries", handlers.GetGallery).Methods("GET")
	protected.HandleFunc("/galleries/{id}", handlers.GetSingleGallery).Methods("GET")
	protected.HandleFunc("/galleries/{id}", handlers.UpdateGallery).Methods("PUT")
	protected.HandleFunc("/galleries/{id}", handlers.DeleteGallery).Methods("DELETE")

	// Lounges
	protected.HandleFunc("/lounges", handlers.CreateLounge).Methods("POST")
	protected.HandleFunc("/lounges", handlers.GetLounge).Methods("GET")
	protected.HandleFunc("/lounges/{id}", handlers.GetSingleLounge).Methods("GET")
	protected.HandleFunc("/lounges/{id}", handlers.UpdateLounge).Methods("PUT")
	protected.HandleFunc("/lounges/{id}", handlers.DeleteLounge).Methods("DELETE")

	// Profiles
	protected.HandleFunc("/profiles", handlers.CreateProfile).Methods("POST")
	protected.HandleFunc("/profiles/{id}", handlers.GetSingleProfile).Methods("GET")
	protected.HandleFunc("/profiles/{id}", handlers.UpdateProfile).Methods("PUT")

	// Recent Activities
	protected.HandleFunc("/recent-activities", handlers.CreateRecent).Methods("POST")
	protected.HandleFunc("/recent-activities", handlers.GetRecentActivities).Methods("GET")
	protected.HandleFunc("/recent-activities/{id}", handlers.GetSingleRecentActivity).Methods("GET")

	// Roles
	protected.HandleFunc("/roles", handlers.CreateRole).Methods("POST")
	protected.HandleFunc("/roles", handlers.GetRole).Methods("GET")
	protected.HandleFunc("/roles/{id}", handlers.UpdateRole).Methods("PUT")
	protected.HandleFunc("/roles/{id}", handlers.DeleteRole).Methods("DELETE")

	// Semesters
	protected.HandleFunc("/semesters", handlers.CreateSemester).Methods("POST")
	protected.HandleFunc("/semesters", handlers.GetSemester).Methods("GET")
	protected.HandleFunc("/semesters/{id}", handlers.GetSingleSemester).Methods("GET")
	protected.HandleFunc("/semesters/{id}", handlers.UpdateSemester).Methods("PUT")
	protected.HandleFunc("/semesters/{id}", handlers.DeleteSemester).Methods("DELETE")

	// Social Links
	protected.HandleFunc("/social-links", handlers.CreateSocialLinks).Methods("POST")
	protected.HandleFunc("/social-links", handlers.GetSocialLinks).Methods("GET")
	protected.HandleFunc("/social-links/{id}", handlers.GetSingleSocailLink).Methods("GET")
	protected.HandleFunc("/social-links/{id}", handlers.UpdateSocialLink).Methods("PUT")
	protected.HandleFunc("/social-links/{id}", handlers.DeleteSocailLink).Methods("DELETE")

	// Student Types
	protected.HandleFunc("/student-types", handlers.CreateStudentType).Methods("POST")
	protected.HandleFunc("/student-types", handlers.GetStudentTypes).Methods("GET")
	protected.HandleFunc("/student-types/{id}", handlers.GetSingleStudentType).Methods("GET")
	protected.HandleFunc("/student-types/{id}", handlers.UpdateStudentType).Methods("PUT")
	protected.HandleFunc("/student-types/{id}", handlers.DeleteStudentType).Methods("DELETE")

	// Universities
	protected.HandleFunc("/universities", handlers.CreateUniversity).Methods("POST")
	protected.HandleFunc("/universities", handlers.GetUniversity).Methods("GET")
	protected.HandleFunc("/universities/{id}", handlers.GetSingleUniversity).Methods("GET")
	protected.HandleFunc("/universities/{id}", handlers.UpdateUniversity).Methods("PUT")
	protected.HandleFunc("/universities/{id}", handlers.DeleteUniversity).Methods("DELETE")

	return router
}
