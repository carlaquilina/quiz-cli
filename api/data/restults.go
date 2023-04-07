package data

import (
	"github.com/spf13/myapp/api/models"
)

// Define a map to store the quiz results
var (
	Results = models.Results{
		M: map[string]int{"user2": 3, "user3": 1, "user4": 0, "user5": 0},
	} //map[username]score
)
