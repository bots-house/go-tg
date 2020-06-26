// Code generated by go-swagger; DO NOT EDIT.

package catalog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// GetLotsURL generates an URL for the get lots operation
type GetLotsURL struct {
	DailyCoverageFrom  *int64
	DailyCoverageTo    *int64
	MembersCountFrom   *int64
	MembersCountTo     *int64
	MonthlyIncomeFrom  *int64
	MonthlyIncomeTo    *int64
	PaybackPeriodFrom  *float64
	PaybackPeriodTo    *float64
	PriceFrom          *int64
	PricePerMemberFrom *float64
	PricePerMemberTo   *float64
	PricePerViewFrom   *float64
	PricePerViewTo     *float64
	PriceTo            *int64
	SortBy             *string
	SortByType         *string
	Topics             []int64

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetLotsURL) WithBasePath(bp string) *GetLotsURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetLotsURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetLotsURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/lots"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/v1"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var dailyCoverageFromQ string
	if o.DailyCoverageFrom != nil {
		dailyCoverageFromQ = swag.FormatInt64(*o.DailyCoverageFrom)
	}
	if dailyCoverageFromQ != "" {
		qs.Set("daily_coverage_from", dailyCoverageFromQ)
	}

	var dailyCoverageToQ string
	if o.DailyCoverageTo != nil {
		dailyCoverageToQ = swag.FormatInt64(*o.DailyCoverageTo)
	}
	if dailyCoverageToQ != "" {
		qs.Set("daily_coverage_to", dailyCoverageToQ)
	}

	var membersCountFromQ string
	if o.MembersCountFrom != nil {
		membersCountFromQ = swag.FormatInt64(*o.MembersCountFrom)
	}
	if membersCountFromQ != "" {
		qs.Set("members_count_from", membersCountFromQ)
	}

	var membersCountToQ string
	if o.MembersCountTo != nil {
		membersCountToQ = swag.FormatInt64(*o.MembersCountTo)
	}
	if membersCountToQ != "" {
		qs.Set("members_count_to", membersCountToQ)
	}

	var monthlyIncomeFromQ string
	if o.MonthlyIncomeFrom != nil {
		monthlyIncomeFromQ = swag.FormatInt64(*o.MonthlyIncomeFrom)
	}
	if monthlyIncomeFromQ != "" {
		qs.Set("monthly_income_from", monthlyIncomeFromQ)
	}

	var monthlyIncomeToQ string
	if o.MonthlyIncomeTo != nil {
		monthlyIncomeToQ = swag.FormatInt64(*o.MonthlyIncomeTo)
	}
	if monthlyIncomeToQ != "" {
		qs.Set("monthly_income_to", monthlyIncomeToQ)
	}

	var paybackPeriodFromQ string
	if o.PaybackPeriodFrom != nil {
		paybackPeriodFromQ = swag.FormatFloat64(*o.PaybackPeriodFrom)
	}
	if paybackPeriodFromQ != "" {
		qs.Set("payback_period_from", paybackPeriodFromQ)
	}

	var paybackPeriodToQ string
	if o.PaybackPeriodTo != nil {
		paybackPeriodToQ = swag.FormatFloat64(*o.PaybackPeriodTo)
	}
	if paybackPeriodToQ != "" {
		qs.Set("payback_period_to", paybackPeriodToQ)
	}

	var priceFromQ string
	if o.PriceFrom != nil {
		priceFromQ = swag.FormatInt64(*o.PriceFrom)
	}
	if priceFromQ != "" {
		qs.Set("price_from", priceFromQ)
	}

	var pricePerMemberFromQ string
	if o.PricePerMemberFrom != nil {
		pricePerMemberFromQ = swag.FormatFloat64(*o.PricePerMemberFrom)
	}
	if pricePerMemberFromQ != "" {
		qs.Set("price_per_member_from", pricePerMemberFromQ)
	}

	var pricePerMemberToQ string
	if o.PricePerMemberTo != nil {
		pricePerMemberToQ = swag.FormatFloat64(*o.PricePerMemberTo)
	}
	if pricePerMemberToQ != "" {
		qs.Set("price_per_member_to", pricePerMemberToQ)
	}

	var pricePerViewFromQ string
	if o.PricePerViewFrom != nil {
		pricePerViewFromQ = swag.FormatFloat64(*o.PricePerViewFrom)
	}
	if pricePerViewFromQ != "" {
		qs.Set("price_per_view_from", pricePerViewFromQ)
	}

	var pricePerViewToQ string
	if o.PricePerViewTo != nil {
		pricePerViewToQ = swag.FormatFloat64(*o.PricePerViewTo)
	}
	if pricePerViewToQ != "" {
		qs.Set("price_per_view_to", pricePerViewToQ)
	}

	var priceToQ string
	if o.PriceTo != nil {
		priceToQ = swag.FormatInt64(*o.PriceTo)
	}
	if priceToQ != "" {
		qs.Set("price_to", priceToQ)
	}

	var sortByQ string
	if o.SortBy != nil {
		sortByQ = *o.SortBy
	}
	if sortByQ != "" {
		qs.Set("sort_by", sortByQ)
	}

	var sortByTypeQ string
	if o.SortByType != nil {
		sortByTypeQ = *o.SortByType
	}
	if sortByTypeQ != "" {
		qs.Set("sort_by_type", sortByTypeQ)
	}

	var topicsIR []string
	for _, topicsI := range o.Topics {
		topicsIS := swag.FormatInt64(topicsI)
		if topicsIS != "" {
			topicsIR = append(topicsIR, topicsIS)
		}
	}

	topics := swag.JoinByFormat(topicsIR, "")

	if len(topics) > 0 {
		qsv := topics[0]
		if qsv != "" {
			qs.Set("topics", qsv)
		}
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetLotsURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetLotsURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetLotsURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetLotsURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetLotsURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetLotsURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
