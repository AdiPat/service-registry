package service_registry

type Id string
type InstanceIP string

type Instance struct {
	id Id         `json:"id"`
	IP InstanceIP `json:"IP"`
}
