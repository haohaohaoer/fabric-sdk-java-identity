#去掉图形界面，一键执行
通过构造函数传入Json对象，获取ip/port，依次执行创建通道->加入entry节点->安装/实例化链码->加入peer节点

Json格式：
{
	"entry":[
	{
		"ip": xxx,
		"port":xxx
	}
	{
		"ip": xxx,
		"port":xxx
	}
	...
		]
	"peer":[
	{
		"ip": xxx,
		"port":xxx
	}
	{
		"ip": xxx,
		"port":xxx
	}
	{
		"ip": xxx,
		"port":xxx
	}
	...
		]
}
