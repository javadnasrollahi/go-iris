package hfuncs

import (
	"strings"
)

func Domain(url string) string {
	var MainService string
	dwsd := strings.SplitN(strings.SplitN(strings.TrimPrefix(strings.TrimPrefix(url, "https://"), "http://"), "?", 2)[0], "/", 2)[0]

	dms := strings.SplitN(dwsd, ".", 3)
	//workes and just find the domain without subdomain
	//but for now , we just need a full url with subdomains due to pre set tokens in other servers.
	if len(dms) > 2 {
		if dms[0] == "www" || dms[0] == "WWW" {
			MainService = dms[1] + "." + dms[2]

		} else {
			MainService = dms[0] + "." + dms[1] + "." + dms[2]
		}

	} else {
		MainService = dwsd
	}
	return MainService
}
