package enum

type Status int

const (
	UserInactive Status = 0
	UserActive   Status = 1
	UserDeleted  Status = 9
)
