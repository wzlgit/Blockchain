package com.wenzil.blockchain;

import com.alibaba.fastjson.JSON;
import com.wenzil.blockchain.bean.Block;
import com.wenzil.blockchain.bean.Transaction;
import com.wenzil.blockchain.util.CryptoUtil;
import com.wenzil.blockchain.service.BlockService;
import java.util.ArrayList;
import java.util.List;

public class Main {
    public static void main(String[] args) {
        // 创建一个空的区块链
        List<Block> blockchain = new ArrayList<>();
        // 生成创世区块
        Block block = new Block(1, System.currentTimeMillis(), new ArrayList<Transaction>(), 1, "1", "1");
        // 加入创世区块到区块链里
        blockchain.add(block);
        System.out.println(JSON.toJSONString(blockchain));

        // 发送方钱包地址
        String sender = "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa";
        // 接收方钱包地址
        String recipient = "1CZPTGSkwX7D9PJgtFaVK7MGQseAnBwQ1h";

        // 创建一个空的交易集合
        List<Transaction> txs = new ArrayList<>();
        // 挖矿
        BlockService.mineBlock(blockchain, txs, sender);

        System.out.println(sender + "钱包的余额为：" + BlockService.getWalletBalance(blockchain, sender));

        // 创建一个空的交易集合
        List<Transaction> txs1 = new ArrayList<>();
        // 已发生但未记账的交易记录，发送者给接收者转账3个比特币
        Transaction tx1 = new Transaction(CryptoUtil.UUID(), sender, recipient, 3);
        // 已发生但未记账的交易记录，发送者给接收者转账1个比特币
        Transaction tx2 = new Transaction(CryptoUtil.UUID(), sender, recipient, 1);
        txs1.add(tx1);
        txs1.add(tx2);
        // 挖矿
        BlockService.mineBlock(blockchain, txs1, sender);
        System.out.println("钱包'" + sender + "'的余额为：" + BlockService.getWalletBalance(blockchain, sender));
        System.out.println("钱包'" + recipient + "'的余额为：" + BlockService.getWalletBalance(blockchain, recipient));
    }
}