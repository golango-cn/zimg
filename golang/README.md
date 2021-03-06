###docker zimg 安装

1.下载镜像
docker pull babymumu/zimg

2.运行
docker run -itd -p 4869:4869 -p 11211:11211 -v /usr/local/src/zimg:/root/zimg-master/bin/img --name zimg babymumu/zimg

3.进到容器内启动服务

cd /root/zimg-master/bin
./zimg conf/zimg.lua


