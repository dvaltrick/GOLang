package services

import (
	"database/sql"
	"fmt"
	
	_ "github.com/lib/pq"
	"github.com/teste/models"
)

//Parameters to connect with a PostgreSql database
const (
	dbuser = "postgres"
	dbpwd = "admin"
	dbname = "goteste" 
)

//Function to save a register of a type University at database
func Grava(university models.University){
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbuser, dbpwd, dbname)
	db, err := sql.Open("postgres", dbinfo)
	if(err != nil){
		fmt.Println(err)
	}
	
	defer db.Close()
	
	fmt.Println("# Inserting values")
	
	var teste int
	err = db.QueryRow("INSERT INTO university(unitid, opeid, opeid6, instnm, city, stabbr, insturl) "+
	                                    "VALUES("+university.UNITID+","+
												"'"+university.OPEID+"'"+","+
												"'"+university.OPEID6+"'"+","+
												"'"+university.INSTNM+"'"+","+
												"'"+university.CITY+"'"+","+
												"'"+university.STABBR+"'"+","+
												"'"+university.INSTURL+"'"+") ;").Scan(&teste)
	
	if(err != nil){
		fmt.Println(err)
	}
}


//Function to get all registers of universities in database
func GetAll() []models.University{
	var uni models.University
	var uniList []models.University
	
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbuser, dbpwd, dbname)
	db, err := sql.Open("postgres", dbinfo)
	if(err != nil){
		fmt.Println(err)
	}
	
	defer db.Close()

	fmt.Println("# Querying")
	rows, err := db.Query("SELECT * FROM university")
	if(err != nil){
		fmt.Println(err)
	}

	for rows.Next() {
		err = rows.Scan(&uni.UNITID, &uni.OPEID, &uni.OPEID6, &uni.INSTNM, &uni.CITY, &uni.STABBR, &uni.INSTURL)
		if(err != nil){
			fmt.Println(err)
		}
		uniList = append(uniList,uni)
	}

	return uniList
}

//Function to get all registers of universities in database that corresponding to 
//filters received from frontend
func Get(uniParam models.University) []models.University{
	var uni models.University
	var uniList []models.University

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbuser, dbpwd, dbname)
	db, err := sql.Open("postgres", dbinfo)
	if(err != nil){
		fmt.Println(err)
	}
	
	defer db.Close()

	rows, err := db.Query("SELECT A.* FROM university A" +
						  " WHERE ('"+uniParam.UNITID+"' = '' OR A.unitid = '"+uniParam.UNITID+"') " +
						  "   AND (Upper(A.instnm) like Upper('%"+uniParam.INSTNM+"%')) " +
						  "   AND (Upper(A.city) like Upper('%"+uniParam.CITY+"%')) " +
						  "   AND (Upper(A.stabbr) like Upper('%"+uniParam.STABBR+"%')) " +
						  " ORDER BY stabbr, city, instnm" )
	if(err != nil){
		fmt.Println(err)
	}

	for rows.Next() {
		err = rows.Scan(&uni.UNITID, &uni.OPEID, &uni.OPEID6, &uni.INSTNM, &uni.CITY, &uni.STABBR, &uni.INSTURL)
		if(err != nil){
			fmt.Println(err)
		}
		uniList = append(uniList,uni)
	}

	return uniList
}

