# 测试说明


使用 kcagent 加密的数据并base64编码后 
通过 python 版本的zk-updater  进行base64解码并decrypter 解密后可以还原得到源数据
由此验证加密及base64编码有效

```
if __name__ == '__main__':
    sid = "xxxx"
    token = "hello world"
    encrypter = Encrypter()
    encrypted_str = encrypter.encrypt(sid, token)
    encrypted_base64_str = base64.b64encode(encrypted_str)
    print(token)
    print(encrypted_base64_str)
    print("decrypt...")
    encrypted_base64_str = "GBCR3SkiZrfzo7m93vv465blGBIdCLae1nZDlaWfWXrtzGgQtAEYEOqvL0w6CB8S0ElIylgwiwQYFHAIg/          apfXyNkJFpk07Ox4dD3nK3AA=="
    encrypted_str = base64.b64decode(encrypted_base64_str)
    decrypted_str = encrypter.decrypt(sid, encrypted_str)
    print(decrypted_str)

[root@c3-mc-sre00 zk-updater]# python testgokc.py
hello world
GBCqDaywLEMe0REfV1IFhsDTGBJ0kyYs7oREyKxcRjBmkBjF0wEYEKgXH3u5uyGMdllMnWcPkRYYFDmEoB+PYzSwinlKMbbbrQ1I6Oe8AA==
decrypt...
hello world
```