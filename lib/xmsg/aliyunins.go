package xmsg

import (
	"encoding/json"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/stringx"
	"gozerobasic/lib/xerr"
)

type AliyunMsgInstance struct {
	regionId string
	accessKeyId string
	accessSecret string
	client *dysmsapi.Client
}

func NewAliyunMsg(regionId,accessKeyId,accessSecret string) *AliyunMsgInstance {
	c ,err := dysmsapi.NewClientWithAccessKey(regionId, accessKeyId, accessSecret)
	if err != nil{
		logx.Errorf("NewAliyunMsg dysmsapi.NewClientWithAccessKey err : %v")
		panic(err)
	}
	return &AliyunMsgInstance{
		client: c,
	}
}

func (m *AliyunMsgInstance)SendSingleMsg(param *SendSingleMsgParams) error {

	return nil

	if stringx.HasEmpty(param.PhoneNumbers) || stringx.HasEmpty(param.SignName) || stringx.HasEmpty(param.TemplateCode){
		logx.Errorf("发送短信失败，缺少必要的参数 param : %v,PhoneNumbers:%s,SignName:%s,TemplateCode:%s",param,param.PhoneNumbers,param.SignName,param.TemplateCode)
		return xerr.NewErrMsg("发送短信失败，缺少必要的参数")
	}

	byteContent,err:=json.Marshal(param.Content)
	if err != nil{
		logx.Errorf("发送短信失败 vclient.SendSms  err : %v ,param:%v ",err,param)
		return xerr.NewErrMsg("发送短信失败，请稍后再试")
	}
	content :=string(byteContent)
	if !stringx.NotEmpty(content){
		logx.Errorf("发送短信内容不能为空 vclient.SendSms  err : %v ,param:%v ",err,param)
		return xerr.NewErrMsg("发送短信内容不能为空")
	}

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https" //默认就用https
	request.PhoneNumbers = param.PhoneNumbers
	request.SignName = param.SignName
	request.TemplateCode = param.TemplateCode
	request.TemplateParam = content
	request.OutId = param.OutId

	response, err := m.client.SendSms(request)
	if err != nil {
		logx.Errorf("发送短信失败 vclient.SendSms  err : %v ,request : %v ,param:%v ,response : %v",err,request,param,response)
		return xerr.NewErrMsg("发送短信失败，请稍后再试")
	}
	return nil
}