//  ___   __   __    __     __   _  _   __   _  _    ____  _  _  ____  ____  __   __ _       ____   __  ____  ____ 
// / __) / _\ (  )  (  )   /  \ / )( \ / _\ ( \/ )  / ___)/ )( \(_  _)(_  _)/  \ (  ( \ _   (___ \ /  \(___ \( __ \
// ( (__ /    \/ (_/\/ (_/\(  O )\ /\ //    \ )  /   \___ \) \/ (  )(    )( (  O )/    /( )   / __/(  0 )/ __/ (__ (
// \___)\_/\_/\____/\____/ \__/ (_/\_)\_/\_/(__/    (____/\____/ (__)  (__) \__/ \_)__)(/   (____) \__/(____)(____/

//  SOFTWARE WARRANTY LICENSE NOTICE
//  (c) 2023 Calloway Sutton. All rights reserved.

//  This software is provided "as is" and without any express or
//  implied warranties, including, but not limited to, the implied
//  warranties of merchantability and fitness for a particular
//  purpose. In no event shall the authors or copyright holders
//  be liable for any claim, damages, or other liability, whether
//  in an action of contract, tort, or otherwise, arising from,
//  out of, or in connection with the software or the use or other
//  dealings in the software.

//  For more information, please contact: me@callowaysutton.com

package main

import (
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/dchest/uniuri"
	"github.com/jung-kurt/gofpdf"
	"github.com/schollz/progressbar/v3"
)

const (
	NumPDFs      = 5000
	ProcessCount = 32 // Number of goroutines to use
)

func generatePDF(i int, wg *sync.WaitGroup, bar *progressbar.ProgressBar, sem chan bool) {
	defer wg.Done()

	name := uniuri.NewLen(8)
	filename := "pdfs/" + name + "_" + strconv.Itoa(i) + ".pdf"
	pdf := gofpdf.New("P", "mm", "A4", "")

	rand.Seed(time.Now().UnixNano())
	numSentences := rand.Intn(101) + 50 // generates number between 50 and 150

	for j := 0; j < numSentences; j++ {
		sentence := uniuri.NewLen(150) // You might want to replace this with actual sentence generation
		pdf.SetFont("Helvetica", "", 12)
		pdf.Text(50, float64(200+(j*14)), sentence)
		pdf.SetFont("Helvetica", "", 10)
		pdf.AddPage()
	}

	err := pdf.OutputFileAndClose(filename)
	if err != nil {
		panic(err)
	}

	bar.Add(1) // increment the progress bar by 1
	<-sem      // release the semaphore
}

func main() {
	// Create the output directory if it doesn't exist
	_ = os.Mkdir("pdfs", 0755)

	var wg sync.WaitGroup

	// Create a progress bar
	bar := progressbar.Default(int64(NumPDFs))

	// Create a semaphore to limit number of concurrent goroutines
	sem := make(chan bool, ProcessCount)

	// Create a pool of goroutines
	for i := 0; i < NumPDFs; i++ {
		wg.Add(1)
		sem <- true
		go generatePDF(i, &wg, bar, sem)
	}

	wg.Wait() // Wait for all goroutines to finish
}
