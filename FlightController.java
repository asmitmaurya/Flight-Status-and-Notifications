import org.springframework.web.bind.annotation.*;

import java.util.ArrayList;
import java.util.List;

@RestController
@RequestMapping("/api")
public class FlightController {

    // Mock data for flights (replace with actual data fetching from airport system)
    private List<Flight> flights = new ArrayList<>();

    public FlightController() {
        // Initialize mock data
        flights.add(new Flight("ABC123", "On time", "A1"));
        flights.add(new Flight("XYZ789", "On time", "B2"));
    }

    // Endpoint to get all flights
    @GetMapping("/flights")
    public List<Flight> getFlights() {
        return flights;
    }

    // Endpoint to update flight status
    @PutMapping("/update_status/{flightNumber}")
    public Flight updateStatus(@PathVariable String flightNumber, @RequestBody Flight updatedFlight) {
        for (Flight flight : flights) {
            if (flight.getFlightNumber().equals(flightNumber)) {
                flight.setStatus(updatedFlight.getStatus());
                // Add notification logic here if needed
                return flight;
            }
        }
        throw new RuntimeException("Flight not found: " + flightNumber);
    }
}
