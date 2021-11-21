package save

type Component struct {
	TypePath         string      `json:"typePath"`
	RootObject       string      `json:"rootObject"`
	InstanceName     string      `json:"instanceName"`
	ParentEntityName string      `json:"parentEntityName"`
	Properties       []*Property `json:"properties"`

	objectDataPos int64
	objectData    []byte
}

func (c *Component) parse(p *parser) error {
	var err error
	c.TypePath, err = p.ReadString()
	if err != nil {
		return err
	}

	c.RootObject, err = p.ReadString()
	if err != nil {
		return err
	}

	c.InstanceName, err = p.ReadString()
	if err != nil {
		return err
	}

	c.ParentEntityName, err = p.ReadString()
	if err != nil {
		return err
	}

	return nil
}

func (c *Component) parseData(p *parser) error {
	// Data length
	//dataSize, err := p.readInt32()
	_, err := p.ReadInt32()
	if err != nil {
		return err
	}

	// Capture all bytes if debug is enabled.
	//if debug {
	//	c.objectDataPos = p.body.Index - 4
	//	c.objectData, err = p.debugReadDataChunk(dataSize)
	//}

	c.Properties, err = p.parseProperties()
	if err != nil {
		return err
	}

	// UNKNOWN_DATA
	_, err = p.ReadInt32()
	if err != nil {
		return err
	}

	return nil
}

func (c *Component) serialize(s *serializer) error {
	err := s.WriteString(c.TypePath)
	if err != nil {
		return err
	}

	err = s.WriteString(c.RootObject)
	if err != nil {
		return err
	}

	err = s.WriteString(c.InstanceName)
	if err != nil {
		return err
	}

	err = s.WriteString(c.ParentEntityName)
	if err != nil {
		return err
	}

	return nil
}

func (c *Component) serializeData(s *serializer) error {
	// Write placeholder length
	lenPos := s.Index()
	err := s.WriteInt32(0)
	if err != nil {
		return err
	}

	m := s.Measure()

	err = serializeProperties(c.Properties, s.Data)
	if err != nil {
		return err
	}

	// UNKNOWN_DATA
	err = s.WriteInt32(0)
	if err != nil {
		return err
	}

	err = s.WriteLen(m(), lenPos)
	if err != nil {
		return err
	}

	return nil
}
