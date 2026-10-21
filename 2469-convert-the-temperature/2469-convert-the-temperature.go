func convertTemperature(celsius float64) []float64 {
    res := make([]float64, 2) 
    res[0] = celsius + 273.15
    res[1] = celsius * 1.80 + 32.00
    return res
}