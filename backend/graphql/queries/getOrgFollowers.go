package queries


type Data_organization_followers struct{
	Data_organizations []struct{
		Organization_name string `json:"organization_name"`
		Follows []struct {
			User struct{
				Email string `json:"email"`
			} `json:"user"`
		}`json:"follows"`
	}`graphql:"data_organizations(where: {organization_id: {_eq: $organization_id}})"`

}
