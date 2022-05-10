package main

type RegisterMeterRes struct {
	Status  bool   `json:"status,omitempty"`
	Message string `json:"messages,omitempty"`
	Data    struct {
		Id        int    `json:"id,omitempty"`
		MeterType string `json:"meter_type,omitempty"`
		MeterNo   string `json:"meter_no,omitempty"`
	} `json:"data,omitempty"`
}

type MeterInfoRes struct {
	Status  bool   `json:"status,omitempty"`
	Message string `json:"messages,omitempty"`
	Data    struct {
		AccountNo string `json:"AccountNo,omitempty"`
		Address   string `json:"Address,omitempty"`
		Charges   struct {
			DemandCharge    int `json:"Demand Charge,omitempty"`
			DemandChargeAll int `json:"Demand Charge-All,omitempty"`
			MeterRent       int `json:"Meter Rent-1P,omitempty"`
		} `json:"Charges,omitempty"`

		CurrentConsumption  string  `json:"CurrentConsumption,omitempty"`
		CustomerName        string  `json:"CustomerName,omitempty"`
		CustomerTel         string  `json:"CustomerTel,omitempty"`
		MaxPower            int     `json:"MaxPower,omitempty"`
		Message             string  `json:"Message,omitempty"`
		MeterNo             string  `json:"MeterNo,omitempty"`
		RemainingBalance    string  `json:"RemainingBalance,omitempty"`
		Status              bool    `json:"Status,omitempty"`
		MeterType           string  `json:"meter_type,omitempty"`
		CurrentConsumptions Balance `json:"CurrentConsumptions,omitempty"`
		RemainingBalances   Balance `json:"RemainingBalances,omitempty"`
		NearOfficeAddress   string  `json:"near_office_address,omitempty"`
		NearOfficeLat       string  `json:"near_office_lat,omitempty"`
		NearOfficeLong      string  `json:"near_office_long,omitempty"`
		ImportAddressTitle  string  `json:"import_address_title,omitempty"`
	} `json:"data,omitempty"`
}

type Balance struct {
	Amount string `json:"amount,omitempty"`
	Unit   string `json:"unit,omitempty"`
	Date   string `json:"date,omitempty"`
	Time   string `json:"time,omitempty"`
}
