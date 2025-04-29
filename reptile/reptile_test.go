package reptile

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/ixugo/efficient_go/reptile/bar"
	"github.com/stretchr/testify/require"
)

const (
	productID1            = "p_5f0ff4f1e4b04349896c44dd" // 黑客网络应用
	productID2            = "p_5f14fb95e4b0d73896b390bb" // 手把手教你 Linux
	productID3            = "p_5f17c814e4b0a1003cae4503" // 企业渗透测试和持续监控视频教程
	productID4            = "p_5f2b5572e4b073cc175693fc" // 安全渗透测试
	productID5            = "p_5a5066704c531_c4SxiL3g"   // OKR目标管理法
	ProductPython1        = "p_5f6b0c32e4b01b26d1bbdf50" // python 1
	ProductPython2        = "p_6119daf0e4b0a27d0e3e1030" // python 2
	ProductPython3        = "p_611a119ce4b0cce271be963c" // python 3
	ProductPython4        = "p_611b6645e4b0cce271bf26ac" // python 4
	ProductPython5        = "p_611c5f54e4b0bf6430075bdd" // python 5
	ProductPythonRefactor = "p_5f474c17e4b0dd4d974b924e" // python 重构 无效
	Product6              = "p_5f6b0997e4b01b26d1bbddfc" // 24 篇算法精讲
	Product7              = "p_5f6b09bde4b0d59c87b7c88b" // 9 篇算法精讲

	Product8 = "p_5f335d21e4b075dc42ad36b7" // 网络课程
	token    = "b95b4e1afd84d340bab337373b74a6b6"
)

const currentProduct = Product8

// TestReptile 测试保存视频到本地
func TestReptile(t *testing.T) {
	r := NewReptile(currentProduct, token, bar.NewMpb())
	details, err := r.GetFullDetails()
	require.NoError(t, err)
	err = r.SaveVideo(details)
	require.NoError(t, err)
}

// 测试读取配置文件保存视频
func TestReadFile(t *testing.T) {
	b, err := os.ReadFile("details.json")
	require.NoError(t, err)
	data := make([]Detail, 0, 80)
	err = json.Unmarshal(b, &data)
	require.NoError(t, err)

	count := len(data)
	// if count > 10 {
	// count = 10
	// }
	r := NewReptile(currentProduct, token, bar.NewMpb())
	err = r.SaveVideo(data[0:count])
	require.NoError(t, err)
}

// TestGetDetail 测试获取一集视频的详情
func TestGetDetail(t *testing.T) {
	r := NewReptile(currentProduct, token, bar.NewMpb())
	c, err := r.GetDetail("v_5f34add7e4b075dc42ad7d1b")
	require.NoError(t, err)
	require.EqualValues(t, c.Code, 0, c.Msg)
	t.Logf("%+v", c.Data)
}
