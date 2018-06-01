### EOS简介与环境搭建

2018年，最火爆且值得期待的区块链项目莫过于EOS和IPFS。IPFS在另外一篇文章《IPFS开发简介与入门实践》有介绍，内容篇幅比较多，写作比较耗时。本篇内容比较少，只对简单介绍、安装要求以及环境搭建等方面作介绍。

EOS官网：https://eos.io/

#### EOS的简单介绍

* EOS是什么

    EOS：可以理解为Enterprise Operation System，即为商用分布式应用设计的一款区块链操作系统。EOS是引入的一种新的区块链架构，旨在实现分布式应用的性能扩展。注意，它并不是像比特币和以太坊那样的货币，而是基于EOS软件项目之上发布的代币，被称为区块链3.0。
**以上内容来自百度百科**


* EOS背后的公司及创始人

   EOS背后的区块链初创公司是Block.One，号称能达到每秒百万级的交易速度。
   
   EOS的创始人BM（Bytemaster，真名为Daniel Larimer），是目前世界上唯一一个连续成功开发了三个（Bitshares，Steem，EOS）基于区块链技术的去中心化系统的人，曾经和比特币的创始人中本聪、以太坊的创始人V神互怼过。

* EOS的火热

  在2017年，EOS ICO创造过ICO记录——5天内筹集了1.85亿美元。目前，EOS市值已达109.5亿美元，成为第五大数字货币。
  ![1.EOS排行](media/15278229023722/1.EOS%E6%8E%92%E8%A1%8C.png)
  
  2018年4月，离主网上线还有2个月时间，堪称“总统之争”的EOS全球21个超级节点之争如火如荼地进行着。随着社区运营为主的引力区、佳能区外，一些明星级别的项目及区块链大咖，包括EOS.CYBEX（暴走恭亲王）、EOSLaoMao（老猫）、InBlockchain（李笑来团队）、薛蛮子等等陆续宣布参与竞选，EOS的价格应声大涨。截至4月25日，全球共有74个参选者，其中中国独占32个，占总参选人数的43.2%。
  ![2.EOS超级节点之争](media/15278229023722/2.EOS%E8%B6%85%E7%BA%A7%E8%8A%82%E7%82%B9%E4%B9%8B%E4%BA%89.png)

  
  相关内容拓展阅读：
  [一文看懂EOS超级节点](http://www.dayqkl.com/21787.html)
  
  
* EOS的优势

    1、采用了DPos（股份权益证明机制）共识算法，通过并行链的方式，最高可以达到数百万TPS。而比特币每秒7笔左右，以太坊每秒15笔左右；
    2、使用免费，而以太坊交易需要花费Gas；
    3、便利地升级和Bug恢复。
    
* EOS漏洞事件
  
    然而，在5月29日中午，360安全卫士在其官方微博上发布长文《360发现区块链史诗级漏洞 可完全控制虚拟货币交易》。受到此漏洞曝光的影响，EOS 应声大幅下跌。BM回应EOS漏洞：360报告制造恐慌，披露前已修复。
 
    相关事件拓展阅读：
 [【重磅】360发现EOS史诗级漏洞，可完全控制虚拟货币交易！](https://xueqiu.com/7686580358/107990177)
    
* EOS相关连接

    [EOS的官方Github](https://github.com/EOSIO/eos)
    [EOS开发文档](https://eosio.github.io/eos/)
    [EOS官方白皮书](https://github.com/EOSIO/Documentation/blob/master/TechnicalWhitePaper.md)
    [中文翻译版白皮书](https://github.com/EOSIO/Documentation/blob/master/zh-CN/TechnicalWhitePaper.md)
    [EOS ÐApps网址](https://eosprojects.org/)
    [EOS安装说明文档](https://github.com/EOSIO/eos/wiki/Local-Environment)
  

#### 2、EOS的安装要求

既然EOS这么火爆，有必要了解一下EOS的开发，先从EOS的安装开始，EOS安装有对应的系统和硬件配置要求。

安装之前，说明下EOS有关的硬件要求。
1、目前支持的操作系统（不支持Windows系统）：
> 
* Amazon 2017.09 and higher.
* Centos 7.
* Fedora 25 and higher (Fedora 27 recommended).
* Mint 18.
* Ubuntu 16.04 (Ubuntu 16.10 recommended).
* MacOS Darwin 10.12 and higher (MacOS 10.13.x recommended).

2、硬件配置要求：
需要8GB内存和20GB的硬盘空间。


#### 3、EOS的环境搭建

如果你的电脑能够满足上面的安装要求，那么可以进行以下的操作，进行EOS的搭建环境。

1、用Git克隆，下载EOS的源码

```
git clone https://github.com/eosio/eos --recursive
```
后面添加--recursive参数会自动全部子模块，从而下载完整的源码。如果没有添加--recursive参数的华，可以进入到项目的根目录执行如下命令继续下载补全子模块。

```
git submodule update --init --recursive
```

下载的时间会比较漫长，请耐心等待。

```
wenzildeiMac:Study wenzil$ pwd
/Users/wenzil/Documents/Study
wenzildeiMac:Study wenzil$ git clone https://github.com/eosio/eos --recursive
Cloning into 'eos'...
remote: Counting objects: 105092, done.
remote: Compressing objects: 100% (51/51), done.
remote: Total 105092 (delta 36), reused 47 (delta 25), pack-reused 105011
Receiving objects: 100% (105092/105092), 90.82 MiB | 80.00 KiB/s, done.
Resolving deltas: 100% (83888/83888), done.
Checking out files: 100% (13238/13238), done.
Submodule 'contracts/libc++/upstream' (https://github.com/EOSIO/libcxx.git) registered for path 'contracts/libc++/upstream'
##########此处省略多行输出内容##########Cloning into '/Users/wenzil/Documents/Study/eos/contracts/libc++/upstream'...
remote: Counting objects: 73628, done.        
remote: Total 73628 (delta 0), reused 0 (delta 0), pack-reused 73628        
Receiving objects: 100% (73628/73628), 22.73 MiB | 385.00 KiB/s, done.
Resolving deltas: 100% (51708/51708), done.
##########此处省略多行输出内容##########
Submodule path 'externals/binaryen/test/emscripten': checked out '1d979a75722a513b586e6705d662ff4ee0ea832b'
Submodule path 'externals/binaryen/test/spec': checked out '668281b3d7dfe9be6cbc3c4500b537c49d6449b9'
Submodule path 'externals/binaryen/test/waterfall': checked out '900646fc880d526160b0df9b78bc9dd4f02dc1ab'
Submodule path 'externals/magic_get': checked out '89fda1da702e6c76a22bfb6233e9e3d0641708ec'
Submodule path 'libraries/appbase': checked out '44d7d88e561cc9708bcac9453686c8d0f0ffa41e'
Submodule path 'libraries/chainbase': checked out 'b40a1ec99ee0eb16e28ee1c1a9577f1af7d48b84'
Submodule path 'libraries/softfloat': checked out '9dff375a6e3a9728a024b30afa127c88fcb40ea2'
wenzildeiMac:Study wenzil$ 
```

2、编译EOS的源码

进入到EOS源码的根目录，然后执行如下命令：

```
./eosio_build.sh darwin full
```

```
wenzildeiMac:eos wenzil$ pwd
/Users/wenzil/Documents/Study/eos

wenzildeiMac:eos wenzil$ ./eosio_build.sh darwin full

	Beginning build version: 1.2
	2018年 6月 1日 星期五 12时34分56秒 UTC
	User: wenzil
	git head id: cbf28a237f3f2605b4ac34c5bb0c98a4d931e99f
	Current branch: master

	ARCHITECTURE: Darwin

	OS name: Darwin
	OS Version: 10.13.2
	CPU speed: 270.00Ghz
	CPU cores: 4
	Physical Memory: 16 Gbytes
	Disk install: /dev/disk0s2
	Disk space total: 930G
	Disk space available: 709G

	Checking xcode-select installation
	xcode-select installation found @ 
	/usr/bin/xcode-select 

	Checking Ruby installation.
	Ruby installation found @ 
	/usr/bin/ruby 

	Checking Home Brew installation
	Home Brew installation found @
	/usr/local/bin/brew

	Checking dependencies.
	Checking automake ... 		 automake found
	Checking Libtool ... 		 Libtool found
	Checking OpenSSL ... 		 OpenSSL found
	Checking llvm ... 		 llvm NOT found.
	Checking wget ... 		 wget NOT found.
	Checking CMake ... 		 CMake found
	Checking GMP ... 		 GMP NOT found.
	Checking gettext ... 		 gettext NOT found.
	Checking MongoDB ... 		 MongoDB found
	Checking Doxygen ... 		 Doxygen NOT found.
	Checking Graphviz ... 		 Graphviz NOT found.
	Checking LCOV ... 		 LCOV NOT found.
	Checking Python3 ... 		 Python3 found

	The following dependencies are required to install EOSIO.

	1. llvm\n\t2. wget\n\t3. GMP\n\t4. gettext\n\t5. Doxygen\n\t6. Graphviz\n\t7. LCOV\n\t

Do you wish to install these packages?
1) Yes
2) No
#? 
```

3、继续编译和安装

弹出安装询问，选择"Yes"。这时不是输入"Yes"，而是是输入"1"，然后输入当前用户密码即可继续安装和编译。
编译需要很长的时间，中途需要多次输入用户密码，花了一个多小时才编译成功。
中间失败过两次，然后可以重新执行"./eosio_build.sh darwin full"命令。

```
Do you wish to install these packages?
1) Yes
2) No
#? Yes
Please type 1 for yes or 2 for no.
#? 1
Password:
	Updating Home Brew.
Updated 2 taps (dart-lang/dart, homebrew/core).
==> New Formulae
dart-lang/dart/dart@1
==> Updated Formulae
##########此处省略多行输出内容##########
==> Installing dependencies for llvm@4: libffi
==> Installing llvm@4 dependency: libffi
##########此处省略多行输出内容##########

Install the project...
-- Install configuration: "Release"
-- Up-to-date: /usr/local/include/bsoncxx/v_noabi/bsoncxx
-- Installing: /usr/local/include/bsoncxx/v_noabi/bsoncxx/array
-- Installing: /usr/local/include/bsoncxx/v_noabi/bsoncxx/array/element.hpp
##########此处省略多行输出内容##########
Scanning dependencies of target LLVMDemangle
Scanning dependencies of target LLVMTableGen
Scanning dependencies of target obj.llvm-tblgen
Scanning dependencies of target LLVMSupport
##########此处省略多行输出内容##########
[100%] Building CXX object lib/Demangle/CMakeFiles/LLVMDemangle.dir/ItaniumDemangle.cpp.o
[100%] Building CXX object lib/TableGen/CMakeFiles/LLVMTableGen.dir/Error.cpp.o
[100%] Building CXX object unittests/CMakeFiles/unit_test.dir/wasm_tests.cpp.o
[100%] Building CXX object unittests/CMakeFiles/unit_test.dir/whitelist_blacklist_tests.cpp.o
[100%] Linking CXX executable unit_test
[100%] Built target unit_test


	 _______  _______  _______ _________ _______
	(  ____ \(  ___  )(  ____ \\__   __/(  ___  )
	| (    \/| (   ) || (    \/   ) (   | (   ) |
	| (__    | |   | || (_____    | |   | |   | |
	|  __)   | |   | |(_____  )   | |   | |   | |
	| (      | |   | |      ) |   | |   | |   | |
	| (____/\| (___) |/\____) |___) (___| (___) |
	(_______/(_______)\_______)\_______/(_______)

	EOSIO has been successfully built. 00:42:15

	To verify your installation run the following commands:

	/usr/local/bin/mongod -f /usr/local/etc/mongod.conf &
	cd /Users/wenzil/Documents/Study/eos/build; make test

	For more information:
	EOSIO website: https://eos.io
	EOSIO Telegram channel @ https://t.me/EOSProject
	EOSIO resources: https://eos.io/resources/
	EOSIO wiki: https://github.com/EOSIO/eos/wiki
	
```


4、查看编译生成的应用程序

编译成功后，会生成多个命令行的应用程序，跟目录路径是"/eos/build/programs"，以下简单说明下对应3个应用程序的作用。
![3.编译生成的3个应用程序](media/15278229023722/3.%E7%BC%96%E8%AF%91%E7%94%9F%E6%88%90%E7%9A%843%E4%B8%AA%E5%BA%94%E7%94%A8%E7%A8%8B%E5%BA%8F.png)


1) keosd：钱包管理组件，可以创建私钥。
2) nodeos：服务器的区块链节点组件，启动Nodeos后，会自动产生区块。
3) cleos：可以操作nodeos和keosd的命令行工具。

5、启动节点

进入到eos的"build/programs/nodeos"目录，然后执行"./nodeos"命令。

```
wenzildeiMac:eos wenzil$ cd build/programs/nodeos
wenzildeiMac:nodeos wenzil$ pwd
/Users/wenzil/Documents/Study/eos/build/programs/nodeos
wenzildeiMac:nodeos wenzil$ ./nodeos
2367002ms thread-0   chain_plugin.cpp:200          plugin_initialize    ] initializing chain plugin
2367003ms thread-0   chain_plugin.cpp:354          plugin_initialize    ] Starting up fresh blockchain with default genesis state.
2367068ms thread-0   http_plugin.cpp:285           plugin_initialize    ] configured http to listen on 127.0.0.1:8888
2367068ms thread-0   net_plugin.cpp:2815           plugin_initialize    ] Initialize net plugin
2367068ms thread-0   net_plugin.cpp:2836           plugin_initialize    ] host: 0.0.0.0 port: 9876 
2367069ms thread-0   net_plugin.cpp:2908           plugin_initialize    ] my node_id is 7b17f1423ac31cfbd02fc43bf105bf5ecf42dc6b342865c8406c8c4ca24da93c
2367069ms thread-0   main.cpp:104                  main                 ] nodeos version cbf28a23
2367069ms thread-0   main.cpp:105                  main                 ] eosio root is /Users/wenzil/Library/Application Support
2367069ms thread-0   controller.cpp:1141           startup              ] No head block in fork db, perhaps we need to replay
2367069ms thread-0   controller.cpp:307            initialize_fork_db   ]  Initializing new blockchain with genesis state                  
2367103ms thread-0   chain_plugin.cpp:409          plugin_startup       ] starting chain in read/write mode
2367103ms thread-0   chain_plugin.cpp:414          plugin_startup       ] Blockchain started; head block is #1, genesis timestamp is 2018-03-02T12:00:00.000
2367103ms thread-0   http_plugin.cpp:323           plugin_startup       ] start listening for http requests
2367104ms thread-0   net_plugin.cpp:2920           plugin_startup       ] starting listener, max clients is 25
2367104ms thread-0   producer_plugin.cpp:577       plugin_startup       ] producer plugin:  plugin_startup() begin
2367104ms thread-0   producer_plugin.cpp:604       plugin_startup       ] producer plugin:  plugin_startup() end
```


6、修改节点配置文件config.ini

路径格式为：
Mac OS: ~/Library/Application Support/eosio/nodeos/config
Linux: ~/.local/share/eosio/nodeos/config
![4.修改节点配置文件](media/15278229023722/4.%E4%BF%AE%E6%94%B9%E8%8A%82%E7%82%B9%E9%85%8D%E7%BD%AE%E6%96%87%E4%BB%B6.png)

参考如下格式对config.ini的内容作修改

原文件的内容：

```
# Enable block production, even if the chain is stale. (eosio::producer_plugin)
enable-stale-production = false
# ID of producer controlled by this node (e.g. inita; may specify multiple times) (eosio::producer_plugin)
# producer-name = 
# Plugin(s) to enable, may be specified multiple times
# plugin = 
```

修改后的内容：

```
genesis-json = "/Users/wenzil/Library/Application\ Support/eosio/nodeos/config/genesis.json"
###格式为-> genesis-json = /path/to/eos/source/genesis.json
enable-stale-production = true
producer-name = eosio
plugin = eosio::producer_plugin
plugin = eosio::wallet_api_plugin
plugin = eosio::chain_api_plugin
plugin = eosio::http_plugin
```

7、"Control+C"快捷键终止第5步的运行操作，然后重新执行"./nodeos"命令

```
wenzildeiMac:nodeos wenzil$ ./nodeos
156065ms thread-0   chain_plugin.cpp:200          plugin_initialize    ] initializing chain plugin
156067ms thread-0   block_log.cpp:123             open                 ] Log is nonempty
156067ms thread-0   block_log.cpp:136             open                 ] my->head->block_num(): 1 
156067ms thread-0   block_log.cpp:142             open                 ] Index is nonempty
156087ms thread-0   wallet_plugin.cpp:41          plugin_initialize    ] initializing wallet plugin
156088ms thread-0   http_plugin.cpp:285           plugin_initialize    ] configured http to listen on 127.0.0.1:8888
156088ms thread-0   net_plugin.cpp:2815           plugin_initialize    ] Initialize net plugin
156088ms thread-0   net_plugin.cpp:2836           plugin_initialize    ] host: 0.0.0.0 port: 9876 
156088ms thread-0   net_plugin.cpp:2908           plugin_initialize    ] my node_id is 67c89ef6761341ccb3a76436f7cd30fdf46f7c254e0c52b1570a21ea86761231
156088ms thread-0   main.cpp:104                  main                 ] nodeos version cbf28a23
156088ms thread-0   main.cpp:105                  main                 ] eosio root is /Users/wenzil/Library/Application Support
156089ms thread-0   chain_plugin.cpp:409          plugin_startup       ] starting chain in read/write mode
156089ms thread-0   chain_plugin.cpp:414          plugin_startup       ] Blockchain started; head block is #1, genesis timestamp is 2018-03-02T12:00:00.000
156089ms thread-0   producer_plugin.cpp:577       plugin_startup       ] producer plugin:  plugin_startup() begin
156089ms thread-0   producer_plugin.cpp:592       plugin_startup       ] Launching block production for 1 producers at 2018-06-01T11:02:36.089.
156089ms thread-0   producer_plugin.cpp:604       plugin_startup       ] producer plugin:  plugin_startup() end
156089ms thread-0   http_plugin.cpp:323           plugin_startup       ] start listening for http requests
156089ms thread-0   wallet_api_plugin.cpp:68      plugin_startup       ] starting wallet_api_plugin
156089ms thread-0   http_plugin.cpp:369           add_handler          ] add api url: /v1/wallet/create
156090ms thread-0   http_plugin.cpp:369           add_handler          ] add api url: /v1/wallet/create_key
156090ms thread-0   http_plugin.cpp:369           add_handler          ] add api url: /v1/wallet/get_public_keys
156090ms thread-0   http_plugin.cpp:369           add_handler          ] add api url: /v1/wallet/import_key
156090ms thread-0   http_plugin.cpp:369           add_handler          ] add api url: /v1/wallet/list_keys
##########此处省略多行输出内容##########

156521ms thread-0   producer_plugin.cpp:1073      produce_block        ] Produced block 00000002b36978e9... #2 @ 2018-06-01T11:02:36.500 signed by eosio [trxs: 0, lib: 0, confirmed: 0]
157004ms thread-0   producer_plugin.cpp:1073      produce_block        ] Produced block 00000003bd212101... #3 @ 2018-06-01T11:02:37.000 signed by eosio [trxs: 0, lib: 2, confirmed: 0]
157505ms thread-0   producer_plugin.cpp:1073      produce_block        ] Produced block 00000004f594ba4a... #4 @ 2018-06-01T11:02:37.500 signed by eosio [trxs: 0, lib: 3, confirmed: 0]
158000ms thread-0   producer_plugin.cpp:1073      produce_block        ] Produced block 00000005ea6790c3... #5 @ 2018-06-01T11:02:38.000 signed by eosio [trxs: 0, lib: 4, confirmed: 0]
158501ms thread-0   producer_plugin.cpp:1073      produce_block        ] Produced block 00000006149f7c7c... #6 @ 2018-06-01T11:02:38.500 signed by eosio [trxs: 0, lib: 5, confirmed: 0]
159002ms thread-0   producer_plugin.cpp:1073      produce_block        ] Produced block 000000076a9040c5... #7 @ 2018-06-01T11:02:39.000 signed by eosio [trxs: 0, lib: 6, confirmed: 0]
159505ms thread-0   producer_plugin.cpp:1073      produce_block        ] Produced block 00000008de9dd847... #8 @ 2018-06-01T11:02:39.500 signed by eosio [trxs: 0, lib: 7, confirmed: 0]
160003ms thread-0   producer_plugin.cpp:1073      produce_block        ] Produced block 0000000911e7b0d6... #9 @ 2018-06-01T11:02:40.000 signed by eosio [trxs: 0, lib: 8, confirmed: 0]
160504ms thread-0   producer_plugin.cpp:1073      produce_block        ] Produced block 0000000a454378c3... #10 @ 2018-06-01T11:02:40.500 signed by eosio [trxs: 0, lib: 9, confirmed: 0]
##########此处省略多行输出内容，后面一直在产生区块##########
```


8、查看区块信息

 新建一个终端，切换到eos目录的"/build/programs/cleos"目录下，然后使用cleos工具查看区块信息，操作如下：


```
wenzildeiMac:cleos wenzil$ pwd
/Users/wenzil/Documents/Study/eos/build/programs/cleos

wenzildeiMac:cleos wenzil$ ./cleos --url http://localhost:8888 get info
{
  "server_version": "cbf28a23",
  "chain_id": "a4fe43c897279a677025624555f3835b949f995f87923c97f0392dcff835bfd0",
  "head_block_num": 1085,
  "last_irreversible_block_num": 1084,
  "last_irreversible_block_id": "0000043c13451164b49b3a6067df9003b60b334ce745a05219ca77ffc46c3db7",
  "head_block_id": "0000043d0a519b2a693947d031f145df12b2a059c482d9f41a8c1a2df4895dbb",
  "head_block_time": "2018-06-01T11:11:38",
  "head_block_producer": "eosio",
  "virtual_block_cpu_limit": 590632,
  "virtual_block_net_limit": 3100815,
  "block_cpu_limit": 199900,
  "block_net_limit": 1048576
}
wenzildeiMac:cleos wenzil$ ./cleos --url http://localhost:8888 get info
{
  "server_version": "cbf28a23",
  "chain_id": "a4fe43c897279a677025624555f3835b949f995f87923c97f0392dcff835bfd0",
  "head_block_num": 1115,
  "last_irreversible_block_num": 1114,
  "last_irreversible_block_id": "0000045a8e727712103d0b9d95d07c518ec0fcacad8f23acdf76f3281de0ea6c",
  "head_block_id": "0000045b2bbd8ba3c08479338e3910c19d28a56b9359ce1a888df799300501b2",
  "head_block_time": "2018-06-01T11:11:53",
  "head_block_producer": "eosio",
  "virtual_block_cpu_limit": 608612,
  "virtual_block_net_limit": 3195281,
  "block_cpu_limit": 199900,
  "block_net_limit": 1048576
}
```

