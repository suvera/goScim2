package scim2

func strArrToIface(values ...string) (ret []interface{}) {
	for _, s := range values {
		ret = append(ret, s)
	}

	return
}
