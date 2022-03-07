package lombokString_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"lombok-toString/pkg/lombokString"
	"testing"
)

// TODO: should use assert.JSONEq instead of assert.Contains
// https://pkg.go.dev/github.com/stretchr/testify/assert#JSONEq
type testCase struct {
	subtestName    string
	lString        string
	shouldContains []string
}

var testCases = []testCase{
	{
		subtestName: "Simple with Array",
		lString:     "Attendance(venue=Jewel Airport, headcount=1023, attendees=[Attendee(name=Taylor Hebert, gender=female), Attendee(name=Blake Thornburn, gender=other)])",
		shouldContains: []string{
			`"venue": "Jewel Airport"`,
			`"headcount": 1023`,
			`"name": "Taylor Hebert"`,
			`"name": "Blake Thornburn"`,
			`"gender": "female"`,
			`"gender": "other"`,
		},
	}, {
		subtestName: "Array with some special characters validation",
		lString:     "Transaction(id=86634, actions=[Action(id=150623, user=Sharon II-2, action=ACTIVATED), Action(id=150625, user=Matt Johnson, action=RETRIEVED), Action(id=150626, user=Jane Doe, email=jane@doe.com, action=ATTEMPTED), Action(id=150627, user=Rock Fair, action=VALIDATED)], tempMap=hello)",
		shouldContains: []string{
			`"tempMap": "hello"`,
			`"id": 86634`,

			`"id": 150623`,
			`"action": "ACTIVATED"`,
			`"user": "Sharon II-2"`,

			`"id": 150625`,
			`"action": "RETRIEVED"`,
			`"user": "Matt Johnson"`,

			`"id": 150626`,
			`"action": "ATTEMPTED"`,
			`"user": "Jane Doe"`,
			`"email": "jane@doe.com"`,

			`"id": 150627`,
			`"action": "VALIDATED"`,
			`"user": "Rock Fair"`,
		},
	}, {
		subtestName: "Nested 3 Levels",
		lString:     "ShippingLogUpdate(attempt=ShippingAttempt(id=182501, session=OrderSession(id=158171, customerRef=ba9d5206-28de-43ec-bd6f-dadfcb5ef890, orderRef=ba9d5206-28de-43ec-bd6f-dadfcb5ef890, shopperName=Jon Snow White, currency=CHF, amount=100.0100, notes=For the Seven Dwarves, description=null, createdAt=2022-03-04T09:08:21.052Z, quantity=4, shipperDescription=null, address=Session.Address(address1=City Square, address2=, city=Singapore, state=, postalCode=810390, country=SG), enableDarkMode=false, status=SHIPPING, updatedAt=2022-03-04 03:08:21.055), paymentMethod=INHOUSE, status=SUCCESS, paymentRef=e4223595-829f-4b2d-be49-a3fb4f9b8044, created=2022-03-04 03:06:44.0, updated=2022-03-04 03:08:21.053), transactionStatusEvent=TransactionStatusEvent(data=TransactionStatusEvent.TransactionStatus(transactionRef=ba9d5206-28de-43ec-bd6f-dadfcb5ef890-158171-182501, status=SUCCESS, orderRef=e4223595-829f-4b2d-be49-a3fb4f9b8044, transactionDescription=null, paymentProvider=BANK)))",
		shouldContains: []string{
			`"transactionRef": "ba9d5206-28de-43ec-bd6f-dadfcb5ef890-158171-182501"`,
			`"address1": "City Square"`,
			`"updatedAt": "2022-03-04 03:08:21.055`,
			`"updated": "2022-03-04 03:08:21.053"`,
			`"quantity": 4`,
			`"description": null`,
		},
	},
}

func TestParseLombokStringAsJSON(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.subtestName, func(t *testing.T) {
			args := lombokString.InterfaceArgs{}
			actual := lombokString.New(tc.lString).ParseAsJSON(args)
			// TODO: remove this utility -- should be supported by assert (but is not currently)
			shouldPrintActual := false

			for _, substring := range tc.shouldContains {
				if !assert.Contains(t, actual, substring) {
					shouldPrintActual = true
				}
			}

			if shouldPrintActual {
				fmt.Printf("\nOutput of { %s } is:\n", tc.lString)
				fmt.Println(actual)
			}
		})
	}
}

func TestParseLombokStringAsJSON_WithExcludeNulls(t *testing.T) {
	tc := testCase{
		subtestName: "Nested 3 Levels",
		lString:     "ShippingLogUpdate(attempt=ShippingAttempt(id=182501, session=OrderSession(id=158171, customerRef=ba9d5206-28de-43ec-bd6f-dadfcb5ef890, orderRef=ba9d5206-28de-43ec-bd6f-dadfcb5ef890, shopperName=Jon Snow White, currency=CHF, amount=100.0100, notes=For the Seven Dwarves, description=null, createdAt=2022-03-04T09:08:21.052Z, quantity=4, shipperDescription=null, address=Session.Address(address1=City Square, address2=, city=Singapore, state=, postalCode=810390, country=SG), enableDarkMode=false, status=SHIPPING, updatedAt=2022-03-04 03:08:21.055), paymentMethod=INHOUSE, status=SUCCESS, paymentRef=e4223595-829f-4b2d-be49-a3fb4f9b8044, created=2022-03-04 03:06:44.0, updated=2022-03-04 03:08:21.053), transactionStatusEvent=TransactionStatusEvent(data=TransactionStatusEvent.TransactionStatus(transactionRef=ba9d5206-28de-43ec-bd6f-dadfcb5ef890-158171-182501, status=SUCCESS, orderRef=e4223595-829f-4b2d-be49-a3fb4f9b8044, transactionDescription=null, paymentProvider=BANK)))",
		shouldContains: []string{
			`"transactionRef": "ba9d5206-28de-43ec-bd6f-dadfcb5ef890-158171-182501"`,
			`"address1": "City Square"`,
			`"updatedAt": "2022-03-04 03:08:21.055`,
			`"updated": "2022-03-04 03:08:21.053"`,
			`"quantity": 4`,
		},
	}
	shouldNotContain := []string{
		`"description"`,
		`"shipperDescription"`,
		`null`,
	}
	args := lombokString.InterfaceArgs{
		ShouldExcludeNulls: true,
	}

	actual := lombokString.New(tc.lString).ParseAsJSON(args)

	for _, substring := range tc.shouldContains {
		assert.Contains(t, actual, substring)
	}
	for _, substring := range shouldNotContain {
		assert.NotContains(t, actual, substring)
	}
}

func TestParseLombokStringAsJSON_WithMinify(t *testing.T) {
	tc := testCase{
		subtestName: "should minify output",
		lString:     "ShippingLogUpdate(attempt=ShippingAttempt(id=182501, session=OrderSession(id=158171, customerRef=ba9d5206-28de-43ec-bd6f-dadfcb5ef890, orderRef=ba9d5206-28de-43ec-bd6f-dadfcb5ef890, shopperName=Jon Snow White, currency=CHF, amount=100.0100, notes=For the Seven Dwarves, description=null, createdAt=2022-03-04T09:08:21.052Z, quantity=4, shipperDescription=null, address=Session.Address(address1=City Square, address2=, city=Singapore, state=, postalCode=810390, country=SG), enableDarkMode=false, status=SHIPPING, updatedAt=2022-03-04 03:08:21.055), paymentMethod=INHOUSE, status=SUCCESS, paymentRef=e4223595-829f-4b2d-be49-a3fb4f9b8044, created=2022-03-04 03:06:44.0, updated=2022-03-04 03:08:21.053), transactionStatusEvent=TransactionStatusEvent(data=TransactionStatusEvent.TransactionStatus(transactionRef=ba9d5206-28de-43ec-bd6f-dadfcb5ef890-158171-182501, status=SUCCESS, orderRef=e4223595-829f-4b2d-be49-a3fb4f9b8044, transactionDescription=null, paymentProvider=BANK)))",
		shouldContains: []string{
			`"transactionRef":"ba9d5206-28de-43ec-bd6f-dadfcb5ef890-158171-182501"`,
			`"address1":"City Square"`,
			`"updatedAt":"2022-03-04 03:08:21.055`,
			`"updated":"2022-03-04 03:08:21.053"`,
			`"quantity":4`,
		},
	}
	args := lombokString.InterfaceArgs{
		ShouldMinify: true,
	}

	actual := lombokString.New(tc.lString).ParseAsJSON(args)

	for _, substring := range tc.shouldContains {
		assert.Contains(t, actual, substring)
	}
	assert.NotContains(t, actual, "  ")
	assert.NotContains(t, actual, "\t")
}
