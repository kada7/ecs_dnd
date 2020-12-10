package ecs_dnd

import "github.com/EngoEngine/ecs"

type MoveSystem struct {
}

func (s *MoveSystem) Update(dt float32) {
	panic("implement me")
}

func (s *MoveSystem) Remove(e ecs.BasicEntity) {
	panic("implement me")
}

func (s *MoveSystem) Add(basic *ecs.BasicEntity, m *MoveComponent) {

}

type HealthSystem struct {
}

func (s *HealthSystem) Update(dt float32) {
	panic("implement me")
}

func (s *HealthSystem) Remove(e ecs.BasicEntity) {
	panic("implement me")
}

func (s *HealthSystem) Add(basic *ecs.BasicEntity, m *HealthComponent) {

}

func main() {
	world := &ecs.World{}
	world.AddSystem(&MoveSystem{})
	world.AddSystem(&HealthSystem{})

	ch := CharacterEntity{BasicEntity: ecs.NewBasic()}

	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *MoveSystem:
			sys.Add(&ch.BasicEntity, &MoveComponent{})
		case *HealthSystem:
			sys.Add(&ch.BasicEntity, &HealthComponent{})
		}
	}
}

type CharacterEntity struct {
	ecs.BasicEntity
	CreatureComponent
	HealthComponent
	MoveComponent
	AbilityComponent
}

type CreatureComponent struct {
	Name      string
	Gender    int
	Alignment string
}

type HealthComponent struct {
	HP    int
	HPMax int
}

type MoveComponent struct {
	Speed int
}

type AbilityComponent struct {
	Abilities int
}
