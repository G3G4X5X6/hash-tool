# Hash-Tool

字符串、文件哈希计算工具。(Power by G3G4X5X6 & Chat-GPT)

### :dragon: 支持的哈希算法

1. MD5 (Message-Digest Algorithm 5)
    - 一种常用的哈希算法，输出结果为128位二进制数，通常用于校验数据完整性。
2. SHA-1 (Secure Hash Algorithm 1)
    - 一种安全性较低的哈希算法，输出结果为160位二进制数，通常用于数字签名。
3. SHA-256 (Secure Hash Algorithm 256)
    - 一种较为安全的哈希算法，输出结果为256位二进制数，通常用于数字签名和数据完整性校验。
4. SHA-512 (Secure Hash Algorithm 512)
    - 一种安全性非常高的哈希算法，输出结果为512位二进制数，通常用于加密和数字签名。
5. CRC (Cyclic Redundancy Check)
    - 一种较为简单的哈希算法，输出结果为32位二进制数，通常用于校验数据的完整性和传输错误检测。

### :camera: 使用帮助
```bash
PS E:\0x02.GitHub\hash-tool\bin> .\hash-windows-amd64.exe -h                                                                                                                                                   
NAME:
   hash - calculate the core value of a file or string

USAGE:
   hash [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --enable value, -e value [ --enable value, -e value ]  enabled hashing algorithms (default: "crc32", "md5", "sha1", "sha256", "sha512")
   --path value, -p value [ --path value, -p value ]      file path
   --thread value, -t value                               number of threads (default: 6)
   --string value, -s value                               computes a string core
   --help, -h                                             show help

——————————————————————————————————————————————————————————
v 1.0.0
——————————————————————————————————————————————————————————
Program runtime: 3.7783ms
——————————————————————————————————————————————————————————

```

### :rocket: 使用示例

#### 计算 `字符串` 哈希

```bash
PS E:\0x02.GitHub\hash-tool\bin> .\hash-windows-amd64.exe -s "666666"

——————————————————————————————————————————————————————————
Text: 666666
——————————————————————————————————————————————————————————
crc32:   4269814406
md5:     F379EAF3C831B04DE153469D1BEC345E
sha1:    1411678A0B9E25EE2F7C8B2F7AC92B6A74B3F9C5
sha256:  94EDF28C6D6DA38FD35D7AD53E485307F89FBEAF120485C8D17A43F323DEEE71
sha512:  4D6698C1E6B230C5FF80935BC26B722E743487B05B0DA398D4716AD43E725C17E8B02CCABC2C593B5DD9547191386AE5ADD75A8873B3162D9544B7EBA873A170
——————————————————————————————————————————————————————————
Done!


——————————————————————————————————————————————————————————
v 1.0.0
——————————————————————————————————————————————————————————
Program runtime: 2.6818ms
——————————————————————————————————————————————————————————

```

#### 计算 `单个文件` 哈希
```bash
PS E:\0x02.GitHub\hash-tool\bin> .\hash-windows-amd64.exe -p "E:\0x04. ISO\CentOS-7-x86_64-DVD-2009.iso"
⠧ Hashing...... (1342623900/-, 42950182 it/s) [31s] 
——————————————————————————————————————————————————————————
Text: E:\0x04. ISO\CentOS-7-x86_64-DVD-2009.iso
——————————————————————————————————————————————————————————
crc32:   7ABB1DA9
md5:     5A3B0B3F4A6654EE195EC6EDB6D938A4
sha1:    6A417E29D97829D038D73FF1584D53208A1415DB
sha256:  E33D7B1EA7A9E2F38C8F693215DD85254C3A4FE446F93F563279715B68D07987
sha512:  7061FA737086370716885439353B50DA47D239AD81B971A0427EBB69B0CA65BCD7309BF211827ECFCCD8A725B10FF5CC2A2CE395AF5BCC559D08D7CE8B56FFE7
——————————————————————————————————————————————————————————

——————————————————————————————————————————————————————————
v 1.0.0
——————————————————————————————————————————————————————————
Program runtime: 31.8060896s
——————————————————————————————————————————————————————————

```

#### 计算 `多个文件` 哈希
```bash
PS E:\0x02.GitHub\hash-tool\bin> .\hash-windows-amd64.exe -p "E:\0x02.GitHub\hash-tool\bin\hash-darwin-amd64,E:\0x02.GitHub\hash-tool\bin\hash-linux-amd64,E:\0x02.GitHub\hash-tool\bin\hash-windows-amd64.exe"

⠋ Hashing...... (0/-, 0 it/hr) [0s] 
——————————————————————————————————————————————————————————
Text: E:\0x02.GitHub\hash-tool\bin\hash-linux-amd64
——————————————————————————————————————————————————————————
crc32:   3AD30B15
md5:     6FFEAE25A05BC4BD8DBE21FB4E01AE77
sha1:    B8399E478F0EB60452139FE8350A23C87C807D8A
sha256:  CBF08AE6E9E18E522B7D822A6E21358D6F4E6ECAEC8FB55D2EFFA1FA44B8A1FD
sha512:  D292068464EDD5D8F89A6E627095ADADC135BD82BA60FAF01450C983CE08D22DF7CE07678C2A910CD2DE8D58D5779A6CFB806D1A0F0D75C8C299FD162028E8BF
——————————————————————————————————————————————————————————

——————————————————————————————————————————————————————————
Text: E:\0x02.GitHub\hash-tool\bin\hash-darwin-amd64
——————————————————————————————————————————————————————————
crc32:   360CED5A
md5:     582875B37F169E056464C1431636794E
sha1:    3841F59E948662142519B1EBEFEA4BCD2447B3AF
sha256:  20057C54AC3C70EF0B25767F41B83FFCE573491045CCEA16BC9E675BCC8A942F
sha512:  1101B7D5439D3C6865A0228382888956E7F879C3C04B7AAB566562CC3089AF984410EAC091BDDECEC787B2A1D3D1D004D4E0350E18AF6309808980E1F27AC717
——————————————————————————————————————————————————————————

——————————————————————————————————————————————————————————
Text: E:\0x02.GitHub\hash-tool\bin\hash-windows-amd64.exe
——————————————————————————————————————————————————————————
crc32:   7EB05CCF
md5:     7A0D1BE6DD6DF81806250A427C8EAE83
sha1:    E863628C6060690A171E6E419705BB603E3CF2A4
sha256:  DB4D07BD805C089F3B63B662F1BE96632369F0B798906601F2FA52EDA005AA21
sha512:  05F513BAD7FD76681EDBA5622601FA8D8B7FB38B85220E603A7A9C2B4CCEE1D515661564FFF7064D932B01DDFEDBA16C248CE33E715ECA80050249CB5F5AF22F
——————————————————————————————————————————————————————————

——————————————————————————————————————————————————————————
v 1.0.0
——————————————————————————————————————————————————————————
Program runtime: 55.1124ms
——————————————————————————————————————————————————————————

```

#### `递归目录` 计算所有文件哈希
```bash
PS E:\0x02.GitHub\hash-tool\bin> .\hash-windows-amd64.exe -p "E:\0x02.GitHub\hash-tool\core"                                                                                                                   
⠋ Hashing...... (0/-, 0 it/hr) [0s] 
——————————————————————————————————————————————————————————
Text: E:\0x02.GitHub\hash-tool\core\file.go
——————————————————————————————————————————————————————————
crc32:   9159FB88
md5:     130991C7DEC6268E4A5B568349E129FF
sha1:    3BBB96E349092799E624867536A9B6BBD3AF40C5
sha256:  8E0A7A4B9C89FDEFE6B2D70FAA175C5393AD1DA83CEAC6B5841403A115A9415D
sha512:  CE17C76D145028ABE69302B71235F1CB8FF632F15759F5A29CB16B00870E83E2B5A7F0031FC4EE0EF6AA4E13A9FC2CE6908CE654D7D9574032B1F4A73995FFE3
——————————————————————————————————————————————————————————


——————————————————————————————————————————————————————————
Text: E:\0x02.GitHub\hash-tool\core\string.go
——————————————————————————————————————————————————————————
crc32:   57E78E92
md5:     D6513F4B929D27D3F71184C494D270B6
sha1:    1FD8F6EA43E24BEB0EEC240542160083A42B1B60
sha256:  CA138517EE76A8CBDB12E580467A7A78C5DAD637FB95AD5E870531DE8EFCF392
sha512:  D346AF0F63061392511B1EC6E365CE2930E1504A5F4FC170E0CEAB987187DE90217215EB3BBBFA73FAE7D1714947E5A88FD1510F5943A0F7C5A5F4C81FFFF997
——————————————————————————————————————————————————————————
——————————————————————————————————————————————————————————
Text: E:\0x02.GitHub\hash-tool\core\struct.go
——————————————————————————————————————————————————————————
crc32:   C09DCD52
md5:     2B76E8C57408EAEB8D01E70691917497
sha1:    9E1B720F8DECA7268EE6A4E94444E00C84F258FD
sha256:  C07D201AE74BFE44C30E1F07F6F1C1BC2ABDBA380C0C2F8DB41FEDCAF0EC907E
sha512:  23AF6F28270A195C3171A12EABBDB7B035C83BDF9178C50F2288830075EE13A1727FEAD7C276D024D8B28D5B15816C4ABD7D9677219EF1B91216AB51F88D771E
——————————————————————————————————————————————————————————

——————————————————————————————————————————————————————————
v 1.0.0
——————————————————————————————————————————————————————————
Program runtime: 13.3592ms
——————————————————————————————————————————————————————————

```

#### 指定哈希算法: ` -e "crc32,md5"`
```bash
PS E:\0x02.GitHub\hash-tool\bin> .\hash-windows-amd64.exe -s "666666" -e "crc32,md5"                                                                                                                           

——————————————————————————————————————————————————————————
Text: 666666
——————————————————————————————————————————————————————————
crc32:   4269814406
md5:     F379EAF3C831B04DE153469D1BEC345E
sha1:
sha256:
sha512:
——————————————————————————————————————————————————————————
Done!


——————————————————————————————————————————————————————————
v 1.0.0
——————————————————————————————————————————————————————————
Program runtime: 2.4479ms
——————————————————————————————————————————————————————————

```
