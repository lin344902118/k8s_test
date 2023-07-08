# k8s_test
测试client_go的部分功能，如创建、删除、查看deployment

# 使用说明
## 使用docker镜像运行
1、将自己的.kube/config目录拷贝到项目下，以便dockerfile打包

2、使用go build -o main将项目编译成可执行文件，以便dockerfile打包

3、docker打包，运行docker build -t xxx:xx .

### 将其部署到k8s集群
1、./main 运行服务 打开服务地址。http://宿主机ip:8000/index

2、填写部署信息，点击部署

备注：如果是使用docker作为容器运行时，那么直接填写上面打包好的地址和版本号即可。
如果使用crictl，需要使用docker save -o xxx.tar xxx:xx将其导出。
然后使用ctr -n=k8s.io images import xxx.tar导入

3、通过kubectl get pod -o wide查看部署情况，并根据返回的ip进行访问
访问地址。返回的ip:8000/index。需要在集群机器上才能访问。如果想在集群外部访问，需要做映射。具体做法不详述

## 直接运行
1、go build -o main编译项目

2、./main运行。通过ip:8000/index可以访问

/index 主页可以部署k8sdeployment
/deployment 可以查看、删除deployment
