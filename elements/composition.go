package elements

type CompositionProperties struct {
	ElementProperties

	// Identical to duration from the common properties, but can be set to "composition".
	Duration interface{} `json:"duration,omitempty"` // number, string, or "composition"

	// Elements that make up the composition.
	Elements []interface{} `json:"elements,omitempty"`
}

type Composition struct {
	BaseElement
}

// Override ToMap to handle nested elements
func (c *Composition) ToMap() map[string]interface{} {
	result := c.BaseElement.ToMap()
	
	// Handle nested elements
	if props, ok := c.Properties.(CompositionProperties); ok {
		if props.Elements != nil {
			elements := make([]interface{}, len(props.Elements))
			for i, elem := range props.Elements {
				if elemWithMap, ok := elem.(interface{ ToMap() map[string]interface{} }); ok {
					elements[i] = elemWithMap.ToMap()
				} else {
					elements[i] = elem
				}
			}
			result["elements"] = elements
		}
	}
	
	return result
}

func NewComposition(props CompositionProperties) *Composition {
	return &Composition{
		BaseElement: BaseElement{
			Type:       "composition",
			Properties: props,
		},
	}
}