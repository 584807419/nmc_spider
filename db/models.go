package db

type Location struct {
	Id        int
	Stationid string
	Country   string
	Province  string
	City      string
	Valid     int
}

type Province struct {
	Id    int
	Name  string
	Abbr  string
	Valid int
}

type EverydayData struct {
	Id                int
	Date              string
	Day_info          string
	Day_temperature   string
	Day_direct        string
	Day_power         string
	Night_info        string
	Night_temperature string
	Night_direct      string
	Night_power       string
}
