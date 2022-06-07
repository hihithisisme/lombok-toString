package lombokString

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:  "Simple Class",
			input: "Attendance(venue=Jewel Airport, headcount=1023)",
			expected: []string{
				"Attendance",
				"(venue=Jewel Airport, headcount=1023)",
			},
		},
		{
			name:  "Simple Class with External Brackets",
			input: "[Attendance(venue=Jewel Airport, headcount=1023)]",
			expected: []string{
				"[Attendance",
				"(venue=Jewel Airport, headcount=1023)",
				"]",
			},
		},
		{
			name:  "Simple Numeric Array",
			input: "[1, 1, 2, 3, 5, 8, 13]",
			expected: []string{
				"[1, 1, 2, 3, 5, 8, 13]",
			},
		},
		{
			name:  "Class within Array",
			input: "[Attendee(name=Taylor Hebert, gender=female), Attendee(name=Blake Thornburn, gender=other)]",
			expected: []string{
				"[Attendee",
				"(name=Taylor Hebert, gender=female),",
				"Attendee",
				"(name=Blake Thornburn, gender=other)",
				"]",
			},
		},
		{
			name:  "Array within Class",
			input: "Attendance(venue=Jewel Airport, headcount=1023, attendees=[Attendee(name=Taylor Hebert, gender=female), Attendee(name=Blake Thornburn, gender=other)])",
			expected: []string{
				"Attendance",
				"(venue=Jewel Airport, headcount=1023, attendees=",
				"[Attendee",
				"(name=Taylor Hebert, gender=female),",
				"Attendee",
				"(name=Blake Thornburn, gender=other)",
				"]",
				")",
			},
		},
		{
			name:  "Nested Class",
			input: "ShippingLogUpdate(attempt=ShippingAttempt(id=182501, session=OrderSesseon(id=158171, address=Address(line1=City Square, postCode=S(112300)), createdAt=2022-03-04T09:08:21.052Z, quantity=4, shipperDescription=null, updatedAt=2022-03-04 03:08:21.055), paymentMethod=INHOUSE), transactionStatusEvent=TransactionStatusEvent(data=TransactionStatusEvent.TransactionStatus(transactionRef=ba9d5206-28de-43ec-bd6f-dadfcb5ef890-158171-182501, status=SUCCESS, orderRef=e4223595/829f/4b2d/be49/a3fb4f9b8044, encoding=encode=YXNkZndlcndlcndlcg==, transactionDescription=null)))",
			expected: []string{
				"ShippingLogUpdate",
				"(attempt=ShippingAttempt",
				"(id=182501, session=OrderSesseon",
				"(id=158171, address=Address",
				"(line1=City Square, postCode=S",
				"(112300)",
				"),",
				"createdAt=2022-03-04T09:08:21.052Z, quantity=4, shipperDescription=null, updatedAt=2022-03-04 03:08:21.055),",
				"paymentMethod=INHOUSE),",
				"transactionStatusEvent=TransactionStatusEvent",
				"(data=TransactionStatusEvent.TransactionStatus",
				"(transactionRef=ba9d5206-28de-43ec-bd6f-dadfcb5ef890-158171-182501, status=SUCCESS, orderRef=e4223595/829f/4b2d/be49/a3fb4f9b8044, encoding=encode=YXNkZndlcndlcndlcg==, transactionDescription=null)",
				")",
				")",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, SegmentByBrackets(test.input))
		})
	}
}
