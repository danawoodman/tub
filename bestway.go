package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	GIZWITS_APP_ID = "98754e684ec045528b073876c34c7348"

	POWER_OFF = 0
	POWER_ON  = 1

	JETS_OFF  = 0
	JETS_LOW  = 51
	JETS_HIGH = 100

	SCREEN_LOCK   = 1
	SCREEN_UNLOCK = 0

	FILTER_OFF = 0
	FILTER_ON  = 2

	HEAT_OFF = 0
	HEAT_ON  = 4
)

type (
	BestWay struct {
		gizwits *Gizwits
		token   string
	}

	BestWayLoginResponse struct {
		Token    string `json:"token"`
		UID      string `json:"uid"`
		ExpireAt int64  `json:"expire_at"`
	}

	BestWayDeviceState struct {
		Power              int `json:"power"`
		SetTemperature     int `json:"Tset"`
		CurrentTemperature int `json:"Tnow"`
		TemperatureUnit    int `json:"Tunit"`
		Filter             int `json:"filter"`
		Heat               int `json:"heat"`
		Jets               int `json:"wave"`
		ScreenLock         int `json:"bit2"`
		// "option7": 10260,
		// "option4": 0,
		// "option5": 27187,
		// "option2": 0,
		// "option3": 30,
		// "option0": 59940,
		// "option1": 59940,
		// "word3": 0,
		// "bit7": 0,
		// "word2": 166,
		// "option6": 10244,
		// "ver": 1174,
		// "E22": 0,
		// "jet": 0,
		// "E23": 0,
		// "E28": 0,
		// "bit6": 0,
		// "bit5": 0,
		// "bit4": 0,
		// "bit3": 0,
		// "word5": 0,
		// "word4": 0,
		// "word7": 39,
		// "word6": 0,
		// "word1": 0,
		// "word0": 0,
		// "E29": 0,
		// "E19": 0,
		// "E18": 0,
		// "E32": 0,
		// "E31": 0,
		// "E30": 0,
		// "E11": 0,
		// "E10": 0,
		// "E13": 0,
		// "E12": 0,
		// "E15": 0,
		// "E14": 0,
		// "E17": 0,
		// "E16": 0,
		// "E24": 0,
		// "E25": 0,
		// "E26": 0,
		// "E27": 0,
		// "E20": 0,
		// "E21": 0,
		// "E08": 0,
		// "E09": 0,
		// "E06": 0,
		// "E07": 0,
		// "E04": 0,
		// "E05": 0,
		// "E02": 0,
		// "E03": 0,
		// "E01": 0,
	}

	BestWayDeviceStatusResponse struct {
		DeviceID    string             `json:"did"`
		UpdatedAt   int                `json:"updated_at"`
		DeviceState BestWayDeviceState `json:"attr"`
	}

	BestWayPowerPayload struct {
		// power 0 (off) or 1 (on)
		Power int `json:"power"`
	}

	BestWayHeatPayload struct {
		// heat 0 (off) or 4 (on)
		Heat int `json:"heat"`
	}

	BestWayJetsPayload struct {
		// jet: 0, 51 (high), 2 (low)
		Jets int `json:"wave"`
	}

	BestWayFilterPayload struct {
		// filter 0 (off) or 2 (on)
		Filter int `json:"filter"`
	}

	BestWayScreenLockPayload struct {
		// screen lock: 0 (off) or 1 (on)
		ScreenLock int `json:"bit2"`
	}

	BestWayTemperaturePayload struct {
		// temperature value
		Temperature int `json:"Tset"`
	}

	BestWayTemperatureUnitPayload struct {
		// temperature units: 0 (F), 1 (C)
		TemperatureUnit int `json:"Tunit"`
	}

	BestWayDevice struct {
		Name        string `json:"dev_alias"`
		ID          string `json:"did"`
		Online      bool   `json:"is_online"`
		ProductName string `json:"product_name"`
		// "protoc": 3,
		// "ws_port": 8080,
		// "port_s": 8883,
		// "gw_did": null,
		// "host": "usm2m.gizwits.com",
		// "sleep_duration": 3600,
		// "port": 1883,
		// "mcu_soft_version": "D4H90227",
		// "product_key": "d3ac9226d983470284b5d133cf4fd6b4",
		// "state_last_timestamp": 1731463263,
		// "role": "owner",
		// "is_sandbox": false,
		// "type": "normal",
		// "is_disabled": false,
		// "mcu_hard_version": "P4960011",
		// "wifi_soft_version": "04X3000B",
		// "mesh_id": null,
		// "dev_label": [],
		// "wss_port": 8880,
		// "remark": "25",
		// "mac": "5432044c9e2c",
		// "passcode": "UOKOUIPIGM",
		// "wifi_hard_version": "0ESP32C3",
		// "is_low_power": false
	}

	BestWayListDevicesResponse struct {
		Devices []BestWayDevice `json:"devices"`
	}
)

func NewBestWay() *BestWay {
	return &BestWay{
		gizwits: NewGizwits(GIZWITS_APP_ID),
	}
}

func (self *BestWay) Login(username, password string) (*BestWayLoginResponse, error) {
	path := "/login"
	payload := map[string]string{"username": username, "password": password}

	resp, err := self.gizwits.Request("POST", path, payload, nil)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	var data BestWayLoginResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	self.SetToken(data.Token)

	return &data, nil
}

func (self *BestWay) SetToken(token string) {
	self.token = token
}

func (self *BestWay) ListDevices() (*BestWayListDevicesResponse, error) {
	resp, err := self.gizwits.AuthRequest(
		"GET",
		"/bindings",
		nil,
		self.token,
	)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	var data BestWayListDevicesResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &data, nil
}

func (self *BestWay) GetDeviceStatus(id string) (*BestWayDeviceStatusResponse, error) {
	if self.token == "" {
		return nil, fmt.Errorf("no token")
	}

	path := fmt.Sprintf("/devdata/%s/latest", id)
	resp, err := self.gizwits.AuthRequest("GET", path, nil, self.token)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	var data BestWayDeviceStatusResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &data, nil
}

func (self *BestWay) SetTemp(deviceID string, temp int) (*http.Response, error) {
	payload := BestWayTemperaturePayload{Temperature: temp}
	return self.gizwits.ControlRequest(deviceID, payload, self.token)
}

func (self *BestWay) SetPower(deviceID string, power bool) (*http.Response, error) {
	powerInt := POWER_OFF
	if power {
		powerInt = POWER_ON
	}
	payload := BestWayPowerPayload{Power: powerInt}
	return self.gizwits.ControlRequest(deviceID, payload, self.token)
}

func (self *BestWay) SetHeat(deviceID string, heat bool) (*http.Response, error) {
	heatInt := HEAT_OFF
	if heat {
		heatInt = HEAT_ON
	}
	payload := BestWayHeatPayload{Heat: heatInt}
	return self.gizwits.ControlRequest(deviceID, payload, self.token)
}

func (self *BestWay) SetJets(deviceID string, jets int) (*http.Response, error) {
	payload := BestWayJetsPayload{Jets: jets}
	return self.gizwits.ControlRequest(deviceID, payload, self.token)
}

func (self *BestWay) SetFilter(deviceID string, filter bool) (*http.Response, error) {
	filterInt := FILTER_OFF
	if filter {
		filterInt = FILTER_ON
	}
	payload := BestWayFilterPayload{Filter: filterInt}
	return self.gizwits.ControlRequest(deviceID, payload, self.token)
}

func (self *BestWay) SetScreenLock(deviceID string, screenLock bool) (*http.Response, error) {
	screenLockInt := SCREEN_LOCK
	if screenLock {
		screenLockInt = SCREEN_UNLOCK
	}
	payload := BestWayScreenLockPayload{ScreenLock: screenLockInt}
	return self.gizwits.ControlRequest(deviceID, payload, self.token)
}
