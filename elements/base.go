package elements

import (
	"encoding/json"
	"reflect"
	"strings"
	
	"github.com/Lakeshore-Labs/creatomate-go/properties"
)

type ValueOrKeyframes[T any] interface{}

type ElementBase interface {
	ToMap() map[string]interface{}
}

type ElementProperties struct {
	// A unique identifier for this element.
	ID string `json:"id,omitempty"`

	// The track number on which this element is placed.
	Track *int `json:"track,omitempty"`

	// The time at which you want the element to appear within its composition.
	Time interface{} `json:"time,omitempty"` // number or string

	// The duration for which you would like the element to appear.
	Duration interface{} `json:"duration,omitempty"` // number or string

	// The x-axis position of the element in the composition.
	X ValueOrKeyframes[interface{}] `json:"x,omitempty"`

	// The y-axis position of the element in the composition.
	Y ValueOrKeyframes[interface{}] `json:"y,omitempty"`

	// The width of the element in relation to the composition.
	Width ValueOrKeyframes[interface{}] `json:"width,omitempty"`

	// The height of the element in relation to the composition.
	Height ValueOrKeyframes[interface{}] `json:"height,omitempty"`

	// Using this property, the element will be constrained to a particular aspect ratio.
	AspectRatio ValueOrKeyframes[float64] `json:"aspect_ratio,omitempty"`

	// Padding of the element on the horizontal axis.
	XPadding ValueOrKeyframes[interface{}] `json:"x_padding,omitempty"`

	// Padding of the element on the vertical axis.
	YPadding ValueOrKeyframes[interface{}] `json:"y_padding,omitempty"`

	// The order in which the elements are rendered.
	ZIndex ValueOrKeyframes[int] `json:"z_index,omitempty"`

	// The element's origin from which its x-axis position, scale, rotate, and skew are applied.
	XAnchor ValueOrKeyframes[interface{}] `json:"x_anchor,omitempty"`

	// The element's origin from which its y-axis position, scale, rotate, and skew are applied.
	YAnchor ValueOrKeyframes[interface{}] `json:"y_anchor,omitempty"`

	// The horizontal scale transformation in percent.
	XScale ValueOrKeyframes[interface{}] `json:"x_scale,omitempty"`

	// The vertical scale transformation in percent.
	YScale ValueOrKeyframes[interface{}] `json:"y_scale,omitempty"`

	// The horizontal skew transformation in degrees.
	XSkew ValueOrKeyframes[interface{}] `json:"x_skew,omitempty"`

	// The vertical skew transformation in degrees.
	YSkew ValueOrKeyframes[interface{}] `json:"y_skew,omitempty"`

	// Rotates the element along the x-axis.
	XRotation ValueOrKeyframes[interface{}] `json:"x_rotation,omitempty"`

	// Rotates the element along the y-axis.
	YRotation ValueOrKeyframes[interface{}] `json:"y_rotation,omitempty"`

	// Rotates the element along the z-axis.
	ZRotation ValueOrKeyframes[interface{}] `json:"z_rotation,omitempty"`

	// The distance between the z=0 plane and the camera.
	Perspective ValueOrKeyframes[interface{}] `json:"perspective,omitempty"`

	// Set to false to hide the backface of the element when rotated.
	BackfaceVisible ValueOrKeyframes[bool] `json:"backface_visible,omitempty"`

	// The position of the element's content on the x-axis.
	XAlignment ValueOrKeyframes[interface{}] `json:"x_alignment,omitempty"`

	// The position of the element's content on the y-axis.
	YAlignment ValueOrKeyframes[interface{}] `json:"y_alignment,omitempty"`

	// The shadow.
	Shadow ValueOrKeyframes[*properties.Shadow] `json:"shadow,omitempty"`

	// The shadow color, or null to disable it.
	ShadowColor ValueOrKeyframes[string] `json:"shadow_color,omitempty"`

	// The blurriness of the shadow.
	ShadowBlur ValueOrKeyframes[interface{}] `json:"shadow_blur,omitempty"`

	// The offset of the shadow on the x-axis.
	ShadowX ValueOrKeyframes[interface{}] `json:"shadow_x,omitempty"`

	// The offset of the shadow on the y-axis.
	ShadowY ValueOrKeyframes[interface{}] `json:"shadow_y,omitempty"`

	// When set to true, the element's content is clipped to its borders.
	Clip ValueOrKeyframes[bool] `json:"clip,omitempty"`

	// The opacity of the element.
	Opacity ValueOrKeyframes[interface{}] `json:"opacity,omitempty"`

	// The blend mode of the element.
	BlendMode ValueOrKeyframes[properties.BlendMode] `json:"blend_mode,omitempty"`

	// The color filter that is applied to the element.
	ColorFilter ValueOrKeyframes[interface{}] `json:"color_filter,omitempty"`

	// This parameter allows you to control the color filter, such as the intensity.
	ColorFilterValue ValueOrKeyframes[float64] `json:"color_filter_value,omitempty"`

	// A color that is applied on top the element.
	ColorOverlay ValueOrKeyframes[string] `json:"color_overlay,omitempty"`

	// The blur.
	Blur ValueOrKeyframes[*properties.Blur] `json:"blur,omitempty"`

	// The radius of the blur that is applied to the element.
	BlurRadius ValueOrKeyframes[float64] `json:"blur_radius,omitempty"`

	// The algorithm used to blur the element.
	BlurMode ValueOrKeyframes[properties.BlurMode] `json:"blur_mode,omitempty"`

	// By setting the mask mode, the element is used as a mask.
	MaskMode ValueOrKeyframes[properties.MaskMode] `json:"mask_mode,omitempty"`

	// When set to true, the element is repeated in its composition.
	Repeat ValueOrKeyframes[bool] `json:"repeat,omitempty"`

	// The warp.
	Warp ValueOrKeyframes[*properties.Warp] `json:"warp,omitempty"`

	// This parameter is used in conjunction with warp_matrix.
	WarpMode ValueOrKeyframes[properties.WarpMode] `json:"warp_mode,omitempty"`

	// Array of points that control the warp effect.
	WarpMatrix ValueOrKeyframes[[][]properties.WarpPoint] `json:"warp_matrix,omitempty"`

	// An animation used as transition between this and the previous element.
	Transition interface{} `json:"transition,omitempty"`

	// An animation that is played at the start.
	Enter interface{} `json:"enter,omitempty"`

	// An animation that is played at the end.
	Exit interface{} `json:"exit,omitempty"`

	// An array of animation keyframes.
	Animations []interface{} `json:"animations,omitempty"`
}

type BaseElement struct {
	Type       string
	Properties interface{}
}

func (e *BaseElement) ToMap() map[string]interface{} {
	// First, manually convert properties to handle special types
	var result map[string]interface{}
	
	// Use reflection to convert struct to map while preserving special types
	if props := structToMapWithExpansion(e.Properties); props != nil {
		result = props
	} else {
		// Fallback to JSON method
		propsJSON, err := json.Marshal(e.Properties)
		if err != nil {
			return map[string]interface{}{"type": e.Type}
		}
		
		if err := json.Unmarshal(propsJSON, &result); err != nil {
			return map[string]interface{}{"type": e.Type}
		}
	}
	
	// Handle animations (enter, exit, transition)
	if animations, ok := result["animations"]; ok {
		if animArray, ok := animations.([]interface{}); ok {
			var expandedAnimations []interface{}
			for _, anim := range animArray {
				if animWithMap, ok := anim.(interface{ ToMap() map[string]interface{} }); ok {
					expandedAnimations = append(expandedAnimations, animWithMap.ToMap())
				} else {
					expandedAnimations = append(expandedAnimations, anim)
				}
			}
			result["animations"] = expandedAnimations
		}
	}
	
	// Handle enter animation
	if enter, ok := result["enter"]; ok {
		if enterWithMap, ok := enter.(interface{ ToMap() map[string]interface{} }); ok {
			enterMap := enterWithMap.ToMap()
			enterMap["time"] = "start"
			if animArray, ok := result["animations"].([]interface{}); ok {
				result["animations"] = append([]interface{}{enterMap}, animArray...)
			} else {
				result["animations"] = []interface{}{enterMap}
			}
			delete(result, "enter")
		}
	}
	
	// Handle exit animation
	if exit, ok := result["exit"]; ok {
		if exitWithMap, ok := exit.(interface{ ToMap() map[string]interface{} }); ok {
			exitMap := exitWithMap.ToMap()
			exitMap["time"] = "end"
			exitMap["reversed"] = true
			if animArray, ok := result["animations"].([]interface{}); ok {
				result["animations"] = append(animArray, exitMap)
			} else {
				result["animations"] = []interface{}{exitMap}
			}
			delete(result, "exit")
		}
	}
	
	// Handle transition animation
	if transition, ok := result["transition"]; ok {
		if transWithMap, ok := transition.(interface{ ToMap() map[string]interface{} }); ok {
			transMap := transWithMap.ToMap()
			transMap["time"] = "start"
			transMap["transition"] = true
			if animArray, ok := result["animations"].([]interface{}); ok {
				result["animations"] = append([]interface{}{transMap}, animArray...)
			} else {
				result["animations"] = []interface{}{transMap}
			}
			delete(result, "transition")
		}
	}
	
	// Remove nil values
	for k, v := range result {
		if v == nil {
			delete(result, k)
		}
	}
	
	result["type"] = e.Type
	return result
}

// expandProperties expands properties that have ToMap method
func expandProperties(properties map[string]interface{}) map[string]interface{} {
	expanded := make(map[string]interface{})
	
	for key, value := range properties {
		if value == nil {
			continue
		}
		
		// Check if value has ToMap method
		if valueWithMap, ok := value.(interface{ ToMap() map[string]interface{} }); ok {
			// Expand the property
			nestedProps := valueWithMap.ToMap()
			for nestedKey, nestedValue := range nestedProps {
				expanded[nestedKey] = nestedValue
			}
		} else {
			// Keep as is
			expanded[key] = value
		}
	}
	
	return expanded
}

// structToMapWithExpansion converts a struct to map using reflection and expands ToMap types
func structToMapWithExpansion(obj interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	
	if v.Kind() != reflect.Struct {
		return nil
	}
	
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		fieldValue := v.Field(i)
		
		// Skip unexported fields
		if !fieldValue.CanInterface() {
			continue
		}
		
		// Get JSON tag
		jsonTag := field.Tag.Get("json")
		if jsonTag == "-" {
			continue
		}
		
		tagParts := strings.Split(jsonTag, ",")
		fieldName := tagParts[0]
		if fieldName == "" {
			fieldName = field.Name
		}
		
		// Skip if omitempty and value is zero
		if len(tagParts) > 1 && tagParts[1] == "omitempty" && isZeroValue(fieldValue) {
			continue
		}
		
		value := fieldValue.Interface()
		
		// Handle embedded structs
		if field.Anonymous && fieldValue.Kind() == reflect.Struct {
			if embedded := structToMapWithExpansion(value); embedded != nil {
				for k, v := range embedded {
					result[k] = v
				}
			}
			continue
		}
		
		// Check if value has ToMap method and should be expanded
		if shouldExpand(fieldName) {
			if valueWithMap, ok := value.(interface{ ToMap() map[string]interface{} }); ok {
				// Expand the property
				nestedProps := valueWithMap.ToMap()
				for nestedKey, nestedValue := range nestedProps {
					result[nestedKey] = nestedValue
				}
				continue
			}
		}
		
		result[fieldName] = value
	}
	
	return result
}

// shouldExpand returns true if the field should be expanded
func shouldExpand(fieldName string) bool {
	// Expand font and background properties
	return fieldName == "font" || fieldName == "background" || fieldName == "text_background"
}

// isZeroValue checks if a reflect.Value is a zero value
func isZeroValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}