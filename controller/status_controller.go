package controller

import (
	"github.com/gin-gonic/gin"
	cron "github.com/robfig/cron/v3"
	"net/http"
	"rincon/config"
	"rincon/service"
	"rincon/utils"
	"strconv"
)

func GetAllServiceStatus(c *gin.Context) {
	returnList := []gin.H{}
	services := service.GetAllServices()
	if len(services) == 0 {
		c.JSON(http.StatusNotFound, returnList)
		return
	}
	for i, s := range services {
		utils.SugarLogger.Infoln(i, s.Name)
		returnList = append(returnList, service.GetServiceStatus(s))
	}
	c.JSON(http.StatusOK, returnList)
}

func GetServiceStatus(c *gin.Context) {
	returnList := []gin.H{}
	services := service.GetServiceByName(c.Param("name"))
	if len(services) == 0 {
		c.JSON(http.StatusNotFound, returnList)
		return
	}
	for i, s := range services {
		utils.SugarLogger.Infoln(i, s.Name)
		returnList = append(returnList, service.GetServiceStatus(s))
	}
	c.JSON(http.StatusOK, returnList)
}

func RegisterStatusCronJob() {
	c := cron.New()
	entryID, err := c.AddFunc("@every "+config.RegistryUpdateDelay+"s", func() {
		go service.Discord.ChannelMessageSend(config.DiscordChannel, ":alarm_clock: Starting Status CRON Job")
		utils.SugarLogger.Infoln("Starting Status CRON Job...")
		services := service.GetAllServices()
		for i, s := range services {
			utils.SugarLogger.Infoln(i, s.Name)
			service.GetServiceStatus(s)
		}
		utils.SugarLogger.Infoln("Finished Status CRON Job!")
		go service.Discord.ChannelMessageSend(config.DiscordChannel, ":white_check_mark: Finished Status CRON Job!")
	})
	if err != nil {
		return
	}
	c.Start()
	utils.SugarLogger.Infoln("Registered CRON Job: " + strconv.Itoa(int(entryID)) + " scheduled for every " + config.RegistryUpdateDelay + "s")
}
