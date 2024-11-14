package footballKit

var (
	mpGoalCn2Num map[string]float64
	mpGoalNum2Cn map[float64]string

	ArrGoalCnHome = []string{"", "平手", "平手/半球", "半球", "半球/一球", "一球", "一球/球半", "球半", "球半/两球", "两球", "两球/两球半",
		"两球半/三球", "三球", "三球/三球半", "三球半/四球", "四球", "四/四球半", "四球半", "四球半/五", "五球", "五/五球半", "五球半", "五球半/六",
		"六球", "六/六球半", "六球半", "六球半/七", "七球", "七/七球半", "七球半", "七球半/八", "八球", "八/八球半", "八球半", "八球半/九", "九球",
		"九/九球半", "九球半", "九球半/十", "十球"}
	ArrGoalCnAway = []string{"", "平手", "受让平手/半球", "受让半球", "受让半球/一球", "受让一球", "受让一球/球半", "受让球半",
		"受让球半/两球", "受让两球", "受让两球/两球半", "受让两球半/三球", "受让三球", "受让三球/三球半", "受让三球半/四球"}
)

func init() {
	mpGoalCn2Num = make(map[string]float64)
	mpGoalNum2Cn = make(map[float64]string)

	mpGoalCn2Num["平手"] = 0.125
	mpGoalCn2Num["平"] = 0.125
	mpGoalCn2Num["平手/半球"] = 0.25
	mpGoalCn2Num["平/半"] = 0.25
	mpGoalCn2Num["半球/一球"] = 0.75
	mpGoalCn2Num["一球"] = 1
	mpGoalCn2Num["一球/球半"] = 1.25
	mpGoalCn2Num["球半"] = 1.5
	mpGoalCn2Num["球半/两球"] = 1.75
	mpGoalCn2Num["两球"] = 2
	mpGoalCn2Num["两球/两球半"] = 2.25
	mpGoalCn2Num["两球半/三球"] = 2.75
	mpGoalCn2Num["三球"] = 3
	mpGoalCn2Num["三球/三球半"] = 3.25
	mpGoalCn2Num["三球半/四球"] = 3.75
	mpGoalCn2Num["四球"] = 4
	mpGoalCn2Num["四球/四球半"] = 4.25
	mpGoalCn2Num["四球半/五球"] = 4.75
	mpGoalCn2Num["五球"] = 5
	mpGoalCn2Num["五球/五球半"] = 5.25
	mpGoalCn2Num["五球半/六球"] = 5.75
	mpGoalCn2Num["六球"] = 6
	mpGoalCn2Num["半球"] = 0.5
	mpGoalCn2Num["两球半"] = 2.5
	mpGoalCn2Num["三球半"] = 3.5
	mpGoalCn2Num["四球半"] = 4.5
	mpGoalCn2Num["五球半"] = 5.5
	mpGoalCn2Num["平手"] = -0.125
	mpGoalCn2Num["受让平手/半球"] = -0.25
	mpGoalCn2Num["受让半球/一球"] = -0.75
	mpGoalCn2Num["受让一球"] = -1
	mpGoalCn2Num["受让一球/球半"] = -1.25
	mpGoalCn2Num["受让球半"] = -1.5
	mpGoalCn2Num["受让球半/两球"] = -1.75
	mpGoalCn2Num["受让两球"] = -2
	mpGoalCn2Num["受让两球/两球半"] = -2.25
	mpGoalCn2Num["受让两球半/三球"] = -2.75
	mpGoalCn2Num["受让三球"] = -3
	mpGoalCn2Num["受让三球/三球半"] = -3.25
	mpGoalCn2Num["受让三球半/四球"] = -3.75
	mpGoalCn2Num["受让四球"] = -4
	mpGoalCn2Num["受让四球/四球半"] = -4.25
	mpGoalCn2Num["受让四球半/五球"] = -4.75
	mpGoalCn2Num["受让五球"] = -5
	mpGoalCn2Num["受让五球/五球半"] = -5.25
	mpGoalCn2Num["受让五球半/六球"] = -5.75
	mpGoalCn2Num["受让六球"] = -6
	mpGoalCn2Num["受让半球"] = -0.5
	mpGoalCn2Num["受让两球半"] = -2.5
	mpGoalCn2Num["受让三球半"] = -3.5
	mpGoalCn2Num["受让四球半"] = -4.5
	mpGoalCn2Num["受让五球半"] = -5.5

	for k, v := range mpGoalCn2Num {
		mpGoalNum2Cn[v] = k
	}
}
func GetGoalCnHome(i int) string {
	if len(ArrGoalCnHome) > i && i > 0 {
		return ArrGoalCnHome[i]
	}

	return ""
}
func GetGoalCnAway(i int) string {
	if len(ArrGoalCnAway) > i && i > 0 {
		return ArrGoalCnAway[i]
	}

	return ""
}

/*
*
//数字盘口转汉汉字
*/
func Goal2GoalCn(goal float64) string {
	return mpGoalNum2Cn[goal]
}

/*
*
汉字盘口转数字盘口
*/
func GoalCn2Goal(cn string) float64 {
	return mpGoalCn2Num[cn]
}

func sortAbc(v1, v2, v3 float64) (float64, float64, float64) {
	if v1 < v2 {
		t := v1
		v1 = v2
		v2 = t
	}
	if v2 < v3 {
		t := v2
		v2 = v3
		v3 = t
	}

	return v1, v2, v3
}

/*
*
盘力指数
*/
func EurOdds2AsiaOdds(pl_b, pl_c, pl_d float64) (float64, float64) {
	v1 := 0.0
	e := 1 / pl_b
	f := 1 / pl_c
	g := 1 / pl_d
	h := 1 / (e + f + g)
	i := h / pl_b
	j := h / pl_c
	k := h / pl_d

	m, n, _ := sortAbc(i, j, k)

	o := (m - n) / n
	p := 0.5
	if o < 0.5 {
		p = 0.25
	}

	o1 := int(o * 1000)
	p1 := int(p * 1000)
	q := float64((p1/o1)*o1) / 1000

	r := p + q
	ac := r * 0.5
	_, _, l := sortAbc(pl_b, pl_c, pl_d)

	if l == pl_d {
		ac = -ac
	}
	v1 = ac
	s := r
	if r < 1 {
		s = 0.5
	}
	t := 0.0
	if s < 1 {
		t = 0.25
	}
	u := r + t
	v := u + p

	x := 0.5
	if u == 0.75 {
		x = 0.25
	}
	w := 1.1
	if r == 0.25 {
		w = 0.925
	}
	y := (w - 1.1*u/v) / x
	z := w - (o-q)*y
	vv := int(z * 1000)
	v2 := float64(vv) / 1000

	return v1, v2
}
