package services

import (
	"encoding/csv"
	"os"
	"fmt"

	"github.com/teste/models"
)

//Function to import the csv file to database
func ImportFromFile(){
	//The file need to be in the root directory
	file, err := os.Open("novo.csv")
	if err != nil {
        fmt.Println("Error", err)
        return
    }
	defer file.Close()
	
	reader := csv.NewReader(file)
	record, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error", err)
	}
	
	for value:= range record{
		var uni models.University
		uni.UNITID = record[value][0]
		uni.OPEID = record[value][1] 
		uni.OPEID6 = record[value][2] 
		uni.INSTNM = record[value][3] 
		uni.CITY = record[value][4] 
		uni.STABBR = record[value][5] 
		uni.INSTURL = record[value][6] 

		fmt.Println(uni.UNITID, uni.OPEID, uni.OPEID6, uni.INSTNM, uni.CITY, uni.STABBR, uni.INSTURL)
		Grava(uni)
	}

}