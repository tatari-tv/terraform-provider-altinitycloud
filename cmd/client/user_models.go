package client

type UserData struct {
	Data []User `json:"data"`
}

type User struct {
	ID               string `json:"id"`
	Login            string `json:"login"`
	Password         string `json:"password"`
	Networks         string `json:"networks"`
	Databases        string `json:"databases"`
	IDCluster        string `json:"id_cluster"`
	IDProfile        string `json:"id_profile"`
	IDQuota          string `json:"id_quota"`
	AccessManagement bool   `json:"accessManagement"`
	System           bool   `json:"system"`
}
