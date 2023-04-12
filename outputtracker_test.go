package outputtracker

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("OutputTracker", func() {
	var tracker *OutputTracker[string]
	var listener *OutputListener[string]

	BeforeEach(func() {
		listener = &OutputListener[string]{}
		tracker = listener.CreateTracker()
	})

	Describe("Data", func() {
		It("should return a copy of the output", func() {
			tracker.add("test data")
			data := tracker.Data()
			data[0] = "modified data"

			Expect(tracker.Data()).To(Equal([]string{"test data"}))
		})
	})

	Describe("Clear", func() {
		It("should clear the output and return the cleared data", func() {
			tracker.add("test data")
			clearedData := tracker.Clear()

			Expect(clearedData).To(Equal([]string{"test data"}))
			Expect(tracker.Data()).To(BeEmpty())
		})
	})

	Describe("Stop", func() {
		It("should remove the tracker from the listener", func() {
			tracker.Stop()

			Expect(listener.listeners).NotTo(ContainElement(tracker))
		})
	})
})
