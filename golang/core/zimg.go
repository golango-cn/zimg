package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"io"
	"io/ioutil"
	"zimg/core/common"
)

type ZImg struct {
	File io.Reader
}

func (z *ZImg) Upload() (string, error) {

	b, err := ioutil.ReadAll(z.File)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s/upload", common.ServerSetting.ZImgServer)

	requ := httplib.Post(url)
	requ.Header("content-type", "jpeg")
	requ.Body(b)

	b, err = requ.Bytes()
	if err != nil {
		return "", err
	}

	var ret *Ret
	err = json.Unmarshal(b, &ret)
	if err != nil {
		return "", err
	}

	if ret.Ret {
		return ret.Info.Md5, nil
	}
	return "", errors.New(ret.Error.Message)

}

type Ret struct {
	Ret  bool `json:"ret"`
	Info struct {
		Md5 string `json:"md5"`
	} `json:"info"`
	Error struct {
		Code    float64 `json:"code"`
		Message string  `json:"message"`
	} `json:"error"`
}
