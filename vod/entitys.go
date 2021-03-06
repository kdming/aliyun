package vod

// AliyunVod 公共参数
type AliyunVod struct {
	Format           string //返回值的类型，支持JSON与XML
	Version          string //API版本号，为日期形式：YYYY-MM-DD，本版本对应为2017-03-21
	AccessKeyID      string `url:"AccessKeyId"` //阿里云颁发给用户的访问服务所用的密钥ID
	SignatureMethod  string //签名方式，目前支持HMAC-SHA1
	Timestamp        string //请求的时间戳。日期格式按照ISO8601标准表示，并需要使用UTC时间。格式为：YYYY-MM-DDThh:mm:ssZ例如，2017-3-29T12:00:00Z(为北京时间2017年3月29日的20点0分0秒
	SignatureVersion string //签名算法版本，目前版本是1.0
	SignatureNonce   string //唯一随机数，用于防止网络重放攻击。用户在不同请求间要使用不同的随机数值
	AccessSecret     string `url:"-"`
}

type baseResposeEntity struct {
	RequestID string `json:"RequestId"`
}

// CreateUploadVideoResposeEntity CreateUploadVideo接口返回信息
type CreateUploadVideoResposeEntity struct {
	baseResposeEntity
	VideoID       string `json:"VideoId"`
	UploadAddress string //上传地址
	UploadAuth    string //上传凭证
}

// PlayAuthResposeEntity PlayAuth返回
type PlayAuthResposeEntity struct {
	baseResposeEntity
	VideoMeta videoDetail
	PlayAuth  string
}

type videoDetail struct {
	VideoID  string `json:"VideoId"`
	Title    string
	Duration float32
	CoverURL string
	Status   string
}

// RefreshUploadVideoResposeEntity 刷新视频上传凭证返回
type RefreshUploadVideoResposeEntity struct {
	baseResposeEntity
	UploadAuth string
}

// ImageType 图片类型
type ImageType string

// ImageExt 图片文件扩展名
type ImageExt string

const (
	// Cover ImageType 封面
	Cover ImageType = "cover"
	// Watermark ImageType 水印
	Watermark = "watermark"
	// Png ImageExt
	Png ImageExt = "png"
	// Jpg ImageExt
	Jpg = "jpg"
	// Jpeg ImageExt
	Jpeg = "jpeg"
)

// CreateUploadImageResposeEntity 获取图片上传地址和凭证返回
type CreateUploadImageResposeEntity struct {
	baseResposeEntity
	UploadAddress string
	UploadAuth    string
	ImageURL      string
}

// GetPlayInfoResposeEntity 获取视频播放地址返回
type GetPlayInfoResposeEntity struct {
	baseResposeEntity
	VideoBase    videoBase
	PlayInfoList playInfoList
}
type videoBase struct {
	VideoID      string `json:"VideoId"`
	Title        string
	Duration     string
	CoverURL     string
	Status       string
	CreationTime string
	MediaType    string
}
type playInfoList struct {
	PlayInfo []playInfo
}
type playInfo struct {
	Bitrate    string
	Definition string
	Duration   string
	Encrypt    int
	PlayURL    string
	Format     string
	Fps        string
	Size       int //视频流大小，单位Byte
	Width      int
	Height     int
}

// GetVideoInfoResposeEntity 获取视频信息返回
type GetVideoInfoResposeEntity struct {
	baseResposeEntity
	Video video
}
type video struct {
	videoDetail
	Description  string
	CreationTime string
	Size         int
	Snapshots    snapshot
	CateID       int `json:"CateId"`
	CateName     string
	Tags         string
}
type snapshot struct {
	Snapshot []string
}

// GetVideoListResposeEntity 获取视频信息列表返回
type GetVideoListResposeEntity struct {
	baseResposeEntity
	VideoList videoList
	Total     int
}
type videoList struct{
	Video []video
}

// UpdateVideoInfoResposeEntity 修改视频信息返回
type UpdateVideoInfoResposeEntity struct{
	baseResposeEntity
}

// DeleteVideoResposeEntity 删除视频返回
type DeleteVideoResposeEntity struct{
	baseResposeEntity
}