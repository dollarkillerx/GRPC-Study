# GRPC-Study
GRPC Study

文档地址:https://www.kancloud.cn/adapa/go-grpc

### 基于TLS鉴权
```
制作私钥 (.key)
# Key considerations for algorithm "RSA" ≥ 2048-bit
openssl genrsa -out server.key 2048
    
# Key considerations for algorithm "ECDSA" ≥ secp384r1
# List ECDSA the supported curves (openssl ecparam -list_curves)
openssl ecparam -genkey -name secp384r1 -out server.key
# 自签名公钥(x509) (PEM-encodings .pem|.crt)
openssl req -new -x509 -sha256 -key server.key -out server.pem -days 3650
```

``` 
Can't load /home/dollarkiller/.rnd into RNG
139622219313600:error:2406F079:random number generator:RAND_load_file:Cannot open file:../crypto/rand/randfile.c:88:Filename=/home/dollarkiller/.rnd
You are about to be asked to enter information that will be incorporated
into your certificate request.
What you are about to enter is what is called a Distinguished Name or a DN.
There are quite a few fields but you can leave some blank
For some fields there will be a default value,
If you enter '.', the field will be left blank.
-----
Country Name (2 letter code) [AU]:USA
string is too long, it needs to be no more than 2 bytes long
Country Name (2 letter code) [AU]:US
State or Province Name (full name) [Some-State]:ssr
Locality Name (eg, city) []:ssr
Organization Name (eg, company) [Internet Widgits Pty Ltd]:ssr
Organizational Unit Name (eg, section) []:ssr
Common Name (e.g. server FQDN or YOUR name) []:ssr    (注意这个名称非常重要)
Email Address []:ssr
```