## docker安装 & 使用
1. 下载安装包 & 安装

        1. 官网url:
        2. 选择免费版 & 系统
        3. 国内镜像(版本更新会慢一点)：get.daocloud.io
        4. 配置镜像加速器（解决docker需要翻墙）
2. docker命令使用 & nginx安装 
        
        1. docker -v
        2. docker run -it alpine sh (linux镜像下载)
            [win10 docker: Error response from daemon: Get  net/http: TLS handshake timeout
                异常解决方案：https://blog.csdn.net/pop0fa/article/details/80029739]如下图
        3. docker run -d -p 80:80 nginx （docker ps 查看运行进程）
            -d:不要退出在后台一直运行
            -p:端口号
            在浏览器中验证：localhost:80
        4. 查看docker安装应用：docker images
        5. 结束进程：docker kill "进程Id"
        6. 查看日志：docker logs "进程Id"
    ![Image text](lang/img/setting.png)
3. docker Elasticsearch安装

        1. docker run -d -p 9200:9200 elasticsearch:6.5.0
        2. 在浏览器中访问：localhost:9200
4. 在IDEA中测试Elasticsearch接口 & 操作数据

        准备工作：
            1. 添加header：Content-Type=application/json
        1. demo：Path=index/type/id (index:数据库，type:table)
        2. PUT|POST —> Path=imooc/course/1 (PUT|POST创建和修改数据，使用POST可省略Id)
            添加数据：{ "name":"golang","instructor":"ccmouse" }
            执行 -> 查看执行结果
            2.2. 查询2插入结果method换成GET —> 执行后查看结果
        3. 查询所有数据：http://localhost:9200/imooc/course/_search/
        4. 查看表结构：/imooc/course/_mapping
5. 向Elasticsearch写入数据（搜索框架：elasticsearch go client）

        1. url：https://olivere.github.io/elastic/
        2. go get github.com/olivere/elastic (安装jar,注意版本对应)
        3. 根据官方文档编写代码
6. 清理Elasticsearch数据
    
        1. docker ps
        2. docker kill ID
        3. clear
        4. docker run -d -p 9200:9200 elasticsearch:6.5.0 （5.6.6）
7. 
8. 