package utilities

func NetPrice(brand_price, ram_price float64, is_dvd bool) float64 {
	if is_dvd {
		return ram_price + brand_price + 3000
	}
	return ram_price + brand_price
}
