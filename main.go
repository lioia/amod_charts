package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

type Solution struct {
	Solver    int
	Instance  int
	Status    int
	Runtime   float64
	Solution  float64
	Heuristic float64
}

func main() {
	sol0 := readCsv("assets/solution-0.csv")
	sol1 := readCsv("assets/solution-1.csv")
	sol2 := readCsv("assets/solution-2.csv")
	sol3 := readCsv("assets/solution-3.csv")
	sol4 := readCsv("assets/solution-4.csv")
	os.Mkdir("output", 0700)
	generateTimedOutCharts(sol0)
	generateRuntime(sol0, sol1, sol2)
	generateRuntimeDifference(sol0, sol3, "Precedence")
	generateRuntimeDifference(sol2, sol4, "Positional")
	generateSolDifference(sol4)
	generateSolDifferenceTimedOut(sol0, sol1)
}

func generateSolDifferenceTimedOut(sol0 []Solution, sol1 []Solution) {
	names := make([]string, 0)
	data0 := make([]opts.LineData, 0)
	data1 := make([]opts.LineData, 0)
	for i := 0; i < len(sol0); i++ {
		if sol0[i].Status == 2 {
			continue
		}
		names = append(names, fmt.Sprintf("%d", i+1))
		data0 = append(data0, opts.LineData{Value: sol0[i].Solution})
		data1 = append(data1, opts.LineData{Value: sol1[i].Solution})
	}
	line := charts.NewLine()
	line.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Solution Difference",
		Subtitle: "Precedence Model",
	}),
		charts.WithColorsOpts(opts.Colors{"blue", "red"}),
		charts.WithToolboxOpts(opts.Toolbox{
			Show:  true,
			Right: "20%",
			Feature: &opts.ToolBoxFeature{
				SaveAsImage: &opts.ToolBoxFeatureSaveAsImage{
					Show:  true,
					Type:  "png",
					Title: "Anything you want",
				},
			}},
		),
	)
	line.SetXAxis(names).
		AddSeries("Timed Out Solution", data0).
		AddSeries("Real Solution", data1)
	f, _ := os.Create("output/precedence_sol_difference.html")
	line.Render(f)
}

func generateSolDifference(sol3 []Solution) {
	names := make([]string, 0)
	data0 := make([]opts.LineData, 0)
	data1 := make([]opts.LineData, 0)
	for i := 0; i < len(sol3); i++ {
		names = append(names, fmt.Sprintf("%d", i+1))
		data0 = append(data0, opts.LineData{Value: sol3[i].Solution})
		data1 = append(data1, opts.LineData{Value: sol3[i].Heuristic})
	}
	line := charts.NewLine()
	line.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Heuristics - Solution",
		Subtitle: "Positional Model",
	}),
		charts.WithColorsOpts(opts.Colors{"blue", "red"}),
		charts.WithToolboxOpts(opts.Toolbox{
			Show:  true,
			Right: "20%",
			Feature: &opts.ToolBoxFeature{
				SaveAsImage: &opts.ToolBoxFeatureSaveAsImage{
					Show:  true,
					Type:  "png",
					Title: "Anything you want",
				},
			}},
		),
	)
	line.SetXAxis(names).
		AddSeries("No Heuristsics Solution", data0).
		AddSeries("Heuristics Solution", data1)
	f, _ := os.Create("output/sol_difference.html")
	line.Render(f)
}

func generateRuntimeDifference(sol_no_heur, sol_heur []Solution, name string) {
	names := make([]string, 0)
	data0 := make([]opts.LineData, 0)
	data1 := make([]opts.LineData, 0)
	for i := 0; i < len(sol_no_heur); i++ {
		names = append(names, fmt.Sprintf("%d", i+1))
		data0 = append(data0, opts.LineData{Value: sol_no_heur[i].Runtime})
		data1 = append(data1, opts.LineData{Value: sol_heur[i].Runtime})
	}
	line := charts.NewLine()
	line.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Heuristics - Runtime",
		Subtitle: fmt.Sprintf("%s Model", name),
	}),
		charts.WithColorsOpts(opts.Colors{"blue", "red"}),
		charts.WithToolboxOpts(opts.Toolbox{
			Show:  true,
			Right: "20%",
			Feature: &opts.ToolBoxFeature{
				SaveAsImage: &opts.ToolBoxFeatureSaveAsImage{
					Show:  true,
					Type:  "png",
					Title: "Anything you want",
				},
			}},
		),
	)
	line.SetXAxis(names).
		AddSeries("No Heuristics Runtime", data0).
		AddSeries("Heuristics Runtime", data1)
	f, _ := os.Create(fmt.Sprintf("output/run_difference_%s.html", name))
	line.Render(f)
}

func generateRuntime(sol0, sol1, sol2 []Solution) {
	names := make([]string, 0)
	data0 := make([]opts.LineData, 0)
	data1 := make([]opts.LineData, 0)
	data2 := make([]opts.LineData, 0)
	for i := 0; i < len(sol0); i++ {
		names = append(names, fmt.Sprintf("%d", i+1))
		data0 = append(data0, opts.LineData{Value: sol0[i].Runtime})
		data1 = append(data1, opts.LineData{Value: sol1[i].Runtime})
		data2 = append(data2, opts.LineData{Value: sol2[i].Runtime})
	}
	line := charts.NewLine()
	line.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title: "Runtime Comparison",
	}),
		charts.WithColorsOpts(opts.Colors{"orange", "blue", "red"}),
		charts.WithToolboxOpts(opts.Toolbox{
			Show:  true,
			Right: "20%",
			Feature: &opts.ToolBoxFeature{
				SaveAsImage: &opts.ToolBoxFeatureSaveAsImage{
					Show:  true,
					Type:  "png",
					Title: "Anything you want",
				},
			}},
		),
	)
	line.SetXAxis(names).
		AddSeries("Precedence Model", data0).
		AddSeries("Positional Model", data1).
		AddSeries("Time-Indexed Model", data2)
	f, _ := os.Create("output/run.html")
	line.Render(f)
}

func generateTimedOutCharts(sol []Solution) {
	completed := 0
	for _, v := range sol {
		if v.Status == 2 {
			completed += 1
		}
	}
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title: "Precedence Model - Status",
	}),
		charts.WithColorsOpts(opts.Colors{"green", "red"}),
		charts.WithToolboxOpts(opts.Toolbox{
			Show:  true,
			Right: "20%",
			Feature: &opts.ToolBoxFeature{
				SaveAsImage: &opts.ToolBoxFeatureSaveAsImage{
					Show:  true,
					Type:  "png",
					Title: "Anything you want",
				},
			}},
		),
	)
	bar.SetXAxis([]string{}).
		AddSeries("Completed", []opts.BarData{{Value: completed}}).
		AddSeries("Timed Out", []opts.BarData{{Value: len(sol) - completed}})
	f, _ := os.Create("output/timed.html")
	bar.Render(f)
}

func readCsv(file string) []Solution {
	f, err := os.Open(file)
	FailOnError(err)
	defer f.Close()
	reader := csv.NewReader(f)
	data, err := reader.ReadAll()
	FailOnError(err)
	var solutions []Solution
	for i, row := range data {
		if i == 0 {
			continue
		}
		solver, err := strconv.Atoi(row[0])
		FailOnError(err)
		instance, err := strconv.Atoi(row[1])
		FailOnError(err)
		status, err := strconv.Atoi(row[2])
		FailOnError(err)
		runtime, err := strconv.ParseFloat(row[3], 64)
		FailOnError(err)
		solution, err := strconv.ParseFloat(row[4], 64)
		FailOnError(err)
		heuristics, err := strconv.ParseFloat(row[5], 64)
		FailOnError(err)
		solutions = append(solutions, Solution{
			Solver:    solver,
			Instance:  instance,
			Status:    status,
			Runtime:   runtime,
			Solution:  solution,
			Heuristic: heuristics,
		})
	}
	return solutions
}

func FailOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
