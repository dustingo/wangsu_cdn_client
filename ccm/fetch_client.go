package ccm

import (
	"github.com/alibabacloud-go/tea/tea"
)

type CcmItemIdFetchRequest struct {
	// {'en':'If you need to prefetch several cached URLs. The submitted URL should meet the following format requirements:
	// 1) The URL must start with http:// or https://, url example: http://www.a.com/image/test.png.
	// 2) Each url has a maximum length of 2000 characters.
	// 3) The domain in the URL is must be the domain of the CDN service.
	// 4) If the url contains special characters such as Chinese characters and spaces, our system will generate multiple push tasks. In addition to pushing the original URL, these special characters will be converted int32o ASCII codes and pushed. If you only want to clean up the transcoded URL, you need to use UTF-8 to complete the transcoding before submitting the URL, and then submit the escaped url to our system.'
	// , 'zh_CN':'要预取到CDN节点的url集合，url格式说明：
	// 1）URL 必须以 http:// 或 https:// 开头，输入示例：http://www.a.com/image/test.png。
	// 2）每个url最大长度 2000 字符。
	// 3）每个url所在的域名必须是在我司加速的域名且有预取权限。
	// 4）url中如果包含中文字符，则提交的url需要是中文转义后的url，采用utf-8方式转义。
	// 5）每日不超过20000条，不超过200G文件大小（账号粒度可调，联系技术支持人员调整）。
	// 6）每次接口调用url的总数不超过400条。'}
	Urls []*string `json:"urls,omitempty" xml:"urls,omitempty" require:"true" type:"Repeated"`
	// {"en":"Only prefetch a range segment of the file header. The user get the file from the beginning, and they will select quickly their int32erested. If the file header is cached, the first pack time of the user's http request will be short.This feature allows users to filter content faster. For example, if a file has 200MB, only the size of the file 0~range is prefetched, instead of prefetching the entire file. Each account can be configured with a size of the range. If you need to modify the size, please contact us. If this element is assigned a value of 1, the default prefetch is 0~512KB.", "zh_CN":"是否需要预取range段。
	//
	// 1）默认为0，表示预取完整的文件；
	// 2）1表示预取文件0~512KB的range段（账号粒度可调，联系技术支持人员调整）。"}
	IsRange *int32 `json:"isRange,omitempty" xml:"isRange,omitempty"`
}

func (s CcmItemIdFetchRequest) String() string {
	return tea.Prettify(s)
}

func (s CcmItemIdFetchRequest) GoString() string {
	return s.String()
}

func (s *CcmItemIdFetchRequest) SetUrls(v []*string) *CcmItemIdFetchRequest {
	s.Urls = v
	return s
}

func (s *CcmItemIdFetchRequest) SetIsRange(v int32) *CcmItemIdFetchRequest {
	s.IsRange = &v
	return s
}

type CcmItemIdFetchResponse struct {
	// {'en':'The status code of the task creation result, 1 means success, 0 means failure.', 'zh_CN':'表示任务创建结果的状态码，1表示任务提交成功，0表示任务提交失败'}
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	// {'en':'Content system response message after submitting the task.', 'zh_CN':'表示任务提交后，系统的响应消息'}
	Message *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	// {'en':'After calling the API once and submitting the task successfully, the content system will return an itemId. This ID is the unique identifier for each submission. You can use itemId to batch query the status (success/failure) of the task.', 'zh_CN':'调用一次接口并提交任务成功后，将返回一个itemId，是当次提交任务的唯一标识，通过itemId可批量查询任务的状态（成功/失败）。'}
	ItemId *string `json:"itemId,omitempty" xml:"itemId,omitempty" require:"true"`
}

func (s CcmItemIdFetchResponse) String() string {
	return tea.Prettify(s)
}

func (s CcmItemIdFetchResponse) GoString() string {
	return s.String()
}

func (s *CcmItemIdFetchResponse) SetCode(v int32) *CcmItemIdFetchResponse {
	s.Code = &v
	return s
}

func (s *CcmItemIdFetchResponse) SetMessage(v string) *CcmItemIdFetchResponse {
	s.Message = &v
	return s
}

func (s *CcmItemIdFetchResponse) SetItemId(v string) *CcmItemIdFetchResponse {
	s.ItemId = &v
	return s
}
