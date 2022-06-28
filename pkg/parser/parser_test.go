package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSegmentByBrackets(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:  "Simple Class",
			input: "Attendance(venue=Jewel Airport, headcount=1023)",
			expected: []string{
				"Attendance(venue=Jewel Airport, headcount=1023)",
			},
		}, {
			name:  "Simple Nested Class",
			input: "Attendance(address=Address(name=Jewel Airport, postalCode=S(103912)), headcount=1023)",
			expected: []string{
				"Attendance(address=",
				"Address(name=Jewel Airport, postalCode=",
				"S(103912)",
				"),",
				"headcount=1023)",
			},
		}, {
			name:  "Simple Nested Array",
			input: "[Attendee(name=Taylor Hebert, gender=female, abilities=[CapeData(codename=skitter, power=[warlord, controls bugs]), CapeData(codename=weaver, power=[controls butterflies]), CapeData(codename=kherpi, power=[you know])]), Attendee(name=Blake Thornburn, gender=other, abilities=[Diabolism, Shamanism, Glamour])]",
			expected: []string{
				"[",
				"Attendee(name=Taylor Hebert, gender=female, abilities=",
				"[",
				"CapeData(codename=skitter, power=",
				"[warlord, controls bugs]",
				"),",
				"CapeData(codename=weaver, power=",
				"[controls butterflies]",
				"),",
				"CapeData(codename=kherpi, power=",
				"[you know]",
				")",
				"]",
				"),",
				"Attendee(name=Blake Thornburn, gender=other, abilities=",
				"[Diabolism, Shamanism, Glamour]",
				")",
				"]",
			},
		}, {
			name:  "Simple Class with External Brackets",
			input: "[Attendance(venue=Jewel Airport, headcount=1023)]",
			expected: []string{
				"[",
				"Attendance(venue=Jewel Airport, headcount=1023)",
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
				"[",
				"Attendee(name=Taylor Hebert, gender=female),",
				"Attendee(name=Blake Thornburn, gender=other)",
				"]",
			},
		},
		{
			name:  "Array within Class",
			input: "Attendance(venue=Jewel Airport, headcount=1023, attendees=[Attendee(name=Taylor Hebert, gender=female), Attendee(name=Blake Thornburn, gender=other)])",
			expected: []string{
				"Attendance(venue=Jewel Airport, headcount=1023, attendees=",
				"[",
				"Attendee(name=Taylor Hebert, gender=female),",
				"Attendee(name=Blake Thornburn, gender=other)",
				"]",
				")",
			},
		},
		{
			name:  "Nested Class",
			input: "ShippingLogUpdate(attempt=ShippingAttempt(id=182501, session=OrderSession(id=158171, address=Address(line1=City Square, postCode=S(112300)), createdAt=2022-03-04T09:08:21.052Z, quantity=4, shipperDescription=null, updatedAt=2022-03-04 03:08:21.055, billingAddress=BillingAddress(line1=Sentosa, postCode=S(612300)), paymentMethod=INHOUSE), transactionStatusEvent=TransactionStatusEvent(data=TransactionStatusEvent.TransactionStatus(transactionRef=ba9d5206-28de-43ec-bd6f-dadfcb5ef890, status=SUCCESS, orderRef=e4223595/829f/4b2d/be49/a3fb4f9b8044, encoding=encode=YXNkZndlcndlcndlcg==, transactionDescription=null))))",
			expected: []string{
				"ShippingLogUpdate(attempt=",
				"ShippingAttempt(id=182501, session=",
				"OrderSession(id=158171, address=",
				"Address(line1=City Square, postCode=",
				"S(112300)",
				"),",
				"createdAt=2022-03-04T09:08:21.052Z, quantity=4, shipperDescription=null, updatedAt=2022-03-04 03:08:21.055, billingAddress=",
				"BillingAddress(line1=Sentosa, postCode=",
				"S(612300)",
				"),",
				"paymentMethod=INHOUSE),",
				"transactionStatusEvent=",
				"TransactionStatusEvent(data=",
				"TransactionStatusEvent.TransactionStatus(transactionRef=ba9d5206-28de-43ec-bd6f-dadfcb5ef890, status=SUCCESS, orderRef=e4223595/829f/4b2d/be49/a3fb4f9b8044, encoding=encode=YXNkZndlcndlcndlcg==, transactionDescription=null)",
				")",
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

func Test_isObject(t *testing.T) {
	tests := []struct {
		name      string
		input     []string
		isClass   []bool
		isArray   []bool
		isInClass []bool
		isInArray []bool
	}{
		{
			name: "Class within Array",
			input: []string{
				"[",
				"Attendee(name=Taylor Hebert, gender=female),",
				"Attendee(name=Blake Thornburn, gender=other)",
				"]",
			},
			isClass: []bool{false, true, true, false},
			isArray: []bool{true, false, false, true},
			//isInClass: []bool{false, false, false, false},
			//isInArray: []bool{false, true, true, false},
		}, {
			name: "Array within Class",
			input: []string{
				"Attendance(venue=Jewel Airport, headcount=1023, attendees=",
				"[",
				"Attendee(name=Taylor Hebert, gender=female),",
				"Attendee(name=Blake Thornburn, gender=other)",
				"]",
				")",
			},
			isClass: []bool{true, false, true, true, false, true},
			isArray: []bool{false, true, false, false, true, false},
			//isInClass: nil,
			//isInArray: nil,
		}, {
			name: "Nested Class",
			input: []string{
				"ShippingLogUpdate(attempt=",
				"ShippingAttempt(id=182501, session=",
				"OrderSession(id=158171, address=",
				"Address(line1=City Square, postCode=",
				"S(112300)",
				"),",
				"createdAt=2022-03-04T09:08:21.052Z, quantity=4, shipperDescription=null, updatedAt=2022-03-04 03:08:21.055, billingAddress=",
				"BillingAddress(line1=Sentosa, postCode=",
				"S(612300)",
				"),",
				"paymentMethod=INHOUSE),",
				"transactionStatusEvent=",
				"TransactionStatusEvent(data=",
				"TransactionStatusEvent.TransactionStatus(transactionRef=ba9d5206-28de-43ec-bd6f-dadfcb5ef890, status=SUCCESS, orderRef=e4223595/829f/4b2d/be49/a3fb4f9b8044, encoding=encode=YXNkZndlcndlcndlcg==, transactionDescription=null)",
				")",
				")",
				")",
			},
			isClass: []bool{true, true, true, true, false, true, false, true, false, true, true, false, true, true, true, true, true},
			isArray: []bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
			//isInClass: []bool{false, true, true, true, true, true, true, true, true, true, true, true, false},
			//isInArray: []bool{false, false, false, false, false, false, false, false, false, false, false, false, false},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for i := range test.input {
				assert.Equalf(t, test.isClass[i], isClass(test.input, i), test.input[i])
				assert.Equalf(t, test.isArray[i], isArray(test.input, i), test.input[i])
				// TODO: determine if I really need isInClass and isInArray
				//assert.Equalf(t, test.isInClass[i], isInClass(test.input, i), test.input[i])
				//assert.Equalf(t, test.isInArray[i], isInArray(test.input, i), test.input[i])
			}
		})
	}
}

func Test_Process(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected interface{}
	}{
		{
			name: "Class within Array",
			input: []string{
				"[",
				"Attendee(name=Taylor Hebert, gender=female),",
				"Attendee(name=Blake Thornburn, gender=other)",
				"]",
			},
			expected: []interface{}{
				map[string]interface{}{
					"name":   "Taylor Hebert",
					"gender": "female",
				},
				map[string]interface{}{
					"name":   "Blake Thornburn",
					"gender": "other",
				},
			},
		}, {
			name: "Nested Array",
			input: []string{
				"[",
				"Attendee(name=Taylor Hebert, gender=female, abilities=",
				"[",
				"CapeData(codename=skitter, power=",
				"[warlord, controls bugs]",
				"),",
				"CapeData(codename=weaver, power=",
				"[controls butterflies]",
				"),",
				"CapeData(codename=kherpi, power=",
				"[you know]",
				")",
				"]",
				"),",
				"Attendee(name=Blake Thornburn, gender=other, abilities=",
				"[Diabolism, Shamanism, Glamour]",
				")",
				"]",
			},
			expected: []interface{}{
				map[string]interface{}{
					"name":   "Taylor Hebert",
					"gender": "female",
					"abilities": []interface{}{
						map[string]interface{}{
							"codename": "skitter",
							"power":    []string{"warlord", "controls bugs"},
						}, map[string]interface{}{
							"codename": "weaver",
							"power":    []string{"controls butterflies"},
						}, map[string]interface{}{
							"codename": "kherpi",
							"power":    []string{"you know"},
						},
					},
				},
				map[string]interface{}{
					"name":      "Blake Thornburn",
					"gender":    "other",
					"abilities": []string{"Diabolism", "Shamanism", "Glamour"},
				},
			},
		}, {
			name: "Simple Nested Class",
			input: []string{
				"Attendance(address=",
				"Address(name=Jewel Airport, postalCode=",
				"S(103912)",
				"),",
				"headcount=1023)",
			},
			expected: map[string]interface{}{
				"address": map[string]interface{}{
					"name":       "Jewel Airport",
					"postalCode": "S(103912)",
				},
				"headcount": "1023",
			},
		}, {
			name: "Nested class",
			input: []string{
				"ShippingLogUpdate(attempt=",
				"ShippingAttempt(id=182501, session=",
				"OrderSession(id=158171, address=",
				"Address(line1=City Square, postCode=",
				"S(112300)",
				"),",
				"createdAt=2022-03-04T09:08:21.052Z, quantity=4, shipperDescription=null, updatedAt=2022-03-04 03:08:21.055, billingAddress=",
				"BillingAddress(line1=Sentosa, postCode=",
				"S(612300)",
				"),",
				"paymentMethod=INHOUSE),",
				"transactionStatusEvent=",
				"TransactionStatusEvent(data=",
				"TransactionStatusEvent.TransactionStatus(transactionRef=ba9d5206-28de-43ec-bd6f-dadfcb5ef890, status=SUCCESS, orderRef=e4223595/829f/4b2d/be49/a3fb4f9b8044, encoding=encode=YXNkZndlcndlcndlcg==, transactionDescription=null)",
				")",
				")",
				")",
			},
			expected: map[string]interface{}{
				"attempt": map[string]interface{}{
					"id": "182501",
					"session": map[string]interface{}{
						"id": "158171",
						"address": map[string]interface{}{
							"line1":    "City Square",
							"postCode": "S(112300)",
						},
						"createdAt":          "2022-03-04T09:08:21.052Z",
						"quantity":           "4",
						"shipperDescription": "null",
						"updatedAt":          "2022-03-04 03:08:21.055",
						"billingAddress": map[string]interface{}{
							"line1":    "Sentosa",
							"postCode": "S(612300)",
						},
						"paymentMethod": "INHOUSE",
					},
					"transactionStatusEvent": map[string]interface{}{
						"data": map[string]interface{}{
							"transactionRef":         "ba9d5206-28de-43ec-bd6f-dadfcb5ef890",
							"status":                 "SUCCESS",
							"orderRef":               "e4223595/829f/4b2d/be49/a3fb4f9b8044",
							"encoding":               "encode=YXNkZndlcndlcndlcg==",
							"transactionDescription": "null",
						},
					},
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, Process(test.input))
		})
	}
}
