package database

import (
	"database/sql"
	"fmt"
	"log"
	"pds-listener-service/types"
)

func GetPersonalInformation(db *sql.DB, username string) (types.EmployeeUser, error) {
	row := db.QueryRow(`SELECT * FROM employees`)
	var Employee types.EmployeeUser

	err := row.Scan(
		&Employee.EmployeeID,
		&Employee.EmployeeEmpID,
		&Employee.Lastname,
		&Employee.Firstname,
		&Employee.Middlename,
		&Employee.Extension,
		&Employee.DateBirth,
		&Employee.PlaceBirth,
		&Employee.Sex,
		&Employee.CivilStatus,
		&Employee.CivilStatusOthers,
		&Employee.Height,
		&Employee.Weight,
		&Employee.BloodType,
		&Employee.GsisIDNo,
		&Employee.GsisPolicyNo,
		&Employee.GsisBPNo,
		&Employee.PagIbigNo,
		&Employee.PhilhealthNo,
		&Employee.SssNo,
		&Employee.TinNo,
		&Employee.LbpAccountNo,
		&Employee.DbpAccountNo,
		&Employee.AgencyEmployeeNo,
		&Employee.Citizenship,
		&Employee.CitizenshipBy,
		&Employee.IndicateCountry,
		&Employee.ResidentialHouseNo,
		&Employee.ResidentialStreet,
		&Employee.ResidentialVillage,
		&Employee.ResidentialBarangayCode,
		&Employee.ResidentialCityCode,
		&Employee.ResidentialProvinceString,
		&Employee.ResidentialBarangayString,
		&Employee.ResidentialCityString,
		&Employee.ResidentialProvinceCode,
		&Employee.ResidentialZipCode,
		&Employee.PermanentHouseNo,
		&Employee.PermanentStreet,
		&Employee.PermanentVillage,
		&Employee.PermanentBarangayString,
		&Employee.PermanentCityString,
		&Employee.PermanentProvinceString,
		&Employee.PermanentBarangayCode,
		&Employee.PermanentCityCode,
		&Employee.PermanentProvinceCode,
		&Employee.PermanentZipCode,
		&Employee.TelephoneNo,
		&Employee.MobileNo,
		&Employee.EmailAddress,
		&Employee.EmployeeStatus,
		&Employee.FirstDayOfService,
		&Employee.ActiveStatus,
		&Employee.Office,
		&Employee.Position,
		&Employee.SalaryGrade,
		&Employee.Step,
		&Employee.SalaryRate,
		&Employee.ID,
		&Employee.PersonNotifiedName,
		&Employee.PersonNotifiedContact,
		&Employee.PersonNotifiedAddress,
		&Employee.Address,
		&Employee.DeletedAt,
		&Employee.CreatedAt,
		&Employee.UpdatedAt,
	)


	if err != nil {
		if err == sql.ErrNoRows {
			return Employee, nil 
		}
		return Employee, err
	}

	return Employee, nil
}

func GetFamilyBackground(db *sql.DB, username string) (types.FamilyInfo, error) {
	row := db.QueryRow(`SELECT 
		employee_id
		,spouse_firstname
		,spouse_lastname
		,spouse_middlename
		,spouse_extension
		,spouse_occupation
		,spouse_employer_business_name
		,spouse_business_address
		,spouse_telephone_number
		,father_firstname
		,father_lastname
		,father_middlename
		,father_extension
		,mother_lastname
		,mother_firstname
		,mother_middlename
		,user_id
		,created_at
		,updated_at
	FROM employee_family_backgrounds`)
	
	var familyBackground types.FamilyInfo

	
	err := row.Scan(
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
		if err == sql.ErrNoRows {
			return familyBackground, nil
		}
		return familyBackground, err
	} else {
		return familyBackground, nil
	}
}

func GetChildren(db *sql.DB, username string) ([]types.SpouseChildren, error) {
	rows, err := db.Query(`SELECT
		employee_id,
		name,
		date_of_birth,
		user_id,
		created_at,
		updated_at
		FROM employee_spouse_childrens
		GROUP BY employee_id, name, date_of_birth
	`)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var children = []types.SpouseChildren{}

	for rows.Next() {
		var child types.SpouseChildren

		err := rows.Scan(&child.EmployeeID, &child.Name, &child.DateOfBirth, &child.UserID, &child.CreatedAt, &child.UpdatedAt)

		if err != nil {
			log.Fatal(err)
		}

		children = append(children, child)
	}
	

	if err != nil {
		return nil, err
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return []types.SpouseChildren{}, nil
		}
		return children, err
	} else {
		return children, nil
	}
}


func UpdatePersonalInformation(sqlServerDatabase *sql.DB, sqliteDatabase *sql.DB, username string) {
	eu, _ := GetPersonalInformation(sqliteDatabase, username)
	if(eu.Firstname.Valid) {
		user := sqlServerDatabase.QueryRow(`SELECT id from users where username = @p1`, username)

		var id int
		user.Scan(&id)
		query := `UPDATE employees SET
		lastname = @p2, firstname = @p3, middlename = @p4, extension = @p5, date_birth = @p6, place_birth = @p7, sex = @p8, civil_status = @p9, 
		civil_status_others = @p10, height = @p11, weight = @p12, blood_type = @p13, gsis_id_no = @p14, gsis_policy_no = @p15, gsis_bp_no = @p16, pag_ibig_no = @p17, 
		philhealth_no = @p18, sss_no = @p19, tin_no = @p20, lbp_account_no = @p21, dbp_account_no = @p22, agency_employee_no = @p23, citizenship = @p24, 
		citizenship_by = @p25, indicate_country = @p26, residential_house_no = @p27, residential_street = @p28, residential_village = @p29, 
		residential_barangay_code = @p30, residential_city_code = @p31, residential_province_string = @p32, residential_barangay_string = @p33, 
		residential_city_string = @p34, residential_province_code = @p35, residential_zip_code = @p36, permanent_house_no = @p37, permanent_street = @p38, 
		permanent_village = @p39, permanent_barangay_string = @p40, permanent_city_string = @p41, permanent_province_string = @p42, 
		permanent_barangay_code = @p43, permanent_city_code = @p44, permanent_province_code = @p45, permanent_zip_code = @p46, telephone_no = @p47, 
		mobile_no = @p48, email_address = @p49
		WHERE user_id = @p1`

		stmt, err := sqlServerDatabase.Prepare(query)

		if err != nil {
			log.Fatal(err)
		}

		defer stmt.Close()

		_, err = stmt.Exec(
			id, eu.Lastname, eu.Firstname, eu.Middlename, eu.Extension, eu.DateBirth, eu.PlaceBirth, eu.Sex, eu.CivilStatus,
			eu.CivilStatusOthers, eu.Height, eu.Weight, eu.BloodType, eu.GsisIDNo, eu.GsisPolicyNo, eu.GsisBPNo, eu.PagIbigNo, eu.PhilhealthNo, eu.SssNo, eu.TinNo, eu.LbpAccountNo, eu.DbpAccountNo, eu.AgencyEmployeeNo, eu.Citizenship, eu.CitizenshipBy, eu.IndicateCountry, eu.ResidentialHouseNo, eu.ResidentialStreet, eu.ResidentialVillage, eu.ResidentialBarangayCode, eu.ResidentialCityCode, eu.ResidentialProvinceString, eu.ResidentialBarangayString, eu.ResidentialCityString, eu.ResidentialProvinceCode, eu.ResidentialZipCode, eu.PermanentHouseNo, eu.PermanentStreet, eu.PermanentVillage, eu.PermanentBarangayString, eu.PermanentCityString, eu.PermanentProvinceString, eu.PermanentBarangayCode, eu.PermanentCityCode, eu.PermanentProvinceCode, eu.PermanentZipCode, eu.TelephoneNo, eu.MobileNo, eu.EmailAddress,
		)
		

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Updated user")
	}
	sqliteDatabase.Close()
}

func UpdateFamilyBackground(sqlServerDatabase *sql.DB, sqliteDatabase *sql.DB, username string) {
	fi, _ := GetFamilyBackground(sqliteDatabase, username)
	user := sqlServerDatabase.QueryRow(`SELECT id from users where username = @p1`, username)
	var id int
	user.Scan(&id)
	if(fi.UserID.Valid && id != 0) {
		query := `UPDATE employee_family_backgrounds SET
			spouse_firstname = @p2,
			spouse_lastname = @p3,
			spouse_middlename = @p4,
			spouse_extension = @p5,
			spouse_occupation = @p6,
			spouse_employer_business_name = @p7,
			spouse_business_address = @p8,
			spouse_telephone_number = @p9,
			father_firstname = @p10,
			father_lastname = @p11,
			father_middlename = @p12,
			father_extension = @p13,
			mother_lastname = @p14,
			mother_firstname = @p15,
			mother_middlename = @p16
		WHERE user_id = @p1`

		stmt, err := sqlServerDatabase.Prepare(query)
		
		if err != nil {
			log.Fatal(err)
			fmt.Println(err)
		}

		defer stmt.Close()

		_, err = stmt.Exec(
			id, fi.SpouseFirstname, fi.SpouseLastname, fi.SpouseMiddlename, fi.SpouseExtension, fi.SpouseOccupation, fi.SpouseEmployerBusiness, fi.SpouseBusinessAddress, fi.SpouseTelephoneNumber, fi.FatherFirstname, fi.FatherLastname, fi.FatherMiddlename, fi.FatherExtension, fi.MotherLastname, fi.MotherFirstname, fi.MotherMiddlename,
		)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Updated family background")
	} 

	sqliteDatabase.Close()
}

func UpdateSpouseChildren(sqlServerDatabase *sql.DB, sqliteDatabase *sql.DB, username string) {
	children, _ := GetChildren(sqliteDatabase, username)
	user := sqlServerDatabase.QueryRow(`SELECT id from users where username = @p1`, username)
	var id int
	user.Scan(&id)
	if(len(children) > 0 && id != 0) {
	// Begin a transaction
	tx, err := sqlServerDatabase.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// Ensure transaction rollback on error
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			log.Fatal(p)
		} else if err != nil {
			tx.Rollback()
			log.Fatal(err)
		} else {
			err = tx.Commit()
		}
	}()

	deleteQuery := `DELETE FROM employee_spouse_childrens WHERE user_id = @p1`
		_, err = tx.Exec(deleteQuery, id)
		if err != nil {
			log.Fatal(err)
		}

		insertQuery := `INSERT INTO employee_spouse_childrens (employee_id, name, date_of_birth, user_id, created_at, updated_at) VALUES (@p1, @p2, @p3, @p4, @p5, @p6)`

		for _, child := range children {
			_, err = tx.Exec(
				insertQuery,
				child.EmployeeID, child.Name, child.DateOfBirth, id, child.CreatedAt, child.UpdatedAt,
			)
			if err != nil {
				log.Fatal(err)
			}
		}

		fmt.Println("Updated spouse children!")
	} 

	sqliteDatabase.Close()
}