package reptile

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const (
	productID1 = "p_5f0ff4f1e4b04349896c44dd" // 黑客网络应用
	productID2 = "p_5f14fb95e4b0d73896b390bb" // 手把手教你 Linux
	productID3 = "p_5f17c814e4b0a1003cae4503" // 企业渗透测试和持续监控视频教程
	productID4 = "p_5f2b5572e4b073cc175693fc" // 安全渗透测试
	productID5 = "p_5a5066704c531_c4SxiL3g"   // OKR目标管理法
	token      = "f12fb8c277d6d211b576e588e22be5a0"
)

func TestReptile(t *testing.T) {
	r := NewReptile(productID5, token)
	details, err := r.GetFullDetails()
	require.NoError(t, err)
	err = r.SaveVideo(details)
	require.NoError(t, err)
}

func TestGetDetail(t *testing.T) {
	r := NewReptile(productID1, token)
	c, err := r.GetDetail("v_5f50a893e4b06a37e03981cb")
	require.NoError(t, err)
	require.EqualValues(t, c.Code, 0, c.Msg)
	t.Logf("%+v", c.Data)
}

func TestNewRequest(t *testing.T) {
	req, err := http.NewRequest(
		http.MethodPost, "http://baidu.com", strings.NewReader(""),
	)
	require.NoError(t, err)
	// 新创建的 header 默认为空
	require.EqualValues(t, len(req.Header), 0)
}

func TestPrintf(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan struct{}, 5)

	arr := make([]string, 100)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				for i, v := range arr {
					if v != "" {
						fmt.Printf("\r%s", v)
					}
					if strings.Contains(v, "OK") {
						arr[i] = ""
					}
				}
			}
		}
	}()

	for i := 0; i < 100; i++ {
		wg.Add(1)
		ch <- struct{}{}
		arr[i] = fmt.Sprintf("%d\tHello", i)

		go func(i int) {
			defer func() {
				<-ch
				wg.Done()
			}()
			time.Sleep(time.Duration(rand.Intn(3)+2) * time.Second)
			arr[i] += "\t\tOK\n"
		}(i)
	}
	wg.Wait()
}
