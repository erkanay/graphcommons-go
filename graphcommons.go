package graphcommons

import (
	"io"
	"bytes"
	"net/url"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

var (
	ContentType = "applications/json"
	ApiUrl      = "https://graphcommons.com/api/v1/"
)

type Request struct {
	*http.Request
	Params url.Values
}

type Graph struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Status      int      `json:"status"`
	Subtitle    string   `json:"subtitle"`
	Signals     []Signal `json:"signals"`
}

type Signal struct {
	Action   string `json:"action"`
	FromName string `json:"from_name"`
	FromType string `json:"from_type"`
	ToName   string `json:"to_name"`
	ToType   string `json:"to_type"`
	Name     string `json:"name"`
	Weight   int    `json:"weight"`
}

type Client struct {
	httpClient  *http.Client
	Url         string
	ApiKey      string
	ContentType string
}

func GraphCommons(ApiKey string) (*Client, error) {
	return &Client{
		httpClient:  &http.Client{},
		Url:         ApiUrl,
		ApiKey:      ApiKey,
		ContentType: ContentType,
	}, nil
}

func EncodeData(body interface{}) *bytes.Buffer {
	data, _ := json.Marshal(body)
	return bytes.NewBuffer(data)
}

func DecodeData(res *http.Response) string {
	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)
	return string(data)
}

func (c *Client) MakeRequest(method, location string, body io.Reader) (*http.Response, error) {
	urlStr := ApiUrl + location
	req, err := http.NewRequest(method, urlStr, body)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authentication", c.ApiKey)
	req.Header.Add("Content-Type", c.ContentType)
	sURL, _ := url.Parse(urlStr)
	req.URL = sURL
	client := c.httpClient
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) Status() string {
	res, err := c.MakeRequest("GET", "status/", nil)
	if err != nil {
		panic(err)
	}
	return DecodeData(res)
}

func (c *Client) Graphs(id string) string {
	res, err := c.MakeRequest("GET", "graphs/" + id, nil)
	if err != nil {
		panic(err)
	}
	return DecodeData(res)
}

func (c *Client) CreateGraph(body interface{}) string {
	res, err := c.MakeRequest("POST", "graphs/", EncodeData(body))
	if err != nil {
		panic(err)
	}
	return DecodeData(res)
}

func (c *Client) UpdateGraph(id string, body interface{}) string {
	res, err := c.MakeRequest("PUT", "graphs/" + id + "/add", EncodeData(body))
	if err != nil {
		panic(err)
	}
	return DecodeData(res)
}

func (c *Client) GetGraphTypes(id string) string {
	res, err := c.MakeRequest("GET", "graphs/" + id + "/types", nil)
	if err != nil {
		panic(err)
	}
	return DecodeData(res)
}

func (c *Client) GetGraphEdges(id, params string) string {
	res, err := c.MakeRequest("GET", "graphs/" + id + "/edges?" + params, nil)
	if err != nil {
		panic(err)
	}
	return DecodeData(res)
}

func (c *Client) GetGraphPaths(id, params string) string {
	res, err := c.MakeRequest("GET", "graphs/" + id + "/paths?" + params, nil)
	if err != nil {
		panic(err)
	}
	return DecodeData(res)
}

func (c *Client) CollabFilterGraphs(id, params string) string {
	res, err := c.MakeRequest("GET", "graphs/" + id + "/collab_filter?" + params, nil)
	if err != nil {
		panic(err)
	}
	return DecodeData(res)
}

func (c *Client) SearchGraphs(id, params string) string {
	res, err := c.MakeRequest("GET", "graphs/search?" + params, nil)
	if err != nil {
		panic(err)
	}
	return DecodeData(res)
}

func (c *Client) DeleteGraph(id string) string {
	res, err := c.MakeRequest("DELETE", "graphs/" + id, nil)
	if err != nil {
		panic(err)
	}
	return DecodeData(res)
}

func (c *Client) GetNode(id string) string {
	res, err := c.MakeRequest("GET", "nodes/" + id, nil)
	if err != nil {
		panic(err)
	}
	return DecodeData(res)
}

func (c *Client) SearchNodes(id, params string) string {
	res, err := c.MakeRequest("GET", "nodes/search?" + params, nil)
	if err != nil {
		panic(err)
	}
	return DecodeData(res)
}

func (c *Client) GetHub(id string) string {
	res, err := c.MakeRequest("GET", "hubs/" + id, nil)
	if err != nil {
		panic(err)
	}
	return DecodeData(res)
}

func (c *Client) GetHubTypes(id string) string {
	res, err := c.MakeRequest("GET", "hubs/" + id + "/types", nil)
	if err != nil {
		panic(err)
	}
	return DecodeData(res)
}

func (c *Client) GetHubPaths(id, params string) string {
	res, err := c.MakeRequest("GET", "hubs/" + id + "/paths?" + params, nil)
	if err != nil {
		panic(err)
	}
	return DecodeData(res)
}

func (c *Client) CollabFilterHubs(id, params string) string {
	res, err := c.MakeRequest("GET", "hubs/" + id + "/collab_filter?" + params, nil)
	if err != nil {
		panic(err)
	}
	return DecodeData(res)
}

func (c *Client) Search(params string) string {
	res, err := c.MakeRequest("GET", "search?" + params, nil)
	if err != nil {
		panic(err)
	}
	return DecodeData(res)
}
