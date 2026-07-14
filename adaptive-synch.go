type Replica struct {
	id    int
	delay int
	state string
	health string
}

func processRequest(id int) string {
	time.Sleep(time.Millisecond * 40)
	return fmt.Sprintf("Update-%d", id)
}

func adaptiveDecision(delay int) int {
	if delay > 300 {
		return delay - 120
	}
	return delay - 40
}

func synchronize(r *Replica, data string, wg *sync.WaitGroup) {
	defer wg.Done()

	adjusted := adaptiveDecision(r.delay)

	time.Sleep(time.Millisecond * time.Duration(adjusted))

	r.state = data
	r.health = "Active"

	fmt.Println("Replica", r.id,
		"Delay", adjusted,
		"State", r.state)
}

func monitor(replicas []Replica) {
	for _, r := range replicas {
		fmt.Println("Node", r.id,
			"Health", r.health,
			"BaseDelay", r.delay)
	}
}

func main() {

	rand.Seed(time.Now().UnixNano())

	replicas := []Replica{
		{1, 420, "", ""},
		{2, 560, "", ""},
		{3, 710, "", ""},
		{4, 860, "", ""},
		{5, 1020, "", ""},
	}

	for request := 1; request <= 5; request++ {

		fmt.Println("Request", request)

		update := processRequest(request)

		fmt.Println("Processed", update)

		var wg sync.WaitGroup

		for i := range replicas {
			wg.Add(1)
			go synchronize(&replicas[i], update, &wg)
		}

		wg.Wait()

		monitor(replicas)

		fmt.Println("----------------------")
	}

	fmt.Println("Adaptive Synchronization Complete")
}
