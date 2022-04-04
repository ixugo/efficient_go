package reptile

// ResponseList 列表
type ResponseList struct {
	Code int          `json:"code"`
	Msg  string       `json:"msg"`
	Data responseList `json:"data"`
}
type responseList struct {
	Total int        `json:"total"`
	Count int        `json:"count"`
	List  []ListTile `json:"list"`
}

type ListTile struct {
	Num              int         `json:"num"`
	ResourceID       string      `json:"resource_id"`
	ResourceType     int         `json:"resource_type"`
	ProductID        string      `json:"product_id"`
	ProductType      int         `json:"product_type"`
	Title            string      `json:"title"`
	ImgURL           string      `json:"img_url"`
	ImgURLCompressed string      `json:"img_url_compressed"`
	Summary          string      `json:"summary"`
	PatchImgURL      string      `json:"patch_img_url"`
	VideoLength      string      `json:"video_length"`
	VideoLengthValue int         `json:"video_length_value"`
	ViewCount        int         `json:"view_count"`
	CommentCount     int         `json:"comment_count"`
	StartAt          string      `json:"start_at"`
	CreatedAt        string      `json:"created_at"`
	IsAvailable      int         `json:"is_available"`
	IsTry            int         `json:"is_try"`
	IsView           int         `json:"is_view"`
	IsCurrView       int         `json:"is_curr_view"`
	CurrTime         int         `json:"curr_time"`
	Progress         string      `json:"progress"`
	IsCrop           int         `json:"is_crop"`
	CropInfo         interface{} `json:"crop_info"`
}

// ResponseDetail 详情
type ResponseDetail struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data Detail `json:"data"`
}
type Detail struct {
	ResourceID       string `json:"resource_id"`
	ResourceType     int    `json:"resource_type"`
	ProductTitle     string `json:"product_title"`
	ProductID        string `json:"product_id"`
	ProductType      int    `json:"product_type"`
	Title            string `json:"title"`
	Desc             string `json:"desc"`
	VideoURL         string `json:"video_url"`
	ImgURL           string `json:"img_url"`
	ImgURLCompressed string `json:"img_url_compressed"`
	Summary          string `json:"summary"`
	PatchImgURL      string `json:"patch_img_url"`
	VideoLength      string `json:"video_length"`
	VideoLengthValue int    `json:"video_length_value"`
	ViewCount        int    `json:"view_count"`
	CommentCount     int    `json:"comment_count"`
	StartAt          string `json:"start_at"`
	CreatedAt        string `json:"created_at"`
	IsAvailable      int    `json:"is_available"`
	IsTry            int    `json:"is_try"`
	CurrTime         int    `json:"curr_time"`
	Progress         string `json:"progress"`
}
