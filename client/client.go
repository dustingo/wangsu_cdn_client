package client

import (
	"fmt"
	"io/ioutil"
	"strings"
	"sync"
	"wangsu_cdn_client/ccm"
	"wangsu_cdn_client/common/auth"
	"wangsu_cdn_client/common/util"
	"wangsu_cdn_client/config"

	"github.com/Wangsu-Cloud-Storage/wcs-go-sdk-v2/wos"
	"github.com/go-kit/log/level"
	"github.com/redmask-hb/GoSimplePrint/goPrint"
)

var WosClient *wos.WosClient
var AllFiles map[string]int64 = make(map[string]int64, 0)

// NewWosClient new wangsu wos client
func NewWosClient(ak, sk, endpoint, region string) (*wos.WosClient, error) {
	var err error
	if WosClient == nil {
		WosClient, err = wos.New(ak, sk, endpoint, wos.WithRegion(region))
		if err != nil {
			return nil, err
		}
	}
	return WosClient, err
}

//ListBucket 列出存储桶
func ListBucket() {
	input := &wos.ListBucketsInput{QueryLocation: true}
	output, err := WosClient.ListBuckets(input)
	if err != nil {
		if wosErr, ok := err.(wos.WosError); ok {
			level.Error(config.Logger).Log("statusCode", wosErr.StatusCode, "message", wosErr.Message)
		} else {
			level.Error(config.Logger).Log("list bucket error", err)
		}
		return
	}
	fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
	fmt.Printf("Owner.DisplayName:%s, Owner.ID:%s\n", output.Owner.DisplayName, output.Owner.ID)
	for index, val := range output.Buckets {
		fmt.Printf("Bucket[%d]-Name:%s,CreationDate:%s,EndPoint:%s,Region:%s\n", index, val.Name, val.CreationDate, val.Endpoint, val.Region)
	}

}

// ListObjects 列出桶内的objects 最大列举1000 没什么意义
func ListObjects(buckerName string) {
	input := &wos.ListObjectsInput{}
	input.Bucket = buckerName
	output, err := WosClient.ListObjects(input)
	if err != nil {
		if wosErr, ok := err.(wos.WosError); ok {
			level.Error(config.Logger).Log("statusCode", wosErr.StatusCode, "message", wosErr.Message)
		} else {
			level.Error(config.Logger).Log("list bucket error", err)
		}
		return
	} else {
		for index, val := range output.Contents {
			fmt.Printf("Index:%d %s,LastModified:%s, Size:%d\n", index, val.Key, val.LastModified, val.Size)
		}
	}
}

// PutObject 适合上传单独的文件 同样需要提前和wos的目录对应
func PutObject(objectKey string) {
	input := &wos.PutObjectInput{}
	input.Bucket = config.Sconfig.Global.Bucket
	input.Key = objectKey
	output, err := WosClient.PutObject(input)
	if err != nil {
		if wosErr, ok := err.(wos.WosError); ok {
			level.Error(config.Logger).Log("statusCode", wosErr.StatusCode, "message", wosErr.Message)
		} else {
			level.Error(config.Logger).Log("list bucket error", err)
		}
		return
	} else {
		level.Info(config.Logger).Log("ETag", output.ETag)
	}
}

// PutALLObject 上传目录下所有所有的文件
func PutAllLObject(dir string) {
	err := ReadDirFile(dir)
	if err != nil {
		level.Error(config.Logger).Log("read dir error", err)
		return
	}
	count := 10
	wg := sync.WaitGroup{}
	c := make(chan struct{}, count)
	defer close(c)
	bar := goPrint.NewBar(len(AllFiles))
	bar.SetNotice("上传进度:")
	bar.SetGraph(">")
	i := 0
	for k, v := range AllFiles {
		wg.Add(1)
		i++
		c <- struct{}{}
		go func(key string, v int64, wg *sync.WaitGroup) {
			defer wg.Done()
			// if v > 8192 {
			// 	fmt.Printf("suggestion: file size greater than 8MB,maybe using multipart upload could be batter. size = %d objectname=%s\n", (v / 1024), key)
			// }
			input := &wos.PutObjectInput{}
			input.Bucket = config.Sconfig.Global.Bucket
			input.Key = key
			_, err := WosClient.PutObject(input)
			if err != nil {
				level.Error(config.Logger).Log("put object error", err)
				return
			}
		}(k, v, &wg)
		bar.PrintBar(i)
	}
	wg.Wait()
	bar.PrintEnd(" 上传完成")
}

// ReadDirFile 边路目录下所有的文件路径
func ReadDirFile(dir string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, file := range files {
		if file.IsDir() {
			ReadDirFile(dir + "/" + file.Name())
		} else {
			//AllFiles = append(AllFiles, dir+"/"+file.Name())
			AllFiles[dir+"/"+file.Name()] = file.Size()
		}
	}
	return nil
}

// FlushCache 刷新缓存
func FlushCache(urls, dirs string) {
	urlsSet := []*string{}
	dirsSet := []*string{}
	var authConfig auth.AkskConfig
	ccmItemIdPurgeRequest := ccm.CcmItemIdPurgeRequest{}
	if urls == "" {
		d := strings.Split(dirs, ",")
		for _, dir := range d {
			dir := dir
			dirsSet = append(dirsSet, &dir)
		}
	} else if dirs == "" {
		u := strings.Split(urls, ",")
		for _, url := range u {
			url := url
			urlsSet = append(urlsSet, &url)
		}
	} else {
		d := strings.Split(dirs, ",")
		for _, dir := range d {
			dir := dir
			dirsSet = append(dirsSet, &dir)
		}
		u := strings.Split(urls, ",")
		for _, url := range u {
			url := url
			urlsSet = append(urlsSet, &url)
		}
	}
	fmt.Println("urls = ", urlsSet)
	fmt.Println("dirs = ", dirsSet)
	ccmItemIdPurgeRequest.SetUrls(urlsSet)
	ccmItemIdPurgeRequest.SetDirs(dirsSet)
	ccmItemIdPurgeRequest.SetUrlAction(config.Sconfig.Global.Urlaction)
	ccmItemIdPurgeRequest.SetDirAction(config.Sconfig.Global.Diraction)
	authConfig.AccessKey = config.Sconfig.Global.Ak
	authConfig.SecretKey = config.Sconfig.Global.Sk
	authConfig.EndPoint = "open.chinanetcenter.com"
	authConfig.Uri = "/ccm/purge/ItemIdReceiver"
	authConfig.Method = "POST"
	response := auth.Invoke(authConfig, ccmItemIdPurgeRequest.String())
	res := util.Format(response)
	fmt.Println(res)
}

func PreFetch(urls string) {
	urlsSet := []*string{}
	var authConfig auth.AkskConfig
	ccmItemIdFetchRequest := ccm.CcmItemIdFetchRequest{}
	u := strings.Split(urls, ",")
	for _, url := range u {
		url := url
		urlsSet = append(urlsSet, &url)
	}
	ccmItemIdFetchRequest.SetUrls(urlsSet)
	ccmItemIdFetchRequest.SetIsRange(0)
	authConfig.AccessKey = config.Sconfig.Global.Ak
	authConfig.SecretKey = config.Sconfig.Global.Sk
	authConfig.EndPoint = "open.chinanetcenter.com"
	authConfig.Uri = "/ccm/fetch/ItemIdReceiver"
	authConfig.Method = "POST"
	response := auth.Invoke(authConfig, ccmItemIdFetchRequest.String())
	res := util.Format(response)
	fmt.Println(res)
}
