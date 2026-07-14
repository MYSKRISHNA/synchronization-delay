type Request struct {
	id int
	data string
}

type Replica struct {
	id    int
	state string
	delay int
}

func processRequest(req Request) string {
	time.Sleep(time.Millisecond * 50)
	return "Processed-" + req.data
}

func synchronize(replica *Replica, state string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Millisecond * time.Duration(replica.delay))
	replica.state = state
	fmt.Println("Replica", replica.id, "Synced")
}

func monitor(replicas []Replica) {
	for _, r := range replicas {
		fmt.Println("Replica", r.id, "Delay", r.delay, "ms")
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	replicas := []Replica{
		{id: 1, delay: 100},
		{id: 2, delay: 180},
		{id: 3, delay: 250},
		{id: 4, delay: 320},
	}

	requests := []Request{
		{id: 1, data: "A"},
		{id: 2, data: "B"},
		{id: 3, data: "C"},
		{id: 4, data: "D"},
	}

	for _, req := range requests {

		fmt.Println("Client Request", req.id)

		primaryState := processRequest(req)

		fmt.Println("Primary Updated", primaryState)

		var wg sync.WaitGroup

		for i := range replicas {
			wg.Add(1)
			go synchronize(&replicas[i], primaryState, &wg)
		}

		wg.Wait()

		monitor(replicas)

		fmt.Println("---------------------")
	}

	fmt.Println("Synchronization Complete")
}
