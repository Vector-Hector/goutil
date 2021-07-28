package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func RequiredQueryCappedInt(c *gin.Context, name string, maxValue int) int {
	i := RequiredQueryInt(c, name)
	if i > maxValue {
		PanicBadRequest("the maximum value of " + name + " is " + strconv.Itoa(maxValue))
	}
	return i
}

func RequiredQueryInt(c *gin.Context, name string) int {
	str := RequiredQueryString(c, name)
	i, err := strconv.Atoi(str)
	if err != nil {
		PanicBadRequest(name + " has to be an int.")
	}
	return i
}

func OptionalQueryCappedInt(c *gin.Context, name string, defaultValue int, maxValue int) int {
	i := OptionalQueryInt(c, name, defaultValue)
	if i > maxValue {
		return defaultValue
	}
	return i
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

func RequiredQueryString(c *gin.Context, name string) string {
	val := c.Query(name)
	if val == "" {
		PanicBadRequest(name + " is not optional")
	}
	return val
}

func OptionalQueryString(c *gin.Context, name string, defaultValue string) string {
	val := c.Query(name)
	if val == "" {
		return defaultValue
	}
	return val
}

func RequiredQueryTime(c *gin.Context, name string) time.Time {
	val := c.Query(name)
	if val == "" {
		PanicBadRequest(name + " is not optional")
	}
	t, err := time.Parse(time.RFC3339, val)
	if err != nil {
		PanicBadRequest(name + " has to be of RFC3339 format.")
	}
	return t
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
