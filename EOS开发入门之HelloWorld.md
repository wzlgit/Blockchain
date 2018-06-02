### EOS开发入门之HelloWorld

EOS的智能合约是基于C++语言的，本篇将进入智能合约的开发阶段，写一个HelloWorld之类的合约。
话不多说，开始敲大🐴。

新建一个"hello_eos"的文件夹，然后创建一个"hello_eos.cpp"的文件，文件夹和文件名一样，内容如下：

1、编写智能合约

```
#include <eosiolib/eosio.hpp>
#include <eosiolib/print.hpp>

using namespace eosio;

class hello_eos : public eosio::contract {

  public:
      using contract::contract;

      /// @abi action
      void test(account_name user) {
         print("Hello ,", name{user});
      }
};
EOSIO_ABI(hello_eos, (test))

```

2、编译智能合约

```
wenzildeiMac:hello_eos wenzil$ pwd
/Users/wenzil/Desktop/study/hello_eos
wenzildeiMac:hello_eos wenzil$ eosiocpp -o hello_eos.wast hello_eos.cpp
-bash: eosiocpp: command not found
```

找不到智能合约编译工具eosiocpp，坑啊。

进入到EOS源码的"build"目录，然后执行如下命令生成eosiocpp工具。

```
wenzildeiMac:build wenzil$ pwd
/Users/wenzil/Documents/Study/eos/build
wenzildeiMac:build wenzil$ sudo make install
Password:
[  0%] Built target binaryen
[  0%] Built target wasm
##########此处省略多行输出内容##########
-- Installing: /usr/local/bin/eosio-abigen
-- Installing: /usr/local/bin/eosiocpp
```

如果失败了，再次编译即可，编译成功后发现多了两个文件。

```
wenzildeiMac:hello_eos wenzil$ ls
hello_eos.cpp	hello_eos.wasm	hello_eos.wast
```

3、生成abi文件

```
wenzildeiMac:hello_eos wenzil$ eosiocpp -g hello_eos.abi hello_eos.cpp 
3376922ms thread-0   abi_generator.hpp:68          ricardian_contracts  ] Warning, no ricardian clauses found for hello_eos

3376922ms thread-0   abi_generator.hpp:75          ricardian_contracts  ] Warning, no ricardian contract found for test

Generated hello_eos.abi ...
```

再次查看，发现多了一个abi文件。

```
wenzildeiMac:hello_eos wenzil$ ls
hello_eos.abi	hello_eos.cpp	hello_eos.wasm	hello_eos.wast
```

4、编译智能合约

user1是之前创建钱包时创建的账户名，这里没有再创建账户了。返回上一个目录，再进行编译。

```
wenzildeiMac:study wenzil$ cleos set contract user1 ./hello_eos
Reading WAST/WASM from ./hello_eos/hello_eos.wasm...
Using already assembled WASM...
Publishing contract...
executed transaction: ac3764e28b27f1d1cadc31be7d25aceba4bee591d4b284b789bab63e120e7016  1800 bytes  536 us
#         eosio <= eosio::setcode               {"account":"user1","vmtype":0,"vmversion":0,"code":"0061736d01000000013b0c60027f7e006000017e60027e7e...
#         eosio <= eosio::setabi                {"account":"user1","abi":"0e656f73696f3a3a6162692f312e300001047465737400010475736572046e616d65010000...
warning: transaction executed locally, but may not be confirmed by the network yet
```

5、调用智能合约

```
wenzildeiMac:build wenzil$ cleos push action user1 test '{"user":"wenzil"}' -p user1
executed transaction: 1d32ad7144dc84602bcbbd518d25ecad068c5364de367cb8d6f5947b6368f081  104 bytes  239 us
#         user1 <= user1::test                  {"user":"wenzil"}
warning: transaction executed locally, but may not be confirmed by the network yet
```

奇怪了，居然没打印出print函数的内容，有空再研究下为啥。





