// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LogreplayV1ReplayStats logreplay v1 replay stats
//
// swagger:model logreplay.v1.ReplayStats
type LogreplayV1ReplayStats struct {

	// number of invalid decisions that could not be replayed (<= entries_evaluated)
	// Required: true
	AnalysisErrors *int64 `json:"analysis_errors"`

	// number of fully analyzed decision batches (<= batches_observed)
	// Required: true
	BatchesAnalyzed *int64 `json:"batches_analyzed"`

	// number of decision batches that could not be downloaded
	// Required: true
	BatchesDownloadErrors *int64 `json:"batches_download_errors"`

	// number of downloaded decision batches (<= batches_scheduled)
	// Required: true
	BatchesDownloaded *int64 `json:"batches_downloaded"`

	// number of downloaded decision batches taken from the cache (<= batches_downloaded)
	// Required: true
	BatchesFromCache *int64 `json:"batches_from_cache"`

	// number of decision batches picked by the analyzers (<= batches_downloaded)
	// Required: true
	BatchesObserved *int64 `json:"batches_observed"`

	// number of decision batches scheduled for download
	// Required: true
	BatchesScheduled *int64 `json:"batches_scheduled"`

	// number of decision batches skipped due to the skip_batches input or pre-filtration (timestamp range, system/stack IDs)
	// Required: true
	BatchesSkipped *int64 `json:"batches_skipped"`

	// number of replayed decisions (<= entries_scheduled)
	// Required: true
	EntriesEvaluated *int64 `json:"entries_evaluated"`

	// number of decisions evaluated to error (<= entries_evaluated)
	// Required: true
	EntriesFailed *int64 `json:"entries_failed"`

	// number of decisions in batches_observed batches
	// Required: true
	EntriesObserved *int64 `json:"entries_observed"`

	// number of decisions scheduled for replay (<= entries_observed)
	// Required: true
	EntriesScheduled *int64 `json:"entries_scheduled"`

	// number of observed decision changes (<= entries_evaluated)
	// Required: true
	ResultsChanged *int64 `json:"results_changed"`
}

// Validate validates this logreplay v1 replay stats
func (m *LogreplayV1ReplayStats) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAnalysisErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBatchesAnalyzed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBatchesDownloadErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBatchesDownloaded(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBatchesFromCache(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBatchesObserved(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBatchesScheduled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBatchesSkipped(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntriesEvaluated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntriesFailed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntriesObserved(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntriesScheduled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResultsChanged(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogreplayV1ReplayStats) validateAnalysisErrors(formats strfmt.Registry) error {

	if err := validate.Required("analysis_errors", "body", m.AnalysisErrors); err != nil {
		return err
	}

	return nil
}

func (m *LogreplayV1ReplayStats) validateBatchesAnalyzed(formats strfmt.Registry) error {

	if err := validate.Required("batches_analyzed", "body", m.BatchesAnalyzed); err != nil {
		return err
	}

	return nil
}

func (m *LogreplayV1ReplayStats) validateBatchesDownloadErrors(formats strfmt.Registry) error {

	if err := validate.Required("batches_download_errors", "body", m.BatchesDownloadErrors); err != nil {
		return err
	}

	return nil
}

func (m *LogreplayV1ReplayStats) validateBatchesDownloaded(formats strfmt.Registry) error {

	if err := validate.Required("batches_downloaded", "body", m.BatchesDownloaded); err != nil {
		return err
	}

	return nil
}

func (m *LogreplayV1ReplayStats) validateBatchesFromCache(formats strfmt.Registry) error {

	if err := validate.Required("batches_from_cache", "body", m.BatchesFromCache); err != nil {
		return err
	}

	return nil
}

func (m *LogreplayV1ReplayStats) validateBatchesObserved(formats strfmt.Registry) error {

	if err := validate.Required("batches_observed", "body", m.BatchesObserved); err != nil {
		return err
	}

	return nil
}

func (m *LogreplayV1ReplayStats) validateBatchesScheduled(formats strfmt.Registry) error {

	if err := validate.Required("batches_scheduled", "body", m.BatchesScheduled); err != nil {
		return err
	}

	return nil
}

func (m *LogreplayV1ReplayStats) validateBatchesSkipped(formats strfmt.Registry) error {

	if err := validate.Required("batches_skipped", "body", m.BatchesSkipped); err != nil {
		return err
	}

	return nil
}

func (m *LogreplayV1ReplayStats) validateEntriesEvaluated(formats strfmt.Registry) error {

	if err := validate.Required("entries_evaluated", "body", m.EntriesEvaluated); err != nil {
		return err
	}

	return nil
}

func (m *LogreplayV1ReplayStats) validateEntriesFailed(formats strfmt.Registry) error {

	if err := validate.Required("entries_failed", "body", m.EntriesFailed); err != nil {
		return err
	}

	return nil
}

func (m *LogreplayV1ReplayStats) validateEntriesObserved(formats strfmt.Registry) error {

	if err := validate.Required("entries_observed", "body", m.EntriesObserved); err != nil {
		return err
	}

	return nil
}

func (m *LogreplayV1ReplayStats) validateEntriesScheduled(formats strfmt.Registry) error {

	if err := validate.Required("entries_scheduled", "body", m.EntriesScheduled); err != nil {
		return err
	}

	return nil
}

func (m *LogreplayV1ReplayStats) validateResultsChanged(formats strfmt.Registry) error {

	if err := validate.Required("results_changed", "body", m.ResultsChanged); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this logreplay v1 replay stats based on context it is used
func (m *LogreplayV1ReplayStats) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LogreplayV1ReplayStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogreplayV1ReplayStats) UnmarshalBinary(b []byte) error {
	var res LogreplayV1ReplayStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
