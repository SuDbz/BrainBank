package main

import (
        "context"
        "fmt"
        "time"

        etcd3 "go.etcd.io/etcd/client/v3"
)

func main() {
        client, err := etcd3.NewClient(etcd3.Config{
                Endpoints:   []string{"http://localhost:2379"},
                DialTimeout: 5 * time.Second,
        })
        if err != nil {
                panic(err)
        }
        defer client.Close()

        lockKey := "/my-lock"

        tryAcquireLock(client, lockKey)
}

func tryAcquireLock(client *etcd3.Client, lockKey string) {
        ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
        defer cancel()

        resp, err := client.Txn(ctx).
                If(etcd3.CmpValue("lock", "=", "")).
                Then(etcd3.OpPut(lockKey, "lock")).
                Else(etcd3.OpGet(lockKey)).
                Commit()
        if err != nil {
                fmt.Println("Error acquiring lock:", err)
                return
        }

        if resp.Header.Revision != 0 {
                // Lock acquired
                fmt.Println("Lock acquired, performing critical section...")
                // Do your work here

                // Release the lock
                _, err = client.Delete(ctx, lockKey)
                if err != nil {
                        fmt.Println("Error releasing lock:", err)
                }
        } else {
                // Lock not acquired
                fmt.Println("Failed to acquire lock, retrying later...")
                time.Sleep(1 * time.Second)
                tryAcquireLock(client, lockKey)
        }
}