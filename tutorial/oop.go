package tutorial

import "fmt"

// Equipment test
type equipment struct {
	name  string
	price int
}

// Weapon test
type weapon struct {
	attack int
	equipment
}

// WeaponBehavior -- Polymorphism
type WeaponBehavior interface {
	doAttack()
}

// toString ...
func (equipment *equipment) toString() {
	fmt.Printf("Equipment{name = %s, price = %d}\n", equipment.name, equipment.price)
}

// toString overload
func (equipment *weapon) toString() {
	fmt.Printf("Weapon{name = %s, price = %d, attack = %d}\n", equipment.name, equipment.price, equipment.attack)
}

// doAttack ...
func (weapon weapon) doAttack() {
	fmt.Printf("Deal %d damage", weapon.attack)
}

// TestOOP ...
func TestOOP() {
	wpn := weapon{
		attack: 23,
		equipment: equipment{
			name:  "fire rod",
			price: 50,
		},
	}
	wpn.toString()
	var wb WeaponBehavior = wpn
	wb.doAttack()
}
