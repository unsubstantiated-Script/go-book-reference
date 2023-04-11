package subscribers

import (
	"fmt"
	"goBook/goBook/magazine"
)

func printSubscriberInfo(s *magazine.Subscriber) {
	fmt.Println("Name: ", s.Name)
	fmt.Println("Monthly rate: ", s.Rate)
	fmt.Println("Active: ", s.Active)
	fmt.Println("Street: ", s.Address.Street)
	fmt.Println("City: ", s.Address.City)
	fmt.Println("State: ", s.Address.State)
	fmt.Println("Postal Code: ", s.Address.PostalCode)

}

func printEmployeeInfo(s *magazine.Employee) {
	fmt.Println("Name: ", s.Name)
	fmt.Println("Yearly salary: ", s.Salary)
}

func defaultSubscriber(name string) *magazine.Subscriber {
	var s magazine.Subscriber
	s.Name = name
	s.Rate = 5.99
	s.Active = true
	return &s
}

func applyDiscount(s *magazine.Subscriber) {
	s.Rate = .99
}

func Magazine() {
	subscriber1 := &magazine.Subscriber{
		Name:   "Aman Singh",
		Rate:   5.00,
		Active: true,
		Address: magazine.Address{
			Street:     "455 West Pest",
			City:       "Spokomption",
			State:      "WA",
			PostalCode: "99444",
		},
	}

	printSubscriberInfo(subscriber1)

	subscriber2 := defaultSubscriber("Skip Jonez")
	printSubscriberInfo(subscriber2)

	subscriber3 := &magazine.Subscriber{Name: "Bob Jones", Rate: 4.99, Active: true}
	printSubscriberInfo(subscriber3)

	//Using pointers
	var subDisc magazine.Subscriber
	applyDiscount(&subDisc)
	fmt.Println(subDisc.Rate)

	employee1 := &magazine.Employee{Name: "Billy Sundae", Salary: 60000}
	printEmployeeInfo(employee1)

}
