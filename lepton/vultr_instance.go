package lepton

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
)

// CreateInstance - Creates instance on Digital Ocean Platform
func (v *Vultr) CreateInstance(ctx *Context) error {
	c := ctx.config

	// you may poll /v1/server/list?SUBID=<SUBID> and check that the "status" field is set to "active"

	createURL := "https://api.vultr.com/v1/server/create"

	token := os.Getenv("TOKEN")

	urlData := url.Values{}
	urlData.Set("DCID", "1")

	// this is the instance size
	// TODO
	urlData.Set("VPSPLANID", "201")

	// id for snapshot
	urlData.Set("OSID", "164")
	urlData.Set("SNAPSHOTID", c.CloudConfig.ImageName)

	req, err := http.NewRequest("POST", createURL, strings.NewReader(urlData.Encode()))
	req.Header.Set("API-Key", token)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	return nil
}

// GetInstanceByID returns the instance with the id passed by argument if it exists
func (v *Vultr) GetInstanceByID(ctx *Context, id string) (*CloudInstance, error) {
	return nil, errors.New("un-implemented")
}

// GetInstances return all instances on Vultr
// TODO
func (v *Vultr) GetInstances(ctx *Context) ([]CloudInstance, error) {
	return nil, errors.New("un-implemented")
}

// ListInstances lists instances on v
func (v *Vultr) ListInstances(ctx *Context) error {

	client := http.Client{}
	req, err := http.NewRequest("GET", "https://api.vultr.com/v1/server/list", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	token := os.Getenv("TOKEN")

	req.Header.Set("API-Key", token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var data map[string]vultrServer

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Status", "Created", "Private Ips", "Public Ips"})
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor})
	table.SetRowLine(true)

	for _, image := range data {
		var row []string
		row = append(row, image.Name)
		row = append(row, image.Status)
		row = append(row, image.CreatedAt)
		row = append(row, image.PrivateIP)
		row = append(row, image.PublicIP)
		table.Append(row)
	}

	table.Render()

	return nil
}

// DeleteInstance deletes instance from v
func (v *Vultr) DeleteInstance(ctx *Context, instanceID string) error {
	destroyInstanceURL := "https://api.vultr.com/v1/server/destroy"

	token := os.Getenv("TOKEN")

	urlData := url.Values{}
	urlData.Set("SUBID", instanceID)

	req, err := http.NewRequest("POST", destroyInstanceURL, strings.NewReader(urlData.Encode()))
	req.Header.Set("API-Key", token)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	return nil
}

// StartInstance starts an instance in v
func (v *Vultr) StartInstance(ctx *Context, instanceID string) error {
	startInstanceURL := "https://api.vultr.com/v1/server/start"

	token := os.Getenv("TOKEN")

	urlData := url.Values{}
	urlData.Set("SUBID", instanceID)

	req, err := http.NewRequest("POST", startInstanceURL, strings.NewReader(urlData.Encode()))
	req.Header.Set("API-Key", token)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	return nil
}

// StopInstance halts instance from v
func (v *Vultr) StopInstance(ctx *Context, instanceID string) error {
	haltInstanceURL := "https://api.vultr.com/v1/server/halt"

	token := os.Getenv("TOKEN")

	urlData := url.Values{}
	urlData.Set("SUBID", instanceID)

	req, err := http.NewRequest("POST", haltInstanceURL, strings.NewReader(urlData.Encode()))
	req.Header.Set("API-Key", token)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	return nil
}

// PrintInstanceLogs writes instance logs to console
func (v *Vultr) PrintInstanceLogs(ctx *Context, instancename string, watch bool) error {
	l, err := v.GetInstanceLogs(ctx, instancename)
	if err != nil {
		return err
	}
	fmt.Printf(l)
	return nil
}

// GetInstanceLogs gets instance related logs
func (v *Vultr) GetInstanceLogs(ctx *Context, instancename string) (string, error) {
	return "", nil
}
