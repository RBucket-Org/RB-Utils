package datatype_parser

import (
	"fmt"
	"strconv"

	"github.com/RBucket-Org/RB-Utils/utils/rest_errors"
	"github.com/gin-gonic/gin"
)

var (
	DataTypeParser ParserInterface = &parser{}
)

type ParsingType int

const (
	Int ParsingType = iota
	Float
	Bool
)

type ParserInterface interface {
	ParseParam(context *gin.Context, param string, parsingType ParsingType) (interface{}, rest_errors.RestError)
	ParseQuery(context *gin.Context, query string, parsingType ParsingType) (interface{}, rest_errors.RestError)
}

type parser struct{}

func (p *parser) ParseParam(context *gin.Context, param string, parsingType ParsingType) (interface{}, rest_errors.RestError) {
	rawValue := context.Param(param)
	return p.parse(rawValue, param, parsingType)
}

func (p *parser) ParseQuery(context *gin.Context, query string, parsingType ParsingType) (interface{}, rest_errors.RestError) {
	rawValue := context.Query(query)
	return p.parse(rawValue, query, parsingType)
}

func (p *parser) parse(rawValue string, key string, parsingType ParsingType) (interface{}, rest_errors.RestError) {
	var setParamType string
	var value interface{}
	var err error
	switch parsingType {
	case Int:
		{
			value, err = strconv.ParseInt(rawValue, 10, 64)
			setParamType = "int"
		}
	case Float:
		{
			value, err = strconv.ParseFloat(rawValue, 64)
			setParamType = "float"
		}
	case Bool:
		{
			value, err = strconv.ParseBool(rawValue)
			setParamType = "bool"
		}
	}

	if err != nil {
		parseErr := rest_errors.NewBadRequestError(fmt.Sprintf("%s should be a %s", key, setParamType))
		return nil, parseErr
	}

	return value, nil
}
