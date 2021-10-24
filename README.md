# My_Cmdb
- 项目骨架介绍
    项目组织的核心思路是: 每个业务模块尽量独立, 方便后期扩展和迁移成独立的服务
		cmd        程序cli工具包
		conf       程序配置对象
		protocol   程序监听的协议
		version    程序自身的版本信息
		pkg        业务领域包
		- host
			- model     业务需要的数据模型
			- interface 业务接口(领域方法)
			- impl      业务具体实现(接口实现)
		- mysql
		- lb
		- ...
		main 程序入口文件