package controllers

import (
	"errors"

	"slashbase.com/backend/internal/daos"
	"slashbase.com/backend/internal/models"
	"slashbase.com/backend/internal/utils"
)

type SettingController struct{}

var settingDao daos.SettingDao

func (sc SettingController) GetSingleSetting(name string) (interface{}, error) {
	setting, err := settingDao.GetSingleSetting(name)
	if err != nil {
		return "", errors.New("there was some problem")
	}
	switch setting.Name {
	case models.SETTING_NAME_APP_ID:
		return setting.UUID().String(), nil
	case models.SETTING_NAME_TELEMETRY_ENABLED:
		return setting.Bool(), nil
	}
	return setting.Value, nil
}

func (sc SettingController) UpdateSingleSetting(name string, value string) error {
	switch name {
	case models.SETTING_NAME_APP_ID:
		return errors.New("cannot update the setting: " + name)
	case models.SETTING_NAME_TELEMETRY_ENABLED:
		if !utils.ContainsString([]string{"true", "false"}, value) {
			return errors.New("cannot update the setting: " + name)
		}
	}
	err := settingDao.UpdateSingleSetting(name, value)
	if err != nil {
		return errors.New("there was some problem updating the setting")
	}
	return nil
}
