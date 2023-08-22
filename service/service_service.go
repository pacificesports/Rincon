package service

import (
	"crypto/rand"
	"rincon/config"
	"rincon/model"
	"rincon/utils"
	"strconv"
)

func GetAllServices() []model.Service {
	var services []model.Service
	result := DB.Find(&services)
	if result.Error != nil {
	}
	return services
}

func GetServiceByID(id int) model.Service {
	var service model.Service
	result := DB.Where("id = ?", id).Find(&service)
	if result.Error != nil {
	}
	return service
}

func GetServiceByName(name string) []model.Service {
	var services []model.Service
	result := DB.Where("UPPER(name) = UPPER(?)", name).Find(&services)
	if result.Error != nil {
	}
	return services
}

func CreateService(service model.Service) (model.Service, error) {
	var existingServices []model.Service
	_ = DB.Where("url = ?", service.URL).Find(&existingServices)
	for i, s := range existingServices {
		utils.SugarLogger.Infoln("Removing existing service (" + strconv.Itoa(i) + ") with duplicate url: " + s.URL)
		_ = RemoveService(existingServices[i])
	}
	id, err := GenerateServiceID(6)
	if err != nil {
		return service, err
	}
	service.ID, _ = strconv.Atoi(id)
	if result := DB.Create(&service); result.Error != nil {
		return service, result.Error
	}
	go Discord.ChannelMessageSend(config.DiscordChannel, "New service ("+strconv.Itoa(service.ID)+") "+service.Name+" added to registry")
	return service, nil
}

func RemoveService(service model.Service) error {
	if result := DB.Delete(&service); result.Error != nil {
		return result.Error
	}
	go Discord.ChannelMessageSend(config.DiscordChannel, "Service ("+strconv.Itoa(service.ID)+") "+service.Name+" removed from registry")
	return nil
}

const idChars = "1234567890"

func GenerateServiceID(length int) (string, error) {
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	otpCharsLength := len(idChars)
	for i := 0; i < length; i++ {
		buffer[i] = idChars[int(buffer[i])%otpCharsLength]
	}

	return string(buffer), nil
}
