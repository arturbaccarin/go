package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	} else {
		return false
	}
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {
	var stringExit string
	if option1 < option2 {
		stringExit = option1 + " is clearly the better choice."
	} else {
		stringExit = option2 + " is clearly the better choice."
	}
	return stringExit
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var finalPrice float64 = 0
	if age < 3.0 {
		finalPrice = originalPrice * 0.8
	} else if 10.0 <= age {
		finalPrice = originalPrice * 0.5
	} else {
		finalPrice = originalPrice * 0.7
	}
	return finalPrice
}
