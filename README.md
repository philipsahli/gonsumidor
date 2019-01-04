Gonsumidor increments the counter by 1 and prints out the counter every 2 seconds. Communication over gRPC to the gontador's service on port `3001`.

# Build

    go build -o gonsumidor main.go

# Run

    SERVER_ADDRESS=localhost:3001 ./gonsumidor

# Undeploy

    oc delete all -l app=gonsumidor
    oc delete template gonsumidor-template