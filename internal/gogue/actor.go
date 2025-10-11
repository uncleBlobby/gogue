package gogue

import (
	"math/rand/v2"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Actor struct {
	Stats    Stats
	Collider rl.Rectangle
}

type AttackResult int

const (
	HIT AttackResult = iota
	CRIT
	MISS
)

type DefendResult int

const (
	NIL DefendResult = iota
	DEFEND
	DODGE
)

func (a *Actor) RollLuck() float32 {
	return rand.Float32() + float32(a.Stats.Luck.Current/100)
}

func (a *Actor) Attack(target *Actor) AttackResult {
	luckRoll := a.RollLuck()
	// if luckroll < 0.1, miss the target
	if luckRoll < 0.1 {
		return AttackResult(MISS)
	}
	// most of the times, attack is just a regular hit
	// doing damage to the target HP equal to the attack value
	//  may be mitigated by defense
	if luckRoll >= 0.1 && luckRoll < 0.9 {
		_ = target.Defend(a.Stats.Attack.Current)
		return AttackResult(HIT)
	}
	// rare crit roll for 1.5x damage
	if luckRoll >= 0.9 {
		_ = target.Defend(int(float32(a.Stats.Attack.Current) * 1.5))
		return AttackResult(CRIT)
	}
	return AttackResult(MISS)
}

func (a *Actor) TakeDamage(finalDamage int) {
	a.Stats.Hitpoints.Current -= finalDamage
}

func (a *Actor) Defend(attackDamage int) DefendResult {
	luckRoll := a.RollLuck()

	if luckRoll < 0.1 {
		a.TakeDamage(attackDamage)
		return DefendResult(NIL)
	}
	// most of the times, attack is just a regular hit
	// doing damage to the target HP equal to the attack value
	if luckRoll >= 0.1 && luckRoll < 0.9 {
		a.TakeDamage(attackDamage - a.Stats.Defense.Current)
		return DefendResult(DEFEND)
	}
	// rare crit roll to dodge all damage
	if luckRoll >= 0.9 {
		a.TakeDamage(0)
		return DefendResult(DODGE)
	}
	return DefendResult(NIL)
}

type Stats struct {
	Hitpoints Hitpoints
	Attack    Attack
	Defense   Defense
	Luck      Stat
	Dodge     Stat
}

func InitBaseCollider(xPos, yPos, width, height float32) rl.Rectangle {
	return rl.Rectangle{X: xPos, Y: yPos, Width: width, Height: height}
}

func InitBaseStats(
	hp int,
	atk int,
	dfs int,
	lk int,
	dg int,
) Stats {
	return Stats{
		Hitpoints: Hitpoints{
			Current: hp,
			Max:     hp,
		},
		Attack: Attack{
			Current: atk,
		},
		Defense: Defense{
			Current: dfs,
		},
		Luck: Stat{
			Current: lk,
		},
		Dodge: Stat{
			Current: dg,
		},
	}
}

type Hitpoints struct {
	Current int
	Max     int
}

type Attack struct {
	Current int
	// Max     int
}

type Defense struct {
	Current int
}

type Stat struct {
	Current int
	Max     int
}

func (a *Attack) Attack(target *Stats) AttackResult {
	// do damage to target stat body

	// roll luck for crit chance, miss chance
	// (for now we just mimic a roll with a random float)

	luckRoll := rand.Float32()

	// if luckroll < 0.1, miss the target
	if luckRoll < 0.1 {
		return AttackResult(MISS)
	}

	// most of the times, attack is just a regular hit
	// doing damage to the target HP equal to the attack value
	if luckRoll >= 0.1 && luckRoll < 0.9 {
		target.Hitpoints.Current -= a.Current
		return AttackResult(HIT)
	}

	// rare crit roll for 1.5x damage
	if luckRoll >= 0.9 {
		target.Hitpoints.Current -= int(float32(a.Current) * 1.5)
		return AttackResult(CRIT)
	}

	return AttackResult(MISS)
}

func (a *Actor) CollidingOtherActor(targ *Actor) bool {

	return rl.CheckCollisionRecs(a.Collider, targ.Collider)
}

// func (a *Actor) UpdateCollider(dt float32) {
// 	a.Collider.X =
// }
