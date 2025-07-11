package types

func NewMsgCreateDenom(
	owner string,
	denom string,
	description string,
	ticker string,
	precision int64,
	url string,
	maxSupply int64,
	supply int64,
	canChangeMaxSupply bool,

) *MsgCreateDenom {
	return &MsgCreateDenom{
		Owner:              owner,
		Denom:              denom,
		Description:        description,
		Ticker:             ticker,
		Precision:          precision,
		Url:                url,
		MaxSupply:          maxSupply,
		Supply:             supply,
		CanChangeMaxSupply: canChangeMaxSupply,
	}
}

func NewMsgUpdateDenom(
	owner string,
	denom string,
	description string,
	ticker string,
	precision int64,
	url string,
	maxSupply int64,
	supply int64,
	canChangeMaxSupply bool,

) *MsgUpdateDenom {
	return &MsgUpdateDenom{
		Owner:              owner,
		Denom:              denom,
		Description:        description,
		Ticker:             ticker,
		Precision:          precision,
		Url:                url,
		MaxSupply:          maxSupply,
		Supply:             supply,
		CanChangeMaxSupply: canChangeMaxSupply,
	}
}

func NewMsgDeleteDenom(
	owner string,
	denom string,

) *MsgDeleteDenom {
	return &MsgDeleteDenom{
		Owner: owner,
		Denom: denom,
	}
}
