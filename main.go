package main

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
	"os"

	"github.com/digitalocean/godo"
	"golang.org/x/oauth2"
)

func main() {
	accessToken := os.Getenv("DO_KEY")
	if accessToken == "" {
		log.Fatal("Usage: DO_KEY environment variable must be set.")
	}

	peerTag := os.Getenv("DO_TAG")
	if peerTag == "" {
		log.Fatal("Usage: DO_TAG environment variable must be set.")
	}

	tmpl, _ := template.New("test").Parse(`## drophosts ##
{{range .}}{{.PrivateIPv4}} {{.Name}}.kubelocal
{{end}}## drophosts ##`)

	oauthClient := oauth2.NewClient(oauth2.NoContext, oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken}))
	client := godo.NewClient(oauthClient)
	droplets, _ := DropletListTags(client.Droplets, peerTag)

	original, err := ioutil.ReadFile("/etc/hosts")
	if err != nil {
		log.Fatal(err)
	}
	var doc bytes.Buffer
	tmpl.Execute(&doc, droplets)

	output := UpdateHosts(string(original), doc.String())
	ioutil.WriteFile("/etc/hosts", []byte(output), 0644)
}

// DropletListTags paginates through the digitalocean API to return a list of
// all droplets with the given tag
func DropletListTags(ds godo.DropletsService, tag string) ([]godo.Droplet, error) {
	// create a list to hold our droplets
	list := []godo.Droplet{}

	// create options. initially, these will be blank
	opt := &godo.ListOptions{}
	for {
		droplets, resp, err := ds.ListByTag(tag, opt)

		if err != nil {
			return nil, err
		}

		// append the current page's droplets to our list
		for _, d := range droplets {
			list = append(list, d)
		}

		// if we are at the last page, break out the for loop
		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}

		page, err := resp.Links.CurrentPage()
		if err != nil {
			return nil, err
		}

		// set the page we want for the next request
		opt.Page = page + 1
	}

	return list, nil
}
