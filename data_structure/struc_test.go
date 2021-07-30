package data_structure

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"
	"testing"
	"time"
)

type MenuExtra struct {
	*MenuExtraThirdApp
	*MenuExtraWx
	*MenuExtraProcessTitleDynamic
}

type MenuExtraThirdApp struct {
	PackageName string `json:"package_name" form:"package_name"` // 安卓包名
	IosScheme   string `json:"ios_scheme" form:"ios_scheme"`     // ios协议
	IosUrl      string `json:"ios_url" form:"ios_url"`           // ios协议
}

type MenuExtraWx struct { //小程序
	Path       string `json:"path" form:"path"`
	OriginalId string `json:"original_id" form:"original_id"`
}

type MenuExtraProcessTitleDynamic struct { // 特殊处理：动态标题倒计时
	EndTime     time.Time `json:"end_time" form:"end_time"`
	Placeholder string    `json:"placeholder" form:"placeholder"`
	EndTitle    string    `json:"end_title" form:"end_title""`
}
type Menu struct {
	Id int `json:"id"`
	MenuExtra
}
func TestMarshalStruct(t *testing.T)  {
	// s :="{\"end_time\": \"2021-04-20 15:35:00\", \"end_title\": \"真真小仙女\", \"placeholder\": \"#countdown#\"}"
	var extra MenuExtra
	// extra = MenuExtra{}
	// err := json.Unmarshal([]byte(s), &extra)
	by, err := json.Marshal(extra)
	if err != nil {
		fmt.Println("json unmarshal error", err)
		return
	}
	if string(by) == "{}" {
		by = nil
	}
	fmt.Printf("extra: %s\n", string(by))
}
func TestUnmarshalStruct(t *testing.T)  {
	a := 123
	fmt.Printf("%02d\n", a)
}

func TestName(t *testing.T) {
	srcTitle := "倒计时#countdown#天"
	type Extra struct {
		EndTime     string `json:"end_time"`
		Placeholder string `json:"placeholder"`
		EndTitle    string `json:"end_title"`
	}
	extra := `{"end_time":"2021-05-01 15:08:09","placeholder":"老郑","end_title":"真真小仙女"}`
	var ep Extra
	if err := json.Unmarshal([]byte(extra), &ep); err != nil {
		fmt.Println("json unmarshal error", err)
	}
	t1, err := time.ParseInLocation("2006-01-02 15:04:05", ep.EndTime, time.Local)
	if err != nil {
		fmt.Println("time parse error:", err)
		return
	}
	remainDays := int(math.Ceil(t1.Sub(time.Now()).Hours() / 24))
	if remainDays <= 0 {
		title := ep.EndTitle
		fmt.Println("title:", title)
	} else {
		title := strings.Replace(srcTitle, ep.Placeholder, fmt.Sprintf("%02d", remainDays), -1)
		fmt.Println("title:", title)
	}
}
