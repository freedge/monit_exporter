package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/prometheus/log"
)

// EjabberdConnectedUserInfo is an item of the list retured by the /api/connected_users_info service
type EjabberdConnectedUserInfo struct {
	jid        string
	connection string
	ip         string
	port       int
	priority   int
	node       string
	uptime     int
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
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Unable to read ejabberd status")
		return nil, err
	}
	var ejabberd []EjabberdConnectedUserInfo
	json.Unmarshal(data, &ejabberd)
	return ejabberd, nil
}
