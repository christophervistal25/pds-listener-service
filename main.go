package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	database "pds-listener-service/database"
	"pds-listener-service/types"
	utilities "pds-listener-service/utilities"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
)

var server = "DESKTOP-LNMCAJC"
var port = "1433"
var user = "sa"
var password = "christopher"
var databaseName = "PDS"
var err error
var db *sql.DB

func getFolderSize(path string) (int64, error) {
    var size int64
    err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() {
            size += info.Size()
        }
        return nil
    })
    return size, err
}

// Struct to hold the state of the folder size monitoring
type FolderMonitor struct {
    path       string
    size       int64
    mu         sync.Mutex
    watcher    *fsnotify.Watcher
    quit       chan struct{}
}

// Function to start monitoring folder size changes
func (fm *FolderMonitor) Start() error {
    var err error
    fm.watcher, err = fsnotify.NewWatcher()
    if err != nil {
        return err
    }
    defer fm.watcher.Close()

    fm.quit = make(chan struct{})
    go fm.watchEvents()

    err = fm.watcher.Add(fm.path)
    if err != nil {
        return err
    }

    // Initial size calculation
    size, err := getFolderSize(fm.path)
    if err != nil {
        return err
    }
    fm.mu.Lock()
    fm.size = size
    fm.mu.Unlock()

    // Periodic size checking every 10 seconds
    go fm.periodicSizeCheck()

    <-fm.quit
    return nil
}

// Function to watch for file system events
func (fm *FolderMonitor) watchEvents() {
    for {
        select {
        case event, ok := <-fm.watcher.Events:
            if !ok {
                return
            }
            if event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Create == fsnotify.Create || event.Op&fsnotify.Remove == fsnotify.Remove {
                fm.updateSize()
            }

            if event.Op&fsnotify.Write == fsnotify.Write && filepath.Ext(event.Name) == ".db" {
                
                databaseSqlite, err := database.ConnectSQLite(filepath.Base(event.Name))
                if(err != nil){
                    log.Fatal(err)
                }

                filename := filepath.Base(event.Name)
                filename = filename[:len(filename)-3]
              
                passkey, e := utilities.Decrypt(filename)
                if(e != nil){
                    log.Fatal(e)
                }

                query := `SELECT TOP (1) e.id AS employee_id, e.employee_id, e.lastname, e.firstname, e.middlename, e.extension, e.date_birth, e.place_birth, e.sex, e.civil_status, e.civil_status_others, e.height, e.weight,
                e.blood_type,
                e.gsis_id_no,
                e.gsis_policy_no,
                e.gsis_bp_no,
                e.pag_ibig_no,
                e.philhealth_no,
                e.sss_no,
                e.tin_no,
                e.lbp_account_no,
                e.dbp_account_no,
                e.agency_employee_no,
                e.citizenship,
                e.citizenship_by,
                e.indicate_country,
                e.residential_house_no,
                e.residential_street,
                e.residential_village,
                e.residential_barangay_code,
                e.residential_city_code,
                e.residential_province_string,
                e.residential_barangay_string,
                e.residential_city_string,
                e.residential_province_code,
                e.residential_zip_code,
                e.permanent_house_no,
                e.permanent_street,
                e.permanent_village,
                e.permanent_barangay_string,
                e.permanent_city_string,
                e.permanent_province_string,
                e.permanent_barangay_code,
                e.permanent_city_code,
                e.permanent_province_code,
                e.permanent_zip_code,
                e.telephone_no,
                e.mobile_no,
                e.email_address,
                e.employee_status,
                e.first_day_of_service,
                e.active_status,
                e.office,
                e.position,
                e.salary_grade,
                e.step,
                e.salary_rate,
                e.user_id,
                e.person_notified_name,
                e.person_notified_contact,
                e.person_notified_address,
                e.address,
                e.deleted_at,
                e.created_at,
                e.updated_at,
                u.email,
                u.username,
                u.lastname AS user_lastname,
                u.firstname AS user_firstname,
                u.middlename AS user_middlename,
                u.suffix,
                u.password,
                u.mobile_no AS user_mobile_no,
                u.date_of_birth AS user_date_of_birth,
                u.account_type,
                u.status AS user_status,
                u.office_code,
                u.office_detailed,
                u.approved_by,
                u.employee_id AS user_employee_id,
                u.created_at AS user_created_at,
                u.updated_at AS user_updated_at 
                FROM employees e INNER JOIN users u ON e.user_id = u.id
                WHERE u.username = @p1
                `

                queryFamilyBackground := `SELECT 
                    e.employee_id,
                    e.spouse_firstname,
                    e.spouse_lastname,
                    e.spouse_middlename,
                    e.spouse_extension,
                    e.spouse_occupation,
                    e.spouse_employer_business_name,
                    e.spouse_business_address,
                    e.spouse_telephone_number,
                    e.father_firstname,
                    e.father_lastname,
                    e.father_middlename,
                    e.father_extension,
                    e.mother_lastname,
                    e.mother_firstname,
                    e.mother_middlename,
                    e.user_id,
                    e.created_at,
                    e.updated_at
                    FROM employee_family_backgrounds e INNER JOIN users u ON e.user_id = u.id
                    WHERE u.username = @p1`
                
                
            queryEducationalBackground := `SELECT
                    e.employee_id
                    ,e.elementary_name
                    ,e.elementary_education
                    ,e.elementary_period_from
                    ,e.elementary_period_to
                    ,e.elementary_highest_level_units_earned
                    ,e.elementary_year_graduated
                    ,e.elementary_scholarship
                    ,e.secondary_name
                    ,e.secondary_education
                    ,e.secondary_period_from
                    ,e.secondary_period_to
                    ,e.secondary_highest_level_units_earned
                    ,e.secondary_year_graduated
                    ,e.secondary_scholarship
                    ,e.vocational_trade_course_name
                    ,e.vocational_education
                    ,e.vocational_trade_course_period_from
                    ,e.vocational_trade_course_period_to
                    ,e.vocational_trade_course_highest_level_units_earned
                    ,e.vocational_trade_course_year_graduated
                    ,e.vocational_trade_course_scholarship
                    ,e.college_name
                    ,e.college_education
                    ,e.college_period_from
                    ,e.college_period_to
                    ,e.college_highest_level_units_earned
                    ,e.college_year_graduated
                    ,e.college_scholarship
                    ,e.graduate_studies_name
                    ,e.graduate_studies_education
                    ,e.graduate_studies_period_from
                    ,e.graduate_studies_period_to
                    ,e.graduate_studies_highest_level_units_earned
                    ,e.graduate_studies_year_graduated
                    ,e.graduate_studies_scholarship
                    ,e.user_id
                    ,e.created_at
                    ,e.updated_at
                    FROM employee_educational_backgrounds e INNER JOIN users u ON e.user_id = u.id
                    WHERE u.username = @p1`

            queryServices := `SELECT 
                e.employee_id,
                e.career_service,
                e.rating,
                e.date_of_examination,
                e.place_of_examination,
                e.license_number,
                e.date_of_validitiy,
                e.user_id,
                e.created_at,
                e.updated_at,
                e.deleted_at
            FROM employee_civil_services e
            INNER JOIN users u ON e.user_id = u.id
            WHERE u.username = @p1 AND deleted_at IS NULL
            `

            queryWorkExperience := `SELECT
                work_experiences.employee_id,
                work_experiences."from",
                work_experiences."to",
                work_experiences.position_title,
                work_experiences.office,
                work_experiences.monthly_salary,
                work_experiences.salary_job_pay_grade,
                work_experiences.status_of_appointment,
                work_experiences.government_service,
                work_experiences.user_id,
                work_experiences.created_at,
                work_experiences.updated_at,
                work_experiences.deleted_at
                FROM employee_work_experiences as work_experiences
                INNER JOIN users u ON work_experiences.user_id = u.id
                WHERE u.username = @p1 AND work_experiences.deleted_at IS NULL
            `

            queryVoluntaryWorks := `SELECT 
                ew.employee_id
                ,ew.name_and_address
                ,ew.inclusive_date_from
                ,ew.inclusive_date_to
                ,ew.no_of_hours
                ,ew.position
                ,ew.user_id
                ,ew.created_at
                ,ew.updated_at
                FROM employee_voluntary_works ew
                INNER JOIN users u ON ew.user_id = u.id
                WHERE u.username = @p1
            `
            

            queryTrainings := `
                SELECT ea.employee_id
                ,ea.title
                ,ea.date_of_attendance_from
                ,ea.date_of_attendance_to
                ,ea.number_of_hours
                ,ea.type_of_id
                ,ea.sponsored_by
                ,ea.user_id
                ,ea.created_at
                ,ea.updated_at AS employee_training_updated_at
                FROM employee_training_attaineds ea
                INNER JOIN users u ON ea.user_id = u.id
                WHERE u.username = @p1
            `

            queryOtherInformation := `SELECT
                ei.employee_id
                ,ei.special_skill
                ,ei.non_academic
                ,ei.organization
                ,ei.user_id
                ,ei.created_at as employee_other_information_created_at
                ,ei.updated_at as employee_other_information_updated_at
                FROM employee_other_information as ei
                INNER JOIN users u ON ei.user_id = u.id
                WHERE u.username = @p1
            `

            queryRelevantQuery := `SELECT
                eq.employee_id
                ,eq.question_34_a_answer
                ,eq.question_34_a_details
                ,eq.question_34_b_answer
                ,eq.question_34_b_details
                ,eq.question_35_a_answer
                ,eq.question_35_a_details
                ,eq.question_35_b_answer
                ,eq.question_35_b_details
                ,eq.question_35_b_date_filled
                ,eq.question_35_b_status_of_cases
                ,eq.question_36_a_answer
                ,eq.question_36_a_details
                ,eq.question_37_a_answer
                ,eq.question_37_a_details
                ,eq.question_38_a_answer
                ,eq.question_38_a_details
                ,eq.question_38_b_answer
                ,eq.question_38_b_details
                ,eq.question_39_a_answer
                ,eq.question_39_a_details
                ,eq.question_40_a_answer
                ,eq.question_40_a_details
                ,eq.question_40_b_answer
                ,eq.question_40_b_details
                ,eq.question_40_c_answer
                ,eq.question_40_c_details
                ,eq.user_id
                ,eq.created_at
                ,eq.updated_at
                FROM employee_relevant_queries eq
                INNER JOIN users u ON eq.user_id = u.id
                WHERE u.username = @p1
            `


            queryReferences := `
                SELECT 
                r.employee_id
                ,r.name
                ,r.address
                ,r.telephone_number
                ,r.user_id
                ,r.created_at
                ,r.updated_at
                FROM employee_references as r
                INNER JOIN users u ON r.user_id = u.id
                WHERE u.username = @p1
            `

              
            queryIssuedID := `SELECT
                issued_id.employee_id
                ,issued_id.id_type
                ,issued_id.id_no
                ,issued_id."date" as issue_date
                ,issued_id.user_id
                ,issued_id.created_at
                ,issued_id.updated_at
                FROM employee_issued_i_d_s as issued_id
                INNER JOIN users u ON issued_id.user_id = u.id
                WHERE u.username = @p1
            `        
        
            
            queryChildrens := `
                SELECT
                    spouse_childrens.employee_id
                    ,spouse_childrens.name
                    ,spouse_childrens.date_of_birth
                    ,spouse_childrens.user_id
                    ,spouse_childrens.created_at
                    ,spouse_childrens.updated_at
                FROM employee_spouse_childrens as spouse_childrens
                INNER JOIN users u ON spouse_childrens.user_id = u.id
                WHERE u.username = @p1 AND deleted_at IS NULL
                `
            
                rows, employeeError := db.Query(query, passkey)
                rowsFamilyBackground, familyBackgroundError := db.Query(queryFamilyBackground, passkey)
                rowEducationalBackground, educationalBackgroundError := db.Query(queryEducationalBackground, passkey)
                rowCivilServices, civilServicesError := db.Query(queryServices, passkey)
                rowWorkExperience, workExperienceError := db.Query(queryWorkExperience, passkey)
                rowVoluntaryWorks, voluntaryWorksError := db.Query(queryVoluntaryWorks, passkey)
                rowTrainings, trainingsError := db.Query(queryTrainings, passkey)
                rowOtherInformation, otherInformationError := db.Query(queryOtherInformation, passkey)
                rowRelevantQuery, relevantQueryError := db.Query(queryRelevantQuery, passkey)
                rowReferences, referencesError := db.Query(queryReferences, passkey)
                rowIssuedID, issuedIDError := db.Query(queryIssuedID, passkey)
                rowChildrens, childrensError := db.Query(queryChildrens, passkey)
                fmt.Println(childrensError)
                if employeeError != nil && familyBackgroundError != nil && educationalBackgroundError != nil && civilServicesError != nil && workExperienceError != nil && voluntaryWorksError != nil && trainingsError != nil && otherInformationError != nil && relevantQueryError != nil  && referencesError != nil && issuedIDError != nil && childrensError != nil {
                    log.Fatal(err)
                }

                defer rows.Close()
                for rows.Next() {
                    var eu types.EmployeeUser
                    err := rows.Scan(
                        &eu.EmployeeID,
                        &eu.EmployeeEmpID,
                        &eu.Lastname,
                        &eu.Firstname,
                        &eu.Middlename,
                        &eu.Extension,
                        &eu.DateBirth,
                        &eu.PlaceBirth,
                        &eu.Sex,
                        &eu.CivilStatus,
                        &eu.CivilStatusOthers,
                        &eu.Height,
                        &eu.Weight,
                        &eu.BloodType,
                        &eu.GsisIDNo,
                        &eu.GsisPolicyNo,
                        &eu.GsisBPNo,
                        &eu.PagIbigNo,
                        &eu.PhilhealthNo,
                        &eu.SssNo,
                        &eu.TinNo,
                        &eu.LbpAccountNo,
                        &eu.DbpAccountNo,
                        &eu.AgencyEmployeeNo,
                        &eu.Citizenship,
                        &eu.CitizenshipBy,
                        &eu.IndicateCountry,
                        &eu.ResidentialHouseNo,
                        &eu.ResidentialStreet,
                        &eu.ResidentialVillage,
                        &eu.ResidentialBarangayCode,
                        &eu.ResidentialCityCode,
                        &eu.ResidentialProvinceString,
                        &eu.ResidentialBarangayString,
                        &eu.ResidentialCityString,
                        &eu.ResidentialProvinceCode,
                        &eu.ResidentialZipCode,
                        &eu.PermanentHouseNo,
                        &eu.PermanentStreet,
                        &eu.PermanentVillage,
                        &eu.PermanentBarangayString,
                        &eu.PermanentCityString,
                        &eu.PermanentProvinceString,
                        &eu.PermanentBarangayCode,
                        &eu.PermanentCityCode,
                        &eu.PermanentProvinceCode,
                        &eu.PermanentZipCode,
                        &eu.TelephoneNo,
                        &eu.MobileNo,
                        &eu.EmailAddress,
                        &eu.EmployeeStatus,
                        &eu.FirstDayOfService,
                        &eu.ActiveStatus,
                        &eu.Office,
                        &eu.Position,
                        &eu.SalaryGrade,
                        &eu.Step,
                        &eu.SalaryRate,
                        &eu.ID,
                        &eu.PersonNotifiedName,
                        &eu.PersonNotifiedContact,
                        &eu.PersonNotifiedAddress,
                        &eu.Address,
                        &eu.DeletedAt,
                        &eu.CreatedAt,
                        &eu.UpdatedAt,
                        &eu.UserEmail,
                        &eu.Username,
                        &eu.UserLastname,
                        &eu.UserFirstname,
                        &eu.UserMiddlename,
                        &eu.Suffix,
                        &eu.Password,
                        &eu.UserMobileNo,
                        &eu.UserDateOfBirth,
                        &eu.AccountType,
                        &eu.UserStatus,
                        &eu.OfficeCode,
                        &eu.OfficeDetailed,
                        &eu.ApprovedBy,
                        &eu.UserEmployeeID,
                        &eu.UserCreatedAt,
                        &eu.UserUpdatedAt,
                    )
                    if err != nil {
                        log.Fatal("Error scanning rows: ", err.Error())
                    }

                    database.CreateUser(databaseSqlite, eu)
                    fm.Pause()
                }

                defer rowsFamilyBackground.Close()
                for rowsFamilyBackground.Next() {
                    
                    var familyBackground types.FamilyInfo
                    err := rowsFamilyBackground.Scan(
                        &familyBackground.EmployeeID,
                        &familyBackground.SpouseFirstname,
                        &familyBackground.SpouseLastname,
                        &familyBackground.SpouseMiddlename,
                        &familyBackground.SpouseExtension,
                        &familyBackground.SpouseOccupation,
                        &familyBackground.SpouseEmployerBusiness,
                        &familyBackground.SpouseBusinessAddress,
                        &familyBackground.SpouseTelephoneNumber,
                        &familyBackground.FatherFirstname,
                        &familyBackground.FatherLastname,
                        &familyBackground.FatherMiddlename,
                        &familyBackground.FatherExtension,
                        &familyBackground.MotherLastname,
                        &familyBackground.MotherFirstname,
                        &familyBackground.MotherMiddlename,
                        &familyBackground.UserID,
                        &familyBackground.CreatedAt,
                        &familyBackground.UpdatedAt,
                    )
                    if err != nil {
                        log.Fatal("Error scanning rows: ", err.Error())
                    }
                    database.CreateFamilyInfo(databaseSqlite, familyBackground)
                    fm.Pause()
                }

                defer rowEducationalBackground.Close()
                for rowEducationalBackground.Next() {
                    var educationalBackground types.EducationalBackground
                    err := rowEducationalBackground.Scan(
                        &educationalBackground.EmployeeID,
                        &educationalBackground.ElementaryName,
                        &educationalBackground.ElementaryEducation,
                        &educationalBackground.ElementaryPeriodFrom,
                        &educationalBackground.ElementaryPeriodTo,
                        &educationalBackground.ElementaryHighestLevelUnitsEarned,
                        &educationalBackground.ElementaryYearGraduated,
                        &educationalBackground.ElementaryScholarship,
                        &educationalBackground.SecondaryName,
                        &educationalBackground.SecondaryEducation,
                        &educationalBackground.SecondaryPeriodFrom,
                        &educationalBackground.SecondaryPeriodTo,
                        &educationalBackground.SecondaryHighestLevelUnitsEarned,
                        &educationalBackground.SecondaryYearGraduated,
                        &educationalBackground.SecondaryScholarship,
                        &educationalBackground.VocationalTradeCourseName,
                        &educationalBackground.VocationalEducation,
                        &educationalBackground.VocationalTradeCoursePeriodFrom,
                        &educationalBackground.VocationalTradeCoursePeriodTo,
                        &educationalBackground.VocationalTradeCourseHighestLevelUnitsEarned,
                        &educationalBackground.VocationalTradeCourseYearGraduated,
                        &educationalBackground.VocationalTradeCourseScholarship,
                        &educationalBackground.CollegeName,
                        &educationalBackground.CollegeEducation,
                        &educationalBackground.CollegePeriodFrom,
                        &educationalBackground.CollegePeriodTo,
                        &educationalBackground.CollegeHighestLevelUnitsEarned,
                        &educationalBackground.CollegeYearGraduated,
                        &educationalBackground.CollegeScholarship,
                        &educationalBackground.GraduateStudiesName,
                        &educationalBackground.GraduateStudiesEducation,
                        &educationalBackground.GraduateStudiesPeriodFrom,
                        &educationalBackground.GraduateStudiesPeriodTo,
                        &educationalBackground.GraduateStudiesHighestLevelUnitsEarned,
                        &educationalBackground.GraduateStudiesYearGraduated,
                        &educationalBackground.GraduateStudiesScholarship,
                        &educationalBackground.UserID,
                        &educationalBackground.CreatedAt,
                        &educationalBackground.UpdatedAt,
                    )
                    if err != nil {
                        log.Fatal("Error scanning rows: ", err.Error())
                    }
                    database.CreateEducationalBackground(databaseSqlite, educationalBackground)
                    fm.Pause()
                }

                defer rowCivilServices.Close()
                for rowCivilServices.Next() {
                    var civilServices types.CareerService
                    err := rowCivilServices.Scan(
                        &civilServices.EmployeeID,
                        &civilServices.CareerService,
                        &civilServices.Rating,
                        &civilServices.DateOfExamination,
                        &civilServices.PlaceOfExamination,
                        &civilServices.LicenseNumber,
                        &civilServices.DateOfValidity,
                        &civilServices.UserID,
                        &civilServices.CreatedAt,
                        &civilServices.UpdatedAt,
                        &civilServices.DeletedAt,
                    )
                    if err != nil {
                        log.Fatal("Error scanning rows: ", err.Error())
                    }
                    database.CreateCivilServices(databaseSqlite, civilServices)
                    fm.Pause()
                }

                defer rowWorkExperience.Close()
                for rowWorkExperience.Next() {
                    var workExperience types.EmployeeWorkExperience

                    err := rowWorkExperience.Scan(
                        &workExperience.EmployeeID,
                        &workExperience.From,
                        &workExperience.To,
                        &workExperience.PositionTitle,
                        &workExperience.Office,
                        &workExperience.MonthlySalary,
                        &workExperience.SalaryJobPayGrade,
                        &workExperience.StatusOfAppointment,
                        &workExperience.GovernmentService,
                        &workExperience.UserID,
                        &workExperience.CreatedAt,
                        &workExperience.UpdatedAt,
                        &workExperience.DeletedAt,
                    )
                        
                    if err != nil {
                        log.Fatal("Error scanning rows: ", err.Error())
                    }
                    database.CreateWorkExperience(databaseSqlite, workExperience)
                    fm.Pause()
                }

                defer rowVoluntaryWorks.Close()
                for rowVoluntaryWorks.Next() {
                    var voluntaryWorks types.VoluntaryWork
                    err := rowVoluntaryWorks.Scan(
                        &voluntaryWorks.EmployeeID,
                        &voluntaryWorks.NameAndAddress,
                        &voluntaryWorks.InclusiveDateFrom,
                        &voluntaryWorks.InclusiveDateTo,
                        &voluntaryWorks.NoOfHours,
                        &voluntaryWorks.Position,
                        &voluntaryWorks.UserID,
                        &voluntaryWorks.CreatedAt,
                        &voluntaryWorks.UpdatedAt,
                    )

                    if err != nil {
                        log.Fatal("Error scanning rows: ", err.Error())
                    }
                    database.CreateVoluntaryWork(databaseSqlite, voluntaryWorks)
                    fm.Pause()
                }

                defer rowTrainings.Close()
                for rowTrainings.Next() {
                    var training types.Tranings
                    err := rowTrainings.Scan(
                        &training.EmployeeID,
                        &training.Title,
                        &training.DateOfAttendanceFrom,
                        &training.DateOfAttendanceTo,
                        &training.NumberOfHours,
                        &training.TypeOfID,
                        &training.SponsoredBy,
                        &training.UserID,
                        &training.CreatedAt,
                        &training.UpdatedAt,
                    )
                    if err != nil {
                        log.Fatal("Error scanning rows: ", err.Error())
                    }
                    database.CreateTraining(databaseSqlite, training)
                    fm.Pause()
                }

                defer rowOtherInformation.Close()
                for rowOtherInformation.Next() {
                    var otherInformation types.OtherInformation
                    err := rowOtherInformation.Scan(
                        &otherInformation.EmployeeID,
                        &otherInformation.SpecialSkill,
                        &otherInformation.NonAcademic,
                        &otherInformation.Organization,
                        &otherInformation.UserID,
                        &otherInformation.CreatedAt,
                        &otherInformation.UpdatedAt,
                    )

                    if err != nil {
                        log.Fatal("Error scanning rows: ", err.Error())
                    }

                    database.CreateOtherInformation(databaseSqlite, otherInformation)
                    fm.Pause()
                }

                defer rowRelevantQuery.Close()
                for rowRelevantQuery.Next() {
                    var relevantQuery types.RelevantQueries
                    err := rowRelevantQuery.Scan(
                        &relevantQuery.EmployeeID,
                        &relevantQuery.Question34AAnswer,
                        &relevantQuery.Question34ADetails,
                        &relevantQuery.Question34BAnswer,
                        &relevantQuery.Question34BDetails,
                        &relevantQuery.Question35AAnswer,
                        &relevantQuery.Question35ADetails,
                        &relevantQuery.Question35BAnswer,
                        &relevantQuery.Question35BDetails,
                        &relevantQuery.Question35BDateFilled,
                        &relevantQuery.Question35BStatusOfCases,
                        &relevantQuery.Question36AAnswer,
                        &relevantQuery.Question36ADetails,
                        &relevantQuery.Question37AAnswer,
                        &relevantQuery.Question37ADetails,
                        &relevantQuery.Question38AAnswer,
                        &relevantQuery.Question38ADetails,
                        &relevantQuery.Question38BAnswer,
                        &relevantQuery.Question38BDetails,
                        &relevantQuery.Question39AAnswer,
                        &relevantQuery.Question39ADetails,
                        &relevantQuery.Question40AAnswer,
                        &relevantQuery.Question40ADetails,
                        &relevantQuery.Question40BAnswer,
                        &relevantQuery.Question40BDetails,
                        &relevantQuery.Question40CAnswer,
                        &relevantQuery.Question40CDetails,
                        &relevantQuery.UserID,
                        &relevantQuery.CreatedAt,
                        &relevantQuery.UpdatedAt,
                    )

                    if err != nil {
                        log.Fatal("Error scanning rows: ", err.Error())
                    }

                    database.CreateRelevantQuery(databaseSqlite, relevantQuery)
                    fm.Pause()
                }

                defer rowReferences.Close()
                for rowReferences.Next() {
                    var references types.EmployeeReference
                    err := rowReferences.Scan(
                        &references.EmployeeID,
                        &references.Name,
                        &references.Address,
                        &references.TelephoneNumber,
                        &references.UserID,
                        &references.CreatedAt,
                        &references.UpdatedAt,
                    )

                    if err != nil {
                        log.Fatal("Error scanning rows: ", err.Error())
                    }

                    database.CreateReferences(databaseSqlite, references)
                    fm.Pause()
                }

                defer rowIssuedID.Close()
                for rowIssuedID.Next() {
                    var issuedID types.IssuedID
                    err := rowIssuedID.Scan(
                        &issuedID.EmployeeID,
                        &issuedID.IDType,
                        &issuedID.IDNo,
                        &issuedID.IssueDate,
                        &issuedID.UserID,
                        &issuedID.CreatedAt,
                        &issuedID.UpdatedAt,
                    )

                    if err != nil {
                        log.Fatal("Error scanning rows: ", err.Error())
                    }

                    database.CreateIssuedID(databaseSqlite, issuedID)
                    fm.Pause()
                }

                defer rowChildrens.Close()
                for rowChildrens.Next() {
                    var children types.SpouseChildren
                    err := rowChildrens.Scan(
                        &children.EmployeeID,
                        &children.Name,
                        &children.DateOfBirth,
                        &children.UserID,
                        &children.CreatedAt,
                        &children.UpdatedAt,
                    )

                    if err != nil {
                        log.Fatal("Error scanning rows: ", err.Error())
                    }

                    database.CreateChildren(databaseSqlite, children)
                }

                fm.Start()

                if err := rows.Err(); err != nil {
                    log.Fatal(err)
                }
            }
        case err, ok := <-fm.watcher.Errors:
            if !ok {
                return
            }
            log.Println("error:", err)
        }
    }
}

// Function to periodically check the folder size
func (fm *FolderMonitor) periodicSizeCheck() {
    ticker := time.NewTicker(1 * time.Second)
    for {
        select {
        case <-ticker.C:
            fm.updateSize()
        case <-fm.quit:
            ticker.Stop()
            return
        }
    }
}

// Function to update the folder size
func (fm *FolderMonitor) updateSize() {
    size, err := getFolderSize(fm.path)
    if err != nil {
        log.Println("error getting folder size:", err)
        return
    }
    fm.mu.Lock()
    if fm.size != size {
        fm.size = size
        fmt.Printf("Folder size changed: %d bytes\n", size)
    }
    fm.mu.Unlock()
}

// Function to stop monitoring
func (fm *FolderMonitor) Stop() {
    close(fm.quit)
}

// function to pause monitoring
func (fm *FolderMonitor) Pause() {
    fm.watcher.Close()
}


func main() {

    if len(os.Args) < 2 {
        log.Fatal("Please specify a folder to monitor")
    }
  
    db, err = database.ConnectSQLServer(server, user, password, port, databaseName)
    if err != nil {
        log.Fatal(err)
    }
    folder := os.Args[1]

    fm := &FolderMonitor{path: folder}
    if err := fm.Start(); err != nil {
        log.Fatal(err)
    }
}
