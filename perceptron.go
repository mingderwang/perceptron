package main

import (
         "encoding/csv"
         "fmt"
         "os"
         "strconv"
 )

type Sample struct {
         Features [3]float64
         Target  int
 }

func main() {

         csvfile, err := os.Open("iris.csv")
	 if err != nil { 
		fmt.Println(err) 
		os.Exit(1) 
	}
	 defer csvfile.Close()

         reader := csv.NewReader(csvfile)

         reader.FieldsPerRecord = -1

         rawCSVdata, err := reader.ReadAll()

         if err != nil {
                 fmt.Println(err)
                 os.Exit(1)
         }

         // sanity check, display to standard output
         for _, each := range rawCSVdata {
                 fmt.Printf("%s %s %s -> %s\n", each[0], each[1], each[2], each[3] )
         }

         // now, safe to move raw CSV data to struct

         var oneRecord Sample

         var allRecords []Sample

         for _, each := range rawCSVdata {
	f1, _ := strconv.ParseFloat(each[0], 64)
	f2, _ := strconv.ParseFloat(each[1], 64)
	f3, _ := strconv.ParseFloat(each[2], 64)
	//f4, _ := strconv.ParseFloat(each[3], 64) 
	oneRecord.Features = [3]float64{f1,f2,f3} 
	if each[3]=="Iris-versicolor" {
                 oneRecord.Target = 1
	} else { 
		oneRecord.Target = 0 
	}
                 allRecords = append(allRecords, oneRecord)
         }


	// now, test weights using perceptron learning algorithm
	
	var w = [3]float64{0,0,0}
	eta := 0.1
	threshold := 0.5
	var result int
	
	for { 
	error := 0;
	for _, each := range allRecords {

             y := each.Features[0]*w[0]+each.Features[1]*w[1]+each.Features[2]*w[2]//+each.Features[3]*w[3]

	     if y >= threshold {
		result = 1	
	} else {
		result = 0 
	}

	tt := eta * float64(each.Target - result)
	fmt.Println(tt);
	if tt == 0 {
	} else {
		error = 1;
		w[0] = w[0] + tt*each.Features[0]
		w[1] = w[1] + tt*each.Features[1]
		w[2] = w[2] + tt*each.Features[2]
		//w[3] = w[3] + tt*each.Features[3]

		fmt.Println(w)
	} 
	}
	if error == 0 {
		break
	}

	}	

}
