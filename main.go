package disposableemail

import (
	_ "embed"
	"encoding/csv"
	"fmt"
	"strings"
)

type CheckResult struct {
	IsDisposable bool `json:"isDisposable"`
}

type NewParams struct {
	Domains []string `json:"domains"`
}

type Service interface {
	Check(emailOrDomain string) *CheckResult
}

type service struct {
	domainsMap map[string]bool
}

// New godoc
// New creates a new disposable e-mail checking service.
func New() (Service, error) {
	domainsMap := map[string]bool{}
	csvReader := csv.NewReader(strings.NewReader(domainsCsvFile))
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read domains.csv: %w", err)
	}

	for i := range records {
		domainsMap[records[i][0]] = true
	}

	return &service{
		domainsMap: domainsMap,
	}, nil
}

// Check godoc
// Check checks if the given e-mail address or domain is disposable.
func (s *service) Check(emailOrDomain string) *CheckResult {
	domain := emailOrDomain

	atIndex := strings.LastIndex(emailOrDomain, "@")
	if atIndex > -1 {
		domain = emailOrDomain[atIndex+1:]
	}

	return &CheckResult{
		IsDisposable: s.domainsMap[domain],
	}
}

//go:embed domains.csv
var domainsCsvFile string
