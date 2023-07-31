package microsoft

import (
	"encoding/json"
)

func (c *microsoftConnector) ExtendPayload(scopes []string, payload []byte, cdata []byte) ([]byte, error) {
	// var cd ConnectorData
	// originalClaims := map[string]interface{}{}
	var originalClaims map[string]interface{}

	c.logger.Info("(microsoft) ExtendPayload scopes", scopes)
	c.logger.Info("(microsoft) ExtendPayload payload", string(payload))
	c.logger.Info("(microsoft) ExtendPayload cdata", string(cdata))

	if err := json.Unmarshal(payload, &originalClaims); err != nil {
		return payload, err
	}
	// if err := json.Unmarshal(cdata, &cd); err != nil {
	// 	return payload, err
	// }
	// for _, scope := range scopes {
	// 	if scope == "federated:id" {
	// 		originalClaims["iam_access_token"] = string(cd.AccessToken)
	// 	}
	// 	// Experimental tenant scoping
	// 	if strings.HasPrefix(scope, "tenant:") {
	// 		group := strings.TrimPrefix(scope, "tenant:")
	// 		if slices.Contains(c.tenantGroups, group) {
	// 			var tenants []string
	// 			// Iterate through introspect and add OrgID as tenant when matched
	// 			for _, org := range cd.Introspect.Organizations.OrganizationList {
	// 				for _, orgGroup := range org.Groups {
	// 					if group == orgGroup {
	// 						tenants = append(tenants, org.OrganizationID)
	// 					}
	// 				}
	// 			}
	// 			if len(tenants) > 0 {
	// 				originalClaims[scope] = tenants
	// 			}
	// 		}
	// 	}
	// }
	// originalClaims["moid"] = cd.Introspect.Organizations.ManagingOrganization
	// // Rewrite subject
	// var orgSubs []string
	// for _, org := range cd.Introspect.Organizations.OrganizationList {
	// 	if org.OrganizationID != cd.TrustedIDPOrg {
	// 		continue
	// 	}
	// 	for _, group := range org.Groups {
	// 		if strings.HasPrefix(group, "sub-") {
	// 			orgSubs = append(orgSubs, fmt.Sprintf("sub:%s", strings.TrimPrefix(group, "sub-")))
	// 		}
	// 	}
	// }
	// // Rewrite name
	// if cd.User.GivenName != "" {
	// 	originalClaims["name"] = fmt.Sprintf("%s %s", cd.User.GivenName, cd.User.FamilyName)
	// }
	// // Inject username
	// if cd.Introspect.Username != "" {
	// 	originalClaims["username"] = cd.Introspect.Username
	// }
	// if len(orgSubs) > 0 {
	// 	subs := strings.Join(orgSubs, ":")
	// 	origSub := originalClaims["sub"].(string)
	// 	originalClaims["sub"] = fmt.Sprintf("%s:id:%s", subs, origSub)
	// }

	originalClaims["aaaaa"] = "bbbbbbb"
	originalClaims["ccccc"] = "bbbbbbb"

	extendedPayload, err := json.Marshal(originalClaims)
	if err != nil {
		return payload, err
	}
	return extendedPayload, nil
}
