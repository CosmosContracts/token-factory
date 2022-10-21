package types

type TokenFactoryQuery struct {
	Token *TokenQuery `json:"token,omitempty"`
}

// See https://github.com/CosmWasm/token-bindings/blob/main/packages/bindings/src/query.rs
type TokenQuery struct {
	/// Given a subdenom minted by a contract via `OsmosisMsg::MintTokens`,
	/// returns the full denom as used by `BankMsg::Send`.
	FullDenom       *FullDenom       `json:"full_denom,omitempty"`
	Metadata        *GetMetadata     `json:"metadata"`
	Admin           *DenomAdmin      `json:"admin"`
	DenomsByCreator *DenomsByCreator `json:"denoms_by_creator"`
	Params          *GetParams       `json:"params"`
}

// query types

type FullDenom struct {
	CreatorAddr string `json:"creator_addr"`
	Subdenom    string `json:"subdenom"`
}

type GetMetadata struct {
	Denom string `json:"denom"`
}

type DenomAdmin struct {
	Denom string `json:"denom"`
}

type DenomsByCreator struct {
	Creator string `json:"creator"`
}

type GetParams struct{}

// responses

type FullDenomResponse struct {
	Denom string `json:"denom"`
}

type AdminResponse struct {
	Admin string `json:"admin"`
}

type MetadataResponse struct {
	Metadata *Metadata `json:"metadata,omitempty"`
}

type DenomsByCreatorResponse struct {
	Denoms []string `json:"denoms"`
}

type ParamsResponse struct {
	Params Params `json:"params"`
}
