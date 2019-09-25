### 大体的目录结构设计
---
```
|-- _webFile
		|-- _js
		|-- _resources
		|-- _css
		|-- _相关目录
		|-- index.html
|-- _server
		|-- _handle						// 接收消息并作出相关操作，操作有operate文件夹内提供
				|-- _每个handle一个目录
				|-- serve_mux.go					// 总handle
		|-- _operate								// 系统操作
				|-- modify_config.go
				|-- open_bard.go
				|-- stop_bard.go
				|-- restart.go
				|-- _office_plugin_config					// 官方插件管理页面
				-------
				|-- _隧道相关操作			  				// 将来bard版本 隧道的安装、设置、开启、关闭
				|-- 模式控制.go					// 手动、全局、pac、自动（先本地直连，后代理，如果成功记录自动pac文件）
		|-- server.go							// 启动server自动打开游览器界面
|-- _bard										// 具体目录结构由bard项目及其子项目决定
		|-- bard.sh							// bard的脚本文件
		|-- ...
|-- _tun2socks
		|-- ...
```