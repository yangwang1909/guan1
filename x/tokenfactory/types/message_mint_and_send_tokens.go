package types

func NewMsgMintAndSendTokens(owner string, denom string, amount int64, recipient string) *MsgMintAndSendTokens {
	return &MsgMintAndSendTokens{
		Owner:     owner,
		Denom:     denom,
		Amount:    amount,
		Recipient: recipient,
	}
}
