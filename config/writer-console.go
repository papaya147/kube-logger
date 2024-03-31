package config

import (
	"context"
	"fmt"

	"github.com/fatih/color"
)

type ConsoleWriter struct {
}

func NewConsoleWriter() Writer {
	return &ConsoleWriter{}
}

func (c *ConsoleWriter) Open(context.Context, string) error {
	return nil
}

var ColorMap map[string]color.Attribute = make(map[string]color.Attribute)

var lastColor color.Attribute

func nextColor() color.Attribute {
	var nextColor color.Attribute
	switch lastColor {
	case color.FgRed:
		nextColor = color.FgMagenta
	case color.FgMagenta:
		nextColor = color.FgYellow
	case color.FgYellow:
		nextColor = color.FgBlue
	case color.FgBlue:
		nextColor = color.FgCyan
	default:
		nextColor = color.FgRed
	}
	lastColor = nextColor
	return nextColor
}

func (c *ConsoleWriter) Write(namespace, pod string, data []byte) error {
	_, ok := ColorMap[namespace+pod]
	if !ok {
		ColorMap[namespace+pod] = nextColor()
	}

	color.Set(ColorMap[namespace+pod])
	fmt.Printf("%s: %s\n", namespace, pod)
	color.Unset()

	color.Set(color.FgWhite)
	fmt.Println(string(data))
	color.Unset()

	return nil
}

func (c *ConsoleWriter) Close(context.Context) error {
	return nil
}
