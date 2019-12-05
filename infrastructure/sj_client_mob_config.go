package infrastructure

import (
	"ddd-go/presentation/config"
)

func GetConfig() config.Config {
	cfg, _ := config.LoadConfiguration()
	return cfg
}

func SjConfigGetAPibaseUrl() string {
	return GetConfig().SupplierConfig.SjApiBaseUrl
}
func SjConfigGetDeviceId() string {
	return GetConfig().SupplierConfig.SjDeviceId
}
func SjConfigGetSubscribeId() string {
	return GetConfig().SupplierConfig.SjSubscribeId
}
func SjConfigGetUserName() string {
	return GetConfig().SupplierConfig.SjUserName
}
func SjConfigGetPassword() string {
	return GetConfig().SupplierConfig.SjUserPassowd
}
func SjConfigGetOs() string {
	return GetConfig().SupplierConfig.SjOs
}
func SjConfigGetAppsName() string {
	return GetConfig().SupplierConfig.SjAppsName
}
func SjConfigGetAppsVersion() string {
	return GetConfig().SupplierConfig.SjAppsVersion
}
func SjConfigGetApiUrl() string {
	return GetConfig().SupplierConfig.SjApiBaseUrl
}
func SjConfigGetUserLogin() string {
	return ""
}
func SjUseMock() bool {
	return GetConfig().SupplierConfig.SjUseMock
}
func SjSearchPath() string {
	return GetConfig().SupplierConfig.SjApiSearchPath
}
func SjBookPath() string {
	return GetConfig().SupplierConfig.SjApiBookPath
}
func SjBookInfoPath() string {
	return GetConfig().SupplierConfig.SjApiBookInfoPath
}
func SjSetPaymentPath() string {
	return GetConfig().SupplierConfig.SjApiSetPaymentPath
}
