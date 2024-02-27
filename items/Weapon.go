package items

type Weapon struct {
	name        string
	attackPower int
}

func NewWeapon(name string, attackPower int) *Weapon {
	return &Weapon{name: name, attackPower: attackPower}
}

func (w Weapon) Name() string {
	return w.name
}

func (w Weapon) AttackPower() int {
	return w.attackPower
}
