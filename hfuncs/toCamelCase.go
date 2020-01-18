package hfuncs

import "strings"

/*
ToCamelCase : convert any thing to camel case
ToCamelCase("Hello !Hamid @I:m reza") // return : HelloHamidIMReza
*/
func ToCamelCase(s string) string {
	t := strings.TrimSpace(s)
	ch := strings.Split(t, " ")
	res := make([]string, len(ch))
	for i, v := range ch {
		ch[i] = strings.Trim(v, ";!@#$%^&*()_-,.}{~<>?|:")
		if ch[i] != "" {
			res = append(res, strings.Title(ch[i]))
		}
	}
	if len(res) > 0 {
		res[0] = strings.ToLower(res[0])
	}
	return strings.Join(res, "")
}
