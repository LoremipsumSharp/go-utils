package string


import (
"strings"
)

type StringAppender struct {
	stringBuilder strings.Builder
}

func NewStringAppender(str string) *StringAppender {
	stringBuilder := strings.Builder{}
	stringBuilder.WriteString(str)
	appender := &StringAppender{
		stringBuilder: strings.Builder{},
	}
	return appender
}

func (appender *StringAppender) Append(stringToAppend string) error {
	_, err := appender.stringBuilder.WriteString(stringToAppend)
	return err
}
func (appender *StringAppender) AppendIf(condition bool, stringToAppend string) error {
	if condition {
		_, err := appender.stringBuilder.WriteString(stringToAppend)
		return err
	}
	return nil
}

func (appender *StringAppender) AppendNewLine() error {
	_, err := appender.stringBuilder.WriteString("\n")
	return err
}

func (appender *StringAppender) ToString() string {
	return appender.stringBuilder.String()
}
