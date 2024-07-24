package main

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

const DEV_COMPOSE_FILE string = "../dev/dev.docker-compose.yaml"
const DEV_SERVER_WORKFLOW_NAME string = "sitenote-server-ci"
const DEV_FRONTEND_WORKFLOW_NAME string = "sitenote-frontend-ci"

const PROD_COMPOSE_FILE string = "../prod/dev.docker-compose.yaml"
const PROD_SERVER_WORKFLOW_NAME string = "sitenote-server-ci-prod"
const PROD_FRONTEND_WORKFLOW_NAME string = "sitenote-frontend-ci-prod"

type githubWebhookRequest struct {
	Action   string   `json:"action"`
	WorkFlow Workflow `json:"workflow"`
}

type Workflow struct {
	Badge_url  string `json:"badge_url"`
	Created_at string `json:"created_at"`
	Html_url   string `json:"html_url"`
	Id         string `json:"id"`
	Name       string `json:"name"`
	Node_id    string `json:"node_id"`
	Path       string `json:"path"`
	State      string `json:"state"`
	Updated_at string `json:updated_at`
	Url        string `json:url`
}

func main() {
	r := gin.Default()
	r.POST("/webhook", deploy)
	r.Run("localhost:9050")
}

func deploy(c *gin.Context) {

	githubReq := githubWebhookRequest{}

	c.BindJSON(&githubReq)

	if githubReq.WorkFlow.Name == PROD_FRONTEND_WORKFLOW_NAME || githubReq.WorkFlow.Name == PROD_SERVER_WORKFLOW_NAME && githubReq.Action == "completed" {
		deployProdEnvironment(PROD_COMPOSE_FILE)
	} else if githubReq.Action == "completed" {
		deployDevEnvironment(DEV_COMPOSE_FILE)
	}

	c.String(http.StatusOK, "success")
}

func deployProdEnvironment(composeFile string) {
	composeDown(composeFile)
	removeImages(composeFile)
	getLatestImages(composeFile)
	composeUp(composeFile)
}

func deployDevEnvironment(composeFile string) {
	composeDown(composeFile)
	removeImages(composeFile)
	getLatestImages(composeFile)
	composeUp(composeFile)
}

func composeDown(composeFile string) {
	fmt.Println("COMPOSING DOWN")
	cmd := exec.Command("docker", "compose", "-f", composeFile, "down")
	commandOutput, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	successMessage := string(commandOutput)
	fmt.Println(successMessage)

}

func removeImages(composeFile string) {

	fmt.Println("REMOVING IMAGES")
	cmd := exec.Command("docker", "compose", "-f", composeFile, "rm", "-f")
	commandOutput, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	successMessage := string(commandOutput)
	fmt.Println(successMessage)

}

func getLatestImages(composeFile string) {
	fmt.Println("GETTING LATEST IMAGES")
	cmd := exec.Command("docker", "compose", "-f", composeFile, "pull")
	commandOutput, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	successMessage := string(commandOutput)
	fmt.Println(successMessage)
}

func composeUp(composeFile string) {
	fmt.Println("COMPOSING UP")
	cmd := exec.Command("docker", "compose", "-f", composeFile, "up", "-d")
	commandOutput, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
		return
	}
	successMessage := string(commandOutput)
	fmt.Println(successMessage)
}
