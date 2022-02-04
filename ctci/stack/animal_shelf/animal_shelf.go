package animalshelf

type Animal struct {
	Order []int
}

type AnimalShelf struct {
	Animal [2]Animal
}

func Constructor() AnimalShelf {
	return AnimalShelf{}
}

func (this *AnimalShelf) Enqueue(animal []int) {
	this.Animal[animal[1]].Order = append(this.Animal[animal[1]].Order, animal[0])
}

func (this *AnimalShelf) DequeueAny() []int {
	cat := this.peek(0)
	dog := this.peek(1)

	if cat[0] == -1 {
		if dog[0] == -1 {
			return cat
		} else {
			this.pop(1)
			return dog
		}
	} else {
		if dog[0] == -1 || cat[0] < dog[0] {
			this.pop(0)
			return cat
		} else {
			this.pop(1)
			return dog
		}
	}
}

func (this *AnimalShelf) DequeueDog() []int {
	return this.dequeue(1)
}

func (this *AnimalShelf) DequeueCat() []int {
	return this.dequeue(0)
}

func (a *AnimalShelf) dequeue(animalType int) []int {
	res := a.peek(animalType)
	if res[0] != -1 {
		a.pop(animalType)
	}
	return res
}

func (a *AnimalShelf) pop(animalType int) {
	a.Animal[animalType].Order = a.Animal[animalType].Order[1:]
}

func (a *AnimalShelf) peek(animalType int) []int {
	n := len(a.Animal[animalType].Order)
	if n == 0 {
		return []int{-1, -1}
	} else {
		res := []int{a.Animal[animalType].Order[0], animalType}
		return res
	}
}

/**
 * Your AnimalShelf object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Enqueue(animal);
 * param_2 := obj.DequeueAny();
 * param_3 := obj.DequeueDog();
 * param_4 := obj.DequeueCat();
 */
