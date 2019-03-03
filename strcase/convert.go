package strcase

import (
	"regexp"
	"strings"
)

type CaseType uint8

const (
	CamelType  = iota
	SnakeType
	PascalType
	HyphenType
)

var StringToCaseType = map[string]CaseType{
	"camel":  CamelType,
	"snake":  SnakeType,
	"pascal": PascalType,
	"hyphen": HyphenType,
}

func Lcfirst(str string) string {
	return strings.ToLower(str[0:1]) + str[1:]
}

func Ucfirst(str string) string {
	return strings.ToUpper(str[0:1]) + str[1:]
}

func Camel(str string) string {
	return Lcfirst(camelRegexp.ReplaceAllStringFunc(str, func(str string) string {
		return Ucfirst(str[1:])
	}))
}

func Pascal(str string) string {
	return Ucfirst(Camel(str))
}

func Snake(str string) string {
	return upperCaseRe.ReplaceAllStringFunc(Camel(str), func(str string) string {
		return "_" + Lcfirst(str)
	})
}

func Hyphen(str string) string {
	return upperCaseRe.ReplaceAllStringFunc(Camel(str), func(str string) string {
		return "-" + Lcfirst(str)
	})
}

func Convert(str string, caseType CaseType) string {
	switch caseType {
	case CamelType:
		return Camel(str)
	case SnakeType:
		return Snake(str)
	case PascalType:
		return Pascal(str)
	case HyphenType:
		return Hyphen(str)
	default:
		return str
	}
}

func ConvertStringType(str string, caseTypeStr string) string {
	_, ok := StringToCaseType[caseTypeStr]
	if ok {
		return Convert(str, StringToCaseType[caseTypeStr])
	}
	return str
}

var camelRegexp = regexp.MustCompile(`[-_](\w)`)
var upperCaseRe = regexp.MustCompile(`([A-Z])`)
