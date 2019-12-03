package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"./messengers"
	_ "github.com/lib/pq"
)

var workOptions = map[string]bool{}

func stringInSlice(stringsMap map[string]bool, list []string) map[string]bool {
	for _, listElement := range list {
		for stringKey, _ := range stringsMap {
			if stringKey == listElement {
				stringsMap[stringKey] = true
			}
		}
	}
	return stringsMap
}

func init() {
	// get program arguments
	argsWithProg := os.Args

	// set default triggers for
	workOptions = map[string]bool{
		"tg":             false,
		"slack":          false,
		"debug":          false,
		"close":          false,
		"withoutCron":    false,
		"loadUsers":      false,
		"loadProjects":   false,
		"install":        false,
		"loadVersions":   false,
		"updateVersions": false,
		"showFeatures":   false,
		"updateProjects": false,
	}

	//debug
	result := stringInSlice(workOptions, argsWithProg)
	fmt.Println(result)
}

func hasOption(key string) bool {
	return workOptions[key] == true
}

func main() {
	if workOptions["close"] {
		return
	}

	if !workOptions["withoutCron"] {
		fmt.Println("Prepare cron jobs")
		// add cron jobs
		//cronJobs := cron.New()
		// every 12 hours at 23:00:00
		//cronJobs.AddFunc("0 0 23 * * *", jiraSystem.LoadProjectsList)
		// every 12 hours at 23:30:00
		//cronJobs.AddFunc("0 30 23 * * *", jiraSystem.LoadProjectsVersionsList)
		//go cronJobs.Start()

	}

	// configure jira API
	time.Sleep(2 * time.Second)

	//jiraSystem.Initialize(JiraClient)

	/*
	if workOptions["loadProjects"] {
		go jiraSystem.LoadProjects(JiraClient)
	}
	if workOptions["loadUsers"] {
		go jiraSystem.LoadUsers(JiraClient)
	}
	if workOptions["loadVersions"] {
		go jiraSystem.LoadVersions()
	}
	if workOptions["updateVersions"] {
		go jiraSystem.UpdateVersions()
	}
	if workOptions["updateProjects"] {
		go jiraSystem.UpdateProjects()
	}
	*/

	/*
	if workOptions["install"] {
		go jiraSystem.Install()
		fmt.Println("Wait when install will be done and then press 'q' to quit")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			exit := scanner.Text()
			if exit == "q" {
				break
			} else {
				fmt.Println("Press 'q' to quit")
			}
		}
	}
	*/

	if !workOptions["withoutMessengers"] {
		messenger.Initialize()
	}

	fmt.Println("Wait when install will be done and then press 'q' to quit")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		exit := scanner.Text()
		if exit == "q" {
			break
		} else {
			fmt.Println("Press 'q' to quit")
		}
	}
}
