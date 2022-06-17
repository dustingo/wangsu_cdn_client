## 基于对象存储、CDN API的上传、刷新工具
平台：Linux、Windows
> 基础功能
- 上传单文件
- 批量上传文件
- 刷新url及目录
- 文件预取 
- 大文件分片上传(todo)

>基本用法
```shell
# 下载，解压，填写配置文件
./wangsu_cdn_client --help
Usage of oss:
  -obj string
        需要上传的文件或目录.单文件格式: key/name,如:client/projectx_v1.13.apk;目录格式: dirname,如: Release
  -putallobjects
        bool默认false,上传目录下所有的文件
  -putobject
        bool默认false，上传单文件
Usage of cdn:
  -dirs string
        需要刷新的dir地址,多dir以','分隔
  -flush
        bool默认false,刷新特定url和目录的节点缓存
  -prefetch
        bool默认false,预取指定的url,使其首次访问即可命中缓存,详细进度需要登陆网页查看
  -urls string
        需要刷新的url地址,多url以','分隔
```
---
> 配置文件
配置文件默认读取目录下的config.toml
```toml
[global]
ak = ""  #账号ak
sk = ""  #账号sk
endpoint = "https://s3-cn-east-1.wcsapi.com" #endpoint地址,可在对象存储的概览中获取
region = "cn-east-1" #同上
bucket = ""          # 分配的bucket名称
urlaction = "delete" # url刷新策略 default、delete、expire
diraction = "expire" # dir刷新策略 delete、expire
```
>使用说明 

主要分为oss对象存储和CDN两部分
- OSS
```shell
# 上传单文件,注意路径，如果目录不存在的话会自动创建；上传成功返回对象的ETag
./wangsu_cdn_client oss --putobject --obj=test/single.txt 

# 上传test(更新时需要和oss中的目录一致)目录下所有文件，会显示上传进度上传进度:[>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>] 100%     3/3
./wangsu_cdn_client oss --putallobjects --obj=test
```
- CDN
```shell
# 缓存刷新需要指定文件urls或目录dirs，二者不能同时为空
# 刷新以下两个url不刷新dir，则可用不传--dirs，只刷新dirs同理
./wangsu_cdn_client cdn --flush --urls=http://xxx.xxx.com/test/xixi/xixi.txt,http://xxx.xx.com/test/haha/haha.txt

# 同时刷新urls和dirs
./wangsu_cdn_client cdn --flush --urls=http://xxx.xx.com/test/xixi/xixi.txt,http://xxx.xx.com/test/haha/haha.txt --dirs=http://xxx.xx.com/test/haha,http://xxx.xx.com/test/xixi

# 大文件预取
./wangsu_cdn_client cdn --prefetct uris=http://xxx.xx.com/xx/xx/xxx.zip
```
> 注意  
     **文件预取针对于大文件，网宿API的返回成功多数只是标识创建目标任务成功，最好执行完毕后可以去平台查看具体的相关进度或联系运维同学帮忙确认。**