### IPFS开发简介

#### 1、IPFS简介
* IPFS是什么

    星际文件系统IPFS（InterPlanetary File System）是一个面向全球的、点对点的分布式版本文件系统，目标是为了补充（甚至是取代）目前统治互联网的超文本传输协议（HTTP），将所有具有相同文件系统的计算设备连接在一起。原理用基于内容的地址替代基于域名的地址，也就是用户寻找的不是某个地址而是储存在某个地方的内容，不需要验证发送者的身份，而只需要验证内容的哈希，通过这样可以让网页的速度更快、更安全、更健壮、更持久。
    **以上内容来自巴比特官网介绍**
    
    IPFS是基于区块链技术的去中心化存储网络，实现了永久性存储。
    
* IPFS的火热

    Filecoin ICO是迄今为止涉及金额最大的ICO，仅一小时就募集了接近2亿美元，打破ICO记录，成为2017年最火爆的区块链项目。IPFS项目吸引了世界各地数字货币投资者和互联网从业者的密切关注，2018年比较值得期待的是IPFS主网的上线。然而就在前段时间，官方公布说上线时间延期到年底。
    注：Filecoin是IPFS激励层的加密数字货币（即代币），有点类似于以太坊平台上的以太币。
    
    IPFS官网：https://ipfs.io/
    Filecoin官网：https://filecoin.io/

* IPFS的应用案例
  GitHub有两款开源项目，且有对应网址，分别是音乐播放器和视频播放器。
    * IPFS音乐播放器
    ![1.IPFS音乐播放器](media/15269926603540/1.IPFS%E9%9F%B3%E4%B9%90%E6%92%AD%E6%94%BE%E5%99%A8.jpg)

    
        IPFS音乐播放器网址：https://diffuse.sh/
        GitHub地址：https://github.com/icidasset/diffuse

    
    * IPFS视频在线播放器
    ![2.IPFS视频在线播放器首页](media/15269926603540/2.IPFS%E8%A7%86%E9%A2%91%E5%9C%A8%E7%BA%BF%E6%92%AD%E6%94%BE%E5%99%A8%E9%A6%96%E9%A1%B5.png)

    ![3.IPFS视频在线播放器播放电影](media/15269926603540/3.IPFS%E8%A7%86%E9%A2%91%E5%9C%A8%E7%BA%BF%E6%92%AD%E6%94%BE%E5%99%A8%E6%92%AD%E6%94%BE%E7%94%B5%E5%BD%B1.png)

    
   号称是国内第一个IPFS应用
   IPFS视频在线播放器网址：http://www.ipfs.guide/
    GitHub地址：https://github.com/download13/ipfstube
        
    可用于测试的电影视频Hash（这里只列举两部部）：
    神秘巨星：QmWBbKvLhVnkryKG6F5YdkcnoVahwD7Qi3CeJeZgM6Tq68
    盗梦空间：QmQATmpxXvSiQgt9c9idz9k3S3gQnh7wYj4DbdMQ9VGyLh

#### 2、IPFS的安装
既然IPFS这么牛，有必要了解一下IPFS的开发，先从环境搭建开始。

* IPFS Desktop

    当然，可以直接安装IPFS节点桌面管理软件来体验一下。该软件可以方便地运行和管理自己的节点，查看IPFS节点资源信息，支持二次开发。该项目是Node.js编写的，已开源。
    ![4.IPFS节点桌面管理软件](media/15269926603540/4.IPFS%E8%8A%82%E7%82%B9%E6%A1%8C%E9%9D%A2%E7%AE%A1%E7%90%86%E8%BD%AF%E4%BB%B6.png)
    
    GitHub地址：
    https://github.com/ipfs-shipyard/ipfs-desktop

* Go-IPFS
    进入ipfs的官网，找到并切换到"Install"页面，点击"Download IPFS for your platform"，会跳转到如下网址（可能需要翻墙）：
    https://dist.ipfs.io/#go-ipfs

    ![6.下载Go-IPFS](media/15269926603540/6.%E4%B8%8B%E8%BD%BDGo-IPFS.png)

    如果打不开，可以试着去Github查看安装方法
GitHub地址：https://github.com/ipfs/go-ipfs

    将网页下载好的文件解压出来，下载的是"go-ipfs_v0.4.13_darwin-amd64.tar.gz"，解压出来的是go-ipfs文件夹，
    然后打开终端，进入到go-ipfs文件夹根目录，复制ipfs文件到系统的bin目录，然后可以通过"ipfs version"来检查，操作如下：
    
    ```
    wenzildeiMac:go-ipfs wenzil$ pwd
/Users/wenzil/Desktop/study/go-ipfs
wenzildeiMac:go-ipfs wenzil$ ls
LICENSE		README.md	build-log	install.sh	ipfs
wenzildeiMac:go-ipfs wenzil$ cp ipfs /usr/local/bin/ipfs
    wenzildeiMac:go-ipfs wenzil$ ipfs version
    ipfs version 0.4.13
    ```
    
    

#### 3、IPFS节点的创建、查看和启动
**以下是安装官网的Go-IPFS步骤下实验的**

* 创建节点

    然后进入到当前用户的根目录，并"ipfs init"命令创建节点，可以用"open ./"打开创建节点生成的".ipfs"目录。
    
    ```
    wenzildeiMac:go-ipfs wenzil$ cd ~/
    wenzildeiMac:~ wenzil$ pwd
    /Users/wenzil
    wenzildeiMac:~ wenzil$ ipfs init
    initializing IPFS node at /Users/wenzil/.ipfs
    generating 2048-bit RSA keypair...done
    peer identity: QmWH4xmGBznY9CJ4SjxpxWqGwMwrFtPdjgeweLCfJxW8j9
    to get started, enter:
    
    	ipfs cat /ipfs/QmS4ustL54uo8FzR9455qaxZwuMiUhyvMcX9Ba8nUH4uVv/readme
    
    wenzildeiMac:~ wenzil$ cd .ipfs
    wenzildeiMac:.ipfs wenzil$ open ./
    ```
![6.IPFS节点的创建](media/15269926603540/6.IPFS%E8%8A%82%E7%82%B9%E7%9A%84%E5%88%9B%E5%BB%BA.png)


* 查看节点
    
    ```
    wenzildeiMac:.ipfs wenzil$ ipfs id
    {
    	"ID": "QmWH4xmGBznY9CJ4SjxpxWqGwMwrFtPdjgeweLCfJxW8j9",
    	"PublicKey": "CAASpgIwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCt98O5zNptBtmSF6xYdOHAMunQsXE9rAmgqBXVPLJk+AoaBjNiAipB+sTXKwaj8stp9c3iuSzz10+dyYI38bq4TPDCsHnQ3PjKSgZwEaT0M6pJhGqJcSbLY57CjQOEwgFe+tw3b4hcoDvvNJMG0dZZg1r9xnfevBz0DNewlj2vtviSTQL6r/ZXlD04GsytxTSQ2pzMJTG3QmYyokeZ37DT7Vwa+IJjDCiDJC5BlZ204zunZBx3tqWUn2Hr2NtEX+4YDdX4SBHa0ZDNqZmzLRe5YlAXjVI0ONHs+BVsC1v1L5we52iSYCSgVHoxRP/pa9EZdHMvgs9QeAa5jO90yxgTAgMBAAE=",
    	"Addresses": null,
    	"AgentVersion": "go-ipfs/0.4.13/",
    	"ProtocolVersion": "ipfs/0.1.0"
    }
    ```

* 修改IPFS默认存储空间
存储节点默认存储大小为10G，如果你想修改存储空间，可以修改config配置文件，找到StorageMax并修改

    ```
  wenzildeiMac:~ wenzil$ export EDITOR=/usr/bin/vim
wenziliangdeMac-mini:~ wenzil$ ipfs config edit  
    ```   
    
       
    
* 启动和同步节点服务器
执行"ipfs daemon"命令，可以同步节点数据到IPFS网络。

    ```
    wenzildeiMac:.ipfs wenzil$ ipfs daemon
    Initializing daemon...
    Adjusting current ulimit to 2048...
    Successfully raised file descriptor limit to 2048.
    Swarm listening on /ip4/127.0.0.1/tcp/4001
    Swarm listening on /ip4/192.168.1.100/tcp/4001
    Swarm listening on /ip6/::1/tcp/4001
    Swarm listening on /p2p-circuit/ipfs/QmWH4xmGBznY9CJ4SjxpxWqGwMwrFtPdjgeweLCfJxW8j9
    Swarm announcing /ip4/100.64.9.179/tcp/55898
    Swarm announcing /ip4/127.0.0.1/tcp/4001
    Swarm announcing /ip4/192.168.1.100/tcp/4001
    Swarm announcing /ip6/::1/tcp/4001
    API server listening on /ip4/127.0.0.1/tcp/5001
    Gateway (readonly) server listening on /ip4/127.0.0.1/tcp/8080
    Daemon is ready
    ```

#### 4、IPFS运行体验
    
* IPFS查看ReadMe
可以回去查看执行"ipfs init"命令的时候，返回的内容。打开一个新的终端，操作如下：

    ```
    wenzildeiMac:~ wenzil$ ipfs cat /ipfs/QmYwAPJzv5CZsnA625s3Xf2nemtYgPpHdWEz79ojWnPbdG/readme
    Hello and Welcome to IPFS!
    
    ██╗██████╗ ███████╗███████╗
    ██║██╔══██╗██╔════╝██╔════╝
    ██║██████╔╝█████╗  ███████╗
    ██║██╔═══╝ ██╔══╝  ╚════██║
    ██║██║     ██║     ███████║
    ╚═╝╚═╝     ╚═╝     ╚══════╝
    
    If you're seeing this, you have successfully installed
    IPFS and are now interfacing with the ipfs merkledag!
    
     -------------------------------------------------------
    | Warning:                                              |
    |   This is alpha software. Use at your own discretion! |
    |   Much is missing or lacking polish. There are bugs.  |
    |   Not yet secure. Read the security notes for more.   |
     -------------------------------------------------------
    
    Check out some of the other files in this directory:
    
      ./about
      ./help
      ./quick-start     <-- usage examples
      ./readme          <-- this file
      ./security-notes
    ```
* Web管理界面：
IPFS节点服务器启动后，可以用浏览器访问：http://localhost:5001/webui
有本地配置、节点连接、本地节点文件等信息，如图：
![7.打开webui](media/15269926603540/7.%E6%89%93%E5%BC%80webui.png)


#### 5、设置跨域资源共享
当我们在前端通过js接口操作ipfs时，会遇到跨域资源访问问题，可以在终端执行以下配置来解决：

```
ipfs config --json API.HTTPHeaders.Access-Control-Allow-Methods '["PUT","GET", "POST", "OPTIONS"]'
ipfs config --json API.HTTPHeaders.Access-Control-Allow-Origin '["*"]'
ipfs config --json API.HTTPHeaders.Access-Control-Allow-Credentials '["true"]'
ipfs config --json API.HTTPHeaders.Access-Control-Allow-Headers '["Authorization"]'
ipfs config --json API.HTTPHeaders.Access-Control-Expose-Headers '["Location"]'
```

#### 6、IPFS的基本操作
* 新建文件

    ```
    wenziliangdeMac-mini:ifps_test wenzil$ vi test.txt
    wenziliangdeMac-mini:ifps_test wenzil$ cat test.txt 
    IPFS测试文件
    ```

* 添加文件到IPFS节点

    ```
    wenziliangdeMac-mini:ifps_test wenzil$ ipfs add test.txt 
    added QmSVKEwPBTzw5QLzGUE8oN8J1r4cadMeieSw4Co1ozm2Ab test.txt
    ```
    添加文件到IPFS节点后，返回了文件的哈希值

* 查看IPFS节点的文件

    ```
    wenziliangdeMac-mini:ifps_test wenzil$ ipfs cat QmSVKEwPBTzw5QLzGUE8oN8J1r4cadMeieSw4Co1ozm2Ab
    IPFS测试文件
    ```
    注意：此时的文件只是添加到了本地的IPFS节点，读取的是本地数据，可以通过如下网址查看。
http://localhost:8080/ipfs/QmSVKEwPBTzw5QLzGUE8oN8J1r4cadMeieSw4Co1ozm2Ab
    
    通过"ipfs daemon"命令启动节点服务器，会将本地节点文件同步到外网，同步成功后，这时就可以访问如下网址来查看（如果访问失败，试着翻墙看看）。
    https://ipfs.io/ipfs/QmSVKEwPBTzw5QLzGUE8oN8J1r4cadMeieSw4Co1ozm2Ab
    ![9.查看IPFS文件内容](media/15269926603540/9.%E6%9F%A5%E7%9C%8BIPFS%E6%96%87%E4%BB%B6%E5%86%85%E5%AE%B9.png)

    
    如果访问成功的话，说明已经存储到IPFS网络中。目前IPFS网络暂未加入代币Filecoin机制，所以存储读取文件免费，速度也比较慢。    
    

* 下载IPFS节点的文件

    ```
    wenziliangdeMac:ifps_test wenzil$ ipfs get QmSVKEwPBTzw5QLzGUE8oN8J1r4cadMeieSw4Co1ozm2Ab
    Saving file(s) to QmSVKEwPBTzw5QLzGUE8oN8J1r4cadMeieSw4Co1ozm2Ab
     25 B / 25 B [========================================================] 100.00% 0s
    ```
    查看当前目录，发现多了一个"QmSV..."的文件"
    

#### 7、IPFS与以太坊DApp的结合

待续。。。


