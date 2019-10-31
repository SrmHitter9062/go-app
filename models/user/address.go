package user


type Address struct{
	House string `db:house_address json:HouseAddress`
	City string `db:city json:city`
	Country string `db:country json:country`
	PinCode uint32 `db:pin_code json:pinCode`
}
