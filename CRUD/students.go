package crud

//Students Connection
func NewStudents() *Repo {
	return &Repo{}
}

// Repo ..
type Repo struct {
}

/*
func (repo *Repo) GetStudentsInfo(ctx context.Context, rb_number string) (*model.StudentInfo, error) {

	rows, err := repo.DBConn.Query("SELECT week_day, start, \"end\" FROM public.timetable where org_id= $1 ORDER BY week_day", rb_number)

	stdInfo := new(model.StudentInfo)
	if err != nil {
		panic(err)
		return stdInfo, err
	}

	if rows.Next() {

	} else {
		return stdInfo, nil
	}

	return stdInfo, nil
}
func (repo *Repo) GetStudentsAssessment(ctx context.Context, rb_number string) (*model.StudentAssessment, error) {

	rows, err := repo.DBConn.Query("SELECT week_day, start, \"end\" FROM public.timetable where org_id= $1 ORDER BY week_day", rb_number)

	stdInfo := new(model.StudentAssessment)
	if err != nil {
		panic(err)
		return stdInfo, err
	}

	if rows.Next() {

	} else {
		return stdInfo, nil
	}

	return stdInfo, nil
}
*/
