## 星云链Nebulas——6.开发和部署星云链DApp

经过前面一系列的基础学习，本篇将进入DApp开发阶段。

### 1、创建星云链NAS钱包

在星云链上，部署智能合约需要星云币，所以需要申请一个星云链的钱包。

#### 1.1 下载官方web钱包

注意：出于安全性的考虑，星云不架设专门Web钱包网站，以免私钥触网导致信息被盗。**将钱包下载到自己的计算机上，断开网络，离线执行这些操作更为安全。**

GitHub地址为：

```bash
https://github.com/nebulasio/web-wallet
```

```bash
wenzildeiMac:study wenzil$ git clone https://github.com/nebulasio/web-wallet.git
Cloning into 'web-wallet'...
remote: Counting objects: 1592, done.
remote: Compressing objects: 100% (43/43), done.
remote: Total 1592 (delta 34), reused 48 (delta 21), pack-reused 1528
Receiving objects: 100% (1592/1592), 1.93 MiB | 156.00 KiB/s, done.
Resolving deltas: 100% (1119/1119), done.
Checking out files: 100% (54/54), done.
```

![1.官方web钱包](media/15286049569309/1.%E5%AE%98%E6%96%B9web%E9%92%B1%E5%8C%85.png)


双击 index.html打开 ，这时会启动浏览器，打开一个保存在本地的网页，即Web版钱包。

![2.新建钱包页面](media/15286049569309/2.%E6%96%B0%E5%BB%BA%E9%92%B1%E5%8C%85%E9%A1%B5%E9%9D%A2.png)

设置密码(密码至少是9位数) ，点击创建新钱包。

特别提醒: 请务必保证秘钥文件的安全，不要忘记密码**（私钥是您数字资产的凭证，请不要告诉任何人，如不慎暴露或丢失，会面临丢失对应数字资产的风险，请谨慎保管）**。


右上方可以选择主网、测试网络还是私有网络，为了方便演示，这里选择了测试网络。

![3.新建钱包下载密码库文件](media/15286049569309/3.%E6%96%B0%E5%BB%BA%E9%92%B1%E5%8C%85%E4%B8%8B%E8%BD%BD%E5%AF%86%E7%A0%81%E5%BA%93%E6%96%87%E4%BB%B6.png)

输入密码后，点击"新建钱包"按钮，然后点击"下载密码库文件"，会生成一个"n1"开头的文件，保存到电脑上，这里包含了这个钱包内的私钥信息，因此请妥善保管，避免遗失或被盗。

![4.密码库文件](media/15286049569309/4.%E5%AF%86%E7%A0%81%E5%BA%93%E6%96%87%E4%BB%B6.png)


注意：密码库文件可以用于断开网络，进行离线交易，这个钱包内的秘钥信息都包含在内，**请妥善保管，如不慎丢失，会面临丢失对应数字资产的风险**。

#### 1.2 使用web钱包

点击"钱包信息"，然后上传钱包文件。

![5.切换到钱包信息](media/15286049569309/5.%E5%88%87%E6%8D%A2%E5%88%B0%E9%92%B1%E5%8C%85%E4%BF%A1%E6%81%AF.png)

上传成功后，输入创建钱包时的密码，点击"解锁"。

![6.解锁钱包](media/15286049569309/6.%E8%A7%A3%E9%94%81%E9%92%B1%E5%8C%85.png)

解锁成功后，会在当前页面出现类似如下页面，可以查看账户的地址、余额、私钥、二维码地址等信息。
![7.解锁后的钱包信息](media/15286049569309/7.%E8%A7%A3%E9%94%81%E5%90%8E%E7%9A%84%E9%92%B1%E5%8C%85%E4%BF%A1%E6%81%AF.png)

发现钱包余额为0，以下介绍如何获取测试网络的星云币。2018年5月22日，星云链推出了星云激励计划，在主网每提交一个有效DApp即可获得100NAS（注：提交的DApp需要官方审核，需要满足一定的条件，可以查看以下链接的内容）。当然，也可以用钱来购买主网的星云币，这里不作介绍，有兴趣的同学自行研究。

[星云激励计划DApp审核规则指南](https://blog.nebulas.io/2018/05/22/nip-dapp-rule/)

#### 1.3 领取测试网络星云币

领取测试网络星云币网址：
https://testnet.nebulas.io/claim/

每天可以领取1NAS的测试币，领取的星云币只能在测试网使用，没有实际价值，与主网没有任何关系。 

打开网址后，输入邮箱和钱包地址，然后"点击按钮进行验证"。
![8.输入邮箱和钱包地址验证](media/15286049569309/8.%E8%BE%93%E5%85%A5%E9%82%AE%E7%AE%B1%E5%92%8C%E9%92%B1%E5%8C%85%E5%9C%B0%E5%9D%80%E9%AA%8C%E8%AF%81.png)

然后，在"Check Balance"一栏输入钱包地址，再点击“Check Balance”按钮检查钱包余额。

![9.检查钱包余额](media/15286049569309/9.%E6%A3%80%E6%9F%A5%E9%92%B1%E5%8C%85%E4%BD%99%E9%A2%9D.png)

成功的话，页面会弹出这样的提示。
![10.钱包多了一个星云币](media/15286049569309/10.%E9%92%B1%E5%8C%85%E5%A4%9A%E4%BA%86%E4%B8%80%E4%B8%AA%E6%98%9F%E4%BA%91%E5%B8%81.png)

说明已经成功领取了一个星云币NAS，接下来就可以用来部署和测试自己编写的智能合约了。

### 2、编写和部署智能合约

这里就不介绍编写和部署智能合约的内容了，直接进入编写部分。

这里以一个简单的"文章发布"DApp为例来讲解，主要的功能是任何区块链上的用户都可以发表一篇标题不重复的文章，如果文章标题被占用，就只能查看这篇文章。

#### 2.1 编写智能合约

先创建一个该DApp的工程目录，暂且命名为"nebulas_article_publish"，并创建一个名为"article_publish.js"的智能合约文件，终端操作如下：

```bash
wenzildeiMac:~ wenzil$ cd /Users/wenzil/Desktop/study 
wenzildeiMac:study wenzil$ pwd
/Users/wenzil/Desktop/study
wenzildeiMac:study wenzil$ mkdir nebulas_article_publish
wenzildeiMac:study wenzil$ cd nebulas_article_publish/
wenzildeiMac:nebulas_article_publish wenzil$ touch article_publish.js
```

然后用JavaScript编写一个智能合约的代码，整个代码如下：

```js
'use strict'

var Article = function(text){
    if(text){
        var obj = JSON.parse(text);
        // 文章作者
        this.author = obj.author;
        // 文章标题
        this.title = obj.title;
        // 文章内容
        this.content = obj.content;
        // 发布者钱包地址
        this.address = obj.address;
        // 文章发布时间
        this.date = obj.date;
    }
};

Article.prototype = {
    toString : function(){
        return JSON.stringify(this)
    }
};

// 获取当前时间
function getCurrentTime() {
    var date = new Date();
    var seperator1 = "-";
    var seperator2 = ":";
    var month = date.getMonth() + 1;
    var strDate = date.getDate();
    if (month >= 1 && month <= 9) {
        month = "0" + month;
    }
    if (strDate >= 0 && strDate <= 9) {
        strDate = "0" + strDate;
    }
    var currentTime = date.getFullYear() + seperator1 + month + seperator1 + strDate
            + " " + date.getHours() + seperator2 + date.getMinutes()
            + seperator2 + date.getSeconds();
    return currentTime;
};

// 智能合约存储区
var ArticleContract = function () {
    // 使用"LocalContractStorage"绑定Map属性
    LocalContractStorage.defineMapProperty(this, "articleData", {
        stringify: function (obj) {
            return obj.toString();
        },
        parse: function (text) {
            return new Article(text);
        }
    });
};

ArticleContract.prototype ={
    init: function(){
      //TODO:
    },

    /**
      * 根据文章标题或者对应的文章
      * author：文章作者
      * title：文章标题
      * content：文章内容
      */
    save: function(author, title, content){
        if(!author || !title || !content){
            throw new Error("文章标题、作者和内容不能为空")
        }

        if(author.length > 20) {
          throw new Error("文章作者名字过长")
        }

        if(title.length > 100) {
          throw new Error("文章标题过长")
        }

        if(content.length > 100000){
            throw new Error("文章内容过长")
        }

        var from = Blockchain.transaction.from;
        var article = this.articleData.get(title);
        if(article){
            throw new Error("该文章已存在，请勿重新发布");
        }

        article = new Article();
        article.author = author;
        article.title = title;
        article.content = content;
        article.address = from;
        article.date = getCurrentTime();
        this.articleData.put(title, article);
    },

    /**
      * 根据文章标题或者对应的文章
      * title：文章标题
      */
    get: function(title){
        if(!title){
            throw new Error("文章标题为空")
        }
        return this.articleData.get(title);
    }
}

module.exports = ArticleContract;
```

#### 2.2 部署智能合约

跟以太坊平台一样，在星云链上部署智能合约写入区块链中需要支付一定的费用（即Gas）。
支付Gas的目的主要有如下几个：
1.保证节点正常运行防止蓄意攻击；
2.用于矿工的奖励；
3.防止恶意发布代码（如执行死循环操作）。

切换到Web钱包的"合约"选项，然后选择"部署"。

![11.切换到合约选项](media/15286049569309/11.%E5%88%87%E6%8D%A2%E5%88%B0%E5%90%88%E7%BA%A6%E9%80%89%E9%A1%B9.png)

将合约代码粘贴到输入框，点击"选择钱包文件"进行上传，输入钱包密码后再点击"解锁"，会出现如下内容。

![12.部署合约](media/15286049569309/12.%E9%83%A8%E7%BD%B2%E5%90%88%E7%BA%A6.png)

点击"测试"按钮后，发现没有返回错误信息。

![13.测试和提交智能合约结果](media/15286049569309/13.%E6%B5%8B%E8%AF%95%E5%92%8C%E6%8F%90%E4%BA%A4%E6%99%BA%E8%83%BD%E5%90%88%E7%BA%A6%E7%BB%93%E6%9E%9C.png)


然后点击右边的"提交"按钮进行部署，成功后会返回交易的哈希以及合约地址（记得做好保存）。

#### 2.3 测试智能合约

切换到"执行"选项，函数输入要测试的名称，如"save"，参数这里传如下内容：

```bash
["Wenzil","开发和部署星云链DApp","经过前面一系列的基础学习，本篇将进入DApp开发阶段。\n\n在星云链上，部署智能合约需要星云币，所以需要申请一个星云链的钱包。。。"]  
```

然后跟之前一样，上传钱包文件，输入密码，点击"解锁"、"测试"和"提交"按钮一系列操作。

![14.测试智能合约save函数](media/15286049569309/14.%E6%B5%8B%E8%AF%95%E6%99%BA%E8%83%BD%E5%90%88%E7%BA%A6save%E5%87%BD%E6%95%B0.png)


合约地址为：
n1yE6rWrGo8yd1E9Wzn6CqZB85gyLF65YdL

然后，可以测试get函数，如下

![15.测试智能合约get函数](media/15286049569309/15.%E6%B5%8B%E8%AF%95%E6%99%BA%E8%83%BD%E5%90%88%E7%BA%A6get%E5%87%BD%E6%95%B0.png)

来到这里，表示已经成功部署合约并测试通过。

### 3、安装相关文件和插件

前端页面中需要用到两个js文件，分别是neb.js和nebPay.js，用于JavaScript中API的调用和支付接口的调用。 

#### 3.1 安装neb.js

neb.js提供了JavaScript开发的API接口，需要从GitHub下载。

1、clone项目neb.js

```bash
git clone https://github.com/nebulasio/neb.js.git
```

终端操作如下：

```bash
wenzildeiMac:study wenzil$ pwd
/Users/wenzil/Desktop/study
wenzildeiMac:study wenzil$ git clone https://github.com/nebulasio/neb.js.git
Cloning into 'neb.js'...
remote: Counting objects: 674, done.
remote: Compressing objects: 100% (95/95), done.
remote: Total 674 (delta 54), reused 64 (delta 20), pack-reused 557
Receiving objects: 100% (674/674), 1.05 MiB | 356.00 KiB/s, done.
Resolving deltas: 100% (377/377), done.
```

2、进入到clone项目的根目录，安装所有依赖


```bash
npm install
```

终端操作如下：

```bash
wenzildeiMac:study wenzil$ cd neb.js
wenzildeiMac:neb.js wenzil$ npm install
```

3、安装自动化构建工具gulp

```bash
npm install gulp
```

4、打包生成neb.js等文件

```bash
gulp
```

如果出现如下错误，用以下另外一种命令执行安装gulp

```bash
-bash: gulp: command not found

sudo npm install --global gulp
```


5、执行成功会生成"dist"文件夹，文件夹中有前端页面要使用的js文件。

![16.安装nebjs](media/15286049569309/16.%E5%AE%89%E8%A3%85nebjs.png)


#### 3.2 安装nebPay.js

NebPay SDK为不同平台的交易提供了统一的支付接口，开发者在Dapp页面中使用NebPay API可以通过浏览器插件钱包、手机App钱包等实现交易支付和合约调用。

github地址：https://github.com/nebulasio/nebPay

安装方法同上，会生成nebPay.js文件

终端操作如下：

```bash
git clone https://github.com/nebulasio/nebPay
cd nebPay/
npm install
gulp
```

![17.安装nebPay](media/15286049569309/17.%E5%AE%89%E8%A3%85nebPay.png)

#### 3.3 安装Chrome钱包插件

调用netPay的api需要唤醒Chrome钱包插件进行支付，因此需要安装钱包插件，查看GitHub有详细的安装说明

GitHub地址：

```bash
https://github.com/ChengOrangeJu/WebExtensionWallet
```

### 4、创建和运行DApp

#### 4.1 创建DApp项目

前面创建了一个工程名为"my_web_wallet"的文件夹，将打包生成的nebulas.js和nebPay.min.js文件放到lib文件夹（jquery-3.3.1.min.js也放到该文件夹），将article_publish.js放到contracts文件夹，新建index.html页面文件，这里用到Bootstrap框架开发简单的前端页面，工程结构如下：

![18.项目工程结构](media/15286049569309/18.%E9%A1%B9%E7%9B%AE%E5%B7%A5%E7%A8%8B%E7%BB%93%E6%9E%84.png)


index.html源码如下：

```bash
<!DOCTYPE html>
<html>
<head>
  <title>文章发布</title>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <link href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.1/css/bootstrap.min.css" rel="stylesheet">
</head>

<body>
  <script type="text/javascript" src="./lib/nebulas.js"></script>
  <script type="text/javascript" src="./lib/nebPay.min.js"></script>
  <script type="text/javascript" src="./lib/jquery-3.3.1.min.js"></script>
  <script src="http://cdn.bootcss.com/bootstrap/3.3.1/js/bootstrap.min.js"></script>

  <div class="container">
    <div class="row clearfix">
      <div class="col-md-12 column">

        <h3>发布文章</h3>

        <form>
          <div class="form-group" style="margin-top:20px">
            <label for="author">文章作者：</label>
            <input id="input_author" type="text" class="form-control" id="author" placeholder="输入文章作者">

            <label for="title" style="margin-top:10px">文章标题：</label>
            <input id="input_title" type="text" class="form-control" id="title" placeholder="输入文章标题">
            
            <label for="content" style="margin-top:10px">文章内容：</label>
            <input id="input_content" type="text" class="form-control" id="content" placeholder="输入文章内容">
          </div>

          <div style="text-align:center;margin-top:20px">
            <button id="post" type="submit" class="btn btn-primary" style="padding:6px 16px">发布</button>
          </div>
        </form>

        <h3 style="margin-bottom:10px">查询文章</h3>

        <div class="input-group">
          <input type="text" class="form-control" placeholder="请输入文章标题" id="text_search" />
          <span class="input-group-btn">
            <button id="button_search" class="btn btn-primary">搜索</button>
          </span>
        </div>

      </div>
      <div class="row clearfix">
        <div class="col-md-12 column">
        </div>
      </div>

      <div style="text-align:center;margin-top:30px">
        <h4 id="article_title"></h4>
      </div>
      <div style="text-align:center">
        <span id="article_author"></span><span id="article_date" style="margin-left:20px"></span>
      </div>
      <div id="article_content" style="margin-top:10px"></div>
    </div>
  </body>

  <script>
  "use strict";

  // 智能合约的地址
  var contactAddress = "n1yE6rWrGo8yd1E9Wzn6CqZB85gyLF65YdL";
  // 访问星云链的API
  var nebulas = require("nebulas"), Account = Account, neb = new nebulas.Neb();
  // 设置使用的网络，主网为"https://mainnet.nebulas.io"，测试网络为
  neb.setRequest(new nebulas.HttpRequest("https://testnet.nebulas.io"));

  // NebPay SDK 为不同平台的交易提供了统一的支付接口
  // 开发者在Dapp页面中使用NebPay API可以通过浏览器插件钱包、手机App钱包等实现交易支付和合约调用。
  var NebPay = require("nebpay");
  var nebPay = new NebPay();

  // 执行合约返回的交易流水号，用于查询交易信息
  var serialNumber;
  $("#button_search").click(function () {
    if (!$("#text_search").val()) {
      alert('搜索文章标题不能为空');
      return;
    }
    $('#article_content').text("");
    var from = contactAddress
    var value = "0";  // 金额
    var nonce = "0";  // 交易序号
    var gas_price = "1000000";  // 手续费价格
    var gas_limit = "2000000";  // 手续费限制
    var callFunction = "get";
    var callArgs = "[\"" + $("#text_search").val() + "\"]";
    console.log(callArgs);
    var contract = {  // 合约
      "function": callFunction,  // 方法名
      "args": callArgs    // 方法的参数
    }

    // 使用api.call执行合约中的方法
    neb.api.call(from, contactAddress, value, nonce, gas_price, gas_limit, contract).then(function (resp) {
      var result = resp.result;

      if (result === 'null') {
        $('#article_content').text("没有找到对应的文章");
        // 清空信息
        $('#article_title').text("");
        $('#article_author').text("");
        $('#article_date').text("");
        return;
      }
      console.log(result);
      result = JSON.parse(result);
      $("#article_title").text(result.title);
      $('#article_content').text(result.content);
      $('#article_author').text("文章作者：" + result.author);
      $('#article_date').text("发布时间：" + result.date);
    }).catch(function (err) {
      console.log("error :" + err.message);
    })
  })
  $('#post').click(function () {
    if (!$("#input_author").val() || !$("#input_title").val() || !$("#input_content").val()) {
      alert('作者、标题、内容输入不能为空');
      return;
    }
    var to = contactAddress;
    var value = "0";
    var callFunction = "save";
    var callArgs = "[\"" + $("#input_author").val() + "\",\""  + $("#input_title").val() + "\",\"" + $("#input_content").val() + "\"]";
    console.log(callArgs);
    serialNumber = nebPay.call(to, value, callFunction, callArgs, {    //使用nebpay的call接口去调用合约,
      listener: function (resp) {
        console.log("nebPay response:" + resp)
      }
    });
  })
</script>

</html>
```

#### 4.2 运行DApp

直接打开indext.html，界面如下：

![19.打开网页](media/15286049569309/19.%E6%89%93%E5%BC%80%E7%BD%91%E9%A1%B5.png)


输入文章作者、标题、内容后，点击"发布"按钮，会弹出Chrome钱包插件。如果之前没有导入钱包文件，需要解锁。

![20.弹出钱包插件](media/15286049569309/20.%E5%BC%B9%E5%87%BA%E9%92%B1%E5%8C%85%E6%8F%92%E4%BB%B6.png)

如果之前导入过钱包，会弹出如下信息，不需要手动填写任何内容。
![21.弹出生成交易信息](media/15286049569309/21.%E5%BC%B9%E5%87%BA%E7%94%9F%E6%88%90%E4%BA%A4%E6%98%93%E4%BF%A1%E6%81%AF.png)


点击"生成交易"按钮，会弹出如下内容
![22.确认交易对话框](media/15286049569309/22.%E7%A1%AE%E8%AE%A4%E4%BA%A4%E6%98%93%E5%AF%B9%E8%AF%9D%E6%A1%86.png)

再点击"确认"按钮，会弹出交易状态详情信息
![23.交易状态详情](media/15286049569309/23.%E4%BA%A4%E6%98%93%E7%8A%B6%E6%80%81%E8%AF%A6%E6%83%85.png)

过了一会，发现交易成功。
这时回到页面，输入提交时填写的文章标题来查询（即"测试标题"）来检查。

![24.检查发布的文章内容](media/15286049569309/24.%E6%A3%80%E6%9F%A5%E5%8F%91%E5%B8%83%E7%9A%84%E6%96%87%E7%AB%A0%E5%86%85%E5%AE%B9.png)

发现，跟发布的内容一样，说明数据已经写入了区块链。当然，发布时间不是东八区的时间，时间格式也有点问题（转换时间格式的JS代码有点问题）。

当然，可以查看部署合约时创建的文章，同时也可以查看打印输出的内容。
![25.查看部署合约时创建的文章](media/15286049569309/25.%E6%9F%A5%E7%9C%8B%E9%83%A8%E7%BD%B2%E5%90%88%E7%BA%A6%E6%97%B6%E5%88%9B%E5%BB%BA%E7%9A%84%E6%96%87%E7%AB%A0.png)


这里运行的只是本地html文件，需要部署到服务器上，别人才能访问。该DApp功能比较简陋，还有很多需要完善的地方，这里不作继续谈论。

至此，整个星云链DApp开发流程已经走完了，可以去主网提交DApp了，哈哈。




