# shorurl
短网址

*演示地址:[vvia.xyz](http://vvia.xyz)*

 API

```
URL: http://vvia.xyz/s
Method: POST
Parmeters: url
Return: JSON
```

示例

```shell
$ curl -d '{"url":"https://github.com/irealing/shorturl"}' vvia.xyz/s
{"err_no":0,"msg":"success","data":"orv2oe"}
```

