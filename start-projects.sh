#! /bin/bash

# INICIAR EL CLIENTE EXPRESS
(cd client && npm run dev) &

# INICIAR EL SERVIDOR FASTAPI
(cd server && python main.py) &

# INICIAR EL TEMPORAL
(cd temp && go run main.go) &

wait


