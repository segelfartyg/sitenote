package main

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

const DEV_COMPOSE_FILE string = "../dev/dev.docker-compose.yaml"
const PROD_COMPOSE_FILE string = "../prod/dev.docker-compose.yaml"

type githubWebhookRequest struct {
}

type workflow struct {
	badge_url  string `json:"badge_url"`
	created_at string `json:"created_at"`
	html_url   string `json:"html_url"`
	id         string `json:"id"`
	name       string `json:"name"`
	node_id    string `json:"node_id"`
	path       string `json:"path"`
	state      string `json:"state"`
	updated_at string `json:updated_at`
	url        string `json:url`
}

func main() {
	r := gin.Default()
	r.GET("/webhook", deploy)
	r.Run("localhost:9050")
}

func deploy(c *gin.Context) {

	githubReq := githubWebhookRequest{}

	c.ShouldBind(&githubReq)

	fmt.Println(githubReq)

	// env := DEV_COMPOSE_FILE

	// if env == PROD_COMPOSE_FILE {
	// 	deployProdEnvironment(PROD_COMPOSE_FILE)
	// } else {
	// 	deployDevEnvironment(DEV_COMPOSE_FILE)
	// }

	res := "deployment succeeded"

	c.String(http.StatusOK, res)
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
