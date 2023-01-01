package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sendmq/Lib"

	"github.com/gin-gonic/gin"
)

type Stage struct {
	StageName   string `json:"stage_name"`
	StageStatus string `json:"stage_status"`
	Time        string `json:"time"`
}

type MyStruct struct {
	BuildID       string `json:"build_id"`
	PipelineType  string `json:"pineline_type"`
	Group         string `json:"group"`
	TargetBranch  string `json:"target_branch"`
	Operator      string `json:"operator"`
	GitCommit     string `json:"git_commit"`
	CommiterName  string `json:"commiter_name"`
	CommiterEmail string `json:"commiter_email"`
	AuthorName    string `json:"author_name"`
	AuthorEmail   string `json:"author_string"`
	Stage         Stage  `json:"stage"`
}

func main() {
	r := gin.Default()

	r.POST("/jenkins", handlePOST)

	c := make(chan error)

	go func() {
		err := r.Run(":8081")
		if err != nil {
			c <- err
		}
	}()
	go func() {
		err := Lib.UserInit()
		if err != nil {
			c <- err
		}
	}()

	err := <-c
	log.Fatal(err)

}

func handlePOST(c *gin.Context) {
	var jsonData MyStruct
	if err := c.ShouldBindJSON(&jsonData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Convert the JSON payload to a byte slice
	jsonBytes, err := json.Marshal(jsonData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	mq := Lib.NewMq()

	err = mq.SendMessage(Lib.ROUTER_KEY_USERREG, Lib.EXCHANGE_USER, jsonBytes)

	defer mq.Channel.Close()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Sent message to RabbitMQ: %s", jsonData)
}
