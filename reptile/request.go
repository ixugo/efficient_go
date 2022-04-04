package reptile

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
)

const (
	reousrceDetailGet = "https://www.hzmedia.com.cn/api/resource_detail_get.ashx" // 资源详情
	resourceListGet   = "https://www.hzmedia.com.cn/api/resource_list_get.ashx"   // 资源列表
)

const (
	listName    = "list.json"
	detailsName = "details.json"
)

// Reptile 某章书院课程下载器
type Reptile struct {
	productID string
	token     string
}

// NewReptile 创建该课程下载器
func NewReptile(PID, token string) *Reptile {
	_ = os.Mkdir(PID, os.ModePerm)
	return &Reptile{
		productID: PID,
		token:     token,
	}
}

// ListName 列表名
func (r *Reptile) ListName() string {
	return r.productID + "/" + listName
}

// DetailsName 每集视频详情
func (r *Reptile) DetailsName() string {
	return r.productID + "/" + detailsName
}

// VideoName 每集视频名称
func (r *Reptile) VideoName(name string) string {
	return r.productID + "/" + name
}

func newRequest(url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}
	// TODO 浏览器是禁用设置 Referer，此处待确定是否有效
	req.Header.Set("Referer", "http://yd.hzmedia.com.cn/")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) MicroMessenger/6.8.0(0x16080000) MacWechat/3.2.2(0x13020210) NetType/WIFI WindowsWechat")
	return req, nil
}

func do(req *http.Request) ([]byte, error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf(resp.Status)
	}
	b, err := io.ReadAll(resp.Body)
	_ = resp.Body.Close()
	return b, err
}

func writeFile[T any](t T, filename string) error {
	b, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, b, os.ModePerm)
}

// GetFullList 获取该课程的全部列表
func (r *Reptile) GetFullList() ([]ListTile, error) {
	resp, err := r.GetList("1", "1")
	if err != nil {
		panic(err)
	}
	if resp.Code != 0 {
		panic(resp.Msg)
	}

	total := resp.Data.Total
	resp, err = r.GetList("1", strconv.Itoa(total))
	if err != nil {
		panic(err)
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf(resp.Msg)
	}
	err = writeFile(resp.Data.List, r.ListName())
	if err != nil {
		fmt.Println(err)
	}
	return resp.Data.List, err
}

// GetList 获取指定列表
func (r *Reptile) GetList(index, size string) (list ResponseList, err error) {
	var (
		req *http.Request
		b   []byte
	)
	data := url.Values{
		"product_id":    []string{r.productID},
		"page_index":    []string{index},
		"page_size":     []string{size},
		"order_by":      []string{"created_at:asc"},
		"resource_type": []string{"0"},
		"is_try":        []string{"-1"},
		"token":         []string{r.token},
	}

	req, err = newRequest(resourceListGet, strings.NewReader(data.Encode()))
	if err != nil {
		return
	}
	b, err = do(req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &list)
	return
}

// GetDetail 获取视频详情
func (r *Reptile) GetDetail(rid string) (c ResponseDetail, err error) {
	var (
		req *http.Request
		b   []byte
	)

	data := url.Values{
		"product_id":  []string{r.productID},
		"resource_id": []string{rid},
		"token":       []string{r.token},
	}
	req, err = newRequest(reousrceDetailGet, strings.NewReader(data.Encode()))
	if err != nil {
		return
	}
	b, err = do(req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &c)
	return
}

// GetFullDetails 获取每部视频的详情，保存到文件中
func (r *Reptile) GetFullDetails() ([]Detail, error) {
	list, err := r.GetFullList()
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		fmt.Println("无列表数据")
		return nil, nil
	}

	details := make([]Detail, 0, len(list))
	for _, v := range list {
		// fmt.Printf("第%d课: %s\n", i, v.Title)
		c, err := r.GetDetail(v.ResourceID)
		if err != nil {
			return nil, err
		}
		if c.Code != 0 {
			return nil, fmt.Errorf(c.Msg)
		}
		details = append(details, c.Data)
	}

	err = writeFile(details, r.DetailsName())
	if err != nil {
		fmt.Println(err)
	}
	return details, nil
}

// SaveVideo 通过详情下载视频，保存到本地
func (r *Reptile) SaveVideo(details []Detail) error {
	ch := make(chan struct{}, 5)
	var wg sync.WaitGroup
	for i, v := range details {
		ch <- struct{}{}
		wg.Add(1)
		go func(i int, v Detail) {
			defer func() {
				<-ch
				wg.Done()
			}()
			fmt.Printf("第%02d课: %s\n", i+1, v.Title)
			resp, err := http.Get(v.VideoURL)
			if err != nil {
				panic("发生错误: " + err.Error())
			}
			b, _ := io.ReadAll(resp.Body)
			_ = resp.Body.Close()
			idx := strings.LastIndex(v.VideoURL, ".")

			_ = os.WriteFile(r.VideoName(v.Title+v.VideoURL[idx:]), b, os.ModePerm)
		}(i, v)
	}
	wg.Wait()
	return nil
}
