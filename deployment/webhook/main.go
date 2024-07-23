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

	cmd := exec.Command("docker", "compose", "-f", "../dev/dev.docker-compose.yaml", "up", "-d")
	commandOutput, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
		return
	}
	successMessage := string(commandOutput)

	c.String(http.StatusOK, successMessage)
}
