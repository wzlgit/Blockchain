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

4、为智能合约创建新账号
智能合约附着在账号上的，因此需要为合约准备一个账号。

于是，先创建key

```
wenzildeiMac:study wenzil$ cleos create key
Private key: 5KcUM9PkRLqRq6KtrGT4ek8dy9mLfPh4cTGR6Pt8fU5tH6Q7y4c
Public key: EOS7caVJ2iDpM7edncLbTGtQqXtTZqyMYxhUP8e1xPzv3mVyL7Gxu
```

然后，导入到钱包

```
wenzildeiMac:study wenzil$ cleos wallet import 5KcUM9PkRLqRq6KtrGT4ek8dy9mLfPh4cTGR6Pt8fU5tH6Q7y4c
imported private key for: EOS7caVJ2iDpM7edncLbTGtQqXtTZqyMYxhUP8e1xPzv3mVyL7Gxu
```

创建账号

```
wenzildeiMac:study wenzil$ cleos create account eosio hello.eos EOS7caVJ2iDpM7edncLbTGtQqXtTZqyMYxhUP8e1xPzv3mVyL7Gxu EOS7caVJ2iDpM7edncLbTGtQqXtTZqyMYxhUP8e1xPzv3mVyL7Gxu
executed transaction: f9f0317694b27749c7eed78a27babe0ea0d92987213f2f6f561c3237ce8e100b  200 bytes  265 us
#         eosio <= eosio::newaccount            {"creator":"eosio","name":"hello.eos","owner":{"threshold":1,"keys":[{"key":"EOS7caVJ2iDpM7edncLbTGt...
warning: transaction executed locally, but may not be confirmed by the network yet
```

5、部署智能合约

部署智能合约需要将合约绑定到账号，先返回到上一个目录

```
wenzildeiMac:hello_eos wenzil$ pwd
/Users/wenzil/Desktop/study/hello_eos
wenzildeiMac:hello_eos wenzil$ cd ..
wenzildeiMac:study wenzil$ pwd
/Users/wenzil/Desktop/study
wenzildeiMac:study wenzil$ cleos set contract hello.eos ./hello_eos -p hello.eos
Reading WAST/WASM from ./hello_eos/hello_eos.wasm...
Using already assembled WASM...
Publishing contract...
executed transaction: 2820024b8513f5c05c82a12444765aaf70f2c9fa717c68993a93a7739f108a8c  1808 bytes  439 us
#         eosio <= eosio::setcode               {"account":"hello.eos","vmtype":0,"vmversion":0,"code":"0061736d01000000013b0c60027f7e006000017e6002...
#         eosio <= eosio::setabi                {"account":"hello.eos","abi":"0e656f73696f3a3a6162692f312e300001047465737400010475736572046e616d6501...
warning: transaction executed locally, but may not be confirmed by the network yet
```


6、调用智能合约

```
wenzildeiMac:study wenzil$ cleos push action hello.eos test '{"user":"hello.eos"}' -p hello.eos 
executed transaction: 23ba97ec481f4b14e533c164fe044aa4d539b413f29d6d608644f862a5414c7c  104 bytes  242 us
#     hello.eos <= hello.eos::test              {"user":"hello.eos"}
warning: transaction executed locally, but may not be confirmed by the network yet
```

奇怪了，居然没打印出test方法中print的内容。





