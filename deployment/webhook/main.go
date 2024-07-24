package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

const DEV_COMPOSE_FILE string = "../dev/dev.docker-compose.yaml"
const DEV_SERVER_WORKFLOW_NAME string = "sitenote-server-ci"
const DEV_FRONTEND_WORKFLOW_NAME string = "sitenote-frontend-ci"

const PROD_COMPOSE_FILE string = "../prod/dev.docker-compose.yaml"
const PROD_SERVER_WORKFLOW_NAME string = "sitenote-server-ci-prod"
const PROD_FRONTEND_WORKFLOW_NAME string = "sitenote-frontend-ci-prod"

var secret string = "hej"

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

	fmt.Println("SECRET:", os.Getenv("SECRET"))
	secret = os.Getenv("SECRET")
	r := gin.Default()
	r.POST("/webhook", deploy)
	r.Run("localhost:9050")
}

func deploy(c *gin.Context) {

	header := c.Request.Header
	signature := header.Get("X-Hub-Signature-256")

	if signature == "" {
		http.Error(c.Writer, "missing signature", http.StatusUnauthorized)
		return
	}

	fmt.Println(header)

	b, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.String(http.StatusBadRequest, "no body")
	}

	expectedMAC := ComputeHMAC(b, []byte(secret))
	receivedMAC := signature[7:]

	fmt.Println(expectedMAC)
	fmt.Println(receivedMAC)

	if !hmac.Equal([]byte(expectedMAC), []byte(receivedMAC)) {
		c.String(http.StatusUnauthorized, "not valid sig")
		return
	}

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

func ComputeHMAC(message, secret []byte) string {
	mac := hmac.New(sha256.New, secret)
	mac.Write(message)
	return hex.EncodeToString(mac.Sum(nil))
}
