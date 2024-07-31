package task

const (
	TaskSendEmail  = "SEND_EMAIL"
	TaskUploadFile = "UPLOAD_FILE"
)

type (
	SendEmail struct {
		From    string `json:"from"`
		To      string `json:"to"`
		Subject string `json:"subject"`
		Content string `json:"content"`
	}

	UploadFile struct {
		Origin      string `json:"origin"`
		Destination string `json:"destination"`
	}
)
