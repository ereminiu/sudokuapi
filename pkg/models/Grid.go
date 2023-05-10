package models

type Grid struct {
	Array [][]string `json:"grid" binding:"required"`
}

type Response struct {
	Sol [][]string `json:"sol" binding:"required"`
}

func NewResponse(a [][]rune) Response {
	n := 9
	res := make([][]string, n)
	for i := 0; i < n; i++ {
		res[i] = make([]string, n)
		for j, val := range a[i] {
			res[i][j] = string(val)
		}
	}

	return Response{Sol: res}
}
