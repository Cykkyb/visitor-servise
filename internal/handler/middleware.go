package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"runtime"
	"time"
)

func debugStartMiddleware(c *gin.Context) {
	fmt.Println("debug start")
	debugTime := time.Now()
	c.Set("debugTime", debugTime)

	memStats := runtime.MemStats{}
	runtime.ReadMemStats(&memStats)
	c.Set("memStatsBefore", memStats)
	c.Next()
}

func debugEndMiddleware(c *gin.Context) {
	debugTime, ok := c.Get("debugTime")
	if !ok {
		return
	}
	elapsed := time.Since(debugTime.(time.Time))
	c.Writer.Header().Set("X-Debug-Time", fmt.Sprintf("%d", int(elapsed/time.Millisecond)))

	memStatsBefore, ok := c.Get("memStatsBefore")
	if !ok {
		return
	}
	memStatsBeforeValue := memStatsBefore.(runtime.MemStats)

	memStats := runtime.MemStats{}
	runtime.ReadMemStats(&memStats)
	debugMemUsed := memStats.TotalAlloc - memStatsBeforeValue.TotalAlloc
	debugMemUsedKB := debugMemUsed / 1024
	c.Writer.Header().Set("X-Debug-Memory", fmt.Sprintf("%d", debugMemUsedKB))

	c.Next()
}
