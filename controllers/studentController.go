package controllers

import (
	"log"
	"strconv"
	"svc-mahasiswa/config"
	"svc-mahasiswa/models"
	"svc-mahasiswa/repository"
	"svc-mahasiswa/repository/impl"

	"github.com/gin-gonic/gin"
)

type Student struct {
	StudentRepo repository.Student
}

func NewDriverController(db *config.DB) Student{
	return Student{
		StudentRepo: impl.NewSQLStudent(db.SQL),
	}
}

func (student *Student)Route(app *gin.Engine){
	app.GET("/api/students/:id", student.Get)
	app.GET("/api/students", student.List)
	app.POST("/api/students", student.Add)
	app.PUT("/api/students/:id", student.Edit)
	app.DELETE("/api/students/:id", student.Delete)
}

func (repo *Student) Get(c *gin.Context){
	idStdn := c.Param("id")
	id, _ := strconv.Atoi(idStdn)
	payload, err := repo.StudentRepo.GetStudent(id)
	if err != nil {
		panic(err)
	}

	c.JSON(200, models.WebResponse{
		Code: 200,
		Status: "OK",
		Data: payload,
	})
}

func (repo *Student) List(c *gin.Context){
	payload, err := repo.StudentRepo.GetStudents()
	if err != nil {
		panic(err)
	}

	c.JSON(200, models.WebResponse{
		Code: 200,
		Status: "OK",
		Data: payload,
	})
}

func (repo *Student) Add(c *gin.Context) {
	var student models.Student
	errBind := c.BindJSON(&student); if errBind != nil {
		c.JSON(400, models.WebResponse{
			Code: 400,
			Status: "Parameter is Required",
		})
	}
	stdn, err := repo.StudentRepo.AddStudent(student)
	 if err != nil {
		 log.Panic(err)
	 }

	 c.JSON(200, models.WebResponse{
		 Code: 200,
		 Status: "OK",
		 Data: stdn,
	 })
}

func (repo *Student) Edit(c *gin.Context) {
	var student models.Student
	idStdn := c.Param("id")
	id, _ := strconv.Atoi(idStdn)
	errBind := c.BindJSON(&student); if errBind != nil {
		c.JSON(400, models.WebResponse{
			Code: 400,
			Status: "Parameter is Required",
		})
	}
	stdn, err := repo.StudentRepo.UpdateStudent(id, student)
	 if err != nil {
		 log.Panic(err)
	 }

	 c.JSON(200, models.WebResponse{
		 Code: 200,
		 Status: "OK",
		 Data: stdn,
	 })
}

func (repo *Student) Delete(c *gin.Context) {
	// var student models.Student
	idStdn := c.Param("id")
	id, _ := strconv.Atoi(idStdn)
	
	 err := repo.StudentRepo.DeleteStudent(id)
	 if err != nil {
		 log.Panic(err)
	 }

	 c.JSON(200, models.WebResponse{
		 Code: 200,
		 Status: "OK",
	 })
}