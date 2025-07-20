package types

func NewMsgUpdateOwner(owner string, denom string, newOwner string) *MsgUpdateOwner {
	return &MsgUpdateOwner{
		Owner:    owner,
		Denom:    denom,
		NewOwner: newOwner,
	}
}
