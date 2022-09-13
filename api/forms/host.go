package forms

type Host struct {
	HID       uint   `json:"hid" form:"hid"`
	Name      string `json:"name" form:"name"`
	Ipaddress string `json:"ipaddress" form:"ipaddress"`
	Password  string `json:"password" form:"password"`
	Cpu       string `json:"cpu" form:"cpu"`
	Memory    string `json:"memory" form:"memory"`
}
