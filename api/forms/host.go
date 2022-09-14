package forms

type Host struct {
	Name      string `json:"name" form:"name"`
	Ipaddress string `json:"ipaddress" form:"ipaddress"`
	Password  string `json:"password" form:"password"`
	Memory    string `json:"memory" form:"memory"`
}
