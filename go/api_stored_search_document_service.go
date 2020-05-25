/*
 * NRF NFDiscovery Service
 *
 * NRF NFDiscovery Service. © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.1.0.alpha-4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"errors"
)

// StoredSearchDocumentApiService is a service that implents the logic for the StoredSearchDocumentApiServicer
// This service should implement the business logic for every endpoint for the StoredSearchDocumentApi API. 
// Include any external packages or services that will be required by this service.
type StoredSearchDocumentApiService struct {
}

// NewStoredSearchDocumentApiService creates a default api service
func NewStoredSearchDocumentApiService() StoredSearchDocumentApiServicer {
	return &StoredSearchDocumentApiService{}
}

// RetrieveStoredSearch - 
func (s *StoredSearchDocumentApiService) RetrieveStoredSearch(searchId string, acceptEncoding string) (interface{}, error) {
	// TODO - update RetrieveStoredSearch with the required logic for this service method.
	// Add api_stored_search_document_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'RetrieveStoredSearch' not implemented")
}