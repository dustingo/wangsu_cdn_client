package ccm

import (
	"github.com/alibabacloud-go/tea/tea"
)

type CcmItemIdPurgeRequest struct {
	// {"en":"If you need to purge several cached URLs. The submitted URL should meet the following format requirements:
	// 1) The URL must start with \'http://\' or \'https://\', url example: http://www.a.com/image/test.png.
	// 2) Each url has a maximum length of 1000 characters.
	// 3) The domain in the URL is must be the domain of the CDN service.
	// Note: URLs and dirs cannot be empty at the same time.","zh_CN":"要清理缓存的url集合，url的格式要求：
	// 1）URL 必须以\'http://\' 或 \'https://\' 开头，输入示例：http://www.a.com/image/test.png。
	// 2）每个url最大长度 2000 字符。
	// 3）每个url所在的域名必须是在我司加速的域名。
	// 4）刷新url每日不超过5000条（账号粒度可调，可联系技术支持调整），
	// 5）每次接口调用urls和dirs的总数不超过500条。
	// 注意：urls和dirs不能同时为空。"}
	Urls []*string `json:"urls,omitempty" xml:"urls,omitempty" type:"Repeated"`
	// {"en":"Need to purge the cached directory collection, the submitted directory should meet the following format requirements:
	// 1) Must start with 'http://domain name/' and end with '/', such as refreshing all files under the directory test, the format of the submitted directory is: http://www.a.com/test/.
	// 2) Each directory has a maximum length of 1000 characters.
	// 3) The domain name in directory must be the domain name of the CDN service.", "zh_CN":"指要清理缓存的目录集合，dir的格式要求：
	// 1）必须以'http://域名/'开始，以'/'结尾， 如刷新目录test下所有文件，输入格式为：http://www.a.com/test/；
	// 2）每个目录最大长度 1000 字符。
	// 3）每个目录所在的域名必须是在我司加速的域名。
	// 4）目录刷新默认过期，不支持目录删除。
	// 5）刷新目录每日不超过500条（账号粒度可调，可联系技术支持调整）
	// 注意：urls和dirs不能同时为空。"}
	Dirs []*string `json:"dirs,omitempty" xml:"dirs,omitempty" type:"Repeated"`
	// {"en":"1) default: default value, the url cache is processed using the pre-configured operation type of domain. When no value is assigned to this element or the element is not submitted, the configuration of domain is read by default.
	// 2) delete: ignore the operation type in the domain configuration, directly delete the cache file of the submitted url.
	// 3) expire: Ignore the operation type in the domain  configuration, and set the file with the cached commit url to expire. When there is a http access for the first time, a request is made to the origin server to check if the file is up-to-date. If the origin has a new file, the new version is directly pulled back from the source station and returned to the user. If there is no update, then the source station responds with http code 304, and the CDN provides the cache file in the edge node to the user.", "zh_CN":"仅对入参'urls'有效，指清理url缓存要以哪种类型操作：
	// 1）default：默认值，以频道预先配置好的操作类型处理url缓存，当不赋值或未传参时，默认取频道配置。
	// 2）delete：忽略频道配置里的操作类型，当前提交urls里的所有url，直接删除节点的缓存文件
	// 3）expire：忽略频道配置里的操作类型，当前提交urls里的所有url，将有缓存的节点置为过期，当第一次有访问时，回源站校验文件是否更新，有更新时从源站重新拉取新版本返回给客户，未有更新时则源站响应304，提供节点缓存文件给客户。"}
	UrlAction *string `json:"urlAction,omitempty" xml:"urlAction,omitempty"`
	// {"en":"It refers to the type of operation to clean up the dir cache, which is only valid for dirs elements. When no assignment or parameters are passed, the default channel configuration is selected. The optional values of this parameter are as follows:
	// 1) delete: ignore the type of operation in the channel configuration, delete the cache directory of the node directly.
	// 2) expire: ignore the type of operation in the channel configuration, set the cached node to expire, when the first visit, check whether the directory of the source station is updated, when there is an update, retrieve the new version from the source station to return to the customer, and when there is no update, the source station responds 304, providing the cached directory of the node to the customer.
	// Note: The use of directory delete function will result in all files cached in the directory empty, all files need to be retrieved, resulting in increased backsource bandwidth. Because of the high risk, the permission does not open by default. If you need to open this function, please contact the corresponding customer service technical support to apply for opening. Only the customer\'s account can be transferred to delete after opening.","zh_CN":"指清理dir缓存要以哪种类型操作，仅对dirs元素有效，当不赋值或未传参时，默认取频道配置。该参数可选值如下：
	// 1）delete：忽略频道配置里的操作类型，直接删除节点的缓存目录。
	// 2）expire：忽略频道配置里的操作类型，将有缓存的节点置为过期，当第一次有访问时，回源站校验目录是否更新，有更新时从源站重新拉取新版本返回给客户，未有更新时则源站响应304，提供节点缓存目录给客户。
	// 注：使用目录删除（delete）功能，会导致该目录下所有文件缓存全部清空，所有文件需要重新回源拉取，造成回源带宽增加。由于风险较大，该权限默认不开启，如需开启该功能，请联系对应的客服技术支持进行申请开启，开通后只有该客户的账号才能传入delete。"}
	DirAction *string `json:"dirAction,omitempty" xml:"dirAction,omitempty"`
}

func (s CcmItemIdPurgeRequest) String() string {
	return tea.Prettify(s)
}

func (s CcmItemIdPurgeRequest) GoString() string {
	return s.String()
}

func (s *CcmItemIdPurgeRequest) SetUrls(v []*string) *CcmItemIdPurgeRequest {
	s.Urls = v
	return s
}

func (s *CcmItemIdPurgeRequest) SetDirs(v []*string) *CcmItemIdPurgeRequest {
	s.Dirs = v
	return s
}

func (s *CcmItemIdPurgeRequest) SetUrlAction(v string) *CcmItemIdPurgeRequest {
	s.UrlAction = &v
	return s
}

func (s *CcmItemIdPurgeRequest) SetDirAction(v string) *CcmItemIdPurgeRequest {
	s.DirAction = &v
	return s
}

type CcmItemIdPurgeResponse struct {
	// {"en":"The status code of the task creation result, 1 means success, 0 means failure.","zh_CN":"表示任务创建结果的状态码，1表示成功，0表示失败"}
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	// {"en":"Content system response message after submitting the task.", "zh_CN":"表示任务提交后，系统的响应消息"}
	Message *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	// {"en":"After calling the API once and submitting the task successfully, the content system will return an itemId. This ID is the unique identifier for each submission. You can use itemId to batch query the status (success/failure) of the task.", "zh_CN":"调用一次接口并提交任务成功后，将返回一个iteamId，是当次提交任务的唯一标识，通过itemId可批量查询任务的状态（成功/失败）。"}
	ItemId *string `json:"itemId,omitempty" xml:"itemId,omitempty" require:"true"`
}

func (s CcmItemIdPurgeResponse) String() string {
	return tea.Prettify(s)
}

func (s CcmItemIdPurgeResponse) GoString() string {
	return s.String()
}

func (s *CcmItemIdPurgeResponse) SetCode(v int32) *CcmItemIdPurgeResponse {
	s.Code = &v
	return s
}

func (s *CcmItemIdPurgeResponse) SetMessage(v string) *CcmItemIdPurgeResponse {
	s.Message = &v
	return s
}

func (s *CcmItemIdPurgeResponse) SetItemId(v string) *CcmItemIdPurgeResponse {
	s.ItemId = &v
	return s
}

type Paths struct {
}

func (s Paths) String() string {
	return tea.Prettify(s)
}

func (s Paths) GoString() string {
	return s.String()
}

type Parameters struct {
}

func (s Parameters) String() string {
	return tea.Prettify(s)
}

func (s Parameters) GoString() string {
	return s.String()
}

type RequestHeader struct {
}

func (s RequestHeader) String() string {
	return tea.Prettify(s)
}

func (s RequestHeader) GoString() string {
	return s.String()
}

type ResponseHeader struct {
}

func (s ResponseHeader) String() string {
	return tea.Prettify(s)
}

func (s ResponseHeader) GoString() string {
	return s.String()
}
