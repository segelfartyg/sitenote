package main

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/webhook", deploy)
	r.Run("localhost:9050")
}

func deploy(c *gin.Context) {

	composeDown()
	removeImages()
	getLatestImages()
	composeUp()

	res := "deployment succeeded"

	c.String(http.StatusOK, res)
}

func composeDown() {

	cmd := exec.Command("docker", "compose", "-f", "../dev/dev.docker-compose.yaml", "down")
	commandOutput, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	successMessage := string(commandOutput)
	fmt.Println(successMessage)

}

func removeImages() {
	cmd := exec.Command("docker", "compose", "-f", "../dev/dev.docker-compose.yaml", "rm", "-f")
	commandOutput, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	successMessage := string(commandOutput)
	fmt.Println(successMessage)

}

func getLatestImages() {
	cmd := exec.Command("docker", "compose", "-f", "../dev/dev.docker-compose.yaml", "pull", "-f")
	commandOutput, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	successMessage := string(commandOutput)
	fmt.Println(successMessage)
}

func composeUp() {
	cmd := exec.Command("docker", "compose", "-f", "../dev/dev.docker-compose.yaml", "up", "-d")
	commandOutput, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
		return
	}
	successMessage := string(commandOutput)
	fmt.Println(successMessage)
}
