/*
* @desc:上传文件model
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/12/6 15:29
 */

package model

type UpFile struct {
	Name     string `json:"name"`
	Url      string `json:"url"`
	FileType string `json:"fileType"`
	Size     uint64 `json:"size"`
}
