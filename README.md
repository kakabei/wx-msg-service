# 背景

一个公众号只能有一个消息回调，但实际开发时，却有开发环境、测试环境、和生产环境。公众号的回调无法满足三个环境。

所以写了一个中间服务，按 OpenId 转到各个不同的环境。这样可以满足开发、测试的需求，同时也不会影响生产环境的正常使用。

# 服务

服务 wx-msg-service  用 go-zore 的框架进行开发。

提供了两个接口：

1、GET /v1/service/wx/wxmsg 设置微信公众号的回调时，微信公众号平台会对接口进行 get 的认证。

2、POST /v1/service/wx/wxmsg 处理回调消息的接口。

服务的两个主要功能：

1、过滤不相关的消息事件

2、转到不同的服务

# 配置

```yaml

WxMsgMgr:
  AllowMsgEvent: ["SCAN","subscribe", ""] # 支持的消息事件, 不在列表中则被过滤掉
  WxOpenIdList:  
    - EnvName : "dev" 
      AllowMsgEvent: ["SCAN","subscribe", ""] # 该环境的转到消息事件
      HandleUrl: http://127.0.0.1:22255/v1/client/wxmsg1 # 回调地址
      OpenIdList: ["o-iUq6qghnE5dHvzRo4JGokOr0lk11"]     # 指定OpenId
    
    - EnvName : "beta" 
      AllowMsgEvent: ["SCAN","subscribe"]
      HandleUrl: http://127.0.0.1:22255/v1/client/wxmsg2
      OpenIdList: ["o-iUq6qghnE5dHvzRo4JGokOr0lk22"]
      
    - EnvName : "idc" 
      AllowMsgEvent: ["SCAN","subscribe", ""] 
      HandleUrl: http://127.0.0.1:22255/v1/client/wxmsg3
      OpenIdList: []

```
