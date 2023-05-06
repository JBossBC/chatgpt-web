package dao

const paramsTable = "params"

type Params struct {
	Key               string `gorm:"key" json:"-"`
	Model             string `gorm:"model" json:"model"`
	Max_tokens        int    `gorm:"max_tokens" json:"max_tokens"`
	Temperature       int    `gorm:"temperature" json:"temperature"`
	Top_p             int    `gorm:"top_p" json:"top_p"`
	Presence_penalty  int    `gorm:"presence_penalty" json:"presence_penalty"`
	Frequency_penalty int    `gorm:"frequency_penalty" json:"frequency_penalty"`
	N                 int    `gorm:"n" json:"n"`
	Stream            bool   `gorm:"stream" json:"stream"`
	Stop              string `gorm:"stop" json:"stop"`
	// Logit_bias        map[string]interface{} `gorm:"type:json" json:"logit_bias"`
	Logit_bias    map[string]interface{} `gorm:"-" json:"-"`
	Role          string                 `gorm:"role" json:"-"`
	Output_format string                 `gorm:"output_format" json:"-"`
	Keep_context  bool                   `gorm:"keep_context" json:"-"`
	IsChange      bool                   `gorm:"-" json:"-"`
}

// init for database
// current can use default configure
var defaultParams Params

func GetDefaultParams() Params {
	return defaultParams
}
func init() {
	newMysqlConn().AutoMigrate(&Params{})
	newMysqlConn().First(&defaultParams)
}

// type paramsDao struct {
// }

// var (
// 	params     *paramsDao
// 	paramsOnce sync.Once
// )

// func NewParamsDao() *paramsDao {
// 	paramsOnce.Do(func() {
// 		params = &paramsDao{}
// 	})
// 	return params
// }
