// Code generated by goctl. DO NOT EDIT.
package types

type CommonRet struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

type CheckWxSignatureReq struct {
	Singnature string `form:"signature"`
	Echostr    string `form:"echostr"`
	Timestamp  int64  `form:"timestamp"`
	Nonce      string `form:"nonce"`
}

type CheckWxSignatureResp struct {
	Echostr string `json:"echostr"`
}

type ReceiveWxMsgReq struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int64  `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Event        string `xml:"Event"`
	EventKey     string `xml:"EventKey"`
	Ticket       string `xml:"Ticket"`
}

type ReceiveWxMsgResp struct {
	XmlData string `json:"xml_data"`
}
