package forms

type Host struct {
	Name      string `json:"name" form:"name"`
	Username  string `json:"username" form:"username"`
	Ipaddress string `json:"ipaddress" form:"ipaddress"`
	Password  string `json:"password" form:"password"`
	Port      uint   `json:"port" form:"port"`
	Memory    string `json:"memory" form:"memory"`
}

type WSHost struct {
	Username  string `json:"username" form:"username"`
	Ipaddress string `json:"ipaddress" form:"ipaddress"`
	Password  string `json:"password" form:"password"`
	Port      uint   `json:"port" form:"port"`
}
