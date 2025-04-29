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
	reousrceDetailGet = "https://cloud.cmpreading.com/api/resource_detail_get.ashx" // 资源详情
	resourceListGet   = "https://cloud.cmpreading.com/api/resource_list_get.ashx"   // 资源列表
)

const (
	listName    = "list.json"
	detailsName = "details.json"
)

// Copier  需要实现 io.copy
// 接口应由调用者提供，而非实现者。
type Copier interface {
	Copy(name string, total int64, dst io.Writer, src io.Reader) (written int64, err error)
}

// Reptile 某章书院课程下载器
type Reptile struct {
	ProductID string
	Token     string
	Copier    // 拷贝
}

// NewReptile 创建该课程下载器
func NewReptile(PID, token string, c Copier) *Reptile {
	_ = os.Mkdir(PID, os.ModePerm)
	return &Reptile{
		ProductID: PID,
		Token:     token,
		Copier:    c,
	}
}

// ListName 列表名
func (r *Reptile) ListName() string {
	return r.ProductID + "/" + listName
}

// DetailsName 每集视频详情
func (r *Reptile) DetailsName() string {
	return r.ProductID + "/" + detailsName
}

// VideoName 每集视频名称
func (r *Reptile) VideoName(name string) string {
	return strings.ReplaceAll(strings.TrimSpace(r.ProductID+"/"+name), "；", "_")
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
		"product_id":    []string{r.ProductID},
		"page_index":    []string{index},
		"page_size":     []string{size},
		"order_by":      []string{"created_at:asc"},
		"resource_type": []string{"0"},
		"is_try":        []string{"-1"},
		"token":         []string{r.Token},
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
		"product_id":  []string{r.ProductID},
		"resource_id": []string{rid},
		"token":       []string{r.Token},
	}
	req, err = newRequest(reousrceDetailGet, strings.NewReader(data.Encode()))
	if err != nil {
		return
	}
	b, err = do(req)
	if err != nil {
		return
	}
	fmt.Println(">>>>>>>>>..", string(b))
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
			// fmt.Printf("第%02d课: %s\n", i+1, v.Title)
			resp, err := http.Get(v.VideoURL)
			defer func() {
				_ = resp.Body.Close()
			}()
			if err != nil {
				panic("发生错误: " + err.Error())
			}
			total, _ := strconv.Atoi(resp.Header.Get("Content-Length"))

			idx := strings.LastIndex(v.VideoURL, ".")
			name := r.VideoName(v.Title + v.VideoURL[idx:])

			tmpName := name + ".tmp"
			file, err := os.Create(tmpName)
			if err != nil {
				panic(err)
			}
			_, err = r.Copy(v.Title, int64(total), file, resp.Body)
			if err != nil {
				panic(err)
			}
			if err = os.Rename(tmpName, name); err != nil {
				panic(err)
			}
		}(i, v)
	}
	wg.Wait()
	return nil
}
