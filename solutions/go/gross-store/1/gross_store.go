package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
    mp := make(map[string]int)
    mp["quarter_of_a_dozen"] = 3
    mp["half_of_a_dozen"] = 6
    mp["dozen"] = 12
    mp["small_gross"] = 120
    mp["gross"] = 144
    mp["great_gross"] = 1728
	return mp 
}

// NewBill creates a new bill.
func NewBill() map[string]int {
    bill := make(map[string]int)
    return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	add, exists := units[unit]
    if exists {
        bill[item] += add
        return true
    }
    return false
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, existsInBill := bill[item]
    itemInUnit, existInUnits := units[unit]

    if !existsInBill || !existInUnits {
        return false
    }

    if bill[item] - itemInUnit < 0 {
        return false
    } else if bill[item] - itemInUnit == 0 {
        delete(bill, item)
        return true
    } else {
        bill[item] -= itemInUnit
        return true
    }
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	value, exists := bill[item]
    return value, exists
}
