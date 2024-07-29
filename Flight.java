public class Flight {
    private String flightNumber;
    private String status;
    private String gate;

    // Constructors, getters, and setters
    public Flight() {}

    public Flight(String flightNumber, String status, String gate) {
        this.flightNumber = flightNumber;
        this.status = status;
        this.gate = gate;
    }

    public String getFlightNumber() {
        return flightNumber;
    }

    public void setFlightNumber(String flightNumber) {
        this.flightNumber = flightNumber;
    }

    public String getStatus() {
        return status;
    }

    public void setStatus(String status) {
        this.status = status;
    }

    public String getGate() {
        return gate;
    }

    public void setGate(String gate) {
        this.gate = gate;
    }
}
