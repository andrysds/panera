package migrate

const (
	SuccessMessage = "success"
	FailMessage    = "fail"
)

func Init() string {
	if result := StandupInit(); result != SuccessMessage {
		return FailMessage
	}
	if result := StandupListInit(); result != SuccessMessage {
		return FailMessage
	}
	return SuccessMessage
}
