package main

import "fmt"

type idea struct {
	title string
}

func (i idea) pitch() {
	fmt.Println("Project is about - ", i.title)
}

type project struct {
	idea
	language string
}

func (pr project) explainExecution() {
	fmt.Println("We are going to build - ", pr.title, " using ", pr.language)
}

func main() {
	pr := project{
		idea: idea{
			title: "Ticket Reservation System",
		},
		language: "Golang",
	}
	pr.pitch()
	pr.explainExecution()
}
