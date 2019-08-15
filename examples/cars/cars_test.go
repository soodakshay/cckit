package cars_test

import (
	"testing"

<<<<<<< HEAD
=======
	"github.com/soodakshay/cckit/examples/cars"
>>>>>>> d9270d9c7e0def8422d5975be671d71af6c232c9
	examplecert "github.com/soodakshay/cckit/examples/cert"
	"github.com/soodakshay/cckit/extensions/owner"
	"github.com/soodakshay/cckit/state"
	testcc "github.com/soodakshay/cckit/testing"
	expectcc "github.com/soodakshay/cckit/testing/expect"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCars(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cars Suite")
}

var _ = Describe(`Cars`, func() {

	//Create chaincode mock
	cc := testcc.NewMockStub(`cars`, cars.New())
	ccWithoutAC := testcc.NewMockStub(`cars`, cars.NewWithoutAccessControl())

	// load actor certificates
	actors, err := testcc.IdentitiesFromFiles(`SOME_MSP`, map[string]string{
		`authority`: `s7techlab.pem`,
		`someone`:   `victor-nosov.pem`}, examplecert.Content)
	if err != nil {
		panic(err)
	}

	BeforeSuite(func() {
		// init chaincode
		expectcc.ResponseOk(cc.From(actors[`authority`]).Init()) // init chaincode from authority
	})

	Describe("Car", func() {

		It("Allow authority to add information about car", func() {
			//invoke chaincode method from authority actor
			expectcc.ResponseOk(cc.From(actors[`authority`]).Invoke(`carRegister`, cars.Payloads[0]))
		})

		It("Disallow non authority to add information about car", func() {
			//invoke chaincode method from non authority actor
			expectcc.ResponseError(
				cc.From(actors[`someone`]).Invoke(`carRegister`, cars.Payloads[0]),
				owner.ErrOwnerOnly) // expect "only owner" error
		})

		It("Allow non authority to add information about car to chaincode without access control", func() {
			//invoke chaincode method from non authority actor
			expectcc.ResponseOk(
				ccWithoutAC.From(actors[`someone`]).Invoke(`carRegister`, cars.Payloads[0]))
		})

		It("Disallow authority to add duplicate information about car", func() {
			expectcc.ResponseError(
				cc.From(actors[`authority`]).Invoke(`carRegister`, cars.Payloads[0]),
				state.ErrKeyAlreadyExists) //expect car id already exists
		})

		It("Allow everyone to retrieve car information", func() {
			car := expectcc.PayloadIs(cc.Invoke(`carGet`, cars.Payloads[0].Id),
				&cars.Car{}).(cars.Car)

			Expect(car.Title).To(Equal(cars.Payloads[0].Title))
			Expect(car.Id).To(Equal(cars.Payloads[0].Id))
		})

		It("Allow everyone to get car list", func() {
			//  &[]Car{} - declares target type for unmarshalling from []byte received from chaincode
			cc := expectcc.PayloadIs(cc.Invoke(`carList`), &[]cars.Car{}).([]cars.Car)

			Expect(cc).To(HaveLen(1))
			Expect(cc[0].Id).To(Equal(cars.Payloads[0].Id))
		})

		It("Allow authority to add more information about car", func() {
			// register second car
			expectcc.ResponseOk(cc.From(actors[`authority`]).Invoke(`carRegister`, cars.Payloads[1]))
			cc := expectcc.PayloadIs(
				cc.From(actors[`authority`]).Invoke(`carList`),
				&[]cars.Car{}).([]cars.Car)

			Expect(cc).To(HaveLen(2))
		})
	})
})
