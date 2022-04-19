package reptile

import (
	"encoding/json"

	"os"
	"testing"

	"github.com/ixugo/efficient_go/reptile/bar"
	"github.com/stretchr/testify/require"
)

const (
	productID1 = "p_5f0ff4f1e4b04349896c44dd" // 黑客网络应用
	productID2 = "p_5f14fb95e4b0d73896b390bb" // 手把手教你 Linux
	productID3 = "p_5f17c814e4b0a1003cae4503" // 企业渗透测试和持续监控视频教程
	productID4 = "p_5f2b5572e4b073cc175693fc" // 安全渗透测试
	productID5 = "p_5a5066704c531_c4SxiL3g"   // OKR目标管理法
	token      = "61fc41edf2cdb44c72c636f1c34f9929"
)

// TestReptile 测试保存视频到本地
func TestReptile(t *testing.T) {
	r := NewReptile(productID5, token, bar.NewMpb())
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
	if count > 10 {
		count = 10
	}
	r := NewReptile(productID5, token, bar.NewMpb())
	err = r.SaveVideo(data[0:count])
	require.NoError(t, err)
}

// TestGetDetail 测试获取一集视频的详情
func TestGetDetail(t *testing.T) {
	r := NewReptile(productID1, token, bar.NewMpb())
	c, err := r.GetDetail("v_5f50a893e4b06a37e03981cb")
	require.NoError(t, err)
	require.EqualValues(t, c.Code, 0, c.Msg)
	t.Logf("%+v", c.Data)
}
