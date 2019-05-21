#修改crypto-config、configtx文件:

1个Orderer、1个Org、5个Peer、1个User



#修改build.sh脚本，重新生成证书:

提前添加bin目录(包含二进制文件)，不然无法执行cryptogen、configtxgen命令；将生成的config、crypto-config目录放到network_resource里面(和docker-compose.yml指定的路径一致)



#在docker-compose.yml里将ca证书换为新生成的:

crypto-config/peerOrganizations/org1.example.com/ca/..._sk



#在docker-compose.yml里新增节点配置:

1个ca、1个orderer、5个peer(peer0—peer4)，修改名字和端口号



#build.sh启动容器(1+1+5)



#修改Config.java文件:

将ORG2的配置注释掉；修改ORG1节点信息（注意端口号）



#新建Manager类，合并main函数

将创建通道、新增节点、链码安装和实例化通过函数调用



#docker-compose文件增加couchdb，验证fabcar链码可用couchdb



#完成链码修改（基于marble）

可安装并且实例化，链码名称为identity.go

#java添加数据前端及功能实现

通过按钮获取输入的数据（name,address,type),调用链码的initIdentity函数进行添加

#java查询数据前端及功能实现

通过名称name，查询数据（name,address,type),调用链码的quaryIdentityByName函数进行查询



#################################



执行过程：

1）进入v8/network目录，执行./build.sh建网（可用docker ps查看已生成的容器，包括5个peer,5个couchdb,1个orderer,1个ca）

2）进入v8/java目录，输入mvn install命令（需安装maven），生成target目录，进入后将block-java-sdk-0.0.1-SNAPSHOT-jar-with-dependencies.jar文件复制到v8/network_resources目录下，可将其重命名为v8.jar

3）进入v8/network_resources目录，执行java -cp v8.jar org.app.network.Manager命令，随即弹出图形界面，按顺序操作即可（节点数最多为5；链码实例化过程较长）
