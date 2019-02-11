package guru

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

type UMLSClient struct {
	client               *http.Client
	ticketGrantingTicket string
}

func NewUMLSClient(username, password string) (UMLSClient, error) {
	var (
		client *http.Client
		umls   UMLSClient
	)

	fmt.Println("getting umls ticket")
	resp, err := client.Post("https://utslogin.nlm.nih.gov/cas/v1/tickets", "application/x-www-form-urlencoded", bytes.NewBufferString(fmt.Sprintf("username=%s&password=%s", username, password)))
	if err != nil {
		return umls, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return umls, err
	}
	re := regexp.MustCompile(`action=".*(?P<Ticket>TGT-.*-cas)"`)
	ticketGrantingTicket := re.FindStringSubmatch(string(b))[1]

	return UMLSClient{
		client:               client,
		ticketGrantingTicket: ticketGrantingTicket,
	}, nil
}

func (c UMLSClient) Preferred(cui string) (string, error) {
	// Request service ticket.
	resp, err := c.client.Post(fmt.Sprintf("https://utslogin.nlm.nih.gov/cas/v1/tickets/%s", c.ticketGrantingTicket), "application/x-www-form-urlencoded", bytes.NewBufferString(fmt.Sprintf("service=%s", "http://umlsks.nlm.nih.gov")))
	if err != nil {
		return "", err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
		return "", err
	}

	ticket := string(b)
	resp, err = c.client.Get(fmt.Sprintf("https://uts-ws.nlm.nih.gov/rest/content/current/CUI/%s/atoms/preferred?ticket=%s", cui, ticket))
	if err != nil {
		panic(err)
		return "", err
	}

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	var r map[string]interface{}
	err = json.Unmarshal(b, &r)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	s := r["result"].(map[string]interface{})["name"].(string)
	return s, nil
}
