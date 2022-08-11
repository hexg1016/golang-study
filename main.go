package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

var count int64

type Header struct {
	len   int32
	index byte
}

func main() {
	data := `{"Key":"2", "3": "4"}`
	var v interface{}
	if err := json.Unmarshal([]byte(data), &v); err != nil {
		log.Fatal(err)
	}

	v1, ok := v.(map[string]interface{})
	if ok {
		keyValue, ok := v1["Key"]
		if ok {
			fmt.Println(reflect.TypeOf(keyValue), keyValue)
		}
	}
	data1 := `{"Key":2, "3": "4"}`
	var v2 interface{}
	if err := json.Unmarshal([]byte(data1), &v2); err != nil {
		log.Fatal(err)
	}

	v3, ok := v2.(map[string]interface{})
	if ok {
		keyValue, ok := v3["Key"]
		if ok {
			fmt.Println(reflect.TypeOf(keyValue), keyValue)
		}
	}
	/*config := api.DefaultConfig()
	config.Address = "127.0.0.1:8501"
	config.Token = "e8868b1b-92a7-6937-8ca8-4365dbbefe50"
	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}

	// Get a handle to the KV API
	kv := client.KV()

	// PUT a new KV pair
	p := &api.KVPair{Key: "REDIS_MAXCLIENTS", Value: []byte("1000")}
	_, err = kv.Put(p, nil)
	if err != nil {
		panic(err)
	}

	// Lookup the pair
	pair, _, err := kv.Get("REDIS_MAXCLIENTS", nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("KV: %v %s\n", pair.Key, pair.Value)

	addr := fmt.Sprintf("%s:9091", localIP())
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Errorf("listen tcp %s failed!", addr)
	}
	defer lis.Close()
	go tcpServer(lis)
	time.Sleep(time.Millisecond * 100)

	opts := []Option{
		WithHeartbeat(true),
		WithHealthCheck(true),
		WithHealthCheckInterval(5),
	}
	r := New(client, opts...)

	version := strconv.FormatInt(time.Now().Unix(), 10)
	svc := &ServiceInstance{
		ID:        "test2233",
		Name:      "test-provider",
		Version:   version,
		Metadata:  map[string]string{"app": "kratos"},
		Endpoints: []string{fmt.Sprintf("tcp://%s?isSecure=false", addr)},
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	err = r.Deregister(ctx, svc)

	err = r.Register(ctx, svc)

	w, err := r.Watch(ctx, "test-provider")

	services, err := w.Next()
	fmt.Println(services)
	registerServer(client)
	/*client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance) // 25893180161173005034

	blockNumber := big.NewInt(5532993)
	balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balanceAt) // 25729324269165216042

	fbalance := new(big.Float)
	fbalance.SetString(balanceAt.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue) // 25.729324269165216041

	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	fmt.Println(pendingBalance) // 25729324269165216042*/
}
