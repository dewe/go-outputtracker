package outputtracker

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("OutputListener", func() {
	var listener *OutputListener[string]

	BeforeEach(func() {
		listener = &OutputListener[string]{}
	})

	Describe("CreateTracker", func() {
		It("should create a new tracker and add it to the list of listeners", func() {
			tracker := listener.CreateTracker()

			Expect(listener.listeners).To(ContainElement(tracker))
		})
	})

	Describe("Track", func() {
		It("should send data to all trackers", func() {
			tracker1 := listener.CreateTracker()
			tracker2 := listener.CreateTracker()

			listener.Track("test data")

			Expect(tracker1.Data()).To(Equal([]string{"test data"}))
			Expect(tracker2.Data()).To(Equal([]string{"test data"}))
		})	
	})	
})
