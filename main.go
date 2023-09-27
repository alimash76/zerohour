package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/getarmy", ArmyAssigner)
	r.Run("0.0.0.0:8080")
}

type Players struct {
	PlayerNames []string `json:"playernames"`
}

func ArmyAssigner(c *gin.Context) {
	var input Players

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	armies := []string{
		"USA",
		"USA Lasser",
		"USA AirForce",
		"USA Superweapon",
		"China",
		"China Nuke",
		"China Tank",
		"China Infantry",
		"GLA",
		"GLA Toxic",
		"GLA Stealth",
		"GLA Demolition",
	}

	match := make(map[string]string)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len(input.PlayerNames); i++ {
		player_name := input.PlayerNames[i]
		player_army := armies[rand.Intn(len(armies))]
		match[player_name] = player_army
	}

	c.JSON(http.StatusOK, gin.H{"message": match})

}
