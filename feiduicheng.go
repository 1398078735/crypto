package main

//非对称加密算法
//定义:加密和解密使用不同的密钥进行操作，非对称加密又称之为公钥加密,
//或者称之为公钥密码加密。
//公钥：可以公开,不担心泄露问题,氪复制可传输
//私钥:私有的，私密的，不可泄露，要妥善保存
//公钥加密,私钥解密适用于信息传输:数据加密!
//私钥加密,公钥解密适用于保证多方接受信息的一致性:数字签名！
//私钥签名,公钥验证,私钥持有者对数据进行数字签名，并将数字签名进行分发，公钥持有者
//适用公钥对接收到的数字签名进行验证，验证通过，表示接收到的信息可信，否则不可信。
//数字签名:使用私钥对数据进行操作室,得到特殊的一段摘要,该摘要内容称之为数字签名。数字签名主用于证明数据由私钥持有者发出
//签名验证:又称为验签，公钥的持有者使用公钥对获取到的数字签名进行验证，如果验证通过，则说明数据可信,没有被篡改，否则数据不可信。

//非对称加密算法:RSA算法,椭圆曲线加密算法

//RSA算法:素数(只能被1和本身整除)
//原理:给定一个大素数s，对素数进行乘式分解，找到两个素数m，n相乘等于s。
//RSA算法:支付宝，微信支付
//一对秘钥:私钥(服务台后台),公钥(支付宝,微信)
//质数:大于1,除1和它本身以外,没有其他因数