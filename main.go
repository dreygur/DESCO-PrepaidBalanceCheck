package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	// Register the Meter
	regMeterInfo, err := registerMeter("13130589")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(regMeterInfo)

	meterId := strconv.Itoa(regMeterInfo.Data.Id)
	// Get the Meter Info
	meterInfo, err := getMeterInfo(meterId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(meterInfo)

	// Delete the Meter
	deletedMeterInfo, err := deleteMeter(meterId)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(deletedMeterInfo.Message)
}

func registerMeter(account string) (*RegisterMeterRes, error) {
	res, err := networkHandler("POST", "meter-info", "meter_type=prepaid&account_no="+account)
	if err != nil {
		return nil, err
	}

	var registerMeterRes *RegisterMeterRes
	err = json.NewDecoder(res.Body).Decode(&registerMeterRes)
	if err != nil {
		return nil, err
	}
	return registerMeterRes, nil
}

func getMeterInfo(id string) (*MeterInfoRes, error) {
	res, err := networkHandler("GET", "meter-info/"+id, "")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var meterInfo *MeterInfoRes
	err = json.NewDecoder(res.Body).Decode(&meterInfo)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return meterInfo, nil
}

func deleteMeter(id string) (*RegisterMeterRes, error) {
	res, err := networkHandler("DELETE", "meter-info/"+id, "")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var meterInfo *RegisterMeterRes
	err = json.NewDecoder(res.Body).Decode(&meterInfo)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return meterInfo, nil
}

func networkHandler(reqType, path, reqBody string) (*http.Response, error) {
	url := "https://descoapp.sslwireless.com/api/v1/" + path
	method := reqType

	payload := strings.NewReader(reqBody)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	req.Header.Add("Host", "descoapp.sslwireless.com")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", "38")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("User-Agent", "okhttp/3.14.9")
	req.Header.Add("Connection", "close")
	req.Header.Add("Authorization", "Bearer wrnj3YQuZ8hiRZ8lpglxwUBWdEgQxAJCSncjYpw5MakfsTwBCe983oZZD4pQ")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// defer res.Body.Close()

	return res, nil
}
