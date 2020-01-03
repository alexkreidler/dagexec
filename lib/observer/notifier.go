package observer

type (
	// Event defines an indication of a point-in-time occurrence.
	Event struct {
		// Data in this case is a simple int, but the actual
		// implementation would depend on the application.
		Data interface{}
	}

	// Observer defines a standard interface for instances that wish to list for
	// the occurrence of a specific event.
	Observer interface {
		// On allows an event to be "published" to interface implementations.
		// In the "real world", error handling would likely be implemented.
		On(Event)
	}

	// Notifier is the instance being observed.
	Notifier interface {
		// Register allows an instance to register itself to listen/observe
		// events.
		Register(Observer)
		// Deregister allows an instance to remove itself from the collection
		// of observers/listeners.
		Deregister(Observer)
		// Publish publishes new events to listeners. The method is not
		// absolutely necessary, as each implementation could define this itself
		// without losing functionality.
		Publish(Event)
	}
)
