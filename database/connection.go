package database

import (
	"database/sql"
	"fmt"
	"log"
	"pds-listener-service/types"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/mattn/go-sqlite3"
)

func ConnectSQLite(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "C:\\Users\\lara\\Desktop\\pds-service\\records\\" + path)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to SQLite database: %v", err)
	}
	fmt.Println("Connected to SQLite database")
	return db, nil
}

func ConnectSQLServer(server, username, password, port, database string) (*sql.DB, error) {
	var db *sql.DB
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;encrypt=disable;",
		server, username, password, port, database)

	var err error
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to SQL Server database")
	return db, nil
}

func CreateUser(db *sql.DB, eu types.EmployeeUser) error {
	query := `INSERT INTO employees (
		employee_id, lastname, firstname, middlename, extension, date_birth, place_birth, sex, civil_status, 
		civil_status_others, height, weight, blood_type, gsis_id_no, gsis_policy_no, gsis_bp_no, pag_ibig_no, 
		philhealth_no, sss_no, tin_no, lbp_account_no, dbp_account_no, agency_employee_no, citizenship, 
		citizenship_by, indicate_country, residential_house_no, residential_street, residential_village, 
		residential_barangay_code, residential_city_code, residential_province_string, residential_barangay_string, 
		residential_city_string, residential_province_code, residential_zip_code, permanent_house_no, permanent_street, 
		permanent_village, permanent_barangay_string, permanent_city_string, permanent_province_string, 
		permanent_barangay_code, permanent_city_code, permanent_province_code, permanent_zip_code, telephone_no, 
		mobile_no, email_address, employee_status, first_day_of_service, active_status, office, position, salary_grade, 
		step, salary_rate, user_id, person_notified_name, person_notified_contact, person_notified_address, address, 
		deleted_at, created_at, updated_at
	) VALUES (@p1, @p2, @p3, @p4, @p5, @p6, @p7, @p8, @p9, @p10, @p11, @p12, @p13, @p14, @p15, @p16, @p17, @p18, @p19, @p20, @p21, @p22, @p23, @p24, @p25, @p26, @p27, @p28, @p29, @p30, @p31, @p32, @p33, @p34, @p35, @p36, @p37, @p38, @p39, @p40, @p41, @p42, @p43, @p44, @p45, @p46, @p47, @p48, @p49, @p50, @p51, @p52, @p53, @p54, @p55, @p56, @p57, @p58, @p59, @p60, @p61, @p62, @p63, @p64, @p65)
`
	
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		eu.EmployeeID, eu.Lastname, eu.Firstname, eu.Middlename, eu.Extension, eu.DateBirth, eu.PlaceBirth, eu.Sex, eu.CivilStatus,
		eu.CivilStatusOthers, eu.Height, eu.Weight, eu.BloodType, eu.GsisIDNo, eu.GsisPolicyNo, eu.GsisBPNo, eu.PagIbigNo, eu.PhilhealthNo, eu.SssNo, eu.TinNo, eu.LbpAccountNo, eu.DbpAccountNo, eu.AgencyEmployeeNo, eu.Citizenship, eu.CitizenshipBy, eu.IndicateCountry, eu.ResidentialHouseNo, eu.ResidentialStreet, eu.ResidentialVillage, eu.ResidentialBarangayCode, eu.ResidentialCityCode, eu.ResidentialProvinceString, eu.ResidentialBarangayString, eu.ResidentialCityString, eu.ResidentialProvinceCode, eu.ResidentialZipCode, eu.PermanentHouseNo, eu.PermanentStreet, eu.PermanentVillage, eu.PermanentBarangayString, eu.PermanentCityString, eu.PermanentProvinceString, eu.PermanentBarangayCode, eu.PermanentCityCode, eu.PermanentProvinceCode, eu.PermanentZipCode, eu.TelephoneNo, eu.MobileNo, eu.EmailAddress, eu.EmployeeStatus, eu.FirstDayOfService, eu.ActiveStatus, eu.Office, eu.Position, eu.SalaryGrade, eu.Step, eu.SalaryRate, eu.UserID, eu.PersonNotifiedName, eu.PersonNotifiedContact, eu.PersonNotifiedAddress, eu.Address, eu.DeletedAt, eu.CreatedAt, eu.UpdatedAt,
	)

	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("Created user")
	return nil
}


func CreateFamilyInfo(db *sql.DB, fi types.FamilyInfo) error {
	query := `INSERT INTO employee_family_backgrounds (employee_id,
		spouse_firstname,
		spouse_lastname,
		spouse_middlename,
		spouse_extension,
		spouse_occupation,
		spouse_employer_business_name,
		spouse_business_address,
		spouse_telephone_number,
		father_firstname,
		father_lastname,
		father_middlename,
		father_extension,
		mother_lastname,
		mother_firstname,
		mother_middlename,
		user_id,
		created_at,
		updated_at
	) VALUES (@p1, @p2, @p3, @p4, @p5, @p6, @p7, @p8, @p9, @p10, @p11, @p12, @p13, @p14, @p15, @p16, @p17, @p18, @p19)`

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		fi.EmployeeID, fi.SpouseFirstname, fi.SpouseLastname, fi.SpouseMiddlename, fi.SpouseExtension, fi.SpouseOccupation, fi.SpouseEmployerBusiness, fi.SpouseBusinessAddress, fi.SpouseTelephoneNumber, fi.FatherFirstname, fi.FatherLastname, fi.FatherMiddlename, fi.FatherExtension, fi.MotherLastname, fi.MotherFirstname, fi.MotherMiddlename, fi.UserID, fi.CreatedAt, fi.UpdatedAt,
	)

	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("Created family info")
	return nil
}

func CreateEducationalBackground(db *sql.DB, eb types.EducationalBackground) error {
	query := `INSERT INTO employee_educational_backgrounds (
		employee_id
		,elementary_name
		,elementary_education
		,elementary_period_from
		,elementary_period_to
		,elementary_highest_level_units_earned
		,elementary_year_graduated
		,elementary_scholarship
		,secondary_name
		,secondary_education
		,secondary_period_from
		,secondary_period_to
		,secondary_highest_level_units_earned
		,secondary_year_graduated
		,secondary_scholarship
		,vocational_trade_course_name
		,vocational_education
		,vocational_trade_course_period_from
		,vocational_trade_course_period_to
		,vocational_trade_course_highest_level_units_earned
		,vocational_trade_course_year_graduated
		,vocational_trade_course_scholarship
		,college_name
		,college_education
		,college_period_from
		,college_period_to
		,college_highest_level_units_earned
		,college_year_graduated
		,college_scholarship
		,graduate_studies_name
		,graduate_studies_education
		,graduate_studies_period_from
		,graduate_studies_period_to
		,graduate_studies_highest_level_units_earned
		,graduate_studies_year_graduated
		,graduate_studies_scholarship
		,user_id
		,created_at
		,updated_at
		) VALUES (@p1, @p2, @p3, @p4, @p5, @p6, @p7, @p8, @p9, @p10, @p11, @p12, @p13, @p14, @p15, @p16, @p17, @p18, @p19, @p20, @p21, @p22, @p23, @p24, @p25, @p26, @p27, @p28, @p29, @p30, @p31, @p32, @p33, @p34, @p35, @p36, @p37, @p38, @p39)`

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		eb.EmployeeID, eb.ElementaryName, eb.ElementaryEducation, eb.ElementaryPeriodFrom, eb.ElementaryPeriodTo, eb.ElementaryHighestLevelUnitsEarned, eb.ElementaryYearGraduated, eb.ElementaryScholarship, eb.SecondaryName, eb.SecondaryEducation, eb.SecondaryPeriodFrom, eb.SecondaryPeriodTo, eb.SecondaryHighestLevelUnitsEarned, eb.SecondaryYearGraduated, eb.SecondaryScholarship, eb.VocationalTradeCourseName, eb.VocationalEducation, eb.VocationalTradeCoursePeriodFrom, eb.VocationalTradeCoursePeriodTo, eb.VocationalTradeCourseHighestLevelUnitsEarned, eb.VocationalTradeCourseYearGraduated, eb.VocationalTradeCourseScholarship, eb.CollegeName, eb.CollegeEducation, eb.CollegePeriodFrom, eb.CollegePeriodTo, eb.CollegeHighestLevelUnitsEarned, eb.CollegeYearGraduated, eb.CollegeScholarship, eb.GraduateStudiesName, eb.GraduateStudiesEducation, eb.GraduateStudiesPeriodFrom, eb.GraduateStudiesPeriodTo, eb.GraduateStudiesHighestLevelUnitsEarned, eb.GraduateStudiesYearGraduated, eb.GraduateStudiesScholarship, eb.UserID, eb.CreatedAt, eb.UpdatedAt,
	)

	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("Created educational background")

	return nil
}

func CreateCivilServices(db *sql.DB, cs types.CareerService) error {
	query := `INSERT INTO employee_civil_services (employee_id
		,career_service
		,rating
		,date_of_examination
		,place_of_examination
		,license_number
		,date_of_validitiy
		,user_id
		,created_at
		,updated_at
		,deleted_at
	) VALUES (@p1, @p2, @p3, @p4, @p5, @p6, @p7, @p8, @p9, @p10, @p11)
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(cs.EmployeeID, cs.CareerService, cs.Rating, cs.DateOfExamination, cs.PlaceOfExamination, cs.LicenseNumber, cs.DateOfValidity, cs.UserID, cs.CreatedAt, cs.UpdatedAt, cs.DeletedAt)

	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("Created civil services")
	return nil
}

func CreateWorkExperience(db *sql.DB, we types.EmployeeWorkExperience) error {

	query := `INSERT INTO employee_work_experiences (
		employee_id,
		date_from,
		date_to,
		position_title,
		department_agency_office_company,
		monthly_salary,
		salary_job_pay_grade,
		status_of_appointment,
		government_service,
		user_id,
		created_at,
		updated_at
	) VALUES (@p1, @p2, @p3, @p4, @p5, @p6, @p7, @p8, @p9, @p10, @p11, @p12)
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(we.EmployeeID, we.From, we.To, we.PositionTitle, we.Office, we.MonthlySalary, we.SalaryJobPayGrade, we.StatusOfAppointment, we.GovernmentService, we.UserID, we.CreatedAt, we.UpdatedAt)

	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("Created work experience")
	return nil
}


func CreateVoluntaryWork(db *sql.DB, vw types.VoluntaryWork) error {
	query := `INSERT INTO employee_voluntary_works (
		employee_id,
		organization_name,
		date_from,
		date_to,
		number_of_hours,
		position_nature_of_work,
		user_id,
		created_at,
		updated_at
	) VALUES (@p1, @p2, @p3, @p4, @p5, @p6, @p7, @p8, @p9)
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(vw.EmployeeID, vw.NameAndAddress, vw.InclusiveDateFrom, vw.InclusiveDateTo, vw.NoOfHours, vw.Position, vw.UserID, vw.CreatedAt, vw.UpdatedAt)

	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("Created voluntary work")
	return nil
}

func CreateTraining(db *sql.DB, tp types.Tranings) error {
	query := `INSERT INTO employee_training_attaineds (
		employee_id,
		title,
		date_of_attendance_from,
		date_of_attendance_to,
		number_of_hours,
		type_of_id,
		sponsored_by,
		user_id,
		created_at,
		updated_at
	) VALUES (@p1, @p2, @p3, @p4, @p5, @p6, @p7, @p8, @p9, @p10)
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(tp.EmployeeID, tp.Title, tp.DateOfAttendanceFrom, tp.DateOfAttendanceTo, tp.NumberOfHours, tp.TypeOfID, tp.SponsoredBy, tp.UserID, tp.CreatedAt, tp.UpdatedAt)

	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("Created training programs")
	return nil
}

func CreateOtherInformation(db *sql.DB, oi types.OtherInformation) error {
	query := `INSERT INTO employee_other_information (
		employee_id,
		special_skill,
		non_academic,
		organization,
		user_id,
		created_at,
		updated_at
	) VALUES (@p1, @p2, @p3, @p4, @p5, @p6, @p7)
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(oi.EmployeeID, oi.SpecialSkill, oi.NonAcademic, oi.Organization, oi.UserID, oi.CreatedAt, oi.UpdatedAt)

	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("Created other information")
	return nil
}

func CreateRelevantQuery(db *sql.DB, rq types.RelevantQueries) error {
	query := `INSERT INTO employee_relevant_queries (
		employee_id,
		question_34_a_answer,
		question_34_a_details,
		question_34_b_answer,
		question_34_b_details,
		question_35_a_answer,
		question_35_a_details,
		question_35_b_answer,
		question_35_b_details,
		question_35_b_date_filled,
		question_35_b_status_of_cases,
		question_36_a_answer,
		question_36_a_details,
		question_37_a_answer,
		question_37_a_details,
		question_38_a_answer,
		question_38_a_details,
		question_38_b_answer,
		question_38_b_details,
		question_39_a_answer,
		question_39_a_details,
		question_40_a_answer,
		question_40_a_details,
		question_40_b_answer,
		question_40_b_details,
		question_40_c_answer,
		question_40_c_details,
		user_id,
		created_at,
		updated_at 
	) VALUES (@p1, @p2, @p3, @p4, @p5, @p6, @p7, @p8, @p9, @p10, @p11, @p12, @p13, @p14, @p15, @p16, @p17, @p18, @p19, @p20, @p21, @p22, @p23, @p24, @p25, @p26, @p27, @p28, @p29, @p30)
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(rq.EmployeeID, rq.Question34AAnswer, rq.Question34ADetails, rq.Question34BAnswer, rq.Question34BDetails, rq.Question35AAnswer, rq.Question35ADetails, rq.Question35BAnswer, rq.Question35BDetails, rq.Question35BDateFilled, rq.Question35BStatusOfCases, rq.Question36AAnswer, rq.Question36ADetails, rq.Question37AAnswer, rq.Question37ADetails, rq.Question38AAnswer, rq.Question38ADetails, rq.Question38BAnswer, rq.Question38BDetails, rq.Question39AAnswer, rq.Question39ADetails, rq.Question40AAnswer, rq.Question40ADetails, rq.Question40BAnswer, rq.Question40BDetails, rq.Question40CAnswer, rq.Question40CDetails, rq.UserID, rq.CreatedAt, rq.UpdatedAt)

	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("Created relevant queries")
	return nil
}

func CreateReferences(db *sql.DB, rf types.EmployeeReference) error {
	query := `INSERT INTO employee_references (
		employee_id,
		name,
		address,
		telephone_number,
		user_id,
		created_at,
		updated_at
	) VALUES (@p1, @p2, @p3, @p4, @p5, @p6, @p7)
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(rf.EmployeeID, rf.Name, rf.Address, rf.TelephoneNumber, rf.UserID, rf.CreatedAt, rf.UpdatedAt)

	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("Created references")
	return nil
}


func CreateIssuedID(db *sql.DB, ch types.IssuedID) error {
	query := `INSERT INTO employee_issued_i_d_s (
		employee_id,
		id_type,
		id_no,
		"date",
		user_id,
		created_at,
		updated_at
	) VALUES (@p1, @p2, @p3, @p4, @p5, @p6, @p7)
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(ch.EmployeeID, ch.IDType, ch.IDNo, ch.IssueDate, ch.UserID, ch.CreatedAt, ch.UpdatedAt)

	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("Created issued ID")
	return nil
}

func CreateChildren(db *sql.DB, ch types.SpouseChildren) error {
	query := `INSERT INTO employee_spouse_childrens (
		employee_id,
		name,
		date_of_birth,
		user_id,
		created_at,
		updated_at
	) VALUES (@p1, @p2, @p3, @p4, @p5, @p6)
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(ch.EmployeeID, ch.Name, ch.DateOfBirth, ch.UserID, ch.CreatedAt, ch.UpdatedAt)

	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("Created children")
	return nil
}