# shorurl
短网址

*演示地址:[vvia.xyz](http://vvia.xyz)*

## API
### 生成短网址
```
URL: http://vvia.xyz/s
Method: POST
Parmeters: url
Return: JSON
```

#### 示例

```shell
$ curl -d '{"url":"https://github.com/irealing/shorturl"}' vvia.xyz/s
{"err_no":0,"msg":"success","data":"orv2oe"}
```

### 查询短网址真实URL

```
URL: http://vvia.xyz/s/\<string:shorted\>
Method: GET
Parmeters: url
Return: JSON
```
#### 示例

``` shell
$ curl -X GET vvia.xyz/s/orv2oe

{"err_no":0,"msg":"success","data":"https://github.com/irealing/shorturl"}
```
