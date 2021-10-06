package provider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
)

//https://yandex.ru/dev/pdd/doc/reference/dns-list.html

type yandexProvider struct {
	PddToken string
}

type YandexErrorResponse struct {
	Domain  string
	success string
	error   string
}

type YandexResponse struct {
	Domain  string         `json:"domain"`
	Records []YandexRecord `json:"records"`
}

type YandexRecord struct {
	Content   string `json:"content"`
	RecordId  int    `json:"record_id"`
	Fqdn      string `json:"fqdn"`
	Ttl       int    `json:"ttl"`
	Domain    string `json:"domain"`
	Priority  int    `json:"priority"`
	Port      int    `json:"port"`
	Weight    int    `json:"weight"`
	Target    string `json:"target"`
	Subdomain string `json:"subdomain"`
	Type      string `json:"type"`
}

type DnsRecord struct {
	Domain         string
	Host           string `json:"host"`
	Type           string `json:"type"`
	Value          string `json:"value"`
	Ttl            int    `json:"ttl"`
	Subdomain      string `json:"subdomain"`
	ExternalId     string `json:"external_id"`
	AdditionalInfo string `json:"additional_info"`
}

func (p yandexProvider) GetRecords(domain string) ([]DnsRecord, error) {

	url := "https://pddimp.yandex.ru/api2/adman/dns/list?domain=" + domain

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("PddToken", p.getToken())

	res, err := p.getClient().Do(req)

	var yaResp YandexResponse
	err = json.NewDecoder(res.Body).Decode(&yaResp)

	defer res.Body.Close()

	var returnAr []DnsRecord

	for _, r := range yaResp.Records {
		returnAr = append(returnAr, DnsRecord{Value: r.Content, Type: r.Type, Host: r.Domain, Subdomain: r.Subdomain, Ttl: r.Ttl, ExternalId: strconv.Itoa(r.RecordId)})
	}

	return returnAr, err
}

func (p yandexProvider) AddRecord(record DnsRecord) error {
	url := "https://pddimp.yandex.ru/api2/adman/dns/add"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	err := writer.WriteField("domain", record.Host)
	err = writer.WriteField("type", record.Type)
	err = writer.WriteField("content", record.Value)
	err = writer.WriteField("subdomain", record.Subdomain)
	err = writer.WriteField("ttl", strconv.Itoa(record.Ttl))

	err = writer.Close()

	req, err := http.NewRequest("POST", url, payload)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("PddToken", p.getToken())

	_, err = p.getClient().Do(req)

	return err
}

func (p yandexProvider) DeleteRecord(record DnsRecord) error {

	url := fmt.Sprintf("https://pddimp.yandex.ru/api2/adman/dns/del?domain=%s&record_id=%s", record.Host, record.ExternalId)

	req, _ := http.NewRequest("POST", url, nil)
	req.Header.Set("PddToken", p.getToken())

	_, err := p.getClient().Do(req)

	return err
}

func (p yandexProvider) getClient() *http.Client {
	c := &http.Client{}
	return c
}

func (p yandexProvider) getToken() string {
	token := p.PddToken
	if strings.HasPrefix(p.PddToken, "ENV_") {
		token = os.Getenv(p.PddToken)
	}
	return token
}
