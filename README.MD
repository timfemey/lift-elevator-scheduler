# Elevator Lift Scheduler System

This project is a Lift Scheduler System designed for safety, implemented in GoLang. It efficiently manages elevator operations, ensuring smooth transitions between floors while preventing duplicate and unordered requests.

## Features

- **Start Lift:** Initiate the elevator system to begin operation.
- **Get Current Floor:** Retrieve the current floor the elevator is stationed at.
- **Get Current Queue:** Obtain the ordered set of pending requests in the elevator's queue.
- **Get Current Direction:** Determine the direction the elevator is currently moving (up, down, stationary).
- **Display Current Floor and Awaiting Next Request:** Visualize the current floor and the next request / queue in the system.

## Safety Measures

- **Mutex:** For Safe concurrent Operations to avoid race conditions and app crashes
- **Ordered Set:** Utilizes an ordered set data structure to prevent duplicate and unordered requests, ensuring efficient and safe operation of the elevator.

## Usage

```go
package main

import (
	"fmt"
	"github.com/timfemey/lift-scheduler"
)

func main() {
	// Initialize the lift scheduler
	scheduler := lift.Lift()

	// Start the lift system
	scheduler.StartLift()

	// Get current floor
	currentFloor := scheduler.GetCurrentFloor()

	// Get current queue
	currentQueue := scheduler.GetCurrentQueue()

	// Get current direction
	currentDirection := scheduler.GetCurrentDirection()

	// Display current floor and awaiting next request
	scheduler.DisplayCurrentFloor()

	fmt.Println("Current Floor:", currentFloor)
	fmt.Println("Current Queue:", currentQueue)
	fmt.Println("Current Direction:", currentDirection)
}
```

## Installation

1. Install the lift package:

   ```bash
   go get -u github.com/timfemey/lift-scheduler
   ```

2. Import the package into your code:

   ```go
   import "github.com/timfemey/lift-scheduler"
   ```

3. Use the package as demonstrated in the example above.

## Contributions

Contributions are welcome! If you find any issues or have suggestions for improvements, please feel free to open an issue or create a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
