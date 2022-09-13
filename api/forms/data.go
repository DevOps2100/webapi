package forms

type Data struct {
	ID           uint   `json:"id" form:"id"`
	Name         string `json:"name" form:"name"`
	DatabaseType string `json:"database_type" form:"database_type"`
	ConnectSite  string `json:"connect_site" form:"connect_site"`
	Username     string `json:"username" form:"username"`
	Passwd       string `json:"passwd" form:"passwd"`
	Port         string `json:"port" form:"port"`
}
