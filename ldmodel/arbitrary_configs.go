package ldmodel

// ArbitraryConfigs represents a configuration that can store arbitrary values
// with their associated metadata like type and version
type ArbitraryConfigs struct {
	// Key uniquely identifies the configuration
	Key string `json:"key" bson:"key"`

	// DataType indicates the type of data stored in Values
	DataType string `json:"dataType" bson:"dataType"`

	// Values stores the arbitrary configuration data
	Values interface{} `json:"values" bson:"values"`

	// Version represents the current version of this configuration
	Version int64 `json:"version" bson:"version"`
}
