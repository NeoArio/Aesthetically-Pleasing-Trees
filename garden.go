package main

type Garden []int

func (g Garden) Solve() int {
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