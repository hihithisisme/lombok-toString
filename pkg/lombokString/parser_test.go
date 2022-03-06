package lombokString_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"lombok-toString/pkg/lombokString"
	"testing"
)

func TestParseLombokStringAsJSON(t *testing.T) {
	t.Run("simple with array", func(t *testing.T) {
		simple := "Attendance(venue=Jewel Airport, headcount=1023, attendees=[Attendee(name=Taylor Hebert, gender=female), Attendee(name=Blake Thornburn, gender=other)])"

		actual := lombokString.New(simple).ParseAsJSON()

		assert.Contains(t, actual, `"venue": "Jewel Airport"`)
		assert.Contains(t, actual, `"headcount": 1023`)
		assert.Contains(t, actual, `"name": "Taylor Hebert"`)
		assert.Contains(t, actual, `"name": "Blake Thornburn"`)
		assert.Contains(t, actual, `"gender": "female"`)
		assert.Contains(t, actual, `"gender": "other"`)
	})

	t.Run("array", func(t *testing.T) {
		array := "Transaction(id=86634, actions=[Action(id=150623, user=Sharon, action=ACTIVATED), Action(id=150624, user=Bob Smith, action=ABANDONED)], tempMap=hello)"

		actual := lombokString.New(array).ParseAsJSON()
		fmt.Println(actual)

		assert.Contains(t, actual, `"tempMap": "hello"`)
		assert.Contains(t, actual, `"id": 86634`)
		assert.Contains(t, actual, `"id": 150623`)
		assert.Contains(t, actual, `"id": 150624`)
		assert.Contains(t, actual, `"user": "Sharon"`)
		assert.Contains(t, actual, `"user": "Bob Smith"`)
		assert.Contains(t, actual, `"action": "ACTIVATED"`)
		assert.Contains(t, actual, `"action": "ABANDONED"`)
	})

	t.Run("complicated and nested", func(t *testing.T) {
		complicatedNested := "[ShippingLogUpdate(attempt=ShippingAttempt(id=182501, session=OrderSession(id=158171, customerRef=ba9d5206-28de-43ec-bd6f-dadfcb5ef890, orderRef=ba9d5206-28de-43ec-bd6f-dadfcb5ef890, shopperName=Jon Snow White, currency=CHF, amount=100.0100, notes=For the Seven Dwarves, description=null, createdAt=2022-03-04T09:08:21.052Z, quantity=4, shipperDescription=null, address=Session.Address(address1=City Square, address2=, city=Singapore, state=, postalCode=810390, country=SG), enableDarkMode=false, status=SHIPPING, updatedAt=2022-03-04 03:08:21.055), paymentMethod=INHOUSE, status=SUCCESS, paymentRef=e4223595-829f-4b2d-be49-a3fb4f9b8044, created=2022-03-04 03:06:44.0, updated=2022-03-04 03:08:21.053), transactionStatusEvent=TransactionStatusEvent(data=TransactionStatusEvent.TransactionStatus(transactionRef=ba9d5206-28de-43ec-bd6f-dadfcb5ef890-158171-182501, status=SUCCESS, orderRef=e4223595-829f-4b2d-be49-a3fb4f9b8044, transactionDescription=null, paymentProvider=BANK)))]"

		actual := lombokString.New(complicatedNested).ParseAsJSON()

		assert.Contains(t, actual, `"transactionRef": "ba9d5206-28de-43ec-bd6f-dadfcb5ef890-158171-182501"`)
		assert.Contains(t, actual, `"address1": "City Square"`)
		assert.Contains(t, actual, `updatedAt": "2022-03-04 03:08:21.055`)
		assert.Contains(t, actual, `"quantity": 4`)
	})
}
