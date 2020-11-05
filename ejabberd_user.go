package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/prometheus/common/log"
)

// EjabberdConnectedUserInfo is an item of the list retured by the /api/connected_users_info service
type EjabberdConnectedUserInfo struct {
	Jid        string
	Connection string
	Ip         string
	Port       int
	Priority   int
	Node       string
	Uptime     int
}

// FetchAndParseEjabberdStatus return the list of connected users from Jabber
func FetchAndParseEjabberdStatus(c *Config) ([]EjabberdConnectedUserInfo, error) {
	req, err := http.NewRequest("POST", c.ejabberd_uri+"/api/connected_users_info", bytes.NewBufferString("{}"))
	if err != nil {
		log.Errorf("Unable to create request: %v", err)
		return nil, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error("Unable to fetch ejabberd status")
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Unable to read ejabberd status")
		return nil, err
	}
	var ejabberd []EjabberdConnectedUserInfo
	json.Unmarshal(data, &ejabberd)
	return ejabberd, nil
}
