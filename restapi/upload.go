package restapi

import (
	"github.com/sohaha/zlsgo/znet"
	"github.com/zlsgo/app_core/common"
)

func HanderUpload(c *znet.Context, subDirName string, opt ...func(o *common.UploadOption)) ([]common.UploadResult, error) {
	resp, err := common.Upload(c, subDirName, opt...)

	return resp, err
}
