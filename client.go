package shumei

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	AppID     string
	AccessKey string
	Region    string
	Timeout   time.Duration
}

func NewClient(appId, accessKey string, options ...ClientOptions) *Client {
	c := &Client{
		AppID:     appId,
		AccessKey: accessKey,
		Region:    RegionBeijing,
		Timeout:   time.Second * 3,
	}

	for _, option := range options {
		if err := option(c); err != nil {
			return nil
		}
	}

	return c
}

// TextDetect 文本检测
func (c *Client) TextDetect(eventId string, req TextDetectReq) (*TextDetectResp, error) {
	buf, _ := json.Marshal(struct {
		CommonReq
		TextDetectReq
	}{
		c.getCommonReq(eventId),
		req,
	})

	respBytes, err := c.httpPostJson(ServiceText, buf)
	if err != nil {
		return nil, err
	}

	resp := TextDetectResp{}
	_ = json.Unmarshal(respBytes, &resp)

	return &resp, nil
}

// ImageSyncDetect 图片单条同步检测
func (c *Client) ImageSyncDetect(eventId string, req ImageSyncDetectReq) (*ImageSyncDetectResp, error) {
	buf, _ := json.Marshal(struct {
		CommonReq
		ImageSyncDetectReq
	}{
		c.getCommonReq(eventId),
		req,
	})

	respBytes, err := c.httpPostJson(ServiceImageSync, buf)
	if err != nil {
		return nil, err
	}

	resp := ImageSyncDetectResp{}
	_ = json.Unmarshal(respBytes, &resp)

	return &resp, nil
}

// ImageAsyncDetect 图片单条异步检测
func (c *Client) ImageAsyncDetect(eventId string, req ImageAsyncDetectReq) (*ImageAsyncDetectResp, error) {
	buf, _ := json.Marshal(struct {
		CommonReq
		ImageAsyncDetectReq
	}{
		c.getCommonReq(eventId),
		req,
	})

	respBytes, err := c.httpPostJson(ServiceImageAsync, buf)
	if err != nil {
		return nil, err
	}

	resp := ImageAsyncDetectResp{}
	_ = json.Unmarshal(respBytes, &resp)

	return &resp, nil
}

// ImagesSyncDetect 图片批量同步检测
func (c *Client) ImagesSyncDetect(eventId string, req ImagesSyncDetectReq) (*ImagesSyncDetectResp, error) {
	buf, _ := json.Marshal(struct {
		CommonReq
		ImagesSyncDetectReq
	}{
		c.getCommonReq(eventId),
		req,
	})

	respBytes, err := c.httpPostJson(ServiceImagesSync, buf)
	if err != nil {
		return nil, err
	}

	resp := ImagesSyncDetectResp{}
	_ = json.Unmarshal(respBytes, &resp)

	return &resp, nil
}

// ImagesAsyncDetect 图片批量异步检测
func (c *Client) ImagesAsyncDetect(eventId string, req ImagesAsyncDetectReq) (*ImagesAsyncDetectResp, error) {
	buf, _ := json.Marshal(struct {
		CommonReq
		ImagesAsyncDetectReq
	}{
		c.getCommonReq(eventId),
		req,
	})

	respBytes, err := c.httpPostJson(ServiceImagesAsync, buf)
	if err != nil {
		return nil, err
	}

	resp := ImagesAsyncDetectResp{}
	_ = json.Unmarshal(respBytes, &resp)

	return &resp, nil
}

// ImageQueryDetect 图片查询
func (c *Client) ImageQueryDetect(eventId string, req ImageQueryReq) (*ImageQueryResp, error) {
	buf, _ := json.Marshal(struct {
		CommonReq
		ImageQueryReq
	}{
		c.getCommonReq(eventId),
		req,
	})

	respBytes, err := c.httpPostJson(ServiceImageQuery, buf)
	if err != nil {
		return nil, err
	}

	resp := ImageQueryResp{}
	_ = json.Unmarshal(respBytes, &resp)

	return &resp, nil
}

// AudioSyncDetectDetect 音频同步检测
func (c *Client) AudioSyncDetectDetect(eventId string, req AudioSyncDetectReq) (*AudioSyncDetectResp, error) {
	buf, _ := json.Marshal(struct {
		CommonReq
		AudioSyncDetectReq
	}{
		c.getCommonReq(eventId),
		req,
	})

	respBytes, err := c.httpPostJson(ServiceAudioSync, buf)
	if err != nil {
		return nil, err
	}

	resp := AudioSyncDetectResp{}
	_ = json.Unmarshal(respBytes, &resp)

	return &resp, nil
}

// AudioAsyncDetectDetect 音频异步检测
func (c *Client) AudioAsyncDetectDetect(eventId string, req AudioAsyncDetectReq) (*AudioAsyncDetectResp, error) {
	buf, _ := json.Marshal(struct {
		CommonReq
		AudioAsyncDetectReq
	}{
		c.getCommonReq(eventId),
		req,
	})

	respBytes, err := c.httpPostJson(ServiceAudioAsync, buf)
	if err != nil {
		return nil, err
	}

	resp := AudioAsyncDetectResp{}
	_ = json.Unmarshal(respBytes, &resp)

	return &resp, nil
}

// AudioQueryDetect 音频查询
func (c *Client) AudioQueryDetect(eventId string, req AudioQueryReq) (*AudioQueryResp, error) {
	buf, _ := json.Marshal(struct {
		CommonReq
		AudioQueryReq
	}{
		c.getCommonReq(eventId),
		req,
	})

	respBytes, err := c.httpPostJson(ServiceAudioQuery, buf)
	if err != nil {
		return nil, err
	}

	resp := AudioQueryResp{}
	_ = json.Unmarshal(respBytes, &resp)

	return &resp, nil
}

// 发送 POST 请求
func (c *Client) httpPostJson(event string, buf []byte) ([]byte, error) {
	endpoint, err := c.getGateway(event)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return nil, err
	}

	respBytes, _ := ioutil.ReadAll(resp.Body)
	return respBytes, nil
}

// 获取服务请求网关
func (c *Client) getGateway(service string) (string, error) {
	var gateway string
	gateways, ok := GatewayMap[service]
	if !ok {
		return gateway, fmt.Errorf("shumei service[%v] not found", service)
	}
	gateway, ok = gateways[c.Region]
	if !ok {
		return gateway, fmt.Errorf("shumei service[%v] gateway[%v] not found", service, c.Region)
	}
	return gateway, nil
}

// 获取公共请求参数
func (c *Client) getCommonReq(eventId string) CommonReq {
	return CommonReq{
		AccessKey: c.AccessKey,
		AppID:     c.AppID,
		EventID:   eventId,
	}
}
