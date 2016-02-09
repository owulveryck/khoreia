/* package event is only used to declare the event type and avoid circular imports */
package event

type Event struct {
	IsDone bool
	Msg    string
}
