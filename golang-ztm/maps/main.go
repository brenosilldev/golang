package main

import "fmt"

// func main() {
//Maps sao chaves e valores

// shoopingList := make(map[string]int)

// shoopingList["banana"] = 10
// shoopingList["apple"] = 5
// shoopingList["orange"] = 3
// fmt.Println(shoopingList)

// shoopingList["banana"] += 20
// fmt.Println(shoopingList)

// delete(shoopingList, "apple")
// fmt.Println(shoopingList)

// fmt.Println(shoopingList["banana"])

// orange, found := shoopingList["orange"]
// if found {
// 	fmt.Println("Orange found:", orange)
// } else {
// 	fmt.Println("Orange not found")
// }

// totalitems := 0

// for _, quantity := range shoopingList {
// 	totalitems += quantity
// }

// fmt.Println("Total items:", totalitems)

//Exercicios

// const (
// 	online     = 0
// 	offilne    = 1
// 	manutencao = 2
// 	aposentado = 3
// )

// // Lista de servidores
// var servers = []string{"darkstar", "aiur", "omicron", "w3375", "centauri"}

// func servioreschange(servers []string) {
// 	for _, server := range servers {

// 	}

// 	servioreschange(servers)
// serversStatus := make(map[string]int)
// totalServers := len(servers)
// fmt.Println("Total de servidores:", totalServers)
// // Adiciona o status de cada servidor no mapa
// for _, server := range servers {
// 	serversStatus[server] = online
// }

// }
const (
	online     = 0
	offline    = 1 // corrigido typo
	manutencao = 2
	aposentado = 3
)

func servioreschange(statusMap map[string]int) {
	total := len(statusMap)
	fmt.Println("Total de servidores:", total)

	stats := make(map[int]int)
	for server, status := range statusMap {
		switch status {
		case online:
			stats[online]++
		case offline:
			stats[offline]++
		case manutencao:
			stats[manutencao]++
		case aposentado:
			stats[aposentado]++
		default:
			fmt.Println("Servidor desconhecido:", server)
		}
	}

	fmt.Println(stats[online], "servidores online")
	fmt.Println(stats[offline], "servidores offline")
	fmt.Println(stats[manutencao], "servidores em manutenção")
	fmt.Println(stats[aposentado], "servidores aposentados")
}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w3375", "centauri", "baseline"}
	serversStatus := make(map[string]int)

	// Inicializa todos como online
	for _, server := range servers {
		serversStatus[server] = manutencao
	}

	// Ajusta alguns status
	serversStatus["darkstar"] = aposentado
	serversStatus["aiur"] = offline

	// Exibe estatísticas
	servioreschange(serversStatus)
	fmt.Println("------------------")

	serversStatus["w3375"] = online
	serversStatus["aiur"] = offline
	serversStatus["baseline"] = offline
	servioreschange(serversStatus)
}
