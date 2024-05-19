package types

import (
	"database/sql"
	"time"
)



type User struct {
    ID            *int
    Email         *string
    Username      *string
    LastName      *string
    FirstName     *string
    MiddleName    *string
    Suffix        *string
    Password      *string
    MobileNo      *string
    DateOfBirth   *time.Time
    AccountType   *string
    Status        *string
    OfficeCode    *string
    OfficeDetailed *string
    ApprovedBy    *string
    EmployeeID    *int
    CreatedAt     *time.Time
    UpdatedAt     *time.Time
}

type Employee struct {
    EmployeeID            *int
    LastName              *string
    FirstName             *string
    MiddleName            *string
    Extension             *string
    DateOfBirth           *time.Time
    PlaceOfBirth          *string
    Sex                   *string
    CivilStatus           *string
    CivilStatusOthers     *string
    Height                *float64
    Weight                *float64
    BloodType             *string
    GSISIDNo              *string
    GSISPolicyNo          *string
    GSISBPNo              *string
    PagIBIGNo             *string
    PhilHealthNo          *string
    SSSNo                 *string
    TINNo                 *string
    LBPAccountNo          *string
    DBPAccountNo          *string
    AgencyEmployeeNo      *string
    Citizenship           *string
    CitizenshipBy         *string
    IndicateCountry       *string
    ResidentialHouseNo    *string
    ResidentialStreet     *string
    ResidentialVillage    *string
    ResidentialBarangayCode *string
    ResidentialCityCode   *string
    ResidentialProvinceString *string
    ResidentialBarangayString *string
    ResidentialCityString *string
    ResidentialProvinceCode *string
    ResidentialZipCode    *string
    PermanentHouseNo      *string
    PermanentStreet       *string
    PermanentVillage      *string
    PermanentBarangayString *string
    PermanentCityString   *string
    PermanentProvinceString *string
    PermanentBarangayCode *string
    PermanentCityCode     *string
    PermanentProvinceCode *string
    PermanentZipCode      *string
    TelephoneNo           *string
    MobileNo              *string
    EmailAddress          *string
    EmployeeStatus        *string
    FirstDayOfService     *time.Time
    ActiveStatus          *string
    Office                *string
    Position              *string
    SalaryGrade           *string
    Step                  *string
    SalaryRate            *float64
    UserID                *int
    PersonNotifiedName    *string
    PersonNotifiedContact *string
    PersonNotifiedAddress *string
    Address               *string
    DeletedAt             *time.Time
    CreatedAt             *time.Time
    UpdatedAt             *time.Time
}

type EmployeeUser struct {
	ID               		sql.NullInt64  `json:"id"`
	EmployeeID               sql.NullInt64  `json:"employee_id"`
	EmployeeEmpID            sql.NullString `json:"employee_emp_id"`
	Lastname                 sql.NullString `json:"lastname"`
	Firstname                sql.NullString `json:"firstname"`
	Middlename               sql.NullString `json:"middlename"`
	Extension                sql.NullString `json:"extension"`
	DateBirth                sql.NullTime   `json:"date_birth"`
	PlaceBirth               sql.NullString `json:"place_birth"`
	Sex                      sql.NullString `json:"sex"`
	CivilStatus              sql.NullString `json:"civil_status"`
	CivilStatusOthers        sql.NullString `json:"civil_status_others"`
	Height                   sql.NullFloat64 `json:"height"`
	Weight                   sql.NullFloat64 `json:"weight"`
	BloodType                sql.NullString `json:"blood_type"`
	GsisIDNo                 sql.NullString `json:"gsis_id_no"`
	GsisPolicyNo             sql.NullString `json:"gsis_policy_no"`
	GsisBPNo                 sql.NullString `json:"gsis_bp_no"`
	PagIbigNo                sql.NullString `json:"pag_ibig_no"`
	PhilhealthNo             sql.NullString `json:"philhealth_no"`
	SssNo                    sql.NullString `json:"sss_no"`
	TinNo                    sql.NullString `json:"tin_no"`
	LbpAccountNo             sql.NullString `json:"lbp_account_no"`
	DbpAccountNo             sql.NullString `json:"dbp_account_no"`
	AgencyEmployeeNo         sql.NullString `json:"agency_employee_no"`
	Citizenship              sql.NullString `json:"citizenship"`
	CitizenshipBy            sql.NullString `json:"citizenship_by"`
	IndicateCountry          sql.NullString `json:"indicate_country"`
	ResidentialHouseNo       sql.NullString `json:"residential_house_no"`
	ResidentialStreet        sql.NullString `json:"residential_street"`
	ResidentialVillage       sql.NullString `json:"residential_village"`
	ResidentialBarangayCode  sql.NullString `json:"residential_barangay_code"`
	ResidentialCityCode      sql.NullString `json:"residential_city_code"`
	ResidentialProvinceString sql.NullString `json:"residential_province_string"`
	ResidentialBarangayString sql.NullString `json:"residential_barangay_string"`
	ResidentialCityString    sql.NullString `json:"residential_city_string"`
	ResidentialProvinceCode  sql.NullString `json:"residential_province_code"`
	ResidentialZipCode       sql.NullString `json:"residential_zip_code"`
	PermanentHouseNo         sql.NullString `json:"permanent_house_no"`
	PermanentStreet          sql.NullString `json:"permanent_street"`
	PermanentVillage         sql.NullString `json:"permanent_village"`
	PermanentBarangayString  sql.NullString `json:"permanent_barangay_string"`
	PermanentCityString      sql.NullString `json:"permanent_city_string"`
	PermanentProvinceString  sql.NullString `json:"permanent_province_string"`
	PermanentBarangayCode    sql.NullString `json:"permanent_barangay_code"`
	PermanentCityCode        sql.NullString `json:"permanent_city_code"`
	PermanentProvinceCode    sql.NullString `json:"permanent_province_code"`
	PermanentZipCode         sql.NullString `json:"permanent_zip_code"`
	TelephoneNo              sql.NullString `json:"telephone_no"`
	MobileNo                 sql.NullString `json:"mobile_no"`
	EmailAddress             sql.NullString `json:"email_address"`
	EmployeeStatus           sql.NullInt64  `json:"employee_status"`
	FirstDayOfService        sql.NullTime   `json:"first_day_of_service"`
	ActiveStatus             sql.NullString `json:"active_status"`
	Office                   sql.NullString `json:"office"`
	Position                 sql.NullString `json:"position"`
	SalaryGrade              sql.NullInt64  `json:"salary_grade"`
	Step                     sql.NullInt64  `json:"step"`
	SalaryRate               sql.NullFloat64 `json:"salary_rate"`
	UserID                   sql.NullInt64  `json:"user_id"`
	PersonNotifiedName       sql.NullString `json:"person_notified_name"`
	PersonNotifiedContact    sql.NullString `json:"person_notified_contact"`
	PersonNotifiedAddress    sql.NullString `json:"person_notified_address"`
	Address                  sql.NullString `json:"address"`
	DeletedAt                sql.NullTime   `json:"deleted_at"`
	CreatedAt                sql.NullTime   `json:"created_at"`
	UpdatedAt                sql.NullTime   `json:"updated_at"`
	UserEmail                sql.NullString `json:"user_email"`
	Username                 sql.NullString `json:"username"`
	UserLastname             sql.NullString `json:"user_lastname"`
	UserFirstname            sql.NullString `json:"user_firstname"`
	UserMiddlename           sql.NullString `json:"user_middlename"`
	Suffix                   sql.NullString `json:"suffix"`
	Password                 sql.NullString `json:"password"`
	UserMobileNo             sql.NullString `json:"user_mobile_no"`
	UserDateOfBirth          sql.NullTime   `json:"user_date_of_birth"`
	AccountType              sql.NullString `json:"account_type"`
	UserStatus               sql.NullString `json:"user_status"`
	OfficeCode               sql.NullString `json:"office_code"`
	OfficeDetailed           sql.NullString `json:"office_detailed"`
	ApprovedBy               sql.NullString `json:"approved_by"`
	UserEmployeeID           sql.NullString `json:"user_employee_id"`
	UserCreatedAt            sql.NullTime   `json:"user_created_at"`
	UserUpdatedAt            sql.NullTime   `json:"user_updated_at"`
}


type FamilyInfo struct {
	EmployeeID              sql.NullInt64      `json:"employee_id"`
	SpouseFirstname         sql.NullString    `json:"spouse_firstname"`
	SpouseLastname          sql.NullString    `json:"spouse_lastname"`
	SpouseMiddlename        sql.NullString    `json:"spouse_middlename"`
	SpouseExtension         sql.NullString    `json:"spouse_extension"`
	SpouseOccupation        sql.NullString    `json:"spouse_occupation"`
	SpouseEmployerBusiness  sql.NullString    `json:"spouse_employer_business_name"`
	SpouseBusinessAddress   sql.NullString    `json:"spouse_business_address"`
	SpouseTelephoneNumber   sql.NullString    `json:"spouse_telephone_number"`
	FatherFirstname         sql.NullString    `json:"father_firstname"`
	FatherLastname          sql.NullString    `json:"father_lastname"`
	FatherMiddlename        sql.NullString    `json:"father_middlename"`
	FatherExtension         sql.NullString    `json:"father_extension"`
	MotherLastname          sql.NullString    `json:"mother_lastname"`
	MotherFirstname         sql.NullString    `json:"mother_firstname"`
	MotherMiddlename        sql.NullString    `json:"mother_middlename"`
	UserID                  sql.NullInt64       `json:"user_id"`
	CreatedAt               sql.NullTime  `json:"created_at"`
	UpdatedAt               sql.NullTime `json:"updated_at"`
}


type EducationalBackground struct {
    EmployeeID                              sql.NullInt64   `json:"employee_id"`
    ElementaryName                          sql.NullString  `json:"elementary_name"`
    ElementaryEducation                     sql.NullString  `json:"elementary_education"`
    ElementaryPeriodFrom                    sql.NullString  `json:"elementary_period_from"`
    ElementaryPeriodTo                      sql.NullString  `json:"elementary_period_to"`
    ElementaryHighestLevelUnitsEarned       sql.NullString  `json:"elementary_highest_level_units_earned"`
    ElementaryYearGraduated                 sql.NullString  `json:"elementary_year_graduated"`
    ElementaryScholarship                   sql.NullString  `json:"elementary_scholarship"`
    SecondaryName                           sql.NullString  `json:"secondary_name"`
    SecondaryEducation                      sql.NullString  `json:"secondary_education"`
    SecondaryPeriodFrom                     sql.NullString  `json:"secondary_period_from"`
    SecondaryPeriodTo                       sql.NullString  `json:"secondary_period_to"`
    SecondaryHighestLevelUnitsEarned        sql.NullString  `json:"secondary_highest_level_units_earned"`
    SecondaryYearGraduated                  sql.NullString  `json:"secondary_year_graduated"`
    SecondaryScholarship                    sql.NullString  `json:"secondary_scholarship"`
    VocationalTradeCourseName               sql.NullString  `json:"vocational_trade_course_name"`
    VocationalEducation                     sql.NullString  `json:"vocational_education"`
    VocationalTradeCoursePeriodFrom         sql.NullString  `json:"vocational_trade_course_period_from"`
    VocationalTradeCoursePeriodTo           sql.NullString  `json:"vocational_trade_course_period_to"`
    VocationalTradeCourseHighestLevelUnitsEarned sql.NullString  `json:"vocational_trade_course_highest_level_units_earned"`
    VocationalTradeCourseYearGraduated      sql.NullString  `json:"vocational_trade_course_year_graduated"`
    VocationalTradeCourseScholarship        sql.NullString  `json:"vocational_trade_course_scholarship"`
    CollegeName                             sql.NullString  `json:"college_name"`
    CollegeEducation                        sql.NullString  `json:"college_education"`
    CollegePeriodFrom                       sql.NullString  `json:"college_period_from"`
    CollegePeriodTo                         sql.NullString  `json:"college_period_to"`
    CollegeHighestLevelUnitsEarned          sql.NullString  `json:"college_highest_level_units_earned"`
    CollegeYearGraduated                    sql.NullString  `json:"college_year_graduated"`
    CollegeScholarship                      sql.NullString  `json:"college_scholarship"`
    GraduateStudiesName                     sql.NullString  `json:"graduate_studies_name"`
    GraduateStudiesEducation                sql.NullString  `json:"graduate_studies_education"`
    GraduateStudiesPeriodFrom               sql.NullString  `json:"graduate_studies_period_from"`
    GraduateStudiesPeriodTo                 sql.NullString  `json:"graduate_studies_period_to"`
    GraduateStudiesHighestLevelUnitsEarned  sql.NullString  `json:"graduate_studies_highest_level_units_earned"`
    GraduateStudiesYearGraduated            sql.NullString  `json:"graduate_studies_year_graduated"`
    GraduateStudiesScholarship              sql.NullString  `json:"graduate_studies_scholarship"`
    UserID                                  sql.NullInt64   `json:"user_id"`
    CreatedAt                               sql.NullTime    `json:"created_at"`
    UpdatedAt                               sql.NullTime    `json:"updated_at"`
}



type CareerService  struct {
    EmployeeID          sql.NullInt64   `json:"employee_id"`
    CareerService       sql.NullString  `json:"career_service"`
    Rating              sql.NullString  `json:"rating"`
    DateOfExamination   sql.NullTime    `json:"date_of_examination"`
    PlaceOfExamination  sql.NullString  `json:"place_of_examination"`
    LicenseNumber       sql.NullString  `json:"license_number"`
    DateOfValidity      sql.NullTime    `json:"date_of_validity"`
    UserID              sql.NullInt64   `json:"user_id"`
    CreatedAt           sql.NullTime    `json:"created_at"`
    UpdatedAt           sql.NullTime    `json:"updated_at"`
    DeletedAt           sql.NullTime    `json:"deleted_at"`
}


type EmployeeWorkExperience struct {
    EmployeeID          sql.NullInt64  `json:"employee_id"`
    From                sql.NullString `json:"from"`
    To                  sql.NullString `json:"to"`
    PositionTitle       sql.NullString `json:"position_title"`
    Office              sql.NullString `json:"office"`
    MonthlySalary       sql.NullString `json:"monthly_salary"`
    SalaryJobPayGrade   sql.NullString `json:"salary_job_pay_grade"`
    StatusOfAppointment sql.NullString `json:"status_of_appointment"`
    GovernmentService   sql.NullString   `json:"government_service"`
    UserID              sql.NullInt64   `json:"user_id"`
    CreatedAt           sql.NullTime    `json:"created_at"`
    UpdatedAt           sql.NullTime    `json:"updated_at"`
    DeletedAt           sql.NullTime   `json:"deleted_at,omitempty"`
}

type VoluntaryWork struct {
	EmployeeID        sql.NullString            `json:"employee_id"`
	NameAndAddress    sql.NullString `json:"name_and_address"`
	InclusiveDateFrom sql.NullTime   `json:"inclusive_date_from"`
	InclusiveDateTo   sql.NullTime   `json:"inclusive_date_to"`
	NoOfHours         sql.NullInt32  `json:"no_of_hours"`
	Position          sql.NullString `json:"position"`
	UserID            sql.NullInt64  `json:"user_id"`
	CreatedAt         sql.NullTime      `json:"created_at"`
	UpdatedAt         sql.NullTime      `json:"updated_at"`
}

type Tranings struct {
	ID                   sql.NullInt64            `json:"id"`
	EmployeeID           sql.NullString `json:"employee_id"`
	Title                sql.NullString `json:"title"`
	DateOfAttendanceFrom sql.NullTime   `json:"date_of_attendance_from"`
	DateOfAttendanceTo   sql.NullTime   `json:"date_of_attendance_to"`
	NumberOfHours        sql.NullInt32  `json:"number_of_hours"`
	TypeOfID             sql.NullString `json:"type_of_id"`
	SponsoredBy          sql.NullString `json:"sponsored_by"`
	UserID               sql.NullInt64            `json:"user_id"`
	CreatedAt            sql.NullTime      `json:"created_at"`
	UpdatedAt            sql.NullTime      `json:"updated_at"`
}

type OtherInformation struct {
    EmployeeID    sql.NullString `json:"employee_id"`
    SpecialSkill  sql.NullString `json:"special_skill"`
    NonAcademic   sql.NullString `json:"non_academic"`
    Organization  sql.NullString `json:"organization"`
    UserID        sql.NullInt64            `json:"user_id"`
    CreatedAt     sql.NullTime      `json:"created_at"`
    UpdatedAt     sql.NullTime      `json:"updated_at"`
}

type RelevantQueries struct {
    EmployeeID              sql.NullString `json:"employee_id"`
    Question34AAnswer       sql.NullString `json:"question_34_a_answer"`
    Question34ADetails      sql.NullString `json:"question_34_a_details"`
    Question34BAnswer       sql.NullString `json:"question_34_b_answer"`
    Question34BDetails      sql.NullString `json:"question_34_b_details"`
    Question35AAnswer       sql.NullString `json:"question_35_a_answer"`
    Question35ADetails      sql.NullString `json:"question_35_a_details"`
    Question35BAnswer       sql.NullString `json:"question_35_b_answer"`
    Question35BDetails      sql.NullString `json:"question_35_b_details"`
    Question35BDateFilled   sql.NullTime   `json:"question_35_b_date_filled"`
    Question35BStatusOfCases sql.NullString `json:"question_35_b_status_of_cases"`
    Question36AAnswer       sql.NullString `json:"question_36_a_answer"`
    Question36ADetails      sql.NullString `json:"question_36_a_details"`
    Question37AAnswer       sql.NullString `json:"question_37_a_answer"`
    Question37ADetails      sql.NullString `json:"question_37_a_details"`
    Question38AAnswer       sql.NullString `json:"question_38_a_answer"`
    Question38ADetails      sql.NullString `json:"question_38_a_details"`
    Question38BAnswer       sql.NullString `json:"question_38_b_answer"`
    Question38BDetails      sql.NullString `json:"question_38_b_details"`
    Question39AAnswer       sql.NullString `json:"question_39_a_answer"`
    Question39ADetails      sql.NullString `json:"question_39_a_details"`
    Question40AAnswer       sql.NullString `json:"question_40_a_answer"`
    Question40ADetails      sql.NullString `json:"question_40_a_details"`
    Question40BAnswer       sql.NullString `json:"question_40_b_answer"`
    Question40BDetails      sql.NullString `json:"question_40_b_details"`
    Question40CAnswer       sql.NullString `json:"question_40_c_answer"`
    Question40CDetails      sql.NullString `json:"question_40_c_details"`
    UserID                  sql.NullInt64            `json:"user_id"`
    CreatedAt               sql.NullTime      `json:"created_at"`
    UpdatedAt               sql.NullTime      `json:"updated_at"`
}

type EmployeeReference struct {
    EmployeeID      sql.NullString `json:"employee_id"`
    Name            sql.NullString `json:"name"`
    Address         sql.NullString `json:"address"`
    TelephoneNumber sql.NullString `json:"telephone_number"`
    UserID          sql.NullInt64            `json:"user_id"`
    CreatedAt       sql.NullTime      `json:"created_at"`
    UpdatedAt       sql.NullTime      `json:"updated_at"`
}

type IssuedID struct {
    EmployeeID sql.NullString `json:"employee_id"`
    IDType    sql.NullString `json:"id_type"`
    IDNo      sql.NullString `json:"id_no"`
    IssueDate sql.NullString   `json:"date"`
    UserID    sql.NullInt64   `json:"user_id"`
    CreatedAt sql.NullTime    `json:"created_at"`
    UpdatedAt sql.NullTime    `json:"updated_at"`
}

type SpouseChildren struct {
    EmployeeID   sql.NullString `json:"employee_id"`
    Name         sql.NullString `json:"name"`
    DateOfBirth  sql.NullTime   `json:"date_of_birth"`
    UserID       sql.NullInt64            `json:"user_id"`
    CreatedAt    sql.NullTime      `json:"created_at"`
    UpdatedAt    sql.NullTime      `json:"updated_at"`
    DeletedAt    sql.NullTime   `json:"deleted_at"`
}