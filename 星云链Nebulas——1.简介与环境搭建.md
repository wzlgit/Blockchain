### 星云链Nebulas——1.简介与环境搭建

#### 1、星云链的简介

* 星云链世界观

    星云链Nebulas（也称星云链NAS），是一个为区块链数据提供去中心化搜索框架的开源公有链。作为全球首个区块链搜索引擎，被称之为“区块链世界中的谷歌”。

    ![2.星云链世界观示意图](media/15281819133101/2.%E6%98%9F%E4%BA%91%E9%93%BE%E4%B8%96%E7%95%8C%E8%A7%82%E7%A4%BA%E6%84%8F%E5%9B%BE.jpg)


    目前有各种主链上有各种去中心化应用DApp，互不联通互不兼容，用户去寻找新应用非常不方便。就像互联网没有出现搜索引擎之前，人们在网上搜集资料很麻烦。后来出现了雅虎、谷歌等搜索引擎，解决了这个问题。
    
    星云链（NAS）意图构建一个能够量化价值尺度、具备自进化能力，并能促进区块链生态建设的区块链系统，主要包括：定义价值尺度的星云指数 Nebulas Rank(NR) ，支持核心协议和智能合约链上升级的星云原力 Nebulas Force(NF) ，开发者激励协议 Developer Incentive Protocol(DIP) ，贡献度证明共识算法 Proof of Devotion(PoD)，去中心化应用的搜索引擎 。通过星云指数Nebulas Rank(NR) 、星云原力 Nebulas Force(NF)、贡献度证明共识算法Proof of Devotion(PoD)，去中心化应用搜索引擎四个模块搭建能实现价值尺度量化、具备自进化能力，并促进区块链生态建设的区块链系统。
    
* 星云链代币NAS

    星云币代币为NAS。星云币作为星云链的原生代币，用于支付星云链上的交易手续费和计算服务费。

    2017年12月，星云链公开发行代币NAS，总量为1亿枚，目前流通量为3000万枚。已于2017年9月28日完成众筹，众筹价为2美元，目前Nas价格为7.02美元，较之上涨了3.5倍，在所有币种中排名52名。

    ![1.星云币排行](media/15281819133101/1.%E6%98%9F%E4%BA%91%E5%B8%81%E6%8E%92%E8%A1%8C.png)

  不同于众多主链的主网集中"六月上线"，星云链于2018年3月30日已经上线主网。

* 星云链三大核心技术

  1、星云指数
  星云指数 Nebulas Rank (NR)，核心的排序算法，已开源。以流动性、传播性、互操作性等体现数据互动关系的因素为基础，可以用来衡量地址、智能合约、去中心化应用等对象的影响力。
  
  2、自进化能力
  星云原力（Nebulas Force）可以为星云链及其应用提供自进化能力，有了NF，开发者可以游刃有余地修改、升级和修复，无需硬分叉，应对变化和风险更加自如。
  
  3、原生激励
  激励是区块链进化第一原动力。星云激励包括开发者激励协议（Developer Incentive Protocol, DIP）和贡献度证明（Proof of Devotion, PoD）。开发者激励协议构建一个反馈循环，激励开发者研发高质量的去中心化应用。通过贡献度证明，用户可以依照星云指数排名参与记账，从中获得星云币（NAS）。
  

* 星云链创始人团队

    徐义吉，星云链（Nebulas）& 小蚁（AntShares）创始人，前蚂蚁金服区块链平台部负责人，前谷歌反作弊小组成员，同济大学计算机学士。
    
    钟馥百，星云链（Nebulas）联合创始人，前蚂蚁金服区块链平台部架构师，前海豚浏览器高级研发总监、游戏业务负责人，毕业于华中科技大学。

    王冠，星云链（Nebulas）和小蚁（NEO）联合创始人，OpenIP&IP圈发起人，毕业于东南大学，区块链行业连续创业者。

    ![3.星云链创始人团队](media/15281819133101/3.%E6%98%9F%E4%BA%91%E9%93%BE%E5%88%9B%E5%A7%8B%E4%BA%BA%E5%9B%A2%E9%98%9F.png)

* 星云链相关链接

    [星云链官网](https://nebulas.io/cn/index.html)
    [星云链非技术白皮书](https://nebulas.io/docs/NebulasWhitepaperZh.pdf)
    [星云链技术白皮书](https://nebulas.io/docs/NebulasTechnicalWhitepaperZh.pdf)
    [星云链Github地址](https://github.com/nebulasio/go-nebulas)
    [星云链浏览器](https://explorer.nebulas.io/#/)
    [星云链钱包](https://wallet.nasscan.io/)
    [领取测试网络星云币](https://testnet.nebulas.io/claim/)
    [星云宠物卡DApp](https://petcard.ruterly.com/#/)
    
#### 2、星云链的安装要求

星云链现阶段只能在Mac和Linux上运行，后续会推出windows版本。

**注意：以下开始介绍星云链的环境搭建，均在MacOS上面操作，Linux系统安装请参考官网Github**

安装前提条件：

* 安装Homebrew

```
/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
```

* 安装Go（Go版本号必须>=1.9.2）：

```
# 安装
brew install go

# 配置环境变量
export GOPATH=/path/to/workspace
```


#### 3、编译星云链

1、下载源码

```
mkdir -p $GOPATH/src/github.com/nebulasio
cd $GOPATH/src/github.com/nebulasio

# 下载源码
git clone https://github.com/nebulasio/go-nebulas.git

# 进入项目目录
cd go-nebulas

# 切换到最稳定的master分支
git checkout master
```

终端操作如下：

```
wenzildeiMac:Study wenzil$ mkdir -p $GOPATH/src/github.com/nebulasio
wenzildeiMac:Study wenzil$ cd $GOPATH/src/github.com/nebulasio
wenzildeiMac:nebulasio wenzil$ git clone https://github.com/nebulasio/go-nebulas.git
Cloning into 'go-nebulas'...
remote: Counting objects: 19448, done.
remote: Compressing objects: 100% (99/99), done.
remote: Total 19448 (delta 73), reused 107 (delta 50), pack-reused 19298
Receiving objects: 100% (19448/19448), 169.79 MiB | 60.00 KiB/s, done.
Resolving deltas: 100% (13666/13666), done.
wenzildeiMac:nebulasio wenzil$ cd go-nebulas
wenzildeiMac:go-nebulas wenzil$ git checkout master
Already on 'master'
Your branch is up-to-date with 'origin/master'.
```


2、安装rocksdb依赖库

```
brew install rocksdb
```

3、安装Go依赖库

在Go-Nebulas中，Go的三方库都通过Dep来做管理，Dep版本号需要>=0.3.1。

```
brew install dep
brew upgrade dep
```

4、下载Go第三方库

切换到Go-Nebulas项目根目录，然后使用Dep来下载三方库。

```
wget http://ory7cn4fx.bkt.clouddn.com/vendor.tar.gz
cd $GOPATH/src/github.com/nebulasio/go-nebulas
tar zxf vendor.tar.gz
```

终端操作如下：

```
wenzildeiMac:go-nebulas wenzil$ wget http://ory7cn4fx.bkt.clouddn.com/vendor.tar.gz
--2018-06-05 17:32:18--  http://ory7cn4fx.bkt.clouddn.com/vendor.tar.gz
正在解析主机 ory7cn4fx.bkt.clouddn.com (ory7cn4fx.bkt.clouddn.com)... 14.215.166.109, 61.140.13.201, 14.215.166.199, ...
正在连接 ory7cn4fx.bkt.clouddn.com (ory7cn4fx.bkt.clouddn.com)|14.215.166.109|:80... 已连接。
已发出 HTTP 请求，正在等待回应... 200 OK
长度：19022460 (18M) [application/x-gzip]
正在保存至: “vendor.tar.gz”

vendor.tar.gz       100%[===================>]  18.14M   948KB/s  用时 20s     

2018-06-05 17:32:38 (946 KB/s) - 已保存 “vendor.tar.gz” [19022460/19022460])

wenzildeiMac:go-nebulas wenzil$ cd $GOPATH/src/github.com/nebulasio/go-nebulas
wenzildeiMac:go-nebulas wenzil$ tar zxf vendor.tar.gz
```

5、安装Chrome V8依赖库

星云虚拟机目前依赖于Chrome的V8引擎，为了大家使用方便，我们已经为Mac OSX和Linux编译好了V8的动态库。运行如下指令就可以完成安装。

```
cd $GOPATH/src/github.com/nebulasio/go-nebulas
make deploy-v8
```

终端操作如下：

```
wenzildeiMac:go-nebulas wenzil$ cd $GOPATH/src/github.com/nebulasio/go-nebulas
wenzildeiMac:go-nebulas wenzil$ make deploy-v8
install nf/nvm/native-lib/*.dylib /usr/local/lib/
```

6、编译可执行文件

完成所有上述依赖库的安装后，进入Go-Nebulas根目录编译星云链的可执行文件。

```
cd $GOPATH/src/github.com/nebulasio/go-nebulas
make build
```

终端操作如下：

```
wenzildeiMac:go-nebulas wenzil$ cd $GOPATH/src/github.com/nebulasio/go-nebulas
wenzildeiMac:go-nebulas wenzil$ make deploy-v8
install nf/nvm/native-lib/*.dylib /usr/local/lib/
wenzildeiMac:go-nebulas wenzil$ cd $GOPATH/src/github.com/nebulasio/go-nebulas
wenzildeiMac:go-nebulas wenzil$ make build
cd cmd/neb; go build -ldflags "-X main.version=1.0.1 -X main.commit=a7aaa0aaef2677038562356e1bbdccfb848d2fb4 -X main.branch=master -X main.compileAt=`date +%s`" -o ../../neb-a7aaa0aaef2677038562356e1bbdccfb848d2fb4
cd cmd/crashreporter; go build -ldflags "-X main.version=1.0.1 -X main.commit=a7aaa0aaef2677038562356e1bbdccfb848d2fb4 -X main.branch=master -X main.compileAt=`date +%s`" -o ../../neb-crashreporter
rm -f neb
ln -s neb-a7aaa0aaef2677038562356e1bbdccfb848d2fb4 neb
```


#### 4、运行星云链

##### 创世区块

在启动一个新的星云链前，我们必须定义创世区块的配置文件。

1、配置创世区块

**（可跳过此步，只是查看一下配置文件）**
进入到下载的星云链的根目录下，找到"nebulasio/go-nebulas/conf/default"目录下的"genesis.conf"文件，可以按照如下配置内容来修改。

```
meta {
  # 每条链的唯一标识
  # 每个区块和交易只会属于一条唯一的链，保证安全性
  chain_id: 100
}

consensus {
  # 在贡献度证明(PoD)被充分验证前，星云链采用DPoS共识算法
  # DPoS共识中，21个人组成一个朝代
  # 每隔一段时间都会切换朝代，每个朝代内，21个矿工轮流出块
  # 由于DPoS只是过渡方案，所以暂时不开放给公众挖矿，即当前版本朝代不会发生变更
  dpos {
    # 初始朝代，包含21个初始矿工地址
    dynasty: [
      [ miner address ],
      ...
    ]
  }
}

# 预分配的代币
token_distribution [
  {
    address: [ allocation address ]
    value: [ amount of allocation tokens ]
  },
  ...
]
```

终端操作如下：

```
###打开"genesis.conf"文件所在目录
wenzildeiMac:go-nebulas wenzil$ open $GOPATH/src/github.com/nebulasio/go-nebulas/conf/default
```

![4.终端执行open命令打开对应文件夹](media/15281819133101/4.%E7%BB%88%E7%AB%AF%E6%89%A7%E8%A1%8Copen%E5%91%BD%E4%BB%A4%E6%89%93%E5%BC%80%E5%AF%B9%E5%BA%94%E6%96%87%E4%BB%B6%E5%A4%B9.png)



##### 配置文件

在启动一个星云节点前，需要定义好该节点的配置文件。

进入到下载的星云链的根目录下，找到"nebulasio/go-nebulas/conf/default"目录下的"config.conf"文件，可以按照如下配置内容来修改。

```
# 网络配置
network {
  # 对于全网第一个节点，不需要配置seed
  # 否则，其他节点启动时需要配置seed，seed节点将会把网络中其他节点的路由信息同步给刚启动的节点
  # 可以配置多个seed, ["...", "..."]
  seed: ["/ip4/127.0.0.1/tcp/8680/ipfs/QmP7HDFcYmJL12Ez4ZNVCKjKedfE7f48f1LAkUc3Whz4jP"]

  # 节点监听网络消息端口，可以配置多个
  listen: ["0.0.0.0:8680"]

  # 网络私钥，用于确认身份节点
  # private_key: "conf/network/id_ed25519"
}

# 链配置
chain {
  # 链的唯一标识
  chain_id: 100

  # 数据存储地址
  datadir: "data.db"

  # 账户keystore文件存储地址
  keydir: "keydir"

  # 创世区块配置
  genesis: "conf/default/genesis.conf"

  # 签名算法，请勿修改
  signature_ciphers: ["ECC_SECP256K1"]

  # 矿工地址，矿工的keystore文件需要放置在配置的keydir下
  miner: "n1XkoVVjswb5Gek3rRufqjKNpwrDdsnQ7Hq"

  # Coinbase地址，该地址用于接收矿工的挖矿奖励，可以和矿工地址一致
  # 该地址的keystore无需暴露，不用放置在配置的keydir下
  coinbase: "n1FF1nz6tarkDVwWQkMnnwFPuPKUaQTdptE"

  # 矿工地址的密码
  passphrase: "passphrase"
}

# API配置
rpc {
    # GRPC服务端口
    rpc_listen: ["127.0.0.1:8684"]

    # HTTP服务端口
    http_listen: ["127.0.0.1:8685"]

    # 开放的API模块
    # API模块包含所有和用户私钥无关的接口
    # Admin模块包含所有和用户私钥相关的接口，需要慎重考虑该模块的访问权限
    http_module: ["api", "admin"]
}

# 日志配置
app {
    # 日志级别: 支持[debug, info, warn, error, fatal]
    log_level: "info"

    # 日志存放位置
    log_file: "logs"

    # 是否打开crash report服务
    enable_crash_report: false
}

# 监控服务配置
stats {
    # 是否打开监控服务
    enable_metrics: false

    # 监控服务将数据上传到Influxdb
    # 配置Influxdb的访问信息
    influxdb: {
        host: "http://localhost:8086"
        db: "nebulas"
        user: "admin"
        password: "admin"
    }
}
```

#### 5、启动星云链

> 此时启动的星云链是本地的私有链，和官方的测试网和主网没有任何相互关联

启动你的第一个星云节点，命令如下

```
cd $GOPATH/src/github.com/nebulasio/go-nebulas
./neb -c conf/default/config.conf
```

启动成功的话，将会看到如下信息，有Started Neblet的日志输出。

![5.启动星云链节点](media/15281819133101/5.%E5%90%AF%E5%8A%A8%E6%98%9F%E4%BA%91%E9%93%BE%E8%8A%82%E7%82%B9.png)

默认情况下，使用配置文件conf/default/config.conf启动的节点不是矿工节点。

接下来，启动你的第一个矿工节点，它的seed节点即我们刚刚启动的第一个节点。

```
cd $GOPATH/src/github.com/nebulasio/go-nebulas
./neb -c conf/example/miner.conf
```

在这个节点启动后，你会先看到如下信息，表示当前节点正在找种子节点同步。

![6.当前节点开始寻找种子节点同步](media/15281819133101/6.%E5%BD%93%E5%89%8D%E8%8A%82%E7%82%B9%E5%BC%80%E5%A7%8B%E5%AF%BB%E6%89%BE%E7%A7%8D%E5%AD%90%E8%8A%82%E7%82%B9%E5%90%8C%E6%AD%A5.png)


等待一会儿，将会看到如下信息，表示当前节点已经连上了seed节点完成了同步。

![7.当前节点已经连上了seed节点完成同步](media/15281819133101/7.%E5%BD%93%E5%89%8D%E8%8A%82%E7%82%B9%E5%B7%B2%E7%BB%8F%E8%BF%9E%E4%B8%8A%E4%BA%86seed%E8%8A%82%E7%82%B9%E5%AE%8C%E6%88%90%E5%90%8C%E6%AD%A5.png)


再等待几分钟，你会看到如下信息，表示当前矿工节点挖出了第一个区块。

![8.当前矿工节点挖出了第一个区块](media/15281819133101/8.%E5%BD%93%E5%89%8D%E7%9F%BF%E5%B7%A5%E8%8A%82%E7%82%B9%E6%8C%96%E5%87%BA%E4%BA%86%E7%AC%AC%E4%B8%80%E4%B8%AA%E5%8C%BA%E5%9D%97.png)

每当矿工挖个区块时，星云链节点会将矿工挖到的区块添加到区块链中
![9.星云链节点添加区块链中](media/15281819133101/9.%E6%98%9F%E4%BA%91%E9%93%BE%E8%8A%82%E7%82%B9%E6%B7%BB%E5%8A%A0%E5%8C%BA%E5%9D%97%E9%93%BE%E4%B8%AD.png)

> 提示: 目前的DPoS共识算法，会有21个节点轮流出块。由于我们只启动了21个矿工节点中的一个矿工节点，所以每隔15*21s才出一个块。你可以启动更多的矿工节点，填补的空缺。但是需要注意，多个节点间的端口号不要相互冲突了。


本文章部分内容来源于如下网址，对部分内容作了删改，全部重新运行并截图，如有侵权联系我删除文章。

参考网址：
http://www.qukuaiwang.com.cn/news/1738.html
https://blog.csdn.net/lwbcxc/article/details/79681902
https://github.com/nebulasio/wiki



