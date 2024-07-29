import React, { useState, useEffect } from 'react';
import axios from 'axios';

const FlightList = () => {
    const [flights, setFlights] = useState([]);

    useEffect(() => {
        fetchFlights();
    }, []);

    const fetchFlights = async () => {
        try {
            const response = await axios.get('/api/flights');
            setFlights(response.data.flights);
        } catch (error) {
            console.error('Error fetching flights:', error);
        }
    };

    const handleStatusUpdate = async (flightNumber, newStatus) => {
        try {
            await axios.put(`/api/update_status/${flightNumber}`, { status: newStatus });
            // After updating status, fetch updated flights
            fetchFlights();
        } catch (error) {
            console.error('Error updating flight status:', error);
        }
    };

    return (
        <div>
            <h2>Flight Status Updates</h2>
            <ul>
                {flights.map(flight => (
                    <li key={flight.flight_number}>
                        <strong>Flight Number:</strong> {flight.flight_number} <br />
                        <strong>Status:</strong> {flight.status} <br />
                        <strong>Gate:</strong> {flight.gate} <br />
                        <button onClick={() => handleStatusUpdate(flight.flight_number, 'Delayed')}>
                            Delay Flight
                        </button>
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default FlightList;
