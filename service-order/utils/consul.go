package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"net/http"
	"service-order/constant"
	"service-order/dto"
)

func CallConsulFindService(service_name string) (dto.SchemaConsulCatalogService, error) {
	var SchemaConsul dto.SchemaConsulCatalogService
	var ListSchemaConsul []dto.SchemaConsulCatalogService
	headers := make(map[string]interface{})

	URL := fmt.Sprintf("http://%s:%s/v1/catalog/service/%v", dto.CfgConsul.ConsulHost, dto.CfgConsul.ConsulPort, service_name)

	request := dto.CallAPIDto{
		Method:      constant.HTTP_GET,
		Url:         URL,
		ContentType: constant.CONTENT_TYPE_APPLICATION_JSON,
		Headers:     headers,
		//BodyRequest: string(reqBytes),
	}

	err := CallAPI(&request)
	if err != nil {
		log.Errorln("Error CallConsulFindService"+URL, err)
		return SchemaConsul, err
	}

	if request.HttpCode == http.StatusOK {

		err = json.Unmarshal([]byte(request.BodyResponse), &ListSchemaConsul)
		if err != nil {
			log.Errorln("Error Unmarshal Consul", err)
			return SchemaConsul, err
		}

	} else {
		log.Errorln("Error Get Catalog "+URL, request.HttpCode)
		return SchemaConsul, errors.New("Error Call API Catalog ")
	}
	if len(ListSchemaConsul) > 0 {
		min := 0
		max := len(ListSchemaConsul)
		SelectService := rand.Intn(max-min) + min
		return ListSchemaConsul[SelectService], nil
	} else {
		return SchemaConsul, errors.New("Error Load Consul Service")
	}

	return SchemaConsul, err
}
