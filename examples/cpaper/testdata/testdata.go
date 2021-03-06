package testdata

import (
	"time"

	"github.com/soodakshay/cckit/testing"

	"github.com/golang/protobuf/ptypes"
	"github.com/soodakshay/cckit/examples/cpaper/schema"
)

var (
	// CPapers commercial paper fixtures
	CPapers = []schema.IssueCommercialPaper{{
		Issuer:       `some-issuer-1`,
		PaperNumber:  `00000001`,
		IssueDate:    ptypes.TimestampNow(),
		MaturityDate: testing.MustProtoTimestamp(time.Now().AddDate(0, 1, 0)),
		FaceValue:    11111,
	}, {
		Issuer:       `some-issuer-2`,
		PaperNumber:  `00000002`,
		IssueDate:    ptypes.TimestampNow(),
		MaturityDate: testing.MustProtoTimestamp(time.Now().AddDate(0, 2, 0)),
		FaceValue:    22222,
	}, {
		Issuer:       `some-issuer-3`,
		PaperNumber:  `00000003`,
		IssueDate:    ptypes.TimestampNow(),
		MaturityDate: testing.MustProtoTimestamp(time.Now().AddDate(0, 3, 0)),
	}}
)
