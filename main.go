package main

import (
	"flag"
	"fmt"
	"os"
	"wangsu_cdn_client/client"
	"wangsu_cdn_client/config"

	"github.com/go-kit/log/level"
)

func main() {
	// flag
	oss := flag.NewFlagSet("oss", flag.PanicOnError)
	putobject := oss.Bool("putobject", false, "bool默认false，上传单文件")
	putallobjects := oss.Bool("putallobjects", false, "bool默认false,上传目录下所有的文件")
	obj := oss.String("obj", "", "需要上传的文件或目录.单文件格式: key/name,如:client/projectx_v1.13.apk;目录格式: dirname,如: Release")
	cdn := flag.NewFlagSet("cdn", flag.PanicOnError)
	flush := cdn.Bool("flush", false, "bool默认false,刷新特定url和目录的节点缓存")
	prefetch := cdn.Bool("prefetch", false, "bool默认false,预取指定的url,使其首次访问即可命中缓存,详细进度需要登陆网页查看")
	urls := cdn.String("urls", "", "需要刷新的url地址,多url以','分隔")
	dirs := cdn.String("dirs", "", "需要刷新的dir地址,多dir以','分隔")
	// init client
	ak := config.Sconfig.Global.Ak
	sk := config.Sconfig.Global.Sk
	endpoint := config.Sconfig.Global.Endpoint
	region := config.Sconfig.Global.Region
	_, err := client.NewWosClient(ak, sk, endpoint, region)
	if err != nil {
		fmt.Println(err)
		return
	}
	//
	if num := len(os.Args); num < 2 {
		oss.Usage()
		cdn.Usage()
		return
	}
	switch os.Args[1] {
	case "oss":
		if err := oss.Parse(os.Args[2:]); err == nil {
			if *putobject && *putallobjects {
				level.Error(config.Logger).Log("parameter error", "putobject and putallobjects are mutual exclusive")
				return
			}
			if *putobject {
				if *obj == "" {
					level.Error(config.Logger).Log("parameter error", "obj is null")
					return
				} else {
					client.PutObject(*obj)
				}
			} else if *putallobjects {
				if *obj == "" {
					level.Error(config.Logger).Log("parameter error", "obj is empty")
					return
				} else {
					client.PutAllLObject(*obj)
				}
			}
		}
	case "cdn":
		if err := cdn.Parse(os.Args[2:]); err == nil {

			if *flush && *prefetch {
				level.Error(config.Logger).Log("parameter error", "flush and prefetch are mutual exclusive")
				return
			}
			if *flush {
				if *urls == "" && *dirs == "" {
					level.Error(config.Logger).Log("parameter error", "urls and dirs  cannot be empty at the same time")
					return
				} else {
					client.FlushCache(*urls, *dirs)
				}
			} else if *prefetch {
				if *urls == "" {
					level.Error(config.Logger).Log("parameter error", "urls cannot be empty")
					return
				}
				client.PreFetch(*urls)
			} else {
				level.Error(config.Logger).Log("parameter error", "if you want to flush cache or prefetch file ,you should set flush or prefetch as true")
			}
		}
	default:
		oss.Usage()
		cdn.Usage()
	}
}
