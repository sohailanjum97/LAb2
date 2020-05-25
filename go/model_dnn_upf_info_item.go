/*
 * NRF NFDiscovery Service
 *
 * NRF NFDiscovery Service. © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.1.0.alpha-4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// DnnUpfInfoItem - Set of parameters supported by UPF for a given DNN
type DnnUpfInfoItem struct {

	Dnn string `json:"dnn"`

	DnaiList []string `json:"dnaiList,omitempty"`

	PduSessionTypes []PduSessionType `json:"pduSessionTypes,omitempty"`

	Ipv4AddressRanges []Ipv4AddressRange `json:"ipv4AddressRanges,omitempty"`

	Ipv6PrefixRanges []Ipv6PrefixRange `json:"ipv6PrefixRanges,omitempty"`
}