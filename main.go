package main

import (
	"encoding/json"
	"fmt"
	"image/jpeg"
	_ "image/jpeg"
	"os"
	"time"

	"github.com/Damione1/thread-art-generator/threadGenerator"
)

func main() {
	tg := new(threadGenerator.ThreadGenerator)

	exportPrefix := time.Now().Format("2006_01_02_15_04_05") + "_"
	outputFolder := "output/" + time.Now().Format("2006_01_02_15_04_05") + "/"

	//create output folder
	err := os.MkdirAll(outputFolder, 0777)
	if err != nil {
		panic(err)
	}

	args := threadGenerator.Args{
		NailsQuantity:     200,
		ImgSize:           1000,
		MaxPaths:          15000,
		StartingNail:      0,
		MinimumDifference: 10,
		BrightnessFactor:  50,
		ImageName:         "source_1.jpg",
	}

	//Use this to calibrate the real nail position
	// args.CalibrationPoints = map[int]float32{
	// 	0:   0.0,
	// 	25:  0.0,
	// 	50:  0.0,
	// 	75:  0.0,
	// 	100: 0.3,
	// 	125: -0.1,
	// 	150: 0.2,
	// 	175: 0.0,
	// }

	stats, err := tg.Generate(args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Save the image
	pathsImage, err := tg.GeneratePathsImage()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	outThreadFile, err := os.Create(outputFolder + exportPrefix + "result_preview.jpg")
	if err != nil {
		panic(err)
	}
	defer outThreadFile.Close()
	err = jpeg.Encode(outThreadFile, pathsImage, nil)
	if err != nil {
		panic(err)
	}

	// Save the thread list
	pathsList := tg.GetPathsList()
	outThreadListFile, err := os.Create(outputFolder + exportPrefix + "result_threads_list.txt")
	if err != nil {
		panic(err)
	}
	defer outThreadListFile.Close()

	for _, line := range pathsList {
		_, err := outThreadListFile.WriteString(fmt.Sprintf("%v	%v\n", line.StartingNail, line.EndingNail))
		if err != nil {
			panic(err)
		}
	}

	// Save the thread list
	gCodeLines := tg.GetGcode()
	outGcodeFile, err := os.Create(outputFolder + exportPrefix + "result_gcode.gcode")
	if err != nil {
		panic(err)
	}
	defer outGcodeFile.Close()

	for _, line := range gCodeLines {
		_, err := outGcodeFile.WriteString(line + "\n")
		if err != nil {
			panic(err)
		}
	}
	outHolesGcodeFile, err := os.Create(outputFolder + exportPrefix + "holes_gcode.gcode")
	if err != nil {
		panic(err)
	}
	defer outHolesGcodeFile.Close()
	for _, line := range tg.GenerateHolesGcode() {
		_, err := outHolesGcodeFile.WriteString(line + "\n")
		if err != nil {
			panic(err)
		}
	}

	// Save the thread list
	calibrationGcodeLines := tg.GenerateCalibrationGcode()
	outCalibrationGcodeFile, err := os.Create(outputFolder + exportPrefix + "calibration_gcode.gcode")
	if err != nil {
		panic(err)
	}
	defer outCalibrationGcodeFile.Close()

	for _, line := range calibrationGcodeLines {
		_, err := outCalibrationGcodeFile.WriteString(line + "\n")
		if err != nil {
			panic(err)
		}
	}

	// Save the config in a file in json
	config := tg.GetConfig()
	outConfigFile, err := os.Create(outputFolder + exportPrefix + "config.json")
	if err != nil {
		panic(err)
	}
	defer outConfigFile.Close()
	configJson, err := json.Marshal(config)
	if err != nil {
		panic(err)
	}
	_, err = outConfigFile.WriteString(string(configJson))
	if err != nil {
		panic(err)
	}

	fmt.Println("Done in", stats.TotalTime)
	fmt.Println("Number of lines", stats.TotalLines)
	fmt.Println(fmt.Sprintf("Estimated thread length: %d meters", stats.ThreadLength))
}
