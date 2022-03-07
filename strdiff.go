package strdiff

type Operation string

const (
	OperationInsert Operation = "insert"
	OperationRemove Operation = "remove"
	OperationModify Operation = "modify"
)

type Step struct {
	Operation Operation `json:"op"`
	Position  int       `json:"po"`
	Value     string    `json:"val"`
}

func NewInsert(po int, val byte) Step {
	return Step{
		Operation: OperationInsert,
		Position:  po,
		Value:     string(val),
	}
}

func NewRemove(po int) Step {
	return Step{
		Operation: OperationRemove,
		Position:  po,
	}
}

func NewModify(po int, to byte) Step {
	return Step{
    Operation: OperationModify,
		Position:  po,
		Value:     string(to),
	}
}

func Diff(a, b string) []Step {
	m := len(a)
	n := len(b)
	dp := make([][][]Step, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([][]Step, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = make([]Step, 0)
		}
	}
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
      if i == 0 && j == 0 {
        continue
      }
			if i == 0 && j != 0 {
				dp[i][j] = append(dp[i][j], dp[i][j-1]...)
        dp[i][j] = append(dp[i][j], NewInsert(j, b[j-1]))
			} else if j == 0 && i != 0 {
				dp[i][j] = append(dp[i][j], dp[i-1][j]...)
				dp[i][j] = append(dp[i][j], NewRemove(0))
			} else if a[i-1] == b[j-1] {
				dp[i][j] = append(dp[i][j], dp[i-1][j-1]...)
			} else {
				insert := append(make([]Step, 0, len(dp[i][j-1])+1), dp[i][j-1]...)
				insert = append(insert, NewInsert(j-2, b[j-1]))
				remove := append(make([]Step, 0, len(dp[i-1][j])+1), dp[i-1][j]...)
				remove = append(remove, NewRemove(j))
				modify := append(make([]Step, 0, len(dp[i-1][j-1])+1), dp[i-1][j-1]...)
				modify = append(modify, NewModify(j-1, b[j-1]))
        dp[i][j] = insert
        if len(remove) < len(dp[i][j]) {
          dp[i][j] = remove
        }
        if len(modify) < len(dp[i][j]) {
          dp[i][j] = modify
        }
			}
		}
	}
  return dp[m][n]
}
