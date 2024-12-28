package identity_server

import "strconv"

type Int64Identity int64

func (i *Int64Identity) Compare(u Unique) bool {
	return i.String() < u.String()
}

func (i *Int64Identity) String() string {
	return strconv.Itoa(int(*i))
}
