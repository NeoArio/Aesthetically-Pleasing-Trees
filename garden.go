package main

type Garden []int

func (g Garden) SolveBySecondSolution() int{
	var UpOrDowns []int

	for i:=0; i < len(g)-2; i++{
		if (g[i]>=g[i+1] && g[i+1]>=g[i+2]) || (g[i]<=g[i+1] && g[i+1]<=g[i+2]){
			UpOrDowns = append(UpOrDowns, i)
		}
	}

	if len(UpOrDowns) == 0{
		return 0
	}

	if len(UpOrDowns) > 1{
		return -1
	}

	i := UpOrDowns[0]

	if i == 0{
		if (g[i]<=g[i+1] && g[i+1]<=g[i+3]) || (g[i]>=g[i+1] && g[i+1]>=g[i+3]){
			return 2
		}else {
			return 3
		}
	}

	cap := len(g)
	if i == cap-3{
		if (g[cap-1]<=g[cap-2] && g[cap-2]<=g[cap-4]) || (g[cap-1]>=g[cap-2] && g[cap-2]>=g[cap-4]){
			return 2
		} else{
			return 3
		}
	}

	maxWays := 3
	if (g[i-1] <= g[i+1] && g[i+1] <= g[i+2]) || (g[i-1] >= g[i+1] && g[i+1] >= g[i+2]){
		maxWays--
	}
	if (g[i] <= g[i+1] && g[i+1] <= g[i+3]) || (g[i] >= g[i+1] && g[i+1] >= g[i+3]){
		maxWays--
	}

	return maxWays
}

func (g Garden) SolveByFirstSolution() int {
	var waysToCut int
	if g.IsAesthetic() {
		return 0
	}

	subGardens := g.FindSubGardens()

	for _, subGarden := range subGardens{
		if subGarden.IsAesthetic(){
			waysToCut++
		}
	}

	if waysToCut == 0 {
		return -1
	}

	return waysToCut
}

func (g Garden) FindSubGardens() (subGardens []Garden) {
	for i, _ := range g {
		var subGarden []int
		for j,_ := range g{
			if j == i{
				continue
			}
			subGarden = append(subGarden, g[j])
		}
		subGardens = append(subGardens, subGarden)
	}
	return subGardens
}

func (g Garden)IsAesthetic() bool {
	for i:=1; i<len(g)-1; i++{
		if (g[i] >= g[i-1] && g[i] <= g[i+1]) || (g[i] <= g[i-1] && g[i] >= g[i+1]){
			return false
		}
	}
	return true
}