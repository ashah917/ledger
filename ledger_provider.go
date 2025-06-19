package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Block represents a basic unit in a DAG-based blockchain
type Block struct {
	ID           string    `json:"id"`            // Unique hash or ID
	ParentIDs    []string  `json:"parent_ids"`    // References to parent blocks
	Timestamp    time.Time `json:"timestamp"`     // Block creation time
	Transactions []string  `json:"transactions"`  // Placeholder: simple transaction strings
	Signature    string    `json:"signature"`     // Optional: mock signature
}

func main() {
	router := gin.Default()

	router.POST("/accept", handleAccept)
	router.GET("/check", handleCheck)

	router.Run(":8080")
}

// POST /accept - Accept a new block into the ledger
func handleAccept(c *gin.Context) {
	var blk Block
	if err := c.ShouldBindJSON(&blk); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid block payload"})
		return
	}

	// TODO: validate and store the block
	c.JSON(http.StatusAccepted, gin.H{
		"message": "Block accepted",
		"block":   blk,
	})
}

// GET /check - Placeholder
func handleCheck(c *gin.Context) {
	// Read query params: n, k, t
	n := c.Query("n")
	k := c.Query("k")
	t := c.Query("t")

	// Just return them for now
	c.JSON(http.StatusOK, gin.H{
		"n": n,
		"k": k,
		"t": t,
		"note": "Check logic not implemented yet",
	})
}

