package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	cloudflare "github.com/cloudflare/cloudflare-go"
)

const (
	publicIPApi = "https://api.ipify.org"
)

var errNotOk = errors.New("status code is not 200")
func task() {
	config, err := loadConfigFromEnv()
	if err != nil {
		fmt.Printf("Could not run task.\n%s\n", err.Error())
		return
	}

	err = executeTask(config)
	if err != nil {
		fmt.Printf("Could not run task.\n%s\n", err.Error())
	}
}

func executeTask(config config) error {
	publicIP, err := getPublicIP()
	if err != nil {
		return err
	}

	fmt.Printf("Current public IP is %s\n", publicIP)

	api, err := cloudflare.NewWithAPIToken(config.APIKey)
	if err != nil {
		return err
	}

	zoneID, err := api.ZoneIDByName(config.Domain)
	if err != nil {
		return err
	}

	records, err := api.DNSRecords(zoneID, cloudflare.DNSRecord{})
	if err != nil {
		return err
	}

	record, err := getDNSRecord(records, config.Domain, config.DNSRecordType)
	if err != nil {
		return err
	}

	fmt.Printf("Current DNS record is %s\n", record.Content)

	if record.Content == publicIP {
		fmt.Printf("No need to change.\n")
		return nil
	}

	record.Content = publicIP
	err = api.UpdateDNSRecord(zoneID, record.ID, record)
	if err != nil {
		return err
	}

	fmt.Printf("DNS record updated.\n")
	return nil
}

func getPublicIP() (string, error) {
	res, err := http.Get(publicIPApi)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return "", errNotOk
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func getDNSRecord(records []cloudflare.DNSRecord, domain string, dnsRecordType string) (cloudflare.DNSRecord, error) {
	for _, record := range records {
		if record.Name == domain && record.Type == dnsRecordType {
			return record, nil
		}
	}

	return cloudflare.DNSRecord{}, errors.New("DNS record not found")
}
