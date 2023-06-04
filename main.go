package main

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

type benchmarkStats struct {
	len      []int64
	nsOp     []float64
	bOp      []float64
	allocsOp []int64
}

type benchmarkStatsMap = map[string]*benchmarkStats

func main() {
	http.HandleFunc("/", httpSearchLogChart)
	http.ListenAndServe(":8081", nil)
}

func httpSearchLogChart(w http.ResponseWriter, _ *http.Request) {
	stats := newBenchmarkStatsMap()
	// create a new line instance
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Алгоритмы поиска",
			Subtitle: "Сравнение LinearSearch и BinarySearch",
		}),
	)

	var xAxisSetted bool

	for fnName, fnStats := range *stats {
		if !xAxisSetted {
			line.SetXAxis(fnStats.len)
			xAxisSetted = true
		}

		// Put data into instance
		lineData := make([]opts.LineData, 0)
		for _, v := range fnStats.nsOp {
			lineData = append(lineData, opts.LineData{Value: v})
		}
		line.AddSeries(fnName, lineData)

	}

	line.SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{ShowSymbol: true}))

	line.Render(w)
}

func newBenchmarkStatsMap() *benchmarkStatsMap {
	m := make(map[string]*benchmarkStats)
	parsedLogFile := parseBenchmarkSearchLogFile()

	for _, v := range parsedLogFile {
		name := v[1]
		len, _ := strconv.ParseInt(v[2], 10, 0)
		nsOp, _ := strconv.ParseFloat(v[3], 64)
		bOp, _ := strconv.ParseFloat(v[4], 64)
		allocsOp, _ := strconv.ParseInt(v[5], 10, 0)

		fnStats, exist := m[name]

		if exist {
			fnStats.len = append(fnStats.len, len)
			fnStats.nsOp = append(fnStats.nsOp, nsOp)
			fnStats.bOp = append(fnStats.bOp, bOp)
			fnStats.allocsOp = append(fnStats.allocsOp, allocsOp)
		} else {
			m[name] = &benchmarkStats{
				append(make([]int64, 0), len),
				append(make([]float64, 0), nsOp),
				append(make([]float64, 0), bOp),
				append(make([]int64, 0), allocsOp),
			}
		}
	}

	return &m
}

func parseBenchmarkSearchLogFile() [][]string {
	b, err := ioutil.ReadFile("./benchmarkSearch.log")
	if err != nil {
		panic(err)
	}
	benchmarkResult := string(b)
	regexBench := regexp.MustCompile(`([a-zA-Z]*)-(\d+)-.* (\d+\.?\d+?)[\s]ns.*[\s](\d+)[\s]B.* (\d+) allocs`)
	matches := regexBench.FindAllStringSubmatch(benchmarkResult, -1)

	return matches
}
