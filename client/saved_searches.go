package client

import (
	"net/http"
	"strings"

	"github.com/splunk/terraform-provider-splunk/client/models"

	"github.com/google/go-querystring/query"
)

func (client *Client) CreateSavedSearches(name, owner, app string, savedSearchObject *models.SavedSearchObject) error {
	values, err := query.Values(savedSearchObject)
	values.Add("name", name)
	if err != nil {
		return err
	}

	endpoint := client.BuildSplunkURL(nil, "servicesNS", owner, app, "saved", "searches")
	resp, err := client.Post(endpoint, values)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (client *Client) ReadSavedSearches(name, owner, app string) (*http.Response, error) {
	endpoint := client.BuildSplunkURL(nil, "servicesNS", owner, app, "saved", "searches", name)
	resp, err := client.Get(endpoint)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (client *Client) UpdateSavedSearches(name string, owner string, app string, savedSearchObject *models.SavedSearchObject, changes []string) error {
	values, err := query.Values(&savedSearchObject)
	if err != nil {
		return err
	}
	for key := range values {
		found := false
		underscoredKey := strings.Replace(key, ".", "_", -1)
		for _, change := range changes {
			if underscoredKey == change {
				found = true
				break
			}
		}
		if !found {
			values.Del(key)
		}
	}
	endpoint := client.BuildSplunkURL(nil, "servicesNS", owner, app, "saved", "searches", name)
	resp, err := client.Post(endpoint, values)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (client *Client) DeleteSavedSearches(name, owner, app string) (*http.Response, error) {
	endpoint := client.BuildSplunkURL(nil, "servicesNS", owner, app, "saved", "searches", name)
	resp, err := client.Delete(endpoint)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// services/saved/searches
func (client *Client) ReadAllSavedSearches() (*http.Response, error) {
	endpoint := client.BuildSplunkURL(nil, "servicesNS", "-", "-", "saved", "searches")
	resp, err := client.Get(endpoint)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
