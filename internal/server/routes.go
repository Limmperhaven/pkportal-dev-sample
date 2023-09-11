package server

import (
	"github.com/Limmperhaven/pkportal-be-v2/internal/controllers"
	"github.com/Limmperhaven/pkportal-be-v2/internal/controllers/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"net/http"
)

func initRoutes(router *gin.Engine, c *controllers.ControllerStorage, m *middlewares.MiddlewareStorage) {
	authorizedNotActivated := router.Group("/")
	authorizedNotActivated.Use(m.AuthMiddleware)
	authorized := authorizedNotActivated.Group("/")
	authorized.Use(m.CheckActivationMiddleware)
	admin := authorizedNotActivated.Group("/")
	admin.Use(m.CheckAdminRoleMiddleware)

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", c.SignUp)
		auth.POST("/sign-in", c.SignIn)
		auth.POST("/recover/:email", c.RecoverPassword)
		auth.GET("/activate/:token", c.ActivateAccount)
		auth.POST("/confirmRecover/:token", c.ConfirmRecover)
	}

	authTd := authorized.Group("/td")
	adminTd := admin.Group("/td")
	{
		adminTd.POST("/create", c.CreateTestDate)
		authTd.GET("/byId/:id", c.GetTestDate)
		adminTd.PUT("/updMaxPersons/:id/:maxPersons", c.UpdateTestDateMaxPersons)
		adminTd.POST("/setStatus/:id/:status", c.SetTestDatePubStatus)
		adminTd.POST("/list", c.ListTestDates)
		authTd.GET("/listAvailable", c.ListAvailableTestDates)
		adminTd.POST("/signUpUser/:userId/:tdId", c.SignUpUserToTestDate)
		authTd.POST("/signUpMe/:tdId", c.SignUpMeToTestDate)
		adminTd.GET("/listCommonLocations", c.ListCommonLocations)
		adminTd.POST("/setAttendance/:userId/:tdId/:attendance", c.SetTestDateAttended)
		adminTd.GET("/regList/:tdId", c.DownloadRegistrationList)
		adminTd.GET("/export/:tdId", c.ExportToXlsx)
	}

	authUser := authorized.Group("/user")
	adminUser := admin.Group("/user")
	aNAUser := authorizedNotActivated.Group("/user")
	{
		adminUser.POST("/create", c.CreateUser)
		adminUser.GET("/byId/:id", c.GetUser)
		adminUser.PUT("/byId/:id", c.UpdateUser)
		adminUser.POST("/list", c.ListUsers)
		adminUser.POST("/setStatus/:userId/:statusId", c.SetUserStatus)
		adminUser.GET("/downloadScreenshot/:userId", c.DownloadScreenshot)
		adminUser.POST("/setUserRole/:userId/:role", c.SetUserRole)
		aNAUser.GET("/me", c.GetMe)
		authUser.POST("/listStatuses", c.ListStatuses)
		authUser.POST("/uploadScreenshot", c.UploadScreenshot)
		authUser.GET("/downloadMyScreenshot", c.DownloadMyScreenshot)
		aNAUser.POST("/resendActivation", c.ResendActivationEmail)
	}

	authProfile := authorized.Group("/profiles")
	adminProfile := admin.Group("/profiles")
	{
		adminProfile.POST("/create", c.CreateProfile)
		authProfile.GET("/byId/:id", c.CreateProfile)
		adminProfile.PUT("/byId/:id", c.UpdateProfile)
		authProfile.GET("/list", c.ListProfiles)
		adminProfile.POST("/setToUser/:userId", c.SetProfilesToUser)
		authProfile.POST("/setToMe", c.SetProfilesToMe)
	}

	authSubject := authorized.Group("/subjects")
	adminSubject := admin.Group("/subjects")
	{
		adminSubject.POST("/create", c.CreateSubject)
		authSubject.GET("/byId/:id", c.GetSubject)
		adminSubject.PUT("/byId/:id", c.UpdateSubject)
		authSubject.POST("/list", c.ListSubjects)
		authSubject.GET("/listFL", c.ListForeignLanguages)
		adminSubject.POST("/setToUser/:userId", c.SetSubjectToUser)
		authSubject.POST("/setToMe", c.SetSubjectToMe)
	}

	authFL := authorized.Group("/fl")
	adminFL := admin.Group("/fl")
	{
		adminFL.POST("/create", c.CreateForeignLanguage)
		authFL.GET("/byId/:id", c.GetForeignLanguage)
		adminFL.POST("/byId/:id", c.UpdateForeignLanguage)
		authFL.GET("/list", c.ListForeignLanguages)
		adminFL.POST("/setToUser/:userId/:flId", c.SetForeignLanguageToUser)
		authFL.POST("/setToMe/:flId", c.SetForeignLanguageToMe)
	}

	adminExams := admin.Group("exams")
	{
		adminExams.POST("/setGrades", c.SetGrades)
	}

	adminNotifications := admin.Group("notifications")
	{
		adminNotifications.POST("/create", c.CreateNotification)
	}

}

func initCors(router *gin.Engine) http.Handler {
	c := cors.New(cors.Options{
		AllowOriginFunc:        func(origin string) bool { return true },
		AllowOriginRequestFunc: func(r *http.Request, origin string) bool { return true },
		AllowedMethods: []string{
			http.MethodGet, http.MethodPost, http.MethodPut,
		},
		AllowedHeaders:      []string{"accept", "authorization", "content-type"},
		ExposedHeaders:      []string{"Set-Cookie", "authorization", "Content-Disposition"},
		AllowCredentials:    true,
		AllowPrivateNetwork: true,
		OptionsPassthrough:  false,
		Debug:               false,
	})
	return c.Handler(router)
}
