package complete

import "sync"

type MultiworkerCompletion struct {
	clients []*Completion
	wg      sync.WaitGroup
}

func NewMultiworkerCompletion(workersNum int, token string, baseUrl string) *MultiworkerCompletion {
	clients := make([]*Completion, 0, workersNum)

	for i := 0; i < workersNum; i++ {
		clients = append(clients, NewCompletion(token, baseUrl))
	}

	return &MultiworkerCompletion{clients: clients}
}

func (m *MultiworkerCompletion) worker(idx int, model string, input <-chan string, output chan<- string) {
	client := m.clients[idx]
	m.wg.Add(1)
	defer m.wg.Done()

	for prompt := range input {
		resp := client.Completion(prompt, model)
		output <- resp
	}
}

func (m *MultiworkerCompletion) Listen(model string, input <-chan string, output chan<- string) {
	for i := range m.clients {
		m.worker(i, model, input, output)
	}

	m.wg.Wait()
	close(output)
}
