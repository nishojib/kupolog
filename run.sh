#!/usr/bin/env zsh

# Navigate to the frontend directory and start the React app
echo "Starting React app..."
cd frontend
pnpm run dev &

# Capture the PID of the React process
REACT_PID=$!

# Navigate to the backend directory and start the Golang app
echo "Starting Golang app..."
cd ../backend
air &

# Capture the PID of the Golang process
GO_PID=$!

# Function to stop both processes when the script is terminated
cleanup() {
  echo "Stopping React app..."
  kill $REACT_PID
  echo "Stopping Golang app..."
  kill $GO_PID
  exit 0
}

# Trap script termination signals and call cleanup
trap cleanup SIGINT SIGTERM

# Wait for both processes to complete
wait $REACT_PID
wait $GO_PID