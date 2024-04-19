package gsr

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
)

func Resty() {
	resp, err := resty.New().R().EnableTrace().Get("https://api.github.com/repos/nxcc/go-script-runner/commits")
	fmt.Println("Error:", err)
	fmt.Println("Trace Info:", resp.Request.TraceInfo())
	fmt.Printf(">>> %v\n", gjson.GetBytes(resp.Body(), "0.commit.message"))
}
