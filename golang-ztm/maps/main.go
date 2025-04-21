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

func servioreschange(servers []string) {
	fmt.Println("Total servidores", len(servers))

	stats := make(map[string]int)
	for _, server := range servers {
		switch stats {
		case online:
			stats[online] += 1
		case "aiur":
			stats[aiur] += 1
		case "omicron":
			stats[omicron] += 1
		case "w3375":
			stats[w3375] += 1
		case "centauri":		
			stats[centauri] += 1
		default:
			fmt.Println("Servidor desconhecido:", server)


	}

	}

}

func main() {
	// Exercicios

	const (
		online     = 0
		offilne    = 1
		manutencao = 2
		aposentado = 3
	)

	// Lista de servidores
	var servers = []string{"darkstar", "aiur", "omicron", "w3375", "centauri"}

	servioreschange(servers)
	serversStatus := make(map[string]int)
	totalServers := len(servers)
	fmt.Println("Total de servidores:", totalServers)
	// Adiciona o status de cada servidor no mapa
	for _, server := range servers {
		serversStatus[server] = online
	}

}
