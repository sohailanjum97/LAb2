/*
 * NRF NFDiscovery Service
 *
 * NRF NFDiscovery Service. © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.1.0.alpha-4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// UpfInfo - Information of an UPF NF Instance
type UpfInfo struct {

	SNssaiUpfInfoList []SnssaiUpfInfoItem `json:"sNssaiUpfInfoList"`

	SmfServingArea []string `json:"smfServingArea,omitempty"`

	InterfaceUpfInfoList []InterfaceUpfInfoItem `json:"interfaceUpfInfoList,omitempty"`

	IwkEpsInd bool `json:"iwkEpsInd,omitempty"`

	PduSessionTypes []PduSessionType `json:"pduSessionTypes,omitempty"`

	AtsssCapability AtsssCapability `json:"atsssCapability,omitempty"`

	UeIpAddrInd bool `json:"ueIpAddrInd,omitempty"`

	TaiList []Tai `json:"taiList,omitempty"`

	WAgfInfo WAgfInfo `json:"wAgfInfo,omitempty"`

	TngfInfo TngfInfo `json:"tngfInfo,omitempty"`

	TwifInfo TwifInfo `json:"twifInfo,omitempty"`

	Priority int32 `json:"priority,omitempty"`

	RedundantGtpu bool `json:"redundantGtpu,omitempty"`
}