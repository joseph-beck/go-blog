package models

type Message struct {
	Model
	Message       string `gorm:"size:500" json:"message"`
	SenderRefer   uint
	Sender        User `gorm:"foreignKey:SenderRefer" json:"senderId"`
	ReceiverRefer uint
	Receiver      User `gorm:"foreignKey:ReceiverRefer" json:"receiverId"`
	// Chat     Chat   `gorm:"foreignKey:chatRefer" json:"chatId"`
}
