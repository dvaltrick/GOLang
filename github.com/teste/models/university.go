package models

//Define an object to University
type University struct {
	UNITID  string `json:"unitid,omitempty"`
	OPEID   string `json:"opeid,omitempty"`
	OPEID6  string `json:"opeid6,omitempty"`
	INSTNM  string `json:"instnm,omitempty"`
	CITY    string `json:"city,omitempty"`
	STABBR  string `json:"stabbr,omitempty"`
	INSTURL string `json:"insturl,omitempty"`
}
