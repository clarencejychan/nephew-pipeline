package scheduler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"time"
)

func CreateSchedulerTask() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		destination := c.PostForm("destination")
		scheduler_time := createTime(c.PostForm("scheduler_time"))
		occurence := createOccurence(c.PostForm("occurence_num"), c.PostForm("occurence_unit"), scheduler_time)
		parameters := createParameters(c.PostForm("keys"), c.PostForm("values"))

		// Call scheduler task here
		c.JSON(http.StatusOK, gin.H{
			"destination":    destination,
			"scheduler_time": scheduler_time,
			"parameters":     parameters,
			"occurence_num":  occurence,
		})
	}
	return gin.HandlerFunc(fn)
}

// Determine proper time zone, right now set to EST
func createTime(time_string string) time.Time {
	time_string += "Z"
	layout := "2006-01-02T15:04Z07:00"
	scheduler_time, err := time.Parse(layout, time_string)

	if err != nil {
		log.Println(err.Error())
	}
	return scheduler_time
}

// Right now it is returning string as a placeholder
func createOccurence(occurence_string string, unit string, scheduler_time time.Time) string {
	// occurence_num, err := strconv.Atoi(occurence_string)

	/* if err != nil {
		log.Println(err.Error())
	} */

	// Place holder for now
	occurence := "Placeholder"

	// Parse date accordinly from model
	switch unit {
	case "Days(s)":
	case "Weeks(s)":
	case "Months(s)":
	case "Years(s)":
	default:
		log.Println("Invalid unit")
	}
	return occurence
}

func createParameters(keys string, values string) map[string]string {
	key_arr := strings.Split(keys, ",")
	value_arr := strings.Split(values, ",")

	if len(key_arr) != len(value_arr) {
		log.Println("Invalid key value format")
	}

	var params map[string]string
	params = make(map[string]string)

	for i := 0; i < len(key_arr); i++ {
		params[key_arr[i]] = value_arr[i]
	}

	return params
}
