package entity

// Project represents the project entity.
type Project struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	//Env        map[string]interface{} `json:"env"`
	//ExpectRes  bool                   `json:"expect_res"`
	//UseDocker  bool                   `json:"use_docker"`
	//LinkDocker string                 `json:"link_docker"`
}
