package disposable_emails

import (
	_ "embed"
	"encoding/csv"
	"fmt"
	"strings"
)

type CheckResult struct {
	IsDisposable bool `json:"isDisposable"`
}

type Service interface {
	Check(emailOrDomain string) *CheckResult
}

type service struct {
	domainsMap map[string]bool
}

func New() (Service, error) {
	csvReader := csv.NewReader(strings.NewReader(domainsCsvFile))
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read domains.csv: %w", err)
	}

	domainsMap := map[string]bool{}
	for i := range records {
		domainsMap[records[i][0]] = true
	}

	return &service{
		domainsMap: domainsMap,
	}, nil
}

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
