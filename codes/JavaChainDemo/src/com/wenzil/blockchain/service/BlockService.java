package com.wenzil.blockchain.service;

import com.alibaba.fastjson.JSON;
import com.wenzil.blockchain.bean.Block;
import com.wenzil.blockchain.bean.Transaction;
import com.wenzil.blockchain.util.CryptoUtil;
import java.util.List;

/**
 * 区块服务
 */
public class BlockService {
    /**
     * 模拟PoW挖矿
     * @param blockchain 整个区块链
     * @param txs 需记账交易记录
     * @param address 矿工钱包地址
     * @return
     */
    public static void mineBlock(List<Block> blockchain, List<Transaction> txs, String address) {
        // 比特币第一个创世区块的奖励设定为50个比特币，此后每新建210,000个区块，奖励减半。
        // 奖励金在2012年当时候从50比特币减半为25比特币，2016年从25个比特币减半为12.5个比特币。大概2020年就会再减半为6.25比特币。
        // 加入系统挖矿奖励的交易，这里默认挖矿奖励50个比特币。
        Transaction sysTx = new Transaction(CryptoUtil.UUID(), "", address, 50);
        txs.add(sysTx);
        // 获取当前区块链里的最后一个区块
        Block latestBlock = blockchain.get(blockchain.size() - 1);
        // 随机数
        int nonce = 1;
        String hash = "";
        // 开始挖矿
        while(true){
            hash = CryptoUtil.SHA256(latestBlock.getHash() + JSON.toJSONString(txs) + nonce);
            if (hash.startsWith("0000")) {
                System.out.println("=====计算结果正确，计算次数为：" +nonce+ "，hash：" + hash);
                break;
            }
            nonce++;
            System.out.println("计算错误，hash:" + hash);
        }

        // 解出难题，可以构造新区块并加入进区块链里
        Block newBlock = new Block(latestBlock.getIndex() + 1, System.currentTimeMillis(), txs, nonce, latestBlock.getHash(), hash);
        blockchain.add(newBlock);
        System.out.println("挖矿后的区块链：" + JSON.toJSONString(blockchain));
    }

    /**
     * 查询钱包的余额
     * @param blockchain 整个区块链
     * @param address 钱包地址
     * @return
     */
    public static int getWalletBalance(List<Block> blockchain, String address) {
        int balance = 0;
        for (Block block : blockchain) {
            List<Transaction> transactions = block.getTransactions();
            for (Transaction transaction : transactions) {
                if (address.equals(transaction.getRecipient())) {
                    balance += transaction.getAmount();
                }
                if (address.equals(transaction.getSender())) {
                    balance -= transaction.getAmount();
                }
            }
        }
        return balance;
    }
}