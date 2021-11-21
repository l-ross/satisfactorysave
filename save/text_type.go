package save

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/save/data"
)

type TextType byte

const (
	NoneTextType TextType = 0xFF
	BaseTextType TextType = 0x00
)

type TextValue interface {
	parse(d *data.Data) error

	serialize(d *data.Data) error
}

func getText(t TextType) func() TextValue {
	return nil
}

//
// None
//

type NoneText struct {
	CultureInvariantString int32
	String                 string
}

func (t *TextPropertyValue) GetNoneText() (*NoneText, error) {
	if v, ok := t.Value.(*NoneText); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong text type: %v", t.Type)
}

func (t *NoneText) parse(d *data.Data) error {
	var err error
	t.CultureInvariantString, err = d.ReadInt32()
	if err != nil {
		return err
	}

	// TODO: Switch on above?
	t.String, err = d.ReadString()
	if err != nil {
		return err
	}

	return nil
}

func (t *NoneText) serialize(d *data.Data) error {
	err := d.WriteInt32(t.CultureInvariantString)
	if err != nil {
		return err
	}

	// TODO: Switch on above?
	err = d.WriteString(t.String)
	if err != nil {
		return err
	}

	return nil
}

//
// Base
//

type BaseText struct {
	Namespace string
	Key       string
	Value     string
}

func (t *TextPropertyValue) GetBaseText() (*BaseText, error) {
	if v, ok := t.Value.(*BaseText); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong text type: %v", t.Type)
}

func (t *BaseText) parse(d *data.Data) error {
	var err error
	t.Namespace, err = d.ReadString()
	if err != nil {
		return err
	}

	t.Key, err = d.ReadString()
	if err != nil {
		return err
	}

	t.Value, err = d.ReadString()
	if err != nil {
		return err
	}

	return nil
}

func (t *BaseText) serialize(d *data.Data) error {
	err := d.WriteString(t.Namespace)
	if err != nil {
		return err
	}

	err = d.WriteString(t.Key)
	if err != nil {
		return err
	}

	err = d.WriteString(t.Value)
	if err != nil {
		return err
	}

	return nil
}
