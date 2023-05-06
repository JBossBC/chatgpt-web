package dao

import (
	"bufio"
	"bytes"
	klog "chatgpt-web/internal/app/log"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

type Body struct {
	Params Params `json:"params"`
	//default keep the 10 length for per session
	Messages []*Message `json:"messages"`
}

const defaultMessageLen = 10

var globalMessagesLen = defaultMessageLen

func (body *Body) Packed(key string, curMessage Message) {
	// inspect the pre message  role
	if curMessage.Role == "" {
		curMessage.Role = body.Params.Role
	}
	if len(body.Messages) >= globalMessagesLen {
		newMessage := make([]*Message, globalMessagesLen)
		copy(newMessage, body.Messages[1:])
		newMessage[globalMessagesLen-1] = &curMessage
	} else {
		body.Messages = append(body.Messages, &curMessage)
	}
	bodyMap.Store(key, body)
}

func (body *Body) Marshal() ([]byte, error) {
	paramsResult, err := json.Marshal(&body.Params)
	if err != nil {
		return []byte{}, err
	}
	sb := strings.Builder{}
	sb.WriteString(string(paramsResult[:len(paramsResult)-1]))
	sb.WriteByte(',')

	messagesResult, err := json.Marshal(&body.Messages)
	if err != nil {
		return []byte{}, err
	}
	sb.WriteString("\"messages\":")
	sb.WriteString(string(messagesResult))
	sb.WriteString("}")
	return []byte(sb.String()), nil

}

type chatgptClient http.Client

var (
	singleChatgptClient *chatgptClient
	chatOnce            sync.Once
)

func GetChatgptClient() *chatgptClient {
	chatOnce.Do(func() {
		singleChatgptClient = &chatgptClient{}
	})
	return singleChatgptClient
}

/**
  must pre-packed
*/
func (client *chatgptClient) send(data *Body) (string, error) {
	req, err := packedRequest(data)
	if err != nil {
		return "", err
	}
	response, err := (*http.Client)(client).Do(req)
	if err != nil {
		return "", nil
	}
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		klog.Error(fmt.Sprintf("%v response read error:%s", result, err.Error()))
		return "", err
	}
	var resultMap ChatResponse
	err = json.NewDecoder(bytes.NewReader([]byte(result))).Decode(&resultMap)
	if err != nil {
		klog.Error(err)
		return "", err
	}
	return string(resultMap.Choices[0].Message.Content), err
}

//TODO need to complete
func (client *chatgptClient) SendQuestion(key string, message Message) (string, error) {
	loadBody, ok := bodyMap.Load(key)
	if !ok {
		loadBody = &Body{
			Params:   defaultParams,
			Messages: make([]*Message, 0, defaultMessageLen),
		}
	}
	loadBody.(*Body).Packed(key, message)
	return client.send(loadBody.(*Body))
}
func (client *chatgptClient) SendQuestionRaw(key string, message Message) (*http.Response, error) {
	loadBody, ok := bodyMap.Load(key)
	if !ok {
		loadBody = &Body{
			Params:   defaultParams,
			Messages: make([]*Message, 0, defaultMessageLen),
		}
	}
	loadBody.(*Body).Packed(key, message)
	return client.sendRaw(loadBody.(*Body))
}
func (client *chatgptClient) sendRaw(data *Body) (*http.Response, error) {
	req, err := packedRequest(data)
	if err != nil {
		return nil, err
	}
	return (*http.Client)(client).Do(req)
}

const defaultChatgptURL = "https://api.openai-proxy.com/v1/chat/completions"

func packedRequest(data *Body) (*http.Request, error) {
	//update the request header
	var err error
	jsonData, err := data.Marshal()
	if err != nil {
		return nil, err
	}
	res, err := http.NewRequest(http.MethodPost, defaultChatgptURL, bufio.NewReader(bytes.NewReader(jsonData)))
	if err != nil {
		return nil, err
	}
	res.Header.Add("Content-Type", "application/json")
	res.Header.Add("Authorization", fmt.Sprintf(" Bearer %s", data.Params.Key))
	return res, nil

}

//global body cache
var bodyMap sync.Map
