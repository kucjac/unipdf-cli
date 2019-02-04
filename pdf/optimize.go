/*
 * This file is subject to the terms and conditions defined in
 * file 'LICENSE.md', which is part of this source code package.
 */

package pdf

import (
	unipdf "github.com/unidoc/unidoc/pdf/model"
	unioptimize "github.com/unidoc/unidoc/pdf/model/optimize"
)

// OptimizeOpts represents the options used for optimizing PDF files.
type OptimizeOpts struct {
	// ImageQuality specifies the quality of the optimized images.
	ImageQuality int

	// ImagePPI specifies the maximum pixels per inch of the optimized images.
	ImagePPI float64
}

// Optimize optimizes the PDF file specified by the inputPath parameter, using
// the provided options and saves the result at the location specified by the
// outputPath parameter. A password can be specified for encrypted input files.
func Optimize(inputPath, outputPath, password string, opts *OptimizeOpts) error {
	// Read input file.
	r, _, _, _, err := readPDF(inputPath, password)
	if err != nil {
		return err
	}

	// Copy input file contents to the output file.
	w := unipdf.NewPdfWriter()
	if err = readerToWriter(r, &w, nil); err != nil {
		return err
	}

	// Add optimizer.
	if opts == nil {
		opts = &OptimizeOpts{
			ImageQuality: 100,
		}
	}

	w.SetOptimizer(unioptimize.New(unioptimize.Options{
		CombineDuplicateDirectObjects:   true,
		CombineIdenticalIndirectObjects: true,
		CombineDuplicateStreams:         true,
		CompressStreams:                 true,
		UseObjectStreams:                true,
		ImageQuality:                    opts.ImageQuality,
		ImageUpperPPI:                   opts.ImagePPI,
	}))

	// Write output file.
	safe := inputPath == outputPath
	return writePDF(outputPath, &w, safe)
}
