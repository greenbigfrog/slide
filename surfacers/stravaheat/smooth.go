package stravaheat

import (
	"slide/utils"
	"slide/utils/smoothsurface"

	geo "github.com/paulmach/go.geo"
)

// Resmooth takes the data pulled in from the tiles and applies a new smoothing
// to it based on a potentially updated `SmoothingStdDev`.
// Basically it clears and resets the LazySmoothSurface.
func (surfacer *Surface) Resmooth() error {
	if surfacer.SmoothSurface == nil {
		return surfacer.smooth()
	}

	surfacer.SmoothSurface.SetKernel(surfacer.smoothKernel())
	return nil
}

// smooth sets up the LazySmoothSurface with a kernel.
func (surfacer *Surface) smooth() error {
	surfacer.SmoothSurface = smoothsurface.New(surfacer.Surface, surfacer.smoothKernel())
	return nil
}

// smoothKernel creates the smoothing kernel that is based on `SmoothingStdDev` (meters).
// See the function utils.Kernel for more information.
func (surfacer *Surface) smoothKernel() []float64 {
	return utils.Kernel(
		surfacer.SmoothingStdDev,
		geo.MercatorScaleFactor(surfacer.lnglatBound.Center().Lat()),
	)
}
