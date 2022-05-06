package goexcel

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"path/filepath"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

type Response struct {
	Message string         `json:"message"`
	Status  string         `json:"status"`
	Data    []DataResponse `json:"data"`
}

type DataResponse struct {
	First   int `json:"first"`
	Seconds int `json:"seconds"`
	Thrids  int `json:"thrids"`
}

func GennerateExcel() {
	client := resty.New()

	var resp *resty.Response
	resp, err := client.R().Get("http://localhost:5000")
	if err != nil {
		logrus.Warning(err)
	}
	var res Response
	if err := json.Unmarshal(resp.Body(), &res); err != nil {
		logrus.Errorln("Error json.Unmarshal ->", err)
	}

	f := excelize.NewFile()
	index := f.NewSheet("Sheet1")
	//สร้าง header ของ file excel
	headExcel := map[string]string{"A1": "first", "B1": "seconds", "C1": "thrinds"}
	for k, v := range headExcel {
		f.SetCellValue("Sheet1", k, v)
	}

	//กำหนด data ของเเต่ละ cell ใน file excel
	for i, v := range res.Data {
		i = i + 2
		itos := strconv.Itoa(i)
		f.SetCellValue("Sheet1", "A"+itos, v.First)
		f.SetCellValue("Sheet1", "B"+itos, v.Seconds)
		f.SetCellValue("Sheet1", "C"+itos, v.Thrids)
	}

	f.SetActiveSheet(index)
	// get abs directory
	pathCurrent, err := os.Getwd()
	if err != nil {
		logrus.Errorln("Error get cwd->", err)
	}
	// สร้างชื่อ file
	fileName := fmt.Sprintf("%s.xlsx", strconv.Itoa(int(time.Now().UnixNano())))
	// สร้าง destination  ในการเก็บ file
	pathFile := filepath.Join(pathCurrent, "excel", fileName)
	// save file
	if err := f.SaveAs(pathFile); err != nil {
		fmt.Println(err)
	}
}
