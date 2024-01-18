package usecase

import (
	"bytes"
	"context"
	eventRepo "customerlabs/business/events/repository"
	"customerlabs/domain/entity"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"sync"
)

type IEventUC interface {
	ProcessEvent(ctx context.Context, req map[string]interface{})
}

type EventUC struct {
	eventChan chan<- map[string]interface{}
	repo eventRepo.IEventRepo
}

func NewEventUC(r eventRepo.IEventRepo, eventChannel chan<- map[string]interface{}) IEventUC {
	return &EventUC{repo: r, eventChan: eventChannel }
}

func (eUC *EventUC) ProcessEvent(ctx context.Context, req map[string]interface{}) {
	eUC.eventChan <- req
} 

func convertEventRequest(req map[string]interface{}) (*entity.ConvertedRequest, error) {
	convRequest := entity.ConvertedRequest{
		Event:           getValue(req, "ev"),
		EventType:       getValue(req, "et"),
		AppID:           getValue(req, "id"),
		UserID:          getValue(req, "uid"),
		MessageID:       getValue(req, "mid"),
		PageTitle:       getValue(req, "t"),
		PageURL:         getValue(req, "p"),
		BrowserLanguage: getValue(req, "l"),
		ScreenSize:      getValue(req, "sc"),
		Attributes:      map[string]entity.Attribute{},
		UserTraits:      map[string]entity.UserTrait{},
	}

	err := processAttributes(req, &convRequest.Attributes)
	if err != nil {
		return nil, err
	}

	err = processUserTraits(req, &convRequest.UserTraits)
	if err != nil {
		return nil, err
	}

	return &convRequest, nil
}

func processAttributes(req map[string]interface{},  attributes *map[string]entity.Attribute) (error) {
	var err error 
	keyPrefix, valuePrefix, typePrefix := "atrk", "atrv", "atrt"
	for key, value := range req {
		if attributeKey, isAttribute := getKey(key, keyPrefix); isAttribute {
			attr := entity.Attribute{}

				attr.Value, err = getStringValue(req, fmt.Sprintf("%s%s", valuePrefix, attributeKey))
				if err != nil {
					return errors.New("Value-attribute type error, not a string")
				}
				attr.Type, err =  getStringValue(req, fmt.Sprintf("%s%s", typePrefix, attributeKey))
				if err != nil {
					return errors.New("Type-attribute type error, not a string")
				}
				if val, ok := value.(string); ok {
					(*attributes)[val] = attr
				} else {
					return errors.New("Key-attribute type error, not a string")
				}
		}
	}
	return nil
}

func processUserTraits(req map[string]interface{}, usertraits *map[string]entity.UserTrait) (error) {
	var err error 
	keyPrefix, valuePrefix, typePrefix := "uatrk", "uatrv", "uatrt"
	for key, value := range req {
		if userTraitKey, isUserTrait := getKey(key, keyPrefix); isUserTrait {
			attr := entity.UserTrait{}

				attr.Value, err = getStringValue(req, fmt.Sprintf("%s%s", valuePrefix, userTraitKey))
				if err != nil {
					return errors.New("Value-user trait type error, not a string")
				}
				attr.Type, err =  getStringValue(req, fmt.Sprintf("%s%s", typePrefix, userTraitKey))
				if err != nil {
					return errors.New("Type-user trait type error, not a string")
				}
				if val, ok := value.(string); ok {
					(*usertraits)[val] = attr
				} else {
					return errors.New("Key-user trait type error, not a string")
				}
		}
	}
	return nil
}
	

func getKey(fullKey, keyPrefix string) (string, bool) {
	if len(fullKey) > len(keyPrefix) && fullKey[:len(keyPrefix)] == keyPrefix {
		return fullKey[len(keyPrefix):], true
	}
	return "", false
}

func getValue(data map[string]interface{}, key string) (string) {
	if value, ok := data[key].(string); ok {
		return value
	}
	return ""
}

func getStringValue(data map[string]interface{}, key string) (string, error) {
	if value, ok := data[key].(string); ok {
		return value, nil
	}
	return "", errors.New("string type error")
}

func Worker(eventChan <-chan map[string]interface{}, shutdown <-chan os.Signal) {
	var wg sync.WaitGroup
	for {
		select {
		case req, ok := <-eventChan:
			if !ok {
				return
			}
			wg.Add(1)
			go func(req map[string]interface{}) {
				defer wg.Done()
				processedReq, err := convertEventRequest(req)
				if err != nil {
					fmt.Println(err)
					return
				}
				err = sendToWebhook(processedReq)
				if err != nil {
					fmt.Println(err)
					return
				}
			}(req)
		case <-shutdown:
			wg.Wait()
			return
		}
	}
}

func sendToWebhook(req *entity.ConvertedRequest) (error){
	jsonData, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("error marshaling EventResponse to JSON: %v", err)
	}
	//webhook.site url generated
	webhookURL := "https://webhook.site/627bb56e-e865-48fb-a005-d4d63c193275"
	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error making HTTP POST request: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	return nil
}