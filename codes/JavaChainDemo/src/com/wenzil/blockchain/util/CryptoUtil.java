package com.wenzil.blockchain.util;

import java.security.MessageDigest;
import java.util.UUID;

/**
 * 加密工具类
 */
public class CryptoUtil {

    private CryptoUtil() {
    }

    /***
     * SHA256加密处理
     * @param str 原始字符串
     * @return 返回加密后的哈希值
     */
    public static String SHA256(String str) {
        // MessageDigest类为应用程序提供的信息摘要算法
        MessageDigest messageDigest;
        String encodeStr = "";
        try {
            // 使用SHA-256算法
            messageDigest = MessageDigest.getInstance("SHA-256");
            // 使用UTF-8编码更新摘要
            messageDigest.update(str.getBytes("UTF-8"));
            // 将bytes数组转换为十六进制值的字符串
            encodeStr = byte2Hex(messageDigest.digest());
        } catch (Exception e) {
            System.out.println("getSHA256 is error" + e.getMessage());
        }
        return encodeStr;
    }

    /**
     * 生成UUID
     * @return
     */
    public static String UUID() {
        return UUID.randomUUID().toString().replaceAll("\\-", "");
    }

    /**
     * 将bytes数组转换为十六进制值的字符串
     * @param bytes
     * @return
     */
    private static String byte2Hex(byte[] bytes) {
        StringBuilder builder = new StringBuilder();
        String temp;
        for (int i = 0; i < bytes.length; i++) {
            temp = Integer.toHexString(bytes[i] & 0xFF);
            if (temp.length() == 1) {
                builder.append("0");
            }
            builder.append(temp);
        }
        return builder.toString();
    }
}