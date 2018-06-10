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
