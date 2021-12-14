package impl

import (
	"fmt"
	"svc-mahasiswa/models"
	repo "svc-mahasiswa/repository"

	"github.com/jmoiron/sqlx"
)

type mySQLStudent struct {
	Conn *sqlx.DB
}

func NewSQLStudent(Conn *sqlx.DB) repo.Student {
	return &mySQLStudent{
		Conn: Conn,
	}
}
func (m *mySQLStudent) GetStudent(id int)([]models.Student, error) {

	var arrStudent []models.Student
	query := fmt.Sprintf("SELECT name, npm, address, city, gander, religion, phone FROM mst_mahasiswa WHERE id=%d", id)
	row, err := m.Conn.Queryx(query)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	for row.Next(){
		var student models.Student
		err := row.StructScan(&student)
		if err != nil {
			return nil, err
		}
		arrStudent = append(arrStudent, student)
	}

	return arrStudent, nil
}

func (m *mySQLStudent) GetStudents()([]models.Student, error) {

	var arrStudent []models.Student
	row, err := m.Conn.Queryx(`SELECT name, npm, address, city, gander, religion, phone FROM mst_mahasiswa`)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	for row.Next(){
		var student models.Student
		err := row.StructScan(&student)
		if err != nil {
			return nil, err
		}
		arrStudent = append(arrStudent, student)
	}

	return arrStudent, nil
}

func (m *mySQLStudent) AddStudent(student models.Student)(models.Student, error) {
	stdn, err := m.Conn.NamedQuery(`INSERT INTO mst_mahasiswa (name, npm, address, city, gander, religion, phone) VALUES (:name, :npm, :address, :city, :gander, :religion, :phone) RETURNING name, npm, address, city, gander, religion, phone`, map[string]interface{}{
		"name" : student.Name,
		"npm" : student.Npm,
		"address" : student.Address,
		"city" : student.City,
		"gander" : student.Gander,
		"religion" : student.Gander,
		"phone" : student.Phone,
	})

	if err != nil {
		return models.Student{}, err
	}

	defer stdn.Close()
	for stdn.Next(){
		stdn.StructScan(&student)
	}

	return student, nil
}
func (m *mySQLStudent) UpdateStudent(id int, student models.Student)(models.Student, error){
	stdn, err := m.Conn.NamedQuery(`UPDATE public.mst_mahasiswa SET name=:name, npm=:npm, address=:address, city=:city, gander=:gender, religion=:religion, phone=:phone WHERE id=:id`, map[string]interface{}{
		"id" : id,
		"name" : student.Name,
		"npm" : student.Npm,
		"address" : student.Address,
		"city" : student.City,
		"gender" : student.Gander,
		"religion" : student.Gander,
		"phone" : student.Phone,
	})

	if err != nil {
		return models.Student{}, err
	}

	defer stdn.Close()
	for stdn.Next(){
		stdn.StructScan(&student)
	}

	return student, nil
}

func (m *mySQLStudent) DeleteStudent(id int) error {
	query := fmt.Sprintf("DELETE FROM mst_mahasiswa WHERE id= %d", id)
	_, err := m.Conn.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

