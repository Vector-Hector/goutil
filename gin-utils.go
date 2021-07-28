package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func RequiredQueryCappedInt(c *gin.Context, name string, maxValue int) (int, bool) {
	i, done := RequiredQueryInt(c, name)
	if done {
		return 0, true
	}
	if i > maxValue {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": "the maximum value of " + name + " is " + strconv.Itoa(maxValue),
		})
		return 0, true
	}
	return i, false
}

func RequiredQueryInt(c *gin.Context, name string) (int, bool) {
	str, done := RequiredQueryString(c, name)
	if done {
		return 0, true
	}
	i, err := strconv.Atoi(str)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": name + " has to be an int.",
		})
		return 0, true
	}
	return i, false
}

func RequiredQueryString(c *gin.Context, name string) (string, bool) {
	val := c.Query(name)
	if val == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": name + " is not optional",
		})
		return "", true
	}
	return val, false
}

func OptionalQueryTime(c *gin.Context, name string, defaultValue time.Time) time.Time {
	val := c.Query(name)
	if val == "" {
		return defaultValue
	}
	t, err := time.Parse(time.RFC3339, val)
	if err != nil {
		fmt.Println(err)
		return defaultValue
	}
	return t
}

func OptionalQueryInt(c *gin.Context, name string, defaultValue int) int {
	val := c.Query(name)
	if val == "" {
		return defaultValue
	}
	i, err := strconv.Atoi(val)
	if err != nil {
		return defaultValue
	}
	return i
}

func OptionalQueryString(c *gin.Context, name string, defaultValue string) string {
	val := c.Query(name)
	if val == "" {
		return defaultValue
	}
	return val
}

